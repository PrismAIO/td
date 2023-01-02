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

// ValidateOrderInfoRequest represents TL type `validateOrderInfo#b5985fcf`.
type ValidateOrderInfoRequest struct {
	// The invoice
	InputInvoice InputInvoiceClass
	// The order information, provided by the user; pass null if empty
	OrderInfo OrderInfo
	// Pass true to save the order information
	AllowSave bool
}

// ValidateOrderInfoRequestTypeID is TL type id of ValidateOrderInfoRequest.
const ValidateOrderInfoRequestTypeID = 0xb5985fcf

// Ensuring interfaces in compile-time for ValidateOrderInfoRequest.
var (
	_ bin.Encoder     = &ValidateOrderInfoRequest{}
	_ bin.Decoder     = &ValidateOrderInfoRequest{}
	_ bin.BareEncoder = &ValidateOrderInfoRequest{}
	_ bin.BareDecoder = &ValidateOrderInfoRequest{}
)

func (v *ValidateOrderInfoRequest) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.InputInvoice == nil) {
		return false
	}
	if !(v.OrderInfo.Zero()) {
		return false
	}
	if !(v.AllowSave == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *ValidateOrderInfoRequest) String() string {
	if v == nil {
		return "ValidateOrderInfoRequest(nil)"
	}
	type Alias ValidateOrderInfoRequest
	return fmt.Sprintf("ValidateOrderInfoRequest%+v", Alias(*v))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ValidateOrderInfoRequest) TypeID() uint32 {
	return ValidateOrderInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ValidateOrderInfoRequest) TypeName() string {
	return "validateOrderInfo"
}

// TypeInfo returns info about TL type.
func (v *ValidateOrderInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "validateOrderInfo",
		ID:   ValidateOrderInfoRequestTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InputInvoice",
			SchemaName: "input_invoice",
		},
		{
			Name:       "OrderInfo",
			SchemaName: "order_info",
		},
		{
			Name:       "AllowSave",
			SchemaName: "allow_save",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *ValidateOrderInfoRequest) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode validateOrderInfo#b5985fcf as nil")
	}
	b.PutID(ValidateOrderInfoRequestTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *ValidateOrderInfoRequest) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode validateOrderInfo#b5985fcf as nil")
	}
	if v.InputInvoice == nil {
		return fmt.Errorf("unable to encode validateOrderInfo#b5985fcf: field input_invoice is nil")
	}
	if err := v.InputInvoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode validateOrderInfo#b5985fcf: field input_invoice: %w", err)
	}
	if err := v.OrderInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode validateOrderInfo#b5985fcf: field order_info: %w", err)
	}
	b.PutBool(v.AllowSave)
	return nil
}

// Decode implements bin.Decoder.
func (v *ValidateOrderInfoRequest) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode validateOrderInfo#b5985fcf to nil")
	}
	if err := b.ConsumeID(ValidateOrderInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode validateOrderInfo#b5985fcf: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *ValidateOrderInfoRequest) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode validateOrderInfo#b5985fcf to nil")
	}
	{
		value, err := DecodeInputInvoice(b)
		if err != nil {
			return fmt.Errorf("unable to decode validateOrderInfo#b5985fcf: field input_invoice: %w", err)
		}
		v.InputInvoice = value
	}
	{
		if err := v.OrderInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode validateOrderInfo#b5985fcf: field order_info: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode validateOrderInfo#b5985fcf: field allow_save: %w", err)
		}
		v.AllowSave = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (v *ValidateOrderInfoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if v == nil {
		return fmt.Errorf("can't encode validateOrderInfo#b5985fcf as nil")
	}
	b.ObjStart()
	b.PutID("validateOrderInfo")
	b.Comma()
	b.FieldStart("input_invoice")
	if v.InputInvoice == nil {
		return fmt.Errorf("unable to encode validateOrderInfo#b5985fcf: field input_invoice is nil")
	}
	if err := v.InputInvoice.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode validateOrderInfo#b5985fcf: field input_invoice: %w", err)
	}
	b.Comma()
	b.FieldStart("order_info")
	if err := v.OrderInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode validateOrderInfo#b5985fcf: field order_info: %w", err)
	}
	b.Comma()
	b.FieldStart("allow_save")
	b.PutBool(v.AllowSave)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (v *ValidateOrderInfoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if v == nil {
		return fmt.Errorf("can't decode validateOrderInfo#b5985fcf to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("validateOrderInfo"); err != nil {
				return fmt.Errorf("unable to decode validateOrderInfo#b5985fcf: %w", err)
			}
		case "input_invoice":
			value, err := DecodeTDLibJSONInputInvoice(b)
			if err != nil {
				return fmt.Errorf("unable to decode validateOrderInfo#b5985fcf: field input_invoice: %w", err)
			}
			v.InputInvoice = value
		case "order_info":
			if err := v.OrderInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode validateOrderInfo#b5985fcf: field order_info: %w", err)
			}
		case "allow_save":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode validateOrderInfo#b5985fcf: field allow_save: %w", err)
			}
			v.AllowSave = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInputInvoice returns value of InputInvoice field.
func (v *ValidateOrderInfoRequest) GetInputInvoice() (value InputInvoiceClass) {
	if v == nil {
		return
	}
	return v.InputInvoice
}

// GetOrderInfo returns value of OrderInfo field.
func (v *ValidateOrderInfoRequest) GetOrderInfo() (value OrderInfo) {
	if v == nil {
		return
	}
	return v.OrderInfo
}

// GetAllowSave returns value of AllowSave field.
func (v *ValidateOrderInfoRequest) GetAllowSave() (value bool) {
	if v == nil {
		return
	}
	return v.AllowSave
}

// ValidateOrderInfo invokes method validateOrderInfo#b5985fcf returning error if any.
func (c *Client) ValidateOrderInfo(ctx context.Context, request *ValidateOrderInfoRequest) (*ValidatedOrderInfo, error) {
	var result ValidatedOrderInfo

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
