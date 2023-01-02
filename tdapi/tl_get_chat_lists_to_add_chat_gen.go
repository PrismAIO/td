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

// GetChatListsToAddChatRequest represents TL type `getChatListsToAddChat#2709d6a1`.
type GetChatListsToAddChatRequest struct {
	// Chat identifier
	ChatID int64
}

// GetChatListsToAddChatRequestTypeID is TL type id of GetChatListsToAddChatRequest.
const GetChatListsToAddChatRequestTypeID = 0x2709d6a1

// Ensuring interfaces in compile-time for GetChatListsToAddChatRequest.
var (
	_ bin.Encoder     = &GetChatListsToAddChatRequest{}
	_ bin.Decoder     = &GetChatListsToAddChatRequest{}
	_ bin.BareEncoder = &GetChatListsToAddChatRequest{}
	_ bin.BareDecoder = &GetChatListsToAddChatRequest{}
)

func (g *GetChatListsToAddChatRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatListsToAddChatRequest) String() string {
	if g == nil {
		return "GetChatListsToAddChatRequest(nil)"
	}
	type Alias GetChatListsToAddChatRequest
	return fmt.Sprintf("GetChatListsToAddChatRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatListsToAddChatRequest) TypeID() uint32 {
	return GetChatListsToAddChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatListsToAddChatRequest) TypeName() string {
	return "getChatListsToAddChat"
}

// TypeInfo returns info about TL type.
func (g *GetChatListsToAddChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatListsToAddChat",
		ID:   GetChatListsToAddChatRequestTypeID,
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
func (g *GetChatListsToAddChatRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatListsToAddChat#2709d6a1 as nil")
	}
	b.PutID(GetChatListsToAddChatRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatListsToAddChatRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatListsToAddChat#2709d6a1 as nil")
	}
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatListsToAddChatRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatListsToAddChat#2709d6a1 to nil")
	}
	if err := b.ConsumeID(GetChatListsToAddChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatListsToAddChat#2709d6a1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatListsToAddChatRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatListsToAddChat#2709d6a1 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatListsToAddChat#2709d6a1: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatListsToAddChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatListsToAddChat#2709d6a1 as nil")
	}
	b.ObjStart()
	b.PutID("getChatListsToAddChat")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatListsToAddChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatListsToAddChat#2709d6a1 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatListsToAddChat"); err != nil {
				return fmt.Errorf("unable to decode getChatListsToAddChat#2709d6a1: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatListsToAddChat#2709d6a1: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatListsToAddChatRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetChatListsToAddChat invokes method getChatListsToAddChat#2709d6a1 returning error if any.
func (c *Client) GetChatListsToAddChat(ctx context.Context, chatid int64) (*ChatLists, error) {
	var result ChatLists

	request := &GetChatListsToAddChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
