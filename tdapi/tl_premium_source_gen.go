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

// PremiumSourceLimitExceeded represents TL type `premiumSourceLimitExceeded#85ae8702`.
type PremiumSourceLimitExceeded struct {
	// Type of the exceeded limit
	LimitType PremiumLimitTypeClass
}

// PremiumSourceLimitExceededTypeID is TL type id of PremiumSourceLimitExceeded.
const PremiumSourceLimitExceededTypeID = 0x85ae8702

// construct implements constructor of PremiumSourceClass.
func (p PremiumSourceLimitExceeded) construct() PremiumSourceClass { return &p }

// Ensuring interfaces in compile-time for PremiumSourceLimitExceeded.
var (
	_ bin.Encoder     = &PremiumSourceLimitExceeded{}
	_ bin.Decoder     = &PremiumSourceLimitExceeded{}
	_ bin.BareEncoder = &PremiumSourceLimitExceeded{}
	_ bin.BareDecoder = &PremiumSourceLimitExceeded{}

	_ PremiumSourceClass = &PremiumSourceLimitExceeded{}
)

func (p *PremiumSourceLimitExceeded) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.LimitType == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumSourceLimitExceeded) String() string {
	if p == nil {
		return "PremiumSourceLimitExceeded(nil)"
	}
	type Alias PremiumSourceLimitExceeded
	return fmt.Sprintf("PremiumSourceLimitExceeded%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumSourceLimitExceeded) TypeID() uint32 {
	return PremiumSourceLimitExceededTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumSourceLimitExceeded) TypeName() string {
	return "premiumSourceLimitExceeded"
}

// TypeInfo returns info about TL type.
func (p *PremiumSourceLimitExceeded) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumSourceLimitExceeded",
		ID:   PremiumSourceLimitExceededTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LimitType",
			SchemaName: "limit_type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumSourceLimitExceeded) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceLimitExceeded#85ae8702 as nil")
	}
	b.PutID(PremiumSourceLimitExceededTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumSourceLimitExceeded) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceLimitExceeded#85ae8702 as nil")
	}
	if p.LimitType == nil {
		return fmt.Errorf("unable to encode premiumSourceLimitExceeded#85ae8702: field limit_type is nil")
	}
	if err := p.LimitType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premiumSourceLimitExceeded#85ae8702: field limit_type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumSourceLimitExceeded) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceLimitExceeded#85ae8702 to nil")
	}
	if err := b.ConsumeID(PremiumSourceLimitExceededTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumSourceLimitExceeded#85ae8702: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumSourceLimitExceeded) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceLimitExceeded#85ae8702 to nil")
	}
	{
		value, err := DecodePremiumLimitType(b)
		if err != nil {
			return fmt.Errorf("unable to decode premiumSourceLimitExceeded#85ae8702: field limit_type: %w", err)
		}
		p.LimitType = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumSourceLimitExceeded) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceLimitExceeded#85ae8702 as nil")
	}
	b.ObjStart()
	b.PutID("premiumSourceLimitExceeded")
	b.Comma()
	b.FieldStart("limit_type")
	if p.LimitType == nil {
		return fmt.Errorf("unable to encode premiumSourceLimitExceeded#85ae8702: field limit_type is nil")
	}
	if err := p.LimitType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode premiumSourceLimitExceeded#85ae8702: field limit_type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumSourceLimitExceeded) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceLimitExceeded#85ae8702 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumSourceLimitExceeded"); err != nil {
				return fmt.Errorf("unable to decode premiumSourceLimitExceeded#85ae8702: %w", err)
			}
		case "limit_type":
			value, err := DecodeTDLibJSONPremiumLimitType(b)
			if err != nil {
				return fmt.Errorf("unable to decode premiumSourceLimitExceeded#85ae8702: field limit_type: %w", err)
			}
			p.LimitType = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLimitType returns value of LimitType field.
