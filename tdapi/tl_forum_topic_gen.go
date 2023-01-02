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

// ForumTopic represents TL type `forumTopic#1213ace6`.
type ForumTopic struct {
	// Basic information about the topic
	Info ForumTopicInfo
	// Last message in the topic; may be null if unknown
	LastMessage Message
	// True, if the topic is pinned in the topic list
	IsPinned bool
	// Number of unread messages in the topic
	UnreadCount int32
	// Identifier of the last read incoming message
	LastReadInboxMessageID int64
	// Identifier of the last read outgoing message
	LastReadOutboxMessageID int64
	// Number of unread messages with a mention/reply in the topic
	UnreadMentionCount int32
	// Number of messages with unread reactions in the topic
	UnreadReactionCount int32
	// Notification settings for the topic
	NotificationSettings ChatNotificationSettings
	// A draft of a message in the topic; may be null
	DraftMessage DraftMessage
}

// ForumTopicTypeID is TL type id of ForumTopic.
const ForumTopicTypeID = 0x1213ace6

// Ensuring interfaces in compile-time for ForumTopic.
var (
	_ bin.Encoder     = &ForumTopic{}
	_ bin.Decoder     = &ForumTopic{}
	_ bin.BareEncoder = &ForumTopic{}
	_ bin.BareDecoder = &ForumTopic{}
)

func (f *ForumTopic) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Info.Zero()) {
		return false
	}
	if !(f.LastMessage.Zero()) {
		return false
	}
	if !(f.IsPinned == false) {
		return false
	}
	if !(f.UnreadCount == 0) {
		return false
	}
	if !(f.LastReadInboxMessageID == 0) {
		return false
	}
	if !(f.LastReadOutboxMessageID == 0) {
		return false
	}
	if !(f.UnreadMentionCount == 0) {
		return false
	}
	if !(f.UnreadReactionCount == 0) {
		return false
	}
	if !(f.NotificationSettings.Zero()) {
		return false
	}
	if !(f.DraftMessage.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *ForumTopic) String() string {
	if f == nil {
		return "ForumTopic(nil)"
	}
	type Alias ForumTopic
	return fmt.Sprintf("ForumTopic%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ForumTopic) TypeID() uint32 {
	return ForumTopicTypeID
}

// TypeName returns name of type in TL schema.
func (*ForumTopic) TypeName() string {
	return "forumTopic"
}

// TypeInfo returns info about TL type.
func (f *ForumTopic) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "forumTopic",
		ID:   ForumTopicTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Info",
			SchemaName: "info",
		},
		{
			Name:       "LastMessage",
			SchemaName: "last_message",
		},
		{
			Name:       "IsPinned",
			SchemaName: "is_pinned",
		},
		{
			Name:       "UnreadCount",
			SchemaName: "unread_count",
		},
		{
			Name:       "LastReadInboxMessageID",
			SchemaName: "last_read_inbox_message_id",
		},
		{
			Name:       "LastReadOutboxMessageID",
			SchemaName: "last_read_outbox_message_id",
		},
		{
			Name:       "UnreadMentionCount",
			SchemaName: "unread_mention_count",
		},
		{
			Name:       "UnreadReactionCount",
			SchemaName: "unread_reaction_count",
		},
		{
			Name:       "NotificationSettings",
			SchemaName: "notification_settings",
		},
		{
			Name:       "DraftMessage",
			SchemaName: "draft_message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *ForumTopic) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode forumTopic#1213ace6 as nil")
	}
	b.PutID(ForumTopicTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *ForumTopic) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode forumTopic#1213ace6 as nil")
	}
	if err := f.Info.Encode(b); err != nil {
		return fmt.Errorf("unable to encode forumTopic#1213ace6: field info: %w", err)
	}
	if err := f.LastMessage.Encode(b); err != nil {
		return fmt.Errorf("unable to encode forumTopic#1213ace6: field last_message: %w", err)
	}
	b.PutBool(f.IsPinned)
	b.PutInt32(f.UnreadCount)
	b.PutInt53(f.LastReadInboxMessageID)
	b.PutInt53(f.LastReadOutboxMessageID)
	b.PutInt32(f.UnreadMentionCount)
	b.PutInt32(f.UnreadReactionCount)
	if err := f.NotificationSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode forumTopic#1213ace6: field notification_settings: %w", err)
	}
	if err := f.DraftMessage.Encode(b); err != nil {
		return fmt.Errorf("unable to encode forumTopic#1213ace6: field draft_message: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *ForumTopic) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode forumTopic#1213ace6 to nil")
	}
	if err := b.ConsumeID(ForumTopicTypeID); err != nil {
		return fmt.Errorf("unable to decode forumTopic#1213ace6: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *ForumTopic) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode forumTopic#1213ace6 to nil")
	}
	{
		if err := f.Info.Decode(b); err != nil {
			return fmt.Errorf("unable to decode forumTopic#1213ace6: field info: %w", err)
		}
	}
	{
		if err := f.LastMessage.Decode(b); err != nil {
			return fmt.Errorf("unable to decode forumTopic#1213ace6: field last_message: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopic#1213ace6: field is_pinned: %w", err)
		}
		f.IsPinned = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopic#1213ace6: field unread_count: %w", err)
		}
		f.UnreadCount = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopic#1213ace6: field last_read_inbox_message_id: %w", err)
		}
		f.LastReadInboxMessageID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopic#1213ace6: field last_read_outbox_message_id: %w", err)
		}
		f.LastReadOutboxMessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopic#1213ace6: field unread_mention_count: %w", err)
		}
		f.UnreadMentionCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopic#1213ace6: field unread_reaction_count: %w", err)
		}
		f.UnreadReactionCount = value
	}
	{
		if err := f.NotificationSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode forumTopic#1213ace6: field notification_settings: %w", err)
		}
	}
	{
		if err := f.DraftMessage.Decode(b); err != nil {
			return fmt.Errorf("unable to decode forumTopic#1213ace6: field draft_message: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *ForumTopic) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode forumTopic#1213ace6 as nil")
	}
	b.ObjStart()
	b.PutID("forumTopic")
	b.Comma()
	b.FieldStart("info")
	if err := f.Info.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode forumTopic#1213ace6: field info: %w", err)
	}
	b.Comma()
	b.FieldStart("last_message")
	if err := f.LastMessage.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode forumTopic#1213ace6: field last_message: %w", err)
	}
	b.Comma()
	b.FieldStart("is_pinned")
	b.PutBool(f.IsPinned)
	b.Comma()
	b.FieldStart("unread_count")
	b.PutInt32(f.UnreadCount)
	b.Comma()
	b.FieldStart("last_read_inbox_message_id")
	b.PutInt53(f.LastReadInboxMessageID)
	b.Comma()
	b.FieldStart("last_read_outbox_message_id")
	b.PutInt53(f.LastReadOutboxMessageID)
	b.Comma()
	b.FieldStart("unread_mention_count")
	b.PutInt32(f.UnreadMentionCount)
	b.Comma()
	b.FieldStart("unread_reaction_count")
	b.PutInt32(f.UnreadReactionCount)
	b.Comma()
	b.FieldStart("notification_settings")
	if err := f.NotificationSettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode forumTopic#1213ace6: field notification_settings: %w", err)
	}
	b.Comma()
	b.FieldStart("draft_message")
	if err := f.DraftMessage.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode forumTopic#1213ace6: field draft_message: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *ForumTopic) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode forumTopic#1213ace6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("forumTopic"); err != nil {
				return fmt.Errorf("unable to decode forumTopic#1213ace6: %w", err)
			}
		case "info":
			if err := f.Info.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode forumTopic#1213ace6: field info: %w", err)
			}
		case "last_message":
			if err := f.LastMessage.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode forumTopic#1213ace6: field last_message: %w", err)
			}
		case "is_pinned":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopic#1213ace6: field is_pinned: %w", err)
			}
			f.IsPinned = value
		case "unread_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopic#1213ace6: field unread_count: %w", err)
			}
			f.UnreadCount = value
		case "last_read_inbox_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopic#1213ace6: field last_read_inbox_message_id: %w", err)
			}
			f.LastReadInboxMessageID = value
		case "last_read_outbox_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopic#1213ace6: field last_read_outbox_message_id: %w", err)
			}
			f.LastReadOutboxMessageID = value
		case "unread_mention_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopic#1213ace6: field unread_mention_count: %w", err)
			}
			f.UnreadMentionCount = value
		case "unread_reaction_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopic#1213ace6: field unread_reaction_count: %w", err)
			}
			f.UnreadReactionCount = value
		case "notification_settings":
			if err := f.NotificationSettings.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode forumTopic#1213ace6: field notification_settings: %w", err)
			}
		case "draft_message":
			if err := f.DraftMessage.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode forumTopic#1213ace6: field draft_message: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInfo returns value of Info field.
