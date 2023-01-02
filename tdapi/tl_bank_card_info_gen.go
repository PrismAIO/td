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

// BankCardInfo represents TL type `bankCardInfo#2bc7da9f`.
type BankCardInfo struct {
	// Title of the bank card description
	Title string
	// Actions that can be done with the bank card number
	Actions []BankCardActionOpenURL
}

// BankCardInfoTypeID is TL type id of BankCardInfo.
const BankCardInfoTypeID = 0x2bc7da9f

// Ensuring interfaces in compile-time for BankCardInfo.
var (
	_ bin.Encoder     = &BankCardInfo{}
	_ bin.Decoder     = &BankCardInfo{}
	_ bin.BareEncoder = &BankCardInfo{}
	_ bin.BareDecoder = &BankCardInfo{}
)

func (b *BankCardInfo) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Title == "") {
		return false
	}
	if !(b.Actions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BankCardInfo) String() string {
	if b == nil {
		return "BankCardInfo(nil)"
	}
	type Alias BankCardInfo
	return fmt.Sprintf("BankCardInfo%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BankCardInfo) TypeID() uint32 {
	return BankCardInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*BankCardInfo) TypeName() string {
	return "bankCardInfo"
}

// TypeInfo returns info about TL type.
func (b *BankCardInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bankCardInfo",
		ID:   BankCardInfoTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Actions",
			SchemaName: "actions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BankCardInfo) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bankCardInfo#2bc7da9f as nil")
	}
	buf.PutID(BankCardInfoTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BankCardInfo) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bankCardInfo#2bc7da9f as nil")
	}
	buf.PutString(b.Title)
	buf.PutInt(len(b.Actions))
	for idx, v := range b.Actions {
		if err := v.EncodeBare(buf); err != nil {
			return fmt.Errorf("unable to encode bare bankCardInfo#2bc7da9f: field actions element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BankCardInfo) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bankCardInfo#2bc7da9f to nil")
	}
	if err := buf.ConsumeID(BankCardInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode bankCardInfo#2bc7da9f: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BankCardInfo) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bankCardInfo#2bc7da9f to nil")
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode bankCardInfo#2bc7da9f: field title: %w", err)
		}
		b.Title = value
	}
	{
		headerLen, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode bankCardInfo#2bc7da9f: field actions: %w", err)
		}

		if headerLen > 0 {
			b.Actions = make([]BankCardActionOpenURL, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BankCardActionOpenURL
			if err := value.DecodeBare(buf); err != nil {
				return fmt.Errorf("unable to decode bare bankCardInfo#2bc7da9f: field actions: %w", err)
			}
			b.Actions = append(b.Actions, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BankCardInfo) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode bankCardInfo#2bc7da9f as nil")
	}
	buf.ObjStart()
	buf.PutID("bankCardInfo")
	buf.Comma()
	buf.FieldStart("title")
	buf.PutString(b.Title)
	buf.Comma()
	buf.FieldStart("actions")
	buf.ArrStart()
	for idx, v := range b.Actions {
		if err := v.EncodeTDLibJSON(buf); err != nil {
			return fmt.Errorf("unable to encode bankCardInfo#2bc7da9f: field actions element with index %d: %w", idx, err)
		}
		buf.Comma()
	}
	buf.StripComma()
	buf.ArrEnd()
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BankCardInfo) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode bankCardInfo#2bc7da9f to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("bankCardInfo"); err != nil {
				return fmt.Errorf("unable to decode bankCardInfo#2bc7da9f: %w", err)
			}
		case "title":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode bankCardInfo#2bc7da9f: field title: %w", err)
			}
			b.Title = value
		case "actions":
			if err := buf.Arr(func(buf tdjson.Decoder) error {
				var value BankCardActionOpenURL
				if err := value.DecodeTDLibJSON(buf); err != nil {
					return fmt.Errorf("unable to decode bankCardInfo#2bc7da9f: field actions: %w", err)
				}
				b.Actions = append(b.Actions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode bankCardInfo#2bc7da9f: field actions: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetTitle returns value of Title field.
func (b *BankCardInfo) GetTitle() (value string) {
	if b == nil {
		return
	}
	return b.Title
}

// GetActions returns value of Actions field.
func (b *BankCardInfo) GetActions() (value []BankCardActionOpenURL) {
	if b == nil {
		return
	}
	return b.Actions
}
