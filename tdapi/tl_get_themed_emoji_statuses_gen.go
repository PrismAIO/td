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

// GetThemedEmojiStatusesRequest represents TL type `getThemedEmojiStatuses#6ac5c8c2`.
type GetThemedEmojiStatusesRequest struct {
}

// GetThemedEmojiStatusesRequestTypeID is TL type id of GetThemedEmojiStatusesRequest.
const GetThemedEmojiStatusesRequestTypeID = 0x6ac5c8c2

// Ensuring interfaces in compile-time for GetThemedEmojiStatusesRequest.
var (
	_ bin.Encoder     = &GetThemedEmojiStatusesRequest{}
	_ bin.Decoder     = &GetThemedEmojiStatusesRequest{}
	_ bin.BareEncoder = &GetThemedEmojiStatusesRequest{}
	_ bin.BareDecoder = &GetThemedEmojiStatusesRequest{}
)

func (g *GetThemedEmojiStatusesRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetThemedEmojiStatusesRequest) String() string {
	if g == nil {
		return "GetThemedEmojiStatusesRequest(nil)"
	}
	type Alias GetThemedEmojiStatusesRequest
	return fmt.Sprintf("GetThemedEmojiStatusesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetThemedEmojiStatusesRequest) TypeID() uint32 {
	return GetThemedEmojiStatusesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetThemedEmojiStatusesRequest) TypeName() string {
	return "getThemedEmojiStatuses"
}

// TypeInfo returns info about TL type.
func (g *GetThemedEmojiStatusesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getThemedEmojiStatuses",
		ID:   GetThemedEmojiStatusesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetThemedEmojiStatusesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getThemedEmojiStatuses#6ac5c8c2 as nil")
	}
	b.PutID(GetThemedEmojiStatusesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetThemedEmojiStatusesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getThemedEmojiStatuses#6ac5c8c2 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetThemedEmojiStatusesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getThemedEmojiStatuses#6ac5c8c2 to nil")
	}
	if err := b.ConsumeID(GetThemedEmojiStatusesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getThemedEmojiStatuses#6ac5c8c2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetThemedEmojiStatusesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getThemedEmojiStatuses#6ac5c8c2 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetThemedEmojiStatusesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getThemedEmojiStatuses#6ac5c8c2 as nil")
	}
	b.ObjStart()
	b.PutID("getThemedEmojiStatuses")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetThemedEmojiStatusesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getThemedEmojiStatuses#6ac5c8c2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getThemedEmojiStatuses"); err != nil {
				return fmt.Errorf("unable to decode getThemedEmojiStatuses#6ac5c8c2: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetThemedEmojiStatuses invokes method getThemedEmojiStatuses#6ac5c8c2 returning error if any.
func (c *Client) GetThemedEmojiStatuses(ctx context.Context) (*EmojiStatuses, error) {
	var result EmojiStatuses

	request := &GetThemedEmojiStatusesRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}