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

// RemoveChatActionBarRequest represents TL type `removeChatActionBar#9d9839fa`.
type RemoveChatActionBarRequest struct {
	// Chat identifier
	ChatID int64
}

// RemoveChatActionBarRequestTypeID is TL type id of RemoveChatActionBarRequest.
const RemoveChatActionBarRequestTypeID = 0x9d9839fa

// Ensuring interfaces in compile-time for RemoveChatActionBarRequest.
var (
	_ bin.Encoder     = &RemoveChatActionBarRequest{}
	_ bin.Decoder     = &RemoveChatActionBarRequest{}
	_ bin.BareEncoder = &RemoveChatActionBarRequest{}
	_ bin.BareDecoder = &RemoveChatActionBarRequest{}
)

func (r *RemoveChatActionBarRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RemoveChatActionBarRequest) String() string {
	if r == nil {
		return "RemoveChatActionBarRequest(nil)"
	}
	type Alias RemoveChatActionBarRequest
	return fmt.Sprintf("RemoveChatActionBarRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RemoveChatActionBarRequest) TypeID() uint32 {
	return RemoveChatActionBarRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RemoveChatActionBarRequest) TypeName() string {
	return "removeChatActionBar"
}

// TypeInfo returns info about TL type.
func (r *RemoveChatActionBarRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "removeChatActionBar",
		ID:   RemoveChatActionBarRequestTypeID,
	}
	if r == nil {
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
func (r *RemoveChatActionBarRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeChatActionBar#9d9839fa as nil")
	}
	b.PutID(RemoveChatActionBarRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RemoveChatActionBarRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeChatActionBar#9d9839fa as nil")
	}
	b.PutInt53(r.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RemoveChatActionBarRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeChatActionBar#9d9839fa to nil")
	}
	if err := b.ConsumeID(RemoveChatActionBarRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode removeChatActionBar#9d9839fa: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RemoveChatActionBarRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeChatActionBar#9d9839fa to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode removeChatActionBar#9d9839fa: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RemoveChatActionBarRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode removeChatActionBar#9d9839fa as nil")
	}
	b.ObjStart()
	b.PutID("removeChatActionBar")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RemoveChatActionBarRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode removeChatActionBar#9d9839fa to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("removeChatActionBar"); err != nil {
				return fmt.Errorf("unable to decode removeChatActionBar#9d9839fa: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode removeChatActionBar#9d9839fa: field chat_id: %w", err)
			}
			r.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *RemoveChatActionBarRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// RemoveChatActionBar invokes method removeChatActionBar#9d9839fa returning error if any.
func (c *Client) RemoveChatActionBar(ctx context.Context, chatid int64) error {
	var ok Ok

	request := &RemoveChatActionBarRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