func (p *PremiumSourceLimitExceeded) GetLimitType() (value PremiumLimitTypeClass) {
	if p == nil {
		return
	}
	return p.LimitType
}

// PremiumSourceFeature represents TL type `premiumSourceFeature#1a929325`.
type PremiumSourceFeature struct {
	// The used feature
	Feature PremiumFeatureClass
}

// PremiumSourceFeatureTypeID is TL type id of PremiumSourceFeature.
const PremiumSourceFeatureTypeID = 0x1a929325

// construct implements constructor of PremiumSourceClass.
func (p PremiumSourceFeature) construct() PremiumSourceClass { return &p }

// Ensuring interfaces in compile-time for PremiumSourceFeature.
var (
	_ bin.Encoder     = &PremiumSourceFeature{}
	_ bin.Decoder     = &PremiumSourceFeature{}
	_ bin.BareEncoder = &PremiumSourceFeature{}
	_ bin.BareDecoder = &PremiumSourceFeature{}

	_ PremiumSourceClass = &PremiumSourceFeature{}
)

func (p *PremiumSourceFeature) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Feature == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumSourceFeature) String() string {
	if p == nil {
		return "PremiumSourceFeature(nil)"
	}
	type Alias PremiumSourceFeature
	return fmt.Sprintf("PremiumSourceFeature%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumSourceFeature) TypeID() uint32 {
	return PremiumSourceFeatureTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumSourceFeature) TypeName() string {
	return "premiumSourceFeature"
}

// TypeInfo returns info about TL type.
func (p *PremiumSourceFeature) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumSourceFeature",
		ID:   PremiumSourceFeatureTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Feature",
			SchemaName: "feature",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumSourceFeature) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceFeature#1a929325 as nil")
	}
	b.PutID(PremiumSourceFeatureTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumSourceFeature) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceFeature#1a929325 as nil")
	}
	if p.Feature == nil {
		return fmt.Errorf("unable to encode premiumSourceFeature#1a929325: field feature is nil")
	}
	if err := p.Feature.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premiumSourceFeature#1a929325: field feature: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumSourceFeature) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceFeature#1a929325 to nil")
	}
	if err := b.ConsumeID(PremiumSourceFeatureTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumSourceFeature#1a929325: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumSourceFeature) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceFeature#1a929325 to nil")
	}
	{
		value, err := DecodePremiumFeature(b)
		if err != nil {
			return fmt.Errorf("unable to decode premiumSourceFeature#1a929325: field feature: %w", err)
		}
		p.Feature = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumSourceFeature) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceFeature#1a929325 as nil")
	}
	b.ObjStart()
	b.PutID("premiumSourceFeature")
	b.Comma()
	b.FieldStart("feature")
	if p.Feature == nil {
		return fmt.Errorf("unable to encode premiumSourceFeature#1a929325: field feature is nil")
	}
	if err := p.Feature.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode premiumSourceFeature#1a929325: field feature: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumSourceFeature) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceFeature#1a929325 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumSourceFeature"); err != nil {
				return fmt.Errorf("unable to decode premiumSourceFeature#1a929325: %w", err)
			}
		case "feature":
			value, err := DecodeTDLibJSONPremiumFeature(b)
			if err != nil {
				return fmt.Errorf("unable to decode premiumSourceFeature#1a929325: field feature: %w", err)
			}
			p.Feature = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFeature returns value of Feature field.
func (p *PremiumSourceFeature) GetFeature() (value PremiumFeatureClass) {
	if p == nil {
		return
	}
	return p.Feature
}

// PremiumSourceLink represents TL type `premiumSourceLink#7f42999c`.
type PremiumSourceLink struct {
	// The referrer from the link
	Referrer string
}

// PremiumSourceLinkTypeID is TL type id of PremiumSourceLink.
const PremiumSourceLinkTypeID = 0x7f42999c

// construct implements constructor of PremiumSourceClass.
func (p PremiumSourceLink) construct() PremiumSourceClass { return &p }

