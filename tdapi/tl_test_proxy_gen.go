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

// TestProxyRequest represents TL type `testProxy#b8a1a29e`.
type TestProxyRequest struct {
	// Proxy server IP address
	Server string
	// Proxy server port
	Port int32
	// Proxy type
	Type ProxyTypeClass
	// Identifier of a datacenter with which to test connection
	DCID int32
	// The maximum overall timeout for the request
	Timeout float64
}

// TestProxyRequestTypeID is TL type id of TestProxyRequest.
const TestProxyRequestTypeID = 0xb8a1a29e

// Ensuring interfaces in compile-time for TestProxyRequest.
var (
	_ bin.Encoder     = &TestProxyRequest{}
	_ bin.Decoder     = &TestProxyRequest{}
	_ bin.BareEncoder = &TestProxyRequest{}
	_ bin.BareDecoder = &TestProxyRequest{}
)

func (t *TestProxyRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Server == "") {
		return false
	}
	if !(t.Port == 0) {
		return false
	}
	if !(t.Type == nil) {
		return false
	}
	if !(t.DCID == 0) {
		return false
	}
	if !(t.Timeout == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestProxyRequest) String() string {
	if t == nil {
		return "TestProxyRequest(nil)"
	}
	type Alias TestProxyRequest
	return fmt.Sprintf("TestProxyRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestProxyRequest) TypeID() uint32 {
	return TestProxyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TestProxyRequest) TypeName() string {
	return "testProxy"
}

// TypeInfo returns info about TL type.
func (t *TestProxyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testProxy",
		ID:   TestProxyRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Server",
			SchemaName: "server",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "Timeout",
			SchemaName: "timeout",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestProxyRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testProxy#b8a1a29e as nil")
	}
	b.PutID(TestProxyRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestProxyRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testProxy#b8a1a29e as nil")
	}
	b.PutString(t.Server)
	b.PutInt32(t.Port)
	if t.Type == nil {
		return fmt.Errorf("unable to encode testProxy#b8a1a29e: field type is nil")
	}
	if err := t.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode testProxy#b8a1a29e: field type: %w", err)
	}
	b.PutInt32(t.DCID)
	b.PutDouble(t.Timeout)
	return nil
}

// Decode implements bin.Decoder.
func (t *TestProxyRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testProxy#b8a1a29e to nil")
	}
	if err := b.ConsumeID(TestProxyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode testProxy#b8a1a29e: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestProxyRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testProxy#b8a1a29e to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode testProxy#b8a1a29e: field server: %w", err)
		}
		t.Server = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode testProxy#b8a1a29e: field port: %w", err)
		}
		t.Port = value
	}
	{
		value, err := DecodeProxyType(b)
		if err != nil {
			return fmt.Errorf("unable to decode testProxy#b8a1a29e: field type: %w", err)
		}
		t.Type = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode testProxy#b8a1a29e: field dc_id: %w", err)
		}
		t.DCID = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode testProxy#b8a1a29e: field timeout: %w", err)
		}
		t.Timeout = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TestProxyRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode testProxy#b8a1a29e as nil")
	}
	b.ObjStart()
	b.PutID("testProxy")
	b.Comma()
	b.FieldStart("server")
	b.PutString(t.Server)
	b.Comma()
	b.FieldStart("port")
	b.PutInt32(t.Port)
	b.Comma()
	b.FieldStart("type")
	if t.Type == nil {
		return fmt.Errorf("unable to encode testProxy#b8a1a29e: field type is nil")
	}
	if err := t.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode testProxy#b8a1a29e: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("dc_id")
	b.PutInt32(t.DCID)
	b.Comma()
	b.FieldStart("timeout")
	b.PutDouble(t.Timeout)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TestProxyRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode testProxy#b8a1a29e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("testProxy"); err != nil {
				return fmt.Errorf("unable to decode testProxy#b8a1a29e: %w", err)
			}
		case "server":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode testProxy#b8a1a29e: field server: %w", err)
			}
			t.Server = value
		case "port":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode testProxy#b8a1a29e: field port: %w", err)
			}
			t.Port = value
		case "type":
			value, err := DecodeTDLibJSONProxyType(b)
			if err != nil {
				return fmt.Errorf("unable to decode testProxy#b8a1a29e: field type: %w", err)
			}
			t.Type = value
		case "dc_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode testProxy#b8a1a29e: field dc_id: %w", err)
			}
			t.DCID = value
		case "timeout":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode testProxy#b8a1a29e: field timeout: %w", err)
			}
			t.Timeout = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetServer returns value of Server field.
func (t *TestProxyRequest) GetServer() (value string) {
	if t == nil {
		return
	}
	return t.Server
}

// GetPort returns value of Port field.
func (t *TestProxyRequest) GetPort() (value int32) {
	if t == nil {
		return
	}
	return t.Port
}

// GetType returns value of Type field.
func (t *TestProxyRequest) GetType() (value ProxyTypeClass) {
	if t == nil {
		return
	}
	return t.Type
}

// GetDCID returns value of DCID field.
func (t *TestProxyRequest) GetDCID() (value int32) {
	if t == nil {
		return
	}
	return t.DCID
}

// GetTimeout returns value of Timeout field.
func (t *TestProxyRequest) GetTimeout() (value float64) {
	if t == nil {
		return
	}
	return t.Timeout
}

// TestProxy invokes method testProxy#b8a1a29e returning error if any.
func (c *Client) TestProxy(ctx context.Context, request *TestProxyRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
