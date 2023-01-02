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

// MessagesClearRecentReactionsRequest represents TL type `messages.clearRecentReactions#9dfeefb4`.
//
// See https://core.telegram.org/method/messages.clearRecentReactions for reference.
type MessagesClearRecentReactionsRequest struct {
}

// MessagesClearRecentReactionsRequestTypeID is TL type id of MessagesClearRecentReactionsRequest.
const MessagesClearRecentReactionsRequestTypeID = 0x9dfeefb4

// Ensuring interfaces in compile-time for MessagesClearRecentReactionsRequest.
var (
	_ bin.Encoder     = &MessagesClearRecentReactionsRequest{}
	_ bin.Decoder     = &MessagesClearRecentReactionsRequest{}
	_ bin.BareEncoder = &MessagesClearRecentReactionsRequest{}
	_ bin.BareDecoder = &MessagesClearRecentReactionsRequest{}
)

func (c *MessagesClearRecentReactionsRequest) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesClearRecentReactionsRequest) String() string {
	if c == nil {
		return "MessagesClearRecentReactionsRequest(nil)"
	}
	type Alias MessagesClearRecentReactionsRequest
	return fmt.Sprintf("MessagesClearRecentReactionsRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesClearRecentReactionsRequest) TypeID() uint32 {
	return MessagesClearRecentReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesClearRecentReactionsRequest) TypeName() string {
	return "messages.clearRecentReactions"
}

// TypeInfo returns info about TL type.
func (c *MessagesClearRecentReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.clearRecentReactions",
		ID:   MessagesClearRecentReactionsRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *MessagesClearRecentReactionsRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.clearRecentReactions#9dfeefb4 as nil")
	}
	b.PutID(MessagesClearRecentReactionsRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *MessagesClearRecentReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.clearRecentReactions#9dfeefb4 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesClearRecentReactionsRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.clearRecentReactions#9dfeefb4 to nil")
	}
	if err := b.ConsumeID(MessagesClearRecentReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.clearRecentReactions#9dfeefb4: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *MessagesClearRecentReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.clearRecentReactions#9dfeefb4 to nil")
	}
	return nil
}

// MessagesClearRecentReactions invokes method messages.clearRecentReactions#9dfeefb4 returning error if any.
//
// See https://core.telegram.org/method/messages.clearRecentReactions for reference.
func (c *Client) MessagesClearRecentReactions(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &MessagesClearRecentReactionsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
