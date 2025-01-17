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

// MsgVpnClientProfile msg vpn client profile
//
// swagger:model MsgVpnClientProfile
type MsgVpnClientProfile struct {

	// Enable or disable allowing Bridge clients using the Client Profile to connect. Changing this setting does not affect existing Bridge client connections. The default value is `false`.
	AllowBridgeConnectionsEnabled bool `json:"allowBridgeConnectionsEnabled,omitempty"`

	// Enable or disable allowing clients using the Client Profile to bind to endpoints with the cut-through forwarding delivery mode. Changing this value does not affect existing client connections. The default value is `false`.
	AllowCutThroughForwardingEnabled bool `json:"allowCutThroughForwardingEnabled,omitempty"`

	// The types of Queues and Topic Endpoints that clients using the client-profile can create. Changing this value does not affect existing client connections. The default value is `"all"`. The allowed values and their meaning are:
	//
	// <pre>
	// "all" - Client can create any type of endpoint.
	// "durable" - Client can create only durable endpoints.
	// "non-durable" - Client can create only non-durable endpoints.
	// </pre>
	//  Available since 2.14.
	// Enum: [all durable non-durable]
	AllowGuaranteedEndpointCreateDurability string `json:"allowGuaranteedEndpointCreateDurability,omitempty"`

	// Enable or disable allowing clients using the Client Profile to create topic endponts or queues. Changing this value does not affect existing client connections. The default value is `false`.
	AllowGuaranteedEndpointCreateEnabled bool `json:"allowGuaranteedEndpointCreateEnabled,omitempty"`

	// Enable or disable allowing clients using the Client Profile to receive guaranteed messages. Changing this setting does not affect existing client connections. The default value is `false`.
	AllowGuaranteedMsgReceiveEnabled bool `json:"allowGuaranteedMsgReceiveEnabled,omitempty"`

	// Enable or disable allowing clients using the Client Profile to send guaranteed messages. Changing this setting does not affect existing client connections. The default value is `false`.
	AllowGuaranteedMsgSendEnabled bool `json:"allowGuaranteedMsgSendEnabled,omitempty"`

	// Enable or disable allowing shared subscriptions. Changing this setting does not affect existing subscriptions. The default value is `false`. Available since 2.11.
	AllowSharedSubscriptionsEnabled bool `json:"allowSharedSubscriptionsEnabled,omitempty"`

	// Enable or disable allowing clients using the Client Profile to establish transacted sessions. Changing this setting does not affect existing client connections. The default value is `false`.
	AllowTransactedSessionsEnabled bool `json:"allowTransactedSessionsEnabled,omitempty"`

	// The name of a queue to copy settings from when a new queue is created by a client using the Client Profile. The referenced queue must exist in the Message VPN. The default value is `""`. Deprecated since 2.14. This attribute has been replaced with `apiQueueManagementCopyFromOnCreateTemplateName`.
	APIQueueManagementCopyFromOnCreateName string `json:"apiQueueManagementCopyFromOnCreateName,omitempty"`

	// The name of a queue template to copy settings from when a new queue is created by a client using the Client Profile. If the referenced queue template does not exist, queue creation will fail when it tries to resolve this template. The default value is `""`. Available since 2.14.
	APIQueueManagementCopyFromOnCreateTemplateName string `json:"apiQueueManagementCopyFromOnCreateTemplateName,omitempty"`

	// The name of a topic endpoint to copy settings from when a new topic endpoint is created by a client using the Client Profile. The referenced topic endpoint must exist in the Message VPN. The default value is `""`. Deprecated since 2.14. This attribute has been replaced with `apiTopicEndpointManagementCopyFromOnCreateTemplateName`.
	APITopicEndpointManagementCopyFromOnCreateName string `json:"apiTopicEndpointManagementCopyFromOnCreateName,omitempty"`

	// The name of a topic endpoint template to copy settings from when a new topic endpoint is created by a client using the Client Profile. If the referenced topic endpoint template does not exist, topic endpoint creation will fail when it tries to resolve this template. The default value is `""`. Available since 2.14.
	APITopicEndpointManagementCopyFromOnCreateTemplateName string `json:"apiTopicEndpointManagementCopyFromOnCreateTemplateName,omitempty"`

	// The name of the Client Profile.
	ClientProfileName string `json:"clientProfileName,omitempty"`

	// Enable or disable allowing clients using the Client Profile to use compression. The default value is `true`. Available since 2.10.
	CompressionEnabled bool `json:"compressionEnabled,omitempty"`

	// The amount of time to delay the delivery of messages to clients using the Client Profile after the initial message has been delivered (the eliding delay interval), in milliseconds. A value of 0 means there is no delay in delivering messages to clients. The default value is `0`.
	ElidingDelay int64 `json:"elidingDelay,omitempty"`

	// Enable or disable message eliding for clients using the Client Profile. The default value is `false`.
	ElidingEnabled bool `json:"elidingEnabled,omitempty"`

	// The maximum number of topics tracked for message eliding per client connection using the Client Profile. The default value is `256`.
	ElidingMaxTopicCount int64 `json:"elidingMaxTopicCount,omitempty"`

	// event client provisioned endpoint spool usage threshold
	EventClientProvisionedEndpointSpoolUsageThreshold *EventThresholdByPercent `json:"eventClientProvisionedEndpointSpoolUsageThreshold,omitempty"`

	// event connection count per client username threshold
	EventConnectionCountPerClientUsernameThreshold *EventThreshold `json:"eventConnectionCountPerClientUsernameThreshold,omitempty"`

	// event egress flow count threshold
	EventEgressFlowCountThreshold *EventThreshold `json:"eventEgressFlowCountThreshold,omitempty"`

	// event endpoint count per client username threshold
	EventEndpointCountPerClientUsernameThreshold *EventThreshold `json:"eventEndpointCountPerClientUsernameThreshold,omitempty"`

	// event ingress flow count threshold
	EventIngressFlowCountThreshold *EventThreshold `json:"eventIngressFlowCountThreshold,omitempty"`

	// event service smf connection count per client username threshold
	EventServiceSmfConnectionCountPerClientUsernameThreshold *EventThreshold `json:"eventServiceSmfConnectionCountPerClientUsernameThreshold,omitempty"`

	// event service web connection count per client username threshold
	EventServiceWebConnectionCountPerClientUsernameThreshold *EventThreshold `json:"eventServiceWebConnectionCountPerClientUsernameThreshold,omitempty"`

	// event subscription count threshold
	EventSubscriptionCountThreshold *EventThreshold `json:"eventSubscriptionCountThreshold,omitempty"`

	// event transacted session count threshold
	EventTransactedSessionCountThreshold *EventThreshold `json:"eventTransactedSessionCountThreshold,omitempty"`

	// event transaction count threshold
	EventTransactionCountThreshold *EventThreshold `json:"eventTransactionCountThreshold,omitempty"`

	// The maximum number of client connections per Client Username using the Client Profile. The default is the maximum value supported by the platform.
	MaxConnectionCountPerClientUsername int64 `json:"maxConnectionCountPerClientUsername,omitempty"`

	// The maximum number of transmit flows that can be created by one client using the Client Profile. The default value is `1000`.
	MaxEgressFlowCount int64 `json:"maxEgressFlowCount,omitempty"`

	// The maximum number of queues and topic endpoints that can be created by clients with the same Client Username using the Client Profile. The default value is `1000`.
	MaxEndpointCountPerClientUsername int64 `json:"maxEndpointCountPerClientUsername,omitempty"`

	// The maximum number of receive flows that can be created by one client using the Client Profile. The default value is `1000`.
	MaxIngressFlowCount int64 `json:"maxIngressFlowCount,omitempty"`

	// The maximum number of publisher and consumer messages combined that is allowed within a transaction for each client associated with this client-profile. Exceeding this limit will result in a transaction prepare or commit failure. Changing this value during operation will not affect existing sessions. It is only validated at transaction creation time. Large transactions consume more resources and are more likely to require retrieving messages from the ADB or from disk to process the transaction prepare or commit requests. The transaction processing rate may diminish if a large number of messages must be retrieved from the ADB or from disk. Care should be taken to not use excessively large transactions needlessly to avoid exceeding resource limits and to avoid reducing the overall broker performance. The default value is `256`. Available since 2.20.
	MaxMsgsPerTransaction int32 `json:"maxMsgsPerTransaction,omitempty"`

	// The maximum number of subscriptions per client using the Client Profile. This limit is not enforced when a client adds a subscription to an endpoint, except for MQTT QoS 1 subscriptions. In addition, this limit is not enforced when a subscription is added using a management interface, such as CLI or SEMP. The default varies by platform.
	MaxSubscriptionCount int64 `json:"maxSubscriptionCount,omitempty"`

	// The maximum number of transacted sessions that can be created by one client using the Client Profile. The default value is `10`.
	MaxTransactedSessionCount int64 `json:"maxTransactedSessionCount,omitempty"`

	// The maximum number of transactions that can be created by one client using the Client Profile. The default varies by platform.
	MaxTransactionCount int64 `json:"maxTransactionCount,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The maximum depth of the "Control 1" (C-1) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is `20000`.
	QueueControl1MaxDepth int32 `json:"queueControl1MaxDepth,omitempty"`

	// The number of messages that are always allowed entry into the "Control 1" (C-1) priority queue, regardless of the `queueControl1MaxDepth` value. The default value is `4`.
	QueueControl1MinMsgBurst int32 `json:"queueControl1MinMsgBurst,omitempty"`

	// The maximum depth of the "Direct 1" (D-1) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is `20000`.
	QueueDirect1MaxDepth int32 `json:"queueDirect1MaxDepth,omitempty"`

	// The number of messages that are always allowed entry into the "Direct 1" (D-1) priority queue, regardless of the `queueDirect1MaxDepth` value. The default value is `4`.
	QueueDirect1MinMsgBurst int32 `json:"queueDirect1MinMsgBurst,omitempty"`

	// The maximum depth of the "Direct 2" (D-2) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is `20000`.
	QueueDirect2MaxDepth int32 `json:"queueDirect2MaxDepth,omitempty"`

	// The number of messages that are always allowed entry into the "Direct 2" (D-2) priority queue, regardless of the `queueDirect2MaxDepth` value. The default value is `4`.
	QueueDirect2MinMsgBurst int32 `json:"queueDirect2MinMsgBurst,omitempty"`

	// The maximum depth of the "Direct 3" (D-3) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is `20000`.
	QueueDirect3MaxDepth int32 `json:"queueDirect3MaxDepth,omitempty"`

	// The number of messages that are always allowed entry into the "Direct 3" (D-3) priority queue, regardless of the `queueDirect3MaxDepth` value. The default value is `4`.
	QueueDirect3MinMsgBurst int32 `json:"queueDirect3MinMsgBurst,omitempty"`

	// The maximum depth of the "Guaranteed 1" (G-1) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is `20000`.
	QueueGuaranteed1MaxDepth int32 `json:"queueGuaranteed1MaxDepth,omitempty"`

	// The number of messages that are always allowed entry into the "Guaranteed 1" (G-3) priority queue, regardless of the `queueGuaranteed1MaxDepth` value. The default value is `255`.
	QueueGuaranteed1MinMsgBurst int32 `json:"queueGuaranteed1MinMsgBurst,omitempty"`

	// Enable or disable the sending of a negative acknowledgement (NACK) to a client using the Client Profile when discarding a guaranteed message due to no matching subscription found. The default value is `false`. Available since 2.2.
	RejectMsgToSenderOnNoSubscriptionMatchEnabled bool `json:"rejectMsgToSenderOnNoSubscriptionMatchEnabled,omitempty"`

	// Enable or disable allowing clients using the Client Profile to connect to the Message VPN when its replication state is standby. The default value is `false`.
	ReplicationAllowClientConnectWhenStandbyEnabled bool `json:"replicationAllowClientConnectWhenStandbyEnabled,omitempty"`

	// The minimum client keepalive timeout which will be enforced for client connections. The default value is `30`. Available since 2.19.
	ServiceMinKeepaliveTimeout int32 `json:"serviceMinKeepaliveTimeout,omitempty"`

	// The maximum number of SMF client connections per Client Username using the Client Profile. The default is the maximum value supported by the platform.
	ServiceSmfMaxConnectionCountPerClientUsername int64 `json:"serviceSmfMaxConnectionCountPerClientUsername,omitempty"`

	// Enable or disable the enforcement of a minimum keepalive timeout for SMF clients. The default value is `false`. Available since 2.19.
	ServiceSmfMinKeepaliveEnabled bool `json:"serviceSmfMinKeepaliveEnabled,omitempty"`

	// The timeout for inactive Web Transport client sessions using the Client Profile, in seconds. The default value is `30`.
	ServiceWebInactiveTimeout int64 `json:"serviceWebInactiveTimeout,omitempty"`

	// The maximum number of Web Transport client connections per Client Username using the Client Profile. The default is the maximum value supported by the platform.
	ServiceWebMaxConnectionCountPerClientUsername int64 `json:"serviceWebMaxConnectionCountPerClientUsername,omitempty"`

	// The maximum Web Transport payload size before fragmentation occurs for clients using the Client Profile, in bytes. The size of the header is not included. The default value is `1000000`.
	ServiceWebMaxPayload int64 `json:"serviceWebMaxPayload,omitempty"`

	// The TCP initial congestion window size for clients using the Client Profile, in multiples of the TCP Maximum Segment Size (MSS). Changing the value from its default of 2 results in non-compliance with RFC 2581. Contact Solace Support before changing this value. The default value is `2`.
	TCPCongestionWindowSize int64 `json:"tcpCongestionWindowSize,omitempty"`

	// The number of TCP keepalive retransmissions to a client using the Client Profile before declaring that it is not available. The default value is `5`.
	TCPKeepaliveCount int64 `json:"tcpKeepaliveCount,omitempty"`

	// The amount of time a client connection using the Client Profile must remain idle before TCP begins sending keepalive probes, in seconds. The default value is `3`.
	TCPKeepaliveIdleTime int64 `json:"tcpKeepaliveIdleTime,omitempty"`

	// The amount of time between TCP keepalive retransmissions to a client using the Client Profile when no acknowledgement is received, in seconds. The default value is `1`.
	TCPKeepaliveInterval int64 `json:"tcpKeepaliveInterval,omitempty"`

	// The TCP maximum segment size for clients using the Client Profile, in bytes. Changes are applied to all existing connections. The default value is `1460`.
	TCPMaxSegmentSize int64 `json:"tcpMaxSegmentSize,omitempty"`

	// The TCP maximum window size for clients using the Client Profile, in kilobytes. Changes are applied to all existing connections. The default value is `256`.
	TCPMaxWindowSize int64 `json:"tcpMaxWindowSize,omitempty"`

	// Enable or disable allowing a client using the Client Profile to downgrade an encrypted connection to plain text. The default value is `true`. Available since 2.8.
	TLSAllowDowngradeToPlainTextEnabled bool `json:"tlsAllowDowngradeToPlainTextEnabled,omitempty"`
}

// Validate validates this msg vpn client profile
func (m *MsgVpnClientProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowGuaranteedEndpointCreateDurability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventClientProvisionedEndpointSpoolUsageThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventConnectionCountPerClientUsernameThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventEgressFlowCountThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventEndpointCountPerClientUsernameThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventIngressFlowCountThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventServiceSmfConnectionCountPerClientUsernameThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventServiceWebConnectionCountPerClientUsernameThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventSubscriptionCountThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTransactedSessionCountThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTransactionCountThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnClientProfileTypeAllowGuaranteedEndpointCreateDurabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","durable","non-durable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnClientProfileTypeAllowGuaranteedEndpointCreateDurabilityPropEnum = append(msgVpnClientProfileTypeAllowGuaranteedEndpointCreateDurabilityPropEnum, v)
	}
}

const (

	// MsgVpnClientProfileAllowGuaranteedEndpointCreateDurabilityAll captures enum value "all"
	MsgVpnClientProfileAllowGuaranteedEndpointCreateDurabilityAll string = "all"

	// MsgVpnClientProfileAllowGuaranteedEndpointCreateDurabilityDurable captures enum value "durable"
	MsgVpnClientProfileAllowGuaranteedEndpointCreateDurabilityDurable string = "durable"

	// MsgVpnClientProfileAllowGuaranteedEndpointCreateDurabilityNonDashDurable captures enum value "non-durable"
	MsgVpnClientProfileAllowGuaranteedEndpointCreateDurabilityNonDashDurable string = "non-durable"
)

// prop value enum
func (m *MsgVpnClientProfile) validateAllowGuaranteedEndpointCreateDurabilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msgVpnClientProfileTypeAllowGuaranteedEndpointCreateDurabilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnClientProfile) validateAllowGuaranteedEndpointCreateDurability(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowGuaranteedEndpointCreateDurability) { // not required
		return nil
	}

	// value enum
	if err := m.validateAllowGuaranteedEndpointCreateDurabilityEnum("allowGuaranteedEndpointCreateDurability", "body", m.AllowGuaranteedEndpointCreateDurability); err != nil {
		return err
	}

	return nil
}

func (m *MsgVpnClientProfile) validateEventClientProvisionedEndpointSpoolUsageThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventClientProvisionedEndpointSpoolUsageThreshold) { // not required
		return nil
	}

	if m.EventClientProvisionedEndpointSpoolUsageThreshold != nil {
		if err := m.EventClientProvisionedEndpointSpoolUsageThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventClientProvisionedEndpointSpoolUsageThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) validateEventConnectionCountPerClientUsernameThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventConnectionCountPerClientUsernameThreshold) { // not required
		return nil
	}

	if m.EventConnectionCountPerClientUsernameThreshold != nil {
		if err := m.EventConnectionCountPerClientUsernameThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventConnectionCountPerClientUsernameThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) validateEventEgressFlowCountThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventEgressFlowCountThreshold) { // not required
		return nil
	}

	if m.EventEgressFlowCountThreshold != nil {
		if err := m.EventEgressFlowCountThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventEgressFlowCountThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) validateEventEndpointCountPerClientUsernameThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventEndpointCountPerClientUsernameThreshold) { // not required
		return nil
	}

	if m.EventEndpointCountPerClientUsernameThreshold != nil {
		if err := m.EventEndpointCountPerClientUsernameThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventEndpointCountPerClientUsernameThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) validateEventIngressFlowCountThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventIngressFlowCountThreshold) { // not required
		return nil
	}

	if m.EventIngressFlowCountThreshold != nil {
		if err := m.EventIngressFlowCountThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventIngressFlowCountThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) validateEventServiceSmfConnectionCountPerClientUsernameThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventServiceSmfConnectionCountPerClientUsernameThreshold) { // not required
		return nil
	}

	if m.EventServiceSmfConnectionCountPerClientUsernameThreshold != nil {
		if err := m.EventServiceSmfConnectionCountPerClientUsernameThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventServiceSmfConnectionCountPerClientUsernameThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) validateEventServiceWebConnectionCountPerClientUsernameThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventServiceWebConnectionCountPerClientUsernameThreshold) { // not required
		return nil
	}

	if m.EventServiceWebConnectionCountPerClientUsernameThreshold != nil {
		if err := m.EventServiceWebConnectionCountPerClientUsernameThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventServiceWebConnectionCountPerClientUsernameThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) validateEventSubscriptionCountThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventSubscriptionCountThreshold) { // not required
		return nil
	}

	if m.EventSubscriptionCountThreshold != nil {
		if err := m.EventSubscriptionCountThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventSubscriptionCountThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) validateEventTransactedSessionCountThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTransactedSessionCountThreshold) { // not required
		return nil
	}

	if m.EventTransactedSessionCountThreshold != nil {
		if err := m.EventTransactedSessionCountThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventTransactedSessionCountThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) validateEventTransactionCountThreshold(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTransactionCountThreshold) { // not required
		return nil
	}

	if m.EventTransactionCountThreshold != nil {
		if err := m.EventTransactionCountThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventTransactionCountThreshold")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this msg vpn client profile based on the context it is used
func (m *MsgVpnClientProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventClientProvisionedEndpointSpoolUsageThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventConnectionCountPerClientUsernameThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventEgressFlowCountThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventEndpointCountPerClientUsernameThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventIngressFlowCountThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventServiceSmfConnectionCountPerClientUsernameThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventServiceWebConnectionCountPerClientUsernameThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventSubscriptionCountThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventTransactedSessionCountThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventTransactionCountThreshold(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsgVpnClientProfile) contextValidateEventClientProvisionedEndpointSpoolUsageThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventClientProvisionedEndpointSpoolUsageThreshold != nil {
		if err := m.EventClientProvisionedEndpointSpoolUsageThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventClientProvisionedEndpointSpoolUsageThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) contextValidateEventConnectionCountPerClientUsernameThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventConnectionCountPerClientUsernameThreshold != nil {
		if err := m.EventConnectionCountPerClientUsernameThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventConnectionCountPerClientUsernameThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) contextValidateEventEgressFlowCountThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventEgressFlowCountThreshold != nil {
		if err := m.EventEgressFlowCountThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventEgressFlowCountThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) contextValidateEventEndpointCountPerClientUsernameThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventEndpointCountPerClientUsernameThreshold != nil {
		if err := m.EventEndpointCountPerClientUsernameThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventEndpointCountPerClientUsernameThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) contextValidateEventIngressFlowCountThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventIngressFlowCountThreshold != nil {
		if err := m.EventIngressFlowCountThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventIngressFlowCountThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) contextValidateEventServiceSmfConnectionCountPerClientUsernameThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventServiceSmfConnectionCountPerClientUsernameThreshold != nil {
		if err := m.EventServiceSmfConnectionCountPerClientUsernameThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventServiceSmfConnectionCountPerClientUsernameThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) contextValidateEventServiceWebConnectionCountPerClientUsernameThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventServiceWebConnectionCountPerClientUsernameThreshold != nil {
		if err := m.EventServiceWebConnectionCountPerClientUsernameThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventServiceWebConnectionCountPerClientUsernameThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) contextValidateEventSubscriptionCountThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventSubscriptionCountThreshold != nil {
		if err := m.EventSubscriptionCountThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventSubscriptionCountThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) contextValidateEventTransactedSessionCountThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventTransactedSessionCountThreshold != nil {
		if err := m.EventTransactedSessionCountThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventTransactedSessionCountThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnClientProfile) contextValidateEventTransactionCountThreshold(ctx context.Context, formats strfmt.Registry) error {

	if m.EventTransactionCountThreshold != nil {
		if err := m.EventTransactionCountThreshold.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventTransactionCountThreshold")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnClientProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnClientProfile) UnmarshalBinary(b []byte) error {
	var res MsgVpnClientProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
