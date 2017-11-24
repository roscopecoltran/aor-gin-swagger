// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Entity entity
// swagger:model Entity
type Entity struct {

	// crud
	Crud []string `json:"crud"`

	// fields
	Fields EntityFields `json:"fields"`

	// id
	ID string `json:"id,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// redirect
	Redirect string `json:"redirect,omitempty"`

	// show in menu
	ShowInMenu bool `json:"showInMenu,omitempty"`
}

// Validate validates this entity
func (m *Entity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrud(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRedirect(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var entityCrudItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["c","r","u","d"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		entityCrudItemsEnum = append(entityCrudItemsEnum, v)
	}
}

func (m *Entity) validateCrudItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, entityCrudItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *Entity) validateCrud(formats strfmt.Registry) error {

	if swag.IsZero(m.Crud) { // not required
		return nil
	}

	for i := 0; i < len(m.Crud); i++ {

		// value enum
		if err := m.validateCrudItemsEnum("crud"+"."+strconv.Itoa(i), "body", m.Crud[i]); err != nil {
			return err
		}

	}

	return nil
}

var entityTypeRedirectPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["edit","show","list"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		entityTypeRedirectPropEnum = append(entityTypeRedirectPropEnum, v)
	}
}

const (
	// EntityRedirectEdit captures enum value "edit"
	EntityRedirectEdit string = "edit"
	// EntityRedirectShow captures enum value "show"
	EntityRedirectShow string = "show"
	// EntityRedirectList captures enum value "list"
	EntityRedirectList string = "list"
)

// prop value enum
func (m *Entity) validateRedirectEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, entityTypeRedirectPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Entity) validateRedirect(formats strfmt.Registry) error {

	if swag.IsZero(m.Redirect) { // not required
		return nil
	}

	// value enum
	if err := m.validateRedirectEnum("redirect", "body", m.Redirect); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Entity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Entity) UnmarshalBinary(b []byte) error {
	var res Entity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
