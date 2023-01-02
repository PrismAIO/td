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

// AttachMenuBotsNotModified represents TL type `attachMenuBotsNotModified#f1d88a5c`.
// The list of bot web apps hasn't changed
//
// See https://core.telegram.org/constructor/attachMenuBotsNotModified for reference.
type AttachMenuBotsNotModified struct {
}

// AttachMenuBotsNotModifiedTypeID is TL type id of AttachMenuBotsNotModified.
const AttachMenuBotsNotModifiedTypeID = 0xf1d88a5c

// construct implements constructor of AttachMenuBotsClass.
func (a AttachMenuBotsNotModified) construct() AttachMenuBotsClass { return &a }

// Ensuring interfaces in compile-time for AttachMenuBotsNotModified.
var (
	_ bin.Encoder     = &AttachMenuBotsNotModified{}
	_ bin.Decoder     = &AttachMenuBotsNotModified{}
	_ bin.BareEncoder = &AttachMenuBotsNotModified{}
	_ bin.BareDecoder = &AttachMenuBotsNotModified{}

	_ AttachMenuBotsClass = &AttachMenuBotsNotModified{}
)

func (a *AttachMenuBotsNotModified) Zero() bool {
	if a == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (a *AttachMenuBotsNotModified) String() string {
	if a == nil {
		return "AttachMenuBotsNotModified(nil)"
	}
	type Alias AttachMenuBotsNotModified
	return fmt.Sprintf("AttachMenuBotsNotModified%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AttachMenuBotsNotModified) TypeID() uint32 {
	return AttachMenuBotsNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*AttachMenuBotsNotModified) TypeName() string {
	return "attachMenuBotsNotModified"
}

// TypeInfo returns info about TL type.
func (a *AttachMenuBotsNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "attachMenuBotsNotModified",
		ID:   AttachMenuBotsNotModifiedTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (a *AttachMenuBotsNotModified) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuBotsNotModified#f1d88a5c as nil")
	}
	b.PutID(AttachMenuBotsNotModifiedTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AttachMenuBotsNotModified) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuBotsNotModified#f1d88a5c as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AttachMenuBotsNotModified) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuBotsNotModified#f1d88a5c to nil")
	}
	if err := b.ConsumeID(AttachMenuBotsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode attachMenuBotsNotModified#f1d88a5c: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AttachMenuBotsNotModified) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuBotsNotModified#f1d88a5c to nil")
	}
	return nil
}

// AttachMenuBots represents TL type `attachMenuBots#3c4301c0`.
// Represents a list of bot web apps that can be launched from the attachment menu »¹
//
// Links:
//  1. https://core.telegram.org/api/bots/attach
//
// See https://core.telegram.org/constructor/attachMenuBots for reference.
type AttachMenuBots struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
	// List of bot web apps that can be launched from the attachment menu »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/attach
	Bots []AttachMenuBot
	// Info about related users/bots
	Users []UserClass
}

// AttachMenuBotsTypeID is TL type id of AttachMenuBots.
const AttachMenuBotsTypeID = 0x3c4301c0

// construct implements constructor of AttachMenuBotsClass.
func (a AttachMenuBots) construct() AttachMenuBotsClass { return &a }

// Ensuring interfaces in compile-time for AttachMenuBots.
var (
	_ bin.Encoder     = &AttachMenuBots{}
	_ bin.Decoder     = &AttachMenuBots{}
	_ bin.BareEncoder = &AttachMenuBots{}
	_ bin.BareDecoder = &AttachMenuBots{}

	_ AttachMenuBotsClass = &AttachMenuBots{}
)

func (a *AttachMenuBots) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Hash == 0) {
		return false
	}
	if !(a.Bots == nil) {
		return false
	}
	if !(a.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AttachMenuBots) String() string {
	if a == nil {
		return "AttachMenuBots(nil)"
	}
	type Alias AttachMenuBots
	return fmt.Sprintf("AttachMenuBots%+v", Alias(*a))
}

// FillFrom fills AttachMenuBots from given interface.
func (a *AttachMenuBots) FillFrom(from interface {
	GetHash() (value int64)
	GetBots() (value []AttachMenuBot)
	GetUsers() (value []UserClass)
}) {
	a.Hash = from.GetHash()
	a.Bots = from.GetBots()
	a.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AttachMenuBots) TypeID() uint32 {
	return AttachMenuBotsTypeID
}

// TypeName returns name of type in TL schema.
func (*AttachMenuBots) TypeName() string {
	return "attachMenuBots"
}

