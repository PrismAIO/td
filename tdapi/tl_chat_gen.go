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

// Chat represents TL type `chat#4ade4c86`.
type Chat struct {
	// Chat unique identifier
	ID int64
	// Type of the chat
	Type ChatTypeClass
	// Chat title
	Title string
	// Chat photo; may be null
	Photo ChatPhotoInfo
	// Actions that non-administrator chat members are allowed to take in the chat
	Permissions ChatPermissions
	// Last message in the chat; may be null
	LastMessage Message
	// Positions of the chat in chat lists
	Positions []ChatPosition
	// Identifier of a user or chat that is selected to send messages in the chat; may be
	// null if the user can't change message sender
	MessageSenderID MessageSenderClass
	// True, if chat content can't be saved locally, forwarded, or copied
	HasProtectedContent bool
	// True, if the chat is marked as unread
	IsMarkedAsUnread bool
	// True, if the chat is blocked by the current user and private messages from the chat
	// can't be received
	IsBlocked bool
	// True, if the chat has scheduled messages
	HasScheduledMessages bool
	// True, if the chat messages can be deleted only for the current user while other users
	// will continue to see the messages
	CanBeDeletedOnlyForSelf bool
	// True, if the chat messages can be deleted for all users
	CanBeDeletedForAllUsers bool
	// True, if the chat can be reported to Telegram moderators through reportChat or
	// reportChatPhoto
	CanBeReported bool
	// Default value of the disable_notification parameter, used when a message is sent to
	// the chat
	DefaultDisableNotification bool
	// Number of unread messages in the chat
	UnreadCount int32
	// Identifier of the last read incoming message
	LastReadInboxMessageID int64
	// Identifier of the last read outgoing message
	LastReadOutboxMessageID int64
	// Number of unread messages with a mention/reply in the chat
	UnreadMentionCount int32
	// Number of messages with unread reactions in the chat
	UnreadReactionCount int32
	// Notification settings for the chat
	NotificationSettings ChatNotificationSettings
	// Types of reaction, available in the chat
	AvailableReactions ChatAvailableReactionsClass
	// Current message auto-delete or self-destruct timer setting for the chat, in seconds; 0
	// if disabled. Self-destruct timer in secret chats starts after the message or its
	// content is viewed. Auto-delete timer in other chats starts from the send date
	MessageAutoDeleteTime int32
	// If non-empty, name of a theme, set for the chat
	ThemeName string
	// Information about actions which must be possible to do through the chat action bar;
	// may be null
	ActionBar ChatActionBarClass
	// Information about video chat of the chat
	VideoChat VideoChat
	// Information about pending join requests; may be null
	PendingJoinRequests ChatJoinRequestsInfo
	// Identifier of the message from which reply markup needs to be used; 0 if there is no
	// default custom reply markup in the chat
	ReplyMarkupMessageID int64
	// A draft of a message in the chat; may be null
	DraftMessage DraftMessage
	// Application-specific data associated with the chat. (For example, the chat scroll
	// position or local chat notification settings can be stored here.) Persistent if the
	// message database is used
	ClientData string
}

// ChatTypeID is TL type id of Chat.
const ChatTypeID = 0x4ade4c86

// Ensuring interfaces in compile-time for Chat.
var (
	_ bin.Encoder     = &Chat{}
	_ bin.Decoder     = &Chat{}
	_ bin.BareEncoder = &Chat{}
	_ bin.BareDecoder = &Chat{}
)

