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

// MsgVpnJndiConnectionFactory msg vpn jndi connection factory
// swagger:model MsgVpnJndiConnectionFactory
type MsgVpnJndiConnectionFactory struct {

	// Enable or disable whether new JMS connections can use the same Client identifier (ID) as an existing connection. The default value is `false`. Available since 2.3.
	AllowDuplicateClientIDEnabled bool `json:"allowDuplicateClientIdEnabled,omitempty"`

	// The description of the Client. The default value is `""`.
	ClientDescription string `json:"clientDescription,omitempty"`

	// The Client identifier (ID). If not specified, a unique value for it will be generated. The default value is `""`.
	ClientID string `json:"clientId,omitempty"`

	// The name of the JMS Connection Factory.
	ConnectionFactoryName string `json:"connectionFactoryName,omitempty"`

	// Enable or disable overriding by the Subscriber (Consumer) of the deliver-to-one (DTO) property on messages. When enabled, the Subscriber can receive all DTO tagged messages. The default value is `true`.
	DtoReceiveOverrideEnabled bool `json:"dtoReceiveOverrideEnabled,omitempty"`

	// The priority for receiving deliver-to-one (DTO) messages by the Subscriber (Consumer) if the messages are published on the local Router that the Subscriber is directly connected to. The default value is `1`.
	DtoReceiveSubscriberLocalPriority int32 `json:"dtoReceiveSubscriberLocalPriority,omitempty"`

	// The priority for receiving deliver-to-one (DTO) messages by the Subscriber (Consumer) if the messages are published on a remote Router. The default value is `1`.
	DtoReceiveSubscriberNetworkPriority int32 `json:"dtoReceiveSubscriberNetworkPriority,omitempty"`

	// Enable or disable the deliver-to-one (DTO) property on messages sent by the Publisher (Producer). The default value is `false`.
	DtoSendEnabled bool `json:"dtoSendEnabled,omitempty"`

	// Enable or disable whether a durable endpoint will be dynamically created on the Router when the client calls "Session.createDurableSubscriber()" or "Session.createQueue()". The created endpoint respects the message time-to-live (TTL) according to the "dynamicEndpointRespectTtlEnabled" property. The default value is `false`.
	DynamicEndpointCreateDurableEnabled bool `json:"dynamicEndpointCreateDurableEnabled,omitempty"`

	// Enable or disable whether dynamically created durable and non-durable endpoints respect the message time-to-live (TTL) property. The default value is `true`.
	DynamicEndpointRespectTTLEnabled bool `json:"dynamicEndpointRespectTtlEnabled,omitempty"`

	// The timeout for sending the acknowledgement (ACK) for guaranteed messages received by the Subscriber (Consumer), in milliseconds. The default value is `1000`.
	GuaranteedReceiveAckTimeout int32 `json:"guaranteedReceiveAckTimeout,omitempty"`

	// The size of the window for guaranteed messages received by the Subscriber (Consumer), in messages. The default value is `18`.
	GuaranteedReceiveWindowSize int32 `json:"guaranteedReceiveWindowSize,omitempty"`

	// The threshold for sending the acknowledgement (ACK) for guaranteed messages received by the Subscriber (Consumer) as a percentage of the "guaranteedReceiveWindowSize" value. The default value is `60`.
	GuaranteedReceiveWindowSizeAckThreshold int32 `json:"guaranteedReceiveWindowSizeAckThreshold,omitempty"`

	// The timeout for receiving the acknowledgement (ACK) for guaranteed messages sent by the Publisher (Producer), in milliseconds. The default value is `2000`.
	GuaranteedSendAckTimeout int32 `json:"guaranteedSendAckTimeout,omitempty"`

	// The size of the window for non-persistent guaranteed messages sent by the Publisher (Producer), in messages. For persistent messages the window size is fixed at 1. The default value is `255`.
	GuaranteedSendWindowSize int32 `json:"guaranteedSendWindowSize,omitempty"`

	// The default delivery mode for messages sent by the Publisher (Producer). The default value is `"persistent"`. The allowed values and their meaning are:
	//
	// <pre>
	// "persistent" - Router spools messages (persists in the Message Spool) as part of the send operation.
	// "non-persistent" - Router does not spool messages (does not persist in the Message Spool) as part of the send operation.
	// </pre>
	//
	// Enum: [persistent non-persistent]
	MessagingDefaultDeliveryMode string `json:"messagingDefaultDeliveryMode,omitempty"`

	// Enable or disable whether messages sent by the Publisher (Producer) are Dead Message Queue (DMQ) eligible by default. The default value is `false`.
	MessagingDefaultDmqEligibleEnabled bool `json:"messagingDefaultDmqEligibleEnabled,omitempty"`

	// Enable or disable whether messages sent by the Publisher (Producer) are Eliding eligible by default. The default value is `false`.
	MessagingDefaultElidingEligibleEnabled bool `json:"messagingDefaultElidingEligibleEnabled,omitempty"`

	// Enable or disable inclusion (adding or replacing) of the JMSXUserID property in messages sent by the Publisher (Producer). The default value is `false`.
	MessagingJmsxUserIDEnabled bool `json:"messagingJmsxUserIdEnabled,omitempty"`

	// Enable or disable encoding of JMS text messages in Publisher (Producer) messages as XML payload. When disabled, JMS text messages are encoded as a binary attachment. The default value is `true`.
	MessagingTextInXMLPayloadEnabled bool `json:"messagingTextInXmlPayloadEnabled,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The ZLIB compression level for the connection to the Router. The value "0" means no compression, and the value "-1" means the compression level is specified in the JNDI Properties file. The default value is `-1`.
	TransportCompressionLevel int32 `json:"transportCompressionLevel,omitempty"`

	// The maximum number of retry attempts to establish an initial connection to the host (Router) or list of hosts (Routers). The value "0" means a single attempt (no retries), and the value "-1" means to retry forever. The default value is `0`.
	TransportConnectRetryCount int32 `json:"transportConnectRetryCount,omitempty"`

	// The maximum number of retry attempts to establish an initial connection to each host (Router) on the list of hosts (Routers). The value "0" means a single attempt (no retries), and the value "-1" means to retry forever. The default value is `0`.
	TransportConnectRetryPerHostCount int32 `json:"transportConnectRetryPerHostCount,omitempty"`

	// The timeout for establishing an initial connection to the Router, in milliseconds. The default value is `30000`.
	TransportConnectTimeout int32 `json:"transportConnectTimeout,omitempty"`

	// Enable or disable usage of the Direct Transport mode for sending non-persistent messages. When disabled, the Guaranteed Transport mode is used. The default value is `true`.
	TransportDirectTransportEnabled bool `json:"transportDirectTransportEnabled,omitempty"`

	// The maximum number of consecutive application-level keepalive messages sent without the Router response before the connection to the Router is closed. The default value is `3`.
	TransportKeepaliveCount int32 `json:"transportKeepaliveCount,omitempty"`

	// Enable or disable usage of application-level keepalive messages to maintain a connection with the Router. The default value is `true`.
	TransportKeepaliveEnabled bool `json:"transportKeepaliveEnabled,omitempty"`

	// The interval between application-level keepalive messages, in milliseconds. The default value is `3000`.
	TransportKeepaliveInterval int32 `json:"transportKeepaliveInterval,omitempty"`

	// Enable or disable delivery of asynchronous messages directly from the I/O thread. Contact Solace Support before enabling this property. The default value is `false`.
	TransportMsgCallbackOnIoThreadEnabled bool `json:"transportMsgCallbackOnIoThreadEnabled,omitempty"`

	// Enable or disable optimization for the Direct Transport delivery mode. If enabled, the client application is limited to one Publisher (Producer) and one non-durable Subscriber (Consumer). The default value is `false`.
	TransportOptimizeDirectEnabled bool `json:"transportOptimizeDirectEnabled,omitempty"`

	// The connection port number on the Router for SMF clients. The value "-1" means the port is specified in the JNDI Properties file. The default value is `-1`.
	TransportPort int32 `json:"transportPort,omitempty"`

	// The timeout for reading a reply from the Router, in milliseconds. The default value is `10000`.
	TransportReadTimeout int32 `json:"transportReadTimeout,omitempty"`

	// The size of the receive socket buffer, in bytes. It corresponds to the SO_RCVBUF socket option. The default value is `65536`.
	TransportReceiveBufferSize int32 `json:"transportReceiveBufferSize,omitempty"`

	// The maximum number of attempts to reconnect to the host (Router) or list of hosts (Routers) after the connection has been lost. The value "-1" means to retry forever. The default value is `3`.
	TransportReconnectRetryCount int32 `json:"transportReconnectRetryCount,omitempty"`

	// The amount of time before making another attempt to connect or reconnect to the host (Router) after the connection has been lost, in milliseconds. The default value is `3000`.
	TransportReconnectRetryWait int32 `json:"transportReconnectRetryWait,omitempty"`

	// The size of the send socket buffer, in bytes. It corresponds to the SO_SNDBUF socket option. The default value is `65536`.
	TransportSendBufferSize int32 `json:"transportSendBufferSize,omitempty"`

	// Enable or disable the TCP_NODELAY option. When enabled, Nagle's algorithm for TCP/IP congestion control (RFC 896) is disabled. The default value is `true`.
	TransportTCPNoDelayEnabled bool `json:"transportTcpNoDelayEnabled,omitempty"`

	// Enable or disable this as an XA Connection Factory. When enabled, the Connection Factory can be cast to "XAConnectionFactory", "XAQueueConnectionFactory" or "XATopicConnectionFactory". The default value is `false`.
	XaEnabled bool `json:"xaEnabled,omitempty"`
}

// Validate validates this msg vpn jndi connection factory
func (m *MsgVpnJndiConnectionFactory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessagingDefaultDeliveryMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnJndiConnectionFactoryTypeMessagingDefaultDeliveryModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["persistent","non-persistent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnJndiConnectionFactoryTypeMessagingDefaultDeliveryModePropEnum = append(msgVpnJndiConnectionFactoryTypeMessagingDefaultDeliveryModePropEnum, v)
	}
}

const (

	// MsgVpnJndiConnectionFactoryMessagingDefaultDeliveryModePersistent captures enum value "persistent"
	MsgVpnJndiConnectionFactoryMessagingDefaultDeliveryModePersistent string = "persistent"

	// MsgVpnJndiConnectionFactoryMessagingDefaultDeliveryModeNonPersistent captures enum value "non-persistent"
	MsgVpnJndiConnectionFactoryMessagingDefaultDeliveryModeNonPersistent string = "non-persistent"
)

// prop value enum
func (m *MsgVpnJndiConnectionFactory) validateMessagingDefaultDeliveryModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, msgVpnJndiConnectionFactoryTypeMessagingDefaultDeliveryModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnJndiConnectionFactory) validateMessagingDefaultDeliveryMode(formats strfmt.Registry) error {

	if swag.IsZero(m.MessagingDefaultDeliveryMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessagingDefaultDeliveryModeEnum("messagingDefaultDeliveryMode", "body", m.MessagingDefaultDeliveryMode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnJndiConnectionFactory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnJndiConnectionFactory) UnmarshalBinary(b []byte) error {
	var res MsgVpnJndiConnectionFactory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
