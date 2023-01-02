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

// MessageReplies represents TL type `messageReplies#83d60fc2`.
// Info about the comment section of a channel post, or a simple message thread¹
//
// Links:
//  1. https://core.telegram.org/api/threads
//
// See https://core.telegram.org/constructor/messageReplies for reference.
type MessageReplies struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this constructor contains information about the comment section of a channel
	// post, or a simple message thread¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	Comments bool
	// Contains the total number of replies in this thread or comment section.
	Replies int
	// PTS¹ of the message that started this thread.
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	RepliesPts int
	// For channel post comments, contains information about the last few comment posters for
	// a specific thread, to show a small list of commenter profile pictures in client
	// previews.
	//
	// Use SetRecentRepliers and GetRecentRepliers helpers.
	RecentRepliers []PeerClass
	// For channel post comments, contains the ID of the associated discussion supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/discussion
	//
	// Use SetChannelID and GetChannelID helpers.
	ChannelID int64
	// ID of the latest message in this thread or comment section.
	//
	// Use SetMaxID and GetMaxID helpers.
	MaxID int
	// Contains the ID of the latest read message in this thread or comment section.
	//
	// Use SetReadMaxID and GetReadMaxID helpers.
	ReadMaxID int
}

// MessageRepliesTypeID is TL type id of MessageReplies.
const MessageRepliesTypeID = 0x83d60fc2

// Ensuring interfaces in compile-time for MessageReplies.
var (
	_ bin.Encoder     = &MessageReplies{}
	_ bin.Decoder     = &MessageReplies{}
	_ bin.BareEncoder = &MessageReplies{}
	_ bin.BareDecoder = &MessageReplies{}
)

func (m *MessageReplies) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Comments == false) {
		return false
	}
	if !(m.Replies == 0) {
		return false
	}
	if !(m.RepliesPts == 0) {
		return false
	}
	if !(m.RecentRepliers == nil) {
		return false
	}
	if !(m.ChannelID == 0) {
		return false
	}
	if !(m.MaxID == 0) {
		return false
	}
	if !(m.ReadMaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageReplies) String() string {
	if m == nil {
		return "MessageReplies(nil)"
	}
	type Alias MessageReplies
	return fmt.Sprintf("MessageReplies%+v", Alias(*m))
}

// FillFrom fills MessageReplies from given interface.
func (m *MessageReplies) FillFrom(from interface {
	GetComments() (value bool)
	GetReplies() (value int)
	GetRepliesPts() (value int)
	GetRecentRepliers() (value []PeerClass, ok bool)
	GetChannelID() (value int64, ok bool)
	GetMaxID() (value int, ok bool)
	GetReadMaxID() (value int, ok bool)
}) {
	m.Comments = from.GetComments()
	m.Replies = from.GetReplies()
	m.RepliesPts = from.GetRepliesPts()
	if val, ok := from.GetRecentRepliers(); ok {
		m.RecentRepliers = val
	}

	if val, ok := from.GetChannelID(); ok {
		m.ChannelID = val
	}

	if val, ok := from.GetMaxID(); ok {
		m.MaxID = val
	}

	if val, ok := from.GetReadMaxID(); ok {
		m.ReadMaxID = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageReplies) TypeID() uint32 {
	return MessageRepliesTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageReplies) TypeName() string {
	return "messageReplies"
}

// TypeInfo returns info about TL type.
func (m *MessageReplies) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageReplies",
		ID:   MessageRepliesTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Comments",
			SchemaName: "comments",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "Replies",
			SchemaName: "replies",
		},
		{
			Name:       "RepliesPts",
			SchemaName: "replies_pts",
		},
		{
			Name:       "RecentRepliers",
			SchemaName: "recent_repliers",
			Null:       !m.Flags.Has(1),
		},
		{
			Name:       "ChannelID",
			SchemaName: "channel_id",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
			Null:       !m.Flags.Has(2),
		},
		{
			Name:       "ReadMaxID",
			SchemaName: "read_max_id",
			Null:       !m.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (m *MessageReplies) SetFlags() {
	if !(m.Comments == false) {
		m.Flags.Set(0)
	}
	if !(m.RecentRepliers == nil) {
		m.Flags.Set(1)
	}
	if !(m.ChannelID == 0) {
		m.Flags.Set(0)
	}
	if !(m.MaxID == 0) {
		m.Flags.Set(2)
	}
	if !(m.ReadMaxID == 0) {
		m.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (m *MessageReplies) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplies#83d60fc2 as nil")
	}
	b.PutID(MessageRepliesTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageReplies) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplies#83d60fc2 as nil")
	}
	m.SetFlags()
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageReplies#83d60fc2: field flags: %w", err)
	}
	b.PutInt(m.Replies)
	b.PutInt(m.RepliesPts)
	if m.Flags.Has(1) {
		b.PutVectorHeader(len(m.RecentRepliers))
		for idx, v := range m.RecentRepliers {
			if v == nil {
				return fmt.Errorf("unable to encode messageReplies#83d60fc2: field recent_repliers element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messageReplies#83d60fc2: field recent_repliers element with index %d: %w", idx, err)
			}
		}
	}
	if m.Flags.Has(0) {
		b.PutLong(m.ChannelID)
	}
	if m.Flags.Has(2) {
		b.PutInt(m.MaxID)
	}
	if m.Flags.Has(3) {
		b.PutInt(m.ReadMaxID)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageReplies) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplies#83d60fc2 to nil")
	}
	if err := b.ConsumeID(MessageRepliesTypeID); err != nil {
		return fmt.Errorf("unable to decode messageReplies#83d60fc2: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageReplies) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplies#83d60fc2 to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageReplies#83d60fc2: field flags: %w", err)
		}
	}
	m.Comments = m.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#83d60fc2: field replies: %w", err)
		}
		m.Replies = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#83d60fc2: field replies_pts: %w", err)
		}
		m.RepliesPts = value
	}
	if m.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#83d60fc2: field recent_repliers: %w", err)
		}

		if headerLen > 0 {
			m.RecentRepliers = make([]PeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode messageReplies#83d60fc2: field recent_repliers: %w", err)
			}
			m.RecentRepliers = append(m.RecentRepliers, value)
		}
	}
	if m.Flags.Has(0) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#83d60fc2: field channel_id: %w", err)
		}
		m.ChannelID = value
	}
	if m.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#83d60fc2: field max_id: %w", err)
		}
		m.MaxID = value
	}
	if m.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#83d60fc2: field read_max_id: %w", err)
		}
		m.ReadMaxID = value
	}
	return nil
}