func (c *Chat) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ID == 0) {
		return false
	}
	if !(c.Type == nil) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.Photo.Zero()) {
		return false
	}
	if !(c.Permissions.Zero()) {
		return false
	}
	if !(c.LastMessage.Zero()) {
		return false
	}
	if !(c.Positions == nil) {
		return false
	}
	if !(c.MessageSenderID == nil) {
		return false
	}
	if !(c.HasProtectedContent == false) {
		return false
	}
	if !(c.IsMarkedAsUnread == false) {
		return false
	}
	if !(c.IsBlocked == false) {
		return false
	}
	if !(c.HasScheduledMessages == false) {
		return false
	}
	if !(c.CanBeDeletedOnlyForSelf == false) {
		return false
	}
	if !(c.CanBeDeletedForAllUsers == false) {
		return false
	}
	if !(c.CanBeReported == false) {
		return false
	}
	if !(c.DefaultDisableNotification == false) {
		return false
	}
	if !(c.UnreadCount == 0) {
		return false
	}
	if !(c.LastReadInboxMessageID == 0) {
		return false
	}
	if !(c.LastReadOutboxMessageID == 0) {
		return false
	}
	if !(c.UnreadMentionCount == 0) {
		return false
	}
	if !(c.UnreadReactionCount == 0) {
		return false
	}
	if !(c.NotificationSettings.Zero()) {
		return false
	}
	if !(c.AvailableReactions == nil) {
		return false
	}
	if !(c.MessageAutoDeleteTime == 0) {
		return false
	}
	if !(c.ThemeName == "") {
		return false
	}
	if !(c.ActionBar == nil) {
		return false
	}
	if !(c.VideoChat.Zero()) {
		return false
	}
	if !(c.PendingJoinRequests.Zero()) {
		return false
	}
	if !(c.ReplyMarkupMessageID == 0) {
		return false
	}
	if !(c.DraftMessage.Zero()) {
		return false
	}
	if !(c.ClientData == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *Chat) String() string {
	if c == nil {
		return "Chat(nil)"
	}
	type Alias Chat
	return fmt.Sprintf("Chat%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Chat) TypeID() uint32 {
	return ChatTypeID
}

// TypeName returns name of type in TL schema.
func (*Chat) TypeName() string {
	return "chat"
}

// TypeInfo returns info about TL type.
func (c *Chat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chat",
		ID:   ChatTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "Permissions",
			SchemaName: "permissions",
		},
		{
			Name:       "LastMessage",
			SchemaName: "last_message",
		},
		{
			Name:       "Positions",
			SchemaName: "positions",
		},
		{
			Name:       "MessageSenderID",
			SchemaName: "message_sender_id",
		},
		{
			Name:       "HasProtectedContent",
			SchemaName: "has_protected_content",
		},
		{
			Name:       "IsMarkedAsUnread",
			SchemaName: "is_marked_as_unread",
		},
		{
			Name:       "IsBlocked",
			SchemaName: "is_blocked",
		},
		{
			Name:       "HasScheduledMessages",
			SchemaName: "has_scheduled_messages",
		},
		{
			Name:       "CanBeDeletedOnlyForSelf",
			SchemaName: "can_be_deleted_only_for_self",
		},
		{
			Name:       "CanBeDeletedForAllUsers",
			SchemaName: "can_be_deleted_for_all_users",
		},
		{
			Name:       "CanBeReported",
			SchemaName: "can_be_reported",
		},
		{
			Name:       "DefaultDisableNotification",
			SchemaName: "default_disable_notification",
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
			Name:       "AvailableReactions",
			SchemaName: "available_reactions",
		},
		{
			Name:       "MessageAutoDeleteTime",
			SchemaName: "message_auto_delete_time",
		},
		{
			Name:       "ThemeName",
			SchemaName: "theme_name",
		},
		{
			Name:       "ActionBar",
			SchemaName: "action_bar",
		},
		{
			Name:       "VideoChat",
			SchemaName: "video_chat",
		},
		{
			Name:       "PendingJoinRequests",
			SchemaName: "pending_join_requests",
		},
		{
			Name:       "ReplyMarkupMessageID",
			SchemaName: "reply_markup_message_id",
		},
		{
			Name:       "DraftMessage",
			SchemaName: "draft_message",
		},
		{
			Name:       "ClientData",
			SchemaName: "client_data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *Chat) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chat#4ade4c86 as nil")
	}
	b.PutID(ChatTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *Chat) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chat#4ade4c86 as nil")
	}
	b.PutInt53(c.ID)
	if c.Type == nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field type is nil")
	}
	if err := c.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field type: %w", err)
	}
	b.PutString(c.Title)
	if err := c.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field photo: %w", err)
	}
	if err := c.Permissions.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field permissions: %w", err)
	}
	if err := c.LastMessage.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field last_message: %w", err)
	}
	b.PutInt(len(c.Positions))
	for idx, v := range c.Positions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chat#4ade4c86: field positions element with index %d: %w", idx, err)
		}
	}
	if c.MessageSenderID == nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field message_sender_id is nil")
	}
	if err := c.MessageSenderID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field message_sender_id: %w", err)
	}
	b.PutBool(c.HasProtectedContent)
	b.PutBool(c.IsMarkedAsUnread)
	b.PutBool(c.IsBlocked)
	b.PutBool(c.HasScheduledMessages)
	b.PutBool(c.CanBeDeletedOnlyForSelf)
	b.PutBool(c.CanBeDeletedForAllUsers)
	b.PutBool(c.CanBeReported)
	b.PutBool(c.DefaultDisableNotification)
	b.PutInt32(c.UnreadCount)
	b.PutInt53(c.LastReadInboxMessageID)
	b.PutInt53(c.LastReadOutboxMessageID)
	b.PutInt32(c.UnreadMentionCount)
	b.PutInt32(c.UnreadReactionCount)
	if err := c.NotificationSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field notification_settings: %w", err)
	}
	if c.AvailableReactions == nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field available_reactions is nil")
	}
	if err := c.AvailableReactions.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field available_reactions: %w", err)
	}
	b.PutInt32(c.MessageAutoDeleteTime)
	b.PutString(c.ThemeName)
	if c.ActionBar == nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field action_bar is nil")
	}
	if err := c.ActionBar.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field action_bar: %w", err)
	}
	if err := c.VideoChat.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field video_chat: %w", err)
	}
	if err := c.PendingJoinRequests.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field pending_join_requests: %w", err)
	}
	b.PutInt53(c.ReplyMarkupMessageID)
	if err := c.DraftMessage.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field draft_message: %w", err)
	}
	b.PutString(c.ClientData)
	return nil
}

