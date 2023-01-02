// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// AccountChangeAuthorizationSettingsRequest represents TL type `account.changeAuthorizationSettings#40f48462`.
// Change authorization settings
//
// See https://core.telegram.org/method/account.changeAuthorizationSettings for reference.
type AccountChangeAuthorizationSettingsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Session ID from the authorization¹ constructor, fetchable using account
	// getAuthorizations²
	//
	// Links:
	//  1) https://core.telegram.org/constructor/authorization
	//  2) https://core.telegram.org/method/account.getAuthorizations
	Hash int64
	// Whether to enable or disable receiving encrypted chats: if the flag is not set, the
	// previous setting is not changed
	//
	// Use SetEncryptedRequestsDisabled and GetEncryptedRequestsDisabled helpers.
	EncryptedRequestsDisabled bool
	// Whether to enable or disable receiving calls: if the flag is not set, the previous
	// setting is not changed
	//
	// Use SetCallRequestsDisabled and GetCallRequestsDisabled helpers.
	CallRequestsDisabled bool
}

// AccountChangeAuthorizationSettingsRequestTypeID is TL type id of AccountChangeAuthorizationSettingsRequest.
const AccountChangeAuthorizationSettingsRequestTypeID = 0x40f48462

// Ensuring interfaces in compile-time for AccountChangeAuthorizationSettingsRequest.
var (
	_ bin.Encoder     = &AccountChangeAuthorizationSettingsRequest{}
	_ bin.Decoder     = &AccountChangeAuthorizationSettingsRequest{}
	_ bin.BareEncoder = &AccountChangeAuthorizationSettingsRequest{}
	_ bin.BareDecoder = &AccountChangeAuthorizationSettingsRequest{}
)

func (c *AccountChangeAuthorizationSettingsRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Hash == 0) {
		return false
	}
	if !(c.EncryptedRequestsDisabled == false) {
		return false
	}
	if !(c.CallRequestsDisabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountChangeAuthorizationSettingsRequest) String() string {
	if c == nil {
		return "AccountChangeAuthorizationSettingsRequest(nil)"
	}
	type Alias AccountChangeAuthorizationSettingsRequest
	return fmt.Sprintf("AccountChangeAuthorizationSettingsRequest%+v", Alias(*c))
}

// FillFrom fills AccountChangeAuthorizationSettingsRequest from given interface.
func (c *AccountChangeAuthorizationSettingsRequest) FillFrom(from interface {
	GetHash() (value int64)
	GetEncryptedRequestsDisabled() (value bool, ok bool)
	GetCallRequestsDisabled() (value bool, ok bool)
}) {
	c.Hash = from.GetHash()
	if val, ok := from.GetEncryptedRequestsDisabled(); ok {
		c.EncryptedRequestsDisabled = val
	}

	if val, ok := from.GetCallRequestsDisabled(); ok {
		c.CallRequestsDisabled = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountChangeAuthorizationSettingsRequest) TypeID() uint32 {
	return AccountChangeAuthorizationSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountChangeAuthorizationSettingsRequest) TypeName() string {
	return "account.changeAuthorizationSettings"
}

// TypeInfo returns info about TL type.
func (c *AccountChangeAuthorizationSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.changeAuthorizationSettings",
		ID:   AccountChangeAuthorizationSettingsRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "EncryptedRequestsDisabled",
			SchemaName: "encrypted_requests_disabled",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "CallRequestsDisabled",
			SchemaName: "call_requests_disabled",
			Null:       !c.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *AccountChangeAuthorizationSettingsRequest) SetFlags() {
	if !(c.EncryptedRequestsDisabled == false) {
		c.Flags.Set(0)
	}
	if !(c.CallRequestsDisabled == false) {
		c.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (c *AccountChangeAuthorizationSettingsRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.changeAuthorizationSettings#40f48462 as nil")
	}
	b.PutID(AccountChangeAuthorizationSettingsRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *AccountChangeAuthorizationSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.changeAuthorizationSettings#40f48462 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.changeAuthorizationSettings#40f48462: field flags: %w", err)
	}
	b.PutLong(c.Hash)
	if c.Flags.Has(0) {
		b.PutBool(c.EncryptedRequestsDisabled)
	}
	if c.Flags.Has(1) {
		b.PutBool(c.CallRequestsDisabled)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *AccountChangeAuthorizationSettingsRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.changeAuthorizationSettings#40f48462 to nil")
	}
	if err := b.ConsumeID(AccountChangeAuthorizationSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.changeAuthorizationSettings#40f48462: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *AccountChangeAuthorizationSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.changeAuthorizationSettings#40f48462 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.changeAuthorizationSettings#40f48462: field flags: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.changeAuthorizationSettings#40f48462: field hash: %w", err)
		}
		c.Hash = value
	}
	if c.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.changeAuthorizationSettings#40f48462: field encrypted_requests_disabled: %w", err)
		}
		c.EncryptedRequestsDisabled = value
	}
	if c.Flags.Has(1) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.changeAuthorizationSettings#40f48462: field call_requests_disabled: %w", err)
		}
		c.CallRequestsDisabled = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (c *AccountChangeAuthorizationSettingsRequest) GetHash() (value int64) {
	if c == nil {
		return
	}
	return c.Hash
}

// SetEncryptedRequestsDisabled sets value of EncryptedRequestsDisabled conditional field.
func (c *AccountChangeAuthorizationSettingsRequest) SetEncryptedRequestsDisabled(value bool) {
	c.Flags.Set(0)
	c.EncryptedRequestsDisabled = value
}

// GetEncryptedRequestsDisabled returns value of EncryptedRequestsDisabled conditional field and
// boolean which is true if field was set.
func (c *AccountChangeAuthorizationSettingsRequest) GetEncryptedRequestsDisabled() (value bool, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.EncryptedRequestsDisabled, true
}

// SetCallRequestsDisabled sets value of CallRequestsDisabled conditional field.
func (c *AccountChangeAuthorizationSettingsRequest) SetCallRequestsDisabled(value bool) {
	c.Flags.Set(1)
	c.CallRequestsDisabled = value
}

// GetCallRequestsDisabled returns value of CallRequestsDisabled conditional field and
// boolean which is true if field was set.
func (c *AccountChangeAuthorizationSettingsRequest) GetCallRequestsDisabled() (value bool, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.CallRequestsDisabled, true
}

// AccountChangeAuthorizationSettings invokes method account.changeAuthorizationSettings#40f48462 returning error if any.
// Change authorization settings
//
// Possible errors:
//
//	400 HASH_INVALID: The provided hash is invalid.
//
// See https://core.telegram.org/method/account.changeAuthorizationSettings for reference.
func (c *Client) AccountChangeAuthorizationSettings(ctx context.Context, request *AccountChangeAuthorizationSettingsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
