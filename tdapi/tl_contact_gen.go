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

// Contact represents TL type `contact#89285774`.
type Contact struct {
	// Phone number of the user
	PhoneNumber string
	// First name of the user; 1-255 characters in length
	FirstName string
	// Last name of the user
	LastName string
	// Additional data about the user in a form of vCard; 0-2048 bytes in length
	Vcard string
	// Identifier of the user, if known; otherwise 0
	UserID int64
}

// ContactTypeID is TL type id of Contact.
const ContactTypeID = 0x89285774

// Ensuring interfaces in compile-time for Contact.
var (
	_ bin.Encoder     = &Contact{}
	_ bin.Decoder     = &Contact{}
	_ bin.BareEncoder = &Contact{}
	_ bin.BareDecoder = &Contact{}
)

func (c *Contact) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.PhoneNumber == "") {
		return false
	}
	if !(c.FirstName == "") {
		return false
	}
	if !(c.LastName == "") {
		return false
	}
	if !(c.Vcard == "") {
		return false
	}
	if !(c.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *Contact) String() string {
	if c == nil {
		return "Contact(nil)"
	}
	type Alias Contact
	return fmt.Sprintf("Contact%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Contact) TypeID() uint32 {
	return ContactTypeID
}

// TypeName returns name of type in TL schema.
func (*Contact) TypeName() string {
	return "contact"
}

// TypeInfo returns info about TL type.
func (c *Contact) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contact",
		ID:   ContactTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "FirstName",
			SchemaName: "first_name",
		},
		{
			Name:       "LastName",
			SchemaName: "last_name",
		},
		{
			Name:       "Vcard",
			SchemaName: "vcard",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *Contact) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contact#89285774 as nil")
	}
	b.PutID(ContactTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *Contact) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contact#89285774 as nil")
	}
	b.PutString(c.PhoneNumber)
	b.PutString(c.FirstName)
	b.PutString(c.LastName)
	b.PutString(c.Vcard)
	b.PutInt53(c.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (c *Contact) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contact#89285774 to nil")
	}
	if err := b.ConsumeID(ContactTypeID); err != nil {
		return fmt.Errorf("unable to decode contact#89285774: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *Contact) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contact#89285774 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contact#89285774: field phone_number: %w", err)
		}
		c.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contact#89285774: field first_name: %w", err)
		}
		c.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contact#89285774: field last_name: %w", err)
		}
		c.LastName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contact#89285774: field vcard: %w", err)
		}
		c.Vcard = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode contact#89285774: field user_id: %w", err)
		}
		c.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *Contact) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode contact#89285774 as nil")
	}
	b.ObjStart()
	b.PutID("contact")
	b.Comma()
	b.FieldStart("phone_number")
	b.PutString(c.PhoneNumber)
	b.Comma()
	b.FieldStart("first_name")
	b.PutString(c.FirstName)
	b.Comma()
	b.FieldStart("last_name")
	b.PutString(c.LastName)
	b.Comma()
	b.FieldStart("vcard")
	b.PutString(c.Vcard)
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(c.UserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *Contact) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode contact#89285774 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("contact"); err != nil {
				return fmt.Errorf("unable to decode contact#89285774: %w", err)
			}
		case "phone_number":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode contact#89285774: field phone_number: %w", err)
			}
			c.PhoneNumber = value
		case "first_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode contact#89285774: field first_name: %w", err)
			}
			c.FirstName = value
		case "last_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode contact#89285774: field last_name: %w", err)
			}
			c.LastName = value
		case "vcard":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode contact#89285774: field vcard: %w", err)
			}
			c.Vcard = value
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode contact#89285774: field user_id: %w", err)
			}
			c.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPhoneNumber returns value of PhoneNumber field.
func (c *Contact) GetPhoneNumber() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumber
}

// GetFirstName returns value of FirstName field.
func (c *Contact) GetFirstName() (value string) {
	if c == nil {
		return
	}
	return c.FirstName
}

// GetLastName returns value of LastName field.
func (c *Contact) GetLastName() (value string) {
	if c == nil {
		return
	}
	return c.LastName
}

// GetVcard returns value of Vcard field.
func (c *Contact) GetVcard() (value string) {
	if c == nil {
		return
	}
	return c.Vcard
}

// GetUserID returns value of UserID field.
func (c *Contact) GetUserID() (value int64) {
	if c == nil {
		return
	}
	return c.UserID
}
