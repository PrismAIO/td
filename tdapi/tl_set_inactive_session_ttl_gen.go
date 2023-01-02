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

// SetInactiveSessionTTLRequest represents TL type `setInactiveSessionTtl#5d9ca950`.
type SetInactiveSessionTTLRequest struct {
	// New number of days of inactivity before sessions will be automatically terminated;
	// 1-366 days
	InactiveSessionTTLDays int32
}

// SetInactiveSessionTTLRequestTypeID is TL type id of SetInactiveSessionTTLRequest.
const SetInactiveSessionTTLRequestTypeID = 0x5d9ca950

// Ensuring interfaces in compile-time for SetInactiveSessionTTLRequest.
var (
	_ bin.Encoder     = &SetInactiveSessionTTLRequest{}
	_ bin.Decoder     = &SetInactiveSessionTTLRequest{}
	_ bin.BareEncoder = &SetInactiveSessionTTLRequest{}
	_ bin.BareDecoder = &SetInactiveSessionTTLRequest{}
)

func (s *SetInactiveSessionTTLRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.InactiveSessionTTLDays == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetInactiveSessionTTLRequest) String() string {
	if s == nil {
		return "SetInactiveSessionTTLRequest(nil)"
	}
	type Alias SetInactiveSessionTTLRequest
	return fmt.Sprintf("SetInactiveSessionTTLRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetInactiveSessionTTLRequest) TypeID() uint32 {
	return SetInactiveSessionTTLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetInactiveSessionTTLRequest) TypeName() string {
	return "setInactiveSessionTtl"
}

// TypeInfo returns info about TL type.
func (s *SetInactiveSessionTTLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setInactiveSessionTtl",
		ID:   SetInactiveSessionTTLRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InactiveSessionTTLDays",
			SchemaName: "inactive_session_ttl_days",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetInactiveSessionTTLRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setInactiveSessionTtl#5d9ca950 as nil")
	}
	b.PutID(SetInactiveSessionTTLRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetInactiveSessionTTLRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setInactiveSessionTtl#5d9ca950 as nil")
	}
	b.PutInt32(s.InactiveSessionTTLDays)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetInactiveSessionTTLRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setInactiveSessionTtl#5d9ca950 to nil")
	}
	if err := b.ConsumeID(SetInactiveSessionTTLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setInactiveSessionTtl#5d9ca950: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetInactiveSessionTTLRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setInactiveSessionTtl#5d9ca950 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setInactiveSessionTtl#5d9ca950: field inactive_session_ttl_days: %w", err)
		}
		s.InactiveSessionTTLDays = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetInactiveSessionTTLRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setInactiveSessionTtl#5d9ca950 as nil")
	}
	b.ObjStart()
	b.PutID("setInactiveSessionTtl")
	b.Comma()
	b.FieldStart("inactive_session_ttl_days")
	b.PutInt32(s.InactiveSessionTTLDays)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetInactiveSessionTTLRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setInactiveSessionTtl#5d9ca950 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setInactiveSessionTtl"); err != nil {
				return fmt.Errorf("unable to decode setInactiveSessionTtl#5d9ca950: %w", err)
			}
		case "inactive_session_ttl_days":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode setInactiveSessionTtl#5d9ca950: field inactive_session_ttl_days: %w", err)
			}
			s.InactiveSessionTTLDays = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInactiveSessionTTLDays returns value of InactiveSessionTTLDays field.
func (s *SetInactiveSessionTTLRequest) GetInactiveSessionTTLDays() (value int32) {
	if s == nil {
		return
	}
	return s.InactiveSessionTTLDays
}

// SetInactiveSessionTTL invokes method setInactiveSessionTtl#5d9ca950 returning error if any.
func (c *Client) SetInactiveSessionTTL(ctx context.Context, inactivesessionttldays int32) error {
	var ok Ok

	request := &SetInactiveSessionTTLRequest{
		InactiveSessionTTLDays: inactivesessionttldays,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
