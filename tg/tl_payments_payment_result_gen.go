// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// PaymentsPaymentResult represents TL type `payments.paymentResult#4e5f810d`.
type PaymentsPaymentResult struct {
	// Updates field of PaymentsPaymentResult.
	Updates UpdatesClass
}

// PaymentsPaymentResultTypeID is TL type id of PaymentsPaymentResult.
const PaymentsPaymentResultTypeID = 0x4e5f810d

// Encode implements bin.Encoder.
func (p *PaymentsPaymentResult) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode payments.paymentResult#4e5f810d as nil")
	}
	b.PutID(PaymentsPaymentResultTypeID)
	if p.Updates == nil {
		return fmt.Errorf("unable to encode payments.paymentResult#4e5f810d: field updates is nil")
	}
	if err := p.Updates.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.paymentResult#4e5f810d: field updates: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentsPaymentResult) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode payments.paymentResult#4e5f810d to nil")
	}
	if err := b.ConsumeID(PaymentsPaymentResultTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.paymentResult#4e5f810d: %w", err)
	}
	{
		value, err := DecodeUpdates(b)
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentResult#4e5f810d: field updates: %w", err)
		}
		p.Updates = value
	}
	return nil
}

// construct implements constructor of PaymentsPaymentResultClass.
func (p PaymentsPaymentResult) construct() PaymentsPaymentResultClass { return &p }

// Ensuring interfaces in compile-time for PaymentsPaymentResult.
var (
	_ bin.Encoder = &PaymentsPaymentResult{}
	_ bin.Decoder = &PaymentsPaymentResult{}

	_ PaymentsPaymentResultClass = &PaymentsPaymentResult{}
)

// PaymentsPaymentVerificationNeeded represents TL type `payments.paymentVerificationNeeded#d8411139`.
type PaymentsPaymentVerificationNeeded struct {
	// URL field of PaymentsPaymentVerificationNeeded.
	URL string
}

// PaymentsPaymentVerificationNeededTypeID is TL type id of PaymentsPaymentVerificationNeeded.
const PaymentsPaymentVerificationNeededTypeID = 0xd8411139

// Encode implements bin.Encoder.
func (p *PaymentsPaymentVerificationNeeded) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode payments.paymentVerificationNeeded#d8411139 as nil")
	}
	b.PutID(PaymentsPaymentVerificationNeededTypeID)
	b.PutString(p.URL)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentsPaymentVerificationNeeded) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode payments.paymentVerificationNeeded#d8411139 to nil")
	}
	if err := b.ConsumeID(PaymentsPaymentVerificationNeededTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.paymentVerificationNeeded#d8411139: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentVerificationNeeded#d8411139: field url: %w", err)
		}
		p.URL = value
	}
	return nil
}

// construct implements constructor of PaymentsPaymentResultClass.
func (p PaymentsPaymentVerificationNeeded) construct() PaymentsPaymentResultClass { return &p }

// Ensuring interfaces in compile-time for PaymentsPaymentVerificationNeeded.
var (
	_ bin.Encoder = &PaymentsPaymentVerificationNeeded{}
	_ bin.Decoder = &PaymentsPaymentVerificationNeeded{}

	_ PaymentsPaymentResultClass = &PaymentsPaymentVerificationNeeded{}
)

// PaymentsPaymentResultClass represents payments.PaymentResult generic type.
//
// Example:
//  g, err := DecodePaymentsPaymentResult(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *PaymentsPaymentResult: // payments.paymentResult#4e5f810d
//  case *PaymentsPaymentVerificationNeeded: // payments.paymentVerificationNeeded#d8411139
//  default: panic(v)
//  }
type PaymentsPaymentResultClass interface {
	bin.Encoder
	bin.Decoder
	construct() PaymentsPaymentResultClass
}

// DecodePaymentsPaymentResult implements binary de-serialization for PaymentsPaymentResultClass.
func DecodePaymentsPaymentResult(buf *bin.Buffer) (PaymentsPaymentResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PaymentsPaymentResultTypeID:
		// Decoding payments.paymentResult#4e5f810d.
		v := PaymentsPaymentResult{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentsPaymentResultClass: %w", err)
		}
		return &v, nil
	case PaymentsPaymentVerificationNeededTypeID:
		// Decoding payments.paymentVerificationNeeded#d8411139.
		v := PaymentsPaymentVerificationNeeded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentsPaymentResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PaymentsPaymentResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// PaymentsPaymentResult boxes the PaymentsPaymentResultClass providing a helper.
type PaymentsPaymentResultBox struct {
	PaymentsPaymentResult PaymentsPaymentResultClass
}

// Decode implements bin.Decoder for PaymentsPaymentResultBox.
func (b *PaymentsPaymentResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PaymentsPaymentResultBox to nil")
	}
	v, err := DecodePaymentsPaymentResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PaymentsPaymentResult = v
	return nil
}

// Encode implements bin.Encode for PaymentsPaymentResultBox.
func (b *PaymentsPaymentResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PaymentsPaymentResult == nil {
		return fmt.Errorf("unable to encode PaymentsPaymentResultClass as nil")
	}
	return b.PaymentsPaymentResult.Encode(buf)
}