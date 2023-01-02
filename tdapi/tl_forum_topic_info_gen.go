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

// ForumTopicInfo represents TL type `forumTopicInfo#8ff3df9e`.
type ForumTopicInfo struct {
	// Message thread identifier of the topic
	MessageThreadID int64
	// Name of the topic
	Name string
	// Icon of the topic
	Icon ForumTopicIcon
	// Date the topic was created
	CreationDate int32
	// Identifier of the creator of the topic
	CreatorID MessageSenderClass
	// True, if the topic is the General topic list
	IsGeneral bool
	// True, if the topic was created by the current user
	IsOutgoing bool
	// True, if the topic is closed
	IsClosed bool
	// True, if the topic is hidden above the topic list and closed; for General topic only
	IsHidden bool
}

// ForumTopicInfoTypeID is TL type id of ForumTopicInfo.
const ForumTopicInfoTypeID = 0x8ff3df9e

// Ensuring interfaces in compile-time for ForumTopicInfo.
var (
	_ bin.Encoder     = &ForumTopicInfo{}
	_ bin.Decoder     = &ForumTopicInfo{}
	_ bin.BareEncoder = &ForumTopicInfo{}
	_ bin.BareDecoder = &ForumTopicInfo{}
)

func (f *ForumTopicInfo) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.MessageThreadID == 0) {
		return false
	}
	if !(f.Name == "") {
		return false
	}
	if !(f.Icon.Zero()) {
		return false
	}
	if !(f.CreationDate == 0) {
		return false
	}
	if !(f.CreatorID == nil) {
		return false
	}
	if !(f.IsGeneral == false) {
		return false
	}
	if !(f.IsOutgoing == false) {
		return false
	}
	if !(f.IsClosed == false) {
		return false
	}
	if !(f.IsHidden == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *ForumTopicInfo) String() string {
	if f == nil {
		return "ForumTopicInfo(nil)"
	}
	type Alias ForumTopicInfo
	return fmt.Sprintf("ForumTopicInfo%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ForumTopicInfo) TypeID() uint32 {
	return ForumTopicInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*ForumTopicInfo) TypeName() string {
	return "forumTopicInfo"
}

// TypeInfo returns info about TL type.
func (f *ForumTopicInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "forumTopicInfo",
		ID:   ForumTopicInfoTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Icon",
			SchemaName: "icon",
		},
		{
			Name:       "CreationDate",
			SchemaName: "creation_date",
		},
		{
			Name:       "CreatorID",
			SchemaName: "creator_id",
		},
		{
			Name:       "IsGeneral",
			SchemaName: "is_general",
		},
		{
			Name:       "IsOutgoing",
			SchemaName: "is_outgoing",
		},
		{
			Name:       "IsClosed",
			SchemaName: "is_closed",
		},
		{
			Name:       "IsHidden",
			SchemaName: "is_hidden",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *ForumTopicInfo) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode forumTopicInfo#8ff3df9e as nil")
	}
	b.PutID(ForumTopicInfoTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *ForumTopicInfo) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode forumTopicInfo#8ff3df9e as nil")
	}
	b.PutInt53(f.MessageThreadID)
	b.PutString(f.Name)
	if err := f.Icon.Encode(b); err != nil {
		return fmt.Errorf("unable to encode forumTopicInfo#8ff3df9e: field icon: %w", err)
	}
	b.PutInt32(f.CreationDate)
	if f.CreatorID == nil {
		return fmt.Errorf("unable to encode forumTopicInfo#8ff3df9e: field creator_id is nil")
	}
	if err := f.CreatorID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode forumTopicInfo#8ff3df9e: field creator_id: %w", err)
	}
	b.PutBool(f.IsGeneral)
	b.PutBool(f.IsOutgoing)
	b.PutBool(f.IsClosed)
	b.PutBool(f.IsHidden)
	return nil
}

