// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessageFwdHeader represents TL type `messageFwdHeader#5f777dce`.
// Info about a forwarded message
//
// See https://core.telegram.org/constructor/messageFwdHeader for reference.
type MessageFwdHeader struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this message was imported from a foreign chat service, click here for more
	// info »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/import
	Imported bool
	// The ID of the user that originally sent the message
	//
	// Use SetFromID and GetFromID helpers.
	FromID PeerClass
	// The name of the user that originally sent the message
	//
	// Use SetFromName and GetFromName helpers.
	FromName string
	// When was the message originally sent
	Date int
	// ID of the channel message that was forwarded
	//
	// Use SetChannelPost and GetChannelPost helpers.
	ChannelPost int
	// For channels and if signatures are enabled, author of the channel message
	//
	// Use SetPostAuthor and GetPostAuthor helpers.
	PostAuthor string
	// Only for messages forwarded to the current user (inputPeerSelf), full info about the
	// user/channel that originally sent the message
	//
	// Use SetSavedFromPeer and GetSavedFromPeer helpers.
	SavedFromPeer PeerClass
	// Only for messages forwarded to the current user (inputPeerSelf), ID of the message
	// that was forwarded from the original user/channel
	//
	// Use SetSavedFromMsgID and GetSavedFromMsgID helpers.
	SavedFromMsgID int
	// PSA type
	//
	// Use SetPsaType and GetPsaType helpers.
	PsaType string
}

// MessageFwdHeaderTypeID is TL type id of MessageFwdHeader.
const MessageFwdHeaderTypeID = 0x5f777dce

// Ensuring interfaces in compile-time for MessageFwdHeader.
var (
	_ bin.Encoder     = &MessageFwdHeader{}
	_ bin.Decoder     = &MessageFwdHeader{}
	_ bin.BareEncoder = &MessageFwdHeader{}
	_ bin.BareDecoder = &MessageFwdHeader{}
)

func (m *MessageFwdHeader) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Imported == false) {
		return false
	}
	if !(m.FromID == nil) {
		return false
	}
	if !(m.FromName == "") {
		return false
	}
	if !(m.Date == 0) {
		return false
	}
	if !(m.ChannelPost == 0) {
		return false
	}
	if !(m.PostAuthor == "") {
		return false
	}
	if !(m.SavedFromPeer == nil) {
		return false
	}
	if !(m.SavedFromMsgID == 0) {
		return false
	}
	if !(m.PsaType == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageFwdHeader) String() string {
	if m == nil {
		return "MessageFwdHeader(nil)"
	}
	type Alias MessageFwdHeader
	return fmt.Sprintf("MessageFwdHeader%+v", Alias(*m))
}

// FillFrom fills MessageFwdHeader from given interface.
func (m *MessageFwdHeader) FillFrom(from interface {
	GetImported() (value bool)
	GetFromID() (value PeerClass, ok bool)
	GetFromName() (value string, ok bool)
	GetDate() (value int)
	GetChannelPost() (value int, ok bool)
	GetPostAuthor() (value string, ok bool)
	GetSavedFromPeer() (value PeerClass, ok bool)
	GetSavedFromMsgID() (value int, ok bool)
	GetPsaType() (value string, ok bool)
}) {
	m.Imported = from.GetImported()
	if val, ok := from.GetFromID(); ok {
		m.FromID = val
	}

	if val, ok := from.GetFromName(); ok {
		m.FromName = val
	}

	m.Date = from.GetDate()
	if val, ok := from.GetChannelPost(); ok {
		m.ChannelPost = val
	}

	if val, ok := from.GetPostAuthor(); ok {
		m.PostAuthor = val
	}

	if val, ok := from.GetSavedFromPeer(); ok {
		m.SavedFromPeer = val
	}

	if val, ok := from.GetSavedFromMsgID(); ok {
		m.SavedFromMsgID = val
	}

	if val, ok := from.GetPsaType(); ok {
		m.PsaType = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageFwdHeader) TypeID() uint32 {
	return MessageFwdHeaderTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageFwdHeader) TypeName() string {
	return "messageFwdHeader"
}

// TypeInfo returns info about TL type.
func (m *MessageFwdHeader) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageFwdHeader",
		ID:   MessageFwdHeaderTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Imported",
			SchemaName: "imported",
			Null:       !m.Flags.Has(7),
		},
		{
			Name:       "FromID",
			SchemaName: "from_id",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "FromName",
			SchemaName: "from_name",
			Null:       !m.Flags.Has(5),
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "ChannelPost",
			SchemaName: "channel_post",
			Null:       !m.Flags.Has(2),
		},
		{
			Name:       "PostAuthor",
			SchemaName: "post_author",
			Null:       !m.Flags.Has(3),
		},
		{
			Name:       "SavedFromPeer",
			SchemaName: "saved_from_peer",
			Null:       !m.Flags.Has(4),
		},
		{
			Name:       "SavedFromMsgID",
			SchemaName: "saved_from_msg_id",
			Null:       !m.Flags.Has(4),
		},
		{
			Name:       "PsaType",
			SchemaName: "psa_type",
			Null:       !m.Flags.Has(6),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (m *MessageFwdHeader) SetFlags() {
	if !(m.Imported == false) {
		m.Flags.Set(7)
	}
	if !(m.FromID == nil) {
		m.Flags.Set(0)
	}
	if !(m.FromName == "") {
		m.Flags.Set(5)
	}
	if !(m.ChannelPost == 0) {
		m.Flags.Set(2)
	}
	if !(m.PostAuthor == "") {
		m.Flags.Set(3)
	}
	if !(m.SavedFromPeer == nil) {
		m.Flags.Set(4)
	}
	if !(m.SavedFromMsgID == 0) {
		m.Flags.Set(4)
	}
	if !(m.PsaType == "") {
		m.Flags.Set(6)
	}
}

// Encode implements bin.Encoder.
func (m *MessageFwdHeader) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageFwdHeader#5f777dce as nil")
	}
	b.PutID(MessageFwdHeaderTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageFwdHeader) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageFwdHeader#5f777dce as nil")
	}
	m.SetFlags()
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageFwdHeader#5f777dce: field flags: %w", err)
	}
	if m.Flags.Has(0) {
		if m.FromID == nil {
			return fmt.Errorf("unable to encode messageFwdHeader#5f777dce: field from_id is nil")
		}
		if err := m.FromID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageFwdHeader#5f777dce: field from_id: %w", err)
		}
	}
	if m.Flags.Has(5) {
		b.PutString(m.FromName)
	}
	b.PutInt(m.Date)
	if m.Flags.Has(2) {
		b.PutInt(m.ChannelPost)
	}
	if m.Flags.Has(3) {
		b.PutString(m.PostAuthor)
	}
	if m.Flags.Has(4) {
		if m.SavedFromPeer == nil {
			return fmt.Errorf("unable to encode messageFwdHeader#5f777dce: field saved_from_peer is nil")
		}
		if err := m.SavedFromPeer.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageFwdHeader#5f777dce: field saved_from_peer: %w", err)
		}
	}
	if m.Flags.Has(4) {
		b.PutInt(m.SavedFromMsgID)
	}
	if m.Flags.Has(6) {
		b.PutString(m.PsaType)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageFwdHeader) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageFwdHeader#5f777dce to nil")
	}
	if err := b.ConsumeID(MessageFwdHeaderTypeID); err != nil {
		return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageFwdHeader) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageFwdHeader#5f777dce to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field flags: %w", err)
		}
	}
	m.Imported = m.Flags.Has(7)
	if m.Flags.Has(0) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field from_id: %w", err)
		}
		m.FromID = value
	}
	if m.Flags.Has(5) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field from_name: %w", err)
		}
		m.FromName = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field date: %w", err)
		}
		m.Date = value
	}
	if m.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field channel_post: %w", err)
		}
		m.ChannelPost = value
	}
	if m.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field post_author: %w", err)
		}
		m.PostAuthor = value
	}
	if m.Flags.Has(4) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field saved_from_peer: %w", err)
		}
		m.SavedFromPeer = value
	}
	if m.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field saved_from_msg_id: %w", err)
		}
		m.SavedFromMsgID = value
	}
	if m.Flags.Has(6) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field psa_type: %w", err)
		}
		m.PsaType = value
	}
	return nil
}