// SetComments sets value of Comments conditional field.
func (m *MessageReplies) SetComments(value bool) {
	if value {
		m.Flags.Set(0)
		m.Comments = true
	} else {
		m.Flags.Unset(0)
		m.Comments = false
	}
}

// GetComments returns value of Comments conditional field.
func (m *MessageReplies) GetComments() (value bool) {
	if m == nil {
		return
	}
	return m.Flags.Has(0)
}

// GetReplies returns value of Replies field.
func (m *MessageReplies) GetReplies() (value int) {
	if m == nil {
		return
	}
	return m.Replies
}

// GetRepliesPts returns value of RepliesPts field.
func (m *MessageReplies) GetRepliesPts() (value int) {
	if m == nil {
		return
	}
	return m.RepliesPts
}

// SetRecentRepliers sets value of RecentRepliers conditional field.
func (m *MessageReplies) SetRecentRepliers(value []PeerClass) {
	m.Flags.Set(1)
	m.RecentRepliers = value
}

// GetRecentRepliers returns value of RecentRepliers conditional field and
// boolean which is true if field was set.
func (m *MessageReplies) GetRecentRepliers() (value []PeerClass, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(1) {
		return value, false
	}
	return m.RecentRepliers, true
}

// SetChannelID sets value of ChannelID conditional field.
func (m *MessageReplies) SetChannelID(value int64) {
	m.Flags.Set(0)
	m.ChannelID = value
}

// GetChannelID returns value of ChannelID conditional field and
// boolean which is true if field was set.
func (m *MessageReplies) GetChannelID() (value int64, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.ChannelID, true
}

// SetMaxID sets value of MaxID conditional field.
func (m *MessageReplies) SetMaxID(value int) {
	m.Flags.Set(2)
	m.MaxID = value
}

// GetMaxID returns value of MaxID conditional field and
// boolean which is true if field was set.
func (m *MessageReplies) GetMaxID() (value int, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(2) {
		return value, false
	}
	return m.MaxID, true
}

// SetReadMaxID sets value of ReadMaxID conditional field.
func (m *MessageReplies) SetReadMaxID(value int) {
	m.Flags.Set(3)
	m.ReadMaxID = value
}

// GetReadMaxID returns value of ReadMaxID conditional field and
// boolean which is true if field was set.
func (m *MessageReplies) GetReadMaxID() (value int, ok bool) {
	if m == nil {
		return
	}
	if !m.Flags.Has(3) {
		return value, false
	}
	return m.ReadMaxID, true
}

// MapRecentRepliers returns field RecentRepliers wrapped in PeerClassArray helper.
func (m *MessageReplies) MapRecentRepliers() (value PeerClassArray, ok bool) {
	if !m.Flags.Has(1) {
		return value, false
	}
	return PeerClassArray(m.RecentRepliers), true
}
