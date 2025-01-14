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

// SponsoredMessage represents TL type `sponsoredMessage#db717e75`.
type SponsoredMessage struct {
	// Message identifier; unique for the chat to which the sponsored message belongs among
	// both ordinary and sponsored messages
	MessageID int64
	// True, if the message needs to be labeled as "recommended" instead of "sponsored"
	IsRecommended bool
	// Sponsor chat identifier; 0 if the sponsor chat is accessible through an invite link
	SponsorChatID int64
	// Information about the sponsor chat; may be null unless sponsor_chat_id == 0
	SponsorChatInfo ChatInviteLinkInfo
	// True, if the sponsor's chat photo must be shown
	ShowChatPhoto bool
	// An internal link to be opened when the sponsored message is clicked; may be null if
	// the sponsor chat needs to be opened instead
	Link InternalLinkTypeClass
	// Content of the message. Currently, can be only of the type messageText
	Content MessageContentClass
}

// SponsoredMessageTypeID is TL type id of SponsoredMessage.
const SponsoredMessageTypeID = 0xdb717e75

// Ensuring interfaces in compile-time for SponsoredMessage.
var (
	_ bin.Encoder     = &SponsoredMessage{}
	_ bin.Decoder     = &SponsoredMessage{}
	_ bin.BareEncoder = &SponsoredMessage{}
	_ bin.BareDecoder = &SponsoredMessage{}
)

