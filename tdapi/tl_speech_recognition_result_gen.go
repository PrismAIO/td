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

// SpeechRecognitionResultPending represents TL type `speechRecognitionResultPending#9ebc8e00`.
type SpeechRecognitionResultPending struct {
	// Partially recognized text
	PartialText string
}

// SpeechRecognitionResultPendingTypeID is TL type id of SpeechRecognitionResultPending.
const SpeechRecognitionResultPendingTypeID = 0x9ebc8e00

// construct implements constructor of SpeechRecognitionResultClass.
func (s SpeechRecognitionResultPending) construct() SpeechRecognitionResultClass { return &s }

// Ensuring interfaces in compile-time for SpeechRecognitionResultPending.
var (
	_ bin.Encoder     = &SpeechRecognitionResultPending{}
	_ bin.Decoder     = &SpeechRecognitionResultPending{}
	_ bin.BareEncoder = &SpeechRecognitionResultPending{}
	_ bin.BareDecoder = &SpeechRecognitionResultPending{}

	_ SpeechRecognitionResultClass = &SpeechRecognitionResultPending{}
)

func (s *SpeechRecognitionResultPending) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PartialText == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SpeechRecognitionResultPending) String() string {
	if s == nil {
		return "SpeechRecognitionResultPending(nil)"
	}
	type Alias SpeechRecognitionResultPending
	return fmt.Sprintf("SpeechRecognitionResultPending%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SpeechRecognitionResultPending) TypeID() uint32 {
	return SpeechRecognitionResultPendingTypeID
}

// TypeName returns name of type in TL schema.
func (*SpeechRecognitionResultPending) TypeName() string {
	return "speechRecognitionResultPending"
}

// TypeInfo returns info about TL type.
func (s *SpeechRecognitionResultPending) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "speechRecognitionResultPending",
		ID:   SpeechRecognitionResultPendingTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PartialText",
			SchemaName: "partial_text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SpeechRecognitionResultPending) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode speechRecognitionResultPending#9ebc8e00 as nil")
	}
	b.PutID(SpeechRecognitionResultPendingTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SpeechRecognitionResultPending) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode speechRecognitionResultPending#9ebc8e00 as nil")
	}
	b.PutString(s.PartialText)
	return nil
}

// Decode implements bin.Decoder.
func (s *SpeechRecognitionResultPending) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode speechRecognitionResultPending#9ebc8e00 to nil")
	}
	if err := b.ConsumeID(SpeechRecognitionResultPendingTypeID); err != nil {
		return fmt.Errorf("unable to decode speechRecognitionResultPending#9ebc8e00: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SpeechRecognitionResultPending) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode speechRecognitionResultPending#9ebc8e00 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode speechRecognitionResultPending#9ebc8e00: field partial_text: %w", err)
		}
		s.PartialText = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SpeechRecognitionResultPending) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode speechRecognitionResultPending#9ebc8e00 as nil")
	}
	b.ObjStart()
	b.PutID("speechRecognitionResultPending")
	b.Comma()
	b.FieldStart("partial_text")
	b.PutString(s.PartialText)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SpeechRecognitionResultPending) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode speechRecognitionResultPending#9ebc8e00 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("speechRecognitionResultPending"); err != nil {
				return fmt.Errorf("unable to decode speechRecognitionResultPending#9ebc8e00: %w", err)
			}
		case "partial_text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode speechRecognitionResultPending#9ebc8e00: field partial_text: %w", err)
			}
			s.PartialText = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPartialText returns value of PartialText field.
func (s *SpeechRecognitionResultPending) GetPartialText() (value string) {
	if s == nil {
		return
	}
	return s.PartialText
}

// SpeechRecognitionResultText represents TL type `speechRecognitionResultText#80e681dd`.
type SpeechRecognitionResultText struct {
	// Recognized text
	Text string
}

// SpeechRecognitionResultTextTypeID is TL type id of SpeechRecognitionResultText.
const SpeechRecognitionResultTextTypeID = 0x80e681dd

// construct implements constructor of SpeechRecognitionResultClass.
func (s SpeechRecognitionResultText) construct() SpeechRecognitionResultClass { return &s }

// Ensuring interfaces in compile-time for SpeechRecognitionResultText.
var (
	_ bin.Encoder     = &SpeechRecognitionResultText{}
	_ bin.Decoder     = &SpeechRecognitionResultText{}
	_ bin.BareEncoder = &SpeechRecognitionResultText{}
	_ bin.BareDecoder = &SpeechRecognitionResultText{}

	_ SpeechRecognitionResultClass = &SpeechRecognitionResultText{}
)

