// Code generated by gotdgen, DO NOT EDIT.

package td

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

// ClientDHInnerData represents TL type `client_DH_inner_data#6643b654`.
//
// See https://localhost:80/doc/constructor/client_DH_inner_data for reference.
type ClientDHInnerData struct {
	// Nonce field of ClientDHInnerData.
	Nonce bin.Int128
	// ServerNonce field of ClientDHInnerData.
	ServerNonce bin.Int128
	// RetryID field of ClientDHInnerData.
	RetryID int64
	// GB field of ClientDHInnerData.
	GB string
}

// ClientDHInnerDataTypeID is TL type id of ClientDHInnerData.
const ClientDHInnerDataTypeID = 0x6643b654

// Ensuring interfaces in compile-time for ClientDHInnerData.
var (
	_ bin.Encoder     = &ClientDHInnerData{}
	_ bin.Decoder     = &ClientDHInnerData{}
	_ bin.BareEncoder = &ClientDHInnerData{}
	_ bin.BareDecoder = &ClientDHInnerData{}
)

func (c *ClientDHInnerData) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Nonce == bin.Int128{}) {
		return false
	}
	if !(c.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(c.RetryID == 0) {
		return false
	}
	if !(c.GB == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ClientDHInnerData) String() string {
	if c == nil {
		return "ClientDHInnerData(nil)"
	}
	type Alias ClientDHInnerData
	return fmt.Sprintf("ClientDHInnerData%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ClientDHInnerData) TypeID() uint32 {
	return ClientDHInnerDataTypeID
}

// TypeName returns name of type in TL schema.
func (*ClientDHInnerData) TypeName() string {
	return "client_DH_inner_data"
}

// TypeInfo returns info about TL type.
func (c *ClientDHInnerData) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "client_DH_inner_data",
		ID:   ClientDHInnerDataTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Nonce",
			SchemaName: "nonce",
		},
		{
			Name:       "ServerNonce",
			SchemaName: "server_nonce",
		},
		{
			Name:       "RetryID",
			SchemaName: "retry_id",
		},
		{
			Name:       "GB",
			SchemaName: "g_b",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ClientDHInnerData) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode client_DH_inner_data#6643b654 as nil")
	}
	b.PutID(ClientDHInnerDataTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ClientDHInnerData) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode client_DH_inner_data#6643b654 as nil")
	}
	b.PutInt128(c.Nonce)
	b.PutInt128(c.ServerNonce)
	b.PutLong(c.RetryID)
	b.PutString(c.GB)
	return nil
}

// Decode implements bin.Decoder.
func (c *ClientDHInnerData) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode client_DH_inner_data#6643b654 to nil")
	}
	if err := b.ConsumeID(ClientDHInnerDataTypeID); err != nil {
		return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ClientDHInnerData) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode client_DH_inner_data#6643b654 to nil")
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field nonce: %w", err)
		}
		c.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field server_nonce: %w", err)
		}
		c.ServerNonce = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field retry_id: %w", err)
		}
		c.RetryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field g_b: %w", err)
		}
		c.GB = value
	}
	return nil
}

// GetNonce returns value of Nonce field.
func (c *ClientDHInnerData) GetNonce() (value bin.Int128) {
	if c == nil {
		return
	}
	return c.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (c *ClientDHInnerData) GetServerNonce() (value bin.Int128) {
	if c == nil {
		return
	}
	return c.ServerNonce
}

// GetRetryID returns value of RetryID field.
func (c *ClientDHInnerData) GetRetryID() (value int64) {
	if c == nil {
		return
	}
	return c.RetryID
}

// GetGB returns value of GB field.
func (c *ClientDHInnerData) GetGB() (value string) {
	if c == nil {
		return
	}
	return c.GB
}
