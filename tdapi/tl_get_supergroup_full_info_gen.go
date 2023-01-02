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

// GetSupergroupFullInfoRequest represents TL type `getSupergroupFullInfo#418d4038`.
type GetSupergroupFullInfoRequest struct {
	// Supergroup or channel identifier
	SupergroupID int64
}

// GetSupergroupFullInfoRequestTypeID is TL type id of GetSupergroupFullInfoRequest.
const GetSupergroupFullInfoRequestTypeID = 0x418d4038

// Ensuring interfaces in compile-time for GetSupergroupFullInfoRequest.
var (
	_ bin.Encoder     = &GetSupergroupFullInfoRequest{}
	_ bin.Decoder     = &GetSupergroupFullInfoRequest{}
	_ bin.BareEncoder = &GetSupergroupFullInfoRequest{}
	_ bin.BareDecoder = &GetSupergroupFullInfoRequest{}
)

func (g *GetSupergroupFullInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.SupergroupID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSupergroupFullInfoRequest) String() string {
	if g == nil {
		return "GetSupergroupFullInfoRequest(nil)"
	}
	type Alias GetSupergroupFullInfoRequest
	return fmt.Sprintf("GetSupergroupFullInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSupergroupFullInfoRequest) TypeID() uint32 {
	return GetSupergroupFullInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSupergroupFullInfoRequest) TypeName() string {
	return "getSupergroupFullInfo"
}

// TypeInfo returns info about TL type.
func (g *GetSupergroupFullInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSupergroupFullInfo",
		ID:   GetSupergroupFullInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSupergroupFullInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupergroupFullInfo#418d4038 as nil")
	}
	b.PutID(GetSupergroupFullInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSupergroupFullInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupergroupFullInfo#418d4038 as nil")
	}
	b.PutInt53(g.SupergroupID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSupergroupFullInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupergroupFullInfo#418d4038 to nil")
	}
	if err := b.ConsumeID(GetSupergroupFullInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSupergroupFullInfo#418d4038: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSupergroupFullInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupergroupFullInfo#418d4038 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getSupergroupFullInfo#418d4038: field supergroup_id: %w", err)
		}
		g.SupergroupID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSupergroupFullInfoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupergroupFullInfo#418d4038 as nil")
	}
	b.ObjStart()
	b.PutID("getSupergroupFullInfo")
	b.Comma()
	b.FieldStart("supergroup_id")
	b.PutInt53(g.SupergroupID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSupergroupFullInfoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupergroupFullInfo#418d4038 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSupergroupFullInfo"); err != nil {
				return fmt.Errorf("unable to decode getSupergroupFullInfo#418d4038: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getSupergroupFullInfo#418d4038: field supergroup_id: %w", err)
			}
			g.SupergroupID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (g *GetSupergroupFullInfoRequest) GetSupergroupID() (value int64) {
	if g == nil {
		return
	}
	return g.SupergroupID
}

// GetSupergroupFullInfo invokes method getSupergroupFullInfo#418d4038 returning error if any.
func (c *Client) GetSupergroupFullInfo(ctx context.Context, supergroupid int64) (*SupergroupFullInfo, error) {
	var result SupergroupFullInfo

	request := &GetSupergroupFullInfoRequest{
		SupergroupID: supergroupid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
