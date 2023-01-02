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

// SearchPublicChatRequest represents TL type `searchPublicChat#3316d9ad`.
type SearchPublicChatRequest struct {
	// Username to be resolved
	Username string
}

// SearchPublicChatRequestTypeID is TL type id of SearchPublicChatRequest.
const SearchPublicChatRequestTypeID = 0x3316d9ad

// Ensuring interfaces in compile-time for SearchPublicChatRequest.
var (
	_ bin.Encoder     = &SearchPublicChatRequest{}
	_ bin.Decoder     = &SearchPublicChatRequest{}
	_ bin.BareEncoder = &SearchPublicChatRequest{}
	_ bin.BareDecoder = &SearchPublicChatRequest{}
)

func (s *SearchPublicChatRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchPublicChatRequest) String() string {
	if s == nil {
		return "SearchPublicChatRequest(nil)"
	}
	type Alias SearchPublicChatRequest
	return fmt.Sprintf("SearchPublicChatRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchPublicChatRequest) TypeID() uint32 {
	return SearchPublicChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchPublicChatRequest) TypeName() string {
	return "searchPublicChat"
}

// TypeInfo returns info about TL type.
func (s *SearchPublicChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchPublicChat",
		ID:   SearchPublicChatRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Username",
			SchemaName: "username",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchPublicChatRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchPublicChat#3316d9ad as nil")
	}
	b.PutID(SearchPublicChatRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchPublicChatRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchPublicChat#3316d9ad as nil")
	}
	b.PutString(s.Username)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchPublicChatRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchPublicChat#3316d9ad to nil")
	}
	if err := b.ConsumeID(SearchPublicChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchPublicChat#3316d9ad: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchPublicChatRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchPublicChat#3316d9ad to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchPublicChat#3316d9ad: field username: %w", err)
		}
		s.Username = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchPublicChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchPublicChat#3316d9ad as nil")
	}
	b.ObjStart()
	b.PutID("searchPublicChat")
	b.Comma()
	b.FieldStart("username")
	b.PutString(s.Username)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchPublicChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchPublicChat#3316d9ad to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchPublicChat"); err != nil {
				return fmt.Errorf("unable to decode searchPublicChat#3316d9ad: %w", err)
			}
		case "username":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchPublicChat#3316d9ad: field username: %w", err)
			}
			s.Username = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUsername returns value of Username field.
func (s *SearchPublicChatRequest) GetUsername() (value string) {
	if s == nil {
		return
	}
	return s.Username
}

// SearchPublicChat invokes method searchPublicChat#3316d9ad returning error if any.
func (c *Client) SearchPublicChat(ctx context.Context, username string) (*Chat, error) {
	var result Chat

	request := &SearchPublicChatRequest{
		Username: username,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
