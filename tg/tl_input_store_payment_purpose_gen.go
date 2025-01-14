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

// InputStorePaymentPremiumSubscription represents TL type `inputStorePaymentPremiumSubscription#a6751e66`.
// Info about a Telegram Premium purchase
//
// See https://core.telegram.org/constructor/inputStorePaymentPremiumSubscription for reference.
type InputStorePaymentPremiumSubscription struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Pass true if this is a restore of a Telegram Premium purchase; only for the App Store
	Restore bool
}

// InputStorePaymentPremiumSubscriptionTypeID is TL type id of InputStorePaymentPremiumSubscription.
const InputStorePaymentPremiumSubscriptionTypeID = 0xa6751e66

// construct implements constructor of InputStorePaymentPurposeClass.
func (i InputStorePaymentPremiumSubscription) construct() InputStorePaymentPurposeClass { return &i }

// Ensuring interfaces in compile-time for InputStorePaymentPremiumSubscription.
var (
	_ bin.Encoder     = &InputStorePaymentPremiumSubscription{}
	_ bin.Decoder     = &InputStorePaymentPremiumSubscription{}
	_ bin.BareEncoder = &InputStorePaymentPremiumSubscription{}
	_ bin.BareDecoder = &InputStorePaymentPremiumSubscription{}

	_ InputStorePaymentPurposeClass = &InputStorePaymentPremiumSubscription{}
)

func (i *InputStorePaymentPremiumSubscription) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Restore == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStorePaymentPremiumSubscription) String() string {
	if i == nil {
		return "InputStorePaymentPremiumSubscription(nil)"
	}
	type Alias InputStorePaymentPremiumSubscription
	return fmt.Sprintf("InputStorePaymentPremiumSubscription%+v", Alias(*i))
}

// FillFrom fills InputStorePaymentPremiumSubscription from given interface.
func (i *InputStorePaymentPremiumSubscription) FillFrom(from interface {
	GetRestore() (value bool)
}) {
	i.Restore = from.GetRestore()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStorePaymentPremiumSubscription) TypeID() uint32 {
	return InputStorePaymentPremiumSubscriptionTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStorePaymentPremiumSubscription) TypeName() string {
	return "inputStorePaymentPremiumSubscription"
}

// TypeInfo returns info about TL type.
func (i *InputStorePaymentPremiumSubscription) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStorePaymentPremiumSubscription",
		ID:   InputStorePaymentPremiumSubscriptionTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Restore",
			SchemaName: "restore",
			Null:       !i.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *InputStorePaymentPremiumSubscription) SetFlags() {
	if !(i.Restore == false) {
		i.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (i *InputStorePaymentPremiumSubscription) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStorePaymentPremiumSubscription#a6751e66 as nil")
	}
	b.PutID(InputStorePaymentPremiumSubscriptionTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStorePaymentPremiumSubscription) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStorePaymentPremiumSubscription#a6751e66 as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStorePaymentPremiumSubscription#a6751e66: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStorePaymentPremiumSubscription) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStorePaymentPremiumSubscription#a6751e66 to nil")
	}
	if err := b.ConsumeID(InputStorePaymentPremiumSubscriptionTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStorePaymentPremiumSubscription#a6751e66: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStorePaymentPremiumSubscription) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStorePaymentPremiumSubscription#a6751e66 to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputStorePaymentPremiumSubscription#a6751e66: field flags: %w", err)
		}
	}
	i.Restore = i.Flags.Has(0)
	return nil
}

// SetRestore sets value of Restore conditional field.
func (i *InputStorePaymentPremiumSubscription) SetRestore(value bool) {
	if value {
		i.Flags.Set(0)
		i.Restore = true
	} else {
		i.Flags.Unset(0)
		i.Restore = false
	}
}

// GetRestore returns value of Restore conditional field.
func (i *InputStorePaymentPremiumSubscription) GetRestore() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(0)
}

// InputStorePaymentGiftPremium represents TL type `inputStorePaymentGiftPremium#616f7fe8`.
// Info about a gifted Telegram Premium purchase
//
// See https://core.telegram.org/constructor/inputStorePaymentGiftPremium for reference.
type InputStorePaymentGiftPremium struct {
	// The user to which the Telegram Premium subscription was gifted
	UserID InputUserClass
	// Three-letter ISO 4217 currency¹ code
	//
	// Links:
	//  1) https://core.telegram.org/bots/payments#supported-currencies
	Currency string
	// Price of the product in the smallest units of the currency (integer, not float/double)
	// For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in
	// currencies.json¹, it shows the number of digits past the decimal point for each
	// currency (2 for the majority of currencies).
	//
	// Links:
	//  1) https://core.telegram.org/bots/payments/currencies.json
	Amount int64
}

// InputStorePaymentGiftPremiumTypeID is TL type id of InputStorePaymentGiftPremium.
const InputStorePaymentGiftPremiumTypeID = 0x616f7fe8

// construct implements constructor of InputStorePaymentPurposeClass.
func (i InputStorePaymentGiftPremium) construct() InputStorePaymentPurposeClass { return &i }

// Ensuring interfaces in compile-time for InputStorePaymentGiftPremium.
var (
	_ bin.Encoder     = &InputStorePaymentGiftPremium{}
	_ bin.Decoder     = &InputStorePaymentGiftPremium{}
	_ bin.BareEncoder = &InputStorePaymentGiftPremium{}
	_ bin.BareDecoder = &InputStorePaymentGiftPremium{}

	_ InputStorePaymentPurposeClass = &InputStorePaymentGiftPremium{}
)

