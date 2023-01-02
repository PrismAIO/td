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

// MessagesGetStickerSetRequest represents TL type `messages.getStickerSet#c8a0ec74`.
// Get info about a stickerset
//
// See https://core.telegram.org/method/messages.getStickerSet for reference.
type MessagesGetStickerSetRequest struct {
	// Stickerset
	Stickerset InputStickerSetClass
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// MessagesGetStickerSetRequestTypeID is TL type id of MessagesGetStickerSetRequest.
const MessagesGetStickerSetRequestTypeID = 0xc8a0ec74

// Ensuring interfaces in compile-time for MessagesGetStickerSetRequest.
var (
	_ bin.Encoder     = &MessagesGetStickerSetRequest{}
	_ bin.Decoder     = &MessagesGetStickerSetRequest{}
	_ bin.BareEncoder = &MessagesGetStickerSetRequest{}
	_ bin.BareDecoder = &MessagesGetStickerSetRequest{}
)

func (g *MessagesGetStickerSetRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Stickerset == nil) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetStickerSetRequest) String() string {
	if g == nil {
		return "MessagesGetStickerSetRequest(nil)"
	}
	type Alias MessagesGetStickerSetRequest
	return fmt.Sprintf("MessagesGetStickerSetRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetStickerSetRequest from given interface.
func (g *MessagesGetStickerSetRequest) FillFrom(from interface {
	GetStickerset() (value InputStickerSetClass)
	GetHash() (value int)
}) {
	g.Stickerset = from.GetStickerset()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetStickerSetRequest) TypeID() uint32 {
	return MessagesGetStickerSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetStickerSetRequest) TypeName() string {
	return "messages.getStickerSet"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetStickerSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getStickerSet",
		ID:   MessagesGetStickerSetRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Stickerset",
			SchemaName: "stickerset",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetStickerSetRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getStickerSet#c8a0ec74 as nil")
	}
	b.PutID(MessagesGetStickerSetRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetStickerSetRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getStickerSet#c8a0ec74 as nil")
	}
	if g.Stickerset == nil {
		return fmt.Errorf("unable to encode messages.getStickerSet#c8a0ec74: field stickerset is nil")
	}
	if err := g.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getStickerSet#c8a0ec74: field stickerset: %w", err)
	}
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetStickerSetRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getStickerSet#c8a0ec74 to nil")
	}
	if err := b.ConsumeID(MessagesGetStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getStickerSet#c8a0ec74: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetStickerSetRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getStickerSet#c8a0ec74 to nil")
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getStickerSet#c8a0ec74: field stickerset: %w", err)
		}
		g.Stickerset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getStickerSet#c8a0ec74: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetStickerset returns value of Stickerset field.
func (g *MessagesGetStickerSetRequest) GetStickerset() (value InputStickerSetClass) {
	if g == nil {
		return
	}
	return g.Stickerset
}

// GetHash returns value of Hash field.
func (g *MessagesGetStickerSetRequest) GetHash() (value int) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetStickerSet invokes method messages.getStickerSet#c8a0ec74 returning error if any.
// Get info about a stickerset
//
// Possible errors:
//
//	400 EMOTICON_STICKERPACK_MISSING: inputStickerSetDice.emoji cannot be empty.
//	406 STICKERSET_INVALID: The provided sticker set is invalid.
//
// See https://core.telegram.org/method/messages.getStickerSet for reference.
// Can be used by bots.
func (c *Client) MessagesGetStickerSet(ctx context.Context, request *MessagesGetStickerSetRequest) (MessagesStickerSetClass, error) {
	var result MessagesStickerSetBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.StickerSet, nil
}
