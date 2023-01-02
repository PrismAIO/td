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

// GetUserProfilePhotosRequest represents TL type `getUserProfilePhotos#c9defe42`.
type GetUserProfilePhotosRequest struct {
	// User identifier
	UserID int64
	// The number of photos to skip; must be non-negative
	Offset int32
	// The maximum number of photos to be returned; up to 100
	Limit int32
}

// GetUserProfilePhotosRequestTypeID is TL type id of GetUserProfilePhotosRequest.
const GetUserProfilePhotosRequestTypeID = 0xc9defe42

// Ensuring interfaces in compile-time for GetUserProfilePhotosRequest.
var (
	_ bin.Encoder     = &GetUserProfilePhotosRequest{}
	_ bin.Decoder     = &GetUserProfilePhotosRequest{}
	_ bin.BareEncoder = &GetUserProfilePhotosRequest{}
	_ bin.BareDecoder = &GetUserProfilePhotosRequest{}
)

func (g *GetUserProfilePhotosRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == 0) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetUserProfilePhotosRequest) String() string {
	if g == nil {
		return "GetUserProfilePhotosRequest(nil)"
	}
	type Alias GetUserProfilePhotosRequest
	return fmt.Sprintf("GetUserProfilePhotosRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetUserProfilePhotosRequest) TypeID() uint32 {
	return GetUserProfilePhotosRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetUserProfilePhotosRequest) TypeName() string {
	return "getUserProfilePhotos"
}

// TypeInfo returns info about TL type.
func (g *GetUserProfilePhotosRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getUserProfilePhotos",
		ID:   GetUserProfilePhotosRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetUserProfilePhotosRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserProfilePhotos#c9defe42 as nil")
	}
	b.PutID(GetUserProfilePhotosRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetUserProfilePhotosRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserProfilePhotos#c9defe42 as nil")
	}
	b.PutInt53(g.UserID)
	b.PutInt32(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetUserProfilePhotosRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserProfilePhotos#c9defe42 to nil")
	}
	if err := b.ConsumeID(GetUserProfilePhotosRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getUserProfilePhotos#c9defe42: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetUserProfilePhotosRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserProfilePhotos#c9defe42 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getUserProfilePhotos#c9defe42: field user_id: %w", err)
		}
		g.UserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getUserProfilePhotos#c9defe42: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getUserProfilePhotos#c9defe42: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetUserProfilePhotosRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserProfilePhotos#c9defe42 as nil")
	}
	b.ObjStart()
	b.PutID("getUserProfilePhotos")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(g.UserID)
	b.Comma()
	b.FieldStart("offset")
	b.PutInt32(g.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetUserProfilePhotosRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserProfilePhotos#c9defe42 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getUserProfilePhotos"); err != nil {
				return fmt.Errorf("unable to decode getUserProfilePhotos#c9defe42: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getUserProfilePhotos#c9defe42: field user_id: %w", err)
			}
			g.UserID = value
		case "offset":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getUserProfilePhotos#c9defe42: field offset: %w", err)
			}
			g.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getUserProfilePhotos#c9defe42: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (g *GetUserProfilePhotosRequest) GetUserID() (value int64) {
	if g == nil {
		return
	}
	return g.UserID
}

// GetOffset returns value of Offset field.
func (g *GetUserProfilePhotosRequest) GetOffset() (value int32) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetUserProfilePhotosRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetUserProfilePhotos invokes method getUserProfilePhotos#c9defe42 returning error if any.
func (c *Client) GetUserProfilePhotos(ctx context.Context, request *GetUserProfilePhotosRequest) (*ChatPhotos, error) {
	var result ChatPhotos

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
