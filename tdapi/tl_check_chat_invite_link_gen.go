// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// CheckChatInviteLinkRequest represents TL type `checkChatInviteLink#e261483b`.
type CheckChatInviteLinkRequest struct {
	// Invite link to be checked
	InviteLink string
}

// CheckChatInviteLinkRequestTypeID is TL type id of CheckChatInviteLinkRequest.
const CheckChatInviteLinkRequestTypeID = 0xe261483b

// Ensuring interfaces in compile-time for CheckChatInviteLinkRequest.
var (
	_ bin.Encoder     = &CheckChatInviteLinkRequest{}
	_ bin.Decoder     = &CheckChatInviteLinkRequest{}
	_ bin.BareEncoder = &CheckChatInviteLinkRequest{}
	_ bin.BareDecoder = &CheckChatInviteLinkRequest{}
)

func (c *CheckChatInviteLinkRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.InviteLink == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckChatInviteLinkRequest) String() string {
	if c == nil {
		return "CheckChatInviteLinkRequest(nil)"
	}
	type Alias CheckChatInviteLinkRequest
	return fmt.Sprintf("CheckChatInviteLinkRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckChatInviteLinkRequest) TypeID() uint32 {
	return CheckChatInviteLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckChatInviteLinkRequest) TypeName() string {
	return "checkChatInviteLink"
}

// TypeInfo returns info about TL type.
func (c *CheckChatInviteLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkChatInviteLink",
		ID:   CheckChatInviteLinkRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckChatInviteLinkRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatInviteLink#e261483b as nil")
	}
	b.PutID(CheckChatInviteLinkRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckChatInviteLinkRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatInviteLink#e261483b as nil")
	}
	b.PutString(c.InviteLink)
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckChatInviteLinkRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatInviteLink#e261483b to nil")
	}
	if err := b.ConsumeID(CheckChatInviteLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkChatInviteLink#e261483b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckChatInviteLinkRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatInviteLink#e261483b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode checkChatInviteLink#e261483b: field invite_link: %w", err)
		}
		c.InviteLink = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckChatInviteLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkChatInviteLink#e261483b as nil")
	}
	b.ObjStart()
	b.PutID("checkChatInviteLink")
	b.Comma()
	b.FieldStart("invite_link")
	b.PutString(c.InviteLink)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckChatInviteLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkChatInviteLink#e261483b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkChatInviteLink"); err != nil {
				return fmt.Errorf("unable to decode checkChatInviteLink#e261483b: %w", err)
			}
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode checkChatInviteLink#e261483b: field invite_link: %w", err)
			}
			c.InviteLink = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInviteLink returns value of InviteLink field.
func (c *CheckChatInviteLinkRequest) GetInviteLink() (value string) {
	if c == nil {
		return
	}
	return c.InviteLink
}

// CheckChatInviteLink invokes method checkChatInviteLink#e261483b returning error if any.
func (c *Client) CheckChatInviteLink(ctx context.Context, invitelink string) (*ChatInviteLinkInfo, error) {
	var result ChatInviteLinkInfo

	request := &CheckChatInviteLinkRequest{
		InviteLink: invitelink,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
