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

// InputPrivacyValueAllowContacts represents TL type `inputPrivacyValueAllowContacts#d09e07b`.
// Allow only contacts
//
// See https://core.telegram.org/constructor/inputPrivacyValueAllowContacts for reference.
type InputPrivacyValueAllowContacts struct {
}

// InputPrivacyValueAllowContactsTypeID is TL type id of InputPrivacyValueAllowContacts.
const InputPrivacyValueAllowContactsTypeID = 0xd09e07b

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowContacts) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowContacts.
var (
	_ bin.Encoder     = &InputPrivacyValueAllowContacts{}
	_ bin.Decoder     = &InputPrivacyValueAllowContacts{}
	_ bin.BareEncoder = &InputPrivacyValueAllowContacts{}
	_ bin.BareDecoder = &InputPrivacyValueAllowContacts{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowContacts{}
)

func (i *InputPrivacyValueAllowContacts) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueAllowContacts) String() string {
	if i == nil {
		return "InputPrivacyValueAllowContacts(nil)"
	}
	type Alias InputPrivacyValueAllowContacts
	return fmt.Sprintf("InputPrivacyValueAllowContacts%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyValueAllowContacts) TypeID() uint32 {
	return InputPrivacyValueAllowContactsTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyValueAllowContacts) TypeName() string {
	return "inputPrivacyValueAllowContacts"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyValueAllowContacts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyValueAllowContacts",
		ID:   InputPrivacyValueAllowContactsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowContacts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowContacts#d09e07b as nil")
	}
	b.PutID(InputPrivacyValueAllowContactsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyValueAllowContacts) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowContacts#d09e07b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowContacts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowContacts#d09e07b to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowContacts#d09e07b: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyValueAllowContacts) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowContacts#d09e07b to nil")
	}
	return nil
}

// InputPrivacyValueAllowAll represents TL type `inputPrivacyValueAllowAll#184b35ce`.
// Allow all users
//
// See https://core.telegram.org/constructor/inputPrivacyValueAllowAll for reference.
type InputPrivacyValueAllowAll struct {
}

// InputPrivacyValueAllowAllTypeID is TL type id of InputPrivacyValueAllowAll.
const InputPrivacyValueAllowAllTypeID = 0x184b35ce

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowAll) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowAll.
var (
	_ bin.Encoder     = &InputPrivacyValueAllowAll{}
	_ bin.Decoder     = &InputPrivacyValueAllowAll{}
	_ bin.BareEncoder = &InputPrivacyValueAllowAll{}
	_ bin.BareDecoder = &InputPrivacyValueAllowAll{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowAll{}
)

func (i *InputPrivacyValueAllowAll) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueAllowAll) String() string {
	if i == nil {
		return "InputPrivacyValueAllowAll(nil)"
	}
	type Alias InputPrivacyValueAllowAll
	return fmt.Sprintf("InputPrivacyValueAllowAll%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyValueAllowAll) TypeID() uint32 {
	return InputPrivacyValueAllowAllTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyValueAllowAll) TypeName() string {
	return "inputPrivacyValueAllowAll"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyValueAllowAll) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyValueAllowAll",
		ID:   InputPrivacyValueAllowAllTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowAll) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowAll#184b35ce as nil")
	}
	b.PutID(InputPrivacyValueAllowAllTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyValueAllowAll) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowAll#184b35ce as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowAll) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowAll#184b35ce to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowAllTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowAll#184b35ce: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyValueAllowAll) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowAll#184b35ce to nil")
	}
	return nil
}

// InputPrivacyValueAllowUsers represents TL type `inputPrivacyValueAllowUsers#131cc67f`.
// Allow only certain users
//
// See https://core.telegram.org/constructor/inputPrivacyValueAllowUsers for reference.
type InputPrivacyValueAllowUsers struct {
	// Allowed users
	Users []InputUserClass
}

// InputPrivacyValueAllowUsersTypeID is TL type id of InputPrivacyValueAllowUsers.
const InputPrivacyValueAllowUsersTypeID = 0x131cc67f

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowUsers) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowUsers.
var (
	_ bin.Encoder     = &InputPrivacyValueAllowUsers{}
	_ bin.Decoder     = &InputPrivacyValueAllowUsers{}
	_ bin.BareEncoder = &InputPrivacyValueAllowUsers{}
	_ bin.BareDecoder = &InputPrivacyValueAllowUsers{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowUsers{}
)

func (i *InputPrivacyValueAllowUsers) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueAllowUsers) String() string {
	if i == nil {
		return "InputPrivacyValueAllowUsers(nil)"
	}
	type Alias InputPrivacyValueAllowUsers
	return fmt.Sprintf("InputPrivacyValueAllowUsers%+v", Alias(*i))
}

