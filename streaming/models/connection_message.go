// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConnectionMessage connection message
//
// swagger:model ConnectionMessage
type ConnectionMessage struct {
	idField int32

	// The connection id
	ConnectionID string `json:"connectionId,omitempty"`
}

// ID gets the id of this subtype
func (m *ConnectionMessage) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *ConnectionMessage) SetID(val int32) {
	m.idField = val
}

// Op gets the op of this subtype
func (m *ConnectionMessage) Op() string {
	return "connection"
}

// SetOp sets the op of this subtype
func (m *ConnectionMessage) SetOp(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ConnectionMessage) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The connection id
		ConnectionID string `json:"connectionId,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ID int32 `json:"id,omitempty"`

		Op string `json:"op,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result ConnectionMessage

	result.idField = base.ID
	if base.Op != result.Op() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid op value: %q", base.Op)
	}

	result.ConnectionID = data.ConnectionID

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ConnectionMessage) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The connection id
		ConnectionID string `json:"connectionId,omitempty"`
	}{

		ConnectionID: m.ConnectionID,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ID int32 `json:"id,omitempty"`

		Op string `json:"op,omitempty"`
	}{

		ID: m.ID(),

		Op: m.Op(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this connection message
func (m *ConnectionMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this connection message based on the context it is used
func (m *ConnectionMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectionMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionMessage) UnmarshalBinary(b []byte) error {
	var res ConnectionMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
