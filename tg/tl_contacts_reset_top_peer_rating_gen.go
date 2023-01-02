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

// ContactsResetTopPeerRatingRequest represents TL type `contacts.resetTopPeerRating#1ae373ac`.
// Reset rating¹ of top peer
//
// Links:
//  1. https://core.telegram.org/api/top-rating
//
// See https://core.telegram.org/method/contacts.resetTopPeerRating for reference.
type ContactsResetTopPeerRatingRequest struct {
	// Top peer category
	Category TopPeerCategoryClass
	// Peer whose rating should be reset
	Peer InputPeerClass
}

// ContactsResetTopPeerRatingRequestTypeID is TL type id of ContactsResetTopPeerRatingRequest.
const ContactsResetTopPeerRatingRequestTypeID = 0x1ae373ac

// Ensuring interfaces in compile-time for ContactsResetTopPeerRatingRequest.
var (
	_ bin.Encoder     = &ContactsResetTopPeerRatingRequest{}
	_ bin.Decoder     = &ContactsResetTopPeerRatingRequest{}
	_ bin.BareEncoder = &ContactsResetTopPeerRatingRequest{}
	_ bin.BareDecoder = &ContactsResetTopPeerRatingRequest{}
)

func (r *ContactsResetTopPeerRatingRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Category == nil) {
		return false
	}
	if !(r.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ContactsResetTopPeerRatingRequest) String() string {
	if r == nil {
		return "ContactsResetTopPeerRatingRequest(nil)"
	}
	type Alias ContactsResetTopPeerRatingRequest
	return fmt.Sprintf("ContactsResetTopPeerRatingRequest%+v", Alias(*r))
}

// FillFrom fills ContactsResetTopPeerRatingRequest from given interface.
func (r *ContactsResetTopPeerRatingRequest) FillFrom(from interface {
	GetCategory() (value TopPeerCategoryClass)
	GetPeer() (value InputPeerClass)
}) {
	r.Category = from.GetCategory()
	r.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsResetTopPeerRatingRequest) TypeID() uint32 {
	return ContactsResetTopPeerRatingRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsResetTopPeerRatingRequest) TypeName() string {
	return "contacts.resetTopPeerRating"
}

// TypeInfo returns info about TL type.
func (r *ContactsResetTopPeerRatingRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.resetTopPeerRating",
		ID:   ContactsResetTopPeerRatingRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Category",
			SchemaName: "category",
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ContactsResetTopPeerRatingRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resetTopPeerRating#1ae373ac as nil")
	}
	b.PutID(ContactsResetTopPeerRatingRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ContactsResetTopPeerRatingRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resetTopPeerRating#1ae373ac as nil")
	}
	if r.Category == nil {
		return fmt.Errorf("unable to encode contacts.resetTopPeerRating#1ae373ac: field category is nil")
	}
	if err := r.Category.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.resetTopPeerRating#1ae373ac: field category: %w", err)
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode contacts.resetTopPeerRating#1ae373ac: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.resetTopPeerRating#1ae373ac: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ContactsResetTopPeerRatingRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resetTopPeerRating#1ae373ac to nil")
	}
	if err := b.ConsumeID(ContactsResetTopPeerRatingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.resetTopPeerRating#1ae373ac: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ContactsResetTopPeerRatingRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resetTopPeerRating#1ae373ac to nil")
	}
	{
		value, err := DecodeTopPeerCategory(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.resetTopPeerRating#1ae373ac: field category: %w", err)
		}
		r.Category = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.resetTopPeerRating#1ae373ac: field peer: %w", err)
		}
		r.Peer = value
	}
	return nil
}

// GetCategory returns value of Category field.
func (r *ContactsResetTopPeerRatingRequest) GetCategory() (value TopPeerCategoryClass) {
	if r == nil {
		return
	}
	return r.Category
}

// GetPeer returns value of Peer field.
func (r *ContactsResetTopPeerRatingRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// ContactsResetTopPeerRating invokes method contacts.resetTopPeerRating#1ae373ac returning error if any.
// Reset rating¹ of top peer
//
// Links:
//  1. https://core.telegram.org/api/top-rating
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/contacts.resetTopPeerRating for reference.
func (c *Client) ContactsResetTopPeerRating(ctx context.Context, request *ContactsResetTopPeerRatingRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
