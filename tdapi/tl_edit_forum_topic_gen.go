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

// EditForumTopicRequest represents TL type `editForumTopic#a7769060`.
type EditForumTopicRequest struct {
	// Identifier of the chat
	ChatID int64
	// Message thread identifier of the forum topic
	MessageThreadID int64
	// New name of the topic; 0-128 characters. If empty, the previous topic name is kept
	Name string
	// Pass true to edit the icon of the topic. Icon of the General topic can't be edited
	EditIconCustomEmoji bool
	// Identifier of the new custom emoji for topic icon; pass 0 to remove the custom emoji.
	// Ignored if edit_icon_custom_emoji is false. Telegram Premium users can use any custom
	// emoji, other users can use only a custom emoji returned by getForumTopicDefaultIcons
	IconCustomEmojiID int64
}

// EditForumTopicRequestTypeID is TL type id of EditForumTopicRequest.
const EditForumTopicRequestTypeID = 0xa7769060

// Ensuring interfaces in compile-time for EditForumTopicRequest.
var (
	_ bin.Encoder     = &EditForumTopicRequest{}
	_ bin.Decoder     = &EditForumTopicRequest{}
	_ bin.BareEncoder = &EditForumTopicRequest{}
	_ bin.BareDecoder = &EditForumTopicRequest{}
)

func (e *EditForumTopicRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.MessageThreadID == 0) {
		return false
	}
	if !(e.Name == "") {
		return false
	}
	if !(e.EditIconCustomEmoji == false) {
		return false
	}
	if !(e.IconCustomEmojiID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditForumTopicRequest) String() string {
	if e == nil {
		return "EditForumTopicRequest(nil)"
	}
	type Alias EditForumTopicRequest
	return fmt.Sprintf("EditForumTopicRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditForumTopicRequest) TypeID() uint32 {
	return EditForumTopicRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditForumTopicRequest) TypeName() string {
	return "editForumTopic"
}

// TypeInfo returns info about TL type.
func (e *EditForumTopicRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editForumTopic",
		ID:   EditForumTopicRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "EditIconCustomEmoji",
			SchemaName: "edit_icon_custom_emoji",
		},
		{
			Name:       "IconCustomEmojiID",
			SchemaName: "icon_custom_emoji_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditForumTopicRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editForumTopic#a7769060 as nil")
	}
	b.PutID(EditForumTopicRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditForumTopicRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editForumTopic#a7769060 as nil")
	}
	b.PutInt53(e.ChatID)
	b.PutInt53(e.MessageThreadID)
	b.PutString(e.Name)
	b.PutBool(e.EditIconCustomEmoji)
	b.PutLong(e.IconCustomEmojiID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EditForumTopicRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editForumTopic#a7769060 to nil")
	}
	if err := b.ConsumeID(EditForumTopicRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editForumTopic#a7769060: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditForumTopicRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editForumTopic#a7769060 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editForumTopic#a7769060: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editForumTopic#a7769060: field message_thread_id: %w", err)
		}
		e.MessageThreadID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode editForumTopic#a7769060: field name: %w", err)
		}
		e.Name = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode editForumTopic#a7769060: field edit_icon_custom_emoji: %w", err)
		}
		e.EditIconCustomEmoji = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode editForumTopic#a7769060: field icon_custom_emoji_id: %w", err)
		}
		e.IconCustomEmojiID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditForumTopicRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editForumTopic#a7769060 as nil")
	}
	b.ObjStart()
	b.PutID("editForumTopic")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(e.ChatID)
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(e.MessageThreadID)
	b.Comma()
	b.FieldStart("name")
	b.PutString(e.Name)
	b.Comma()
	b.FieldStart("edit_icon_custom_emoji")
	b.PutBool(e.EditIconCustomEmoji)
	b.Comma()
	b.FieldStart("icon_custom_emoji_id")
	b.PutLong(e.IconCustomEmojiID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditForumTopicRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editForumTopic#a7769060 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editForumTopic"); err != nil {
				return fmt.Errorf("unable to decode editForumTopic#a7769060: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editForumTopic#a7769060: field chat_id: %w", err)
			}
			e.ChatID = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editForumTopic#a7769060: field message_thread_id: %w", err)
			}
			e.MessageThreadID = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode editForumTopic#a7769060: field name: %w", err)
			}
			e.Name = value
		case "edit_icon_custom_emoji":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode editForumTopic#a7769060: field edit_icon_custom_emoji: %w", err)
			}
			e.EditIconCustomEmoji = value
		case "icon_custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode editForumTopic#a7769060: field icon_custom_emoji_id: %w", err)
			}
			e.IconCustomEmojiID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (e *EditForumTopicRequest) GetChatID() (value int64) {
	if e == nil {
		return
	}
	return e.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (e *EditForumTopicRequest) GetMessageThreadID() (value int64) {
	if e == nil {
		return
	}
	return e.MessageThreadID
}

// GetName returns value of Name field.
func (e *EditForumTopicRequest) GetName() (value string) {
	if e == nil {
		return
	}
	return e.Name
}

// GetEditIconCustomEmoji returns value of EditIconCustomEmoji field.
func (e *EditForumTopicRequest) GetEditIconCustomEmoji() (value bool) {
	if e == nil {
		return
	}
	return e.EditIconCustomEmoji
}

// GetIconCustomEmojiID returns value of IconCustomEmojiID field.
func (e *EditForumTopicRequest) GetIconCustomEmojiID() (value int64) {
	if e == nil {
		return
	}
	return e.IconCustomEmojiID
}

// EditForumTopic invokes method editForumTopic#a7769060 returning error if any.
func (c *Client) EditForumTopic(ctx context.Context, request *EditForumTopicRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}