// FillFrom fills InputPrivacyValueAllowUsers from given interface.
func (i *InputPrivacyValueAllowUsers) FillFrom(from interface {
	GetUsers() (value []InputUserClass)
}) {
	i.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyValueAllowUsers) TypeID() uint32 {
	return InputPrivacyValueAllowUsersTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyValueAllowUsers) TypeName() string {
	return "inputPrivacyValueAllowUsers"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyValueAllowUsers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyValueAllowUsers",
		ID:   InputPrivacyValueAllowUsersTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowUsers) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowUsers#131cc67f as nil")
	}
	b.PutID(InputPrivacyValueAllowUsersTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyValueAllowUsers) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowUsers#131cc67f as nil")
	}
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode inputPrivacyValueAllowUsers#131cc67f: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputPrivacyValueAllowUsers#131cc67f: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowUsers) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowUsers#131cc67f to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowUsers#131cc67f: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyValueAllowUsers) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowUsers#131cc67f to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueAllowUsers#131cc67f: field users: %w", err)
		}

		if headerLen > 0 {
			i.Users = make([]InputUserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueAllowUsers#131cc67f: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// GetUsers returns value of Users field.
func (i *InputPrivacyValueAllowUsers) GetUsers() (value []InputUserClass) {
	if i == nil {
		return
	}
	return i.Users
}

// MapUsers returns field Users wrapped in InputUserClassArray helper.
func (i *InputPrivacyValueAllowUsers) MapUsers() (value InputUserClassArray) {
	return InputUserClassArray(i.Users)
}

// InputPrivacyValueDisallowContacts represents TL type `inputPrivacyValueDisallowContacts#ba52007`.
// Disallow only contacts
//
// See https://core.telegram.org/constructor/inputPrivacyValueDisallowContacts for reference.
type InputPrivacyValueDisallowContacts struct {
}

// InputPrivacyValueDisallowContactsTypeID is TL type id of InputPrivacyValueDisallowContacts.
const InputPrivacyValueDisallowContactsTypeID = 0xba52007

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowContacts) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowContacts.
var (
	_ bin.Encoder     = &InputPrivacyValueDisallowContacts{}
	_ bin.Decoder     = &InputPrivacyValueDisallowContacts{}
	_ bin.BareEncoder = &InputPrivacyValueDisallowContacts{}
	_ bin.BareDecoder = &InputPrivacyValueDisallowContacts{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowContacts{}
)

func (i *InputPrivacyValueDisallowContacts) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueDisallowContacts) String() string {
	if i == nil {
		return "InputPrivacyValueDisallowContacts(nil)"
	}
	type Alias InputPrivacyValueDisallowContacts
	return fmt.Sprintf("InputPrivacyValueDisallowContacts%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyValueDisallowContacts) TypeID() uint32 {
	return InputPrivacyValueDisallowContactsTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyValueDisallowContacts) TypeName() string {
	return "inputPrivacyValueDisallowContacts"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyValueDisallowContacts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyValueDisallowContacts",
		ID:   InputPrivacyValueDisallowContactsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowContacts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowContacts#ba52007 as nil")
	}
	b.PutID(InputPrivacyValueDisallowContactsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyValueDisallowContacts) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowContacts#ba52007 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowContacts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowContacts#ba52007 to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowContacts#ba52007: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyValueDisallowContacts) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowContacts#ba52007 to nil")
	}
	return nil
}

// InputPrivacyValueDisallowAll represents TL type `inputPrivacyValueDisallowAll#d66b66c9`.
// Disallow all
//
// See https://core.telegram.org/constructor/inputPrivacyValueDisallowAll for reference.
type InputPrivacyValueDisallowAll struct {
}

