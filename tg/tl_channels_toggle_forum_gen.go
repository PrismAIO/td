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

// ChannelsToggleForumRequest represents TL type `channels.toggleForum#a4298b29`.
//
// See https://core.telegram.org/method/channels.toggleForum for reference.
type ChannelsToggleForumRequest struct {
	// Channel field of ChannelsToggleForumRequest.
	Channel InputChannelClass
	// Enabled field of ChannelsToggleForumRequest.
	Enabled bool
}

// ChannelsToggleForumRequestTypeID is TL type id of ChannelsToggleForumRequest.
const ChannelsToggleForumRequestTypeID = 0xa4298b29

// Ensuring interfaces in compile-time for ChannelsToggleForumRequest.
var (
	_ bin.Encoder     = &ChannelsToggleForumRequest{}
	_ bin.Decoder     = &ChannelsToggleForumRequest{}
	_ bin.BareEncoder = &ChannelsToggleForumRequest{}
	_ bin.BareDecoder = &ChannelsToggleForumRequest{}
)

func (t *ChannelsToggleForumRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Channel == nil) {
		return false
	}
	if !(t.Enabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ChannelsToggleForumRequest) String() string {
	if t == nil {
		return "ChannelsToggleForumRequest(nil)"
	}
	type Alias ChannelsToggleForumRequest
	return fmt.Sprintf("ChannelsToggleForumRequest%+v", Alias(*t))
}

// FillFrom fills ChannelsToggleForumRequest from given interface.
func (t *ChannelsToggleForumRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetEnabled() (value bool)
}) {
	t.Channel = from.GetChannel()
	t.Enabled = from.GetEnabled()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsToggleForumRequest) TypeID() uint32 {
	return ChannelsToggleForumRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsToggleForumRequest) TypeName() string {
	return "channels.toggleForum"
}

// TypeInfo returns info about TL type.
func (t *ChannelsToggleForumRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.toggleForum",
		ID:   ChannelsToggleForumRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Enabled",
			SchemaName: "enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ChannelsToggleForumRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleForum#a4298b29 as nil")
	}
	b.PutID(ChannelsToggleForumRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ChannelsToggleForumRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleForum#a4298b29 as nil")
	}
	if t.Channel == nil {
		return fmt.Errorf("unable to encode channels.toggleForum#a4298b29: field channel is nil")
	}
	if err := t.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.toggleForum#a4298b29: field channel: %w", err)
	}
	b.PutBool(t.Enabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *ChannelsToggleForumRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleForum#a4298b29 to nil")
	}
	if err := b.ConsumeID(ChannelsToggleForumRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.toggleForum#a4298b29: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ChannelsToggleForumRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleForum#a4298b29 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleForum#a4298b29: field channel: %w", err)
		}
		t.Channel = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleForum#a4298b29: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (t *ChannelsToggleForumRequest) GetChannel() (value InputChannelClass) {
	if t == nil {
		return
	}
	return t.Channel
}

// GetEnabled returns value of Enabled field.
func (t *ChannelsToggleForumRequest) GetEnabled() (value bool) {
	if t == nil {
		return
	}
	return t.Enabled
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (t *ChannelsToggleForumRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return t.Channel.AsNotEmpty()
}

// ChannelsToggleForum invokes method channels.toggleForum#a4298b29 returning error if any.
//
// See https://core.telegram.org/method/channels.toggleForum for reference.
func (c *Client) ChannelsToggleForum(ctx context.Context, request *ChannelsToggleForumRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
