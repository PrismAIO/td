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

// GetMenuButtonRequest represents TL type `getMenuButton#e5eef440`.
type GetMenuButtonRequest struct {
	// Identifier of the user or 0 to get the default menu button
	UserID int64
}

// GetMenuButtonRequestTypeID is TL type id of GetMenuButtonRequest.
const GetMenuButtonRequestTypeID = 0xe5eef440

// Ensuring interfaces in compile-time for GetMenuButtonRequest.
var (
	_ bin.Encoder     = &GetMenuButtonRequest{}
	_ bin.Decoder     = &GetMenuButtonRequest{}
	_ bin.BareEncoder = &GetMenuButtonRequest{}
	_ bin.BareDecoder = &GetMenuButtonRequest{}
)

func (g *GetMenuButtonRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMenuButtonRequest) String() string {
	if g == nil {
		return "GetMenuButtonRequest(nil)"
	}
	type Alias GetMenuButtonRequest
	return fmt.Sprintf("GetMenuButtonRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMenuButtonRequest) TypeID() uint32 {
	return GetMenuButtonRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMenuButtonRequest) TypeName() string {
	return "getMenuButton"
}

// TypeInfo returns info about TL type.
func (g *GetMenuButtonRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMenuButton",
		ID:   GetMenuButtonRequestTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMenuButtonRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMenuButton#e5eef440 as nil")
	}
	b.PutID(GetMenuButtonRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMenuButtonRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMenuButton#e5eef440 as nil")
	}
	b.PutInt53(g.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMenuButtonRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMenuButton#e5eef440 to nil")
	}
	if err := b.ConsumeID(GetMenuButtonRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMenuButton#e5eef440: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMenuButtonRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMenuButton#e5eef440 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMenuButton#e5eef440: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMenuButtonRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMenuButton#e5eef440 as nil")
	}
	b.ObjStart()
	b.PutID("getMenuButton")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(g.UserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMenuButtonRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMenuButton#e5eef440 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMenuButton"); err != nil {
				return fmt.Errorf("unable to decode getMenuButton#e5eef440: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMenuButton#e5eef440: field user_id: %w", err)
			}
			g.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (g *GetMenuButtonRequest) GetUserID() (value int64) {
	if g == nil {
		return
	}
	return g.UserID
}

// GetMenuButton invokes method getMenuButton#e5eef440 returning error if any.
func (c *Client) GetMenuButton(ctx context.Context, userid int64) (*BotMenuButton, error) {
	var result BotMenuButton

	request := &GetMenuButtonRequest{
		UserID: userid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
