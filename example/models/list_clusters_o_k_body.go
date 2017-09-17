// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ListClustersOKBody list clusters o k body
// swagger:model listClustersOKBody

type ListClustersOKBody struct {

	// items
	Items ListClustersOKBodyItems `json:"items"`
}

/* polymorph listClustersOKBody items false */

// Validate validates this list clusters o k body
func (m *ListClustersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ListClustersOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListClustersOKBody) UnmarshalBinary(b []byte) error {
	var res ListClustersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
