// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MsgVpnMqttSessionSubscriptionsResponse msg vpn mqtt session subscriptions response
//
// swagger:model MsgVpnMqttSessionSubscriptionsResponse
type MsgVpnMqttSessionSubscriptionsResponse struct {

	// data
	Data []*MsgVpnMqttSessionSubscription `json:"data"`

	// links
	Links []*MsgVpnMqttSessionSubscriptionLinks `json:"links"`

	// meta
	// Required: true
	Meta *SempMeta `json:"meta"`
}

// Validate validates this msg vpn mqtt session subscriptions response
func (m *MsgVpnMqttSessionSubscriptionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsgVpnMqttSessionSubscriptionsResponse) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsgVpnMqttSessionSubscriptionsResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsgVpnMqttSessionSubscriptionsResponse) validateMeta(formats strfmt.Registry) error {

	if err := validate.Required("meta", "body", m.Meta); err != nil {
		return err
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this msg vpn mqtt session subscriptions response based on the context it is used
func (m *MsgVpnMqttSessionSubscriptionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsgVpnMqttSessionSubscriptionsResponse) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {
			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsgVpnMqttSessionSubscriptionsResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsgVpnMqttSessionSubscriptionsResponse) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {
		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnMqttSessionSubscriptionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnMqttSessionSubscriptionsResponse) UnmarshalBinary(b []byte) error {
	var res MsgVpnMqttSessionSubscriptionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
