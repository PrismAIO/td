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

// TopPeerCategoryPeers represents TL type `topPeerCategoryPeers#fb834291`.
// Top peer category
//
// See https://core.telegram.org/constructor/topPeerCategoryPeers for reference.
type TopPeerCategoryPeers struct {
	// Top peer category of peers
	Category TopPeerCategoryClass
	// Count of peers
	Count int
	// Peers
	Peers []TopPeer
}

// TopPeerCategoryPeersTypeID is TL type id of TopPeerCategoryPeers.
const TopPeerCategoryPeersTypeID = 0xfb834291

// Ensuring interfaces in compile-time for TopPeerCategoryPeers.
var (
	_ bin.Encoder     = &TopPeerCategoryPeers{}
	_ bin.Decoder     = &TopPeerCategoryPeers{}
	_ bin.BareEncoder = &TopPeerCategoryPeers{}
	_ bin.BareDecoder = &TopPeerCategoryPeers{}
)

func (t *TopPeerCategoryPeers) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Category == nil) {
		return false
	}
	if !(t.Count == 0) {
		return false
	}
	if !(t.Peers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TopPeerCategoryPeers) String() string {
	if t == nil {
		return "TopPeerCategoryPeers(nil)"
	}
	type Alias TopPeerCategoryPeers
	return fmt.Sprintf("TopPeerCategoryPeers%+v", Alias(*t))
}

// FillFrom fills TopPeerCategoryPeers from given interface.
func (t *TopPeerCategoryPeers) FillFrom(from interface {
	GetCategory() (value TopPeerCategoryClass)
	GetCount() (value int)
	GetPeers() (value []TopPeer)
}) {
	t.Category = from.GetCategory()
	t.Count = from.GetCount()
	t.Peers = from.GetPeers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TopPeerCategoryPeers) TypeID() uint32 {
	return TopPeerCategoryPeersTypeID
}

// TypeName returns name of type in TL schema.
func (*TopPeerCategoryPeers) TypeName() string {
	return "topPeerCategoryPeers"
}

// TypeInfo returns info about TL type.
func (t *TopPeerCategoryPeers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "topPeerCategoryPeers",
		ID:   TopPeerCategoryPeersTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Category",
			SchemaName: "category",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Peers",
			SchemaName: "peers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TopPeerCategoryPeers) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryPeers#fb834291 as nil")
	}
	b.PutID(TopPeerCategoryPeersTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TopPeerCategoryPeers) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode topPeerCategoryPeers#fb834291 as nil")
	}
	if t.Category == nil {
		return fmt.Errorf("unable to encode topPeerCategoryPeers#fb834291: field category is nil")
	}
	if err := t.Category.Encode(b); err != nil {
		return fmt.Errorf("unable to encode topPeerCategoryPeers#fb834291: field category: %w", err)
	}
	b.PutInt(t.Count)
	b.PutVectorHeader(len(t.Peers))
	for idx, v := range t.Peers {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode topPeerCategoryPeers#fb834291: field peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TopPeerCategoryPeers) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryPeers#fb834291 to nil")
	}
	if err := b.ConsumeID(TopPeerCategoryPeersTypeID); err != nil {
		return fmt.Errorf("unable to decode topPeerCategoryPeers#fb834291: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TopPeerCategoryPeers) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode topPeerCategoryPeers#fb834291 to nil")
	}
	{
		value, err := DecodeTopPeerCategory(b)
		if err != nil {
			return fmt.Errorf("unable to decode topPeerCategoryPeers#fb834291: field category: %w", err)
		}
		t.Category = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode topPeerCategoryPeers#fb834291: field count: %w", err)
		}
		t.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode topPeerCategoryPeers#fb834291: field peers: %w", err)
		}

		if headerLen > 0 {
			t.Peers = make([]TopPeer, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TopPeer
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode topPeerCategoryPeers#fb834291: field peers: %w", err)
			}
			t.Peers = append(t.Peers, value)
		}
	}
	return nil
}

// GetCategory returns value of Category field.
func (t *TopPeerCategoryPeers) GetCategory() (value TopPeerCategoryClass) {
	if t == nil {
		return
	}
	return t.Category
}

// GetCount returns value of Count field.
func (t *TopPeerCategoryPeers) GetCount() (value int) {
	if t == nil {
		return
	}
	return t.Count
}

// GetPeers returns value of Peers field.
func (t *TopPeerCategoryPeers) GetPeers() (value []TopPeer) {
	if t == nil {
		return
	}
	return t.Peers
}
