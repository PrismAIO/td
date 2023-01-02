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

// ChannelsEditBannedRequest represents TL type `channels.editBanned#96e6cd81`.
// Ban/unban/kick a user in a supergroup/channel¹.
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.editBanned for reference.
type ChannelsEditBannedRequest struct {
	// The supergroup/channel¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// Participant to ban
	Participant InputPeerClass
	// The banned rights
	BannedRights ChatBannedRights
}

// ChannelsEditBannedRequestTypeID is TL type id of ChannelsEditBannedRequest.
const ChannelsEditBannedRequestTypeID = 0x96e6cd81

// Ensuring interfaces in compile-time for ChannelsEditBannedRequest.
var (
	_ bin.Encoder     = &ChannelsEditBannedRequest{}
	_ bin.Decoder     = &ChannelsEditBannedRequest{}
	_ bin.BareEncoder = &ChannelsEditBannedRequest{}
	_ bin.BareDecoder = &ChannelsEditBannedRequest{}
)

func (e *ChannelsEditBannedRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Channel == nil) {
		return false
	}
	if !(e.Participant == nil) {
		return false
	}
	if !(e.BannedRights.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ChannelsEditBannedRequest) String() string {
	if e == nil {
		return "ChannelsEditBannedRequest(nil)"
	}
	type Alias ChannelsEditBannedRequest
	return fmt.Sprintf("ChannelsEditBannedRequest%+v", Alias(*e))
}

// FillFrom fills ChannelsEditBannedRequest from given interface.
func (e *ChannelsEditBannedRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetParticipant() (value InputPeerClass)
	GetBannedRights() (value ChatBannedRights)
}) {
	e.Channel = from.GetChannel()
	e.Participant = from.GetParticipant()
	e.BannedRights = from.GetBannedRights()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsEditBannedRequest) TypeID() uint32 {
	return ChannelsEditBannedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsEditBannedRequest) TypeName() string {
	return "channels.editBanned"
}

// TypeInfo returns info about TL type.
func (e *ChannelsEditBannedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.editBanned",
		ID:   ChannelsEditBannedRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Participant",
			SchemaName: "participant",
		},
		{
			Name:       "BannedRights",
			SchemaName: "banned_rights",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *ChannelsEditBannedRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editBanned#96e6cd81 as nil")
	}
	b.PutID(ChannelsEditBannedRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ChannelsEditBannedRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editBanned#96e6cd81 as nil")
	}
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editBanned#96e6cd81: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editBanned#96e6cd81: field channel: %w", err)
	}
	if e.Participant == nil {
		return fmt.Errorf("unable to encode channels.editBanned#96e6cd81: field participant is nil")
	}
	if err := e.Participant.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editBanned#96e6cd81: field participant: %w", err)
	}
	if err := e.BannedRights.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editBanned#96e6cd81: field banned_rights: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditBannedRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editBanned#96e6cd81 to nil")
	}
	if err := b.ConsumeID(ChannelsEditBannedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editBanned#96e6cd81: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ChannelsEditBannedRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editBanned#96e6cd81 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editBanned#96e6cd81: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editBanned#96e6cd81: field participant: %w", err)
		}
		e.Participant = value
	}
	{
		if err := e.BannedRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.editBanned#96e6cd81: field banned_rights: %w", err)
		}
	}
	return nil
}

// GetChannel returns value of Channel field.
func (e *ChannelsEditBannedRequest) GetChannel() (value InputChannelClass) {
	if e == nil {
		return
	}
	return e.Channel
}

// GetParticipant returns value of Participant field.
func (e *ChannelsEditBannedRequest) GetParticipant() (value InputPeerClass) {
	if e == nil {
		return
	}
	return e.Participant
}

// GetBannedRights returns value of BannedRights field.
func (e *ChannelsEditBannedRequest) GetBannedRights() (value ChatBannedRights) {
	if e == nil {
		return
	}
	return e.BannedRights
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (e *ChannelsEditBannedRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return e.Channel.AsNotEmpty()
}

// ChannelsEditBanned invokes method channels.editBanned#96e6cd81 returning error if any.
// Ban/unban/kick a user in a supergroup/channel¹.
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PARTICIPANT_ID_INVALID: The specified participant ID is invalid.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 USER_ADMIN_INVALID: You're not an admin.
//	400 USER_ID_INVALID: The provided user ID is invalid.
//
// See https://core.telegram.org/method/channels.editBanned for reference.
// Can be used by bots.
func (c *Client) ChannelsEditBanned(ctx context.Context, request *ChannelsEditBannedRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
