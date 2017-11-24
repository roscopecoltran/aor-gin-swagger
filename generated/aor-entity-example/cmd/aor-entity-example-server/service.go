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
	"github.com/k0kubun/pp"
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

func (s *AorEntityExampleServer) AddApplyUsingOPTIONS(ctx *gin.Context, params *apply_controller.AddApplyUsingOPTIONSParams) *api.Response {

	// debug - tmpl vars

	// Method: OPTIONS
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [apply-controller]

	// Package: apply_controller
	// Path: /apply
	// RootPackage: operations
	// Authorized: true

	// Name: addApplyUsingOPTIONS
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "OPTIONS"? true

	// ExtractString "add" "Using": Apply
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Apply
	// ModelsPackage: models
	// Name: apply

	pp.Println("params: ", params)
	response := make(map[string]string)
	var apply models.Apply
	var queryRes []models.Apply
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddApplyUsingPOST(ctx *gin.Context, params *apply_controller.AddApplyUsingPOSTParams) *api.Response {

	// debug - tmpl vars

	// Method: POST
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [apply-controller]

	// Package: apply_controller
	// Path: /apply
	// RootPackage: operations
	// Authorized: true

	// Name: addApplyUsingPOST
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "POST"? true

	// ExtractString "add" "Using": Apply
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Apply
	// ModelsPackage: models
	// Name: apply

	pp.Println("params: ", params)
	response := make(map[string]string)
	var apply models.Apply
	var queryRes []models.Apply
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddDataSourceUsingOPTIONS(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingOPTIONSParams) *api.Response {

	// debug - tmpl vars

	// Method: OPTIONS
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [data-source-controller]

	// Package: data_source_controller
	// Path: /datasource/_datasource
	// RootPackage: operations
	// Authorized: true

	// Name: addDataSourceUsingOPTIONS
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "OPTIONS"? true

	// ExtractString "add" "Using": DataSource
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: DataSource
	// ModelsPackage: models
	// Name: dataSource

	pp.Println("params: ", params)
	response := make(map[string]string)
	var dataSource models.DataSource
	var queryRes []models.DataSource
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddDataSourceUsingPOST(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingPOSTParams) *api.Response {

	// debug - tmpl vars

	// Method: POST
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [data-source-controller]

	// Package: data_source_controller
	// Path: /datasource/_datasource
	// RootPackage: operations
	// Authorized: true

	// Name: addDataSourceUsingPOST
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "POST"? true

	// ExtractString "add" "Using": DataSource
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: DataSource
	// ModelsPackage: models
	// Name: dataSource

	pp.Println("params: ", params)
	response := make(map[string]string)
	var dataSource models.DataSource
	var queryRes []models.DataSource
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddEntityUsingOPTIONS(ctx *gin.Context, params *schema_controller.AddEntityUsingOPTIONSParams) *api.Response {

	// debug - tmpl vars

	// Method: OPTIONS
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/_entitys
	// RootPackage: operations
	// Authorized: true

	// Name: addEntityUsingOPTIONS
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "OPTIONS"? true

	// ExtractString "add" "Using": Entity
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Entity
	// ModelsPackage: models
	// Name: entity

	pp.Println("params: ", params)
	response := make(map[string]string)
	var entity models.Entity
	var queryRes []models.Entity
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddEntityUsingPOST(ctx *gin.Context, params *schema_controller.AddEntityUsingPOSTParams) *api.Response {

	// debug - tmpl vars

	// Method: POST
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/_entitys
	// RootPackage: operations
	// Authorized: true

	// Name: addEntityUsingPOST
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "POST"? true

	// ExtractString "add" "Using": Entity
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Entity
	// ModelsPackage: models
	// Name: entity

	pp.Println("params: ", params)
	response := make(map[string]string)
	var entity models.Entity
	var queryRes []models.Entity
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddFieldUsingOPTIONS(ctx *gin.Context, params *schema_controller.AddFieldUsingOPTIONSParams) *api.Response {

	// debug - tmpl vars

	// Method: OPTIONS
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/_fields
	// RootPackage: operations
	// Authorized: true

	// Name: addFieldUsingOPTIONS
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "OPTIONS"? true

	// ExtractString "add" "Using": Field
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Field
	// ModelsPackage: models
	// Name: field

	pp.Println("params: ", params)
	response := make(map[string]string)
	var field models.Field
	var queryRes []models.Field
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddFieldUsingPOST(ctx *gin.Context, params *schema_controller.AddFieldUsingPOSTParams) *api.Response {

	// debug - tmpl vars

	// Method: POST
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/_fields
	// RootPackage: operations
	// Authorized: true

	// Name: addFieldUsingPOST
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "POST"? true

	// ExtractString "add" "Using": Field
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Field
	// ModelsPackage: models
	// Name: field

	pp.Println("params: ", params)
	response := make(map[string]string)
	var field models.Field
	var queryRes []models.Field
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddPermissionUsingOPTIONS(ctx *gin.Context, params *permission_controller.AddPermissionUsingOPTIONSParams) *api.Response {

	// debug - tmpl vars

	// Method: OPTIONS
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [permission-controller]

	// Package: permission_controller
	// Path: /permission/_permission
	// RootPackage: operations
	// Authorized: true

	// Name: addPermissionUsingOPTIONS
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "OPTIONS"? true

	// ExtractString "add" "Using": Permission
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Permission
	// ModelsPackage: models
	// Name: permission

	pp.Println("params: ", params)
	response := make(map[string]string)
	var permission models.Permission
	var queryRes []models.Permission
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddPermissionUsingPOST(ctx *gin.Context, params *permission_controller.AddPermissionUsingPOSTParams) *api.Response {

	// debug - tmpl vars

	// Method: POST
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [permission-controller]

	// Package: permission_controller
	// Path: /permission/_permission
	// RootPackage: operations
	// Authorized: true

	// Name: addPermissionUsingPOST
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "POST"? true

	// ExtractString "add" "Using": Permission
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Permission
	// ModelsPackage: models
	// Name: permission

	pp.Println("params: ", params)
	response := make(map[string]string)
	var permission models.Permission
	var queryRes []models.Permission
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddRoleUsingOPTIONS(ctx *gin.Context, params *role_controller.AddRoleUsingOPTIONSParams) *api.Response {

	// debug - tmpl vars

	// Method: OPTIONS
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [role-controller]

	// Package: role_controller
	// Path: /role/_roles
	// RootPackage: operations
	// Authorized: true

	// Name: addRoleUsingOPTIONS
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "OPTIONS"? true

	// ExtractString "add" "Using": Role
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Role
	// ModelsPackage: models
	// Name: role

	pp.Println("params: ", params)
	response := make(map[string]string)
	var role models.Role
	var queryRes []models.Role
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddRoleUsingPOST(ctx *gin.Context, params *role_controller.AddRoleUsingPOSTParams) *api.Response {

	// debug - tmpl vars

	// Method: POST
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [role-controller]

	// Package: role_controller
	// Path: /role/_roles
	// RootPackage: operations
	// Authorized: true

	// Name: addRoleUsingPOST
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "POST"? true

	// ExtractString "add" "Using": Role
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Role
	// ModelsPackage: models
	// Name: role

	pp.Println("params: ", params)
	response := make(map[string]string)
	var role models.Role
	var queryRes []models.Role
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddUserUsingOPTIONS(ctx *gin.Context, params *user_controller.AddUserUsingOPTIONSParams) *api.Response {

	// debug - tmpl vars

	// Method: OPTIONS
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [user-controller]

	// Package: user_controller
	// Path: /user/_users
	// RootPackage: operations
	// Authorized: true

	// Name: addUserUsingOPTIONS
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "OPTIONS"? true

	// ExtractString "add" "Using": User
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: User
	// ModelsPackage: models
	// Name: user

	pp.Println("params: ", params)
	response := make(map[string]string)
	var user models.User
	var queryRes []models.User
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) AddUserUsingPOST(ctx *gin.Context, params *user_controller.AddUserUsingPOSTParams) *api.Response {

	// debug - tmpl vars

	// Method: POST
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [user-controller]

	// Package: user_controller
	// Path: /user/_users
	// RootPackage: operations
	// Authorized: true

	// Name: addUserUsingPOST
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? true
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? true
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "POST"? true

	// ExtractString "add" "Using": User
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: User
	// ModelsPackage: models
	// Name: user

	pp.Println("params: ", params)
	response := make(map[string]string)
	var user models.User
	var queryRes []models.User
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) CreateAuthenticationTokenUsingOPTIONS(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingOPTIONSParams) *api.Response {

	// debug - tmpl vars

	// Method: OPTIONS
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [authentication-rest-controller]

	// Package: authentication_rest_controller
	// Path: /auth
	// RootPackage: operations
	// Authorized: true

	// Name: createAuthenticationTokenUsingOPTIONS
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? true
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? true
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "OPTIONS"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: AuthenticationRequest
	// ModelsPackage: models
	// Name: authenticationRequest

	pp.Println("params: ", params)
	response := make(map[string]string)
	var authenticationRequest models.AuthenticationRequest
	var queryRes []models.AuthenticationRequest
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) CreateAuthenticationTokenUsingPOST(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingPOSTParams) *api.Response {

	// debug - tmpl vars

	// Method: POST
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [authentication-rest-controller]

	// Package: authentication_rest_controller
	// Path: /auth
	// RootPackage: operations
	// Authorized: true

	// Name: createAuthenticationTokenUsingPOST
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? true
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? true
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "POST"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: AuthenticationRequest
	// ModelsPackage: models
	// Name: authenticationRequest

	pp.Println("params: ", params)
	response := make(map[string]string)
	var authenticationRequest models.AuthenticationRequest
	var queryRes []models.AuthenticationRequest
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) DataMutationUsingDELETE(ctx *gin.Context, params *data_controller.DataMutationUsingDELETEParams) *api.Response {

	// debug - tmpl vars

	// Method: DELETE
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [data-controller]

	// Package: data_controller
	// Path: /api/{entity}/{id}
	// RootPackage: operations
	// Authorized: true

	// Name: dataMutationUsingDELETE
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "DELETE"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Entity
	// ModelsPackage: models
	// Name: entity

	pp.Println("params: ", params)
	response := make(map[string]string)
	var entity models.Entity
	var queryRes []models.Entity
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: ID
	// ModelsPackage: models
	// Name: id

	pp.Println("params: ", params)
	response := make(map[string]string)
	var id models.ID
	var queryRes []models.ID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} Entity entity models "entity" o.Entity i o path  entity   <nil>  <nil> <nil> <nil> <nil> false  false map[]} {{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} ID id models "id" o.ID i o path  id   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) DataMutationUsingGET(ctx *gin.Context, params *data_controller.DataMutationUsingGETParams) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [data-controller]

	// Package: data_controller
	// Path: /api/{entity}/{id}
	// RootPackage: operations
	// Authorized: true

	// Name: dataMutationUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Entity
	// ModelsPackage: models
	// Name: entity

	pp.Println("params: ", params)
	response := make(map[string]string)
	var entity models.Entity
	var queryRes []models.Entity
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: ID
	// ModelsPackage: models
	// Name: id

	pp.Println("params: ", params)
	response := make(map[string]string)
	var id models.ID
	var queryRes []models.ID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} Entity entity models "entity" o.Entity i o path  entity   <nil>  <nil> <nil> <nil> <nil> false  false map[]} {{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} ID id models "id" o.ID i o path  id   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) DataMutationUsingOPTIONS(ctx *gin.Context, params *data_controller.DataMutationUsingOPTIONSParams) *api.Response {

	// debug - tmpl vars

	// Method: OPTIONS
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [data-controller]

	// Package: data_controller
	// Path: /api/{entity}
	// RootPackage: operations
	// Authorized: true

	// Name: dataMutationUsingOPTIONS
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "OPTIONS"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: AllRequestParams
	// ModelsPackage: models
	// Name: allRequestParams

	pp.Println("params: ", params)
	response := make(map[string]string)
	var allRequestParams models.AllRequestParams
	var queryRes []models.AllRequestParams
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: Entity
	// ModelsPackage: models
	// Name: entity

	pp.Println("params: ", params)
	response := make(map[string]string)
	var entity models.Entity
	var queryRes []models.Entity
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} Entity entity models "entity" o.Entity i o path  entity   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) DataMutationUsingPOST(ctx *gin.Context, params *data_controller.DataMutationUsingPOSTParams) *api.Response {

	// debug - tmpl vars

	// Method: POST
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [data-controller]

	// Package: data_controller
	// Path: /api/{entity}
	// RootPackage: operations
	// Authorized: true

	// Name: dataMutationUsingPOST
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "POST"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: AllRequestParams
	// ModelsPackage: models
	// Name: allRequestParams

	pp.Println("params: ", params)
	response := make(map[string]string)
	var allRequestParams models.AllRequestParams
	var queryRes []models.AllRequestParams
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: Entity
	// ModelsPackage: models
	// Name: entity

	pp.Println("params: ", params)
	response := make(map[string]string)
	var entity models.Entity
	var queryRes []models.Entity
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} Entity entity models "entity" o.Entity i o path  entity   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) DataMutationUsingPUT(ctx *gin.Context, params *data_controller.DataMutationUsingPUTParams) *api.Response {

	// debug - tmpl vars

	// Method: PUT
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [data-controller]

	// Package: data_controller
	// Path: /api/{entity}/{id}
	// RootPackage: operations
	// Authorized: true

	// Name: dataMutationUsingPUT
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "PUT"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: AllRequestParams
	// ModelsPackage: models
	// Name: allRequestParams

	pp.Println("params: ", params)
	response := make(map[string]string)
	var allRequestParams models.AllRequestParams
	var queryRes []models.AllRequestParams
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: Entity
	// ModelsPackage: models
	// Name: entity

	pp.Println("params: ", params)
	response := make(map[string]string)
	var entity models.Entity
	var queryRes []models.Entity
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: ID
	// ModelsPackage: models
	// Name: id

	pp.Println("params: ", params)
	response := make(map[string]string)
	var id models.ID
	var queryRes []models.ID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} Entity entity models "entity" o.Entity i o path  entity   <nil>  <nil> <nil> <nil> <nil> false  false map[]} {{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} ID id models "id" o.ID i o path  id   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditDataSourceUsingPUT(ctx *gin.Context, params *data_source_controller.EditDataSourceUsingPUTParams) *api.Response {

	// debug - tmpl vars

	// Method: PUT
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [data-source-controller]

	// Package: data_source_controller
	// Path: /datasource/_datasource/put/{id}
	// RootPackage: operations
	// Authorized: true

	// Name: editDataSourceUsingPUT
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? true
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? true
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "PUT"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: DataSource
	// ModelsPackage: models
	// Name: dataSource

	pp.Println("params: ", params)
	response := make(map[string]string)
	var dataSource models.DataSource
	var queryRes []models.DataSource
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: ID
	// ModelsPackage: models
	// Name: id

	pp.Println("params: ", params)
	response := make(map[string]string)
	var id models.ID
	var queryRes []models.ID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} ID id models "id" o.ID i o path  id   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditEntityUsingPUT(ctx *gin.Context, params *schema_controller.EditEntityUsingPUTParams) *api.Response {

	// debug - tmpl vars

	// Method: PUT
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/_entitys/put/{id}
	// RootPackage: operations
	// Authorized: true

	// Name: editEntityUsingPUT
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? true
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? true
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "PUT"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Entity
	// ModelsPackage: models
	// Name: entity

	pp.Println("params: ", params)
	response := make(map[string]string)
	var entity models.Entity
	var queryRes []models.Entity
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: ID
	// ModelsPackage: models
	// Name: id

	pp.Println("params: ", params)
	response := make(map[string]string)
	var id models.ID
	var queryRes []models.ID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} ID id models "id" o.ID i o path  id   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditPermissionUsingPUT(ctx *gin.Context, params *permission_controller.EditPermissionUsingPUTParams) *api.Response {

	// debug - tmpl vars

	// Method: PUT
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [permission-controller]

	// Package: permission_controller
	// Path: /permission/_permission/{id}
	// RootPackage: operations
	// Authorized: true

	// Name: editPermissionUsingPUT
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? true
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? true
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "PUT"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: ID
	// ModelsPackage: models
	// Name: id

	pp.Println("params: ", params)
	response := make(map[string]string)
	var id models.ID
	var queryRes []models.ID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: Permission
	// ModelsPackage: models
	// Name: permission

	pp.Println("params: ", params)
	response := make(map[string]string)
	var permission models.Permission
	var queryRes []models.Permission
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} ID id models "id" o.ID i o path  id   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditRoleUsingPUT(ctx *gin.Context, params *role_controller.EditRoleUsingPUTParams) *api.Response {

	// debug - tmpl vars

	// Method: PUT
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [role-controller]

	// Package: role_controller
	// Path: /role/_roles/put/{id}
	// RootPackage: operations
	// Authorized: true

	// Name: editRoleUsingPUT
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? true
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? true
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "PUT"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: ID
	// ModelsPackage: models
	// Name: id

	pp.Println("params: ", params)
	response := make(map[string]string)
	var id models.ID
	var queryRes []models.ID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: Role
	// ModelsPackage: models
	// Name: role

	pp.Println("params: ", params)
	response := make(map[string]string)
	var role models.Role
	var queryRes []models.Role
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} ID id models "id" o.ID i o path  id   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditSchemaFieldUsingPUT(ctx *gin.Context, params *schema_controller.EditSchemaFieldUsingPUTParams) *api.Response {

	// debug - tmpl vars

	// Method: PUT
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/_fields/put/{id}
	// RootPackage: operations
	// Authorized: true

	// Name: editSchemaFieldUsingPUT
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? true
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? true
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "PUT"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Field
	// ModelsPackage: models
	// Name: field

	pp.Println("params: ", params)
	response := make(map[string]string)
	var field models.Field
	var queryRes []models.Field
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: ID
	// ModelsPackage: models
	// Name: id

	pp.Println("params: ", params)
	response := make(map[string]string)
	var id models.ID
	var queryRes []models.ID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} ID id models "id" o.ID i o path  id   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) EditUserFieldUsingPUT(ctx *gin.Context, params *user_controller.EditUserFieldUsingPUTParams) *api.Response {

	// debug - tmpl vars

	// Method: PUT
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [user-controller]

	// Package: user_controller
	// Path: /user/_users/put/{id}
	// RootPackage: operations
	// Authorized: true

	// Name: editUserFieldUsingPUT
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? true
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? true
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "PUT"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: ID
	// ModelsPackage: models
	// Name: id

	pp.Println("params: ", params)
	response := make(map[string]string)
	var id models.ID
	var queryRes []models.ID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// ID: User
	// ModelsPackage: models
	// Name: user

	pp.Println("params: ", params)
	response := make(map[string]string)
	var user models.User
	var queryRes []models.User
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} ID id models "id" o.ID i o path  id   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindDataSourceUsingGET(ctx *gin.Context, params *data_source_controller.FindDataSourceUsingGETParams) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [data-source-controller]

	// Package: data_source_controller
	// Path: /datasource/_datasource/{datasourceId}
	// RootPackage: operations
	// Authorized: true

	// Name: findDataSourceUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: DatasourceID
	// ModelsPackage: models
	// Name: datasourceId

	pp.Println("params: ", params)
	response := make(map[string]string)
	var datasourceId models.DatasourceID
	var queryRes []models.DatasourceID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} DatasourceID datasourceId models "datasourceId" o.DatasourceID i o path  datasourceId   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindEntityFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindEntityFieldsUsingGETParams) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: true
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/fields
	// RootPackage: operations
	// Authorized: true

	// Name: findEntityFieldsUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: EntityName
	// ModelsPackage: models
	// Name: entityName

	pp.Println("params: ", params)
	response := make(map[string]string)
	var entityName models.EntityName
	var queryRes []models.EntityName
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} EntityName entityName models "entityName" o.EntityName i o query  entityName   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindFieldUsingGET(ctx *gin.Context, params *schema_controller.FindFieldUsingGETParams) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/_fields/{fid}
	// RootPackage: operations
	// Authorized: true

	// Name: findFieldUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Fid
	// ModelsPackage: models
	// Name: fid

	pp.Println("params: ", params)
	response := make(map[string]string)
	var fid models.Fid
	var queryRes []models.Fid
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} Fid fid models "fid" o.Fid i o path  fid   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindFieldsUsingGETParams) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: true
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/_fields
	// RootPackage: operations
	// Authorized: true

	// Name: findFieldsUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Eid
	// ModelsPackage: models
	// Name: eid

	pp.Println("params: ", params)
	response := make(map[string]string)
	var eid models.Eid
	var queryRes []models.Eid
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} Eid eid models "eid" o.Eid i o query  eid   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindPermissionUsingGET(ctx *gin.Context, params *permission_controller.FindPermissionUsingGETParams) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [permission-controller]

	// Package: permission_controller
	// Path: /permission/_permission/{id}
	// RootPackage: operations
	// Authorized: true

	// Name: findPermissionUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: ID
	// ModelsPackage: models
	// Name: id

	pp.Println("params: ", params)
	response := make(map[string]string)
	var id models.ID
	var queryRes []models.ID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} ID id models "id" o.ID i o path  id   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindRoleUsingGET(ctx *gin.Context, params *role_controller.FindRoleUsingGETParams) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [role-controller]

	// Package: role_controller
	// Path: /role/_roles/{roleId}
	// RootPackage: operations
	// Authorized: true

	// Name: findRoleUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: RoleID
	// ModelsPackage: models
	// Name: roleId

	pp.Println("params: ", params)
	response := make(map[string]string)
	var roleId models.RoleID
	var queryRes []models.RoleID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} RoleID roleId models "roleId" o.RoleID i o path  roleId   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindSchemaEntityGET(ctx *gin.Context, params *schema_controller.FindSchemaEntityGETParams) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/_entitys/{eid}
	// RootPackage: operations
	// Authorized: true

	// Name: findSchemaEntityGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: Eid
	// ModelsPackage: models
	// Name: eid

	pp.Println("params: ", params)
	response := make(map[string]string)
	var eid models.Eid
	var queryRes []models.Eid
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} Eid eid models "eid" o.Eid i o path  eid   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) FindUserUsingGET(ctx *gin.Context, params *user_controller.FindUserUsingGETParams) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [user-controller]

	// Package: user_controller
	// Path: /user/_users/{userId}
	// RootPackage: operations
	// Authorized: true

	// Name: findUserUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: UserID
	// ModelsPackage: models
	// Name: userId

	pp.Println("params: ", params)
	response := make(map[string]string)
	var userId models.UserID
	var queryRes []models.UserID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} UserID userId models "userId" o.UserID i o path  userId   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) GetAuthenticatedUserUsingGET(ctx *gin.Context) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [user-controller]

	// Package: user_controller
	// Path: /user/me
	// RootPackage: operations
	// Authorized: true

	// Name: getAuthenticatedUserUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? true

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? true
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using": AuthenticatedUser

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) GetDataMutationUsingGET(ctx *gin.Context, params *data_controller.GetDataMutationUsingGETParams) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [data-controller]

	// Package: data_controller
	// Path: /api/{entity}
	// RootPackage: operations
	// Authorized: true

	// Name: getDataMutationUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? true

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? true
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using": DataMutation

	// ID: Entity
	// ModelsPackage: models
	// Name: entity

	pp.Println("params: ", params)
	response := make(map[string]string)
	var entity models.Entity
	var queryRes []models.Entity
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} Entity entity models "entity" o.Entity i o path  entity   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) GetRoles(ctx *gin.Context) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [role-controller]

	// Package: role_controller
	// Path: /role/_roles
	// RootPackage: operations
	// Authorized: true

	// Name: getRoles
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? true

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? true
	// HasPrefix "delete"? false

	// HasSuffix "GET"? false

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) GetSchemasUsingGET(ctx *gin.Context) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/_entitys
	// RootPackage: operations
	// Authorized: true

	// Name: getSchemasUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? true

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? true
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using": Schemas

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) ListDataSourceUsingGET(ctx *gin.Context) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [data-source-controller]

	// Package: data_source_controller
	// Path: /datasource/_datasource
	// RootPackage: operations
	// Authorized: true

	// Name: listDataSourceUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? true
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? true
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using": DataSource
	// ExtractString "get" "Using":

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) ListPermissionUsingGET(ctx *gin.Context, params *permission_controller.ListPermissionUsingGETParams) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: true
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [permission-controller]

	// Package: permission_controller
	// Path: /permission/_permission
	// RootPackage: operations
	// Authorized: true

	// Name: listPermissionUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? true
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? true
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using": Permission
	// ExtractString "get" "Using":

	// ID: RoleID
	// ModelsPackage: models
	// Name: roleId

	pp.Println("params: ", params)
	response := make(map[string]string)
	var roleId models.RoleID
	var queryRes []models.RoleID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} RoleID roleId models "roleId" o.RoleID i o query  roleId   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) ListUsersGET(ctx *gin.Context) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [user-controller]

	// Package: user_controller
	// Path: /user/_users
	// RootPackage: operations
	// Authorized: true

	// Name: listUsersGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? true
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? true
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) RefreshAndGetAuthenticationTokenUsingGET(ctx *gin.Context) *api.Response {

	// debug - tmpl vars

	// Method: GET
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [authentication-rest-controller]

	// Package: authentication_rest_controller
	// Path: /refresh
	// RootPackage: operations
	// Authorized: true

	// Name: refreshAndGetAuthenticationTokenUsingGET
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? true
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? true
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "GET"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// QueryParams: []
	// PathParams: []
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) ResetCurrentDsUsingPUT(ctx *gin.Context, params *schema_controller.ResetCurrentDsUsingPUTParams) *api.Response {

	// debug - tmpl vars

	// Method: PUT
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/resetCurrentDs/{dataSourceId}
	// RootPackage: operations
	// Authorized: true

	// Name: resetCurrentDsUsingPUT
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? false
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? false
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "PUT"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: DataSourceID
	// ModelsPackage: models
	// Name: dataSourceId

	pp.Println("params: ", params)
	response := make(map[string]string)
	var dataSourceId models.DataSourceID
	var queryRes []models.DataSourceID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} DataSourceID dataSourceId models "dataSourceId" o.DataSourceID i o path  dataSourceId   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AorEntityExampleServer) SyncSchemasUsingPUT(ctx *gin.Context, params *schema_controller.SyncSchemasUsingPUTParams) *api.Response {

	// debug - tmpl vars

	// Method: PUT
	// HasQueryParams: false
	// HasFormParams: false
	// HasFormValueParams: false
	// HasFileParams: false
	// HasStreamingResponse: false
	// WithContext: false
	// HasFileParams: false
	// Tags: [schema-controller]

	// Package: schema_controller
	// Path: /schemas/sync/{dataSourceId}
	// RootPackage: operations
	// Authorized: true

	// Name: syncSchemasUsingPUT
	// StrContains "findAll"? false
	// StrContains "findOne"? false
	// StrContains "WithID"? false

	// StrContains "refresh"? false
	// StrContains "sync"? true
	// StrContains "update"? false
	// StrContains "edit"? false
	// StrContains "create"? false
	// StrContains "add"? false
	// StrContains "list"? false
	// StrContains "get"? false

	// HasPrefix "refresh"? false
	// HasPrefix "update"? false
	// HasPrefix "edit"? false
	// HasPrefix "sync"? true
	// HasPrefix "create"? false
	// HasPrefix "add"? false
	// HasPrefix "list"? false
	// HasPrefix "get"? false
	// HasPrefix "delete"? false

	// HasSuffix "PUT"? true

	// ExtractString "add" "Using":
	// ExtractString "list" "Using":
	// ExtractString "get" "Using":

	// ID: DataSourceID
	// ModelsPackage: models
	// Name: dataSourceId

	pp.Println("params: ", params)
	response := make(map[string]string)
	var dataSourceId models.DataSourceID
	var queryRes []models.DataSourceID
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? AND db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}

	// QueryParams: []
	// PathParams: [{{false false false false true false false false false false false false false false false string    string  map[] <nil>} {true <nil> <nil>  <nil> <nil> <nil> false false [] [] false <nil> <nil> false false false false false} DataSourceID dataSourceId models "dataSourceId" o.DataSourceID i o path  dataSourceId   <nil>  <nil> <nil> <nil> <nil> false  false map[]}]
	// HeaderParams: []
	// FormParams: []

	/*
		db := ctx.MustGet("db").(*gorm.DB)
		var results []models.User
		response := make(map[string]interface{})
		if err := db.Select("*").Find(&results).Error; err != nil {
			response["error"] = err.Error()
			return &api.Response{Code: 400, Body: response}
		}
		response["results"] = results
		response["status"] = "success"
		return &api.Response{Code: http.StatusOK, Body: response}
	*/

	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}
