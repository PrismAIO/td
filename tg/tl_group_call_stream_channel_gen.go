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

// GroupCallStreamChannel represents TL type `groupCallStreamChannel#80eb48af`.
// Info about an RTMP stream in a group call or livestream
//
// See https://core.telegram.org/constructor/groupCallStreamChannel for reference.
type GroupCallStreamChannel struct {
	// Channel ID
	Channel int
	// Specifies the duration of the video segment to fetch in milliseconds, by bitshifting
	// 1000 to the right scale times: duration_ms := 1000 >> scale.
	Scale int
	// Last seen timestamp to easily start fetching livestream chunks using
	// inputGroupCallStream¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/inputGroupCallStream
	LastTimestampMs int64
}

// GroupCallStreamChannelTypeID is TL type id of GroupCallStreamChannel.
const GroupCallStreamChannelTypeID = 0x80eb48af

// Ensuring interfaces in compile-time for GroupCallStreamChannel.
var (
	_ bin.Encoder     = &GroupCallStreamChannel{}
	_ bin.Decoder     = &GroupCallStreamChannel{}
	_ bin.BareEncoder = &GroupCallStreamChannel{}
	_ bin.BareDecoder = &GroupCallStreamChannel{}
)

func (g *GroupCallStreamChannel) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Channel == 0) {
		return false
	}
	if !(g.Scale == 0) {
		return false
	}
	if !(g.LastTimestampMs == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GroupCallStreamChannel) String() string {
	if g == nil {
		return "GroupCallStreamChannel(nil)"
	}
	type Alias GroupCallStreamChannel
	return fmt.Sprintf("GroupCallStreamChannel%+v", Alias(*g))
}

// FillFrom fills GroupCallStreamChannel from given interface.
func (g *GroupCallStreamChannel) FillFrom(from interface {
	GetChannel() (value int)
	GetScale() (value int)
	GetLastTimestampMs() (value int64)
}) {
	g.Channel = from.GetChannel()
	g.Scale = from.GetScale()
	g.LastTimestampMs = from.GetLastTimestampMs()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GroupCallStreamChannel) TypeID() uint32 {
	return GroupCallStreamChannelTypeID
}

// TypeName returns name of type in TL schema.
func (*GroupCallStreamChannel) TypeName() string {
	return "groupCallStreamChannel"
}

// TypeInfo returns info about TL type.
func (g *GroupCallStreamChannel) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "groupCallStreamChannel",
		ID:   GroupCallStreamChannelTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Scale",
			SchemaName: "scale",
		},
		{
			Name:       "LastTimestampMs",
			SchemaName: "last_timestamp_ms",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GroupCallStreamChannel) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallStreamChannel#80eb48af as nil")
	}
	b.PutID(GroupCallStreamChannelTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GroupCallStreamChannel) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallStreamChannel#80eb48af as nil")
	}
	b.PutInt(g.Channel)
	b.PutInt(g.Scale)
	b.PutLong(g.LastTimestampMs)
	return nil
}

// Decode implements bin.Decoder.
func (g *GroupCallStreamChannel) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallStreamChannel#80eb48af to nil")
	}
	if err := b.ConsumeID(GroupCallStreamChannelTypeID); err != nil {
		return fmt.Errorf("unable to decode groupCallStreamChannel#80eb48af: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GroupCallStreamChannel) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallStreamChannel#80eb48af to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallStreamChannel#80eb48af: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallStreamChannel#80eb48af: field scale: %w", err)
		}
		g.Scale = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallStreamChannel#80eb48af: field last_timestamp_ms: %w", err)
		}
		g.LastTimestampMs = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (g *GroupCallStreamChannel) GetChannel() (value int) {
	if g == nil {
		return
	}
	return g.Channel
}

// GetScale returns value of Scale field.
func (g *GroupCallStreamChannel) GetScale() (value int) {
	if g == nil {
		return
	}
	return g.Scale
}

// GetLastTimestampMs returns value of LastTimestampMs field.
func (g *GroupCallStreamChannel) GetLastTimestampMs() (value int64) {
	if g == nil {
		return
	}
	return g.LastTimestampMs
}
