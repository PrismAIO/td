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

// DateRange represents TL type `dateRange#51150c66`.
type DateRange struct {
	// Point in time (Unix timestamp) at which the date range begins
	StartDate int32
	// Point in time (Unix timestamp) at which the date range ends
	EndDate int32
}

// DateRangeTypeID is TL type id of DateRange.
const DateRangeTypeID = 0x51150c66

// Ensuring interfaces in compile-time for DateRange.
var (
	_ bin.Encoder     = &DateRange{}
	_ bin.Decoder     = &DateRange{}
	_ bin.BareEncoder = &DateRange{}
	_ bin.BareDecoder = &DateRange{}
)

func (d *DateRange) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.StartDate == 0) {
		return false
	}
	if !(d.EndDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DateRange) String() string {
	if d == nil {
		return "DateRange(nil)"
	}
	type Alias DateRange
	return fmt.Sprintf("DateRange%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DateRange) TypeID() uint32 {
	return DateRangeTypeID
}

// TypeName returns name of type in TL schema.
func (*DateRange) TypeName() string {
	return "dateRange"
}

// TypeInfo returns info about TL type.
func (d *DateRange) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dateRange",
		ID:   DateRangeTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StartDate",
			SchemaName: "start_date",
		},
		{
			Name:       "EndDate",
			SchemaName: "end_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DateRange) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dateRange#51150c66 as nil")
	}
	b.PutID(DateRangeTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DateRange) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dateRange#51150c66 as nil")
	}
	b.PutInt32(d.StartDate)
	b.PutInt32(d.EndDate)
	return nil
}

// Decode implements bin.Decoder.
func (d *DateRange) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dateRange#51150c66 to nil")
	}
	if err := b.ConsumeID(DateRangeTypeID); err != nil {
		return fmt.Errorf("unable to decode dateRange#51150c66: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DateRange) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dateRange#51150c66 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dateRange#51150c66: field start_date: %w", err)
		}
		d.StartDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode dateRange#51150c66: field end_date: %w", err)
		}
		d.EndDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DateRange) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode dateRange#51150c66 as nil")
	}
	b.ObjStart()
	b.PutID("dateRange")
	b.Comma()
	b.FieldStart("start_date")
	b.PutInt32(d.StartDate)
	b.Comma()
	b.FieldStart("end_date")
	b.PutInt32(d.EndDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DateRange) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode dateRange#51150c66 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("dateRange"); err != nil {
				return fmt.Errorf("unable to decode dateRange#51150c66: %w", err)
			}
		case "start_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode dateRange#51150c66: field start_date: %w", err)
			}
			d.StartDate = value
		case "end_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode dateRange#51150c66: field end_date: %w", err)
			}
			d.EndDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStartDate returns value of StartDate field.
func (d *DateRange) GetStartDate() (value int32) {
	if d == nil {
		return
	}
	return d.StartDate
}

// GetEndDate returns value of EndDate field.
func (d *DateRange) GetEndDate() (value int32) {
	if d == nil {
		return
	}
	return d.EndDate
}
