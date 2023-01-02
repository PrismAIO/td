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

// SetSupergroupStickerSetRequest represents TL type `setSupergroupStickerSet#856ead69`.
type SetSupergroupStickerSetRequest struct {
	// Identifier of the supergroup
	SupergroupID int64
	// New value of the supergroup sticker set identifier. Use 0 to remove the supergroup
	// sticker set
	StickerSetID int64
}

// SetSupergroupStickerSetRequestTypeID is TL type id of SetSupergroupStickerSetRequest.
const SetSupergroupStickerSetRequestTypeID = 0x856ead69

// Ensuring interfaces in compile-time for SetSupergroupStickerSetRequest.
var (
	_ bin.Encoder     = &SetSupergroupStickerSetRequest{}
	_ bin.Decoder     = &SetSupergroupStickerSetRequest{}
	_ bin.BareEncoder = &SetSupergroupStickerSetRequest{}
	_ bin.BareDecoder = &SetSupergroupStickerSetRequest{}
)

func (s *SetSupergroupStickerSetRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.SupergroupID == 0) {
		return false
	}
	if !(s.StickerSetID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetSupergroupStickerSetRequest) String() string {
	if s == nil {
		return "SetSupergroupStickerSetRequest(nil)"
	}
	type Alias SetSupergroupStickerSetRequest
	return fmt.Sprintf("SetSupergroupStickerSetRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetSupergroupStickerSetRequest) TypeID() uint32 {
	return SetSupergroupStickerSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetSupergroupStickerSetRequest) TypeName() string {
	return "setSupergroupStickerSet"
}

// TypeInfo returns info about TL type.
func (s *SetSupergroupStickerSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setSupergroupStickerSet",
		ID:   SetSupergroupStickerSetRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
		{
			Name:       "StickerSetID",
			SchemaName: "sticker_set_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetSupergroupStickerSetRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setSupergroupStickerSet#856ead69 as nil")
	}
	b.PutID(SetSupergroupStickerSetRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetSupergroupStickerSetRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setSupergroupStickerSet#856ead69 as nil")
	}
	b.PutInt53(s.SupergroupID)
	b.PutLong(s.StickerSetID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetSupergroupStickerSetRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setSupergroupStickerSet#856ead69 to nil")
	}
	if err := b.ConsumeID(SetSupergroupStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setSupergroupStickerSet#856ead69: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetSupergroupStickerSetRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setSupergroupStickerSet#856ead69 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setSupergroupStickerSet#856ead69: field supergroup_id: %w", err)
		}
		s.SupergroupID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode setSupergroupStickerSet#856ead69: field sticker_set_id: %w", err)
		}
		s.StickerSetID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetSupergroupStickerSetRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setSupergroupStickerSet#856ead69 as nil")
	}
	b.ObjStart()
	b.PutID("setSupergroupStickerSet")
	b.Comma()
	b.FieldStart("supergroup_id")
	b.PutInt53(s.SupergroupID)
	b.Comma()
	b.FieldStart("sticker_set_id")
	b.PutLong(s.StickerSetID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetSupergroupStickerSetRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setSupergroupStickerSet#856ead69 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setSupergroupStickerSet"); err != nil {
				return fmt.Errorf("unable to decode setSupergroupStickerSet#856ead69: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setSupergroupStickerSet#856ead69: field supergroup_id: %w", err)
			}
			s.SupergroupID = value
		case "sticker_set_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode setSupergroupStickerSet#856ead69: field sticker_set_id: %w", err)
			}
			s.StickerSetID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (s *SetSupergroupStickerSetRequest) GetSupergroupID() (value int64) {
	if s == nil {
		return
	}
	return s.SupergroupID
}

// GetStickerSetID returns value of StickerSetID field.
func (s *SetSupergroupStickerSetRequest) GetStickerSetID() (value int64) {
	if s == nil {
		return
	}
	return s.StickerSetID
}

// SetSupergroupStickerSet invokes method setSupergroupStickerSet#856ead69 returning error if any.
func (c *Client) SetSupergroupStickerSet(ctx context.Context, request *SetSupergroupStickerSetRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
