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

// ChannelsSetStickersRequest represents TL type `channels.setStickers#ea8ca4f9`.
// Associate a stickerset to the supergroup
//
// See https://core.telegram.org/method/channels.setStickers for reference.
type ChannelsSetStickersRequest struct {
	// Supergroup
	Channel InputChannelClass
	// The stickerset to associate
	Stickerset InputStickerSetClass
}

// ChannelsSetStickersRequestTypeID is TL type id of ChannelsSetStickersRequest.
const ChannelsSetStickersRequestTypeID = 0xea8ca4f9

// Ensuring interfaces in compile-time for ChannelsSetStickersRequest.
var (
	_ bin.Encoder     = &ChannelsSetStickersRequest{}
	_ bin.Decoder     = &ChannelsSetStickersRequest{}
	_ bin.BareEncoder = &ChannelsSetStickersRequest{}
	_ bin.BareDecoder = &ChannelsSetStickersRequest{}
)

func (s *ChannelsSetStickersRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Channel == nil) {
		return false
	}
	if !(s.Stickerset == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ChannelsSetStickersRequest) String() string {
	if s == nil {
		return "ChannelsSetStickersRequest(nil)"
	}
	type Alias ChannelsSetStickersRequest
	return fmt.Sprintf("ChannelsSetStickersRequest%+v", Alias(*s))
}

// FillFrom fills ChannelsSetStickersRequest from given interface.
func (s *ChannelsSetStickersRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetStickerset() (value InputStickerSetClass)
}) {
	s.Channel = from.GetChannel()
	s.Stickerset = from.GetStickerset()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsSetStickersRequest) TypeID() uint32 {
	return ChannelsSetStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsSetStickersRequest) TypeName() string {
	return "channels.setStickers"
}

// TypeInfo returns info about TL type.
func (s *ChannelsSetStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.setStickers",
		ID:   ChannelsSetStickersRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Stickerset",
			SchemaName: "stickerset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *ChannelsSetStickersRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode channels.setStickers#ea8ca4f9 as nil")
	}
	b.PutID(ChannelsSetStickersRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *ChannelsSetStickersRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode channels.setStickers#ea8ca4f9 as nil")
	}
	if s.Channel == nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field channel is nil")
	}
	if err := s.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field channel: %w", err)
	}
	if s.Stickerset == nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field stickerset is nil")
	}
	if err := s.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.setStickers#ea8ca4f9: field stickerset: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *ChannelsSetStickersRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode channels.setStickers#ea8ca4f9 to nil")
	}
	if err := b.ConsumeID(ChannelsSetStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.setStickers#ea8ca4f9: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *ChannelsSetStickersRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode channels.setStickers#ea8ca4f9 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.setStickers#ea8ca4f9: field channel: %w", err)
		}
		s.Channel = value
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.setStickers#ea8ca4f9: field stickerset: %w", err)
		}
		s.Stickerset = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (s *ChannelsSetStickersRequest) GetChannel() (value InputChannelClass) {
	if s == nil {
		return
	}
	return s.Channel
}

// GetStickerset returns value of Stickerset field.
func (s *ChannelsSetStickersRequest) GetStickerset() (value InputStickerSetClass) {
	if s == nil {
		return
	}
	return s.Stickerset
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (s *ChannelsSetStickersRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return s.Channel.AsNotEmpty()
}

// ChannelsSetStickers invokes method channels.setStickers#ea8ca4f9 returning error if any.
// Associate a stickerset to the supergroup
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 PARTICIPANTS_TOO_FEW: Not enough participants.
//	406 STICKERSET_OWNER_ANONYMOUS: Provided stickerset can't be installed as group stickerset to prevent admin deanonymization.
//
// See https://core.telegram.org/method/channels.setStickers for reference.
// Can be used by bots.
func (c *Client) ChannelsSetStickers(ctx context.Context, request *ChannelsSetStickersRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
