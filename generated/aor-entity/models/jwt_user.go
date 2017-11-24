// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// JwtUser jwt user
// swagger:model JwtUser
type JwtUser struct {

	// authorities
	Authorities JwtUserAuthorities `json:"authorities"`

	// email
	Email string `json:"email,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this jwt user
func (m *JwtUser) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *JwtUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JwtUser) UnmarshalBinary(b []byte) error {
	var res JwtUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}