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

// Double represents TL type `double#2210c154`.
type Double struct {
}

// DoubleTypeID is TL type id of Double.
const DoubleTypeID = 0x2210c154

// Ensuring interfaces in compile-time for Double.
var (
	_ bin.Encoder     = &Double{}
	_ bin.Decoder     = &Double{}
	_ bin.BareEncoder = &Double{}
	_ bin.BareDecoder = &Double{}
)

func (d *Double) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *Double) String() string {
	if d == nil {
		return "Double(nil)"
	}
	type Alias Double
	return fmt.Sprintf("Double%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Double) TypeID() uint32 {
	return DoubleTypeID
}

// TypeName returns name of type in TL schema.
func (*Double) TypeName() string {
	return "double"
}

// TypeInfo returns info about TL type.
func (d *Double) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "double",
		ID:   DoubleTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (d *Double) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode double#2210c154 as nil")
	}
	b.PutID(DoubleTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *Double) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode double#2210c154 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *Double) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode double#2210c154 to nil")
	}
	if err := b.ConsumeID(DoubleTypeID); err != nil {
		return fmt.Errorf("unable to decode double#2210c154: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *Double) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode double#2210c154 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *Double) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode double#2210c154 as nil")
	}
	b.ObjStart()
	b.PutID("double")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *Double) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode double#2210c154 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("double"); err != nil {
				return fmt.Errorf("unable to decode double#2210c154: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}
