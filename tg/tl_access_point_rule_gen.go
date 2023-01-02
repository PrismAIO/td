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

// AccessPointRule represents TL type `accessPointRule#4679b65f`.
//
// See https://core.telegram.org/constructor/accessPointRule for reference.
type AccessPointRule struct {
	// PhonePrefixRules field of AccessPointRule.
	PhonePrefixRules string
	// DCID field of AccessPointRule.
	DCID int
	// IPs field of AccessPointRule.
	IPs []IPPortClass
}

// AccessPointRuleTypeID is TL type id of AccessPointRule.
const AccessPointRuleTypeID = 0x4679b65f

// Ensuring interfaces in compile-time for AccessPointRule.
var (
	_ bin.Encoder     = &AccessPointRule{}
	_ bin.Decoder     = &AccessPointRule{}
	_ bin.BareEncoder = &AccessPointRule{}
	_ bin.BareDecoder = &AccessPointRule{}
)

func (a *AccessPointRule) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.PhonePrefixRules == "") {
		return false
	}
	if !(a.DCID == 0) {
		return false
	}
	if !(a.IPs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AccessPointRule) String() string {
	if a == nil {
		return "AccessPointRule(nil)"
	}
	type Alias AccessPointRule
	return fmt.Sprintf("AccessPointRule%+v", Alias(*a))
}

// FillFrom fills AccessPointRule from given interface.
func (a *AccessPointRule) FillFrom(from interface {
	GetPhonePrefixRules() (value string)
	GetDCID() (value int)
	GetIPs() (value []IPPortClass)
}) {
	a.PhonePrefixRules = from.GetPhonePrefixRules()
	a.DCID = from.GetDCID()
	a.IPs = from.GetIPs()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccessPointRule) TypeID() uint32 {
	return AccessPointRuleTypeID
}

// TypeName returns name of type in TL schema.
func (*AccessPointRule) TypeName() string {
	return "accessPointRule"
}

// TypeInfo returns info about TL type.
func (a *AccessPointRule) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "accessPointRule",
		ID:   AccessPointRuleTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhonePrefixRules",
			SchemaName: "phone_prefix_rules",
		},
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "IPs",
			SchemaName: "ips",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AccessPointRule) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode accessPointRule#4679b65f as nil")
	}
	b.PutID(AccessPointRuleTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AccessPointRule) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode accessPointRule#4679b65f as nil")
	}
	b.PutString(a.PhonePrefixRules)
	b.PutInt(a.DCID)
	b.PutInt(len(a.IPs))
	for idx, v := range a.IPs {
		if v == nil {
			return fmt.Errorf("unable to encode accessPointRule#4679b65f: field ips element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare accessPointRule#4679b65f: field ips element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AccessPointRule) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode accessPointRule#4679b65f to nil")
	}
	if err := b.ConsumeID(AccessPointRuleTypeID); err != nil {
		return fmt.Errorf("unable to decode accessPointRule#4679b65f: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AccessPointRule) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode accessPointRule#4679b65f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode accessPointRule#4679b65f: field phone_prefix_rules: %w", err)
		}
		a.PhonePrefixRules = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode accessPointRule#4679b65f: field dc_id: %w", err)
		}
		a.DCID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode accessPointRule#4679b65f: field ips: %w", err)
		}

		if headerLen > 0 {
			a.IPs = make([]IPPortClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeIPPort(b)
			if err != nil {
				return fmt.Errorf("unable to decode accessPointRule#4679b65f: field ips: %w", err)
			}
			a.IPs = append(a.IPs, value)
		}
	}
	return nil
}

// GetPhonePrefixRules returns value of PhonePrefixRules field.
func (a *AccessPointRule) GetPhonePrefixRules() (value string) {
	if a == nil {
		return
	}
	return a.PhonePrefixRules
}

// GetDCID returns value of DCID field.
func (a *AccessPointRule) GetDCID() (value int) {
	if a == nil {
		return
	}
	return a.DCID
}

// GetIPs returns value of IPs field.
func (a *AccessPointRule) GetIPs() (value []IPPortClass) {
	if a == nil {
		return
	}
	return a.IPs
}

// MapIPs returns field IPs wrapped in IPPortClassArray helper.
func (a *AccessPointRule) MapIPs() (value IPPortClassArray) {
	return IPPortClassArray(a.IPs)
}