// SetImported sets value of Imported conditional field.
func (m *MessageFwdHeader) SetImported(value bool) {
	if value {
		m.Flags.Set(7)
		m.Imported = true
	} else {
		m.Flags.Unset(7)
		m.Imported = false
	}
}

// GetImported returns value of Imported conditional field.
func (m *MessageFwdHeader) GetImported() (value bool) {
	if m == nil {
		return
	}
	return m.Flags.Has(7)
}

// SetFromID sets value of FromID conditional field.
func (m *MessageFwdHeader) SetFromID(value PeerClass) {
	m.Flags.Set(0)
	m.FromID = value
}

// GetFromID returns value of FromID conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetFromID() (value PeerClass, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.FromID, true
}

// SetFromName sets value of FromName conditional field.
func (m *MessageFwdHeader) SetFromName(value string) {
	m.Flags.Set(5)
	m.FromName = value
}

// GetFromName returns value of FromName conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetFromName() (value string, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(5) {
		return value, false
	}
	return m.FromName, true
}

// GetDate returns value of Date field.
func (m *MessageFwdHeader) GetDate() (value int) {
	if m == nil {
		return
	}
	return m.Date
}

// SetChannelPost sets value of ChannelPost conditional field.
func (m *MessageFwdHeader) SetChannelPost(value int) {
	m.Flags.Set(2)
	m.ChannelPost = value
}

// GetChannelPost returns value of ChannelPost conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetChannelPost() (value int, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(2) {
		return value, false
	}
	return m.ChannelPost, true
}

// SetPostAuthor sets value of PostAuthor conditional field.
func (m *MessageFwdHeader) SetPostAuthor(value string) {
	m.Flags.Set(3)
	m.PostAuthor = value
}

// GetPostAuthor returns value of PostAuthor conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetPostAuthor() (value string, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(3) {
		return value, false
	}
	return m.PostAuthor, true
}

// SetSavedFromPeer sets value of SavedFromPeer conditional field.
func (m *MessageFwdHeader) SetSavedFromPeer(value PeerClass) {
	m.Flags.Set(4)
	m.SavedFromPeer = value
}

// GetSavedFromPeer returns value of SavedFromPeer conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetSavedFromPeer() (value PeerClass, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(4) {
		return value, false
	}
	return m.SavedFromPeer, true
}

// SetSavedFromMsgID sets value of SavedFromMsgID conditional field.
func (m *MessageFwdHeader) SetSavedFromMsgID(value int) {
	m.Flags.Set(4)
	m.SavedFromMsgID = value
}

// GetSavedFromMsgID returns value of SavedFromMsgID conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetSavedFromMsgID() (value int, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(4) {
		return value, false
	}
	return m.SavedFromMsgID, true
}

// SetPsaType sets value of PsaType conditional field.
func (m *MessageFwdHeader) SetPsaType(value string) {
	m.Flags.Set(6)
	m.PsaType = value
}

// GetPsaType returns value of PsaType conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetPsaType() (value string, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(6) {
		return value, false
	}
	return m.PsaType, true
}
