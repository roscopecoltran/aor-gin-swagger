// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResponseEntity response entity
// swagger:model ResponseEntity
type ResponseEntity struct {

	// body
	Body interface{} `json:"body,omitempty"`

	// status code
	StatusCode string `json:"statusCode,omitempty"`

	// status code value
	StatusCodeValue int32 `json:"statusCodeValue,omitempty"`
}

// Validate validates this response entity
func (m *ResponseEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatusCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var responseEntityTypeStatusCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["100","101","102","103","200","201","202","203","204","205","206","207","208","226","300","301","302","303","304","305","307","308","400","401","402","403","404","405","406","407","408","409","410","411","412","413","414","415","416","417","418","419","420","421","422","423","424","426","428","429","431","451","500","501","502","503","504","505","506","507","508","509","510","511"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responseEntityTypeStatusCodePropEnum = append(responseEntityTypeStatusCodePropEnum, v)
	}
}

const (
	// ResponseEntityStatusCodeNr100 captures enum value "100"
	ResponseEntityStatusCodeNr100 string = "100"
	// ResponseEntityStatusCodeNr101 captures enum value "101"
	ResponseEntityStatusCodeNr101 string = "101"
	// ResponseEntityStatusCodeNr102 captures enum value "102"
	ResponseEntityStatusCodeNr102 string = "102"
	// ResponseEntityStatusCodeNr103 captures enum value "103"
	ResponseEntityStatusCodeNr103 string = "103"
	// ResponseEntityStatusCodeNr200 captures enum value "200"
	ResponseEntityStatusCodeNr200 string = "200"
	// ResponseEntityStatusCodeNr201 captures enum value "201"
	ResponseEntityStatusCodeNr201 string = "201"
	// ResponseEntityStatusCodeNr202 captures enum value "202"
	ResponseEntityStatusCodeNr202 string = "202"
	// ResponseEntityStatusCodeNr203 captures enum value "203"
	ResponseEntityStatusCodeNr203 string = "203"
	// ResponseEntityStatusCodeNr204 captures enum value "204"
	ResponseEntityStatusCodeNr204 string = "204"
	// ResponseEntityStatusCodeNr205 captures enum value "205"
	ResponseEntityStatusCodeNr205 string = "205"
	// ResponseEntityStatusCodeNr206 captures enum value "206"
	ResponseEntityStatusCodeNr206 string = "206"
	// ResponseEntityStatusCodeNr207 captures enum value "207"
	ResponseEntityStatusCodeNr207 string = "207"
	// ResponseEntityStatusCodeNr208 captures enum value "208"
	ResponseEntityStatusCodeNr208 string = "208"
	// ResponseEntityStatusCodeNr226 captures enum value "226"
	ResponseEntityStatusCodeNr226 string = "226"
	// ResponseEntityStatusCodeNr300 captures enum value "300"
	ResponseEntityStatusCodeNr300 string = "300"
	// ResponseEntityStatusCodeNr301 captures enum value "301"
	ResponseEntityStatusCodeNr301 string = "301"
	// ResponseEntityStatusCodeNr302 captures enum value "302"
	ResponseEntityStatusCodeNr302 string = "302"
	// ResponseEntityStatusCodeNr303 captures enum value "303"
	ResponseEntityStatusCodeNr303 string = "303"
	// ResponseEntityStatusCodeNr304 captures enum value "304"
	ResponseEntityStatusCodeNr304 string = "304"
	// ResponseEntityStatusCodeNr305 captures enum value "305"
	ResponseEntityStatusCodeNr305 string = "305"
	// ResponseEntityStatusCodeNr307 captures enum value "307"
	ResponseEntityStatusCodeNr307 string = "307"
	// ResponseEntityStatusCodeNr308 captures enum value "308"
	ResponseEntityStatusCodeNr308 string = "308"
	// ResponseEntityStatusCodeNr400 captures enum value "400"
	ResponseEntityStatusCodeNr400 string = "400"
	// ResponseEntityStatusCodeNr401 captures enum value "401"
	ResponseEntityStatusCodeNr401 string = "401"
	// ResponseEntityStatusCodeNr402 captures enum value "402"
	ResponseEntityStatusCodeNr402 string = "402"
	// ResponseEntityStatusCodeNr403 captures enum value "403"
	ResponseEntityStatusCodeNr403 string = "403"
	// ResponseEntityStatusCodeNr404 captures enum value "404"
	ResponseEntityStatusCodeNr404 string = "404"
	// ResponseEntityStatusCodeNr405 captures enum value "405"
	ResponseEntityStatusCodeNr405 string = "405"
	// ResponseEntityStatusCodeNr406 captures enum value "406"
	ResponseEntityStatusCodeNr406 string = "406"
	// ResponseEntityStatusCodeNr407 captures enum value "407"
	ResponseEntityStatusCodeNr407 string = "407"
	// ResponseEntityStatusCodeNr408 captures enum value "408"
	ResponseEntityStatusCodeNr408 string = "408"
	// ResponseEntityStatusCodeNr409 captures enum value "409"
	ResponseEntityStatusCodeNr409 string = "409"
	// ResponseEntityStatusCodeNr410 captures enum value "410"
	ResponseEntityStatusCodeNr410 string = "410"
	// ResponseEntityStatusCodeNr411 captures enum value "411"
	ResponseEntityStatusCodeNr411 string = "411"
	// ResponseEntityStatusCodeNr412 captures enum value "412"
	ResponseEntityStatusCodeNr412 string = "412"
	// ResponseEntityStatusCodeNr413 captures enum value "413"
	ResponseEntityStatusCodeNr413 string = "413"
	// ResponseEntityStatusCodeNr414 captures enum value "414"
	ResponseEntityStatusCodeNr414 string = "414"
	// ResponseEntityStatusCodeNr415 captures enum value "415"
	ResponseEntityStatusCodeNr415 string = "415"
	// ResponseEntityStatusCodeNr416 captures enum value "416"
	ResponseEntityStatusCodeNr416 string = "416"
	// ResponseEntityStatusCodeNr417 captures enum value "417"
	ResponseEntityStatusCodeNr417 string = "417"
	// ResponseEntityStatusCodeNr418 captures enum value "418"
	ResponseEntityStatusCodeNr418 string = "418"
	// ResponseEntityStatusCodeNr419 captures enum value "419"
	ResponseEntityStatusCodeNr419 string = "419"
	// ResponseEntityStatusCodeNr420 captures enum value "420"
	ResponseEntityStatusCodeNr420 string = "420"
	// ResponseEntityStatusCodeNr421 captures enum value "421"
	ResponseEntityStatusCodeNr421 string = "421"
	// ResponseEntityStatusCodeNr422 captures enum value "422"
	ResponseEntityStatusCodeNr422 string = "422"
	// ResponseEntityStatusCodeNr423 captures enum value "423"
	ResponseEntityStatusCodeNr423 string = "423"
	// ResponseEntityStatusCodeNr424 captures enum value "424"
	ResponseEntityStatusCodeNr424 string = "424"
	// ResponseEntityStatusCodeNr426 captures enum value "426"
	ResponseEntityStatusCodeNr426 string = "426"
	// ResponseEntityStatusCodeNr428 captures enum value "428"
	ResponseEntityStatusCodeNr428 string = "428"
	// ResponseEntityStatusCodeNr429 captures enum value "429"
	ResponseEntityStatusCodeNr429 string = "429"
	// ResponseEntityStatusCodeNr431 captures enum value "431"
	ResponseEntityStatusCodeNr431 string = "431"
	// ResponseEntityStatusCodeNr451 captures enum value "451"
	ResponseEntityStatusCodeNr451 string = "451"
	// ResponseEntityStatusCodeNr500 captures enum value "500"
	ResponseEntityStatusCodeNr500 string = "500"
	// ResponseEntityStatusCodeNr501 captures enum value "501"
	ResponseEntityStatusCodeNr501 string = "501"
	// ResponseEntityStatusCodeNr502 captures enum value "502"
	ResponseEntityStatusCodeNr502 string = "502"
	// ResponseEntityStatusCodeNr503 captures enum value "503"
	ResponseEntityStatusCodeNr503 string = "503"
	// ResponseEntityStatusCodeNr504 captures enum value "504"
	ResponseEntityStatusCodeNr504 string = "504"
	// ResponseEntityStatusCodeNr505 captures enum value "505"
	ResponseEntityStatusCodeNr505 string = "505"
	// ResponseEntityStatusCodeNr506 captures enum value "506"
	ResponseEntityStatusCodeNr506 string = "506"
	// ResponseEntityStatusCodeNr507 captures enum value "507"
	ResponseEntityStatusCodeNr507 string = "507"
	// ResponseEntityStatusCodeNr508 captures enum value "508"
	ResponseEntityStatusCodeNr508 string = "508"
	// ResponseEntityStatusCodeNr509 captures enum value "509"
	ResponseEntityStatusCodeNr509 string = "509"
	// ResponseEntityStatusCodeNr510 captures enum value "510"
	ResponseEntityStatusCodeNr510 string = "510"
	// ResponseEntityStatusCodeNr511 captures enum value "511"
	ResponseEntityStatusCodeNr511 string = "511"
)

// prop value enum
func (m *ResponseEntity) validateStatusCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, responseEntityTypeStatusCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResponseEntity) validateStatusCode(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusCodeEnum("statusCode", "body", m.StatusCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseEntity) UnmarshalBinary(b []byte) error {
	var res ResponseEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