func (s *SponsoredMessage) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.MessageID == 0) {
		return false
	}
	if !(s.IsRecommended == false) {
		return false
	}
	if !(s.SponsorChatID == 0) {
		return false
	}
	if !(s.SponsorChatInfo.Zero()) {
		return false
	}
	if !(s.ShowChatPhoto == false) {
		return false
	}
	if !(s.Link == nil) {
		return false
	}
	if !(s.Content == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SponsoredMessage) String() string {
	if s == nil {
		return "SponsoredMessage(nil)"
	}
	type Alias SponsoredMessage
	return fmt.Sprintf("SponsoredMessage%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SponsoredMessage) TypeID() uint32 {
	return SponsoredMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*SponsoredMessage) TypeName() string {
	return "sponsoredMessage"
}

// TypeInfo returns info about TL type.
func (s *SponsoredMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sponsoredMessage",
		ID:   SponsoredMessageTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "IsRecommended",
			SchemaName: "is_recommended",
		},
		{
			Name:       "SponsorChatID",
			SchemaName: "sponsor_chat_id",
		},
		{
			Name:       "SponsorChatInfo",
			SchemaName: "sponsor_chat_info",
		},
		{
			Name:       "ShowChatPhoto",
			SchemaName: "show_chat_photo",
		},
		{
			Name:       "Link",
			SchemaName: "link",
		},
		{
			Name:       "Content",
			SchemaName: "content",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SponsoredMessage) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#db717e75 as nil")
	}
	b.PutID(SponsoredMessageTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SponsoredMessage) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#db717e75 as nil")
	}
	b.PutInt53(s.MessageID)
	b.PutBool(s.IsRecommended)
	b.PutInt53(s.SponsorChatID)
	if err := s.SponsorChatInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#db717e75: field sponsor_chat_info: %w", err)
	}
	b.PutBool(s.ShowChatPhoto)
	if s.Link == nil {
		return fmt.Errorf("unable to encode sponsoredMessage#db717e75: field link is nil")
	}
	if err := s.Link.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#db717e75: field link: %w", err)
	}
	if s.Content == nil {
		return fmt.Errorf("unable to encode sponsoredMessage#db717e75: field content is nil")
	}
	if err := s.Content.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#db717e75: field content: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SponsoredMessage) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#db717e75 to nil")
	}
	if err := b.ConsumeID(SponsoredMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode sponsoredMessage#db717e75: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SponsoredMessage) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#db717e75 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field message_id: %w", err)
		}
		s.MessageID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field is_recommended: %w", err)
		}
		s.IsRecommended = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field sponsor_chat_id: %w", err)
		}
		s.SponsorChatID = value
	}
	{
		if err := s.SponsorChatInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field sponsor_chat_info: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field show_chat_photo: %w", err)
		}
		s.ShowChatPhoto = value
	}
	{
		value, err := DecodeInternalLinkType(b)
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field link: %w", err)
		}
		s.Link = value
	}
	{
		value, err := DecodeMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field content: %w", err)
		}
		s.Content = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SponsoredMessage) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#db717e75 as nil")
	}
	b.ObjStart()
	b.PutID("sponsoredMessage")
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(s.MessageID)
	b.Comma()
	b.FieldStart("is_recommended")
	b.PutBool(s.IsRecommended)
	b.Comma()
	b.FieldStart("sponsor_chat_id")
	b.PutInt53(s.SponsorChatID)
	b.Comma()
	b.FieldStart("sponsor_chat_info")
	if err := s.SponsorChatInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#db717e75: field sponsor_chat_info: %w", err)
	}
	b.Comma()
	b.FieldStart("show_chat_photo")
	b.PutBool(s.ShowChatPhoto)
	b.Comma()
	b.FieldStart("link")
	if s.Link == nil {
		return fmt.Errorf("unable to encode sponsoredMessage#db717e75: field link is nil")
	}
	if err := s.Link.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#db717e75: field link: %w", err)
	}
	b.Comma()
	b.FieldStart("content")
	if s.Content == nil {
		return fmt.Errorf("unable to encode sponsoredMessage#db717e75: field content is nil")
	}
	if err := s.Content.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#db717e75: field content: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SponsoredMessage) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#db717e75 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sponsoredMessage"); err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#db717e75: %w", err)
			}
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field message_id: %w", err)
			}
			s.MessageID = value
		case "is_recommended":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field is_recommended: %w", err)
			}
			s.IsRecommended = value
		case "sponsor_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field sponsor_chat_id: %w", err)
			}
			s.SponsorChatID = value
		case "sponsor_chat_info":
			if err := s.SponsorChatInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field sponsor_chat_info: %w", err)
			}
		case "show_chat_photo":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field show_chat_photo: %w", err)
			}
			s.ShowChatPhoto = value
		case "link":
			value, err := DecodeTDLibJSONInternalLinkType(b)
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field link: %w", err)
			}
			s.Link = value
		case "content":
			value, err := DecodeTDLibJSONMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#db717e75: field content: %w", err)
			}
			s.Content = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetMessageID returns value of MessageID field.
func (s *SponsoredMessage) GetMessageID() (value int64) {
	if s == nil {
		return
	}
	return s.MessageID
}

// GetIsRecommended returns value of IsRecommended field.
func (s *SponsoredMessage) GetIsRecommended() (value bool) {
	if s == nil {
		return
	}
	return s.IsRecommended
}

// GetSponsorChatID returns value of SponsorChatID field.
func (s *SponsoredMessage) GetSponsorChatID() (value int64) {
	if s == nil {
		return
	}
	return s.SponsorChatID
}

// GetSponsorChatInfo returns value of SponsorChatInfo field.
func (s *SponsoredMessage) GetSponsorChatInfo() (value ChatInviteLinkInfo) {
	if s == nil {
		return
	}
	return s.SponsorChatInfo
}

// GetShowChatPhoto returns value of ShowChatPhoto field.
func (s *SponsoredMessage) GetShowChatPhoto() (value bool) {
	if s == nil {
		return
	}
	return s.ShowChatPhoto
}

// GetLink returns value of Link field.
func (s *SponsoredMessage) GetLink() (value InternalLinkTypeClass) {
	if s == nil {
		return
	}
	return s.Link
}

// GetContent returns value of Content field.
func (s *SponsoredMessage) GetContent() (value MessageContentClass) {
	if s == nil {
		return
	}
	return s.Content
}
