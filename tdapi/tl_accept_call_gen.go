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

// AcceptCallRequest represents TL type `acceptCall#d97562d0`.
type AcceptCallRequest struct {
	// Call identifier
	CallID int32
	// The call protocols supported by the application
	Protocol CallProtocol
}

// AcceptCallRequestTypeID is TL type id of AcceptCallRequest.
const AcceptCallRequestTypeID = 0xd97562d0

// Ensuring interfaces in compile-time for AcceptCallRequest.
var (
	_ bin.Encoder     = &AcceptCallRequest{}
	_ bin.Decoder     = &AcceptCallRequest{}
	_ bin.BareEncoder = &AcceptCallRequest{}
	_ bin.BareDecoder = &AcceptCallRequest{}
)

func (a *AcceptCallRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.CallID == 0) {
		return false
	}
	if !(a.Protocol.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AcceptCallRequest) String() string {
	if a == nil {
		return "AcceptCallRequest(nil)"
	}
	type Alias AcceptCallRequest
	return fmt.Sprintf("AcceptCallRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AcceptCallRequest) TypeID() uint32 {
	return AcceptCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AcceptCallRequest) TypeName() string {
	return "acceptCall"
}

// TypeInfo returns info about TL type.
func (a *AcceptCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "acceptCall",
		ID:   AcceptCallRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CallID",
			SchemaName: "call_id",
		},
		{
			Name:       "Protocol",
			SchemaName: "protocol",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AcceptCallRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode acceptCall#d97562d0 as nil")
	}
	b.PutID(AcceptCallRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AcceptCallRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode acceptCall#d97562d0 as nil")
	}
	b.PutInt32(a.CallID)
	if err := a.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode acceptCall#d97562d0: field protocol: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AcceptCallRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode acceptCall#d97562d0 to nil")
	}
	if err := b.ConsumeID(AcceptCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode acceptCall#d97562d0: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AcceptCallRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode acceptCall#d97562d0 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode acceptCall#d97562d0: field call_id: %w", err)
		}
		a.CallID = value
	}
	{
		if err := a.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode acceptCall#d97562d0: field protocol: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AcceptCallRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode acceptCall#d97562d0 as nil")
	}
	b.ObjStart()
	b.PutID("acceptCall")
	b.Comma()
	b.FieldStart("call_id")
	b.PutInt32(a.CallID)
	b.Comma()
	b.FieldStart("protocol")
	if err := a.Protocol.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode acceptCall#d97562d0: field protocol: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AcceptCallRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode acceptCall#d97562d0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("acceptCall"); err != nil {
				return fmt.Errorf("unable to decode acceptCall#d97562d0: %w", err)
			}
		case "call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode acceptCall#d97562d0: field call_id: %w", err)
			}
			a.CallID = value
		case "protocol":
			if err := a.Protocol.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode acceptCall#d97562d0: field protocol: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCallID returns value of CallID field.
func (a *AcceptCallRequest) GetCallID() (value int32) {
	if a == nil {
		return
	}
	return a.CallID
}

// GetProtocol returns value of Protocol field.
func (a *AcceptCallRequest) GetProtocol() (value CallProtocol) {
	if a == nil {
		return
	}
	return a.Protocol
}

// AcceptCall invokes method acceptCall#d97562d0 returning error if any.
func (c *Client) AcceptCall(ctx context.Context, request *AcceptCallRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