// InputPrivacyValueDisallowAllTypeID is TL type id of InputPrivacyValueDisallowAll.
const InputPrivacyValueDisallowAllTypeID = 0xd66b66c9

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowAll) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowAll.
var (
	_ bin.Encoder     = &InputPrivacyValueDisallowAll{}
	_ bin.Decoder     = &InputPrivacyValueDisallowAll{}
	_ bin.BareEncoder = &InputPrivacyValueDisallowAll{}
	_ bin.BareDecoder = &InputPrivacyValueDisallowAll{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowAll{}
)

func (i *InputPrivacyValueDisallowAll) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueDisallowAll) String() string {
	if i == nil {
		return "InputPrivacyValueDisallowAll(nil)"
	}
	type Alias InputPrivacyValueDisallowAll
	return fmt.Sprintf("InputPrivacyValueDisallowAll%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyValueDisallowAll) TypeID() uint32 {
	return InputPrivacyValueDisallowAllTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyValueDisallowAll) TypeName() string {
	return "inputPrivacyValueDisallowAll"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyValueDisallowAll) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyValueDisallowAll",
		ID:   InputPrivacyValueDisallowAllTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowAll) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowAll#d66b66c9 as nil")
	}
	b.PutID(InputPrivacyValueDisallowAllTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyValueDisallowAll) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowAll#d66b66c9 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowAll) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowAll#d66b66c9 to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowAllTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowAll#d66b66c9: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyValueDisallowAll) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowAll#d66b66c9 to nil")
	}
	return nil
}

// InputPrivacyValueDisallowUsers represents TL type `inputPrivacyValueDisallowUsers#90110467`.
// Disallow only certain users
//
// See https://core.telegram.org/constructor/inputPrivacyValueDisallowUsers for reference.
type InputPrivacyValueDisallowUsers struct {
	// Users to disallow
	Users []InputUserClass
}

// InputPrivacyValueDisallowUsersTypeID is TL type id of InputPrivacyValueDisallowUsers.
const InputPrivacyValueDisallowUsersTypeID = 0x90110467

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowUsers) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowUsers.
var (
	_ bin.Encoder     = &InputPrivacyValueDisallowUsers{}
	_ bin.Decoder     = &InputPrivacyValueDisallowUsers{}
	_ bin.BareEncoder = &InputPrivacyValueDisallowUsers{}
	_ bin.BareDecoder = &InputPrivacyValueDisallowUsers{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowUsers{}
)

func (i *InputPrivacyValueDisallowUsers) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueDisallowUsers) String() string {
	if i == nil {
		return "InputPrivacyValueDisallowUsers(nil)"
	}
	type Alias InputPrivacyValueDisallowUsers
	return fmt.Sprintf("InputPrivacyValueDisallowUsers%+v", Alias(*i))
}

// FillFrom fills InputPrivacyValueDisallowUsers from given interface.
func (i *InputPrivacyValueDisallowUsers) FillFrom(from interface {
	GetUsers() (value []InputUserClass)
}) {
	i.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyValueDisallowUsers) TypeID() uint32 {
	return InputPrivacyValueDisallowUsersTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyValueDisallowUsers) TypeName() string {
	return "inputPrivacyValueDisallowUsers"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyValueDisallowUsers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyValueDisallowUsers",
		ID:   InputPrivacyValueDisallowUsersTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowUsers) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowUsers#90110467 as nil")
	}
	b.PutID(InputPrivacyValueDisallowUsersTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyValueDisallowUsers) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowUsers#90110467 as nil")
	}
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode inputPrivacyValueDisallowUsers#90110467: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputPrivacyValueDisallowUsers#90110467: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowUsers) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowUsers#90110467 to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowUsers#90110467: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyValueDisallowUsers) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowUsers#90110467 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueDisallowUsers#90110467: field users: %w", err)
		}

		if headerLen > 0 {
			i.Users = make([]InputUserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueDisallowUsers#90110467: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// GetUsers returns value of Users field.
func (i *InputPrivacyValueDisallowUsers) GetUsers() (value []InputUserClass) {
	if i == nil {
		return
	}
	return i.Users
}

// MapUsers returns field Users wrapped in InputUserClassArray helper.
func (i *InputPrivacyValueDisallowUsers) MapUsers() (value InputUserClassArray) {
	return InputUserClassArray(i.Users)
}

// InputPrivacyValueAllowChatParticipants represents TL type `inputPrivacyValueAllowChatParticipants#840649cf`.
// Allow only participants of certain chats
//
// See https://core.telegram.org/constructor/inputPrivacyValueAllowChatParticipants for reference.
type InputPrivacyValueAllowChatParticipants struct {
	// Allowed chat IDs
	Chats []int64
}

// InputPrivacyValueAllowChatParticipantsTypeID is TL type id of InputPrivacyValueAllowChatParticipants.
const InputPrivacyValueAllowChatParticipantsTypeID = 0x840649cf

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowChatParticipants) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowChatParticipants.
var (
	_ bin.Encoder     = &InputPrivacyValueAllowChatParticipants{}
	_ bin.Decoder     = &InputPrivacyValueAllowChatParticipants{}
	_ bin.BareEncoder = &InputPrivacyValueAllowChatParticipants{}
	_ bin.BareDecoder = &InputPrivacyValueAllowChatParticipants{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowChatParticipants{}
)

func (i *InputPrivacyValueAllowChatParticipants) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Chats == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueAllowChatParticipants) String() string {
	if i == nil {
		return "InputPrivacyValueAllowChatParticipants(nil)"
	}
	type Alias InputPrivacyValueAllowChatParticipants
	return fmt.Sprintf("InputPrivacyValueAllowChatParticipants%+v", Alias(*i))
}

// FillFrom fills InputPrivacyValueAllowChatParticipants from given interface.
func (i *InputPrivacyValueAllowChatParticipants) FillFrom(from interface {
	GetChats() (value []int64)
}) {
	i.Chats = from.GetChats()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyValueAllowChatParticipants) TypeID() uint32 {
	return InputPrivacyValueAllowChatParticipantsTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyValueAllowChatParticipants) TypeName() string {
	return "inputPrivacyValueAllowChatParticipants"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyValueAllowChatParticipants) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyValueAllowChatParticipants",
		ID:   InputPrivacyValueAllowChatParticipantsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowChatParticipants) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowChatParticipants#840649cf as nil")
	}
	b.PutID(InputPrivacyValueAllowChatParticipantsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyValueAllowChatParticipants) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowChatParticipants#840649cf as nil")
	}
	b.PutVectorHeader(len(i.Chats))
	for _, v := range i.Chats {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowChatParticipants) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowChatParticipants#840649cf to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowChatParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowChatParticipants#840649cf: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyValueAllowChatParticipants) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowChatParticipants#840649cf to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueAllowChatParticipants#840649cf: field chats: %w", err)
		}

		if headerLen > 0 {
			i.Chats = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueAllowChatParticipants#840649cf: field chats: %w", err)
			}
			i.Chats = append(i.Chats, value)
		}
	}
	return nil
}

