package data_controller

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

// BusinessLogicDataMutationUsingDELETE executes the core logic of the related
// route endpoint.
func BusinessLogicDataMutationUsingDELETE(f func(ctx *gin.Context, params *DataMutationUsingDELETEParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// generate params from request
		params := &DataMutationUsingDELETEParams{}
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

// DataMutationUsingDELETEParams contains all the bound params for the data mutation using d e l e t e operation
// typically these are obtained from a http.Request
//
// swagger:parameters dataMutationUsingDELETE
type DataMutationUsingDELETEParams struct {

	/*entity
	  Required: true
	  In: path
	*/
	Entity string
	/*id
	  Required: true
	  In: path
	*/
	ID string
}

// DataMutationUsingDELETEParamsFromCtx gets the params struct from the gin context.
func DataMutationUsingDELETEParamsFromCtx(ctx *gin.Context) *DataMutationUsingDELETEParams {
	params, _ := ctx.Get("params")
	return params.(*DataMutationUsingDELETEParams)
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DataMutationUsingDELETEParams) bindRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	rEntity := []string{ctx.Param("entity")}
	if err := o.bindEntity(rEntity, true, formats); err != nil {
		res = append(res, err)
	}

	rID := []string{ctx.Param("id")}
	if err := o.bindID(rID, true, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DataMutationUsingDELETEParams) bindEntity(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Entity = raw

	return nil
}

func (o *DataMutationUsingDELETEParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}

// vim: ft=go