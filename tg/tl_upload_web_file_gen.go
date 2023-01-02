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

// UploadWebFile represents TL type `upload.webFile#21e753bc`.
// Represents a chunk of an HTTP webfile¹ downloaded through telegram's secure MTProto
// servers
//
// Links:
//  1. https://core.telegram.org/api/files
//
// See https://core.telegram.org/constructor/upload.webFile for reference.
type UploadWebFile struct {
	// File size
	Size int
	// Mime type
	MimeType string
	// File type
	FileType StorageFileTypeClass
	// Modified time
	Mtime int
	// Data
	Bytes []byte
}

// UploadWebFileTypeID is TL type id of UploadWebFile.
const UploadWebFileTypeID = 0x21e753bc

// Ensuring interfaces in compile-time for UploadWebFile.
var (
	_ bin.Encoder     = &UploadWebFile{}
	_ bin.Decoder     = &UploadWebFile{}
	_ bin.BareEncoder = &UploadWebFile{}
	_ bin.BareDecoder = &UploadWebFile{}
)

func (w *UploadWebFile) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Size == 0) {
		return false
	}
	if !(w.MimeType == "") {
		return false
	}
	if !(w.FileType == nil) {
		return false
	}
	if !(w.Mtime == 0) {
		return false
	}
	if !(w.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *UploadWebFile) String() string {
	if w == nil {
		return "UploadWebFile(nil)"
	}
	type Alias UploadWebFile
	return fmt.Sprintf("UploadWebFile%+v", Alias(*w))
}

// FillFrom fills UploadWebFile from given interface.
func (w *UploadWebFile) FillFrom(from interface {
	GetSize() (value int)
	GetMimeType() (value string)
	GetFileType() (value StorageFileTypeClass)
	GetMtime() (value int)
	GetBytes() (value []byte)
}) {
	w.Size = from.GetSize()
	w.MimeType = from.GetMimeType()
	w.FileType = from.GetFileType()
	w.Mtime = from.GetMtime()
	w.Bytes = from.GetBytes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadWebFile) TypeID() uint32 {
	return UploadWebFileTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadWebFile) TypeName() string {
	return "upload.webFile"
}

// TypeInfo returns info about TL type.
func (w *UploadWebFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.webFile",
		ID:   UploadWebFileTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
		{
			Name:       "FileType",
			SchemaName: "file_type",
		},
		{
			Name:       "Mtime",
			SchemaName: "mtime",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *UploadWebFile) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode upload.webFile#21e753bc as nil")
	}
	b.PutID(UploadWebFileTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *UploadWebFile) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode upload.webFile#21e753bc as nil")
	}
	b.PutInt(w.Size)
	b.PutString(w.MimeType)
	if w.FileType == nil {
		return fmt.Errorf("unable to encode upload.webFile#21e753bc: field file_type is nil")
	}
	if err := w.FileType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upload.webFile#21e753bc: field file_type: %w", err)
	}
	b.PutInt(w.Mtime)
	b.PutBytes(w.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (w *UploadWebFile) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode upload.webFile#21e753bc to nil")
	}
	if err := b.ConsumeID(UploadWebFileTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.webFile#21e753bc: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *UploadWebFile) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode upload.webFile#21e753bc to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.webFile#21e753bc: field size: %w", err)
		}
		w.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode upload.webFile#21e753bc: field mime_type: %w", err)
		}
		w.MimeType = value
	}
	{
		value, err := DecodeStorageFileType(b)
		if err != nil {
			return fmt.Errorf("unable to decode upload.webFile#21e753bc: field file_type: %w", err)
		}
		w.FileType = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.webFile#21e753bc: field mtime: %w", err)
		}
		w.Mtime = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.webFile#21e753bc: field bytes: %w", err)
		}
		w.Bytes = value
	}
	return nil
}

// GetSize returns value of Size field.
func (w *UploadWebFile) GetSize() (value int) {
	if w == nil {
		return
	}
	return w.Size
}

// GetMimeType returns value of MimeType field.
func (w *UploadWebFile) GetMimeType() (value string) {
	if w == nil {
		return
	}
	return w.MimeType
}

// GetFileType returns value of FileType field.
func (w *UploadWebFile) GetFileType() (value StorageFileTypeClass) {
	if w == nil {
		return
	}
	return w.FileType
}

// GetMtime returns value of Mtime field.
func (w *UploadWebFile) GetMtime() (value int) {
	if w == nil {
		return
	}
	return w.Mtime
}

// GetBytes returns value of Bytes field.
func (w *UploadWebFile) GetBytes() (value []byte) {
	if w == nil {
		return
	}
	return w.Bytes
}
