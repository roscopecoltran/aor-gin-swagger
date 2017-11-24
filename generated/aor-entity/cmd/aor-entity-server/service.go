package main

import (
	"net/http"

	// "github.com/roscopecoltran/gin-swagger/database"
	"github.com/roscopecoltran/gin-swagger/api"

	// DefaultImports
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/models"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/apply_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/authentication_rest_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/data_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/data_source_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/permission_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/role_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/schema_controller"
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/restapi/operations/user_controller"

	// Imports

	"github.com/jinzhu/gorm"

	// "github.com/gin-gonic/contrib/jwt"

	"github.com/gin-gonic/gin"
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

func (s *addDataSourceUsingOPTIONSServer) AddDataSourceUsingOPTIONS(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *addDataSourceUsingPOSTServer) AddDataSourceUsingPOST(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *addEntityUsingOPTIONSServer) AddEntityUsingOPTIONS(ctx *gin.Context, params *schema_controller.AddEntityUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *addEntityUsingPOSTServer) AddEntityUsingPOST(ctx *gin.Context, params *schema_controller.AddEntityUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *addFieldUsingOPTIONSServer) AddFieldUsingOPTIONS(ctx *gin.Context, params *schema_controller.AddFieldUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *addFieldUsingPOSTServer) AddFieldUsingPOST(ctx *gin.Context, params *schema_controller.AddFieldUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *addPermissionUsingOPTIONSServer) AddPermissionUsingOPTIONS(ctx *gin.Context, params *permission_controller.AddPermissionUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *addPermissionUsingPOSTServer) AddPermissionUsingPOST(ctx *gin.Context, params *permission_controller.AddPermissionUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *addRoleUsingOPTIONSServer) AddRoleUsingOPTIONS(ctx *gin.Context, params *role_controller.AddRoleUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *addRoleUsingPOSTServer) AddRoleUsingPOST(ctx *gin.Context, params *role_controller.AddRoleUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *addUserUsingOPTIONSServer) AddUserUsingOPTIONS(ctx *gin.Context, params *user_controller.AddUserUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *addUserUsingPOSTServer) AddUserUsingPOST(ctx *gin.Context, params *user_controller.AddUserUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *applyUsingOPTIONSServer) ApplyUsingOPTIONS(ctx *gin.Context, params *apply_controller.ApplyUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *applyUsingPOSTServer) ApplyUsingPOST(ctx *gin.Context, params *apply_controller.ApplyUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *createAuthenticationTokenUsingOPTIONSServer) CreateAuthenticationTokenUsingOPTIONS(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *createAuthenticationTokenUsingPOSTServer) CreateAuthenticationTokenUsingPOST(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *dataMutationUsingDELETEServer) DataMutationUsingDELETE(ctx *gin.Context, params *data_controller.DataMutationUsingDELETEParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *dataMutationUsingOPTIONSServer) DataMutationUsingOPTIONS(ctx *gin.Context, params *data_controller.DataMutationUsingOPTIONSParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *dataMutationUsingPOSTServer) DataMutationUsingPOST(ctx *gin.Context, params *data_controller.DataMutationUsingPOSTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *dataMutationUsingPUTServer) DataMutationUsingPUT(ctx *gin.Context, params *data_controller.DataMutationUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *dataQueryUsingGETServer) DataQueryUsingGET(ctx *gin.Context, params *data_controller.DataQueryUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *editDataSourceUsingPUTServer) EditDataSourceUsingPUT(ctx *gin.Context, params *data_source_controller.EditDataSourceUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *editEntityUsingPUTServer) EditEntityUsingPUT(ctx *gin.Context, params *schema_controller.EditEntityUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *editFieldUsingPUTServer) EditFieldUsingPUT(ctx *gin.Context, params *schema_controller.EditFieldUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *editFieldUsingPUT_1Server) EditFieldUsingPUT1(ctx *gin.Context, params *user_controller.EditFieldUsingPUT1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *editPermissionUsingPUTServer) EditPermissionUsingPUT(ctx *gin.Context, params *permission_controller.EditPermissionUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *editRoleUsingPUTServer) EditRoleUsingPUT(ctx *gin.Context, params *role_controller.EditRoleUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *findAllFieldsUsingGETServer) FindAllFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindAllFieldsUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *findEntityFieldsUsingGETServer) FindEntityFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindEntityFieldsUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *findOneFieldUsingGETServer) FindOneFieldUsingGET(ctx *gin.Context, params *schema_controller.FindOneFieldUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *findOneUsingGETServer) FindOneUsingGET(ctx *gin.Context, params *data_controller.FindOneUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *findRoleUsingGETServer) FindRoleUsingGET(ctx *gin.Context, params *data_source_controller.FindRoleUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *findRoleUsingGET_1Server) FindRoleUsingGET1(ctx *gin.Context, params *role_controller.FindRoleUsingGET1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *findSchemaEntityGETServer) FindSchemaEntityGET(ctx *gin.Context, params *schema_controller.FindSchemaEntityGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *findUserUsingGETServer) FindUserUsingGET(ctx *gin.Context, params *permission_controller.FindUserUsingGETParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *findUserUsingGET_1Server) FindUserUsingGET1(ctx *gin.Context, params *user_controller.FindUserUsingGET1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *getAuthenticatedUserUsingGETServer) GetAuthenticatedUserUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *getRolesServer) GetRoles(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *getSchemasUsingGETServer) GetSchemasUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *listUsersGETServer) ListUsersGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *listUsingGETServer) ListUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *listUsingGET_1Server) ListUsingGET1(ctx *gin.Context, params *permission_controller.ListUsingGET1Params) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *refreshAndGetAuthenticationTokenUsingGETServer) RefreshAndGetAuthenticationTokenUsingGET(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *resetCurrentDsUsingPUTServer) ResetCurrentDsUsingPUT(ctx *gin.Context, params *schema_controller.ResetCurrentDsUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *syncSchemasUsingPUTServer) SyncSchemasUsingPUT(ctx *gin.Context, params *schema_controller.SyncSchemasUsingPUTParams) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}
