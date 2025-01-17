// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnDistributedCacheClusterTopic msg vpn distributed cache cluster topic
//
// swagger:model MsgVpnDistributedCacheClusterTopic
type MsgVpnDistributedCacheClusterTopic struct {

	// The name of the Distributed Cache.
	CacheName string `json:"cacheName,omitempty"`

	// The name of the Cache Cluster.
	ClusterName string `json:"clusterName,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The value of the Topic in the form a/b/c.
	Topic string `json:"topic,omitempty"`
}

// Validate validates this msg vpn distributed cache cluster topic
func (m *MsgVpnDistributedCacheClusterTopic) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn distributed cache cluster topic based on context it is used
func (m *MsgVpnDistributedCacheClusterTopic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnDistributedCacheClusterTopic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnDistributedCacheClusterTopic) UnmarshalBinary(b []byte) error {
	var res MsgVpnDistributedCacheClusterTopic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
