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

// SetEmojiStatusRequest represents TL type `setEmojiStatus#b1e09a07`.
type SetEmojiStatusRequest struct {
	// New emoji status; pass null to switch to the default badge
	EmojiStatus EmojiStatus
	// Duration of the status, in seconds; pass 0 to keep the status active until it will be
	// changed manually
	Duration int32
}

// SetEmojiStatusRequestTypeID is TL type id of SetEmojiStatusRequest.
const SetEmojiStatusRequestTypeID = 0xb1e09a07

// Ensuring interfaces in compile-time for SetEmojiStatusRequest.
var (
	_ bin.Encoder     = &SetEmojiStatusRequest{}
	_ bin.Decoder     = &SetEmojiStatusRequest{}
	_ bin.BareEncoder = &SetEmojiStatusRequest{}
	_ bin.BareDecoder = &SetEmojiStatusRequest{}
)

func (s *SetEmojiStatusRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.EmojiStatus.Zero()) {
		return false
	}
	if !(s.Duration == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetEmojiStatusRequest) String() string {
	if s == nil {
		return "SetEmojiStatusRequest(nil)"
	}
	type Alias SetEmojiStatusRequest
	return fmt.Sprintf("SetEmojiStatusRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetEmojiStatusRequest) TypeID() uint32 {
	return SetEmojiStatusRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetEmojiStatusRequest) TypeName() string {
	return "setEmojiStatus"
}

// TypeInfo returns info about TL type.
func (s *SetEmojiStatusRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setEmojiStatus",
		ID:   SetEmojiStatusRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EmojiStatus",
			SchemaName: "emoji_status",
		},
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetEmojiStatusRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setEmojiStatus#b1e09a07 as nil")
	}
	b.PutID(SetEmojiStatusRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetEmojiStatusRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setEmojiStatus#b1e09a07 as nil")
	}
	if err := s.EmojiStatus.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setEmojiStatus#b1e09a07: field emoji_status: %w", err)
	}
	b.PutInt32(s.Duration)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetEmojiStatusRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setEmojiStatus#b1e09a07 to nil")
	}
	if err := b.ConsumeID(SetEmojiStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setEmojiStatus#b1e09a07: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetEmojiStatusRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setEmojiStatus#b1e09a07 to nil")
	}
	{
		if err := s.EmojiStatus.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setEmojiStatus#b1e09a07: field emoji_status: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setEmojiStatus#b1e09a07: field duration: %w", err)
		}
		s.Duration = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetEmojiStatusRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setEmojiStatus#b1e09a07 as nil")
	}
	b.ObjStart()
	b.PutID("setEmojiStatus")
	b.Comma()
	b.FieldStart("emoji_status")
	if err := s.EmojiStatus.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setEmojiStatus#b1e09a07: field emoji_status: %w", err)
	}
	b.Comma()
	b.FieldStart("duration")
	b.PutInt32(s.Duration)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetEmojiStatusRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setEmojiStatus#b1e09a07 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setEmojiStatus"); err != nil {
				return fmt.Errorf("unable to decode setEmojiStatus#b1e09a07: %w", err)
			}
		case "emoji_status":
			if err := s.EmojiStatus.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setEmojiStatus#b1e09a07: field emoji_status: %w", err)
			}
		case "duration":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode setEmojiStatus#b1e09a07: field duration: %w", err)
			}
			s.Duration = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEmojiStatus returns value of EmojiStatus field.
func (s *SetEmojiStatusRequest) GetEmojiStatus() (value EmojiStatus) {
	if s == nil {
		return
	}
	return s.EmojiStatus
}

// GetDuration returns value of Duration field.
func (s *SetEmojiStatusRequest) GetDuration() (value int32) {
	if s == nil {
		return
	}
	return s.Duration
}

// SetEmojiStatus invokes method setEmojiStatus#b1e09a07 returning error if any.
func (c *Client) SetEmojiStatus(ctx context.Context, request *SetEmojiStatusRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
