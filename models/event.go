// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Event Event is a report of an event somewhere in the cluster.
//
// swagger:model Event
type Event struct {

	// Annotations that can be added to the resource
	Annotations map[string]string `json:"annotations,omitempty"`

	// The number of times this event has occurred.
	Count int32 `json:"count,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was created.
	// Format: date-time
	CreationTimestamp strfmt.DateTime `json:"creationTimestamp,omitempty"`

	// DeletionTimestamp is a timestamp representing the server time when this object was deleted.
	// Format: date-time
	DeletionTimestamp strfmt.DateTime `json:"deletionTimestamp,omitempty"`

	// ID unique value that identifies the resource generated by the server. Read-Only.
	ID string `json:"id,omitempty"`

	// The time at which the most recent occurrence of this event was recorded.
	// Format: date-time
	LastTimestamp strfmt.DateTime `json:"lastTimestamp,omitempty"`

	// A human-readable description of the status of this operation.
	Message string `json:"message,omitempty"`

	// Name represents human readable name for the resource
	Name string `json:"name,omitempty"`

	// Type of this event (i.e. normal or warning). New types could be added in the future.
	Type string `json:"type,omitempty"`

	// involved object
	InvolvedObject *ObjectReferenceResource `json:"involvedObject,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvolvedObject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateCreationTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("creationTimestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateDeletionTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletionTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("deletionTimestamp", "body", "date-time", m.DeletionTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateLastTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("lastTimestamp", "body", "date-time", m.LastTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateInvolvedObject(formats strfmt.Registry) error {
	if swag.IsZero(m.InvolvedObject) { // not required
		return nil
	}

	if m.InvolvedObject != nil {
		if err := m.InvolvedObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("involvedObject")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this event based on the context it is used
func (m *Event) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInvolvedObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) contextValidateInvolvedObject(ctx context.Context, formats strfmt.Registry) error {

	if m.InvolvedObject != nil {
		if err := m.InvolvedObject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("involvedObject")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
