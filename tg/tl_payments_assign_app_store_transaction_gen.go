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

// PaymentsAssignAppStoreTransactionRequest represents TL type `payments.assignAppStoreTransaction#80ed747d`.
// Informs server about a purchase made through the App Store: for official applications
// only.
//
// See https://core.telegram.org/method/payments.assignAppStoreTransaction for reference.
type PaymentsAssignAppStoreTransactionRequest struct {
	// Receipt
	Receipt []byte
	// Payment purpose
	Purpose InputStorePaymentPurposeClass
}

// PaymentsAssignAppStoreTransactionRequestTypeID is TL type id of PaymentsAssignAppStoreTransactionRequest.
const PaymentsAssignAppStoreTransactionRequestTypeID = 0x80ed747d

// Ensuring interfaces in compile-time for PaymentsAssignAppStoreTransactionRequest.
var (
	_ bin.Encoder     = &PaymentsAssignAppStoreTransactionRequest{}
	_ bin.Decoder     = &PaymentsAssignAppStoreTransactionRequest{}
	_ bin.BareEncoder = &PaymentsAssignAppStoreTransactionRequest{}
	_ bin.BareDecoder = &PaymentsAssignAppStoreTransactionRequest{}
)

func (a *PaymentsAssignAppStoreTransactionRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Receipt == nil) {
		return false
	}
	if !(a.Purpose == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *PaymentsAssignAppStoreTransactionRequest) String() string {
	if a == nil {
		return "PaymentsAssignAppStoreTransactionRequest(nil)"
	}
	type Alias PaymentsAssignAppStoreTransactionRequest
	return fmt.Sprintf("PaymentsAssignAppStoreTransactionRequest%+v", Alias(*a))
}

// FillFrom fills PaymentsAssignAppStoreTransactionRequest from given interface.
func (a *PaymentsAssignAppStoreTransactionRequest) FillFrom(from interface {
	GetReceipt() (value []byte)
	GetPurpose() (value InputStorePaymentPurposeClass)
}) {
	a.Receipt = from.GetReceipt()
	a.Purpose = from.GetPurpose()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsAssignAppStoreTransactionRequest) TypeID() uint32 {
	return PaymentsAssignAppStoreTransactionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsAssignAppStoreTransactionRequest) TypeName() string {
	return "payments.assignAppStoreTransaction"
}

// TypeInfo returns info about TL type.
func (a *PaymentsAssignAppStoreTransactionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.assignAppStoreTransaction",
		ID:   PaymentsAssignAppStoreTransactionRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Receipt",
			SchemaName: "receipt",
		},
		{
			Name:       "Purpose",
			SchemaName: "purpose",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *PaymentsAssignAppStoreTransactionRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode payments.assignAppStoreTransaction#80ed747d as nil")
	}
	b.PutID(PaymentsAssignAppStoreTransactionRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *PaymentsAssignAppStoreTransactionRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode payments.assignAppStoreTransaction#80ed747d as nil")
	}
	b.PutBytes(a.Receipt)
	if a.Purpose == nil {
		return fmt.Errorf("unable to encode payments.assignAppStoreTransaction#80ed747d: field purpose is nil")
	}
	if err := a.Purpose.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.assignAppStoreTransaction#80ed747d: field purpose: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *PaymentsAssignAppStoreTransactionRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode payments.assignAppStoreTransaction#80ed747d to nil")
	}
	if err := b.ConsumeID(PaymentsAssignAppStoreTransactionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.assignAppStoreTransaction#80ed747d: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *PaymentsAssignAppStoreTransactionRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode payments.assignAppStoreTransaction#80ed747d to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode payments.assignAppStoreTransaction#80ed747d: field receipt: %w", err)
		}
		a.Receipt = value
	}
	{
		value, err := DecodeInputStorePaymentPurpose(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.assignAppStoreTransaction#80ed747d: field purpose: %w", err)
		}
		a.Purpose = value
	}
	return nil
}

// GetReceipt returns value of Receipt field.
func (a *PaymentsAssignAppStoreTransactionRequest) GetReceipt() (value []byte) {
	if a == nil {
		return
	}
	return a.Receipt
}

// GetPurpose returns value of Purpose field.
func (a *PaymentsAssignAppStoreTransactionRequest) GetPurpose() (value InputStorePaymentPurposeClass) {
	if a == nil {
		return
	}
	return a.Purpose
}

// PaymentsAssignAppStoreTransaction invokes method payments.assignAppStoreTransaction#80ed747d returning error if any.
// Informs server about a purchase made through the App Store: for official applications
// only.
//
// See https://core.telegram.org/method/payments.assignAppStoreTransaction for reference.
func (c *Client) PaymentsAssignAppStoreTransaction(ctx context.Context, request *PaymentsAssignAppStoreTransactionRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
