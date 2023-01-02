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

// PaymentReceipt represents TL type `paymentReceipt#e819e6c1`.
type PaymentReceipt struct {
	// Product title
	Title string
	// Contains information about a successful payment
	Description FormattedText
	// Product photo; may be null
	Photo Photo
	// Point in time (Unix timestamp) when the payment was made
	Date int32
	// User identifier of the seller bot
	SellerBotUserID int64
	// User identifier of the payment provider bot
	PaymentProviderUserID int64
	// Information about the invoice
	Invoice Invoice
	// Order information; may be null
	OrderInfo OrderInfo
	// Chosen shipping option; may be null
	ShippingOption ShippingOption
	// Title of the saved credentials chosen by the buyer
	CredentialsTitle string
	// The amount of tip chosen by the buyer in the smallest units of the currency
	TipAmount int64
}

// PaymentReceiptTypeID is TL type id of PaymentReceipt.
const PaymentReceiptTypeID = 0xe819e6c1

// Ensuring interfaces in compile-time for PaymentReceipt.
var (
	_ bin.Encoder     = &PaymentReceipt{}
	_ bin.Decoder     = &PaymentReceipt{}
	_ bin.BareEncoder = &PaymentReceipt{}
	_ bin.BareDecoder = &PaymentReceipt{}
)

func (p *PaymentReceipt) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Title == "") {
		return false
	}
	if !(p.Description.Zero()) {
		return false
	}
	if !(p.Photo.Zero()) {
		return false
	}
	if !(p.Date == 0) {
		return false
	}
	if !(p.SellerBotUserID == 0) {
		return false
	}
	if !(p.PaymentProviderUserID == 0) {
		return false
	}
	if !(p.Invoice.Zero()) {
		return false
	}
	if !(p.OrderInfo.Zero()) {
		return false
	}
	if !(p.ShippingOption.Zero()) {
		return false
	}
	if !(p.CredentialsTitle == "") {
		return false
	}
	if !(p.TipAmount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentReceipt) String() string {
	if p == nil {
		return "PaymentReceipt(nil)"
	}
	type Alias PaymentReceipt
	return fmt.Sprintf("PaymentReceipt%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentReceipt) TypeID() uint32 {
	return PaymentReceiptTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentReceipt) TypeName() string {
	return "paymentReceipt"
}

// TypeInfo returns info about TL type.
func (p *PaymentReceipt) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentReceipt",
		ID:   PaymentReceiptTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "SellerBotUserID",
			SchemaName: "seller_bot_user_id",
		},
		{
			Name:       "PaymentProviderUserID",
			SchemaName: "payment_provider_user_id",
		},
		{
			Name:       "Invoice",
			SchemaName: "invoice",
		},
		{
			Name:       "OrderInfo",
			SchemaName: "order_info",
		},
		{
			Name:       "ShippingOption",
			SchemaName: "shipping_option",
		},
		{
			Name:       "CredentialsTitle",
			SchemaName: "credentials_title",
		},
		{
			Name:       "TipAmount",
			SchemaName: "tip_amount",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentReceipt) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentReceipt#e819e6c1 as nil")
	}
	b.PutID(PaymentReceiptTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentReceipt) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentReceipt#e819e6c1 as nil")
	}
	b.PutString(p.Title)
	if err := p.Description.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#e819e6c1: field description: %w", err)
	}
	if err := p.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#e819e6c1: field photo: %w", err)
	}
	b.PutInt32(p.Date)
	b.PutInt53(p.SellerBotUserID)
	b.PutInt53(p.PaymentProviderUserID)
	if err := p.Invoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#e819e6c1: field invoice: %w", err)
	}
	if err := p.OrderInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#e819e6c1: field order_info: %w", err)
	}
	if err := p.ShippingOption.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#e819e6c1: field shipping_option: %w", err)
	}
	b.PutString(p.CredentialsTitle)
	b.PutInt53(p.TipAmount)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentReceipt) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentReceipt#e819e6c1 to nil")
	}
	if err := b.ConsumeID(PaymentReceiptTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentReceipt) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentReceipt#e819e6c1 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field title: %w", err)
		}
		p.Title = value
	}
	{
		if err := p.Description.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field description: %w", err)
		}
	}
	{
		if err := p.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field photo: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field date: %w", err)
		}
		p.Date = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field seller_bot_user_id: %w", err)
		}
		p.SellerBotUserID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field payment_provider_user_id: %w", err)
		}
		p.PaymentProviderUserID = value
	}
	{
		if err := p.Invoice.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field invoice: %w", err)
		}
	}
	{
		if err := p.OrderInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field order_info: %w", err)
		}
	}
	{
		if err := p.ShippingOption.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field shipping_option: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field credentials_title: %w", err)
		}
		p.CredentialsTitle = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field tip_amount: %w", err)
		}
		p.TipAmount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PaymentReceipt) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentReceipt#e819e6c1 as nil")
	}
	b.ObjStart()
	b.PutID("paymentReceipt")
	b.Comma()
	b.FieldStart("title")
	b.PutString(p.Title)
	b.Comma()
	b.FieldStart("description")
	if err := p.Description.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#e819e6c1: field description: %w", err)
	}
	b.Comma()
	b.FieldStart("photo")
	if err := p.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#e819e6c1: field photo: %w", err)
	}
	b.Comma()
	b.FieldStart("date")
	b.PutInt32(p.Date)
	b.Comma()
	b.FieldStart("seller_bot_user_id")
	b.PutInt53(p.SellerBotUserID)
	b.Comma()
	b.FieldStart("payment_provider_user_id")
	b.PutInt53(p.PaymentProviderUserID)
	b.Comma()
	b.FieldStart("invoice")
	if err := p.Invoice.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#e819e6c1: field invoice: %w", err)
	}
	b.Comma()
	b.FieldStart("order_info")
	if err := p.OrderInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#e819e6c1: field order_info: %w", err)
	}
	b.Comma()
	b.FieldStart("shipping_option")
	if err := p.ShippingOption.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#e819e6c1: field shipping_option: %w", err)
	}
	b.Comma()
	b.FieldStart("credentials_title")
	b.PutString(p.CredentialsTitle)
	b.Comma()
	b.FieldStart("tip_amount")
	b.PutInt53(p.TipAmount)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PaymentReceipt) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentReceipt#e819e6c1 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("paymentReceipt"); err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: %w", err)
			}
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field title: %w", err)
			}
			p.Title = value
		case "description":
			if err := p.Description.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field description: %w", err)
			}
		case "photo":
			if err := p.Photo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field photo: %w", err)
			}
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field date: %w", err)
			}
			p.Date = value
		case "seller_bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field seller_bot_user_id: %w", err)
			}
			p.SellerBotUserID = value
		case "payment_provider_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field payment_provider_user_id: %w", err)
			}
			p.PaymentProviderUserID = value
		case "invoice":
			if err := p.Invoice.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field invoice: %w", err)
			}
		case "order_info":
			if err := p.OrderInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field order_info: %w", err)
			}
		case "shipping_option":
			if err := p.ShippingOption.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field shipping_option: %w", err)
			}
		case "credentials_title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field credentials_title: %w", err)
			}
			p.CredentialsTitle = value
		case "tip_amount":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode paymentReceipt#e819e6c1: field tip_amount: %w", err)
			}
			p.TipAmount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTitle returns value of Title field.
