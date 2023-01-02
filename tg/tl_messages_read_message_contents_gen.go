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

// MessagesReadMessageContentsRequest represents TL type `messages.readMessageContents#36a73f77`.
// Notifies the sender about the recipient having listened a voice message or watched a
// video.
//
// See https://core.telegram.org/method/messages.readMessageContents for reference.
type MessagesReadMessageContentsRequest struct {
	// Message ID list
	ID []int
}

// MessagesReadMessageContentsRequestTypeID is TL type id of MessagesReadMessageContentsRequest.
const MessagesReadMessageContentsRequestTypeID = 0x36a73f77

// Ensuring interfaces in compile-time for MessagesReadMessageContentsRequest.
var (
	_ bin.Encoder     = &MessagesReadMessageContentsRequest{}
	_ bin.Decoder     = &MessagesReadMessageContentsRequest{}
	_ bin.BareEncoder = &MessagesReadMessageContentsRequest{}
	_ bin.BareDecoder = &MessagesReadMessageContentsRequest{}
)

func (r *MessagesReadMessageContentsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReadMessageContentsRequest) String() string {
	if r == nil {
		return "MessagesReadMessageContentsRequest(nil)"
	}
	type Alias MessagesReadMessageContentsRequest
	return fmt.Sprintf("MessagesReadMessageContentsRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReadMessageContentsRequest from given interface.
func (r *MessagesReadMessageContentsRequest) FillFrom(from interface {
	GetID() (value []int)
}) {
	r.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReadMessageContentsRequest) TypeID() uint32 {
	return MessagesReadMessageContentsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReadMessageContentsRequest) TypeName() string {
	return "messages.readMessageContents"
}

// TypeInfo returns info about TL type.
func (r *MessagesReadMessageContentsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.readMessageContents",
		ID:   MessagesReadMessageContentsRequestTypeID,
	}
	if r == nil {
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
func (r *MessagesReadMessageContentsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readMessageContents#36a73f77 as nil")
	}
	b.PutID(MessagesReadMessageContentsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReadMessageContentsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readMessageContents#36a73f77 as nil")
	}
	b.PutVectorHeader(len(r.ID))
	for _, v := range r.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReadMessageContentsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readMessageContents#36a73f77 to nil")
	}
	if err := b.ConsumeID(MessagesReadMessageContentsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.readMessageContents#36a73f77: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReadMessageContentsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readMessageContents#36a73f77 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.readMessageContents#36a73f77: field id: %w", err)
		}

		if headerLen > 0 {
			r.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.readMessageContents#36a73f77: field id: %w", err)
			}
			r.ID = append(r.ID, value)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (r *MessagesReadMessageContentsRequest) GetID() (value []int) {
	if r == nil {
		return
	}
	return r.ID
}

// MessagesReadMessageContents invokes method messages.readMessageContents#36a73f77 returning error if any.
// Notifies the sender about the recipient having listened a voice message or watched a
// video.
//
// See https://core.telegram.org/method/messages.readMessageContents for reference.
func (c *Client) MessagesReadMessageContents(ctx context.Context, id []int) (*MessagesAffectedMessages, error) {
	var result MessagesAffectedMessages

	request := &MessagesReadMessageContentsRequest{
		ID: id,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