// GetChats returns value of Chats field.
func (i *InputPrivacyValueAllowChatParticipants) GetChats() (value []int64) {
	if i == nil {
		return
	}
	return i.Chats
}

// InputPrivacyValueDisallowChatParticipants represents TL type `inputPrivacyValueDisallowChatParticipants#e94f0f86`.
// Disallow only participants of certain chats
//
// See https://core.telegram.org/constructor/inputPrivacyValueDisallowChatParticipants for reference.
type InputPrivacyValueDisallowChatParticipants struct {
	// Disallowed chat IDs
	Chats []int64
}

// InputPrivacyValueDisallowChatParticipantsTypeID is TL type id of InputPrivacyValueDisallowChatParticipants.
const InputPrivacyValueDisallowChatParticipantsTypeID = 0xe94f0f86

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowChatParticipants) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowChatParticipants.
var (
	_ bin.Encoder     = &InputPrivacyValueDisallowChatParticipants{}
	_ bin.Decoder     = &InputPrivacyValueDisallowChatParticipants{}
	_ bin.BareEncoder = &InputPrivacyValueDisallowChatParticipants{}
	_ bin.BareDecoder = &InputPrivacyValueDisallowChatParticipants{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowChatParticipants{}
)

func (i *InputPrivacyValueDisallowChatParticipants) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Chats == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyValueDisallowChatParticipants) String() string {
	if i == nil {
		return "InputPrivacyValueDisallowChatParticipants(nil)"
	}
	type Alias InputPrivacyValueDisallowChatParticipants
	return fmt.Sprintf("InputPrivacyValueDisallowChatParticipants%+v", Alias(*i))
}

// FillFrom fills InputPrivacyValueDisallowChatParticipants from given interface.
func (i *InputPrivacyValueDisallowChatParticipants) FillFrom(from interface {
	GetChats() (value []int64)
}) {
	i.Chats = from.GetChats()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyValueDisallowChatParticipants) TypeID() uint32 {
	return InputPrivacyValueDisallowChatParticipantsTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyValueDisallowChatParticipants) TypeName() string {
	return "inputPrivacyValueDisallowChatParticipants"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyValueDisallowChatParticipants) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyValueDisallowChatParticipants",
		ID:   InputPrivacyValueDisallowChatParticipantsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowChatParticipants) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowChatParticipants#e94f0f86 as nil")
	}
	b.PutID(InputPrivacyValueDisallowChatParticipantsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyValueDisallowChatParticipants) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowChatParticipants#e94f0f86 as nil")
	}
	b.PutVectorHeader(len(i.Chats))
	for _, v := range i.Chats {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowChatParticipants) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowChatParticipants#e94f0f86 to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowChatParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowChatParticipants#e94f0f86: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyValueDisallowChatParticipants) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowChatParticipants#e94f0f86 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueDisallowChatParticipants#e94f0f86: field chats: %w", err)
		}

		if headerLen > 0 {
			i.Chats = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueDisallowChatParticipants#e94f0f86: field chats: %w", err)
			}
			i.Chats = append(i.Chats, value)
		}
	}
	return nil
}

