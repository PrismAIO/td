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

// ContactsBlockRequest represents TL type `contacts.block#68cc1411`.
// Adds the user to the blacklist.
//
// See https://core.telegram.org/method/contacts.block for reference.
type ContactsBlockRequest struct {
	// User ID
	ID InputPeerClass
}

// ContactsBlockRequestTypeID is TL type id of ContactsBlockRequest.
const ContactsBlockRequestTypeID = 0x68cc1411

// Ensuring interfaces in compile-time for ContactsBlockRequest.
var (
	_ bin.Encoder     = &ContactsBlockRequest{}
	_ bin.Decoder     = &ContactsBlockRequest{}
	_ bin.BareEncoder = &ContactsBlockRequest{}
	_ bin.BareDecoder = &ContactsBlockRequest{}
)

func (b *ContactsBlockRequest) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *ContactsBlockRequest) String() string {
	if b == nil {
		return "ContactsBlockRequest(nil)"
	}
	type Alias ContactsBlockRequest
	return fmt.Sprintf("ContactsBlockRequest%+v", Alias(*b))
}

// FillFrom fills ContactsBlockRequest from given interface.
func (b *ContactsBlockRequest) FillFrom(from interface {
	GetID() (value InputPeerClass)
}) {
	b.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsBlockRequest) TypeID() uint32 {
	return ContactsBlockRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsBlockRequest) TypeName() string {
	return "contacts.block"
}

// TypeInfo returns info about TL type.
func (b *ContactsBlockRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.block",
		ID:   ContactsBlockRequestTypeID,
	}
	if b == nil {
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
func (b *ContactsBlockRequest) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode contacts.block#68cc1411 as nil")
	}
	buf.PutID(ContactsBlockRequestTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *ContactsBlockRequest) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode contacts.block#68cc1411 as nil")
	}
	if b.ID == nil {
		return fmt.Errorf("unable to encode contacts.block#68cc1411: field id is nil")
	}
	if err := b.ID.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode contacts.block#68cc1411: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *ContactsBlockRequest) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode contacts.block#68cc1411 to nil")
	}
	if err := buf.ConsumeID(ContactsBlockRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.block#68cc1411: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *ContactsBlockRequest) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode contacts.block#68cc1411 to nil")
	}
	{
		value, err := DecodeInputPeer(buf)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.block#68cc1411: field id: %w", err)
		}
		b.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (b *ContactsBlockRequest) GetID() (value InputPeerClass) {
	if b == nil {
		return
	}
	return b.ID
}

// ContactsBlock invokes method contacts.block#68cc1411 returning error if any.
// Adds the user to the blacklist.
//
// Possible errors:
//
//	400 CONTACT_ID_INVALID: The provided contact ID is invalid.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/contacts.block for reference.
func (c *Client) ContactsBlock(ctx context.Context, id InputPeerClass) (bool, error) {
	var result BoolBox

	request := &ContactsBlockRequest{
		ID: id,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
