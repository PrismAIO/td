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

// UsersGetUsersRequest represents TL type `users.getUsers#d91a548`.
// Returns basic user info according to their identifiers.
//
// See https://core.telegram.org/method/users.getUsers for reference.
type UsersGetUsersRequest struct {
	// List of user identifiers
	ID []InputUserClass
}

// UsersGetUsersRequestTypeID is TL type id of UsersGetUsersRequest.
const UsersGetUsersRequestTypeID = 0xd91a548

// Ensuring interfaces in compile-time for UsersGetUsersRequest.
var (
	_ bin.Encoder     = &UsersGetUsersRequest{}
	_ bin.Decoder     = &UsersGetUsersRequest{}
	_ bin.BareEncoder = &UsersGetUsersRequest{}
	_ bin.BareDecoder = &UsersGetUsersRequest{}
)

func (g *UsersGetUsersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *UsersGetUsersRequest) String() string {
	if g == nil {
		return "UsersGetUsersRequest(nil)"
	}
	type Alias UsersGetUsersRequest
	return fmt.Sprintf("UsersGetUsersRequest%+v", Alias(*g))
}

// FillFrom fills UsersGetUsersRequest from given interface.
func (g *UsersGetUsersRequest) FillFrom(from interface {
	GetID() (value []InputUserClass)
}) {
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UsersGetUsersRequest) TypeID() uint32 {
	return UsersGetUsersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UsersGetUsersRequest) TypeName() string {
	return "users.getUsers"
}

// TypeInfo returns info about TL type.
func (g *UsersGetUsersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "users.getUsers",
		ID:   UsersGetUsersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *UsersGetUsersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode users.getUsers#d91a548 as nil")
	}
	b.PutID(UsersGetUsersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *UsersGetUsersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode users.getUsers#d91a548 as nil")
	}
	b.PutVectorHeader(len(g.ID))
	for idx, v := range g.ID {
		if v == nil {
			return fmt.Errorf("unable to encode users.getUsers#d91a548: field id element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode users.getUsers#d91a548: field id element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *UsersGetUsersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode users.getUsers#d91a548 to nil")
	}
	if err := b.ConsumeID(UsersGetUsersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode users.getUsers#d91a548: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *UsersGetUsersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode users.getUsers#d91a548 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode users.getUsers#d91a548: field id: %w", err)
		}

		if headerLen > 0 {
			g.ID = make([]InputUserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode users.getUsers#d91a548: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (g *UsersGetUsersRequest) GetID() (value []InputUserClass) {
	if g == nil {
		return
	}
	return g.ID
}

// MapID returns field ID wrapped in InputUserClassArray helper.
func (g *UsersGetUsersRequest) MapID() (value InputUserClassArray) {
	return InputUserClassArray(g.ID)
}

// UsersGetUsers invokes method users.getUsers#d91a548 returning error if any.
// Returns basic user info according to their identifiers.
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 FROM_MESSAGE_BOT_DISABLED: Bots can't use fromMessage min constructors.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels.
//
// See https://core.telegram.org/method/users.getUsers for reference.
// Can be used by bots.
func (c *Client) UsersGetUsers(ctx context.Context, id []InputUserClass) ([]UserClass, error) {
	var result UserClassVector

	request := &UsersGetUsersRequest{
		ID: id,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []UserClass(result.Elems), nil
}
