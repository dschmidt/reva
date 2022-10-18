// Copyright 2018-2021 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package ocdav

import (
	"net/http"
	"strings"

	"github.com/cs3org/reva/v2/internal/http/interceptors/appctx"
	"github.com/cs3org/reva/v2/internal/http/interceptors/auth"
	revaLogMiddleware "github.com/cs3org/reva/v2/internal/http/interceptors/log"
	"github.com/cs3org/reva/v2/internal/http/services/owncloud/ocdav"
	"github.com/cs3org/reva/v2/pkg/rgrpc/todo/pool"
	"github.com/cs3org/reva/v2/pkg/rhttp/global"
	"github.com/cs3org/reva/v2/pkg/storage/favorite/memory"
	rtrace "github.com/cs3org/reva/v2/pkg/trace"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpServer "github.com/go-micro/plugins/v4/server/http"
	"github.com/owncloud/ocis/v2/ocis-pkg/registry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	// ServerName to use when announcing the service to the registry
	ServerName = "ocdav"
)

// Service initializes the ocdav service and underlying http server.
func Service(opts ...Option) (micro.Service, error) {

	sopts := newOptions(opts...)

	// set defaults
	if err := setDefaults(&sopts); err != nil {
		return nil, err
	}

	sopts.Logger = sopts.Logger.With().Str("name", sopts.Name).Logger()

	srv := httpServer.NewServer(
		server.Broker(sopts.Broker),
		server.TLSConfig(sopts.TLSConfig),
		server.Name(sopts.Name),
		server.Address(sopts.Address), // Address defaults to ":0" and will pick any free port
		server.Version(sopts.config.VersionString),
	)

	tp := rtrace.GetTracerProvider(sopts.TracingEnabled, sopts.TracingCollector, sopts.TracingEndpoint, sopts.Name)
	revaService, err := ocdav.NewWith(&sopts.config, sopts.FavoriteManager, sopts.lockSystem, &sopts.Logger, tp)
	if err != nil {
		return nil, err
	}

	// Comment back in after resolving the issue in go-chi.
	// See comment in line 87.
	// register additional webdav verbs
	// chi.RegisterMethod(ocdav.MethodPropfind)
	// chi.RegisterMethod(ocdav.MethodProppatch)
	// chi.RegisterMethod(ocdav.MethodLock)
	// chi.RegisterMethod(ocdav.MethodUnlock)
	// chi.RegisterMethod(ocdav.MethodCopy)
	// chi.RegisterMethod(ocdav.MethodMove)
	// chi.RegisterMethod(ocdav.MethodMkcol)
	// chi.RegisterMethod(ocdav.MethodReport)
	r := chi.NewRouter()

	if err := useMiddlewares(r, &sopts, revaService, tp); err != nil {
		return nil, err
	}

	r.Handle("/*", revaService.Handler())
	// This is a workaround for the go-chi concurrent map read write issue.
	// After the issue has been solved upstream in go-chi we should switch
	// back to using `chi.RegisterMethod`.
	r.MethodNotAllowed(http.HandlerFunc(revaService.Handler().ServeHTTP))

	hd := srv.NewHandler(r)
	if err := srv.Handle(hd); err != nil {
		return nil, err
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(registry.GetRegistry()),
	)

	// Init the service? make that optional?
	service.Init()

	// finally, return the service so it can be Run() by the caller himself
	return service, nil
}

func setDefaults(sopts *Options) error {

	// set defaults
	if sopts.Name == "" {
		sopts.Name = ServerName
	}
	if sopts.lockSystem == nil {
		client, err := pool.GetGatewayServiceClient(sopts.config.GatewaySvc)
		if err != nil {
			return err
		}
		sopts.lockSystem = ocdav.NewCS3LS(client)
	}
	if sopts.FavoriteManager == nil {
		sopts.FavoriteManager, _ = memory.New(map[string]interface{}{})
	}
	if !strings.HasPrefix(sopts.config.Prefix, "/") {
		sopts.config.Prefix = "/" + sopts.config.Prefix
	}
	if sopts.config.VersionString == "" {
		sopts.config.VersionString = "0.0.0"
	}
	return nil
}

func useMiddlewares(r *chi.Mux, sopts *Options, svc global.Service, tp trace.TracerProvider) error {
	// auth
	for _, v := range svc.Unprotected() {
		sopts.Logger.Info().Str("url", v).Msg("unprotected URL")
	}
	authMiddle, err := auth.New(map[string]interface{}{
		"gatewaysvc": sopts.config.GatewaySvc,
		"token_managers": map[string]interface{}{
			"jwt": map[string]interface{}{
				"secret": sopts.JWTSecret,
			},
		},
	}, svc.Unprotected(), tp)
	if err != nil {
		return err
	}

	// log
	lm := revaLogMiddleware.New()

	// tracing
	tm := func(h http.Handler) http.Handler { return h }
	if sopts.TracingEnabled {
		tm = traceHandler(tp, "ocdav")
	}

	// metrics
	pm := func(h http.Handler) http.Handler { return h }
	if sopts.MetricsEnabled {
		namespace := sopts.MetricsNamespace
		if namespace == "" {
			namespace = "reva"
		}
		subsystem := sopts.MetricsSubsystem
		if subsystem == "" {
			subsystem = "ocdav"
		}
		counter := promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "http_requests_total",
			Help:      "The total number of processed " + subsystem + " HTTP requests for " + namespace,
		})
		pm = func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				h.ServeHTTP(w, r)
				counter.Inc()
			})
		}
	}

	// ctx
	cm := appctx.New(sopts.Logger, tp)

	// request-id
	rm := middleware.RequestID

	// actually register
	r.Use(pm, tm, lm, authMiddle, rm, cm)
	return nil
}

func traceHandler(tp trace.TracerProvider, name string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := rtrace.Propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
			t := tp.Tracer("reva")
			ctx, span := t.Start(ctx, name)
			defer span.End()

			rtrace.Propagator.Inject(ctx, propagation.HeaderCarrier(r.Header))
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