func (i *InputStorePaymentGiftPremium) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.UserID == nil) {
		return false
	}
	if !(i.Currency == "") {
		return false
	}
	if !(i.Amount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStorePaymentGiftPremium) String() string {
	if i == nil {
		return "InputStorePaymentGiftPremium(nil)"
	}
	type Alias InputStorePaymentGiftPremium
	return fmt.Sprintf("InputStorePaymentGiftPremium%+v", Alias(*i))
}

// FillFrom fills InputStorePaymentGiftPremium from given interface.
func (i *InputStorePaymentGiftPremium) FillFrom(from interface {
	GetUserID() (value InputUserClass)
	GetCurrency() (value string)
	GetAmount() (value int64)
}) {
	i.UserID = from.GetUserID()
	i.Currency = from.GetCurrency()
	i.Amount = from.GetAmount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStorePaymentGiftPremium) TypeID() uint32 {
	return InputStorePaymentGiftPremiumTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStorePaymentGiftPremium) TypeName() string {
	return "inputStorePaymentGiftPremium"
}

// TypeInfo returns info about TL type.
func (i *InputStorePaymentGiftPremium) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStorePaymentGiftPremium",
		ID:   InputStorePaymentGiftPremiumTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Currency",
			SchemaName: "currency",
		},
		{
			Name:       "Amount",
			SchemaName: "amount",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStorePaymentGiftPremium) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStorePaymentGiftPremium#616f7fe8 as nil")
	}
	b.PutID(InputStorePaymentGiftPremiumTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStorePaymentGiftPremium) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStorePaymentGiftPremium#616f7fe8 as nil")
	}
	if i.UserID == nil {
		return fmt.Errorf("unable to encode inputStorePaymentGiftPremium#616f7fe8: field user_id is nil")
	}
	if err := i.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStorePaymentGiftPremium#616f7fe8: field user_id: %w", err)
	}
	b.PutString(i.Currency)
	b.PutLong(i.Amount)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStorePaymentGiftPremium) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStorePaymentGiftPremium#616f7fe8 to nil")
	}
	if err := b.ConsumeID(InputStorePaymentGiftPremiumTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStorePaymentGiftPremium#616f7fe8: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStorePaymentGiftPremium) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStorePaymentGiftPremium#616f7fe8 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStorePaymentGiftPremium#616f7fe8: field user_id: %w", err)
		}
		i.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStorePaymentGiftPremium#616f7fe8: field currency: %w", err)
		}
		i.Currency = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputStorePaymentGiftPremium#616f7fe8: field amount: %w", err)
		}
		i.Amount = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (i *InputStorePaymentGiftPremium) GetUserID() (value InputUserClass) {
	if i == nil {
		return
	}
	return i.UserID
}

// GetCurrency returns value of Currency field.
func (i *InputStorePaymentGiftPremium) GetCurrency() (value string) {
	if i == nil {
		return
	}
	return i.Currency
}

// GetAmount returns value of Amount field.
func (i *InputStorePaymentGiftPremium) GetAmount() (value int64) {
	if i == nil {
		return
	}
	return i.Amount
}

// InputStorePaymentPurposeClassName is schema name of InputStorePaymentPurposeClass.
const InputStorePaymentPurposeClassName = "InputStorePaymentPurpose"

// InputStorePaymentPurposeClass represents InputStorePaymentPurpose generic type.
//
// See https://core.telegram.org/type/InputStorePaymentPurpose for reference.
//
// Example:
//
//	g, err := tg.DecodeInputStorePaymentPurpose(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputStorePaymentPremiumSubscription: // inputStorePaymentPremiumSubscription#a6751e66
//	case *tg.InputStorePaymentGiftPremium: // inputStorePaymentGiftPremium#616f7fe8
//	default: panic(v)
//	}
type InputStorePaymentPurposeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputStorePaymentPurposeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputStorePaymentPurpose implements binary de-serialization for InputStorePaymentPurposeClass.
func DecodeInputStorePaymentPurpose(buf *bin.Buffer) (InputStorePaymentPurposeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputStorePaymentPremiumSubscriptionTypeID:
		// Decoding inputStorePaymentPremiumSubscription#a6751e66.
		v := InputStorePaymentPremiumSubscription{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStorePaymentPurposeClass: %w", err)
		}
		return &v, nil
	case InputStorePaymentGiftPremiumTypeID:
		// Decoding inputStorePaymentGiftPremium#616f7fe8.
		v := InputStorePaymentGiftPremium{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStorePaymentPurposeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStorePaymentPurposeClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputStorePaymentPurpose boxes the InputStorePaymentPurposeClass providing a helper.
type InputStorePaymentPurposeBox struct {
	InputStorePaymentPurpose InputStorePaymentPurposeClass
}

// Decode implements bin.Decoder for InputStorePaymentPurposeBox.
func (b *InputStorePaymentPurposeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStorePaymentPurposeBox to nil")
	}
	v, err := DecodeInputStorePaymentPurpose(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputStorePaymentPurpose = v
	return nil
}

// Encode implements bin.Encode for InputStorePaymentPurposeBox.
func (b *InputStorePaymentPurposeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputStorePaymentPurpose == nil {
		return fmt.Errorf("unable to encode InputStorePaymentPurposeClass as nil")
	}
	return b.InputStorePaymentPurpose.Encode(buf)
}