func (p *PaymentReceipt) GetTitle() (value string) {
	if p == nil {
		return
	}
	return p.Title
}

// GetDescription returns value of Description field.
func (p *PaymentReceipt) GetDescription() (value FormattedText) {
	if p == nil {
		return
	}
	return p.Description
}

// GetPhoto returns value of Photo field.
func (p *PaymentReceipt) GetPhoto() (value Photo) {
	if p == nil {
		return
	}
	return p.Photo
}

// GetDate returns value of Date field.
func (p *PaymentReceipt) GetDate() (value int32) {
	if p == nil {
		return
	}
	return p.Date
}

// GetSellerBotUserID returns value of SellerBotUserID field.
func (p *PaymentReceipt) GetSellerBotUserID() (value int64) {
	if p == nil {
		return
	}
	return p.SellerBotUserID
}

// GetPaymentProviderUserID returns value of PaymentProviderUserID field.
func (p *PaymentReceipt) GetPaymentProviderUserID() (value int64) {
	if p == nil {
		return
	}
	return p.PaymentProviderUserID
}

// GetInvoice returns value of Invoice field.
func (p *PaymentReceipt) GetInvoice() (value Invoice) {
	if p == nil {
		return
	}
	return p.Invoice
}

// GetOrderInfo returns value of OrderInfo field.
func (p *PaymentReceipt) GetOrderInfo() (value OrderInfo) {
	if p == nil {
		return
	}
	return p.OrderInfo
}

// GetShippingOption returns value of ShippingOption field.
func (p *PaymentReceipt) GetShippingOption() (value ShippingOption) {
	if p == nil {
		return
	}
	return p.ShippingOption
}

// GetCredentialsTitle returns value of CredentialsTitle field.
func (p *PaymentReceipt) GetCredentialsTitle() (value string) {
	if p == nil {
		return
	}
	return p.CredentialsTitle
}

// GetTipAmount returns value of TipAmount field.
func (p *PaymentReceipt) GetTipAmount() (value int64) {
	if p == nil {
		return
	}
	return p.TipAmount
}
