// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/PrismAIO/td/bin"
	"github.com/PrismAIO/td/tdjson"
	"github.com/PrismAIO/td/tdp"
	"github.com/PrismAIO/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// Message represents TL type `message#ec200d96`.
//
// See https://localhost:80/doc/constructor/message for reference.
type Message struct {
	// Err field of Message.
	Err Error
}

// MessageTypeID is TL type id of Message.
const MessageTypeID = 0xec200d96

// Ensuring interfaces in compile-time for Message.
var (
	_ bin.Encoder     = &Message{}
	_ bin.Decoder     = &Message{}
	_ bin.BareEncoder = &Message{}
	_ bin.BareDecoder = &Message{}
)

func (m *Message) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Err.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *Message) String() string {
	if m == nil {
		return "Message(nil)"
	}
	type Alias Message
	return fmt.Sprintf("Message%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Message) TypeID() uint32 {
	return MessageTypeID
}

// TypeName returns name of type in TL schema.
func (*Message) TypeName() string {
	return "message"
}

// TypeInfo returns info about TL type.
func (m *Message) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "message",
		ID:   MessageTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Err",
			SchemaName: "err",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *Message) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode message#ec200d96 as nil")
	}
	b.PutID(MessageTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *Message) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode message#ec200d96 as nil")
	}
	if err := m.Err.Encode(b); err != nil {
		return fmt.Errorf("unable to encode message#ec200d96: field err: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *Message) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode message#ec200d96 to nil")
	}
	if err := b.ConsumeID(MessageTypeID); err != nil {
		return fmt.Errorf("unable to decode message#ec200d96: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *Message) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode message#ec200d96 to nil")
	}
	{
		if err := m.Err.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#ec200d96: field err: %w", err)
		}
	}
	return nil
}

// GetErr returns value of Err field.
func (m *Message) GetErr() (value Error) {
	if m == nil {
		return
	}
	return m.Err
}
