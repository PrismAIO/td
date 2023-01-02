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

// InputPeerNotifySettings represents TL type `inputPeerNotifySettings#df1f002b`.
// Notification settings.
//
// See https://core.telegram.org/constructor/inputPeerNotifySettings for reference.
type InputPeerNotifySettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If the text of the message shall be displayed in notification
	//
	// Use SetShowPreviews and GetShowPreviews helpers.
	ShowPreviews bool
	// Peer was muted?
	//
	// Use SetSilent and GetSilent helpers.
	Silent bool
	// Date until which all notifications shall be switched off
	//
	// Use SetMuteUntil and GetMuteUntil helpers.
	MuteUntil int
	// Name of an audio file for notification
	//
	// Use SetSound and GetSound helpers.
	Sound NotificationSoundClass
}

// InputPeerNotifySettingsTypeID is TL type id of InputPeerNotifySettings.
const InputPeerNotifySettingsTypeID = 0xdf1f002b

// Ensuring interfaces in compile-time for InputPeerNotifySettings.
var (
	_ bin.Encoder     = &InputPeerNotifySettings{}
	_ bin.Decoder     = &InputPeerNotifySettings{}
	_ bin.BareEncoder = &InputPeerNotifySettings{}
	_ bin.BareDecoder = &InputPeerNotifySettings{}
)

func (i *InputPeerNotifySettings) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.ShowPreviews == false) {
		return false
	}
	if !(i.Silent == false) {
		return false
	}
	if !(i.MuteUntil == 0) {
		return false
	}
	if !(i.Sound == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerNotifySettings) String() string {
	if i == nil {
		return "InputPeerNotifySettings(nil)"
	}
	type Alias InputPeerNotifySettings
	return fmt.Sprintf("InputPeerNotifySettings%+v", Alias(*i))
}

// FillFrom fills InputPeerNotifySettings from given interface.
func (i *InputPeerNotifySettings) FillFrom(from interface {
	GetShowPreviews() (value bool, ok bool)
	GetSilent() (value bool, ok bool)
	GetMuteUntil() (value int, ok bool)
	GetSound() (value NotificationSoundClass, ok bool)
}) {
	if val, ok := from.GetShowPreviews(); ok {
		i.ShowPreviews = val
	}

	if val, ok := from.GetSilent(); ok {
		i.Silent = val
	}

	if val, ok := from.GetMuteUntil(); ok {
		i.MuteUntil = val
	}

	if val, ok := from.GetSound(); ok {
		i.Sound = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPeerNotifySettings) TypeID() uint32 {
	return InputPeerNotifySettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPeerNotifySettings) TypeName() string {
	return "inputPeerNotifySettings"
}

// TypeInfo returns info about TL type.
func (i *InputPeerNotifySettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPeerNotifySettings",
		ID:   InputPeerNotifySettingsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShowPreviews",
			SchemaName: "show_previews",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !i.Flags.Has(1),
		},
		{
			Name:       "MuteUntil",
			SchemaName: "mute_until",
			Null:       !i.Flags.Has(2),
		},
		{
			Name:       "Sound",
			SchemaName: "sound",
			Null:       !i.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *InputPeerNotifySettings) SetFlags() {
	if !(i.ShowPreviews == false) {
		i.Flags.Set(0)
	}
	if !(i.Silent == false) {
		i.Flags.Set(1)
	}
	if !(i.MuteUntil == 0) {
		i.Flags.Set(2)
	}
	if !(i.Sound == nil) {
		i.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (i *InputPeerNotifySettings) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerNotifySettings#df1f002b as nil")
	}
	b.PutID(InputPeerNotifySettingsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPeerNotifySettings) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerNotifySettings#df1f002b as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPeerNotifySettings#df1f002b: field flags: %w", err)
	}
	if i.Flags.Has(0) {
		b.PutBool(i.ShowPreviews)
	}
	if i.Flags.Has(1) {
		b.PutBool(i.Silent)
	}
	if i.Flags.Has(2) {
		b.PutInt(i.MuteUntil)
	}
	if i.Flags.Has(3) {
		if i.Sound == nil {
			return fmt.Errorf("unable to encode inputPeerNotifySettings#df1f002b: field sound is nil")
		}
		if err := i.Sound.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputPeerNotifySettings#df1f002b: field sound: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPeerNotifySettings) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerNotifySettings#df1f002b to nil")
	}
	if err := b.ConsumeID(InputPeerNotifySettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerNotifySettings#df1f002b: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPeerNotifySettings) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerNotifySettings#df1f002b to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPeerNotifySettings#df1f002b: field flags: %w", err)
		}
	}
	if i.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerNotifySettings#df1f002b: field show_previews: %w", err)
		}
		i.ShowPreviews = value
	}
	if i.Flags.Has(1) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerNotifySettings#df1f002b: field silent: %w", err)
		}
		i.Silent = value
	}
	if i.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerNotifySettings#df1f002b: field mute_until: %w", err)
		}
		i.MuteUntil = value
	}
	if i.Flags.Has(3) {
		value, err := DecodeNotificationSound(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerNotifySettings#df1f002b: field sound: %w", err)
		}
		i.Sound = value
	}
	return nil
}

// SetShowPreviews sets value of ShowPreviews conditional field.
func (i *InputPeerNotifySettings) SetShowPreviews(value bool) {
	i.Flags.Set(0)
	i.ShowPreviews = value
}

// GetShowPreviews returns value of ShowPreviews conditional field and
// boolean which is true if field was set.
func (i *InputPeerNotifySettings) GetShowPreviews() (value bool, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.ShowPreviews, true
}

// SetSilent sets value of Silent conditional field.
func (i *InputPeerNotifySettings) SetSilent(value bool) {
	i.Flags.Set(1)
	i.Silent = value
}

// GetSilent returns value of Silent conditional field and
// boolean which is true if field was set.
func (i *InputPeerNotifySettings) GetSilent() (value bool, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Silent, true
}

// SetMuteUntil sets value of MuteUntil conditional field.
func (i *InputPeerNotifySettings) SetMuteUntil(value int) {
	i.Flags.Set(2)
	i.MuteUntil = value
}

// GetMuteUntil returns value of MuteUntil conditional field and
// boolean which is true if field was set.
func (i *InputPeerNotifySettings) GetMuteUntil() (value int, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.MuteUntil, true
}

// SetSound sets value of Sound conditional field.
func (i *InputPeerNotifySettings) SetSound(value NotificationSoundClass) {
	i.Flags.Set(3)
	i.Sound = value
}

// GetSound returns value of Sound conditional field and
// boolean which is true if field was set.
func (i *InputPeerNotifySettings) GetSound() (value NotificationSoundClass, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(3) {
		return value, false
	}
	return i.Sound, true
}
