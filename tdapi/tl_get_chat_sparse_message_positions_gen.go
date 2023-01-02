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

// GetChatSparseMessagePositionsRequest represents TL type `getChatSparseMessagePositions#e472f784`.
type GetChatSparseMessagePositionsRequest struct {
	// Identifier of the chat in which to return information about message positions
	ChatID int64
	// Filter for message content. Filters searchMessagesFilterEmpty,
	// searchMessagesFilterMention, searchMessagesFilterUnreadMention, and
	// searchMessagesFilterUnreadReaction are unsupported in this function
	Filter SearchMessagesFilterClass
	// The message identifier from which to return information about message positions
	FromMessageID int64
	// The expected number of message positions to be returned; 50-2000. A smaller number of
	// positions can be returned, if there are not enough appropriate messages
	Limit int32
}

// GetChatSparseMessagePositionsRequestTypeID is TL type id of GetChatSparseMessagePositionsRequest.
const GetChatSparseMessagePositionsRequestTypeID = 0xe472f784

// Ensuring interfaces in compile-time for GetChatSparseMessagePositionsRequest.
var (
	_ bin.Encoder     = &GetChatSparseMessagePositionsRequest{}
	_ bin.Decoder     = &GetChatSparseMessagePositionsRequest{}
	_ bin.BareEncoder = &GetChatSparseMessagePositionsRequest{}
	_ bin.BareDecoder = &GetChatSparseMessagePositionsRequest{}
)

func (g *GetChatSparseMessagePositionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.FromMessageID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatSparseMessagePositionsRequest) String() string {
	if g == nil {
		return "GetChatSparseMessagePositionsRequest(nil)"
	}
	type Alias GetChatSparseMessagePositionsRequest
	return fmt.Sprintf("GetChatSparseMessagePositionsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatSparseMessagePositionsRequest) TypeID() uint32 {
	return GetChatSparseMessagePositionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatSparseMessagePositionsRequest) TypeName() string {
	return "getChatSparseMessagePositions"
}

// TypeInfo returns info about TL type.
func (g *GetChatSparseMessagePositionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatSparseMessagePositions",
		ID:   GetChatSparseMessagePositionsRequestTypeID,
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
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "FromMessageID",
			SchemaName: "from_message_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatSparseMessagePositionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatSparseMessagePositions#e472f784 as nil")
	}
	b.PutID(GetChatSparseMessagePositionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatSparseMessagePositionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatSparseMessagePositions#e472f784 as nil")
	}
	b.PutInt53(g.ChatID)
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getChatSparseMessagePositions#e472f784: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatSparseMessagePositions#e472f784: field filter: %w", err)
	}
	b.PutInt53(g.FromMessageID)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatSparseMessagePositionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatSparseMessagePositions#e472f784 to nil")
	}
	if err := b.ConsumeID(GetChatSparseMessagePositionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatSparseMessagePositions#e472f784: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatSparseMessagePositionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatSparseMessagePositions#e472f784 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatSparseMessagePositions#e472f784: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := DecodeSearchMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode getChatSparseMessagePositions#e472f784: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatSparseMessagePositions#e472f784: field from_message_id: %w", err)
		}
		g.FromMessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatSparseMessagePositions#e472f784: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatSparseMessagePositionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatSparseMessagePositions#e472f784 as nil")
	}
	b.ObjStart()
	b.PutID("getChatSparseMessagePositions")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("filter")
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getChatSparseMessagePositions#e472f784: field filter is nil")
	}
	if err := g.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatSparseMessagePositions#e472f784: field filter: %w", err)
	}
	b.Comma()
	b.FieldStart("from_message_id")
	b.PutInt53(g.FromMessageID)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatSparseMessagePositionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatSparseMessagePositions#e472f784 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatSparseMessagePositions"); err != nil {
				return fmt.Errorf("unable to decode getChatSparseMessagePositions#e472f784: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatSparseMessagePositions#e472f784: field chat_id: %w", err)
			}
			g.ChatID = value
		case "filter":
			value, err := DecodeTDLibJSONSearchMessagesFilter(b)
			if err != nil {
				return fmt.Errorf("unable to decode getChatSparseMessagePositions#e472f784: field filter: %w", err)
			}
			g.Filter = value
		case "from_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatSparseMessagePositions#e472f784: field from_message_id: %w", err)
			}
			g.FromMessageID = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatSparseMessagePositions#e472f784: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatSparseMessagePositionsRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetFilter returns value of Filter field.
func (g *GetChatSparseMessagePositionsRequest) GetFilter() (value SearchMessagesFilterClass) {
	if g == nil {
		return
	}
	return g.Filter
}

// GetFromMessageID returns value of FromMessageID field.
func (g *GetChatSparseMessagePositionsRequest) GetFromMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.FromMessageID
}

// GetLimit returns value of Limit field.
func (g *GetChatSparseMessagePositionsRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetChatSparseMessagePositions invokes method getChatSparseMessagePositions#e472f784 returning error if any.
func (c *Client) GetChatSparseMessagePositions(ctx context.Context, request *GetChatSparseMessagePositionsRequest) (*MessagePositions, error) {
	var result MessagePositions

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
