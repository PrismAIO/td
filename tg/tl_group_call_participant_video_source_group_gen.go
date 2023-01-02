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

// GroupCallParticipantVideoSourceGroup represents TL type `groupCallParticipantVideoSourceGroup#dcb118b7`.
// Describes a group of video synchronization source identifiers
//
// See https://core.telegram.org/constructor/groupCallParticipantVideoSourceGroup for reference.
type GroupCallParticipantVideoSourceGroup struct {
	// SDP semantics
	Semantics string
	// Source IDs
	Sources []int
}

// GroupCallParticipantVideoSourceGroupTypeID is TL type id of GroupCallParticipantVideoSourceGroup.
const GroupCallParticipantVideoSourceGroupTypeID = 0xdcb118b7

// Ensuring interfaces in compile-time for GroupCallParticipantVideoSourceGroup.
var (
	_ bin.Encoder     = &GroupCallParticipantVideoSourceGroup{}
	_ bin.Decoder     = &GroupCallParticipantVideoSourceGroup{}
	_ bin.BareEncoder = &GroupCallParticipantVideoSourceGroup{}
	_ bin.BareDecoder = &GroupCallParticipantVideoSourceGroup{}
)

func (g *GroupCallParticipantVideoSourceGroup) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Semantics == "") {
		return false
	}
	if !(g.Sources == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GroupCallParticipantVideoSourceGroup) String() string {
	if g == nil {
		return "GroupCallParticipantVideoSourceGroup(nil)"
	}
	type Alias GroupCallParticipantVideoSourceGroup
	return fmt.Sprintf("GroupCallParticipantVideoSourceGroup%+v", Alias(*g))
}

// FillFrom fills GroupCallParticipantVideoSourceGroup from given interface.
func (g *GroupCallParticipantVideoSourceGroup) FillFrom(from interface {
	GetSemantics() (value string)
	GetSources() (value []int)
}) {
	g.Semantics = from.GetSemantics()
	g.Sources = from.GetSources()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GroupCallParticipantVideoSourceGroup) TypeID() uint32 {
	return GroupCallParticipantVideoSourceGroupTypeID
}

// TypeName returns name of type in TL schema.
func (*GroupCallParticipantVideoSourceGroup) TypeName() string {
	return "groupCallParticipantVideoSourceGroup"
}

// TypeInfo returns info about TL type.
func (g *GroupCallParticipantVideoSourceGroup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "groupCallParticipantVideoSourceGroup",
		ID:   GroupCallParticipantVideoSourceGroupTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Semantics",
			SchemaName: "semantics",
		},
		{
			Name:       "Sources",
			SchemaName: "sources",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GroupCallParticipantVideoSourceGroup) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallParticipantVideoSourceGroup#dcb118b7 as nil")
	}
	b.PutID(GroupCallParticipantVideoSourceGroupTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GroupCallParticipantVideoSourceGroup) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallParticipantVideoSourceGroup#dcb118b7 as nil")
	}
	b.PutString(g.Semantics)
	b.PutVectorHeader(len(g.Sources))
	for _, v := range g.Sources {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GroupCallParticipantVideoSourceGroup) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallParticipantVideoSourceGroup#dcb118b7 to nil")
	}
	if err := b.ConsumeID(GroupCallParticipantVideoSourceGroupTypeID); err != nil {
		return fmt.Errorf("unable to decode groupCallParticipantVideoSourceGroup#dcb118b7: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GroupCallParticipantVideoSourceGroup) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallParticipantVideoSourceGroup#dcb118b7 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipantVideoSourceGroup#dcb118b7: field semantics: %w", err)
		}
		g.Semantics = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallParticipantVideoSourceGroup#dcb118b7: field sources: %w", err)
		}

		if headerLen > 0 {
			g.Sources = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode groupCallParticipantVideoSourceGroup#dcb118b7: field sources: %w", err)
			}
			g.Sources = append(g.Sources, value)
		}
	}
	return nil
}

// GetSemantics returns value of Semantics field.
func (g *GroupCallParticipantVideoSourceGroup) GetSemantics() (value string) {
	if g == nil {
		return
	}
	return g.Semantics
}

// GetSources returns value of Sources field.
func (g *GroupCallParticipantVideoSourceGroup) GetSources() (value []int) {
	if g == nil {
		return
	}
	return g.Sources
}
