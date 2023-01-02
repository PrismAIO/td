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

// SearchChatRecentLocationMessagesRequest represents TL type `searchChatRecentLocationMessages#38a37ee6`.
type SearchChatRecentLocationMessagesRequest struct {
	// Chat identifier
	ChatID int64
	// The maximum number of messages to be returned
	Limit int32
}

// SearchChatRecentLocationMessagesRequestTypeID is TL type id of SearchChatRecentLocationMessagesRequest.
const SearchChatRecentLocationMessagesRequestTypeID = 0x38a37ee6

// Ensuring interfaces in compile-time for SearchChatRecentLocationMessagesRequest.
var (
	_ bin.Encoder     = &SearchChatRecentLocationMessagesRequest{}
	_ bin.Decoder     = &SearchChatRecentLocationMessagesRequest{}
	_ bin.BareEncoder = &SearchChatRecentLocationMessagesRequest{}
	_ bin.BareDecoder = &SearchChatRecentLocationMessagesRequest{}
)

func (s *SearchChatRecentLocationMessagesRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchChatRecentLocationMessagesRequest) String() string {
	if s == nil {
		return "SearchChatRecentLocationMessagesRequest(nil)"
	}
	type Alias SearchChatRecentLocationMessagesRequest
	return fmt.Sprintf("SearchChatRecentLocationMessagesRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchChatRecentLocationMessagesRequest) TypeID() uint32 {
	return SearchChatRecentLocationMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchChatRecentLocationMessagesRequest) TypeName() string {
	return "searchChatRecentLocationMessages"
}

// TypeInfo returns info about TL type.
func (s *SearchChatRecentLocationMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchChatRecentLocationMessages",
		ID:   SearchChatRecentLocationMessagesRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchChatRecentLocationMessagesRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchChatRecentLocationMessages#38a37ee6 as nil")
	}
	b.PutID(SearchChatRecentLocationMessagesRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchChatRecentLocationMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchChatRecentLocationMessages#38a37ee6 as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt32(s.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchChatRecentLocationMessagesRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchChatRecentLocationMessages#38a37ee6 to nil")
	}
	if err := b.ConsumeID(SearchChatRecentLocationMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchChatRecentLocationMessages#38a37ee6: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchChatRecentLocationMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchChatRecentLocationMessages#38a37ee6 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode searchChatRecentLocationMessages#38a37ee6: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchChatRecentLocationMessages#38a37ee6: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchChatRecentLocationMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchChatRecentLocationMessages#38a37ee6 as nil")
	}
	b.ObjStart()
	b.PutID("searchChatRecentLocationMessages")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchChatRecentLocationMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchChatRecentLocationMessages#38a37ee6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchChatRecentLocationMessages"); err != nil {
				return fmt.Errorf("unable to decode searchChatRecentLocationMessages#38a37ee6: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode searchChatRecentLocationMessages#38a37ee6: field chat_id: %w", err)
			}
			s.ChatID = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchChatRecentLocationMessages#38a37ee6: field limit: %w", err)
			}
			s.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SearchChatRecentLocationMessagesRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetLimit returns value of Limit field.
func (s *SearchChatRecentLocationMessagesRequest) GetLimit() (value int32) {
	if s == nil {
		return
	}
	return s.Limit
}

// SearchChatRecentLocationMessages invokes method searchChatRecentLocationMessages#38a37ee6 returning error if any.
func (c *Client) SearchChatRecentLocationMessages(ctx context.Context, request *SearchChatRecentLocationMessagesRequest) (*Messages, error) {
	var result Messages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
