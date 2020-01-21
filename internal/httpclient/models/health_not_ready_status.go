// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// HealthNotReadyStatus HealthNotReadyStatus health not ready status
// swagger:model healthNotReadyStatus
type HealthNotReadyStatus struct {

	// Errors contains a list of errors that caused the not ready status.
	Errors map[string]string `json:"errors,omitempty"`
}

// Validate validates this health not ready status
func (m *HealthNotReadyStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HealthNotReadyStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthNotReadyStatus) UnmarshalBinary(b []byte) error {
	var res HealthNotReadyStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