// Ensuring interfaces in compile-time for PremiumSourceLink.
var (
	_ bin.Encoder     = &PremiumSourceLink{}
	_ bin.Decoder     = &PremiumSourceLink{}
	_ bin.BareEncoder = &PremiumSourceLink{}
	_ bin.BareDecoder = &PremiumSourceLink{}

	_ PremiumSourceClass = &PremiumSourceLink{}
)

func (p *PremiumSourceLink) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Referrer == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumSourceLink) String() string {
	if p == nil {
		return "PremiumSourceLink(nil)"
	}
	type Alias PremiumSourceLink
	return fmt.Sprintf("PremiumSourceLink%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumSourceLink) TypeID() uint32 {
	return PremiumSourceLinkTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumSourceLink) TypeName() string {
	return "premiumSourceLink"
}

// TypeInfo returns info about TL type.
func (p *PremiumSourceLink) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumSourceLink",
		ID:   PremiumSourceLinkTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Referrer",
			SchemaName: "referrer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumSourceLink) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceLink#7f42999c as nil")
	}
	b.PutID(PremiumSourceLinkTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumSourceLink) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceLink#7f42999c as nil")
	}
	b.PutString(p.Referrer)
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumSourceLink) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceLink#7f42999c to nil")
	}
	if err := b.ConsumeID(PremiumSourceLinkTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumSourceLink#7f42999c: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumSourceLink) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceLink#7f42999c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode premiumSourceLink#7f42999c: field referrer: %w", err)
		}
		p.Referrer = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumSourceLink) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceLink#7f42999c as nil")
	}
	b.ObjStart()
	b.PutID("premiumSourceLink")
	b.Comma()
	b.FieldStart("referrer")
	b.PutString(p.Referrer)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumSourceLink) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceLink#7f42999c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumSourceLink"); err != nil {
				return fmt.Errorf("unable to decode premiumSourceLink#7f42999c: %w", err)
			}
		case "referrer":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode premiumSourceLink#7f42999c: field referrer: %w", err)
			}
			p.Referrer = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetReferrer returns value of Referrer field.
func (p *PremiumSourceLink) GetReferrer() (value string) {
	if p == nil {
		return
	}
	return p.Referrer
}

// PremiumSourceSettings represents TL type `premiumSourceSettings#eef88535`.
type PremiumSourceSettings struct {
}

// PremiumSourceSettingsTypeID is TL type id of PremiumSourceSettings.
const PremiumSourceSettingsTypeID = 0xeef88535

// construct implements constructor of PremiumSourceClass.
func (p PremiumSourceSettings) construct() PremiumSourceClass { return &p }

// Ensuring interfaces in compile-time for PremiumSourceSettings.
var (
	_ bin.Encoder     = &PremiumSourceSettings{}
	_ bin.Decoder     = &PremiumSourceSettings{}
	_ bin.BareEncoder = &PremiumSourceSettings{}
	_ bin.BareDecoder = &PremiumSourceSettings{}

	_ PremiumSourceClass = &PremiumSourceSettings{}
)

