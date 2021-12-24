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

// GetChatSponsoredMessageRequest represents TL type `getChatSponsoredMessage#72c10b`.
type GetChatSponsoredMessageRequest struct {
	// Identifier of the chat
	ChatID int64
}

// GetChatSponsoredMessageRequestTypeID is TL type id of GetChatSponsoredMessageRequest.
const GetChatSponsoredMessageRequestTypeID = 0x72c10b

// Ensuring interfaces in compile-time for GetChatSponsoredMessageRequest.
var (
	_ bin.Encoder     = &GetChatSponsoredMessageRequest{}
	_ bin.Decoder     = &GetChatSponsoredMessageRequest{}
	_ bin.BareEncoder = &GetChatSponsoredMessageRequest{}
	_ bin.BareDecoder = &GetChatSponsoredMessageRequest{}
)

func (g *GetChatSponsoredMessageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatSponsoredMessageRequest) String() string {
	if g == nil {
		return "GetChatSponsoredMessageRequest(nil)"
	}
	type Alias GetChatSponsoredMessageRequest
	return fmt.Sprintf("GetChatSponsoredMessageRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatSponsoredMessageRequest) TypeID() uint32 {
	return GetChatSponsoredMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatSponsoredMessageRequest) TypeName() string {
	return "getChatSponsoredMessage"
}

// TypeInfo returns info about TL type.
func (g *GetChatSponsoredMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatSponsoredMessage",
		ID:   GetChatSponsoredMessageRequestTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatSponsoredMessageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatSponsoredMessage#72c10b as nil")
	}
	b.PutID(GetChatSponsoredMessageRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatSponsoredMessageRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatSponsoredMessage#72c10b as nil")
	}
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatSponsoredMessageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatSponsoredMessage#72c10b to nil")
	}
	if err := b.ConsumeID(GetChatSponsoredMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatSponsoredMessage#72c10b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatSponsoredMessageRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatSponsoredMessage#72c10b to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatSponsoredMessage#72c10b: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatSponsoredMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatSponsoredMessage#72c10b as nil")
	}
	b.ObjStart()
	b.PutID("getChatSponsoredMessage")
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatSponsoredMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatSponsoredMessage#72c10b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatSponsoredMessage"); err != nil {
				return fmt.Errorf("unable to decode getChatSponsoredMessage#72c10b: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatSponsoredMessage#72c10b: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatSponsoredMessageRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetChatSponsoredMessage invokes method getChatSponsoredMessage#72c10b returning error if any.
func (c *Client) GetChatSponsoredMessage(ctx context.Context, chatid int64) (*SponsoredMessage, error) {
	var result SponsoredMessage

	request := &GetChatSponsoredMessageRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
