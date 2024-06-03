package events

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"

	"code.cloudfoundry.org/bbs/models"
	"github.com/vito/go-sse/sse"
	"google.golang.org/protobuf/proto"
)

var (
	ErrUnrecognizedEventType = errors.New("unrecognized event type")
	ErrSourceClosed          = errors.New("source closed")
	ErrNoData                = errors.New("event with no data")
)

type invalidPayloadError struct {
	payloadType string
	protoErr    error
}

func NewInvalidPayloadError(payloadType string, protoErr error) error {
	return invalidPayloadError{payloadType: payloadType, protoErr: protoErr}
}

func (e invalidPayloadError) Error() string {
	return fmt.Sprintf("invalid protobuf payload of type %s: %s", e.payloadType, e.protoErr.Error())
}

type rawEventSourceError struct {
	rawError error
}

func NewRawEventSourceError(rawError error) error {
	return rawEventSourceError{rawError: rawError}
}

func (e rawEventSourceError) Error() string {
	return fmt.Sprintf("raw event source error: %s", e.rawError.Error())
}

type closeError struct {
	err error
}

func NewCloseError(err error) error {
	return closeError{err: err}
}

func (e closeError) Error() string {
	return fmt.Sprintf("error closing raw source: %s", e.err.Error())
}

func NewEventFromModelEvent(eventID int, event models.Event) (sse.Event, error) {
	payload, err := proto.Marshal(event.ToEventProto())
	if err != nil {
		return sse.Event{}, err
	}

	encodedPayload := base64.StdEncoding.EncodeToString(payload)
	return sse.Event{
		ID:   strconv.Itoa(eventID),
		Name: string(event.EventType()),
		Data: []byte(encodedPayload),
	}, nil
}

//go:generate counterfeiter -generate

//counterfeiter:generate -o eventfakes/fake_event_source.go . EventSource

// EventSource provides sequential access to a stream of events.
type EventSource interface {
	// Next reads the next event from the source. If the connection is lost, it
	// automatically reconnects.
	//
	// If the end of the stream is reached cleanly (which should actually never
	// happen), io.EOF is returned. If called after or during Close,
	// ErrSourceClosed is returned.
	Next() (models.Event, error)

	// Close releases the underlying response, interrupts any in-flight Next, and
	// prevents further calls to Next.
	Close() error
}

//counterfeiter:generate -o eventfakes/fake_raw_event_source.go . RawEventSource

type RawEventSource interface {
	Next() (sse.Event, error)
	Close() error
}

type eventSource struct {
	rawEventSource RawEventSource
}

func NewEventSource(raw RawEventSource) EventSource {
	return &eventSource{
		rawEventSource: raw,
	}
}

func (e *eventSource) Next() (models.Event, error) {
	rawEvent, err := e.rawEventSource.Next()
	if err != nil {
		switch err {
		case io.EOF:
			return nil, err

		case sse.ErrSourceClosed:
			return nil, ErrSourceClosed

		default:
			return nil, NewRawEventSourceError(err)
		}
	}

	return parseRawEvent(rawEvent)
}

func (e *eventSource) Close() error {
	err := e.rawEventSource.Close()
	if err != nil {
		return NewCloseError(err)
	}

	return nil
}

func parseRawEvent(rawEvent sse.Event) (models.Event, error) {
	data, err := base64.StdEncoding.DecodeString(string(rawEvent.Data))
	if len(data) == 0 {
		return nil, NewInvalidPayloadError(rawEvent.Name, ErrNoData)
	} else if err != nil {
		return nil, NewInvalidPayloadError(rawEvent.Name, err)
	}

	switch rawEvent.Name {
	case models.EventTypeDesiredLRPCreated:
		event := new(models.DesiredLRPCreatedEvent)
		protoEvent := new(models.ProtoDesiredLRPCreatedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

	case models.EventTypeDesiredLRPChanged:
		event := new(models.DesiredLRPChangedEvent)
		protoEvent := new(models.ProtoDesiredLRPChangedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

	case models.EventTypeDesiredLRPRemoved:
		event := new(models.DesiredLRPRemovedEvent)
		protoEvent := new(models.ProtoDesiredLRPRemovedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

		//lint:ignore SA1019 - need to support this event until the deprecation becomes deletion
	case models.EventTypeActualLRPCreated:
		//lint:ignore SA1019 - need to support this event until the deprecation becomes deletion
		event := new(models.ActualLRPCreatedEvent)
		protoEvent := new(models.ProtoActualLRPCreatedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

		//lint:ignore SA1019 - need to support this event until the deprecation becomes deletion
	case models.EventTypeActualLRPChanged:
		//lint:ignore SA1019 - need to support this event until the deprecation becomes deletion
		event := new(models.ActualLRPChangedEvent)
		protoEvent := new(models.ProtoActualLRPChangedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

		//lint:ignore SA1019 - need to support this event until the deprecation becomes deletion
	case models.EventTypeActualLRPRemoved:
		//lint:ignore SA1019 - need to support this event until the deprecation becomes deletion
		event := new(models.ActualLRPRemovedEvent)
		protoEvent := new(models.ProtoActualLRPRemovedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

	case models.EventTypeActualLRPCrashed:
		event := new(models.ActualLRPCrashedEvent)
		protoEvent := new(models.ProtoActualLRPCrashedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

	case models.EventTypeTaskCreated:
		event := new(models.TaskCreatedEvent)
		protoEvent := new(models.ProtoTaskCreatedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

	case models.EventTypeTaskChanged:
		event := new(models.TaskChangedEvent)
		protoEvent := new(models.ProtoTaskChangedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

	case models.EventTypeTaskRemoved:
		event := new(models.TaskRemovedEvent)
		protoEvent := new(models.ProtoTaskRemovedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

	case models.EventTypeActualLRPInstanceCreated:
		event := new(models.ActualLRPInstanceCreatedEvent)
		protoEvent := new(models.ProtoActualLRPInstanceCreatedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

	case models.EventTypeActualLRPInstanceChanged:
		event := new(models.ActualLRPInstanceChangedEvent)
		protoEvent := new(models.ProtoActualLRPInstanceChangedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil

	case models.EventTypeActualLRPInstanceRemoved:
		event := new(models.ActualLRPInstanceRemovedEvent)
		protoEvent := new(models.ProtoActualLRPInstanceRemovedEvent)
		err := proto.Unmarshal(data, protoEvent)
		if err != nil {
			return nil, NewInvalidPayloadError(rawEvent.Name, err)
		}
		event = protoEvent.FromProto()

		return event, nil
	}

	return nil, ErrUnrecognizedEventType
}