func (s *SpeechRecognitionResultText) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SpeechRecognitionResultText) String() string {
	if s == nil {
		return "SpeechRecognitionResultText(nil)"
	}
	type Alias SpeechRecognitionResultText
	return fmt.Sprintf("SpeechRecognitionResultText%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SpeechRecognitionResultText) TypeID() uint32 {
	return SpeechRecognitionResultTextTypeID
}

// TypeName returns name of type in TL schema.
func (*SpeechRecognitionResultText) TypeName() string {
	return "speechRecognitionResultText"
}

// TypeInfo returns info about TL type.
func (s *SpeechRecognitionResultText) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "speechRecognitionResultText",
		ID:   SpeechRecognitionResultTextTypeID,
	}
	if s == nil {
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
func (s *SpeechRecognitionResultText) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode speechRecognitionResultText#80e681dd as nil")
	}
	b.PutID(SpeechRecognitionResultTextTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SpeechRecognitionResultText) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode speechRecognitionResultText#80e681dd as nil")
	}
	b.PutString(s.Text)
	return nil
}

// Decode implements bin.Decoder.
func (s *SpeechRecognitionResultText) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode speechRecognitionResultText#80e681dd to nil")
	}
	if err := b.ConsumeID(SpeechRecognitionResultTextTypeID); err != nil {
		return fmt.Errorf("unable to decode speechRecognitionResultText#80e681dd: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SpeechRecognitionResultText) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode speechRecognitionResultText#80e681dd to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode speechRecognitionResultText#80e681dd: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SpeechRecognitionResultText) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode speechRecognitionResultText#80e681dd as nil")
	}
	b.ObjStart()
	b.PutID("speechRecognitionResultText")
	b.Comma()
	b.FieldStart("text")
	b.PutString(s.Text)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SpeechRecognitionResultText) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode speechRecognitionResultText#80e681dd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("speechRecognitionResultText"); err != nil {
				return fmt.Errorf("unable to decode speechRecognitionResultText#80e681dd: %w", err)
			}
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode speechRecognitionResultText#80e681dd: field text: %w", err)
			}
			s.Text = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (s *SpeechRecognitionResultText) GetText() (value string) {
	if s == nil {
		return
	}
	return s.Text
}

// SpeechRecognitionResultError represents TL type `speechRecognitionResultError#9d243fc`.
type SpeechRecognitionResultError struct {
	// Received error
	Error Error
}

// SpeechRecognitionResultErrorTypeID is TL type id of SpeechRecognitionResultError.
const SpeechRecognitionResultErrorTypeID = 0x9d243fc

// construct implements constructor of SpeechRecognitionResultClass.
func (s SpeechRecognitionResultError) construct() SpeechRecognitionResultClass { return &s }

// Ensuring interfaces in compile-time for SpeechRecognitionResultError.
var (
	_ bin.Encoder     = &SpeechRecognitionResultError{}
	_ bin.Decoder     = &SpeechRecognitionResultError{}
	_ bin.BareEncoder = &SpeechRecognitionResultError{}
	_ bin.BareDecoder = &SpeechRecognitionResultError{}

	_ SpeechRecognitionResultClass = &SpeechRecognitionResultError{}
)

func (s *SpeechRecognitionResultError) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Error.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SpeechRecognitionResultError) String() string {
	if s == nil {
		return "SpeechRecognitionResultError(nil)"
	}
	type Alias SpeechRecognitionResultError
	return fmt.Sprintf("SpeechRecognitionResultError%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SpeechRecognitionResultError) TypeID() uint32 {
	return SpeechRecognitionResultErrorTypeID
}

// TypeName returns name of type in TL schema.
func (*SpeechRecognitionResultError) TypeName() string {
	return "speechRecognitionResultError"
}

// TypeInfo returns info about TL type.
func (s *SpeechRecognitionResultError) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "speechRecognitionResultError",
		ID:   SpeechRecognitionResultErrorTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Error",
			SchemaName: "error",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SpeechRecognitionResultError) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode speechRecognitionResultError#9d243fc as nil")
	}
	b.PutID(SpeechRecognitionResultErrorTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SpeechRecognitionResultError) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode speechRecognitionResultError#9d243fc as nil")
	}
	if err := s.Error.Encode(b); err != nil {
		return fmt.Errorf("unable to encode speechRecognitionResultError#9d243fc: field error: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SpeechRecognitionResultError) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode speechRecognitionResultError#9d243fc to nil")
	}
	if err := b.ConsumeID(SpeechRecognitionResultErrorTypeID); err != nil {
		return fmt.Errorf("unable to decode speechRecognitionResultError#9d243fc: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SpeechRecognitionResultError) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode speechRecognitionResultError#9d243fc to nil")
	}
	{
		if err := s.Error.Decode(b); err != nil {
			return fmt.Errorf("unable to decode speechRecognitionResultError#9d243fc: field error: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SpeechRecognitionResultError) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode speechRecognitionResultError#9d243fc as nil")
	}
	b.ObjStart()
	b.PutID("speechRecognitionResultError")
	b.Comma()
	b.FieldStart("error")
	if err := s.Error.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode speechRecognitionResultError#9d243fc: field error: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SpeechRecognitionResultError) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode speechRecognitionResultError#9d243fc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("speechRecognitionResultError"); err != nil {
				return fmt.Errorf("unable to decode speechRecognitionResultError#9d243fc: %w", err)
			}
		case "error":
			if err := s.Error.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode speechRecognitionResultError#9d243fc: field error: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetError returns value of Error field.
