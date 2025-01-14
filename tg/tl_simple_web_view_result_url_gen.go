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

// SimpleWebViewResultURL represents TL type `simpleWebViewResultUrl#882f76bb`.
// Contains the webview URL with appropriate theme parameters added
//
// See https://core.telegram.org/constructor/simpleWebViewResultUrl for reference.
type SimpleWebViewResultURL struct {
	// URL
	URL string
}

// SimpleWebViewResultURLTypeID is TL type id of SimpleWebViewResultURL.
const SimpleWebViewResultURLTypeID = 0x882f76bb

// Ensuring interfaces in compile-time for SimpleWebViewResultURL.
var (
	_ bin.Encoder     = &SimpleWebViewResultURL{}
	_ bin.Decoder     = &SimpleWebViewResultURL{}
	_ bin.BareEncoder = &SimpleWebViewResultURL{}
	_ bin.BareDecoder = &SimpleWebViewResultURL{}
)

func (s *SimpleWebViewResultURL) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SimpleWebViewResultURL) String() string {
	if s == nil {
		return "SimpleWebViewResultURL(nil)"
	}
	type Alias SimpleWebViewResultURL
	return fmt.Sprintf("SimpleWebViewResultURL%+v", Alias(*s))
}

// FillFrom fills SimpleWebViewResultURL from given interface.
func (s *SimpleWebViewResultURL) FillFrom(from interface {
	GetURL() (value string)
}) {
	s.URL = from.GetURL()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SimpleWebViewResultURL) TypeID() uint32 {
	return SimpleWebViewResultURLTypeID
}

// TypeName returns name of type in TL schema.
func (*SimpleWebViewResultURL) TypeName() string {
	return "simpleWebViewResultUrl"
}

// TypeInfo returns info about TL type.
func (s *SimpleWebViewResultURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "simpleWebViewResultUrl",
		ID:   SimpleWebViewResultURLTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SimpleWebViewResultURL) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode simpleWebViewResultUrl#882f76bb as nil")
	}
	b.PutID(SimpleWebViewResultURLTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SimpleWebViewResultURL) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode simpleWebViewResultUrl#882f76bb as nil")
	}
	b.PutString(s.URL)
	return nil
}

// Decode implements bin.Decoder.
func (s *SimpleWebViewResultURL) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode simpleWebViewResultUrl#882f76bb to nil")
	}
	if err := b.ConsumeID(SimpleWebViewResultURLTypeID); err != nil {
		return fmt.Errorf("unable to decode simpleWebViewResultUrl#882f76bb: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SimpleWebViewResultURL) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode simpleWebViewResultUrl#882f76bb to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode simpleWebViewResultUrl#882f76bb: field url: %w", err)
		}
		s.URL = value
	}
	return nil
}

// GetURL returns value of URL field.
func (s *SimpleWebViewResultURL) GetURL() (value string) {
	if s == nil {
		return
	}
	return s.URL
}
