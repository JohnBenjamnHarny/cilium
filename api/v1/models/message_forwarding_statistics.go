// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MessageForwardingStatistics Statistics of a message forwarding entity
//
// swagger:model MessageForwardingStatistics
type MessageForwardingStatistics struct {

	// Number of messages denied
	Denied int64 `json:"denied,omitempty"`

	// Number of errors while parsing messages
	Error int64 `json:"error,omitempty"`

	// Number of messages forwarded
	Forwarded int64 `json:"forwarded,omitempty"`

	// Number of messages received
	Received int64 `json:"received,omitempty"`
}

// Validate validates this message forwarding statistics
func (m *MessageForwardingStatistics) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MessageForwardingStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageForwardingStatistics) UnmarshalBinary(b []byte) error {
	var res MessageForwardingStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