func (s *SpeechRecognitionResultError) GetError() (value Error) {
	if s == nil {
		return
	}
	return s.Error
}

// SpeechRecognitionResultClassName is schema name of SpeechRecognitionResultClass.
const SpeechRecognitionResultClassName = "SpeechRecognitionResult"

// SpeechRecognitionResultClass represents SpeechRecognitionResult generic type.
//
// Example:
//
//	g, err := tdapi.DecodeSpeechRecognitionResult(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.SpeechRecognitionResultPending: // speechRecognitionResultPending#9ebc8e00
//	case *tdapi.SpeechRecognitionResultText: // speechRecognitionResultText#80e681dd
//	case *tdapi.SpeechRecognitionResultError: // speechRecognitionResultError#9d243fc
//	default: panic(v)
//	}
type SpeechRecognitionResultClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() SpeechRecognitionResultClass

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

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeSpeechRecognitionResult implements binary de-serialization for SpeechRecognitionResultClass.
func DecodeSpeechRecognitionResult(buf *bin.Buffer) (SpeechRecognitionResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SpeechRecognitionResultPendingTypeID:
		// Decoding speechRecognitionResultPending#9ebc8e00.
		v := SpeechRecognitionResultPending{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SpeechRecognitionResultClass: %w", err)
		}
		return &v, nil
	case SpeechRecognitionResultTextTypeID:
		// Decoding speechRecognitionResultText#80e681dd.
		v := SpeechRecognitionResultText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SpeechRecognitionResultClass: %w", err)
		}
		return &v, nil
	case SpeechRecognitionResultErrorTypeID:
		// Decoding speechRecognitionResultError#9d243fc.
		v := SpeechRecognitionResultError{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SpeechRecognitionResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SpeechRecognitionResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONSpeechRecognitionResult implements binary de-serialization for SpeechRecognitionResultClass.
func DecodeTDLibJSONSpeechRecognitionResult(buf tdjson.Decoder) (SpeechRecognitionResultClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "speechRecognitionResultPending":
		// Decoding speechRecognitionResultPending#9ebc8e00.
		v := SpeechRecognitionResultPending{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SpeechRecognitionResultClass: %w", err)
		}
		return &v, nil
	case "speechRecognitionResultText":
		// Decoding speechRecognitionResultText#80e681dd.
		v := SpeechRecognitionResultText{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SpeechRecognitionResultClass: %w", err)
		}
		return &v, nil
	case "speechRecognitionResultError":
		// Decoding speechRecognitionResultError#9d243fc.
		v := SpeechRecognitionResultError{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SpeechRecognitionResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SpeechRecognitionResultClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// SpeechRecognitionResult boxes the SpeechRecognitionResultClass providing a helper.
type SpeechRecognitionResultBox struct {
	SpeechRecognitionResult SpeechRecognitionResultClass
}

// Decode implements bin.Decoder for SpeechRecognitionResultBox.
func (b *SpeechRecognitionResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SpeechRecognitionResultBox to nil")
	}
	v, err := DecodeSpeechRecognitionResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SpeechRecognitionResult = v
	return nil
}

// Encode implements bin.Encode for SpeechRecognitionResultBox.
func (b *SpeechRecognitionResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SpeechRecognitionResult == nil {
		return fmt.Errorf("unable to encode SpeechRecognitionResultClass as nil")
	}
	return b.SpeechRecognitionResult.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for SpeechRecognitionResultBox.
func (b *SpeechRecognitionResultBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode SpeechRecognitionResultBox to nil")
	}
	v, err := DecodeTDLibJSONSpeechRecognitionResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SpeechRecognitionResult = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for SpeechRecognitionResultBox.
func (b *SpeechRecognitionResultBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.SpeechRecognitionResult == nil {
		return fmt.Errorf("unable to encode SpeechRecognitionResultClass as nil")
	}
	return b.SpeechRecognitionResult.EncodeTDLibJSON(buf)
}