// GetChats returns value of Chats field.
func (i *InputPrivacyValueDisallowChatParticipants) GetChats() (value []int64) {
	if i == nil {
		return
	}
	return i.Chats
}

// InputPrivacyRuleClassName is schema name of InputPrivacyRuleClass.
const InputPrivacyRuleClassName = "InputPrivacyRule"

// InputPrivacyRuleClass represents InputPrivacyRule generic type.
//
// See https://core.telegram.org/type/InputPrivacyRule for reference.
//
// Example:
//
//	g, err := tg.DecodeInputPrivacyRule(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputPrivacyValueAllowContacts: // inputPrivacyValueAllowContacts#d09e07b
//	case *tg.InputPrivacyValueAllowAll: // inputPrivacyValueAllowAll#184b35ce
//	case *tg.InputPrivacyValueAllowUsers: // inputPrivacyValueAllowUsers#131cc67f
//	case *tg.InputPrivacyValueDisallowContacts: // inputPrivacyValueDisallowContacts#ba52007
//	case *tg.InputPrivacyValueDisallowAll: // inputPrivacyValueDisallowAll#d66b66c9
//	case *tg.InputPrivacyValueDisallowUsers: // inputPrivacyValueDisallowUsers#90110467
//	case *tg.InputPrivacyValueAllowChatParticipants: // inputPrivacyValueAllowChatParticipants#840649cf
//	case *tg.InputPrivacyValueDisallowChatParticipants: // inputPrivacyValueDisallowChatParticipants#e94f0f86
//	default: panic(v)
//	}
type InputPrivacyRuleClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputPrivacyRuleClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputPrivacyRule implements binary de-serialization for InputPrivacyRuleClass.
func DecodeInputPrivacyRule(buf *bin.Buffer) (InputPrivacyRuleClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPrivacyValueAllowContactsTypeID:
		// Decoding inputPrivacyValueAllowContacts#d09e07b.
		v := InputPrivacyValueAllowContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueAllowAllTypeID:
		// Decoding inputPrivacyValueAllowAll#184b35ce.
		v := InputPrivacyValueAllowAll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueAllowUsersTypeID:
		// Decoding inputPrivacyValueAllowUsers#131cc67f.
		v := InputPrivacyValueAllowUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowContactsTypeID:
		// Decoding inputPrivacyValueDisallowContacts#ba52007.
		v := InputPrivacyValueDisallowContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowAllTypeID:
		// Decoding inputPrivacyValueDisallowAll#d66b66c9.
		v := InputPrivacyValueDisallowAll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowUsersTypeID:
		// Decoding inputPrivacyValueDisallowUsers#90110467.
		v := InputPrivacyValueDisallowUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueAllowChatParticipantsTypeID:
		// Decoding inputPrivacyValueAllowChatParticipants#840649cf.
		v := InputPrivacyValueAllowChatParticipants{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowChatParticipantsTypeID:
		// Decoding inputPrivacyValueDisallowChatParticipants#e94f0f86.
		v := InputPrivacyValueDisallowChatParticipants{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPrivacyRule boxes the InputPrivacyRuleClass providing a helper.
type InputPrivacyRuleBox struct {
	InputPrivacyRule InputPrivacyRuleClass
}

// Decode implements bin.Decoder for InputPrivacyRuleBox.
func (b *InputPrivacyRuleBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPrivacyRuleBox to nil")
	}
	v, err := DecodeInputPrivacyRule(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPrivacyRule = v
	return nil
}

// Encode implements bin.Encode for InputPrivacyRuleBox.
func (b *InputPrivacyRuleBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPrivacyRule == nil {
		return fmt.Errorf("unable to encode InputPrivacyRuleClass as nil")
	}
	return b.InputPrivacyRule.Encode(buf)
}
