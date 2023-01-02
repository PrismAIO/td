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

// SetClientDHParamsRequest represents TL type `set_client_DH_params#f5045f1f`.
type SetClientDHParamsRequest struct {
	// Nonce field of SetClientDHParamsRequest.
	Nonce bin.Int128
	// ServerNonce field of SetClientDHParamsRequest.
	ServerNonce bin.Int128
	// EncryptedData field of SetClientDHParamsRequest.
	EncryptedData []byte
}

// SetClientDHParamsRequestTypeID is TL type id of SetClientDHParamsRequest.
const SetClientDHParamsRequestTypeID = 0xf5045f1f

// Ensuring interfaces in compile-time for SetClientDHParamsRequest.
var (
	_ bin.Encoder     = &SetClientDHParamsRequest{}
	_ bin.Decoder     = &SetClientDHParamsRequest{}
	_ bin.BareEncoder = &SetClientDHParamsRequest{}
	_ bin.BareDecoder = &SetClientDHParamsRequest{}
)

func (s *SetClientDHParamsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Nonce == bin.Int128{}) {
		return false
	}
	if !(s.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(s.EncryptedData == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetClientDHParamsRequest) String() string {
	if s == nil {
		return "SetClientDHParamsRequest(nil)"
	}
	type Alias SetClientDHParamsRequest
	return fmt.Sprintf("SetClientDHParamsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetClientDHParamsRequest) TypeID() uint32 {
	return SetClientDHParamsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetClientDHParamsRequest) TypeName() string {
	return "set_client_DH_params"
}

// TypeInfo returns info about TL type.
func (s *SetClientDHParamsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "set_client_DH_params",
		ID:   SetClientDHParamsRequestTypeID,
	}
	if s == nil {
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
			Name:       "EncryptedData",
			SchemaName: "encrypted_data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetClientDHParamsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode set_client_DH_params#f5045f1f as nil")
	}
	b.PutID(SetClientDHParamsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetClientDHParamsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode set_client_DH_params#f5045f1f as nil")
	}
	b.PutInt128(s.Nonce)
	b.PutInt128(s.ServerNonce)
	b.PutBytes(s.EncryptedData)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetClientDHParamsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode set_client_DH_params#f5045f1f to nil")
	}
	if err := b.ConsumeID(SetClientDHParamsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode set_client_DH_params#f5045f1f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetClientDHParamsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode set_client_DH_params#f5045f1f to nil")
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode set_client_DH_params#f5045f1f: field nonce: %w", err)
		}
		s.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode set_client_DH_params#f5045f1f: field server_nonce: %w", err)
		}
		s.ServerNonce = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode set_client_DH_params#f5045f1f: field encrypted_data: %w", err)
		}
		s.EncryptedData = value
	}
	return nil
}

// GetNonce returns value of Nonce field.
func (s *SetClientDHParamsRequest) GetNonce() (value bin.Int128) {
	if s == nil {
		return
	}
	return s.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (s *SetClientDHParamsRequest) GetServerNonce() (value bin.Int128) {
	if s == nil {
		return
	}
	return s.ServerNonce
}

// GetEncryptedData returns value of EncryptedData field.
func (s *SetClientDHParamsRequest) GetEncryptedData() (value []byte) {
	if s == nil {
		return
	}
	return s.EncryptedData
}
