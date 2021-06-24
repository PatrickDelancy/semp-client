// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MsgVpnQueueTemplate msg vpn queue template
//
// swagger:model MsgVpnQueueTemplate
type MsgVpnQueueTemplate struct {

	// The access type for delivering messages to consumer flows. The default value is `"exclusive"`. The allowed values and their meaning are:
	//
	// <pre>
	// "exclusive" - Exclusive delivery of messages to the first bound consumer flow.
	// "non-exclusive" - Non-exclusive delivery of messages to all bound consumer flows in a round-robin fashion.
	// </pre>
	//
	// Enum: [exclusive non-exclusive]
	AccessType string `json:"accessType,omitempty"`

	// Enable or disable the propagation of consumer acknowledgements (ACKs) received on the active replication Message VPN to the standby replication Message VPN. The default value is `true`.
	ConsumerAckPropagationEnabled bool `json:"consumerAckPropagationEnabled,omitempty"`

	// The name of the Dead Message Queue (DMQ). The default value is `"#DEAD_MSG_QUEUE"`.
	DeadMsgQueue string `json:"deadMsgQueue,omitempty"`

	// Controls the durability of queues created from this template. If non-durable, the created queue will be non-durable, regardless of the specified durability. If none, the created queue will have the requested durability. The default value is `"none"`. The allowed values and their meaning are:
	//
	// <pre>
	// "none" - The durability of the endpoint will be as requested on create.
	// "non-durable" - The durability of the created queue will be non-durable, regardless of what was requested.
	// </pre>
	//
	// Enum: [none non-durable]
	DurabilityOverride string `json:"durabilityOverride,omitempty"`

	// event bind count threshold
	EventBindCountThreshold *EventThreshold `json:"eventBindCountThreshold,omitempty"`

	// event msg spool usage threshold
	EventMsgSpoolUsageThreshold *EventThreshold `json:"eventMsgSpoolUsageThreshold,omitempty"`

	// event reject low priority msg limit threshold
	EventRejectLowPriorityMsgLimitThreshold *EventThreshold `json:"eventRejectLowPriorityMsgLimitThreshold,omitempty"`

	// The maximum number of consumer flows that can bind. The default value is `1000`.
	MaxBindCount int64 `json:"maxBindCount,omitempty"`

	// The maximum number of messages delivered but not acknowledged per flow. The default value is `10000`.
	MaxDeliveredUnackedMsgsPerFlow int64 `json:"maxDeliveredUnackedMsgsPerFlow,omitempty"`

	// The maximum message size allowed, in bytes (B). The default value is `10000000`.
	MaxMsgSize int32 `json:"maxMsgSize,omitempty"`

	// The maximum message spool usage allowed, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. The default value is `5000`.
	MaxMsgSpoolUsage int64 `json:"maxMsgSpoolUsage,omitempty"`

	// The maximum number of message redelivery attempts that will occur prior to the message being discarded or moved to the DMQ. A value of 0 means to retry forever. The default value is `0`.
	MaxRedeliveryCount int64 `json:"maxRedeliveryCount,omitempty"`

	// The maximum time in seconds a message can stay in a Queue when `respectTtlEnabled` is `"true"`. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the `maxTtl` configured for the Queue, is exceeded. A value of 0 disables expiry. The default value is `0`.
	MaxTTL int64 `json:"maxTtl,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The permission level for all consumers, excluding the owner. The default value is `"no-access"`. The allowed values and their meaning are:
	//
	// <pre>
	// "no-access" - Disallows all access.
	// "read-only" - Read-only access to the messages.
	// "consume" - Consume (read and remove) messages.
	// "modify-topic" - Consume messages or modify the topic/selector.
	// "delete" - Consume messages, modify the topic/selector or delete the Client created endpoint altogether.
	// </pre>
	//
	// Enum: [no-access read-only consume modify-topic delete]
	Permission string `json:"permission,omitempty"`

	// A wildcardable pattern used to determine which Queues use settings from this Template. Two different wildcards are supported: * and >. Similar to topic filters or subscription patterns, a > matches anything (but only when used at the end), and a * matches zero or more characters but never a slash (/). A > is only a wildcard when used at the end, after a /. A * is only allowed at the end, after a slash (/). The default value is `""`.
	QueueNameFilter string `json:"queueNameFilter,omitempty"`

	// The name of the Queue Template.
	QueueTemplateName string `json:"queueTemplateName,omitempty"`

	// Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by maxRedeliveryCount. When disabled, the message will never be delivered from the queue more than once. The default value is `true`. Available since 2.18.
	RedeliveryEnabled bool `json:"redeliveryEnabled,omitempty"`

	// Enable or disable the checking of low priority messages against the `rejectLowPriorityMsgLimit`. This may only be enabled if `rejectMsgToSenderOnDiscardBehavior` does not have a value of `"never"`. The default value is `false`.
	RejectLowPriorityMsgEnabled bool `json:"rejectLowPriorityMsgEnabled,omitempty"`

	// The number of messages of any priority above which low priority messages are not admitted but higher priority messages are allowed. The default value is `0`.
	RejectLowPriorityMsgLimit int64 `json:"rejectLowPriorityMsgLimit,omitempty"`

	// Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs prevent the message from being delivered to any destination and Transacted Session commits to fail. The default value is `"when-queue-enabled"`. The allowed values and their meaning are:
	//
	// <pre>
	// "always" - Always return a negative acknowledgment (NACK) to the sending client on message discard.
	// "when-queue-enabled" - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Queue is enabled.
	// "never" - Never return a negative acknowledgment (NACK) to the sending client on message discard.
	// </pre>
	//
	// Enum: [always when-queue-enabled never]
	RejectMsgToSenderOnDiscardBehavior string `json:"rejectMsgToSenderOnDiscardBehavior,omitempty"`

	// Enable or disable the respecting of message priority. When enabled, messages are delivered in priority order, from 9 (highest) to 0 (lowest). The default value is `false`.
	RespectMsgPriorityEnabled bool `json:"respectMsgPriorityEnabled,omitempty"`

	// Enable or disable the respecting of the time-to-live (TTL) for messages. When enabled, expired messages are discarded or moved to the DMQ. The default value is `false`.
	RespectTTLEnabled bool `json:"respectTtlEnabled,omitempty"`
}

// Validate validates this msg vpn queue template
func (m *MsgVpnQueueTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDurabilityOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventBindCountThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventMsgSpoolUsageThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventRejectLowPriorityMsgLimitThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejectMsgToSenderOnDiscardBehavior(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnQueueTemplateTypeAccessTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["exclusive","non-exclusive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnQueueTemplateTypeAccessTypePropEnum = append(msgVpnQueueTemplateTypeAccessTypePropEnum, v)
	}
}

const (

	// MsgVpnQueueTemplateAccessTypeExclusive captures enum value "exclusive"
	MsgVpnQueueTemplateAccessTypeExclusive string = "exclusive"

	// MsgVpnQueueTemplateAccessTypeNonDashExclusive captures enum value "non-exclusive"
	MsgVpnQueueTemplateAccessTypeNonDashExclusive string = "non-exclusive"
)

// prop value enum
func (m *MsgVpnQueueTemplate) validateAccessTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msgVpnQueueTemplateTypeAccessTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnQueueTemplate) validateAccessType(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessTypeEnum("accessType", "body", m.AccessType); err != nil {
		return err
	}

	return nil
}

var msgVpnQueueTemplateTypeDurabilityOverridePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","non-durable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnQueueTemplateTypeDurabilityOverridePropEnum = append(msgVpnQueueTemplateTypeDurabilityOverridePropEnum, v)
	}
}

const (

	// MsgVpnQueueTemplateDurabilityOverrideNone captures enum value "none"
	MsgVpnQueueTemplateDurabilityOverrideNone string = "none"

	// MsgVpnQueueTemplateDurabilityOverrideNonDashDurable captures enum value "non-durable"
	MsgVpnQueueTemplateDurabilityOverrideNonDashDurable string = "non-durable"
)

// prop value enum
func (m *MsgVpnQueueTemplate) validateDurabilityOverrideEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msgVpnQueueTemplateTypeDurabilityOverridePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnQueueTemplate) validateDurabilityOverride(formats strfmt.Registry) error {
	if swag.IsZero(m.DurabilityOverride) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurabilityOverrideEnum("durabilityOverride", "body", m.DurabilityOverride); err != nil {
		return err
	}

	return nil
}

func (m *MsgVpnQueueTemplate) validateEventBindCountThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventBindCountThreshold) { // not required
		return nil
	}

	if m.EventBindCountThreshold != nil {
		if err := m.EventBindCountThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventBindCountThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnQueueTemplate) validateEventMsgSpoolUsageThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventMsgSpoolUsageThreshold) { // not required
		return nil
	}

	if m.EventMsgSpoolUsageThreshold != nil {
		if err := m.EventMsgSpoolUsageThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventMsgSpoolUsageThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnQueueTemplate) validateEventRejectLowPriorityMsgLimitThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventRejectLowPriorityMsgLimitThreshold) { // not required
		return nil
	}

	if m.EventRejectLowPriorityMsgLimitThreshold != nil {
		if err := m.EventRejectLowPriorityMsgLimitThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventRejectLowPriorityMsgLimitThreshold")
			}
			return err
		}
	}

	return nil
}

var msgVpnQueueTemplateTypePermissionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["no-access","read-only","consume","modify-topic","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnQueueTemplateTypePermissionPropEnum = append(msgVpnQueueTemplateTypePermissionPropEnum, v)
	}
}

const (

	// MsgVpnQueueTemplatePermissionNoDashAccess captures enum value "no-access"
	MsgVpnQueueTemplatePermissionNoDashAccess string = "no-access"

	// MsgVpnQueueTemplatePermissionReadDashOnly captures enum value "read-only"
	MsgVpnQueueTemplatePermissionReadDashOnly string = "read-only"

	// MsgVpnQueueTemplatePermissionConsume captures enum value "consume"
	MsgVpnQueueTemplatePermissionConsume string = "consume"

	// MsgVpnQueueTemplatePermissionModifyDashTopic captures enum value "modify-topic"
	MsgVpnQueueTemplatePermissionModifyDashTopic string = "modify-topic"

	// MsgVpnQueueTemplatePermissionDelete captures enum value "delete"
	MsgVpnQueueTemplatePermissionDelete string = "delete"
)

// prop value enum
func (m *MsgVpnQueueTemplate) validatePermissionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msgVpnQueueTemplateTypePermissionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnQueueTemplate) validatePermission(formats strfmt.Registry) error {
	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	// value enum
	if err := m.validatePermissionEnum("permission", "body", m.Permission); err != nil {
		return err
	}

	return nil
}

var msgVpnQueueTemplateTypeRejectMsgToSenderOnDiscardBehaviorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["always","when-queue-enabled","never"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnQueueTemplateTypeRejectMsgToSenderOnDiscardBehaviorPropEnum = append(msgVpnQueueTemplateTypeRejectMsgToSenderOnDiscardBehaviorPropEnum, v)
	}
}

const (

	// MsgVpnQueueTemplateRejectMsgToSenderOnDiscardBehaviorAlways captures enum value "always"
	MsgVpnQueueTemplateRejectMsgToSenderOnDiscardBehaviorAlways string = "always"

	// MsgVpnQueueTemplateRejectMsgToSenderOnDiscardBehaviorWhenDashQueueDashEnabled captures enum value "when-queue-enabled"
	MsgVpnQueueTemplateRejectMsgToSenderOnDiscardBehaviorWhenDashQueueDashEnabled string = "when-queue-enabled"

	// MsgVpnQueueTemplateRejectMsgToSenderOnDiscardBehaviorNever captures enum value "never"
	MsgVpnQueueTemplateRejectMsgToSenderOnDiscardBehaviorNever string = "never"
)

// prop value enum
func (m *MsgVpnQueueTemplate) validateRejectMsgToSenderOnDiscardBehaviorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msgVpnQueueTemplateTypeRejectMsgToSenderOnDiscardBehaviorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnQueueTemplate) validateRejectMsgToSenderOnDiscardBehavior(formats strfmt.Registry) error {
	if swag.IsZero(m.RejectMsgToSenderOnDiscardBehavior) { // not required
		return nil
	}

	// value enum
	if err := m.validateRejectMsgToSenderOnDiscardBehaviorEnum("rejectMsgToSenderOnDiscardBehavior", "body", m.RejectMsgToSenderOnDiscardBehavior); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this msg vpn queue template based on the context it is used
func (m *MsgVpnQueueTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventBindCountThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventMsgSpoolUsageThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventRejectLowPriorityMsgLimitThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsgVpnQueueTemplate) contextValidateEventBindCountThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventBindCountThreshold != nil {
		if err := m.EventBindCountThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventBindCountThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnQueueTemplate) contextValidateEventMsgSpoolUsageThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventMsgSpoolUsageThreshold != nil {
		if err := m.EventMsgSpoolUsageThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventMsgSpoolUsageThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnQueueTemplate) contextValidateEventRejectLowPriorityMsgLimitThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventRejectLowPriorityMsgLimitThreshold != nil {
		if err := m.EventRejectLowPriorityMsgLimitThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventRejectLowPriorityMsgLimitThreshold")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnQueueTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnQueueTemplate) UnmarshalBinary(b []byte) error {
	var res MsgVpnQueueTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