// Decode implements bin.Decoder.
func (c *Chat) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chat#4ade4c86 to nil")
	}
	if err := b.ConsumeID(ChatTypeID); err != nil {
		return fmt.Errorf("unable to decode chat#4ade4c86: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *Chat) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chat#4ade4c86 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := DecodeChatType(b)
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field type: %w", err)
		}
		c.Type = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field title: %w", err)
		}
		c.Title = value
	}
	{
		if err := c.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field photo: %w", err)
		}
	}
	{
		if err := c.Permissions.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field permissions: %w", err)
		}
	}
	{
		if err := c.LastMessage.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field last_message: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field positions: %w", err)
		}

		if headerLen > 0 {
			c.Positions = make([]ChatPosition, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatPosition
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chat#4ade4c86: field positions: %w", err)
			}
			c.Positions = append(c.Positions, value)
		}
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field message_sender_id: %w", err)
		}
		c.MessageSenderID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field has_protected_content: %w", err)
		}
		c.HasProtectedContent = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field is_marked_as_unread: %w", err)
		}
		c.IsMarkedAsUnread = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field is_blocked: %w", err)
		}
		c.IsBlocked = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field has_scheduled_messages: %w", err)
		}
		c.HasScheduledMessages = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field can_be_deleted_only_for_self: %w", err)
		}
		c.CanBeDeletedOnlyForSelf = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field can_be_deleted_for_all_users: %w", err)
		}
		c.CanBeDeletedForAllUsers = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field can_be_reported: %w", err)
		}
		c.CanBeReported = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field default_disable_notification: %w", err)
		}
		c.DefaultDisableNotification = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field unread_count: %w", err)
		}
		c.UnreadCount = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field last_read_inbox_message_id: %w", err)
		}
		c.LastReadInboxMessageID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field last_read_outbox_message_id: %w", err)
		}
		c.LastReadOutboxMessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field unread_mention_count: %w", err)
		}
		c.UnreadMentionCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field unread_reaction_count: %w", err)
		}
		c.UnreadReactionCount = value
	}
	{
		if err := c.NotificationSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field notification_settings: %w", err)
		}
	}
	{
		value, err := DecodeChatAvailableReactions(b)
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field available_reactions: %w", err)
		}
		c.AvailableReactions = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field message_auto_delete_time: %w", err)
		}
		c.MessageAutoDeleteTime = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field theme_name: %w", err)
		}
		c.ThemeName = value
	}
	{
		value, err := DecodeChatActionBar(b)
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field action_bar: %w", err)
		}
		c.ActionBar = value
	}
	{
		if err := c.VideoChat.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field video_chat: %w", err)
		}
	}
	{
		if err := c.PendingJoinRequests.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field pending_join_requests: %w", err)
		}
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field reply_markup_message_id: %w", err)
		}
		c.ReplyMarkupMessageID = value
	}
	{
		if err := c.DraftMessage.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field draft_message: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chat#4ade4c86: field client_data: %w", err)
		}
		c.ClientData = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *Chat) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chat#4ade4c86 as nil")
	}
	b.ObjStart()
	b.PutID("chat")
	b.Comma()
	b.FieldStart("id")
	b.PutInt53(c.ID)
	b.Comma()
	b.FieldStart("type")
	if c.Type == nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field type is nil")
	}
	if err := c.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("title")
	b.PutString(c.Title)
	b.Comma()
	b.FieldStart("photo")
	if err := c.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field photo: %w", err)
	}
	b.Comma()
	b.FieldStart("permissions")
	if err := c.Permissions.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field permissions: %w", err)
	}
	b.Comma()
	b.FieldStart("last_message")
	if err := c.LastMessage.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field last_message: %w", err)
	}
	b.Comma()
	b.FieldStart("positions")
	b.ArrStart()
	for idx, v := range c.Positions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chat#4ade4c86: field positions element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("message_sender_id")
	if c.MessageSenderID == nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field message_sender_id is nil")
	}
	if err := c.MessageSenderID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field message_sender_id: %w", err)
	}
	b.Comma()
	b.FieldStart("has_protected_content")
	b.PutBool(c.HasProtectedContent)
	b.Comma()
	b.FieldStart("is_marked_as_unread")
	b.PutBool(c.IsMarkedAsUnread)
	b.Comma()
	b.FieldStart("is_blocked")
	b.PutBool(c.IsBlocked)
	b.Comma()
	b.FieldStart("has_scheduled_messages")
	b.PutBool(c.HasScheduledMessages)
	b.Comma()
	b.FieldStart("can_be_deleted_only_for_self")
	b.PutBool(c.CanBeDeletedOnlyForSelf)
	b.Comma()
	b.FieldStart("can_be_deleted_for_all_users")
	b.PutBool(c.CanBeDeletedForAllUsers)
	b.Comma()
	b.FieldStart("can_be_reported")
	b.PutBool(c.CanBeReported)
	b.Comma()
	b.FieldStart("default_disable_notification")
	b.PutBool(c.DefaultDisableNotification)
	b.Comma()
	b.FieldStart("unread_count")
	b.PutInt32(c.UnreadCount)
	b.Comma()
	b.FieldStart("last_read_inbox_message_id")
	b.PutInt53(c.LastReadInboxMessageID)
	b.Comma()
	b.FieldStart("last_read_outbox_message_id")
	b.PutInt53(c.LastReadOutboxMessageID)
	b.Comma()
	b.FieldStart("unread_mention_count")
	b.PutInt32(c.UnreadMentionCount)
	b.Comma()
	b.FieldStart("unread_reaction_count")
	b.PutInt32(c.UnreadReactionCount)
	b.Comma()
	b.FieldStart("notification_settings")
	if err := c.NotificationSettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field notification_settings: %w", err)
	}
	b.Comma()
	b.FieldStart("available_reactions")
	if c.AvailableReactions == nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field available_reactions is nil")
	}
	if err := c.AvailableReactions.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field available_reactions: %w", err)
	}
	b.Comma()
	b.FieldStart("message_auto_delete_time")
	b.PutInt32(c.MessageAutoDeleteTime)
	b.Comma()
	b.FieldStart("theme_name")
	b.PutString(c.ThemeName)
	b.Comma()
	b.FieldStart("action_bar")
	if c.ActionBar == nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field action_bar is nil")
	}
	if err := c.ActionBar.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field action_bar: %w", err)
	}
	b.Comma()
	b.FieldStart("video_chat")
	if err := c.VideoChat.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field video_chat: %w", err)
	}
	b.Comma()
	b.FieldStart("pending_join_requests")
	if err := c.PendingJoinRequests.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field pending_join_requests: %w", err)
	}
	b.Comma()
	b.FieldStart("reply_markup_message_id")
	b.PutInt53(c.ReplyMarkupMessageID)
	b.Comma()
	b.FieldStart("draft_message")
	if err := c.DraftMessage.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chat#4ade4c86: field draft_message: %w", err)
	}
	b.Comma()
	b.FieldStart("client_data")
	b.PutString(c.ClientData)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *Chat) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chat#4ade4c86 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chat"); err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: %w", err)
			}
		case "id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field id: %w", err)
			}
			c.ID = value
		case "type":
			value, err := DecodeTDLibJSONChatType(b)
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field type: %w", err)
			}
			c.Type = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field title: %w", err)
			}
			c.Title = value
		case "photo":
			if err := c.Photo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field photo: %w", err)
			}
		case "permissions":
			if err := c.Permissions.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field permissions: %w", err)
			}
		case "last_message":
			if err := c.LastMessage.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field last_message: %w", err)
			}
		case "positions":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value ChatPosition
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode chat#4ade4c86: field positions: %w", err)
				}
				c.Positions = append(c.Positions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field positions: %w", err)
			}
		case "message_sender_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field message_sender_id: %w", err)
			}
			c.MessageSenderID = value
		case "has_protected_content":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field has_protected_content: %w", err)
			}
			c.HasProtectedContent = value
		case "is_marked_as_unread":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field is_marked_as_unread: %w", err)
			}
			c.IsMarkedAsUnread = value
		case "is_blocked":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field is_blocked: %w", err)
			}
			c.IsBlocked = value
		case "has_scheduled_messages":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field has_scheduled_messages: %w", err)
			}
			c.HasScheduledMessages = value
		case "can_be_deleted_only_for_self":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field can_be_deleted_only_for_self: %w", err)
			}
			c.CanBeDeletedOnlyForSelf = value
		case "can_be_deleted_for_all_users":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field can_be_deleted_for_all_users: %w", err)
			}
			c.CanBeDeletedForAllUsers = value
		case "can_be_reported":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field can_be_reported: %w", err)
			}
			c.CanBeReported = value
		case "default_disable_notification":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field default_disable_notification: %w", err)
			}
			c.DefaultDisableNotification = value
		case "unread_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field unread_count: %w", err)
			}
			c.UnreadCount = value
		case "last_read_inbox_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field last_read_inbox_message_id: %w", err)
			}
			c.LastReadInboxMessageID = value
		case "last_read_outbox_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field last_read_outbox_message_id: %w", err)
			}
			c.LastReadOutboxMessageID = value
		case "unread_mention_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field unread_mention_count: %w", err)
			}
			c.UnreadMentionCount = value
		case "unread_reaction_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field unread_reaction_count: %w", err)
			}
			c.UnreadReactionCount = value
		case "notification_settings":
			if err := c.NotificationSettings.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field notification_settings: %w", err)
			}
		case "available_reactions":
			value, err := DecodeTDLibJSONChatAvailableReactions(b)
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field available_reactions: %w", err)
			}
			c.AvailableReactions = value
		case "message_auto_delete_time":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field message_auto_delete_time: %w", err)
			}
			c.MessageAutoDeleteTime = value
		case "theme_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field theme_name: %w", err)
			}
			c.ThemeName = value
		case "action_bar":
			value, err := DecodeTDLibJSONChatActionBar(b)
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field action_bar: %w", err)
			}
			c.ActionBar = value
		case "video_chat":
			if err := c.VideoChat.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field video_chat: %w", err)
			}
		case "pending_join_requests":
			if err := c.PendingJoinRequests.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field pending_join_requests: %w", err)
			}
		case "reply_markup_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field reply_markup_message_id: %w", err)
			}
			c.ReplyMarkupMessageID = value
		case "draft_message":
			if err := c.DraftMessage.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field draft_message: %w", err)
			}
		case "client_data":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chat#4ade4c86: field client_data: %w", err)
			}
			c.ClientData = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (c *Chat) GetID() (value int64) {
	if c == nil {
		return
	}
	return c.ID
}

