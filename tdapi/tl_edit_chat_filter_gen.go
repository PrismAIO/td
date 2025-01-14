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

// EditChatFilterRequest represents TL type `editChatFilter#9c2caf82`.
type EditChatFilterRequest struct {
	// Chat filter identifier
	ChatFilterID int32
	// The edited chat filter
	Filter ChatFilter
}

// EditChatFilterRequestTypeID is TL type id of EditChatFilterRequest.
const EditChatFilterRequestTypeID = 0x9c2caf82

// Ensuring interfaces in compile-time for EditChatFilterRequest.
var (
	_ bin.Encoder     = &EditChatFilterRequest{}
	_ bin.Decoder     = &EditChatFilterRequest{}
	_ bin.BareEncoder = &EditChatFilterRequest{}
	_ bin.BareDecoder = &EditChatFilterRequest{}
)

func (e *EditChatFilterRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatFilterID == 0) {
		return false
	}
	if !(e.Filter.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditChatFilterRequest) String() string {
	if e == nil {
		return "EditChatFilterRequest(nil)"
	}
	type Alias EditChatFilterRequest
	return fmt.Sprintf("EditChatFilterRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditChatFilterRequest) TypeID() uint32 {
	return EditChatFilterRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditChatFilterRequest) TypeName() string {
	return "editChatFilter"
}

// TypeInfo returns info about TL type.
func (e *EditChatFilterRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editChatFilter",
		ID:   EditChatFilterRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatFilterID",
			SchemaName: "chat_filter_id",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditChatFilterRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editChatFilter#9c2caf82 as nil")
	}
	b.PutID(EditChatFilterRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditChatFilterRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editChatFilter#9c2caf82 as nil")
	}
	b.PutInt32(e.ChatFilterID)
	if err := e.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editChatFilter#9c2caf82: field filter: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditChatFilterRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editChatFilter#9c2caf82 to nil")
	}
	if err := b.ConsumeID(EditChatFilterRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editChatFilter#9c2caf82: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditChatFilterRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editChatFilter#9c2caf82 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode editChatFilter#9c2caf82: field chat_filter_id: %w", err)
		}
		e.ChatFilterID = value
	}
	{
		if err := e.Filter.Decode(b); err != nil {
			return fmt.Errorf("unable to decode editChatFilter#9c2caf82: field filter: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditChatFilterRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editChatFilter#9c2caf82 as nil")
	}
	b.ObjStart()
	b.PutID("editChatFilter")
	b.Comma()
	b.FieldStart("chat_filter_id")
	b.PutInt32(e.ChatFilterID)
	b.Comma()
	b.FieldStart("filter")
	if err := e.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editChatFilter#9c2caf82: field filter: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditChatFilterRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editChatFilter#9c2caf82 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editChatFilter"); err != nil {
				return fmt.Errorf("unable to decode editChatFilter#9c2caf82: %w", err)
			}
		case "chat_filter_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode editChatFilter#9c2caf82: field chat_filter_id: %w", err)
			}
			e.ChatFilterID = value
		case "filter":
			if err := e.Filter.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode editChatFilter#9c2caf82: field filter: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatFilterID returns value of ChatFilterID field.
func (e *EditChatFilterRequest) GetChatFilterID() (value int32) {
	if e == nil {
		return
	}
	return e.ChatFilterID
}

// GetFilter returns value of Filter field.
func (e *EditChatFilterRequest) GetFilter() (value ChatFilter) {
	if e == nil {
		return
	}
	return e.Filter
}

// EditChatFilter invokes method editChatFilter#9c2caf82 returning error if any.
func (c *Client) EditChatFilter(ctx context.Context, request *EditChatFilterRequest) (*ChatFilterInfo, error) {
	var result ChatFilterInfo

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
