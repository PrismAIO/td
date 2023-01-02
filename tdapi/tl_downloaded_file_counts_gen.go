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

// DownloadedFileCounts represents TL type `downloadedFileCounts#8a572842`.
type DownloadedFileCounts struct {
	// Number of active file downloads found, including paused
	ActiveCount int32
	// Number of paused file downloads found
	PausedCount int32
	// Number of completed file downloads found
	CompletedCount int32
}

// DownloadedFileCountsTypeID is TL type id of DownloadedFileCounts.
const DownloadedFileCountsTypeID = 0x8a572842

// Ensuring interfaces in compile-time for DownloadedFileCounts.
var (
	_ bin.Encoder     = &DownloadedFileCounts{}
	_ bin.Decoder     = &DownloadedFileCounts{}
	_ bin.BareEncoder = &DownloadedFileCounts{}
	_ bin.BareDecoder = &DownloadedFileCounts{}
)

func (d *DownloadedFileCounts) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ActiveCount == 0) {
		return false
	}
	if !(d.PausedCount == 0) {
		return false
	}
	if !(d.CompletedCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DownloadedFileCounts) String() string {
	if d == nil {
		return "DownloadedFileCounts(nil)"
	}
	type Alias DownloadedFileCounts
	return fmt.Sprintf("DownloadedFileCounts%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DownloadedFileCounts) TypeID() uint32 {
	return DownloadedFileCountsTypeID
}

// TypeName returns name of type in TL schema.
func (*DownloadedFileCounts) TypeName() string {
	return "downloadedFileCounts"
}

// TypeInfo returns info about TL type.
func (d *DownloadedFileCounts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "downloadedFileCounts",
		ID:   DownloadedFileCountsTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ActiveCount",
			SchemaName: "active_count",
		},
		{
			Name:       "PausedCount",
			SchemaName: "paused_count",
		},
		{
			Name:       "CompletedCount",
			SchemaName: "completed_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DownloadedFileCounts) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode downloadedFileCounts#8a572842 as nil")
	}
	b.PutID(DownloadedFileCountsTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DownloadedFileCounts) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode downloadedFileCounts#8a572842 as nil")
	}
	b.PutInt32(d.ActiveCount)
	b.PutInt32(d.PausedCount)
	b.PutInt32(d.CompletedCount)
	return nil
}

// Decode implements bin.Decoder.
func (d *DownloadedFileCounts) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode downloadedFileCounts#8a572842 to nil")
	}
	if err := b.ConsumeID(DownloadedFileCountsTypeID); err != nil {
		return fmt.Errorf("unable to decode downloadedFileCounts#8a572842: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DownloadedFileCounts) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode downloadedFileCounts#8a572842 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode downloadedFileCounts#8a572842: field active_count: %w", err)
		}
		d.ActiveCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode downloadedFileCounts#8a572842: field paused_count: %w", err)
		}
		d.PausedCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode downloadedFileCounts#8a572842: field completed_count: %w", err)
		}
		d.CompletedCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DownloadedFileCounts) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode downloadedFileCounts#8a572842 as nil")
	}
	b.ObjStart()
	b.PutID("downloadedFileCounts")
	b.Comma()
	b.FieldStart("active_count")
	b.PutInt32(d.ActiveCount)
	b.Comma()
	b.FieldStart("paused_count")
	b.PutInt32(d.PausedCount)
	b.Comma()
	b.FieldStart("completed_count")
	b.PutInt32(d.CompletedCount)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DownloadedFileCounts) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode downloadedFileCounts#8a572842 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("downloadedFileCounts"); err != nil {
				return fmt.Errorf("unable to decode downloadedFileCounts#8a572842: %w", err)
			}
		case "active_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode downloadedFileCounts#8a572842: field active_count: %w", err)
			}
			d.ActiveCount = value
		case "paused_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode downloadedFileCounts#8a572842: field paused_count: %w", err)
			}
			d.PausedCount = value
		case "completed_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode downloadedFileCounts#8a572842: field completed_count: %w", err)
			}
			d.CompletedCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetActiveCount returns value of ActiveCount field.
func (d *DownloadedFileCounts) GetActiveCount() (value int32) {
	if d == nil {
		return
	}
	return d.ActiveCount
}

// GetPausedCount returns value of PausedCount field.
func (d *DownloadedFileCounts) GetPausedCount() (value int32) {
	if d == nil {
		return
	}
	return d.PausedCount
}

// GetCompletedCount returns value of CompletedCount field.
func (d *DownloadedFileCounts) GetCompletedCount() (value int32) {
	if d == nil {
		return
	}
	return d.CompletedCount
}
