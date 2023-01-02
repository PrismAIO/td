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

// TestString represents TL type `testString#fe56688c`.
type TestString struct {
	// String
	Value string
}

// TestStringTypeID is TL type id of TestString.
const TestStringTypeID = 0xfe56688c

// Ensuring interfaces in compile-time for TestString.
var (
	_ bin.Encoder     = &TestString{}
	_ bin.Decoder     = &TestString{}
	_ bin.BareEncoder = &TestString{}
	_ bin.BareDecoder = &TestString{}
)

func (t *TestString) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Value == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestString) String() string {
	if t == nil {
		return "TestString(nil)"
	}
	type Alias TestString
	return fmt.Sprintf("TestString%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestString) TypeID() uint32 {
	return TestStringTypeID
}

// TypeName returns name of type in TL schema.
func (*TestString) TypeName() string {
	return "testString"
}

// TypeInfo returns info about TL type.
func (t *TestString) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testString",
		ID:   TestStringTypeID,
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
func (t *TestString) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testString#fe56688c as nil")
	}
	b.PutID(TestStringTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestString) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testString#fe56688c as nil")
	}
	b.PutString(t.Value)
	return nil
}

// Decode implements bin.Decoder.
func (t *TestString) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testString#fe56688c to nil")
	}
	if err := b.ConsumeID(TestStringTypeID); err != nil {
		return fmt.Errorf("unable to decode testString#fe56688c: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestString) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testString#fe56688c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode testString#fe56688c: field value: %w", err)
		}
		t.Value = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TestString) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode testString#fe56688c as nil")
	}
	b.ObjStart()
	b.PutID("testString")
	b.Comma()
	b.FieldStart("value")
	b.PutString(t.Value)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TestString) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode testString#fe56688c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("testString"); err != nil {
				return fmt.Errorf("unable to decode testString#fe56688c: %w", err)
			}
		case "value":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode testString#fe56688c: field value: %w", err)
			}
			t.Value = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetValue returns value of Value field.
func (t *TestString) GetValue() (value string) {
	if t == nil {
		return
	}
	return t.Value
}
