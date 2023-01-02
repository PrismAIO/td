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

// SearchStickersRequest represents TL type `searchStickers#5cbb2f43`.
type SearchStickersRequest struct {
	// String representation of emoji; must be non-empty
	Emoji string
	// The maximum number of stickers to be returned; 0-100
	Limit int32
}

// SearchStickersRequestTypeID is TL type id of SearchStickersRequest.
const SearchStickersRequestTypeID = 0x5cbb2f43

// Ensuring interfaces in compile-time for SearchStickersRequest.
var (
	_ bin.Encoder     = &SearchStickersRequest{}
	_ bin.Decoder     = &SearchStickersRequest{}
	_ bin.BareEncoder = &SearchStickersRequest{}
	_ bin.BareDecoder = &SearchStickersRequest{}
)

func (s *SearchStickersRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Emoji == "") {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchStickersRequest) String() string {
	if s == nil {
		return "SearchStickersRequest(nil)"
	}
	type Alias SearchStickersRequest
	return fmt.Sprintf("SearchStickersRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchStickersRequest) TypeID() uint32 {
	return SearchStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchStickersRequest) TypeName() string {
	return "searchStickers"
}

// TypeInfo returns info about TL type.
func (s *SearchStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchStickers",
		ID:   SearchStickersRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Emoji",
			SchemaName: "emoji",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchStickersRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchStickers#5cbb2f43 as nil")
	}
	b.PutID(SearchStickersRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchStickersRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchStickers#5cbb2f43 as nil")
	}
	b.PutString(s.Emoji)
	b.PutInt32(s.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchStickersRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchStickers#5cbb2f43 to nil")
	}
	if err := b.ConsumeID(SearchStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchStickers#5cbb2f43: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchStickersRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchStickers#5cbb2f43 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchStickers#5cbb2f43: field emoji: %w", err)
		}
		s.Emoji = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchStickers#5cbb2f43: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchStickersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchStickers#5cbb2f43 as nil")
	}
	b.ObjStart()
	b.PutID("searchStickers")
	b.Comma()
	b.FieldStart("emoji")
	b.PutString(s.Emoji)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchStickersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchStickers#5cbb2f43 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchStickers"); err != nil {
				return fmt.Errorf("unable to decode searchStickers#5cbb2f43: %w", err)
			}
		case "emoji":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchStickers#5cbb2f43: field emoji: %w", err)
			}
			s.Emoji = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchStickers#5cbb2f43: field limit: %w", err)
			}
			s.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEmoji returns value of Emoji field.
func (s *SearchStickersRequest) GetEmoji() (value string) {
	if s == nil {
		return
	}
	return s.Emoji
}

// GetLimit returns value of Limit field.
func (s *SearchStickersRequest) GetLimit() (value int32) {
	if s == nil {
		return
	}
	return s.Limit
}

// SearchStickers invokes method searchStickers#5cbb2f43 returning error if any.
func (c *Client) SearchStickers(ctx context.Context, request *SearchStickersRequest) (*Stickers, error) {
	var result Stickers

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