// GetType returns value of Type field.
func (c *Chat) GetType() (value ChatTypeClass) {
	if c == nil {
		return
	}
	return c.Type
}

// GetTitle returns value of Title field.
func (c *Chat) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

// GetPhoto returns value of Photo field.
func (c *Chat) GetPhoto() (value ChatPhotoInfo) {
	if c == nil {
		return
	}
	return c.Photo
}

// GetPermissions returns value of Permissions field.
func (c *Chat) GetPermissions() (value ChatPermissions) {
	if c == nil {
		return
	}
	return c.Permissions
}

// GetLastMessage returns value of LastMessage field.
func (c *Chat) GetLastMessage() (value Message) {
	if c == nil {
		return
	}
	return c.LastMessage
}

// GetPositions returns value of Positions field.
func (c *Chat) GetPositions() (value []ChatPosition) {
	if c == nil {
		return
	}
	return c.Positions
}

// GetMessageSenderID returns value of MessageSenderID field.
func (c *Chat) GetMessageSenderID() (value MessageSenderClass) {
	if c == nil {
		return
	}
	return c.MessageSenderID
}

// GetHasProtectedContent returns value of HasProtectedContent field.
func (c *Chat) GetHasProtectedContent() (value bool) {
	if c == nil {
		return
	}
	return c.HasProtectedContent
}

