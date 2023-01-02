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

// MessagesGetPeerDialogsRequest represents TL type `messages.getPeerDialogs#e470bcfd`.
// Get dialog info of specified peers
//
// See https://core.telegram.org/method/messages.getPeerDialogs for reference.
type MessagesGetPeerDialogsRequest struct {
	// Peers
	Peers []InputDialogPeerClass
}

// MessagesGetPeerDialogsRequestTypeID is TL type id of MessagesGetPeerDialogsRequest.
const MessagesGetPeerDialogsRequestTypeID = 0xe470bcfd

// Ensuring interfaces in compile-time for MessagesGetPeerDialogsRequest.
var (
	_ bin.Encoder     = &MessagesGetPeerDialogsRequest{}
	_ bin.Decoder     = &MessagesGetPeerDialogsRequest{}
	_ bin.BareEncoder = &MessagesGetPeerDialogsRequest{}
	_ bin.BareDecoder = &MessagesGetPeerDialogsRequest{}
)

func (g *MessagesGetPeerDialogsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetPeerDialogsRequest) String() string {
	if g == nil {
		return "MessagesGetPeerDialogsRequest(nil)"
	}
	type Alias MessagesGetPeerDialogsRequest
	return fmt.Sprintf("MessagesGetPeerDialogsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetPeerDialogsRequest from given interface.
func (g *MessagesGetPeerDialogsRequest) FillFrom(from interface {
	GetPeers() (value []InputDialogPeerClass)
}) {
	g.Peers = from.GetPeers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetPeerDialogsRequest) TypeID() uint32 {
	return MessagesGetPeerDialogsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetPeerDialogsRequest) TypeName() string {
	return "messages.getPeerDialogs"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetPeerDialogsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getPeerDialogs",
		ID:   MessagesGetPeerDialogsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peers",
			SchemaName: "peers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetPeerDialogsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPeerDialogs#e470bcfd as nil")
	}
	b.PutID(MessagesGetPeerDialogsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetPeerDialogsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPeerDialogs#e470bcfd as nil")
	}
	b.PutVectorHeader(len(g.Peers))
	for idx, v := range g.Peers {
		if v == nil {
			return fmt.Errorf("unable to encode messages.getPeerDialogs#e470bcfd: field peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.getPeerDialogs#e470bcfd: field peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetPeerDialogsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPeerDialogs#e470bcfd to nil")
	}
	if err := b.ConsumeID(MessagesGetPeerDialogsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getPeerDialogs#e470bcfd: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetPeerDialogsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPeerDialogs#e470bcfd to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPeerDialogs#e470bcfd: field peers: %w", err)
		}

		if headerLen > 0 {
			g.Peers = make([]InputDialogPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputDialogPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.getPeerDialogs#e470bcfd: field peers: %w", err)
			}
			g.Peers = append(g.Peers, value)
		}
	}
	return nil
}

// GetPeers returns value of Peers field.
func (g *MessagesGetPeerDialogsRequest) GetPeers() (value []InputDialogPeerClass) {
	if g == nil {
		return
	}
	return g.Peers
}

// MapPeers returns field Peers wrapped in InputDialogPeerClassArray helper.
func (g *MessagesGetPeerDialogsRequest) MapPeers() (value InputDialogPeerClassArray) {
	return InputDialogPeerClassArray(g.Peers)
}

// MessagesGetPeerDialogs invokes method messages.getPeerDialogs#e470bcfd returning error if any.
// Get dialog info of specified peers
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.getPeerDialogs for reference.
func (c *Client) MessagesGetPeerDialogs(ctx context.Context, peers []InputDialogPeerClass) (*MessagesPeerDialogs, error) {
	var result MessagesPeerDialogs

	request := &MessagesGetPeerDialogsRequest{
		Peers: peers,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
