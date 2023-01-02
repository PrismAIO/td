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

// RecoveryEmailAddress represents TL type `recoveryEmailAddress#4cebddeb`.
type RecoveryEmailAddress struct {
	// Recovery email address
	RecoveryEmailAddress string
}

// RecoveryEmailAddressTypeID is TL type id of RecoveryEmailAddress.
const RecoveryEmailAddressTypeID = 0x4cebddeb

// Ensuring interfaces in compile-time for RecoveryEmailAddress.
var (
	_ bin.Encoder     = &RecoveryEmailAddress{}
	_ bin.Decoder     = &RecoveryEmailAddress{}
	_ bin.BareEncoder = &RecoveryEmailAddress{}
	_ bin.BareDecoder = &RecoveryEmailAddress{}
)

func (r *RecoveryEmailAddress) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.RecoveryEmailAddress == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecoveryEmailAddress) String() string {
	if r == nil {
		return "RecoveryEmailAddress(nil)"
	}
	type Alias RecoveryEmailAddress
	return fmt.Sprintf("RecoveryEmailAddress%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecoveryEmailAddress) TypeID() uint32 {
	return RecoveryEmailAddressTypeID
}

// TypeName returns name of type in TL schema.
func (*RecoveryEmailAddress) TypeName() string {
	return "recoveryEmailAddress"
}

// TypeInfo returns info about TL type.
func (r *RecoveryEmailAddress) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recoveryEmailAddress",
		ID:   RecoveryEmailAddressTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RecoveryEmailAddress",
			SchemaName: "recovery_email_address",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecoveryEmailAddress) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recoveryEmailAddress#4cebddeb as nil")
	}
	b.PutID(RecoveryEmailAddressTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecoveryEmailAddress) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recoveryEmailAddress#4cebddeb as nil")
	}
	b.PutString(r.RecoveryEmailAddress)
	return nil
}

// Decode implements bin.Decoder.
func (r *RecoveryEmailAddress) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recoveryEmailAddress#4cebddeb to nil")
	}
	if err := b.ConsumeID(RecoveryEmailAddressTypeID); err != nil {
		return fmt.Errorf("unable to decode recoveryEmailAddress#4cebddeb: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecoveryEmailAddress) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recoveryEmailAddress#4cebddeb to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recoveryEmailAddress#4cebddeb: field recovery_email_address: %w", err)
		}
		r.RecoveryEmailAddress = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RecoveryEmailAddress) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode recoveryEmailAddress#4cebddeb as nil")
	}
	b.ObjStart()
	b.PutID("recoveryEmailAddress")
	b.Comma()
	b.FieldStart("recovery_email_address")
	b.PutString(r.RecoveryEmailAddress)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RecoveryEmailAddress) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode recoveryEmailAddress#4cebddeb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("recoveryEmailAddress"); err != nil {
				return fmt.Errorf("unable to decode recoveryEmailAddress#4cebddeb: %w", err)
			}
		case "recovery_email_address":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode recoveryEmailAddress#4cebddeb: field recovery_email_address: %w", err)
			}
			r.RecoveryEmailAddress = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRecoveryEmailAddress returns value of RecoveryEmailAddress field.
func (r *RecoveryEmailAddress) GetRecoveryEmailAddress() (value string) {
	if r == nil {
		return
	}
	return r.RecoveryEmailAddress
}
