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

// AccountPasswordSettings represents TL type `account.passwordSettings#9a5c33e5`.
// Private info associated to the password info (recovery email, telegram passport¹ info
// & so on)
//
// Links:
//  1. https://core.telegram.org/passport
//
// See https://core.telegram.org/constructor/account.passwordSettings for reference.
type AccountPasswordSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// 2FA Recovery email¹
	//
	// Links:
	//  1) https://core.telegram.org/api/srp#email-verification
	//
	// Use SetEmail and GetEmail helpers.
	Email string
	// Telegram passport¹ settings
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetSecureSettings and GetSecureSettings helpers.
	SecureSettings SecureSecretSettings
}

// AccountPasswordSettingsTypeID is TL type id of AccountPasswordSettings.
const AccountPasswordSettingsTypeID = 0x9a5c33e5

// Ensuring interfaces in compile-time for AccountPasswordSettings.
var (
	_ bin.Encoder     = &AccountPasswordSettings{}
	_ bin.Decoder     = &AccountPasswordSettings{}
	_ bin.BareEncoder = &AccountPasswordSettings{}
	_ bin.BareDecoder = &AccountPasswordSettings{}
)

func (p *AccountPasswordSettings) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Email == "") {
		return false
	}
	if !(p.SecureSettings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *AccountPasswordSettings) String() string {
	if p == nil {
		return "AccountPasswordSettings(nil)"
	}
	type Alias AccountPasswordSettings
	return fmt.Sprintf("AccountPasswordSettings%+v", Alias(*p))
}

// FillFrom fills AccountPasswordSettings from given interface.
func (p *AccountPasswordSettings) FillFrom(from interface {
	GetEmail() (value string, ok bool)
	GetSecureSettings() (value SecureSecretSettings, ok bool)
}) {
	if val, ok := from.GetEmail(); ok {
		p.Email = val
	}

	if val, ok := from.GetSecureSettings(); ok {
		p.SecureSettings = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountPasswordSettings) TypeID() uint32 {
	return AccountPasswordSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountPasswordSettings) TypeName() string {
	return "account.passwordSettings"
}

// TypeInfo returns info about TL type.
func (p *AccountPasswordSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.passwordSettings",
		ID:   AccountPasswordSettingsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Email",
			SchemaName: "email",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "SecureSettings",
			SchemaName: "secure_settings",
			Null:       !p.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (p *AccountPasswordSettings) SetFlags() {
	if !(p.Email == "") {
		p.Flags.Set(0)
	}
	if !(p.SecureSettings.Zero()) {
		p.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (p *AccountPasswordSettings) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode account.passwordSettings#9a5c33e5 as nil")
	}
	b.PutID(AccountPasswordSettingsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *AccountPasswordSettings) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode account.passwordSettings#9a5c33e5 as nil")
	}
	p.SetFlags()
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.passwordSettings#9a5c33e5: field flags: %w", err)
	}
	if p.Flags.Has(0) {
		b.PutString(p.Email)
	}
	if p.Flags.Has(1) {
		if err := p.SecureSettings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.passwordSettings#9a5c33e5: field secure_settings: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *AccountPasswordSettings) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode account.passwordSettings#9a5c33e5 to nil")
	}
	if err := b.ConsumeID(AccountPasswordSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode account.passwordSettings#9a5c33e5: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *AccountPasswordSettings) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode account.passwordSettings#9a5c33e5 to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.passwordSettings#9a5c33e5: field flags: %w", err)
		}
	}
	if p.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.passwordSettings#9a5c33e5: field email: %w", err)
		}
		p.Email = value
	}
	if p.Flags.Has(1) {
		if err := p.SecureSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.passwordSettings#9a5c33e5: field secure_settings: %w", err)
		}
	}
	return nil
}

// SetEmail sets value of Email conditional field.
func (p *AccountPasswordSettings) SetEmail(value string) {
	p.Flags.Set(0)
	p.Email = value
}

// GetEmail returns value of Email conditional field and
// boolean which is true if field was set.
func (p *AccountPasswordSettings) GetEmail() (value string, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.Email, true
}

// SetSecureSettings sets value of SecureSettings conditional field.
func (p *AccountPasswordSettings) SetSecureSettings(value SecureSecretSettings) {
	p.Flags.Set(1)
	p.SecureSettings = value
}

// GetSecureSettings returns value of SecureSettings conditional field and
// boolean which is true if field was set.
func (p *AccountPasswordSettings) GetSecureSettings() (value SecureSecretSettings, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.SecureSettings, true
}
