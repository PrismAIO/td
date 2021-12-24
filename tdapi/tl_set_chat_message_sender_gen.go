// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
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

// SetChatMessageSenderRequest represents TL type `setChatMessageSender#ab456b7e`.
type SetChatMessageSenderRequest struct {
	// Chat identifier
	ChatID int64
	// New message sender for the chat
	MessageSenderID MessageSenderClass
}

// SetChatMessageSenderRequestTypeID is TL type id of SetChatMessageSenderRequest.
const SetChatMessageSenderRequestTypeID = 0xab456b7e

// Ensuring interfaces in compile-time for SetChatMessageSenderRequest.
var (
	_ bin.Encoder     = &SetChatMessageSenderRequest{}
	_ bin.Decoder     = &SetChatMessageSenderRequest{}
	_ bin.BareEncoder = &SetChatMessageSenderRequest{}
	_ bin.BareDecoder = &SetChatMessageSenderRequest{}
)

func (s *SetChatMessageSenderRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.MessageSenderID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatMessageSenderRequest) String() string {
	if s == nil {
		return "SetChatMessageSenderRequest(nil)"
	}
	type Alias SetChatMessageSenderRequest
	return fmt.Sprintf("SetChatMessageSenderRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatMessageSenderRequest) TypeID() uint32 {
	return SetChatMessageSenderRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatMessageSenderRequest) TypeName() string {
	return "setChatMessageSender"
}

// TypeInfo returns info about TL type.
func (s *SetChatMessageSenderRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatMessageSender",
		ID:   SetChatMessageSenderRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageSenderID",
			SchemaName: "message_sender_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatMessageSenderRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatMessageSender#ab456b7e as nil")
	}
	b.PutID(SetChatMessageSenderRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatMessageSenderRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatMessageSender#ab456b7e as nil")
	}
	b.PutInt53(s.ChatID)
	if s.MessageSenderID == nil {
		return fmt.Errorf("unable to encode setChatMessageSender#ab456b7e: field message_sender_id is nil")
	}
	if err := s.MessageSenderID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setChatMessageSender#ab456b7e: field message_sender_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatMessageSenderRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatMessageSender#ab456b7e to nil")
	}
	if err := b.ConsumeID(SetChatMessageSenderRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatMessageSender#ab456b7e: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatMessageSenderRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatMessageSender#ab456b7e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setChatMessageSender#ab456b7e: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode setChatMessageSender#ab456b7e: field message_sender_id: %w", err)
		}
		s.MessageSenderID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatMessageSenderRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatMessageSender#ab456b7e as nil")
	}
	b.ObjStart()
	b.PutID("setChatMessageSender")
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.FieldStart("message_sender_id")
	if s.MessageSenderID == nil {
		return fmt.Errorf("unable to encode setChatMessageSender#ab456b7e: field message_sender_id is nil")
	}
	if err := s.MessageSenderID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setChatMessageSender#ab456b7e: field message_sender_id: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatMessageSenderRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatMessageSender#ab456b7e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatMessageSender"); err != nil {
				return fmt.Errorf("unable to decode setChatMessageSender#ab456b7e: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setChatMessageSender#ab456b7e: field chat_id: %w", err)
			}
			s.ChatID = value
		case "message_sender_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode setChatMessageSender#ab456b7e: field message_sender_id: %w", err)
			}
			s.MessageSenderID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatMessageSenderRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetMessageSenderID returns value of MessageSenderID field.
func (s *SetChatMessageSenderRequest) GetMessageSenderID() (value MessageSenderClass) {
	if s == nil {
		return
	}
	return s.MessageSenderID
}

// SetChatMessageSender invokes method setChatMessageSender#ab456b7e returning error if any.
func (c *Client) SetChatMessageSender(ctx context.Context, request *SetChatMessageSenderRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
