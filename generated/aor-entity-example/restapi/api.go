package restapi

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/oauth2"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	ginoauth2 "github.com/zalando/gin-oauth2"

	"github.com/roscopecoltran/aor-gin-swagger/api"
	"github.com/roscopecoltran/aor-gin-swagger/config"
	"github.com/roscopecoltran/aor-gin-swagger/database"
	"github.com/roscopecoltran/aor-gin-swagger/middleware"

	// zap
	// logrus
	log "github.com/sirupsen/logrus"

	// DefaultImports

	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/apply_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/authentication_rest_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/data_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/data_source_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/permission_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/role_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/schema_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/user_controller"
	// Imports
)

// Routes defines all the routes of the API service.
type Routes struct {
	*gin.Engine
	AddApplyUsingOPTIONS struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddApplyUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddDataSourceUsingOPTIONS struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddDataSourceUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddEntityUsingOPTIONS struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddEntityUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddFieldUsingOPTIONS struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddFieldUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddPermissionUsingOPTIONS struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddPermissionUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddRoleUsingOPTIONS struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddRoleUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddUserUsingOPTIONS struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddUserUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	CreateAuthenticationTokenUsingOPTIONS struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	CreateAuthenticationTokenUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	DataMutationUsingDELETE struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	DataMutationUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	DataMutationUsingOPTIONS struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	DataMutationUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	DataMutationUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditDataSourceUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditEntityUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditPermissionUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditRoleUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditSchemaFieldUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditUserFieldUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindDataSourceUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindEntityFieldsUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindFieldUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindFieldsUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindPermissionUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindRoleUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindSchemaEntityGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindUserUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	GetAuthenticatedUserUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	GetDataMutationUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	GetRoles struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	GetSchemasUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	ListDataSourceUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	ListPermissionUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	ListUsersGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	RefreshAndGetAuthenticationTokenUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	ResetCurrentDsUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	SyncSchemasUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
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
	AddApplyUsingOPTIONS(ctx *gin.Context, params *apply_controller.AddApplyUsingOPTIONSParams) *api.Response
	AddApplyUsingPOST(ctx *gin.Context, params *apply_controller.AddApplyUsingPOSTParams) *api.Response
	AddDataSourceUsingOPTIONS(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingOPTIONSParams) *api.Response
	AddDataSourceUsingPOST(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingPOSTParams) *api.Response
	AddEntityUsingOPTIONS(ctx *gin.Context, params *schema_controller.AddEntityUsingOPTIONSParams) *api.Response
	AddEntityUsingPOST(ctx *gin.Context, params *schema_controller.AddEntityUsingPOSTParams) *api.Response
	AddFieldUsingOPTIONS(ctx *gin.Context, params *schema_controller.AddFieldUsingOPTIONSParams) *api.Response
	AddFieldUsingPOST(ctx *gin.Context, params *schema_controller.AddFieldUsingPOSTParams) *api.Response
	AddPermissionUsingOPTIONS(ctx *gin.Context, params *permission_controller.AddPermissionUsingOPTIONSParams) *api.Response
	AddPermissionUsingPOST(ctx *gin.Context, params *permission_controller.AddPermissionUsingPOSTParams) *api.Response
	AddRoleUsingOPTIONS(ctx *gin.Context, params *role_controller.AddRoleUsingOPTIONSParams) *api.Response
	AddRoleUsingPOST(ctx *gin.Context, params *role_controller.AddRoleUsingPOSTParams) *api.Response
	AddUserUsingOPTIONS(ctx *gin.Context, params *user_controller.AddUserUsingOPTIONSParams) *api.Response
	AddUserUsingPOST(ctx *gin.Context, params *user_controller.AddUserUsingPOSTParams) *api.Response
	CreateAuthenticationTokenUsingOPTIONS(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingOPTIONSParams) *api.Response
	CreateAuthenticationTokenUsingPOST(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingPOSTParams) *api.Response
	DataMutationUsingDELETE(ctx *gin.Context, params *data_controller.DataMutationUsingDELETEParams) *api.Response
	DataMutationUsingGET(ctx *gin.Context, params *data_controller.DataMutationUsingGETParams) *api.Response
	DataMutationUsingOPTIONS(ctx *gin.Context, params *data_controller.DataMutationUsingOPTIONSParams) *api.Response
	DataMutationUsingPOST(ctx *gin.Context, params *data_controller.DataMutationUsingPOSTParams) *api.Response
	DataMutationUsingPUT(ctx *gin.Context, params *data_controller.DataMutationUsingPUTParams) *api.Response
	EditDataSourceUsingPUT(ctx *gin.Context, params *data_source_controller.EditDataSourceUsingPUTParams) *api.Response
	EditEntityUsingPUT(ctx *gin.Context, params *schema_controller.EditEntityUsingPUTParams) *api.Response
	EditPermissionUsingPUT(ctx *gin.Context, params *permission_controller.EditPermissionUsingPUTParams) *api.Response
	EditRoleUsingPUT(ctx *gin.Context, params *role_controller.EditRoleUsingPUTParams) *api.Response
	EditSchemaFieldUsingPUT(ctx *gin.Context, params *schema_controller.EditSchemaFieldUsingPUTParams) *api.Response
	EditUserFieldUsingPUT(ctx *gin.Context, params *user_controller.EditUserFieldUsingPUTParams) *api.Response
	FindDataSourceUsingGET(ctx *gin.Context, params *data_source_controller.FindDataSourceUsingGETParams) *api.Response
	FindEntityFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindEntityFieldsUsingGETParams) *api.Response
	FindFieldUsingGET(ctx *gin.Context, params *schema_controller.FindFieldUsingGETParams) *api.Response
	FindFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindFieldsUsingGETParams) *api.Response
	FindPermissionUsingGET(ctx *gin.Context, params *permission_controller.FindPermissionUsingGETParams) *api.Response
	FindRoleUsingGET(ctx *gin.Context, params *role_controller.FindRoleUsingGETParams) *api.Response
	FindSchemaEntityGET(ctx *gin.Context, params *schema_controller.FindSchemaEntityGETParams) *api.Response
	FindUserUsingGET(ctx *gin.Context, params *user_controller.FindUserUsingGETParams) *api.Response
	GetAuthenticatedUserUsingGET(ctx *gin.Context) *api.Response
	GetDataMutationUsingGET(ctx *gin.Context, params *data_controller.GetDataMutationUsingGETParams) *api.Response
	GetRoles(ctx *gin.Context) *api.Response
	GetSchemasUsingGET(ctx *gin.Context) *api.Response
	ListDataSourceUsingGET(ctx *gin.Context) *api.Response
	ListPermissionUsingGET(ctx *gin.Context, params *permission_controller.ListPermissionUsingGETParams) *api.Response
	ListUsersGET(ctx *gin.Context) *api.Response
	RefreshAndGetAuthenticationTokenUsingGET(ctx *gin.Context) *api.Response
	ResetCurrentDsUsingPUT(ctx *gin.Context, params *schema_controller.ResetCurrentDsUsingPUTParams) *api.Response
	SyncSchemasUsingPUT(ctx *gin.Context, params *schema_controller.SyncSchemasUsingPUTParams) *api.Response
}

