// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatacenterSpecPacket DatacenterSpecPacket describes a Packet datacenter.
//
// swagger:model DatacenterSpecPacket
type DatacenterSpecPacket struct {

	// The list of enabled facilities, for example "ams1", for a full list of available
	// facilities see https://metal.equinix.com/developers/docs/locations/facilities/
	Facilities []string `json:"facilities"`

	// Metros are facilities that are grouped together geographically and share capacity
	// and networking features, see https://metal.equinix.com/developers/docs/locations/metros/
	Metro string `json:"metro,omitempty"`
}

// Validate validates this datacenter spec packet
func (m *DatacenterSpecPacket) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this datacenter spec packet based on context it is used
func (m *DatacenterSpecPacket) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatacenterSpecPacket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatacenterSpecPacket) UnmarshalBinary(b []byte) error {
	var res DatacenterSpecPacket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
