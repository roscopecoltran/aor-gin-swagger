package main

import (
	"net/http"

	"github.com/roscopecoltran/aor-gin-swagger/api"

	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/models"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/apply_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/authentication_rest_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/data_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/data_source_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/permission_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/role_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/schema_controller"
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/restapi/operations/user_controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/k0kubun/pp"
)

func CreateOrUpdateJwtToken(db *gorm.DB, jwtUser *models.JwtUser) (bool, error) {
	var existing models.JwtUser
	if db.Where("user_name = ? AND email = ?", jwtUser.Username, jwtUser.Email).First(&existing).RecordNotFound() {
		err := db.Create(jwtUser).Error
		return err == nil, err
	}
	jwtUser.Username = existing.Username
	jwtUser.Email = existing.Email
	// jwtUser.ID = existing.ID
	// jwtUser.CreatedAt = existing.CreatedAt
	return false, db.Save(jwtUser).Error
}

type AorEntityExampleServer struct {
	Health bool
}

func (s *AorEntityExampleServer) Healthy() bool {
	return s.Health
}

func (s *AorEntityExampleServer) AddDataSourceUsingOPTIONS(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddDataSourceUsingPOST(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddEntityUsingOPTIONS(ctx *gin.Context, params *schema_controller.AddEntityUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddEntityUsingPOST(ctx *gin.Context, params *schema_controller.AddEntityUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddFieldUsingOPTIONS(ctx *gin.Context, params *schema_controller.AddFieldUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddFieldUsingPOST(ctx *gin.Context, params *schema_controller.AddFieldUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddPermissionUsingOPTIONS(ctx *gin.Context, params *permission_controller.AddPermissionUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddPermissionUsingPOST(ctx *gin.Context, params *permission_controller.AddPermissionUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddRoleUsingOPTIONS(ctx *gin.Context, params *role_controller.AddRoleUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddRoleUsingPOST(ctx *gin.Context, params *role_controller.AddRoleUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddUserUsingOPTIONS(ctx *gin.Context, params *user_controller.AddUserUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddUserUsingPOST(ctx *gin.Context, params *user_controller.AddUserUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) ApplyUsingOPTIONS(ctx *gin.Context, params *apply_controller.ApplyUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) ApplyUsingPOST(ctx *gin.Context, params *apply_controller.ApplyUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) CreateAuthenticationTokenUsingOPTIONS(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) CreateAuthenticationTokenUsingPOST(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) DataMutationUsingDELETE(ctx *gin.Context, params *data_controller.DataMutationUsingDELETEParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) DataMutationUsingOPTIONS(ctx *gin.Context, params *data_controller.DataMutationUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) DataMutationUsingPOST(ctx *gin.Context, params *data_controller.DataMutationUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) DataMutationUsingPUT(ctx *gin.Context, params *data_controller.DataMutationUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) DataQueryUsingGET(ctx *gin.Context, params *data_controller.DataQueryUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditDataSourceUsingPUT(ctx *gin.Context, params *data_source_controller.EditDataSourceUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditEntityUsingPUT(ctx *gin.Context, params *schema_controller.EditEntityUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditFieldUsingPUT(ctx *gin.Context, params *schema_controller.EditFieldUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditFieldUsingPUT1(ctx *gin.Context, params *user_controller.EditFieldUsingPUT1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditPermissionUsingPUT(ctx *gin.Context, params *permission_controller.EditPermissionUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditRoleUsingPUT(ctx *gin.Context, params *role_controller.EditRoleUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindAllFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindAllFieldsUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindEntityFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindEntityFieldsUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindOneFieldUsingGET(ctx *gin.Context, params *schema_controller.FindOneFieldUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindOneUsingGET(ctx *gin.Context, params *data_controller.FindOneUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindRoleUsingGET(ctx *gin.Context, params *data_source_controller.FindRoleUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindRoleUsingGET1(ctx *gin.Context, params *role_controller.FindRoleUsingGET1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindSchemaEntityGET(ctx *gin.Context, params *schema_controller.FindSchemaEntityGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindUserUsingGET(ctx *gin.Context, params *permission_controller.FindUserUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindUserUsingGET1(ctx *gin.Context, params *user_controller.FindUserUsingGET1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) GetAuthenticatedUserUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) GetRoles(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) GetSchemasUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) ListUsersGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) ListUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) ListUsingGET1(ctx *gin.Context, params *permission_controller.ListUsingGET1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) RefreshAndGetAuthenticationTokenUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) ResetCurrentDsUsingPUT(ctx *gin.Context, params *schema_controller.ResetCurrentDsUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) SyncSchemasUsingPUT(ctx *gin.Context, params *schema_controller.SyncSchemasUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}