// GetIsMarkedAsUnread returns value of IsMarkedAsUnread field.
func (c *Chat) GetIsMarkedAsUnread() (value bool) {
	if c == nil {
		return
	}
	return c.IsMarkedAsUnread
}

// GetIsBlocked returns value of IsBlocked field.
func (c *Chat) GetIsBlocked() (value bool) {
	if c == nil {
		return
	}
	return c.IsBlocked
}

// GetHasScheduledMessages returns value of HasScheduledMessages field.
func (c *Chat) GetHasScheduledMessages() (value bool) {
	if c == nil {
		return
	}
	return c.HasScheduledMessages
}

// GetCanBeDeletedOnlyForSelf returns value of CanBeDeletedOnlyForSelf field.
func (c *Chat) GetCanBeDeletedOnlyForSelf() (value bool) {
	if c == nil {
		return
	}
	return c.CanBeDeletedOnlyForSelf
}

// GetCanBeDeletedForAllUsers returns value of CanBeDeletedForAllUsers field.
func (c *Chat) GetCanBeDeletedForAllUsers() (value bool) {
	if c == nil {
		return
	}
	return c.CanBeDeletedForAllUsers
}

// GetCanBeReported returns value of CanBeReported field.
func (c *Chat) GetCanBeReported() (value bool) {
	if c == nil {
		return
	}
	return c.CanBeReported
}

