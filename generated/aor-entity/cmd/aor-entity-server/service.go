package main

import (
	"net/http"

	"github.com/roscopecoltran/gin-swagger/api"

	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/models"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/apply_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/authentication_rest_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/data_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/data_source_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/permission_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/role_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/schema_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/user_controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/k0kubun/pp"
)

func CreateOrUpdateJwtToken(db *gorm.DB, jwtUser *models.JwtUser) (bool, error) {
	// Get existing by remote ID and service ID
	var existing models.JwtUser
	if db.Where("user_name = ? AND email = ?", jwtUser.Username, jwtUser.Email).First(&existing).RecordNotFound() {
		err := db.Create(jwtUser).Error
		return err == nil, err
	}
	jwtUser.ID = existing.ID
	jwtUser.CreatedAt = existing.CreatedAt
	return false, db.Save(jwtUser).Error
}

type AorEntityServer struct {
	Health bool
}

func (s *AorEntityServer) Healthy() bool {
	return s.Health
}

func (s *AorEntityServer) AddDataSourceUsingOPTIONS(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) AddDataSourceUsingPOST(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) AddEntityUsingOPTIONS(ctx *gin.Context, params *schema_controller.AddEntityUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) AddEntityUsingPOST(ctx *gin.Context, params *schema_controller.AddEntityUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) AddFieldUsingOPTIONS(ctx *gin.Context, params *schema_controller.AddFieldUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) AddFieldUsingPOST(ctx *gin.Context, params *schema_controller.AddFieldUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) AddPermissionUsingOPTIONS(ctx *gin.Context, params *permission_controller.AddPermissionUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) AddPermissionUsingPOST(ctx *gin.Context, params *permission_controller.AddPermissionUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) AddRoleUsingOPTIONS(ctx *gin.Context, params *role_controller.AddRoleUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) AddRoleUsingPOST(ctx *gin.Context, params *role_controller.AddRoleUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) AddUserUsingOPTIONS(ctx *gin.Context, params *user_controller.AddUserUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) AddUserUsingPOST(ctx *gin.Context, params *user_controller.AddUserUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) ApplyUsingOPTIONS(ctx *gin.Context, params *apply_controller.ApplyUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) ApplyUsingPOST(ctx *gin.Context, params *apply_controller.ApplyUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) CreateAuthenticationTokenUsingOPTIONS(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) CreateAuthenticationTokenUsingPOST(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) DataMutationUsingDELETE(ctx *gin.Context, params *data_controller.DataMutationUsingDELETEParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) DataMutationUsingOPTIONS(ctx *gin.Context, params *data_controller.DataMutationUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) DataMutationUsingPOST(ctx *gin.Context, params *data_controller.DataMutationUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) DataMutationUsingPUT(ctx *gin.Context, params *data_controller.DataMutationUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) DataQueryUsingGET(ctx *gin.Context, params *data_controller.DataQueryUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) EditDataSourceUsingPUT(ctx *gin.Context, params *data_source_controller.EditDataSourceUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) EditEntityUsingPUT(ctx *gin.Context, params *schema_controller.EditEntityUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) EditFieldUsingPUT(ctx *gin.Context, params *schema_controller.EditFieldUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) EditFieldUsingPUT1(ctx *gin.Context, params *user_controller.EditFieldUsingPUT1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) EditPermissionUsingPUT(ctx *gin.Context, params *permission_controller.EditPermissionUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) EditRoleUsingPUT(ctx *gin.Context, params *role_controller.EditRoleUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) FindAllFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindAllFieldsUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) FindEntityFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindEntityFieldsUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) FindOneFieldUsingGET(ctx *gin.Context, params *schema_controller.FindOneFieldUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) FindOneUsingGET(ctx *gin.Context, params *data_controller.FindOneUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) FindRoleUsingGET(ctx *gin.Context, params *data_source_controller.FindRoleUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) FindRoleUsingGET1(ctx *gin.Context, params *role_controller.FindRoleUsingGET1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) FindSchemaEntityGET(ctx *gin.Context, params *schema_controller.FindSchemaEntityGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) FindUserUsingGET(ctx *gin.Context, params *permission_controller.FindUserUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) FindUserUsingGET1(ctx *gin.Context, params *user_controller.FindUserUsingGET1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) GetAuthenticatedUserUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) GetRoles(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) GetSchemasUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) ListUsersGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) ListUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) ListUsingGET1(ctx *gin.Context, params *permission_controller.ListUsingGET1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) RefreshAndGetAuthenticationTokenUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) ResetCurrentDsUsingPUT(ctx *gin.Context, params *schema_controller.ResetCurrentDsUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityServer) SyncSchemasUsingPUT(ctx *gin.Context, params *schema_controller.SyncSchemasUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}
