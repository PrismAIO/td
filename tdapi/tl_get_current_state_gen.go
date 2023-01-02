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

// GetCurrentStateRequest represents TL type `getCurrentState#b8fc6889`.
type GetCurrentStateRequest struct {
}

// GetCurrentStateRequestTypeID is TL type id of GetCurrentStateRequest.
const GetCurrentStateRequestTypeID = 0xb8fc6889

// Ensuring interfaces in compile-time for GetCurrentStateRequest.
var (
	_ bin.Encoder     = &GetCurrentStateRequest{}
	_ bin.Decoder     = &GetCurrentStateRequest{}
	_ bin.BareEncoder = &GetCurrentStateRequest{}
	_ bin.BareDecoder = &GetCurrentStateRequest{}
)

func (g *GetCurrentStateRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetCurrentStateRequest) String() string {
	if g == nil {
		return "GetCurrentStateRequest(nil)"
	}
	type Alias GetCurrentStateRequest
	return fmt.Sprintf("GetCurrentStateRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetCurrentStateRequest) TypeID() uint32 {
	return GetCurrentStateRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetCurrentStateRequest) TypeName() string {
	return "getCurrentState"
}

// TypeInfo returns info about TL type.
func (g *GetCurrentStateRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getCurrentState",
		ID:   GetCurrentStateRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetCurrentStateRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCurrentState#b8fc6889 as nil")
	}
	b.PutID(GetCurrentStateRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetCurrentStateRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCurrentState#b8fc6889 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetCurrentStateRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCurrentState#b8fc6889 to nil")
	}
	if err := b.ConsumeID(GetCurrentStateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getCurrentState#b8fc6889: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetCurrentStateRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCurrentState#b8fc6889 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetCurrentStateRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getCurrentState#b8fc6889 as nil")
	}
	b.ObjStart()
	b.PutID("getCurrentState")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetCurrentStateRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getCurrentState#b8fc6889 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getCurrentState"); err != nil {
				return fmt.Errorf("unable to decode getCurrentState#b8fc6889: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCurrentState invokes method getCurrentState#b8fc6889 returning error if any.
func (c *Client) GetCurrentState(ctx context.Context) (*Updates, error) {
	var result Updates

	request := &GetCurrentStateRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