// TypeInfo returns info about TL type.
func (a *AttachMenuBots) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "attachMenuBots",
		ID:   AttachMenuBotsTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Bots",
			SchemaName: "bots",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AttachMenuBots) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuBots#3c4301c0 as nil")
	}
	b.PutID(AttachMenuBotsTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AttachMenuBots) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode attachMenuBots#3c4301c0 as nil")
	}
	b.PutLong(a.Hash)
	b.PutVectorHeader(len(a.Bots))
	for idx, v := range a.Bots {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode attachMenuBots#3c4301c0: field bots element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(a.Users))
	for idx, v := range a.Users {
		if v == nil {
			return fmt.Errorf("unable to encode attachMenuBots#3c4301c0: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode attachMenuBots#3c4301c0: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AttachMenuBots) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuBots#3c4301c0 to nil")
	}
	if err := b.ConsumeID(AttachMenuBotsTypeID); err != nil {
		return fmt.Errorf("unable to decode attachMenuBots#3c4301c0: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AttachMenuBots) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode attachMenuBots#3c4301c0 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode attachMenuBots#3c4301c0: field hash: %w", err)
		}
		a.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode attachMenuBots#3c4301c0: field bots: %w", err)
		}

		if headerLen > 0 {
			a.Bots = make([]AttachMenuBot, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value AttachMenuBot
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode attachMenuBots#3c4301c0: field bots: %w", err)
			}
			a.Bots = append(a.Bots, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode attachMenuBots#3c4301c0: field users: %w", err)
		}

		if headerLen > 0 {
			a.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode attachMenuBots#3c4301c0: field users: %w", err)
			}
			a.Users = append(a.Users, value)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (a *AttachMenuBots) GetHash() (value int64) {
	if a == nil {
		return
	}
	return a.Hash
}

// GetBots returns value of Bots field.
func (a *AttachMenuBots) GetBots() (value []AttachMenuBot) {
	if a == nil {
		return
	}
	return a.Bots
}

// GetUsers returns value of Users field.
func (a *AttachMenuBots) GetUsers() (value []UserClass) {
	if a == nil {
		return
	}
	return a.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (a *AttachMenuBots) MapUsers() (value UserClassArray) {
	return UserClassArray(a.Users)
}

// AttachMenuBotsClassName is schema name of AttachMenuBotsClass.
const AttachMenuBotsClassName = "AttachMenuBots"

// AttachMenuBotsClass represents AttachMenuBots generic type.
//
// See https://core.telegram.org/type/AttachMenuBots for reference.
//
// Example:
//
//	g, err := tg.DecodeAttachMenuBots(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.AttachMenuBotsNotModified: // attachMenuBotsNotModified#f1d88a5c
//	case *tg.AttachMenuBots: // attachMenuBots#3c4301c0
//	default: panic(v)
//	}
type AttachMenuBotsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AttachMenuBotsClass

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

	// AsModified tries to map AttachMenuBotsClass to AttachMenuBots.
	AsModified() (*AttachMenuBots, bool)
}

// AsModified tries to map AttachMenuBotsNotModified to AttachMenuBots.
func (a *AttachMenuBotsNotModified) AsModified() (*AttachMenuBots, bool) {
	return nil, false
}

// AsModified tries to map AttachMenuBots to AttachMenuBots.
func (a *AttachMenuBots) AsModified() (*AttachMenuBots, bool) {
	return a, true
}

// DecodeAttachMenuBots implements binary de-serialization for AttachMenuBotsClass.
func DecodeAttachMenuBots(buf *bin.Buffer) (AttachMenuBotsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AttachMenuBotsNotModifiedTypeID:
		// Decoding attachMenuBotsNotModified#f1d88a5c.
		v := AttachMenuBotsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AttachMenuBotsClass: %w", err)
		}
		return &v, nil
	case AttachMenuBotsTypeID:
		// Decoding attachMenuBots#3c4301c0.
		v := AttachMenuBots{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AttachMenuBotsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AttachMenuBotsClass: %w", bin.NewUnexpectedID(id))
	}
}

// AttachMenuBots boxes the AttachMenuBotsClass providing a helper.
type AttachMenuBotsBox struct {
	AttachMenuBots AttachMenuBotsClass
}

// Decode implements bin.Decoder for AttachMenuBotsBox.
func (b *AttachMenuBotsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AttachMenuBotsBox to nil")
	}
	v, err := DecodeAttachMenuBots(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AttachMenuBots = v
	return nil
}

// Encode implements bin.Encode for AttachMenuBotsBox.
func (b *AttachMenuBotsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.AttachMenuBots == nil {
		return fmt.Errorf("unable to encode AttachMenuBotsClass as nil")
	}
	return b.AttachMenuBots.Encode(buf)
}
