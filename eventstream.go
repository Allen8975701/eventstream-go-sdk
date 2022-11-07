/*
 * Copyright 2019 AccelByte Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package eventstream

import (
	"context"
	"errors"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	eventStreamNull   = "none"
	eventStreamStdout = "stdout"
	eventStreamKafka  = "kafka"
)

const (
	separator      = "."
	defaultVersion = 1
)

// log level
const (
	OffLevel   = "off"
	InfoLevel  = "info"
	DebugLevel = "debug"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

// Event defines the structure of event
type Event struct {
	ID               string                 `json:"id,omitempty"`
	EventName        string                 `json:"name,omitempty"`
	Namespace        string                 `json:"namespace,omitempty"`
	ParentNamespace  string                 `json:"parent_namespace,omitempty"`
	ClientID         string                 `json:"clientId,omitempty"`
	TraceID          string                 `json:"traceId,omitempty"`
	SpanContext      string                 `json:"spanContext,omitempty"`
	UserID           string                 `json:"userId,omitempty"`
	SessionID        string                 `json:"sessionId,omitempty"`
	Timestamp        string                 `json:"timestamp,omitempty"`
	Version          int                    `json:"version,omitempty"`
	EventID          int                    `json:"event_id,omitempty"`
	EventType        int                    `json:"event_type,omitempty"`
	EventLevel       int                    `json:"event_level,omitempty"`
	ServiceName      string                 `json:"service,omitempty"`
	ClientIDs        []string               `json:"client_ids,omitempty"`
	TargetUserIDs    []string               `json:"target_user_ids,omitempty"`
	TargetNamespace  string                 `json:"target_namespace,omitempty"`
	Privacy          bool                   `json:"privacy,omitempty"`
	Topic            string                 `json:"topic,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	Payload          map[string]interface{} `json:"payload,omitempty"`

	Partition int    `json:",omitempty"`
	Offset    int64  `json:",omitempty"`
	Key       string `json:",omitempty"`
}

var (
	NotificationEventNamePath       = "name"
	FreeformNotificationUserIDsPath = []string{"payload", "userIds"}
)

// BrokerConfig is custom configuration for message broker
type BrokerConfig struct {
	StrictValidation bool
	CACertFile       string
	DialTimeout      time.Duration
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	Balancer         kafka.Balancer
	SecurityConfig   *SecurityConfig
}

// SecurityConfig contains security configuration for message broker
type SecurityConfig struct {
	AuthenticationType string
	SASLUsername       string
	SASLPassword       string
}

// PublishBuilder defines the structure of message which is sent through message broker
type PublishBuilder struct {
	topic            []string
	eventName        string
	namespace        string
	parentNamespace  string
	clientID         string
	traceID          string
	spanContext      string
	userID           string
	sessionID        string
	version          int
	eventID          int
	eventType        int
	eventLevel       int
	serviceName      string
	clientIDs        []string
	targetUserIDs    []string
	targetNamespace  string
	privacy          bool
	additionalFields map[string]interface{}
	key              string
	payload          map[string]interface{}
	errorCallback    func(event *Event, err error)
	ctx              context.Context
}

// NewPublish create new PublishBuilder instance
func NewPublish() *PublishBuilder {
	return &PublishBuilder{
		version:       defaultVersion,
		ctx:           context.Background(),
		errorCallback: nil,
	}
}

// Topic set channel / topic name
func (p *PublishBuilder) Topic(topics ...string) *PublishBuilder {
	p.topic = append(p.topic, topics...)
	return p
}

// EventName set name of published event
func (p *PublishBuilder) EventName(eventName string) *PublishBuilder {
	p.eventName = eventName
	return p
}

// Namespace set namespace of published event
func (p *PublishBuilder) Namespace(namespace string) *PublishBuilder {
	p.namespace = namespace
	return p
}

func (p *PublishBuilder) ParentNamespace(parentNamespace string) *PublishBuilder {
	p.parentNamespace = parentNamespace
	return p
}

// ClientID set clientID of publisher event
func (p *PublishBuilder) ClientID(clientID string) *PublishBuilder {
	p.clientID = clientID
	return p
}

// TraceID set traceID of publisher event
func (p *PublishBuilder) TraceID(traceID string) *PublishBuilder {
	p.traceID = traceID
	return p
}

// SpanContext set jaeger spanContext of publisher event
func (p *PublishBuilder) SpanContext(spanID string) *PublishBuilder {
	p.spanContext = spanID
	return p
}

// SessionID set sessionID of publisher event
func (p *PublishBuilder) SessionID(sessionID string) *PublishBuilder {
	p.sessionID = sessionID
	return p
}

// UserID set userID of publisher event
func (p *PublishBuilder) UserID(userID string) *PublishBuilder {
	p.userID = userID
	return p
}

// Version set event schema version
func (p *PublishBuilder) Version(version int) *PublishBuilder {
	p.version = version
	return p
}

// EventID set eventID of publisher event
func (p *PublishBuilder) EventID(eventID int) *PublishBuilder {
	p.eventID = eventID
	return p
}

// EventType set eventType of publisher event
func (p *PublishBuilder) EventType(eventType int) *PublishBuilder {
	p.eventType = eventType
	return p
}

// EventLevel set eventLevel of publisher event
func (p *PublishBuilder) EventLevel(eventLevel int) *PublishBuilder {
	p.eventLevel = eventLevel
	return p
}

// ServiceName set serviceName of publisher event
func (p *PublishBuilder) ServiceName(serviceName string) *PublishBuilder {
	p.serviceName = serviceName
	return p
}

// ClientIDs set clientIDs of publisher event
func (p *PublishBuilder) ClientIDs(clientIDs []string) *PublishBuilder {
	p.clientIDs = clientIDs
	return p
}

// TargetUserIDs set targetUserIDs of publisher event
func (p *PublishBuilder) TargetUserIDs(targetUserIDs []string) *PublishBuilder {
	p.targetUserIDs = targetUserIDs
	return p
}

// TargetNamespace set targetNamespace of publisher event
func (p *PublishBuilder) TargetNamespace(targetNamespace string) *PublishBuilder {
	p.targetNamespace = targetNamespace
	return p
}

// Privacy set privacy of publisher event
func (p *PublishBuilder) Privacy(privacy bool) *PublishBuilder {
	p.privacy = privacy
	return p
}

// AdditionalFields set AdditionalFields of publisher event
func (p *PublishBuilder) AdditionalFields(additionalFields map[string]interface{}) *PublishBuilder {
	p.additionalFields = additionalFields
	return p
}

// Key is a message key that used to determine the partition of the topic
// if client require strong order for the events
func (p *PublishBuilder) Key(key string) *PublishBuilder {
	p.key = key
	return p
}

// Payload is a event payload that will be published
func (p *PublishBuilder) Payload(payload map[string]interface{}) *PublishBuilder {
	p.payload = payload
	return p
}

// ErrorCallback function to handle the event when failed to publish
func (p *PublishBuilder) ErrorCallback(errorCallback func(event *Event, err error)) *PublishBuilder {
	p.errorCallback = errorCallback
	return p
}

// Context define client context when publish event.
// default: context.Background()
func (p *PublishBuilder) Context(ctx context.Context) *PublishBuilder {
	p.ctx = ctx
	return p
}

// SubscribeBuilder defines the structure of message which is sent through message broker
type SubscribeBuilder struct {
	topic       string
	groupID     string
	offset      int64
	callback    func(ctx context.Context, event *Event, err error) error
	eventName   string
	ctx         context.Context
	callbackRaw func(ctx context.Context, msgValue []byte, err error) error
}

// NewSubscribe create new SubscribeBuilder instance
func NewSubscribe() *SubscribeBuilder {
	return &SubscribeBuilder{
		ctx:    context.Background(),
		offset: kafka.LastOffset,
	}
}

// Topic set topic that will be subscribe
func (s *SubscribeBuilder) Topic(topic string) *SubscribeBuilder {
	s.topic = topic
	return s
}

// Offset set Offset of the event to start
func (s *SubscribeBuilder) Offset(offset int64) *SubscribeBuilder {
	s.offset = offset
	return s
}

// GroupID set subscriber groupID or queue group name
func (s *SubscribeBuilder) GroupID(groupID string) *SubscribeBuilder {
	s.groupID = groupID
	return s
}

// EventName set event name that will be subscribe
func (s *SubscribeBuilder) EventName(eventName string) *SubscribeBuilder {
	s.eventName = eventName
	return s
}

// Callback to do when the event received
func (s *SubscribeBuilder) Callback(
	callback func(ctx context.Context, event *Event, err error) error,
) *SubscribeBuilder {
	s.callback = callback
	return s
}

// CallbackRaw callback that receives the undecoded payload
func (s *SubscribeBuilder) CallbackRaw(
	f func(ctx context.Context, msgValue []byte, err error) error,
) *SubscribeBuilder {
	s.callbackRaw = f
	return s
}

// Context define client context when subscribe event.
// default: context.Background()
func (s *SubscribeBuilder) Context(ctx context.Context) *SubscribeBuilder {
	s.ctx = ctx
	return s
}

func NewClient(prefix, stream string, brokers []string, config ...*BrokerConfig) (Client, error) {
	switch stream {
	case eventStreamNull:
		return newBlackholeClient(), nil
	case eventStreamStdout:
		return newStdoutClient(prefix), nil
	case eventStreamKafka:
		return newKafkaClient(brokers, prefix, config...)
	default:
		return nil, errors.New("unsupported stream")
	}
}

// Client is an interface for event stream functionality
type Client interface {
	Publish(publishBuilder *PublishBuilder) error
	PublishSync(publishBuilder *PublishBuilder) error
	Register(subscribeBuilder *SubscribeBuilder) error
}