func (f *ForumTopic) GetInfo() (value ForumTopicInfo) {
	if f == nil {
		return
	}
	return f.Info
}

// GetLastMessage returns value of LastMessage field.
func (f *ForumTopic) GetLastMessage() (value Message) {
	if f == nil {
		return
	}
	return f.LastMessage
}

// GetIsPinned returns value of IsPinned field.
func (f *ForumTopic) GetIsPinned() (value bool) {
	if f == nil {
		return
	}
	return f.IsPinned
}

// GetUnreadCount returns value of UnreadCount field.
func (f *ForumTopic) GetUnreadCount() (value int32) {
	if f == nil {
		return
	}
	return f.UnreadCount
}

// GetLastReadInboxMessageID returns value of LastReadInboxMessageID field.
func (f *ForumTopic) GetLastReadInboxMessageID() (value int64) {
	if f == nil {
		return
	}
	return f.LastReadInboxMessageID
}

// GetLastReadOutboxMessageID returns value of LastReadOutboxMessageID field.
func (f *ForumTopic) GetLastReadOutboxMessageID() (value int64) {
	if f == nil {
		return
	}
	return f.LastReadOutboxMessageID
}

// GetUnreadMentionCount returns value of UnreadMentionCount field.
func (f *ForumTopic) GetUnreadMentionCount() (value int32) {
	if f == nil {
		return
	}
	return f.UnreadMentionCount
}

// GetUnreadReactionCount returns value of UnreadReactionCount field.
func (f *ForumTopic) GetUnreadReactionCount() (value int32) {
	if f == nil {
		return
	}
	return f.UnreadReactionCount
}

// GetNotificationSettings returns value of NotificationSettings field.
func (f *ForumTopic) GetNotificationSettings() (value ChatNotificationSettings) {
	if f == nil {
		return
	}
	return f.NotificationSettings
}

// GetDraftMessage returns value of DraftMessage field.
func (f *ForumTopic) GetDraftMessage() (value DraftMessage) {
	if f == nil {
		return
	}
	return f.DraftMessage
}
