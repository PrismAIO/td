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

// MessagesTranslateNoResult represents TL type `messages.translateNoResult#67ca4737`.
// No translation is available
//
// See https://core.telegram.org/constructor/messages.translateNoResult for reference.
type MessagesTranslateNoResult struct {
}

// MessagesTranslateNoResultTypeID is TL type id of MessagesTranslateNoResult.
const MessagesTranslateNoResultTypeID = 0x67ca4737

// construct implements constructor of MessagesTranslatedTextClass.
func (t MessagesTranslateNoResult) construct() MessagesTranslatedTextClass { return &t }

// Ensuring interfaces in compile-time for MessagesTranslateNoResult.
var (
	_ bin.Encoder     = &MessagesTranslateNoResult{}
	_ bin.Decoder     = &MessagesTranslateNoResult{}
	_ bin.BareEncoder = &MessagesTranslateNoResult{}
	_ bin.BareDecoder = &MessagesTranslateNoResult{}

	_ MessagesTranslatedTextClass = &MessagesTranslateNoResult{}
)

func (t *MessagesTranslateNoResult) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesTranslateNoResult) String() string {
	if t == nil {
		return "MessagesTranslateNoResult(nil)"
	}
	type Alias MessagesTranslateNoResult
	return fmt.Sprintf("MessagesTranslateNoResult%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesTranslateNoResult) TypeID() uint32 {
	return MessagesTranslateNoResultTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesTranslateNoResult) TypeName() string {
	return "messages.translateNoResult"
}

// TypeInfo returns info about TL type.
func (t *MessagesTranslateNoResult) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.translateNoResult",
		ID:   MessagesTranslateNoResultTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *MessagesTranslateNoResult) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.translateNoResult#67ca4737 as nil")
	}
	b.PutID(MessagesTranslateNoResultTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesTranslateNoResult) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.translateNoResult#67ca4737 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesTranslateNoResult) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.translateNoResult#67ca4737 to nil")
	}
	if err := b.ConsumeID(MessagesTranslateNoResultTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.translateNoResult#67ca4737: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesTranslateNoResult) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.translateNoResult#67ca4737 to nil")
	}
	return nil
}

// MessagesTranslateResultText represents TL type `messages.translateResultText#a214f7d0`.
// Translated text
//
// See https://core.telegram.org/constructor/messages.translateResultText for reference.
type MessagesTranslateResultText struct {
	// Translated text
	Text string
}

// MessagesTranslateResultTextTypeID is TL type id of MessagesTranslateResultText.
const MessagesTranslateResultTextTypeID = 0xa214f7d0

// construct implements constructor of MessagesTranslatedTextClass.
func (t MessagesTranslateResultText) construct() MessagesTranslatedTextClass { return &t }

// Ensuring interfaces in compile-time for MessagesTranslateResultText.
var (
	_ bin.Encoder     = &MessagesTranslateResultText{}
	_ bin.Decoder     = &MessagesTranslateResultText{}
	_ bin.BareEncoder = &MessagesTranslateResultText{}
	_ bin.BareDecoder = &MessagesTranslateResultText{}

	_ MessagesTranslatedTextClass = &MessagesTranslateResultText{}
)

func (t *MessagesTranslateResultText) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesTranslateResultText) String() string {
	if t == nil {
		return "MessagesTranslateResultText(nil)"
	}
	type Alias MessagesTranslateResultText
	return fmt.Sprintf("MessagesTranslateResultText%+v", Alias(*t))
}

// FillFrom fills MessagesTranslateResultText from given interface.
func (t *MessagesTranslateResultText) FillFrom(from interface {
	GetText() (value string)
}) {
	t.Text = from.GetText()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesTranslateResultText) TypeID() uint32 {
	return MessagesTranslateResultTextTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesTranslateResultText) TypeName() string {
	return "messages.translateResultText"
}

// TypeInfo returns info about TL type.
func (t *MessagesTranslateResultText) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.translateResultText",
		ID:   MessagesTranslateResultTextTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *MessagesTranslateResultText) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.translateResultText#a214f7d0 as nil")
	}
	b.PutID(MessagesTranslateResultTextTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesTranslateResultText) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.translateResultText#a214f7d0 as nil")
	}
	b.PutString(t.Text)
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesTranslateResultText) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.translateResultText#a214f7d0 to nil")
	}
	if err := b.ConsumeID(MessagesTranslateResultTextTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.translateResultText#a214f7d0: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesTranslateResultText) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.translateResultText#a214f7d0 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.translateResultText#a214f7d0: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// GetText returns value of Text field.
func (t *MessagesTranslateResultText) GetText() (value string) {
	if t == nil {
		return
	}
	return t.Text
}

// MessagesTranslatedTextClassName is schema name of MessagesTranslatedTextClass.
const MessagesTranslatedTextClassName = "messages.TranslatedText"

// MessagesTranslatedTextClass represents messages.TranslatedText generic type.
//
// See https://core.telegram.org/type/messages.TranslatedText for reference.
//
// Example:
//
//	g, err := tg.DecodeMessagesTranslatedText(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessagesTranslateNoResult: // messages.translateNoResult#67ca4737
//	case *tg.MessagesTranslateResultText: // messages.translateResultText#a214f7d0
//	default: panic(v)
//	}
type MessagesTranslatedTextClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesTranslatedTextClass

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

// DecodeMessagesTranslatedText implements binary de-serialization for MessagesTranslatedTextClass.
func DecodeMessagesTranslatedText(buf *bin.Buffer) (MessagesTranslatedTextClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesTranslateNoResultTypeID:
		// Decoding messages.translateNoResult#67ca4737.
		v := MessagesTranslateNoResult{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesTranslatedTextClass: %w", err)
		}
		return &v, nil
	case MessagesTranslateResultTextTypeID:
		// Decoding messages.translateResultText#a214f7d0.
		v := MessagesTranslateResultText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesTranslatedTextClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesTranslatedTextClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesTranslatedText boxes the MessagesTranslatedTextClass providing a helper.
type MessagesTranslatedTextBox struct {
	TranslatedText MessagesTranslatedTextClass
}

// Decode implements bin.Decoder for MessagesTranslatedTextBox.
func (b *MessagesTranslatedTextBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesTranslatedTextBox to nil")
	}
	v, err := DecodeMessagesTranslatedText(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TranslatedText = v
	return nil
}

// Encode implements bin.Encode for MessagesTranslatedTextBox.
func (b *MessagesTranslatedTextBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.TranslatedText == nil {
		return fmt.Errorf("unable to encode MessagesTranslatedTextClass as nil")
	}
	return b.TranslatedText.Encode(buf)
}
