// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// GetCallbackQueryMessageRequest represents TL type `getCallbackQueryMessage#bd209172`.
type GetCallbackQueryMessageRequest struct {
	// Identifier of the chat the message belongs to
	ChatID int64
	// Message identifier
	MessageID int64
	// Identifier of the callback query
	CallbackQueryID int64
}

// GetCallbackQueryMessageRequestTypeID is TL type id of GetCallbackQueryMessageRequest.
const GetCallbackQueryMessageRequestTypeID = 0xbd209172

// Ensuring interfaces in compile-time for GetCallbackQueryMessageRequest.
var (
	_ bin.Encoder     = &GetCallbackQueryMessageRequest{}
	_ bin.Decoder     = &GetCallbackQueryMessageRequest{}
	_ bin.BareEncoder = &GetCallbackQueryMessageRequest{}
	_ bin.BareDecoder = &GetCallbackQueryMessageRequest{}
)

func (g *GetCallbackQueryMessageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.CallbackQueryID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetCallbackQueryMessageRequest) String() string {
	if g == nil {
		return "GetCallbackQueryMessageRequest(nil)"
	}
	type Alias GetCallbackQueryMessageRequest
	return fmt.Sprintf("GetCallbackQueryMessageRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetCallbackQueryMessageRequest) TypeID() uint32 {
	return GetCallbackQueryMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetCallbackQueryMessageRequest) TypeName() string {
	return "getCallbackQueryMessage"
}

// TypeInfo returns info about TL type.
func (g *GetCallbackQueryMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getCallbackQueryMessage",
		ID:   GetCallbackQueryMessageRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "CallbackQueryID",
			SchemaName: "callback_query_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetCallbackQueryMessageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCallbackQueryMessage#bd209172 as nil")
	}
	b.PutID(GetCallbackQueryMessageRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetCallbackQueryMessageRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCallbackQueryMessage#bd209172 as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageID)
	b.PutLong(g.CallbackQueryID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetCallbackQueryMessageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCallbackQueryMessage#bd209172 to nil")
	}
	if err := b.ConsumeID(GetCallbackQueryMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getCallbackQueryMessage#bd209172: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetCallbackQueryMessageRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCallbackQueryMessage#bd209172 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getCallbackQueryMessage#bd209172: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getCallbackQueryMessage#bd209172: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getCallbackQueryMessage#bd209172: field callback_query_id: %w", err)
		}
		g.CallbackQueryID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetCallbackQueryMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getCallbackQueryMessage#bd209172 as nil")
	}
	b.ObjStart()
	b.PutID("getCallbackQueryMessage")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.Comma()
	b.FieldStart("callback_query_id")
	b.PutLong(g.CallbackQueryID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetCallbackQueryMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getCallbackQueryMessage#bd209172 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getCallbackQueryMessage"); err != nil {
				return fmt.Errorf("unable to decode getCallbackQueryMessage#bd209172: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getCallbackQueryMessage#bd209172: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getCallbackQueryMessage#bd209172: field message_id: %w", err)
			}
			g.MessageID = value
		case "callback_query_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getCallbackQueryMessage#bd209172: field callback_query_id: %w", err)
			}
			g.CallbackQueryID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetCallbackQueryMessageRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetCallbackQueryMessageRequest) GetMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.MessageID
}

// GetCallbackQueryID returns value of CallbackQueryID field.
func (g *GetCallbackQueryMessageRequest) GetCallbackQueryID() (value int64) {
	if g == nil {
		return
	}
	return g.CallbackQueryID
}

// GetCallbackQueryMessage invokes method getCallbackQueryMessage#bd209172 returning error if any.
func (c *Client) GetCallbackQueryMessage(ctx context.Context, request *GetCallbackQueryMessageRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
