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

// StickersAddStickerToSetRequest represents TL type `stickers.addStickerToSet#8653febe`.
// Add a sticker to a stickerset, bots only. The sticker set must have been created by
// the bot.
//
// See https://core.telegram.org/method/stickers.addStickerToSet for reference.
type StickersAddStickerToSetRequest struct {
	// The stickerset
	Stickerset InputStickerSetClass
	// The sticker
	Sticker InputStickerSetItem
}

// StickersAddStickerToSetRequestTypeID is TL type id of StickersAddStickerToSetRequest.
const StickersAddStickerToSetRequestTypeID = 0x8653febe

// Ensuring interfaces in compile-time for StickersAddStickerToSetRequest.
var (
	_ bin.Encoder     = &StickersAddStickerToSetRequest{}
	_ bin.Decoder     = &StickersAddStickerToSetRequest{}
	_ bin.BareEncoder = &StickersAddStickerToSetRequest{}
	_ bin.BareDecoder = &StickersAddStickerToSetRequest{}
)

func (a *StickersAddStickerToSetRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Stickerset == nil) {
		return false
	}
	if !(a.Sticker.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *StickersAddStickerToSetRequest) String() string {
	if a == nil {
		return "StickersAddStickerToSetRequest(nil)"
	}
	type Alias StickersAddStickerToSetRequest
	return fmt.Sprintf("StickersAddStickerToSetRequest%+v", Alias(*a))
}

// FillFrom fills StickersAddStickerToSetRequest from given interface.
func (a *StickersAddStickerToSetRequest) FillFrom(from interface {
	GetStickerset() (value InputStickerSetClass)
	GetSticker() (value InputStickerSetItem)
}) {
	a.Stickerset = from.GetStickerset()
	a.Sticker = from.GetSticker()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickersAddStickerToSetRequest) TypeID() uint32 {
	return StickersAddStickerToSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StickersAddStickerToSetRequest) TypeName() string {
	return "stickers.addStickerToSet"
}

// TypeInfo returns info about TL type.
func (a *StickersAddStickerToSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickers.addStickerToSet",
		ID:   StickersAddStickerToSetRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Stickerset",
			SchemaName: "stickerset",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *StickersAddStickerToSetRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stickers.addStickerToSet#8653febe as nil")
	}
	b.PutID(StickersAddStickerToSetRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *StickersAddStickerToSetRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stickers.addStickerToSet#8653febe as nil")
	}
	if a.Stickerset == nil {
		return fmt.Errorf("unable to encode stickers.addStickerToSet#8653febe: field stickerset is nil")
	}
	if err := a.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.addStickerToSet#8653febe: field stickerset: %w", err)
	}
	if err := a.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.addStickerToSet#8653febe: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *StickersAddStickerToSetRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stickers.addStickerToSet#8653febe to nil")
	}
	if err := b.ConsumeID(StickersAddStickerToSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.addStickerToSet#8653febe: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *StickersAddStickerToSetRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stickers.addStickerToSet#8653febe to nil")
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.addStickerToSet#8653febe: field stickerset: %w", err)
		}
		a.Stickerset = value
	}
	{
		if err := a.Sticker.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickers.addStickerToSet#8653febe: field sticker: %w", err)
		}
	}
	return nil
}

// GetStickerset returns value of Stickerset field.
func (a *StickersAddStickerToSetRequest) GetStickerset() (value InputStickerSetClass) {
	if a == nil {
		return
	}
	return a.Stickerset
}

// GetSticker returns value of Sticker field.
func (a *StickersAddStickerToSetRequest) GetSticker() (value InputStickerSetItem) {
	if a == nil {
		return
	}
	return a.Sticker
}

// StickersAddStickerToSet invokes method stickers.addStickerToSet#8653febe returning error if any.
// Add a sticker to a stickerset, bots only. The sticker set must have been created by
// the bot.
//
// Possible errors:
//
//	400 BOT_MISSING: Only bots can call this method, please use @stickers if you're a user.
//	400 STICKERPACK_STICKERS_TOO_MUCH: There are too many stickers in this stickerpack, you can't add any more.
//	400 STICKERSET_INVALID: The provided sticker set is invalid.
//	400 STICKERS_TOO_MUCH: There are too many stickers in this stickerpack, you can't add any more.
//	400 STICKER_PNG_NOPNG: One of the specified stickers is not a valid PNG file.
//	400 STICKER_TGS_NOTGS: Invalid TGS sticker provided.
//
// See https://core.telegram.org/method/stickers.addStickerToSet for reference.
// Can be used by bots.
func (c *Client) StickersAddStickerToSet(ctx context.Context, request *StickersAddStickerToSetRequest) (MessagesStickerSetClass, error) {
	var result MessagesStickerSetBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.StickerSet, nil
}