// Decode implements bin.Decoder.
func (f *ForumTopicInfo) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode forumTopicInfo#8ff3df9e to nil")
	}
	if err := b.ConsumeID(ForumTopicInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *ForumTopicInfo) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode forumTopicInfo#8ff3df9e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field message_thread_id: %w", err)
		}
		f.MessageThreadID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field name: %w", err)
		}
		f.Name = value
	}
	{
		if err := f.Icon.Decode(b); err != nil {
			return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field icon: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field creation_date: %w", err)
		}
		f.CreationDate = value
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field creator_id: %w", err)
		}
		f.CreatorID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field is_general: %w", err)
		}
		f.IsGeneral = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field is_outgoing: %w", err)
		}
		f.IsOutgoing = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field is_closed: %w", err)
		}
		f.IsClosed = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field is_hidden: %w", err)
		}
		f.IsHidden = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *ForumTopicInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode forumTopicInfo#8ff3df9e as nil")
	}
	b.ObjStart()
	b.PutID("forumTopicInfo")
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(f.MessageThreadID)
	b.Comma()
	b.FieldStart("name")
	b.PutString(f.Name)
	b.Comma()
	b.FieldStart("icon")
	if err := f.Icon.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode forumTopicInfo#8ff3df9e: field icon: %w", err)
	}
	b.Comma()
	b.FieldStart("creation_date")
	b.PutInt32(f.CreationDate)
	b.Comma()
	b.FieldStart("creator_id")
	if f.CreatorID == nil {
		return fmt.Errorf("unable to encode forumTopicInfo#8ff3df9e: field creator_id is nil")
	}
	if err := f.CreatorID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode forumTopicInfo#8ff3df9e: field creator_id: %w", err)
	}
	b.Comma()
	b.FieldStart("is_general")
	b.PutBool(f.IsGeneral)
	b.Comma()
	b.FieldStart("is_outgoing")
	b.PutBool(f.IsOutgoing)
	b.Comma()
	b.FieldStart("is_closed")
	b.PutBool(f.IsClosed)
	b.Comma()
	b.FieldStart("is_hidden")
	b.PutBool(f.IsHidden)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *ForumTopicInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode forumTopicInfo#8ff3df9e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("forumTopicInfo"); err != nil {
				return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: %w", err)
			}
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field message_thread_id: %w", err)
			}
			f.MessageThreadID = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field name: %w", err)
			}
			f.Name = value
		case "icon":
			if err := f.Icon.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field icon: %w", err)
			}
		case "creation_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field creation_date: %w", err)
			}
			f.CreationDate = value
		case "creator_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field creator_id: %w", err)
			}
			f.CreatorID = value
		case "is_general":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field is_general: %w", err)
			}
			f.IsGeneral = value
		case "is_outgoing":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field is_outgoing: %w", err)
			}
			f.IsOutgoing = value
		case "is_closed":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field is_closed: %w", err)
			}
			f.IsClosed = value
		case "is_hidden":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode forumTopicInfo#8ff3df9e: field is_hidden: %w", err)
			}
			f.IsHidden = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetMessageThreadID returns value of MessageThreadID field.
func (f *ForumTopicInfo) GetMessageThreadID() (value int64) {
	if f == nil {
		return
	}
	return f.MessageThreadID
}

// GetName returns value of Name field.
func (f *ForumTopicInfo) GetName() (value string) {
	if f == nil {
		return
	}
	return f.Name
}

// GetIcon returns value of Icon field.
func (f *ForumTopicInfo) GetIcon() (value ForumTopicIcon) {
	if f == nil {
		return
	}
	return f.Icon
}

// GetCreationDate returns value of CreationDate field.
func (f *ForumTopicInfo) GetCreationDate() (value int32) {
	if f == nil {
		return
	}
	return f.CreationDate
}

// GetCreatorID returns value of CreatorID field.
func (f *ForumTopicInfo) GetCreatorID() (value MessageSenderClass) {
	if f == nil {
		return
	}
	return f.CreatorID
}

// GetIsGeneral returns value of IsGeneral field.
func (f *ForumTopicInfo) GetIsGeneral() (value bool) {
	if f == nil {
		return
	}
	return f.IsGeneral
}

// GetIsOutgoing returns value of IsOutgoing field.
func (f *ForumTopicInfo) GetIsOutgoing() (value bool) {
	if f == nil {
		return
	}
	return f.IsOutgoing
}

// GetIsClosed returns value of IsClosed field.
func (f *ForumTopicInfo) GetIsClosed() (value bool) {
	if f == nil {
		return
	}
	return f.IsClosed
}

// GetIsHidden returns value of IsHidden field.
func (f *ForumTopicInfo) GetIsHidden() (value bool) {
	if f == nil {
		return
	}
	return f.IsHidden
}