func ginizePath(path string) string {
	return strings.Replace(strings.Replace(path, "{", ":", -1), "}", "", -1)
}

// configureRoutes configures the routes for the API service.
func configureRoutes(service Service, enableAuth bool, tokenURL string) *Routes {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.LogrusLogger())

	svcConfig := config.Config

	if svcConfig.API.CORS.Enabled {
		engine.Use(middleware.CORS())
	}

	if svcConfig.DB.Enabled {
		engine.Use(middleware.DatabaseContext(database.DB))
	}

	if svcConfig.API.JWT.Enabled {
		fmt.Printf("UseJWT? %t \n", svcConfig.API.JWT.Enabled)
		fmt.Println("SecretKey: ", svcConfig.API.JWT.SecretKey)
	}
	routes := &Routes{Engine: engine}

	routes.AddApplyUsingOPTIONS.RouterGroup = routes.Group("")
	routes.AddApplyUsingOPTIONS.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddApplyUsingOPTIONS.RouterGroup.Use(routes.AddApplyUsingOPTIONS.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddApplyUsingOPTIONS.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddApplyUsingOPTIONS.RouterGroup.Use(routes.AddApplyUsingOPTIONS.Auth)

		routes.AddApplyUsingOPTIONS.RouterGroup.Use(routes.AddApplyUsingOPTIONS.Auth)

	}
	routes.AddApplyUsingOPTIONS.Post = routes.AddApplyUsingOPTIONS.Group("")
	routes.AddApplyUsingOPTIONS.Post.OPTIONS(ginizePath("/apply"), apply_controller.BusinessLogicAddApplyUsingOPTIONS(service.AddApplyUsingOPTIONS))

	routes.AddApplyUsingPOST.RouterGroup = routes.Group("")
	routes.AddApplyUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddApplyUsingPOST.RouterGroup.Use(routes.AddApplyUsingPOST.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddApplyUsingPOST.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddApplyUsingPOST.RouterGroup.Use(routes.AddApplyUsingPOST.Auth)

		routes.AddApplyUsingPOST.RouterGroup.Use(routes.AddApplyUsingPOST.Auth)

	}
	routes.AddApplyUsingPOST.Post = routes.AddApplyUsingPOST.Group("")
	routes.AddApplyUsingPOST.Post.POST(ginizePath("/apply"), apply_controller.BusinessLogicAddApplyUsingPOST(service.AddApplyUsingPOST))

	routes.AddDataSourceUsingOPTIONS.RouterGroup = routes.Group("")
	routes.AddDataSourceUsingOPTIONS.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddDataSourceUsingOPTIONS.RouterGroup.Use(routes.AddDataSourceUsingOPTIONS.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddDataSourceUsingOPTIONS.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddDataSourceUsingOPTIONS.RouterGroup.Use(routes.AddDataSourceUsingOPTIONS.Auth)

		routes.AddDataSourceUsingOPTIONS.RouterGroup.Use(routes.AddDataSourceUsingOPTIONS.Auth)

	}
	routes.AddDataSourceUsingOPTIONS.Post = routes.AddDataSourceUsingOPTIONS.Group("")
	routes.AddDataSourceUsingOPTIONS.Post.OPTIONS(ginizePath("/datasource/_datasource"), data_source_controller.BusinessLogicAddDataSourceUsingOPTIONS(service.AddDataSourceUsingOPTIONS))

	routes.AddDataSourceUsingPOST.RouterGroup = routes.Group("")
	routes.AddDataSourceUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddDataSourceUsingPOST.RouterGroup.Use(routes.AddDataSourceUsingPOST.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddDataSourceUsingPOST.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddDataSourceUsingPOST.RouterGroup.Use(routes.AddDataSourceUsingPOST.Auth)

		routes.AddDataSourceUsingPOST.RouterGroup.Use(routes.AddDataSourceUsingPOST.Auth)

	}
	routes.AddDataSourceUsingPOST.Post = routes.AddDataSourceUsingPOST.Group("")
	routes.AddDataSourceUsingPOST.Post.POST(ginizePath("/datasource/_datasource"), data_source_controller.BusinessLogicAddDataSourceUsingPOST(service.AddDataSourceUsingPOST))

	routes.AddEntityUsingOPTIONS.RouterGroup = routes.Group("")
	routes.AddEntityUsingOPTIONS.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddEntityUsingOPTIONS.RouterGroup.Use(routes.AddEntityUsingOPTIONS.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddEntityUsingOPTIONS.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddEntityUsingOPTIONS.RouterGroup.Use(routes.AddEntityUsingOPTIONS.Auth)

		routes.AddEntityUsingOPTIONS.RouterGroup.Use(routes.AddEntityUsingOPTIONS.Auth)

	}
	routes.AddEntityUsingOPTIONS.Post = routes.AddEntityUsingOPTIONS.Group("")
	routes.AddEntityUsingOPTIONS.Post.OPTIONS(ginizePath("/schemas/_entitys"), schema_controller.BusinessLogicAddEntityUsingOPTIONS(service.AddEntityUsingOPTIONS))

	routes.AddEntityUsingPOST.RouterGroup = routes.Group("")
	routes.AddEntityUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddEntityUsingPOST.RouterGroup.Use(routes.AddEntityUsingPOST.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddEntityUsingPOST.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddEntityUsingPOST.RouterGroup.Use(routes.AddEntityUsingPOST.Auth)

		routes.AddEntityUsingPOST.RouterGroup.Use(routes.AddEntityUsingPOST.Auth)

	}
	routes.AddEntityUsingPOST.Post = routes.AddEntityUsingPOST.Group("")
	routes.AddEntityUsingPOST.Post.POST(ginizePath("/schemas/_entitys"), schema_controller.BusinessLogicAddEntityUsingPOST(service.AddEntityUsingPOST))

	routes.AddFieldUsingOPTIONS.RouterGroup = routes.Group("")
	routes.AddFieldUsingOPTIONS.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddFieldUsingOPTIONS.RouterGroup.Use(routes.AddFieldUsingOPTIONS.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddFieldUsingOPTIONS.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddFieldUsingOPTIONS.RouterGroup.Use(routes.AddFieldUsingOPTIONS.Auth)

		routes.AddFieldUsingOPTIONS.RouterGroup.Use(routes.AddFieldUsingOPTIONS.Auth)

	}
	routes.AddFieldUsingOPTIONS.Post = routes.AddFieldUsingOPTIONS.Group("")
	routes.AddFieldUsingOPTIONS.Post.OPTIONS(ginizePath("/schemas/_fields"), schema_controller.BusinessLogicAddFieldUsingOPTIONS(service.AddFieldUsingOPTIONS))

	routes.AddFieldUsingPOST.RouterGroup = routes.Group("")
	routes.AddFieldUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddFieldUsingPOST.RouterGroup.Use(routes.AddFieldUsingPOST.Auth)

		routes.AddFieldUsingPOST.RouterGroup.Use(routes.AddFieldUsingPOST.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddFieldUsingPOST.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddFieldUsingPOST.RouterGroup.Use(routes.AddFieldUsingPOST.Auth)

	}
	routes.AddFieldUsingPOST.Post = routes.AddFieldUsingPOST.Group("")
	routes.AddFieldUsingPOST.Post.POST(ginizePath("/schemas/_fields"), schema_controller.BusinessLogicAddFieldUsingPOST(service.AddFieldUsingPOST))

	routes.AddPermissionUsingOPTIONS.RouterGroup = routes.Group("")
	routes.AddPermissionUsingOPTIONS.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddPermissionUsingOPTIONS.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddPermissionUsingOPTIONS.RouterGroup.Use(routes.AddPermissionUsingOPTIONS.Auth)

		routes.AddPermissionUsingOPTIONS.RouterGroup.Use(routes.AddPermissionUsingOPTIONS.Auth)

		routes.AddPermissionUsingOPTIONS.RouterGroup.Use(routes.AddPermissionUsingOPTIONS.Auth)

	}
	routes.AddPermissionUsingOPTIONS.Post = routes.AddPermissionUsingOPTIONS.Group("")
	routes.AddPermissionUsingOPTIONS.Post.OPTIONS(ginizePath("/permission/_permission"), permission_controller.BusinessLogicAddPermissionUsingOPTIONS(service.AddPermissionUsingOPTIONS))

	routes.AddPermissionUsingPOST.RouterGroup = routes.Group("")
	routes.AddPermissionUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddPermissionUsingPOST.RouterGroup.Use(routes.AddPermissionUsingPOST.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddPermissionUsingPOST.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddPermissionUsingPOST.RouterGroup.Use(routes.AddPermissionUsingPOST.Auth)

		routes.AddPermissionUsingPOST.RouterGroup.Use(routes.AddPermissionUsingPOST.Auth)

	}
	routes.AddPermissionUsingPOST.Post = routes.AddPermissionUsingPOST.Group("")
	routes.AddPermissionUsingPOST.Post.POST(ginizePath("/permission/_permission"), permission_controller.BusinessLogicAddPermissionUsingPOST(service.AddPermissionUsingPOST))

	routes.AddRoleUsingOPTIONS.RouterGroup = routes.Group("")
	routes.AddRoleUsingOPTIONS.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddRoleUsingOPTIONS.RouterGroup.Use(routes.AddRoleUsingOPTIONS.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddRoleUsingOPTIONS.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddRoleUsingOPTIONS.RouterGroup.Use(routes.AddRoleUsingOPTIONS.Auth)

		routes.AddRoleUsingOPTIONS.RouterGroup.Use(routes.AddRoleUsingOPTIONS.Auth)

	}
	routes.AddRoleUsingOPTIONS.Post = routes.AddRoleUsingOPTIONS.Group("")
	routes.AddRoleUsingOPTIONS.Post.OPTIONS(ginizePath("/role/_roles"), role_controller.BusinessLogicAddRoleUsingOPTIONS(service.AddRoleUsingOPTIONS))

	routes.AddRoleUsingPOST.RouterGroup = routes.Group("")
	routes.AddRoleUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddRoleUsingPOST.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddRoleUsingPOST.RouterGroup.Use(routes.AddRoleUsingPOST.Auth)

		routes.AddRoleUsingPOST.RouterGroup.Use(routes.AddRoleUsingPOST.Auth)

		routes.AddRoleUsingPOST.RouterGroup.Use(routes.AddRoleUsingPOST.Auth)

	}
	routes.AddRoleUsingPOST.Post = routes.AddRoleUsingPOST.Group("")
	routes.AddRoleUsingPOST.Post.POST(ginizePath("/role/_roles"), role_controller.BusinessLogicAddRoleUsingPOST(service.AddRoleUsingPOST))

	routes.AddUserUsingOPTIONS.RouterGroup = routes.Group("")
	routes.AddUserUsingOPTIONS.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddUserUsingOPTIONS.RouterGroup.Use(routes.AddUserUsingOPTIONS.Auth)

		routes.AddUserUsingOPTIONS.RouterGroup.Use(routes.AddUserUsingOPTIONS.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddUserUsingOPTIONS.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddUserUsingOPTIONS.RouterGroup.Use(routes.AddUserUsingOPTIONS.Auth)

	}
	routes.AddUserUsingOPTIONS.Post = routes.AddUserUsingOPTIONS.Group("")
	routes.AddUserUsingOPTIONS.Post.OPTIONS(ginizePath("/user/_users"), user_controller.BusinessLogicAddUserUsingOPTIONS(service.AddUserUsingOPTIONS))

	routes.AddUserUsingPOST.RouterGroup = routes.Group("")
	routes.AddUserUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.AddUserUsingPOST.RouterGroup.Use(routes.AddUserUsingPOST.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddUserUsingPOST.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.AddUserUsingPOST.RouterGroup.Use(routes.AddUserUsingPOST.Auth)

		routes.AddUserUsingPOST.RouterGroup.Use(routes.AddUserUsingPOST.Auth)

	}
	routes.AddUserUsingPOST.Post = routes.AddUserUsingPOST.Group("")
	routes.AddUserUsingPOST.Post.POST(ginizePath("/user/_users"), user_controller.BusinessLogicAddUserUsingPOST(service.AddUserUsingPOST))

	routes.CreateAuthenticationTokenUsingOPTIONS.RouterGroup = routes.Group("")
	routes.CreateAuthenticationTokenUsingOPTIONS.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.CreateAuthenticationTokenUsingOPTIONS.RouterGroup.Use(routes.CreateAuthenticationTokenUsingOPTIONS.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.CreateAuthenticationTokenUsingOPTIONS.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.CreateAuthenticationTokenUsingOPTIONS.RouterGroup.Use(routes.CreateAuthenticationTokenUsingOPTIONS.Auth)

		routes.CreateAuthenticationTokenUsingOPTIONS.RouterGroup.Use(routes.CreateAuthenticationTokenUsingOPTIONS.Auth)

	}
	routes.CreateAuthenticationTokenUsingOPTIONS.Post = routes.CreateAuthenticationTokenUsingOPTIONS.Group("")
	routes.CreateAuthenticationTokenUsingOPTIONS.Post.OPTIONS(ginizePath("/auth"), authentication_rest_controller.BusinessLogicCreateAuthenticationTokenUsingOPTIONS(service.CreateAuthenticationTokenUsingOPTIONS))

	routes.CreateAuthenticationTokenUsingPOST.RouterGroup = routes.Group("")
	routes.CreateAuthenticationTokenUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.CreateAuthenticationTokenUsingPOST.RouterGroup.Use(routes.CreateAuthenticationTokenUsingPOST.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.CreateAuthenticationTokenUsingPOST.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.CreateAuthenticationTokenUsingPOST.RouterGroup.Use(routes.CreateAuthenticationTokenUsingPOST.Auth)

		routes.CreateAuthenticationTokenUsingPOST.RouterGroup.Use(routes.CreateAuthenticationTokenUsingPOST.Auth)

	}
	routes.CreateAuthenticationTokenUsingPOST.Post = routes.CreateAuthenticationTokenUsingPOST.Group("")
	routes.CreateAuthenticationTokenUsingPOST.Post.POST(ginizePath("/auth"), authentication_rest_controller.BusinessLogicCreateAuthenticationTokenUsingPOST(service.CreateAuthenticationTokenUsingPOST))

	routes.DataMutationUsingDELETE.RouterGroup = routes.Group("")
	routes.DataMutationUsingDELETE.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.DataMutationUsingDELETE.RouterGroup.Use(routes.DataMutationUsingDELETE.Auth)

		routes.DataMutationUsingDELETE.RouterGroup.Use(routes.DataMutationUsingDELETE.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.DataMutationUsingDELETE.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.DataMutationUsingDELETE.RouterGroup.Use(routes.DataMutationUsingDELETE.Auth)

	}
	routes.DataMutationUsingDELETE.Post = routes.DataMutationUsingDELETE.Group("")
	routes.DataMutationUsingDELETE.Post.DELETE(ginizePath("/api/{entity}/{id}"), data_controller.BusinessLogicDataMutationUsingDELETE(service.DataMutationUsingDELETE))

	routes.DataMutationUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.DataMutationUsingGET.RouterGroup.Use(routes.DataMutationUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.DataMutationUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.DataMutationUsingGET.RouterGroup.Use(routes.DataMutationUsingGET.Auth)

		routes.DataMutationUsingGET.RouterGroup.Use(routes.DataMutationUsingGET.Auth)

	}
	routes.DataMutationUsingGET.Post = routes.DataMutationUsingGET.Group("")
	routes.DataMutationUsingGET.Post.GET(ginizePath("/api/{entity}/{id}"), data_controller.BusinessLogicDataMutationUsingGET(service.DataMutationUsingGET))

	routes.DataMutationUsingOPTIONS.RouterGroup = routes.Group("")
	routes.DataMutationUsingOPTIONS.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.DataMutationUsingOPTIONS.RouterGroup.Use(routes.DataMutationUsingOPTIONS.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.DataMutationUsingOPTIONS.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.DataMutationUsingOPTIONS.RouterGroup.Use(routes.DataMutationUsingOPTIONS.Auth)

		routes.DataMutationUsingOPTIONS.RouterGroup.Use(routes.DataMutationUsingOPTIONS.Auth)

	}
	routes.DataMutationUsingOPTIONS.Post = routes.DataMutationUsingOPTIONS.Group("")
	routes.DataMutationUsingOPTIONS.Post.OPTIONS(ginizePath("/api/{entity}"), data_controller.BusinessLogicDataMutationUsingOPTIONS(service.DataMutationUsingOPTIONS))

	routes.DataMutationUsingPOST.RouterGroup = routes.Group("")
	routes.DataMutationUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.DataMutationUsingPOST.RouterGroup.Use(routes.DataMutationUsingPOST.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.DataMutationUsingPOST.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.DataMutationUsingPOST.RouterGroup.Use(routes.DataMutationUsingPOST.Auth)

		routes.DataMutationUsingPOST.RouterGroup.Use(routes.DataMutationUsingPOST.Auth)

	}
	routes.DataMutationUsingPOST.Post = routes.DataMutationUsingPOST.Group("")
	routes.DataMutationUsingPOST.Post.POST(ginizePath("/api/{entity}"), data_controller.BusinessLogicDataMutationUsingPOST(service.DataMutationUsingPOST))

	routes.DataMutationUsingPUT.RouterGroup = routes.Group("")
	routes.DataMutationUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.DataMutationUsingPUT.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.DataMutationUsingPUT.RouterGroup.Use(routes.DataMutationUsingPUT.Auth)

		routes.DataMutationUsingPUT.RouterGroup.Use(routes.DataMutationUsingPUT.Auth)

		routes.DataMutationUsingPUT.RouterGroup.Use(routes.DataMutationUsingPUT.Auth)

	}
	routes.DataMutationUsingPUT.Post = routes.DataMutationUsingPUT.Group("")
	routes.DataMutationUsingPUT.Post.PUT(ginizePath("/api/{entity}/{id}"), data_controller.BusinessLogicDataMutationUsingPUT(service.DataMutationUsingPUT))

	routes.EditDataSourceUsingPUT.RouterGroup = routes.Group("")
	routes.EditDataSourceUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.EditDataSourceUsingPUT.RouterGroup.Use(routes.EditDataSourceUsingPUT.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.EditDataSourceUsingPUT.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.EditDataSourceUsingPUT.RouterGroup.Use(routes.EditDataSourceUsingPUT.Auth)

		routes.EditDataSourceUsingPUT.RouterGroup.Use(routes.EditDataSourceUsingPUT.Auth)

	}
	routes.EditDataSourceUsingPUT.Post = routes.EditDataSourceUsingPUT.Group("")
	routes.EditDataSourceUsingPUT.Post.PUT(ginizePath("/datasource/_datasource/put/{id}"), data_source_controller.BusinessLogicEditDataSourceUsingPUT(service.EditDataSourceUsingPUT))

	routes.EditEntityUsingPUT.RouterGroup = routes.Group("")
	routes.EditEntityUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.EditEntityUsingPUT.RouterGroup.Use(routes.EditEntityUsingPUT.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.EditEntityUsingPUT.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.EditEntityUsingPUT.RouterGroup.Use(routes.EditEntityUsingPUT.Auth)

		routes.EditEntityUsingPUT.RouterGroup.Use(routes.EditEntityUsingPUT.Auth)

	}
	routes.EditEntityUsingPUT.Post = routes.EditEntityUsingPUT.Group("")
	routes.EditEntityUsingPUT.Post.PUT(ginizePath("/schemas/_entitys/put/{id}"), schema_controller.BusinessLogicEditEntityUsingPUT(service.EditEntityUsingPUT))

	routes.EditPermissionUsingPUT.RouterGroup = routes.Group("")
	routes.EditPermissionUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.EditPermissionUsingPUT.RouterGroup.Use(routes.EditPermissionUsingPUT.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.EditPermissionUsingPUT.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.EditPermissionUsingPUT.RouterGroup.Use(routes.EditPermissionUsingPUT.Auth)

		routes.EditPermissionUsingPUT.RouterGroup.Use(routes.EditPermissionUsingPUT.Auth)

	}
	routes.EditPermissionUsingPUT.Post = routes.EditPermissionUsingPUT.Group("")
	routes.EditPermissionUsingPUT.Post.PUT(ginizePath("/permission/_permission/{id}"), permission_controller.BusinessLogicEditPermissionUsingPUT(service.EditPermissionUsingPUT))

	routes.EditRoleUsingPUT.RouterGroup = routes.Group("")
	routes.EditRoleUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.EditRoleUsingPUT.RouterGroup.Use(routes.EditRoleUsingPUT.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.EditRoleUsingPUT.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.EditRoleUsingPUT.RouterGroup.Use(routes.EditRoleUsingPUT.Auth)

		routes.EditRoleUsingPUT.RouterGroup.Use(routes.EditRoleUsingPUT.Auth)

	}
	routes.EditRoleUsingPUT.Post = routes.EditRoleUsingPUT.Group("")
	routes.EditRoleUsingPUT.Post.PUT(ginizePath("/role/_roles/put/{id}"), role_controller.BusinessLogicEditRoleUsingPUT(service.EditRoleUsingPUT))

	routes.EditSchemaFieldUsingPUT.RouterGroup = routes.Group("")
	routes.EditSchemaFieldUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.EditSchemaFieldUsingPUT.RouterGroup.Use(routes.EditSchemaFieldUsingPUT.Auth)

		routes.EditSchemaFieldUsingPUT.RouterGroup.Use(routes.EditSchemaFieldUsingPUT.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.EditSchemaFieldUsingPUT.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.EditSchemaFieldUsingPUT.RouterGroup.Use(routes.EditSchemaFieldUsingPUT.Auth)

	}
	routes.EditSchemaFieldUsingPUT.Post = routes.EditSchemaFieldUsingPUT.Group("")
	routes.EditSchemaFieldUsingPUT.Post.PUT(ginizePath("/schemas/_fields/put/{id}"), schema_controller.BusinessLogicEditSchemaFieldUsingPUT(service.EditSchemaFieldUsingPUT))

	routes.EditUserFieldUsingPUT.RouterGroup = routes.Group("")
	routes.EditUserFieldUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.EditUserFieldUsingPUT.RouterGroup.Use(routes.EditUserFieldUsingPUT.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.EditUserFieldUsingPUT.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.EditUserFieldUsingPUT.RouterGroup.Use(routes.EditUserFieldUsingPUT.Auth)

		routes.EditUserFieldUsingPUT.RouterGroup.Use(routes.EditUserFieldUsingPUT.Auth)

	}
	routes.EditUserFieldUsingPUT.Post = routes.EditUserFieldUsingPUT.Group("")
	routes.EditUserFieldUsingPUT.Post.PUT(ginizePath("/user/_users/put/{id}"), user_controller.BusinessLogicEditUserFieldUsingPUT(service.EditUserFieldUsingPUT))

	routes.FindDataSourceUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.FindDataSourceUsingGET.RouterGroup.Use(routes.FindDataSourceUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.FindDataSourceUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.FindDataSourceUsingGET.RouterGroup.Use(routes.FindDataSourceUsingGET.Auth)

		routes.FindDataSourceUsingGET.RouterGroup.Use(routes.FindDataSourceUsingGET.Auth)

	}
	routes.FindDataSourceUsingGET.Post = routes.FindDataSourceUsingGET.Group("")
	routes.FindDataSourceUsingGET.Post.GET(ginizePath("/datasource/_datasource/{datasourceId}"), data_source_controller.BusinessLogicFindDataSourceUsingGET(service.FindDataSourceUsingGET))

	routes.FindEntityFieldsUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.FindEntityFieldsUsingGET.RouterGroup.Use(routes.FindEntityFieldsUsingGET.Auth)

		routes.FindEntityFieldsUsingGET.RouterGroup.Use(routes.FindEntityFieldsUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.FindEntityFieldsUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.FindEntityFieldsUsingGET.RouterGroup.Use(routes.FindEntityFieldsUsingGET.Auth)

	}
	routes.FindEntityFieldsUsingGET.Post = routes.FindEntityFieldsUsingGET.Group("")
	routes.FindEntityFieldsUsingGET.Post.GET(ginizePath("/schemas/fields"), schema_controller.BusinessLogicFindEntityFieldsUsingGET(service.FindEntityFieldsUsingGET))

	routes.FindFieldUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.FindFieldUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.FindFieldUsingGET.RouterGroup.Use(routes.FindFieldUsingGET.Auth)

		routes.FindFieldUsingGET.RouterGroup.Use(routes.FindFieldUsingGET.Auth)

		routes.FindFieldUsingGET.RouterGroup.Use(routes.FindFieldUsingGET.Auth)

	}
	routes.FindFieldUsingGET.Post = routes.FindFieldUsingGET.Group("")
	routes.FindFieldUsingGET.Post.GET(ginizePath("/schemas/_fields/{fid}"), schema_controller.BusinessLogicFindFieldUsingGET(service.FindFieldUsingGET))

	routes.FindFieldsUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.FindFieldsUsingGET.RouterGroup.Use(routes.FindFieldsUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.FindFieldsUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.FindFieldsUsingGET.RouterGroup.Use(routes.FindFieldsUsingGET.Auth)

		routes.FindFieldsUsingGET.RouterGroup.Use(routes.FindFieldsUsingGET.Auth)

	}
	routes.FindFieldsUsingGET.Post = routes.FindFieldsUsingGET.Group("")
	routes.FindFieldsUsingGET.Post.GET(ginizePath("/schemas/_fields"), schema_controller.BusinessLogicFindFieldsUsingGET(service.FindFieldsUsingGET))

	routes.FindPermissionUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.FindPermissionUsingGET.RouterGroup.Use(routes.FindPermissionUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.FindPermissionUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.FindPermissionUsingGET.RouterGroup.Use(routes.FindPermissionUsingGET.Auth)

		routes.FindPermissionUsingGET.RouterGroup.Use(routes.FindPermissionUsingGET.Auth)

	}
	routes.FindPermissionUsingGET.Post = routes.FindPermissionUsingGET.Group("")
	routes.FindPermissionUsingGET.Post.GET(ginizePath("/permission/_permission/{id}"), permission_controller.BusinessLogicFindPermissionUsingGET(service.FindPermissionUsingGET))

	routes.FindRoleUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.FindRoleUsingGET.RouterGroup.Use(routes.FindRoleUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.FindRoleUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.FindRoleUsingGET.RouterGroup.Use(routes.FindRoleUsingGET.Auth)

		routes.FindRoleUsingGET.RouterGroup.Use(routes.FindRoleUsingGET.Auth)

	}
	routes.FindRoleUsingGET.Post = routes.FindRoleUsingGET.Group("")
	routes.FindRoleUsingGET.Post.GET(ginizePath("/role/_roles/{roleId}"), role_controller.BusinessLogicFindRoleUsingGET(service.FindRoleUsingGET))

	routes.FindSchemaEntityGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.FindSchemaEntityGET.RouterGroup.Use(routes.FindSchemaEntityGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.FindSchemaEntityGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.FindSchemaEntityGET.RouterGroup.Use(routes.FindSchemaEntityGET.Auth)

		routes.FindSchemaEntityGET.RouterGroup.Use(routes.FindSchemaEntityGET.Auth)

	}
	routes.FindSchemaEntityGET.Post = routes.FindSchemaEntityGET.Group("")
	routes.FindSchemaEntityGET.Post.GET(ginizePath("/schemas/_entitys/{eid}"), schema_controller.BusinessLogicFindSchemaEntityGET(service.FindSchemaEntityGET))

	routes.FindUserUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.FindUserUsingGET.RouterGroup.Use(routes.FindUserUsingGET.Auth)

		routes.FindUserUsingGET.RouterGroup.Use(routes.FindUserUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.FindUserUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.FindUserUsingGET.RouterGroup.Use(routes.FindUserUsingGET.Auth)

	}
	routes.FindUserUsingGET.Post = routes.FindUserUsingGET.Group("")
	routes.FindUserUsingGET.Post.GET(ginizePath("/user/_users/{userId}"), user_controller.BusinessLogicFindUserUsingGET(service.FindUserUsingGET))

	routes.GetAuthenticatedUserUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.GetAuthenticatedUserUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.GetAuthenticatedUserUsingGET.RouterGroup.Use(routes.GetAuthenticatedUserUsingGET.Auth)

		routes.GetAuthenticatedUserUsingGET.RouterGroup.Use(routes.GetAuthenticatedUserUsingGET.Auth)

		routes.GetAuthenticatedUserUsingGET.RouterGroup.Use(routes.GetAuthenticatedUserUsingGET.Auth)

	}
	routes.GetAuthenticatedUserUsingGET.Post = routes.GetAuthenticatedUserUsingGET.Group("")
	routes.GetAuthenticatedUserUsingGET.Post.GET(ginizePath("/user/me"), user_controller.BusinessLogicGetAuthenticatedUserUsingGET(service.GetAuthenticatedUserUsingGET))

	routes.GetDataMutationUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.GetDataMutationUsingGET.RouterGroup.Use(routes.GetDataMutationUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.GetDataMutationUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.GetDataMutationUsingGET.RouterGroup.Use(routes.GetDataMutationUsingGET.Auth)

		routes.GetDataMutationUsingGET.RouterGroup.Use(routes.GetDataMutationUsingGET.Auth)

	}
	routes.GetDataMutationUsingGET.Post = routes.GetDataMutationUsingGET.Group("")
	routes.GetDataMutationUsingGET.Post.GET(ginizePath("/api/{entity}"), data_controller.BusinessLogicGetDataMutationUsingGET(service.GetDataMutationUsingGET))

	routes.GetRoles.RouterGroup = routes.Group("")
	if enableAuth {

		routes.GetRoles.RouterGroup.Use(routes.GetRoles.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.GetRoles.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.GetRoles.RouterGroup.Use(routes.GetRoles.Auth)

		routes.GetRoles.RouterGroup.Use(routes.GetRoles.Auth)

	}
	routes.GetRoles.Post = routes.GetRoles.Group("")
	routes.GetRoles.Post.GET(ginizePath("/role/_roles"), role_controller.BusinessLogicGetRoles(service.GetRoles))

	routes.GetSchemasUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.GetSchemasUsingGET.RouterGroup.Use(routes.GetSchemasUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.GetSchemasUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.GetSchemasUsingGET.RouterGroup.Use(routes.GetSchemasUsingGET.Auth)

		routes.GetSchemasUsingGET.RouterGroup.Use(routes.GetSchemasUsingGET.Auth)

	}
	routes.GetSchemasUsingGET.Post = routes.GetSchemasUsingGET.Group("")
	routes.GetSchemasUsingGET.Post.GET(ginizePath("/schemas/_entitys"), schema_controller.BusinessLogicGetSchemasUsingGET(service.GetSchemasUsingGET))

	routes.ListDataSourceUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.ListDataSourceUsingGET.RouterGroup.Use(routes.ListDataSourceUsingGET.Auth)

		routes.ListDataSourceUsingGET.RouterGroup.Use(routes.ListDataSourceUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.ListDataSourceUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.ListDataSourceUsingGET.RouterGroup.Use(routes.ListDataSourceUsingGET.Auth)

	}
	routes.ListDataSourceUsingGET.Post = routes.ListDataSourceUsingGET.Group("")
	routes.ListDataSourceUsingGET.Post.GET(ginizePath("/datasource/_datasource"), data_source_controller.BusinessLogicListDataSourceUsingGET(service.ListDataSourceUsingGET))

	routes.ListPermissionUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.ListPermissionUsingGET.RouterGroup.Use(routes.ListPermissionUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.ListPermissionUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.ListPermissionUsingGET.RouterGroup.Use(routes.ListPermissionUsingGET.Auth)

		routes.ListPermissionUsingGET.RouterGroup.Use(routes.ListPermissionUsingGET.Auth)

	}
	routes.ListPermissionUsingGET.Post = routes.ListPermissionUsingGET.Group("")
	routes.ListPermissionUsingGET.Post.GET(ginizePath("/permission/_permission"), permission_controller.BusinessLogicListPermissionUsingGET(service.ListPermissionUsingGET))

	routes.ListUsersGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.ListUsersGET.RouterGroup.Use(routes.ListUsersGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.ListUsersGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.ListUsersGET.RouterGroup.Use(routes.ListUsersGET.Auth)

		routes.ListUsersGET.RouterGroup.Use(routes.ListUsersGET.Auth)

	}
	routes.ListUsersGET.Post = routes.ListUsersGET.Group("")
	routes.ListUsersGET.Post.GET(ginizePath("/user/_users"), user_controller.BusinessLogicListUsersGET(service.ListUsersGET))

	routes.RefreshAndGetAuthenticationTokenUsingGET.RouterGroup = routes.Group("")
	if enableAuth {

		routes.RefreshAndGetAuthenticationTokenUsingGET.RouterGroup.Use(routes.RefreshAndGetAuthenticationTokenUsingGET.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.RefreshAndGetAuthenticationTokenUsingGET.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.RefreshAndGetAuthenticationTokenUsingGET.RouterGroup.Use(routes.RefreshAndGetAuthenticationTokenUsingGET.Auth)

		routes.RefreshAndGetAuthenticationTokenUsingGET.RouterGroup.Use(routes.RefreshAndGetAuthenticationTokenUsingGET.Auth)

	}
	routes.RefreshAndGetAuthenticationTokenUsingGET.Post = routes.RefreshAndGetAuthenticationTokenUsingGET.Group("")
	routes.RefreshAndGetAuthenticationTokenUsingGET.Post.GET(ginizePath("/refresh"), authentication_rest_controller.BusinessLogicRefreshAndGetAuthenticationTokenUsingGET(service.RefreshAndGetAuthenticationTokenUsingGET))

	routes.ResetCurrentDsUsingPUT.RouterGroup = routes.Group("")
	routes.ResetCurrentDsUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.ResetCurrentDsUsingPUT.RouterGroup.Use(routes.ResetCurrentDsUsingPUT.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.ResetCurrentDsUsingPUT.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.ResetCurrentDsUsingPUT.RouterGroup.Use(routes.ResetCurrentDsUsingPUT.Auth)

		routes.ResetCurrentDsUsingPUT.RouterGroup.Use(routes.ResetCurrentDsUsingPUT.Auth)

	}
	routes.ResetCurrentDsUsingPUT.Post = routes.ResetCurrentDsUsingPUT.Group("")
	routes.ResetCurrentDsUsingPUT.Post.PUT(ginizePath("/schemas/resetCurrentDs/{dataSourceId}"), schema_controller.BusinessLogicResetCurrentDsUsingPUT(service.ResetCurrentDsUsingPUT))

	routes.SyncSchemasUsingPUT.RouterGroup = routes.Group("")
	routes.SyncSchemasUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {

		routes.SyncSchemasUsingPUT.RouterGroup.Use(routes.SyncSchemasUsingPUT.Auth)

		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "https://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.SyncSchemasUsingPUT.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)

		routes.SyncSchemasUsingPUT.RouterGroup.Use(routes.SyncSchemasUsingPUT.Auth)

		routes.SyncSchemasUsingPUT.RouterGroup.Use(routes.SyncSchemasUsingPUT.Auth)

	}
	routes.SyncSchemasUsingPUT.Post = routes.SyncSchemasUsingPUT.Group("")
	routes.SyncSchemasUsingPUT.Post.PUT(ginizePath("/schemas/sync/{dataSourceId}"), schema_controller.BusinessLogicSyncSchemasUsingPUT(service.SyncSchemasUsingPUT))

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
		Title:   "admin-on-rest entity APIs",
		Version: "1.0",
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
