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

// ToggleChatHasProtectedContentRequest represents TL type `toggleChatHasProtectedContent#3a20d94d`.
type ToggleChatHasProtectedContentRequest struct {
	// Chat identifier
	ChatID int64
	// New value of has_protected_content
	HasProtectedContent bool
}

// ToggleChatHasProtectedContentRequestTypeID is TL type id of ToggleChatHasProtectedContentRequest.
const ToggleChatHasProtectedContentRequestTypeID = 0x3a20d94d

// Ensuring interfaces in compile-time for ToggleChatHasProtectedContentRequest.
var (
	_ bin.Encoder     = &ToggleChatHasProtectedContentRequest{}
	_ bin.Decoder     = &ToggleChatHasProtectedContentRequest{}
	_ bin.BareEncoder = &ToggleChatHasProtectedContentRequest{}
	_ bin.BareDecoder = &ToggleChatHasProtectedContentRequest{}
)

func (t *ToggleChatHasProtectedContentRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ChatID == 0) {
		return false
	}
	if !(t.HasProtectedContent == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleChatHasProtectedContentRequest) String() string {
	if t == nil {
		return "ToggleChatHasProtectedContentRequest(nil)"
	}
	type Alias ToggleChatHasProtectedContentRequest
	return fmt.Sprintf("ToggleChatHasProtectedContentRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleChatHasProtectedContentRequest) TypeID() uint32 {
	return ToggleChatHasProtectedContentRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleChatHasProtectedContentRequest) TypeName() string {
	return "toggleChatHasProtectedContent"
}

// TypeInfo returns info about TL type.
func (t *ToggleChatHasProtectedContentRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleChatHasProtectedContent",
		ID:   ToggleChatHasProtectedContentRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "HasProtectedContent",
			SchemaName: "has_protected_content",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleChatHasProtectedContentRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatHasProtectedContent#3a20d94d as nil")
	}
	b.PutID(ToggleChatHasProtectedContentRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleChatHasProtectedContentRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatHasProtectedContent#3a20d94d as nil")
	}
	b.PutInt53(t.ChatID)
	b.PutBool(t.HasProtectedContent)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleChatHasProtectedContentRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatHasProtectedContent#3a20d94d to nil")
	}
	if err := b.ConsumeID(ToggleChatHasProtectedContentRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleChatHasProtectedContent#3a20d94d: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleChatHasProtectedContentRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatHasProtectedContent#3a20d94d to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatHasProtectedContent#3a20d94d: field chat_id: %w", err)
		}
		t.ChatID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleChatHasProtectedContent#3a20d94d: field has_protected_content: %w", err)
		}
		t.HasProtectedContent = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleChatHasProtectedContentRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleChatHasProtectedContent#3a20d94d as nil")
	}
	b.ObjStart()
	b.PutID("toggleChatHasProtectedContent")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(t.ChatID)
	b.Comma()
	b.FieldStart("has_protected_content")
	b.PutBool(t.HasProtectedContent)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleChatHasProtectedContentRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleChatHasProtectedContent#3a20d94d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleChatHasProtectedContent"); err != nil {
				return fmt.Errorf("unable to decode toggleChatHasProtectedContent#3a20d94d: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatHasProtectedContent#3a20d94d: field chat_id: %w", err)
			}
			t.ChatID = value
		case "has_protected_content":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleChatHasProtectedContent#3a20d94d: field has_protected_content: %w", err)
			}
			t.HasProtectedContent = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (t *ToggleChatHasProtectedContentRequest) GetChatID() (value int64) {
	if t == nil {
		return
	}
	return t.ChatID
}

// GetHasProtectedContent returns value of HasProtectedContent field.
func (t *ToggleChatHasProtectedContentRequest) GetHasProtectedContent() (value bool) {
	if t == nil {
		return
	}
	return t.HasProtectedContent
}

// ToggleChatHasProtectedContent invokes method toggleChatHasProtectedContent#3a20d94d returning error if any.
func (c *Client) ToggleChatHasProtectedContent(ctx context.Context, request *ToggleChatHasProtectedContentRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
