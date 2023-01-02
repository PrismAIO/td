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

// HelpSupport represents TL type `help.support#17c6b5f6`.
// Info on support user.
//
// See https://core.telegram.org/constructor/help.support for reference.
type HelpSupport struct {
	// Phone number
	PhoneNumber string
	// User
	User UserClass
}

// HelpSupportTypeID is TL type id of HelpSupport.
const HelpSupportTypeID = 0x17c6b5f6

// Ensuring interfaces in compile-time for HelpSupport.
var (
	_ bin.Encoder     = &HelpSupport{}
	_ bin.Decoder     = &HelpSupport{}
	_ bin.BareEncoder = &HelpSupport{}
	_ bin.BareDecoder = &HelpSupport{}
)

func (s *HelpSupport) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PhoneNumber == "") {
		return false
	}
	if !(s.User == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *HelpSupport) String() string {
	if s == nil {
		return "HelpSupport(nil)"
	}
	type Alias HelpSupport
	return fmt.Sprintf("HelpSupport%+v", Alias(*s))
}

// FillFrom fills HelpSupport from given interface.
func (s *HelpSupport) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetUser() (value UserClass)
}) {
	s.PhoneNumber = from.GetPhoneNumber()
	s.User = from.GetUser()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpSupport) TypeID() uint32 {
	return HelpSupportTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpSupport) TypeName() string {
	return "help.support"
}

// TypeInfo returns info about TL type.
func (s *HelpSupport) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.support",
		ID:   HelpSupportTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "User",
			SchemaName: "user",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *HelpSupport) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode help.support#17c6b5f6 as nil")
	}
	b.PutID(HelpSupportTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *HelpSupport) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode help.support#17c6b5f6 as nil")
	}
	b.PutString(s.PhoneNumber)
	if s.User == nil {
		return fmt.Errorf("unable to encode help.support#17c6b5f6: field user is nil")
	}
	if err := s.User.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.support#17c6b5f6: field user: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *HelpSupport) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode help.support#17c6b5f6 to nil")
	}
	if err := b.ConsumeID(HelpSupportTypeID); err != nil {
		return fmt.Errorf("unable to decode help.support#17c6b5f6: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *HelpSupport) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode help.support#17c6b5f6 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.support#17c6b5f6: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		value, err := DecodeUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.support#17c6b5f6: field user: %w", err)
		}
		s.User = value
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (s *HelpSupport) GetPhoneNumber() (value string) {
	if s == nil {
		return
	}
	return s.PhoneNumber
}

// GetUser returns value of User field.
func (s *HelpSupport) GetUser() (value UserClass) {
	if s == nil {
		return
	}
	return s.User
}

// GetUserAsNotEmpty returns mapped value of User field.
func (s *HelpSupport) GetUserAsNotEmpty() (*User, bool) {
	return s.User.AsNotEmpty()
}
