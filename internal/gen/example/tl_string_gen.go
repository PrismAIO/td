// Code generated by gotdgen, DO NOT EDIT.

package td

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

// String represents TL type `string#b5286e24`.
//
// See https://localhost:80/doc/constructor/string for reference.
type String struct {
}

// StringTypeID is TL type id of String.
const StringTypeID = 0xb5286e24

// Ensuring interfaces in compile-time for String.
var (
	_ bin.Encoder     = &String{}
	_ bin.Decoder     = &String{}
	_ bin.BareEncoder = &String{}
	_ bin.BareDecoder = &String{}
)

func (s *String) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *String) String() string {
	if s == nil {
		return "String(nil)"
	}
	type Alias String
	return fmt.Sprintf("String%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*String) TypeID() uint32 {
	return StringTypeID
}

// TypeName returns name of type in TL schema.
func (*String) TypeName() string {
	return "string"
}

// TypeInfo returns info about TL type.
func (s *String) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "string",
		ID:   StringTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *String) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode string#b5286e24 as nil")
	}
	b.PutID(StringTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *String) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode string#b5286e24 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *String) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode string#b5286e24 to nil")
	}
	if err := b.ConsumeID(StringTypeID); err != nil {
		return fmt.Errorf("unable to decode string#b5286e24: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *String) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode string#b5286e24 to nil")
	}
	return nil
}
