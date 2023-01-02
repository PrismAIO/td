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

// Int32 represents TL type `int32#5cb934fa`.
type Int32 struct {
}

// Int32TypeID is TL type id of Int32.
const Int32TypeID = 0x5cb934fa

// Ensuring interfaces in compile-time for Int32.
var (
	_ bin.Encoder     = &Int32{}
	_ bin.Decoder     = &Int32{}
	_ bin.BareEncoder = &Int32{}
	_ bin.BareDecoder = &Int32{}
)

func (i *Int32) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *Int32) String() string {
	if i == nil {
		return "Int32(nil)"
	}
	type Alias Int32
	return fmt.Sprintf("Int32%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Int32) TypeID() uint32 {
	return Int32TypeID
}

// TypeName returns name of type in TL schema.
func (*Int32) TypeName() string {
	return "int32"
}

// TypeInfo returns info about TL type.
func (i *Int32) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "int32",
		ID:   Int32TypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *Int32) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode int32#5cb934fa as nil")
	}
	b.PutID(Int32TypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *Int32) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode int32#5cb934fa as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *Int32) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode int32#5cb934fa to nil")
	}
	if err := b.ConsumeID(Int32TypeID); err != nil {
		return fmt.Errorf("unable to decode int32#5cb934fa: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *Int32) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode int32#5cb934fa to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *Int32) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode int32#5cb934fa as nil")
	}
	b.ObjStart()
	b.PutID("int32")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *Int32) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode int32#5cb934fa to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("int32"); err != nil {
				return fmt.Errorf("unable to decode int32#5cb934fa: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}
