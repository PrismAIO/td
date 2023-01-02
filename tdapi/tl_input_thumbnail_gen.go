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

// InputThumbnail represents TL type `inputThumbnail#5e515024`.
type InputThumbnail struct {
	// Thumbnail file to send. Sending thumbnails by file_id is currently not supported
	Thumbnail InputFileClass
	// Thumbnail width, usually shouldn't exceed 320. Use 0 if unknown
	Width int32
	// Thumbnail height, usually shouldn't exceed 320. Use 0 if unknown
	Height int32
}

// InputThumbnailTypeID is TL type id of InputThumbnail.
const InputThumbnailTypeID = 0x5e515024

// Ensuring interfaces in compile-time for InputThumbnail.
var (
	_ bin.Encoder     = &InputThumbnail{}
	_ bin.Decoder     = &InputThumbnail{}
	_ bin.BareEncoder = &InputThumbnail{}
	_ bin.BareDecoder = &InputThumbnail{}
)

func (i *InputThumbnail) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Thumbnail == nil) {
		return false
	}
	if !(i.Width == 0) {
		return false
	}
	if !(i.Height == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputThumbnail) String() string {
	if i == nil {
		return "InputThumbnail(nil)"
	}
	type Alias InputThumbnail
	return fmt.Sprintf("InputThumbnail%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputThumbnail) TypeID() uint32 {
	return InputThumbnailTypeID
}

// TypeName returns name of type in TL schema.
func (*InputThumbnail) TypeName() string {
	return "inputThumbnail"
}

// TypeInfo returns info about TL type.
func (i *InputThumbnail) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputThumbnail",
		ID:   InputThumbnailTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Thumbnail",
			SchemaName: "thumbnail",
		},
		{
			Name:       "Width",
			SchemaName: "width",
		},
		{
			Name:       "Height",
			SchemaName: "height",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputThumbnail) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputThumbnail#5e515024 as nil")
	}
	b.PutID(InputThumbnailTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputThumbnail) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputThumbnail#5e515024 as nil")
	}
	if i.Thumbnail == nil {
		return fmt.Errorf("unable to encode inputThumbnail#5e515024: field thumbnail is nil")
	}
	if err := i.Thumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputThumbnail#5e515024: field thumbnail: %w", err)
	}
	b.PutInt32(i.Width)
	b.PutInt32(i.Height)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputThumbnail) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputThumbnail#5e515024 to nil")
	}
	if err := b.ConsumeID(InputThumbnailTypeID); err != nil {
		return fmt.Errorf("unable to decode inputThumbnail#5e515024: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputThumbnail) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputThumbnail#5e515024 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputThumbnail#5e515024: field thumbnail: %w", err)
		}
		i.Thumbnail = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inputThumbnail#5e515024: field width: %w", err)
		}
		i.Width = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode inputThumbnail#5e515024: field height: %w", err)
		}
		i.Height = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputThumbnail) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputThumbnail#5e515024 as nil")
	}
	b.ObjStart()
	b.PutID("inputThumbnail")
	b.Comma()
	b.FieldStart("thumbnail")
	if i.Thumbnail == nil {
		return fmt.Errorf("unable to encode inputThumbnail#5e515024: field thumbnail is nil")
	}
	if err := i.Thumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputThumbnail#5e515024: field thumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("width")
	b.PutInt32(i.Width)
	b.Comma()
	b.FieldStart("height")
	b.PutInt32(i.Height)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputThumbnail) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputThumbnail#5e515024 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputThumbnail"); err != nil {
				return fmt.Errorf("unable to decode inputThumbnail#5e515024: %w", err)
			}
		case "thumbnail":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputThumbnail#5e515024: field thumbnail: %w", err)
			}
			i.Thumbnail = value
		case "width":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inputThumbnail#5e515024: field width: %w", err)
			}
			i.Width = value
		case "height":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode inputThumbnail#5e515024: field height: %w", err)
			}
			i.Height = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetThumbnail returns value of Thumbnail field.
func (i *InputThumbnail) GetThumbnail() (value InputFileClass) {
	if i == nil {
		return
	}
	return i.Thumbnail
}

// GetWidth returns value of Width field.
func (i *InputThumbnail) GetWidth() (value int32) {
	if i == nil {
		return
	}
	return i.Width
}

// GetHeight returns value of Height field.
func (i *InputThumbnail) GetHeight() (value int32) {
	if i == nil {
		return
	}
	return i.Height
}
