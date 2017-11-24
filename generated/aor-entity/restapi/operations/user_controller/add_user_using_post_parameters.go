package user_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/roscopecoltran/gin-swagger/api"

	strfmt "github.com/go-openapi/strfmt"

	// DefaultImports
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/models"
	// Imports
)

// BusinessLogicAddUserUsingPOST executes the core logic of the related
// route endpoint.
func BusinessLogicAddUserUsingPOST(f func(ctx *gin.Context, params *AddUserUsingPOSTParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// generate params from request
		params := &AddUserUsingPOSTParams{}
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

// AddUserUsingPOSTParams contains all the bound params for the add user using p o s t operation
// typically these are obtained from a http.Request
//
// swagger:parameters addUserUsingPOST
type AddUserUsingPOSTParams struct {

	/*user
	  Required: true
	  In: body
	*/
	User *models.User
}

// AddUserUsingPOSTParamsFromCtx gets the params struct from the gin context.
func AddUserUsingPOSTParamsFromCtx(ctx *gin.Context) *AddUserUsingPOSTParams {
	params, _ := ctx.Get("params")
	return params.(*AddUserUsingPOSTParams)
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddUserUsingPOSTParams) bindRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	if runtime.HasBody(ctx.Request) {
		var body models.User
		if err := ctx.BindJSON(&body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("user", "body"))
			} else {
				res = append(res, errors.NewParseError("user", "body", "", err))
			}

		} else {
			if err := body.Validate(formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.User = &body
			}
		}

	} else {
		res = append(res, errors.Required("user", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// vim: ft=go
