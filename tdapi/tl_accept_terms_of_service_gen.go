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

// AcceptTermsOfServiceRequest represents TL type `acceptTermsOfService#7efe03e4`.
type AcceptTermsOfServiceRequest struct {
	// Terms of service identifier
	TermsOfServiceID string
}

// AcceptTermsOfServiceRequestTypeID is TL type id of AcceptTermsOfServiceRequest.
const AcceptTermsOfServiceRequestTypeID = 0x7efe03e4

// Ensuring interfaces in compile-time for AcceptTermsOfServiceRequest.
var (
	_ bin.Encoder     = &AcceptTermsOfServiceRequest{}
	_ bin.Decoder     = &AcceptTermsOfServiceRequest{}
	_ bin.BareEncoder = &AcceptTermsOfServiceRequest{}
	_ bin.BareDecoder = &AcceptTermsOfServiceRequest{}
)

func (a *AcceptTermsOfServiceRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.TermsOfServiceID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AcceptTermsOfServiceRequest) String() string {
	if a == nil {
		return "AcceptTermsOfServiceRequest(nil)"
	}
	type Alias AcceptTermsOfServiceRequest
	return fmt.Sprintf("AcceptTermsOfServiceRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AcceptTermsOfServiceRequest) TypeID() uint32 {
	return AcceptTermsOfServiceRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AcceptTermsOfServiceRequest) TypeName() string {
	return "acceptTermsOfService"
}

// TypeInfo returns info about TL type.
func (a *AcceptTermsOfServiceRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "acceptTermsOfService",
		ID:   AcceptTermsOfServiceRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TermsOfServiceID",
			SchemaName: "terms_of_service_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AcceptTermsOfServiceRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode acceptTermsOfService#7efe03e4 as nil")
	}
	b.PutID(AcceptTermsOfServiceRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AcceptTermsOfServiceRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode acceptTermsOfService#7efe03e4 as nil")
	}
	b.PutString(a.TermsOfServiceID)
	return nil
}

// Decode implements bin.Decoder.
func (a *AcceptTermsOfServiceRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode acceptTermsOfService#7efe03e4 to nil")
	}
	if err := b.ConsumeID(AcceptTermsOfServiceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode acceptTermsOfService#7efe03e4: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AcceptTermsOfServiceRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode acceptTermsOfService#7efe03e4 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode acceptTermsOfService#7efe03e4: field terms_of_service_id: %w", err)
		}
		a.TermsOfServiceID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AcceptTermsOfServiceRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode acceptTermsOfService#7efe03e4 as nil")
	}
	b.ObjStart()
	b.PutID("acceptTermsOfService")
	b.Comma()
	b.FieldStart("terms_of_service_id")
	b.PutString(a.TermsOfServiceID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AcceptTermsOfServiceRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode acceptTermsOfService#7efe03e4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("acceptTermsOfService"); err != nil {
				return fmt.Errorf("unable to decode acceptTermsOfService#7efe03e4: %w", err)
			}
		case "terms_of_service_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode acceptTermsOfService#7efe03e4: field terms_of_service_id: %w", err)
			}
			a.TermsOfServiceID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTermsOfServiceID returns value of TermsOfServiceID field.
func (a *AcceptTermsOfServiceRequest) GetTermsOfServiceID() (value string) {
	if a == nil {
		return
	}
	return a.TermsOfServiceID
}

// AcceptTermsOfService invokes method acceptTermsOfService#7efe03e4 returning error if any.
func (c *Client) AcceptTermsOfService(ctx context.Context, termsofserviceid string) error {
	var ok Ok

	request := &AcceptTermsOfServiceRequest{
		TermsOfServiceID: termsofserviceid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
