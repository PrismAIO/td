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

// AddStickerToSetRequest represents TL type `addStickerToSet#b015455d`.
type AddStickerToSetRequest struct {
	// Sticker set owner
	UserID int64
	// Sticker set name
	Name string
	// Sticker to add to the set
	Sticker InputSticker
}

// AddStickerToSetRequestTypeID is TL type id of AddStickerToSetRequest.
const AddStickerToSetRequestTypeID = 0xb015455d

// Ensuring interfaces in compile-time for AddStickerToSetRequest.
var (
	_ bin.Encoder     = &AddStickerToSetRequest{}
	_ bin.Decoder     = &AddStickerToSetRequest{}
	_ bin.BareEncoder = &AddStickerToSetRequest{}
	_ bin.BareDecoder = &AddStickerToSetRequest{}
)

func (a *AddStickerToSetRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.UserID == 0) {
		return false
	}
	if !(a.Name == "") {
		return false
	}
	if !(a.Sticker.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddStickerToSetRequest) String() string {
	if a == nil {
		return "AddStickerToSetRequest(nil)"
	}
	type Alias AddStickerToSetRequest
	return fmt.Sprintf("AddStickerToSetRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddStickerToSetRequest) TypeID() uint32 {
	return AddStickerToSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddStickerToSetRequest) TypeName() string {
	return "addStickerToSet"
}

// TypeInfo returns info about TL type.
func (a *AddStickerToSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addStickerToSet",
		ID:   AddStickerToSetRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddStickerToSetRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addStickerToSet#b015455d as nil")
	}
	b.PutID(AddStickerToSetRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddStickerToSetRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addStickerToSet#b015455d as nil")
	}
	b.PutInt53(a.UserID)
	b.PutString(a.Name)
	if err := a.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode addStickerToSet#b015455d: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AddStickerToSetRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addStickerToSet#b015455d to nil")
	}
	if err := b.ConsumeID(AddStickerToSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addStickerToSet#b015455d: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddStickerToSetRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addStickerToSet#b015455d to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode addStickerToSet#b015455d: field user_id: %w", err)
		}
		a.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode addStickerToSet#b015455d: field name: %w", err)
		}
		a.Name = value
	}
	{
		if err := a.Sticker.Decode(b); err != nil {
			return fmt.Errorf("unable to decode addStickerToSet#b015455d: field sticker: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AddStickerToSetRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addStickerToSet#b015455d as nil")
	}
	b.ObjStart()
	b.PutID("addStickerToSet")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(a.UserID)
	b.Comma()
	b.FieldStart("name")
	b.PutString(a.Name)
	b.Comma()
	b.FieldStart("sticker")
	if err := a.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode addStickerToSet#b015455d: field sticker: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AddStickerToSetRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addStickerToSet#b015455d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("addStickerToSet"); err != nil {
				return fmt.Errorf("unable to decode addStickerToSet#b015455d: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode addStickerToSet#b015455d: field user_id: %w", err)
			}
			a.UserID = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode addStickerToSet#b015455d: field name: %w", err)
			}
			a.Name = value
		case "sticker":
			if err := a.Sticker.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode addStickerToSet#b015455d: field sticker: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (a *AddStickerToSetRequest) GetUserID() (value int64) {
	if a == nil {
		return
	}
	return a.UserID
}

// GetName returns value of Name field.
func (a *AddStickerToSetRequest) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

// GetSticker returns value of Sticker field.
func (a *AddStickerToSetRequest) GetSticker() (value InputSticker) {
	if a == nil {
		return
	}
	return a.Sticker
}

// AddStickerToSet invokes method addStickerToSet#b015455d returning error if any.
func (c *Client) AddStickerToSet(ctx context.Context, request *AddStickerToSetRequest) (*StickerSet, error) {
	var result StickerSet

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
