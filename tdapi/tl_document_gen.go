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

// Document represents TL type `document#af19afd8`.
type Document struct {
	// Original name of the file; as defined by the sender
	FileName string
	// MIME type of the file; as defined by the sender
	MimeType string
	// Document minithumbnail; may be null
	Minithumbnail Minithumbnail
	// Document thumbnail in JPEG or PNG format (PNG will be used only for background
	// patterns); as defined by the sender; may be null
	Thumbnail Thumbnail
	// File containing the document
	Document File
}

// DocumentTypeID is TL type id of Document.
const DocumentTypeID = 0xaf19afd8

// Ensuring interfaces in compile-time for Document.
var (
	_ bin.Encoder     = &Document{}
	_ bin.Decoder     = &Document{}
	_ bin.BareEncoder = &Document{}
	_ bin.BareDecoder = &Document{}
)

func (d *Document) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.FileName == "") {
		return false
	}
	if !(d.MimeType == "") {
		return false
	}
	if !(d.Minithumbnail.Zero()) {
		return false
	}
	if !(d.Thumbnail.Zero()) {
		return false
	}
	if !(d.Document.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *Document) String() string {
	if d == nil {
		return "Document(nil)"
	}
	type Alias Document
	return fmt.Sprintf("Document%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Document) TypeID() uint32 {
	return DocumentTypeID
}

// TypeName returns name of type in TL schema.
func (*Document) TypeName() string {
	return "document"
}

// TypeInfo returns info about TL type.
func (d *Document) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "document",
		ID:   DocumentTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileName",
			SchemaName: "file_name",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
		{
			Name:       "Minithumbnail",
			SchemaName: "minithumbnail",
		},
		{
			Name:       "Thumbnail",
			SchemaName: "thumbnail",
		},
		{
			Name:       "Document",
			SchemaName: "document",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *Document) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode document#af19afd8 as nil")
	}
	b.PutID(DocumentTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *Document) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode document#af19afd8 as nil")
	}
	b.PutString(d.FileName)
	b.PutString(d.MimeType)
	if err := d.Minithumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode document#af19afd8: field minithumbnail: %w", err)
	}
	if err := d.Thumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode document#af19afd8: field thumbnail: %w", err)
	}
	if err := d.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode document#af19afd8: field document: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *Document) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode document#af19afd8 to nil")
	}
	if err := b.ConsumeID(DocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode document#af19afd8: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *Document) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode document#af19afd8 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode document#af19afd8: field file_name: %w", err)
		}
		d.FileName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode document#af19afd8: field mime_type: %w", err)
		}
		d.MimeType = value
	}
	{
		if err := d.Minithumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode document#af19afd8: field minithumbnail: %w", err)
		}
	}
	{
		if err := d.Thumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode document#af19afd8: field thumbnail: %w", err)
		}
	}
	{
		if err := d.Document.Decode(b); err != nil {
			return fmt.Errorf("unable to decode document#af19afd8: field document: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *Document) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode document#af19afd8 as nil")
	}
	b.ObjStart()
	b.PutID("document")
	b.Comma()
	b.FieldStart("file_name")
	b.PutString(d.FileName)
	b.Comma()
	b.FieldStart("mime_type")
	b.PutString(d.MimeType)
	b.Comma()
	b.FieldStart("minithumbnail")
	if err := d.Minithumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode document#af19afd8: field minithumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("thumbnail")
	if err := d.Thumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode document#af19afd8: field thumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("document")
	if err := d.Document.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode document#af19afd8: field document: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *Document) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode document#af19afd8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("document"); err != nil {
				return fmt.Errorf("unable to decode document#af19afd8: %w", err)
			}
		case "file_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode document#af19afd8: field file_name: %w", err)
			}
			d.FileName = value
		case "mime_type":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode document#af19afd8: field mime_type: %w", err)
			}
			d.MimeType = value
		case "minithumbnail":
			if err := d.Minithumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode document#af19afd8: field minithumbnail: %w", err)
			}
		case "thumbnail":
			if err := d.Thumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode document#af19afd8: field thumbnail: %w", err)
			}
		case "document":
			if err := d.Document.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode document#af19afd8: field document: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFileName returns value of FileName field.
func (d *Document) GetFileName() (value string) {
	if d == nil {
		return
	}
	return d.FileName
}

// GetMimeType returns value of MimeType field.
func (d *Document) GetMimeType() (value string) {
	if d == nil {
		return
	}
	return d.MimeType
}

// GetMinithumbnail returns value of Minithumbnail field.
func (d *Document) GetMinithumbnail() (value Minithumbnail) {
	if d == nil {
		return
	}
	return d.Minithumbnail
}

// GetThumbnail returns value of Thumbnail field.
func (d *Document) GetThumbnail() (value Thumbnail) {
	if d == nil {
		return
	}
	return d.Thumbnail
}

// GetDocument returns value of Document field.
func (d *Document) GetDocument() (value File) {
	if d == nil {
		return
	}
	return d.Document
}
