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

// TestInt represents TL type `testInt#ddbd2c09`.
type TestInt struct {
	// Number
	Value int32
}

// TestIntTypeID is TL type id of TestInt.
const TestIntTypeID = 0xddbd2c09

// Ensuring interfaces in compile-time for TestInt.
var (
	_ bin.Encoder     = &TestInt{}
	_ bin.Decoder     = &TestInt{}
	_ bin.BareEncoder = &TestInt{}
	_ bin.BareDecoder = &TestInt{}
)

func (t *TestInt) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Value == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestInt) String() string {
	if t == nil {
		return "TestInt(nil)"
	}
	type Alias TestInt
	return fmt.Sprintf("TestInt%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestInt) TypeID() uint32 {
	return TestIntTypeID
}

// TypeName returns name of type in TL schema.
func (*TestInt) TypeName() string {
	return "testInt"
}

// TypeInfo returns info about TL type.
func (t *TestInt) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testInt",
		ID:   TestIntTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestInt) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testInt#ddbd2c09 as nil")
	}
	b.PutID(TestIntTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestInt) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testInt#ddbd2c09 as nil")
	}
	b.PutInt32(t.Value)
	return nil
}

// Decode implements bin.Decoder.
func (t *TestInt) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testInt#ddbd2c09 to nil")
	}
	if err := b.ConsumeID(TestIntTypeID); err != nil {
		return fmt.Errorf("unable to decode testInt#ddbd2c09: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestInt) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testInt#ddbd2c09 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode testInt#ddbd2c09: field value: %w", err)
		}
		t.Value = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TestInt) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode testInt#ddbd2c09 as nil")
	}
	b.ObjStart()
	b.PutID("testInt")
	b.Comma()
	b.FieldStart("value")
	b.PutInt32(t.Value)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TestInt) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode testInt#ddbd2c09 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("testInt"); err != nil {
				return fmt.Errorf("unable to decode testInt#ddbd2c09: %w", err)
			}
		case "value":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode testInt#ddbd2c09: field value: %w", err)
			}
			t.Value = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetValue returns value of Value field.
func (t *TestInt) GetValue() (value int32) {
	if t == nil {
		return
	}
	return t.Value
}
