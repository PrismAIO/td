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

// Username represents TL type `username#b4073647`.
//
// See https://core.telegram.org/constructor/username for reference.
type Username struct {
	// Flags field of Username.
	Flags bin.Fields
	// Editable field of Username.
	Editable bool
	// Active field of Username.
	Active bool
	// Username field of Username.
	Username string
}

// UsernameTypeID is TL type id of Username.
const UsernameTypeID = 0xb4073647

// Ensuring interfaces in compile-time for Username.
var (
	_ bin.Encoder     = &Username{}
	_ bin.Decoder     = &Username{}
	_ bin.BareEncoder = &Username{}
	_ bin.BareDecoder = &Username{}
)

func (u *Username) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Editable == false) {
		return false
	}
	if !(u.Active == false) {
		return false
	}
	if !(u.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *Username) String() string {
	if u == nil {
		return "Username(nil)"
	}
	type Alias Username
	return fmt.Sprintf("Username%+v", Alias(*u))
}

// FillFrom fills Username from given interface.
func (u *Username) FillFrom(from interface {
	GetEditable() (value bool)
	GetActive() (value bool)
	GetUsername() (value string)
}) {
	u.Editable = from.GetEditable()
	u.Active = from.GetActive()
	u.Username = from.GetUsername()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Username) TypeID() uint32 {
	return UsernameTypeID
}

// TypeName returns name of type in TL schema.
func (*Username) TypeName() string {
	return "username"
}

// TypeInfo returns info about TL type.
func (u *Username) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "username",
		ID:   UsernameTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Editable",
			SchemaName: "editable",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "Active",
			SchemaName: "active",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "Username",
			SchemaName: "username",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *Username) SetFlags() {
	if !(u.Editable == false) {
		u.Flags.Set(0)
	}
	if !(u.Active == false) {
		u.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (u *Username) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode username#b4073647 as nil")
	}
	b.PutID(UsernameTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *Username) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode username#b4073647 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode username#b4073647: field flags: %w", err)
	}
	b.PutString(u.Username)
	return nil
}

// Decode implements bin.Decoder.
func (u *Username) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode username#b4073647 to nil")
	}
	if err := b.ConsumeID(UsernameTypeID); err != nil {
		return fmt.Errorf("unable to decode username#b4073647: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *Username) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode username#b4073647 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode username#b4073647: field flags: %w", err)
		}
	}
	u.Editable = u.Flags.Has(0)
	u.Active = u.Flags.Has(1)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode username#b4073647: field username: %w", err)
		}
		u.Username = value
	}
	return nil
}

// SetEditable sets value of Editable conditional field.
func (u *Username) SetEditable(value bool) {
	if value {
		u.Flags.Set(0)
		u.Editable = true
	} else {
		u.Flags.Unset(0)
		u.Editable = false
	}
}

// GetEditable returns value of Editable conditional field.
func (u *Username) GetEditable() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// SetActive sets value of Active conditional field.
func (u *Username) SetActive(value bool) {
	if value {
		u.Flags.Set(1)
		u.Active = true
	} else {
		u.Flags.Unset(1)
		u.Active = false
	}
}

// GetActive returns value of Active conditional field.
func (u *Username) GetActive() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(1)
}

// GetUsername returns value of Username field.
func (u *Username) GetUsername() (value string) {
	if u == nil {
		return
	}
	return u.Username
}
