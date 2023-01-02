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

// NotificationSound represents TL type `notificationSound#f4ef6137`.
type NotificationSound struct {
	// Unique identifier of the notification sound
	ID int64
	// Duration of the sound, in seconds
	Duration int32
	// Point in time (Unix timestamp) when the sound was created
	Date int32
	// Title of the notification sound
	Title string
	// Arbitrary data, defined while the sound was uploaded
	Data string
	// File containing the sound
	Sound File
}

// NotificationSoundTypeID is TL type id of NotificationSound.
const NotificationSoundTypeID = 0xf4ef6137

// Ensuring interfaces in compile-time for NotificationSound.
var (
	_ bin.Encoder     = &NotificationSound{}
	_ bin.Decoder     = &NotificationSound{}
	_ bin.BareEncoder = &NotificationSound{}
	_ bin.BareDecoder = &NotificationSound{}
)

func (n *NotificationSound) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.ID == 0) {
		return false
	}
	if !(n.Duration == 0) {
		return false
	}
	if !(n.Date == 0) {
		return false
	}
	if !(n.Title == "") {
		return false
	}
	if !(n.Data == "") {
		return false
	}
	if !(n.Sound.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotificationSound) String() string {
	if n == nil {
		return "NotificationSound(nil)"
	}
	type Alias NotificationSound
	return fmt.Sprintf("NotificationSound%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NotificationSound) TypeID() uint32 {
	return NotificationSoundTypeID
}

// TypeName returns name of type in TL schema.
func (*NotificationSound) TypeName() string {
	return "notificationSound"
}

// TypeInfo returns info about TL type.
func (n *NotificationSound) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "notificationSound",
		ID:   NotificationSoundTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Data",
			SchemaName: "data",
		},
		{
			Name:       "Sound",
			SchemaName: "sound",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NotificationSound) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationSound#f4ef6137 as nil")
	}
	b.PutID(NotificationSoundTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NotificationSound) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationSound#f4ef6137 as nil")
	}
	b.PutLong(n.ID)
	b.PutInt32(n.Duration)
	b.PutInt32(n.Date)
	b.PutString(n.Title)
	b.PutString(n.Data)
	if err := n.Sound.Encode(b); err != nil {
		return fmt.Errorf("unable to encode notificationSound#f4ef6137: field sound: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NotificationSound) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationSound#f4ef6137 to nil")
	}
	if err := b.ConsumeID(NotificationSoundTypeID); err != nil {
		return fmt.Errorf("unable to decode notificationSound#f4ef6137: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NotificationSound) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationSound#f4ef6137 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode notificationSound#f4ef6137: field id: %w", err)
		}
		n.ID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode notificationSound#f4ef6137: field duration: %w", err)
		}
		n.Duration = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode notificationSound#f4ef6137: field date: %w", err)
		}
		n.Date = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode notificationSound#f4ef6137: field title: %w", err)
		}
		n.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode notificationSound#f4ef6137: field data: %w", err)
		}
		n.Data = value
	}
	{
		if err := n.Sound.Decode(b); err != nil {
			return fmt.Errorf("unable to decode notificationSound#f4ef6137: field sound: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NotificationSound) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationSound#f4ef6137 as nil")
	}
	b.ObjStart()
	b.PutID("notificationSound")
	b.Comma()
	b.FieldStart("id")
	b.PutLong(n.ID)
	b.Comma()
	b.FieldStart("duration")
	b.PutInt32(n.Duration)
	b.Comma()
	b.FieldStart("date")
	b.PutInt32(n.Date)
	b.Comma()
	b.FieldStart("title")
	b.PutString(n.Title)
	b.Comma()
	b.FieldStart("data")
	b.PutString(n.Data)
	b.Comma()
	b.FieldStart("sound")
	if err := n.Sound.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode notificationSound#f4ef6137: field sound: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NotificationSound) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationSound#f4ef6137 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("notificationSound"); err != nil {
				return fmt.Errorf("unable to decode notificationSound#f4ef6137: %w", err)
			}
		case "id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode notificationSound#f4ef6137: field id: %w", err)
			}
			n.ID = value
		case "duration":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode notificationSound#f4ef6137: field duration: %w", err)
			}
			n.Duration = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode notificationSound#f4ef6137: field date: %w", err)
			}
			n.Date = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode notificationSound#f4ef6137: field title: %w", err)
			}
			n.Title = value
		case "data":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode notificationSound#f4ef6137: field data: %w", err)
			}
			n.Data = value
		case "sound":
			if err := n.Sound.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode notificationSound#f4ef6137: field sound: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (n *NotificationSound) GetID() (value int64) {
	if n == nil {
		return
	}
	return n.ID
}

// GetDuration returns value of Duration field.
func (n *NotificationSound) GetDuration() (value int32) {
	if n == nil {
		return
	}
	return n.Duration
}

// GetDate returns value of Date field.
func (n *NotificationSound) GetDate() (value int32) {
	if n == nil {
		return
	}
	return n.Date
}

// GetTitle returns value of Title field.
func (n *NotificationSound) GetTitle() (value string) {
	if n == nil {
		return
	}
	return n.Title
}

// GetData returns value of Data field.
func (n *NotificationSound) GetData() (value string) {
	if n == nil {
		return
	}
	return n.Data
}

// GetSound returns value of Sound field.
func (n *NotificationSound) GetSound() (value File) {
	if n == nil {
		return
	}
	return n.Sound
}
