package restapi

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/roscopecoltran/gin-swagger/api"
	"github.com/roscopecoltran/gin-swagger/generated/searchkit-apis/restapi/operations"
	"github.com/roscopecoltran/gin-swagger/middleware"
	log "github.com/sirupsen/logrus"
)

// Routes defines all the routes of the API service.
type Routes struct {
	*gin.Engine
	SearchAPISpecs struct {
		*gin.RouterGroup
		Post *gin.RouterGroup
	}
}

// configureWellKnown enables and configures /.well-known endpoints.
func (r *Routes) configureWellKnown(healthFunc func() bool) {
	wellKnown := r.Group("/.well-known")
	{
		wellKnown.GET("/schema-discovery", func(ctx *gin.Context) {
			discovery := struct {
				SchemaURL  string `json:"schema_url"`
				SchemaType string `json:"schema_type"`
				UIURL      string `json:"ui_url"`
			}{
				SchemaURL:  "/swagger.json",
				SchemaType: "swagger-2.0",
			}
			ctx.JSON(http.StatusOK, &discovery)
		})
		wellKnown.GET("/health", healthHandler(healthFunc))
	}

	r.GET("/swagger.json", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, string(SwaggerJSON))
	})
}

// healthHandler is the health HTTP handler used for the /.well-known/health
// route if enabled.
func healthHandler(healthFunc func() bool) gin.HandlerFunc {
	healthy := healthFunc
	return func(ctx *gin.Context) {
		health := struct {
			Health bool `json:"health"`
		}{
			Health: healthy(),
		}

		if !health.Health {
			ctx.JSON(http.StatusServiceUnavailable, &health)
		} else {
			ctx.JSON(http.StatusOK, &health)
		}
	}
}

// Service is the interface that must be implemented in order to provide
// business logic for the API service.
type Service interface {
	Healthy() bool
	SearchAPISpecs(ctx *gin.Context, params *operations.SearchAPISpecsParams) *api.Response
}

func ginizePath(path string) string {
	return strings.Replace(strings.Replace(path, "{", ":", -1), "}", "", -1)
}

// configureRoutes configures the routes for the API service.
func configureRoutes(service Service, enableAuth bool, tokenURL string) *Routes {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.LogrusLogger())
	routes := &Routes{Engine: engine}

	routes.SearchAPISpecs.RouterGroup = routes.Group("")
	routes.SearchAPISpecs.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.SearchAPISpecs.Post = routes.SearchAPISpecs.Group("")
	routes.SearchAPISpecs.Post.POST(ginizePath("/specs/_search"), operations.BusinessLogicSearchAPISpecs(service.SearchAPISpecs))

	return routes
}

// API defines the API service.
type API struct {
	Routes  *Routes
	config  *Config
	server  *http.Server
	Title   string
	Version string
}

// NewAPI initializes a new API service.
func NewAPI(svc Service, config *Config) *API {
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	api := &API{
		Routes:  configureRoutes(svc, !config.AuthDisabled, config.TokenURL),
		config:  config,
		Title:   "searchkit-apis",
		Version: "0.0.1",
	}

	// enable pprof http endpoints in debug mode
	if config.Debug {
		pprof.Register(api.Routes.Engine, nil)
	}

	// set logrus logger to TextFormatter with no colors
	log.SetFormatter(&log.TextFormatter{DisableColors: true})

	api.server = &http.Server{
		Addr:         config.Address,
		Handler:      api.Routes.Engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if !config.WellKnownDisabled {
		api.Routes.configureWellKnown(svc.Healthy)
	}

	// configure healthz endpoint
	api.Routes.GET("/healthz", healthHandler(svc.Healthy))

	return api
}

// Run runs the API server it will listen on either HTTP or HTTPS depending on
// the config passed to NewAPI.
func (a *API) Run() error {
	log.Infof("Serving '%s - %s' on address %s", a.Title, a.Version, a.server.Addr)
	if a.config.InsecureHTTP {
		return a.server.ListenAndServe()
	}
	return a.server.ListenAndServeTLS(a.config.TLSCertFile, a.config.TLSKeyFile)
}

// Shutdown will gracefully shutdown the API server.
func (a *API) Shutdown() error {
	return a.server.Shutdown(context.Background())
}

// RunWithSigHandler runs the API server with SIGTERM handling automatically
// enabled. The server will listen for a SIGTERM signal and gracefully shutdown
// the web server.
// It's possible to optionally pass any number shutdown functions which will
// execute one by one after the webserver has been shutdown successfully.
func (a *API) RunWithSigHandler(shutdown ...func() error) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)

	go func() {
		<-sigCh
		a.Shutdown()
	}()

	err := a.Run()
	if err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}

	for _, fn := range shutdown {
		err := fn()
		if err != nil {
			return err
		}
	}

	return nil
}

// vim: ft=go
