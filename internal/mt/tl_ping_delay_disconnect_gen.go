// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// PingDelayDisconnectRequest represents TL type `ping_delay_disconnect#f3427b8c`.
type PingDelayDisconnectRequest struct {
	// PingID field of PingDelayDisconnectRequest.
	PingID int64
	// DisconnectDelay field of PingDelayDisconnectRequest.
	DisconnectDelay int
}

// PingDelayDisconnectRequestTypeID is TL type id of PingDelayDisconnectRequest.
const PingDelayDisconnectRequestTypeID = 0xf3427b8c

// Ensuring interfaces in compile-time for PingDelayDisconnectRequest.
var (
	_ bin.Encoder     = &PingDelayDisconnectRequest{}
	_ bin.Decoder     = &PingDelayDisconnectRequest{}
	_ bin.BareEncoder = &PingDelayDisconnectRequest{}
	_ bin.BareDecoder = &PingDelayDisconnectRequest{}
)

func (p *PingDelayDisconnectRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PingID == 0) {
		return false
	}
	if !(p.DisconnectDelay == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PingDelayDisconnectRequest) String() string {
	if p == nil {
		return "PingDelayDisconnectRequest(nil)"
	}
	type Alias PingDelayDisconnectRequest
	return fmt.Sprintf("PingDelayDisconnectRequest%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PingDelayDisconnectRequest) TypeID() uint32 {
	return PingDelayDisconnectRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PingDelayDisconnectRequest) TypeName() string {
	return "ping_delay_disconnect"
}

// TypeInfo returns info about TL type.
func (p *PingDelayDisconnectRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "ping_delay_disconnect",
		ID:   PingDelayDisconnectRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PingID",
			SchemaName: "ping_id",
		},
		{
			Name:       "DisconnectDelay",
			SchemaName: "disconnect_delay",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PingDelayDisconnectRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode ping_delay_disconnect#f3427b8c as nil")
	}
	b.PutID(PingDelayDisconnectRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PingDelayDisconnectRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode ping_delay_disconnect#f3427b8c as nil")
	}
	b.PutLong(p.PingID)
	b.PutInt(p.DisconnectDelay)
	return nil
}

// Decode implements bin.Decoder.
func (p *PingDelayDisconnectRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode ping_delay_disconnect#f3427b8c to nil")
	}
	if err := b.ConsumeID(PingDelayDisconnectRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode ping_delay_disconnect#f3427b8c: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PingDelayDisconnectRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode ping_delay_disconnect#f3427b8c to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode ping_delay_disconnect#f3427b8c: field ping_id: %w", err)
		}
		p.PingID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode ping_delay_disconnect#f3427b8c: field disconnect_delay: %w", err)
		}
		p.DisconnectDelay = value
	}
	return nil
}

// GetPingID returns value of PingID field.
func (p *PingDelayDisconnectRequest) GetPingID() (value int64) {
	if p == nil {
		return
	}
	return p.PingID
}

// GetDisconnectDelay returns value of DisconnectDelay field.
func (p *PingDelayDisconnectRequest) GetDisconnectDelay() (value int) {
	if p == nil {
		return
	}
	return p.DisconnectDelay
}
