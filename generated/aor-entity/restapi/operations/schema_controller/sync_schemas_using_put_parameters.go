package schema_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/roscopecoltran/gin-swagger/api"

	strfmt "github.com/go-openapi/strfmt"
	// DefaultImports
	// Imports
)

// BusinessLogicSyncSchemasUsingPUT executes the core logic of the related
// route endpoint.
func BusinessLogicSyncSchemasUsingPUT(f func(ctx *gin.Context, params *SyncSchemasUsingPUTParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// generate params from request
		params := &SyncSchemasUsingPUTParams{}
		err := params.bindRequest(ctx)
		if err != nil {
			errObj := err.(*errors.CompositeError)
			problem := api.Problem{
				Title:  "Unprocessable Entity.",
				Status: int(errObj.Code()),
				Detail: errObj.Error(),
			}
			ctx.Writer.Header().Set("Content-Type", "application/problem+json")
			ctx.JSON(problem.Status, problem)
			return
		}

		resp := f(ctx, params)
		switch resp.Code {
		case http.StatusNoContent:
			ctx.AbortWithStatus(resp.Code)
		default:
			ctx.JSON(resp.Code, resp.Body)
		}
	}
}

// SyncSchemasUsingPUTParams contains all the bound params for the sync schemas using p u t operation
// typically these are obtained from a http.Request
//
// swagger:parameters syncSchemasUsingPUT
type SyncSchemasUsingPUTParams struct {

	/*dataSourceId
	  Required: true
	  In: path
	*/
	DataSourceID string
}

// SyncSchemasUsingPUTParamsFromCtx gets the params struct from the gin context.
func SyncSchemasUsingPUTParamsFromCtx(ctx *gin.Context) *SyncSchemasUsingPUTParams {
	params, _ := ctx.Get("params")
	return params.(*SyncSchemasUsingPUTParams)
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *SyncSchemasUsingPUTParams) bindRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	rDataSourceID := []string{ctx.Param("dataSourceId")}
	if err := o.bindDataSourceID(rDataSourceID, true, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SyncSchemasUsingPUTParams) bindDataSourceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.DataSourceID = raw

	return nil
}

// vim: ft=go