// GetDefaultDisableNotification returns value of DefaultDisableNotification field.
func (c *Chat) GetDefaultDisableNotification() (value bool) {
	if c == nil {
		return
	}
	return c.DefaultDisableNotification
}

// GetUnreadCount returns value of UnreadCount field.
func (c *Chat) GetUnreadCount() (value int32) {
	if c == nil {
		return
	}
	return c.UnreadCount
}

// GetLastReadInboxMessageID returns value of LastReadInboxMessageID field.
func (c *Chat) GetLastReadInboxMessageID() (value int64) {
	if c == nil {
		return
	}
	return c.LastReadInboxMessageID
}

// GetLastReadOutboxMessageID returns value of LastReadOutboxMessageID field.
func (c *Chat) GetLastReadOutboxMessageID() (value int64) {
	if c == nil {
		return
	}
	return c.LastReadOutboxMessageID
}

// GetUnreadMentionCount returns value of UnreadMentionCount field.
func (c *Chat) GetUnreadMentionCount() (value int32) {
	if c == nil {
		return
	}
	return c.UnreadMentionCount
}

// GetUnreadReactionCount returns value of UnreadReactionCount field.
func (c *Chat) GetUnreadReactionCount() (value int32) {
	if c == nil {
		return
	}
	return c.UnreadReactionCount
}

// GetNotificationSettings returns value of NotificationSettings field.
func (c *Chat) GetNotificationSettings() (value ChatNotificationSettings) {
	if c == nil {
		return
	}
	return c.NotificationSettings
}

// GetAvailableReactions returns value of AvailableReactions field.
func (c *Chat) GetAvailableReactions() (value ChatAvailableReactionsClass) {
	if c == nil {
		return
	}
	return c.AvailableReactions
}

// GetMessageAutoDeleteTime returns value of MessageAutoDeleteTime field.
func (c *Chat) GetMessageAutoDeleteTime() (value int32) {
	if c == nil {
		return
	}
	return c.MessageAutoDeleteTime
}

// GetThemeName returns value of ThemeName field.
func (c *Chat) GetThemeName() (value string) {
	if c == nil {
		return
	}
	return c.ThemeName
}

// GetActionBar returns value of ActionBar field.
func (c *Chat) GetActionBar() (value ChatActionBarClass) {
	if c == nil {
		return
	}
	return c.ActionBar
}

// GetVideoChat returns value of VideoChat field.
func (c *Chat) GetVideoChat() (value VideoChat) {
	if c == nil {
		return
	}
	return c.VideoChat
}

// GetPendingJoinRequests returns value of PendingJoinRequests field.
func (c *Chat) GetPendingJoinRequests() (value ChatJoinRequestsInfo) {
	if c == nil {
		return
	}
	return c.PendingJoinRequests
}

// GetReplyMarkupMessageID returns value of ReplyMarkupMessageID field.
func (c *Chat) GetReplyMarkupMessageID() (value int64) {
	if c == nil {
		return
	}
	return c.ReplyMarkupMessageID
}

// GetDraftMessage returns value of DraftMessage field.
func (c *Chat) GetDraftMessage() (value DraftMessage) {
	if c == nil {
		return
	}
	return c.DraftMessage
}

// GetClientData returns value of ClientData field.
func (c *Chat) GetClientData() (value string) {
	if c == nil {
		return
	}
	return c.ClientData
}
