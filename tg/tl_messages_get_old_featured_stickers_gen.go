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

// MessagesGetOldFeaturedStickersRequest represents TL type `messages.getOldFeaturedStickers#7ed094a1`.
// Method for fetching previously featured stickers
//
// See https://core.telegram.org/method/messages.getOldFeaturedStickers for reference.
type MessagesGetOldFeaturedStickersRequest struct {
	// Offset
	Offset int
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// MessagesGetOldFeaturedStickersRequestTypeID is TL type id of MessagesGetOldFeaturedStickersRequest.
const MessagesGetOldFeaturedStickersRequestTypeID = 0x7ed094a1

// Ensuring interfaces in compile-time for MessagesGetOldFeaturedStickersRequest.
var (
	_ bin.Encoder     = &MessagesGetOldFeaturedStickersRequest{}
	_ bin.Decoder     = &MessagesGetOldFeaturedStickersRequest{}
	_ bin.BareEncoder = &MessagesGetOldFeaturedStickersRequest{}
	_ bin.BareDecoder = &MessagesGetOldFeaturedStickersRequest{}
)

func (g *MessagesGetOldFeaturedStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetOldFeaturedStickersRequest) String() string {
	if g == nil {
		return "MessagesGetOldFeaturedStickersRequest(nil)"
	}
	type Alias MessagesGetOldFeaturedStickersRequest
	return fmt.Sprintf("MessagesGetOldFeaturedStickersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetOldFeaturedStickersRequest from given interface.
func (g *MessagesGetOldFeaturedStickersRequest) FillFrom(from interface {
	GetOffset() (value int)
	GetLimit() (value int)
	GetHash() (value int64)
}) {
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetOldFeaturedStickersRequest) TypeID() uint32 {
	return MessagesGetOldFeaturedStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetOldFeaturedStickersRequest) TypeName() string {
	return "messages.getOldFeaturedStickers"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetOldFeaturedStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getOldFeaturedStickers",
		ID:   MessagesGetOldFeaturedStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetOldFeaturedStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getOldFeaturedStickers#7ed094a1 as nil")
	}
	b.PutID(MessagesGetOldFeaturedStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetOldFeaturedStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getOldFeaturedStickers#7ed094a1 as nil")
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetOldFeaturedStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getOldFeaturedStickers#7ed094a1 to nil")
	}
	if err := b.ConsumeID(MessagesGetOldFeaturedStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getOldFeaturedStickers#7ed094a1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetOldFeaturedStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getOldFeaturedStickers#7ed094a1 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getOldFeaturedStickers#7ed094a1: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getOldFeaturedStickers#7ed094a1: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getOldFeaturedStickers#7ed094a1: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetOffset returns value of Offset field.
func (g *MessagesGetOldFeaturedStickersRequest) GetOffset() (value int) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *MessagesGetOldFeaturedStickersRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetHash returns value of Hash field.
func (g *MessagesGetOldFeaturedStickersRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetOldFeaturedStickers invokes method messages.getOldFeaturedStickers#7ed094a1 returning error if any.
// Method for fetching previously featured stickers
//
// See https://core.telegram.org/method/messages.getOldFeaturedStickers for reference.
func (c *Client) MessagesGetOldFeaturedStickers(ctx context.Context, request *MessagesGetOldFeaturedStickersRequest) (MessagesFeaturedStickersClass, error) {
	var result MessagesFeaturedStickersBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.FeaturedStickers, nil
}