func (p *PremiumSourceSettings) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumSourceSettings) String() string {
	if p == nil {
		return "PremiumSourceSettings(nil)"
	}
	type Alias PremiumSourceSettings
	return fmt.Sprintf("PremiumSourceSettings%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumSourceSettings) TypeID() uint32 {
	return PremiumSourceSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumSourceSettings) TypeName() string {
	return "premiumSourceSettings"
}

// TypeInfo returns info about TL type.
func (p *PremiumSourceSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumSourceSettings",
		ID:   PremiumSourceSettingsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumSourceSettings) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceSettings#eef88535 as nil")
	}
	b.PutID(PremiumSourceSettingsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumSourceSettings) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceSettings#eef88535 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumSourceSettings) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceSettings#eef88535 to nil")
	}
	if err := b.ConsumeID(PremiumSourceSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumSourceSettings#eef88535: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumSourceSettings) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceSettings#eef88535 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumSourceSettings) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumSourceSettings#eef88535 as nil")
	}
	b.ObjStart()
	b.PutID("premiumSourceSettings")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumSourceSettings) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumSourceSettings#eef88535 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumSourceSettings"); err != nil {
				return fmt.Errorf("unable to decode premiumSourceSettings#eef88535: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// PremiumSourceClassName is schema name of PremiumSourceClass.
const PremiumSourceClassName = "PremiumSource"

// PremiumSourceClass represents PremiumSource generic type.
//
// Example:
//
//	g, err := tdapi.DecodePremiumSource(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.PremiumSourceLimitExceeded: // premiumSourceLimitExceeded#85ae8702
//	case *tdapi.PremiumSourceFeature: // premiumSourceFeature#1a929325
//	case *tdapi.PremiumSourceLink: // premiumSourceLink#7f42999c
//	case *tdapi.PremiumSourceSettings: // premiumSourceSettings#eef88535
//	default: panic(v)
//	}
type PremiumSourceClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PremiumSourceClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodePremiumSource implements binary de-serialization for PremiumSourceClass.
func DecodePremiumSource(buf *bin.Buffer) (PremiumSourceClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PremiumSourceLimitExceededTypeID:
		// Decoding premiumSourceLimitExceeded#85ae8702.
		v := PremiumSourceLimitExceeded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumSourceClass: %w", err)
		}
		return &v, nil
	case PremiumSourceFeatureTypeID:
		// Decoding premiumSourceFeature#1a929325.
		v := PremiumSourceFeature{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumSourceClass: %w", err)
		}
		return &v, nil
	case PremiumSourceLinkTypeID:
		// Decoding premiumSourceLink#7f42999c.
		v := PremiumSourceLink{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumSourceClass: %w", err)
		}
		return &v, nil
	case PremiumSourceSettingsTypeID:
		// Decoding premiumSourceSettings#eef88535.
		v := PremiumSourceSettings{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumSourceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PremiumSourceClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONPremiumSource implements binary de-serialization for PremiumSourceClass.
func DecodeTDLibJSONPremiumSource(buf tdjson.Decoder) (PremiumSourceClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "premiumSourceLimitExceeded":
		// Decoding premiumSourceLimitExceeded#85ae8702.
		v := PremiumSourceLimitExceeded{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumSourceClass: %w", err)
		}
		return &v, nil
	case "premiumSourceFeature":
		// Decoding premiumSourceFeature#1a929325.
		v := PremiumSourceFeature{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumSourceClass: %w", err)
		}
		return &v, nil
	case "premiumSourceLink":
		// Decoding premiumSourceLink#7f42999c.
		v := PremiumSourceLink{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumSourceClass: %w", err)
		}
		return &v, nil
	case "premiumSourceSettings":
		// Decoding premiumSourceSettings#eef88535.
		v := PremiumSourceSettings{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PremiumSourceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PremiumSourceClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// PremiumSource boxes the PremiumSourceClass providing a helper.
type PremiumSourceBox struct {
	PremiumSource PremiumSourceClass
}

// Decode implements bin.Decoder for PremiumSourceBox.
func (b *PremiumSourceBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PremiumSourceBox to nil")
	}
	v, err := DecodePremiumSource(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PremiumSource = v
	return nil
}

// Encode implements bin.Encode for PremiumSourceBox.
func (b *PremiumSourceBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PremiumSource == nil {
		return fmt.Errorf("unable to encode PremiumSourceClass as nil")
	}
	return b.PremiumSource.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for PremiumSourceBox.
func (b *PremiumSourceBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode PremiumSourceBox to nil")
	}
	v, err := DecodeTDLibJSONPremiumSource(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PremiumSource = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for PremiumSourceBox.
func (b *PremiumSourceBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.PremiumSource == nil {
		return fmt.Errorf("unable to encode PremiumSourceClass as nil")
	}
	return b.PremiumSource.EncodeTDLibJSON(buf)
}
