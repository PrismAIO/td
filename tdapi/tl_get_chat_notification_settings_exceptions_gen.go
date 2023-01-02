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

// GetChatNotificationSettingsExceptionsRequest represents TL type `getChatNotificationSettingsExceptions#bfe0e11`.
type GetChatNotificationSettingsExceptionsRequest struct {
	// If specified, only chats from the scope will be returned; pass null to return chats
	// from all scopes
	Scope NotificationSettingsScopeClass
	// Pass true to include in the response chats with only non-default sound
	CompareSound bool
}

// GetChatNotificationSettingsExceptionsRequestTypeID is TL type id of GetChatNotificationSettingsExceptionsRequest.
const GetChatNotificationSettingsExceptionsRequestTypeID = 0xbfe0e11

// Ensuring interfaces in compile-time for GetChatNotificationSettingsExceptionsRequest.
var (
	_ bin.Encoder     = &GetChatNotificationSettingsExceptionsRequest{}
	_ bin.Decoder     = &GetChatNotificationSettingsExceptionsRequest{}
	_ bin.BareEncoder = &GetChatNotificationSettingsExceptionsRequest{}
	_ bin.BareDecoder = &GetChatNotificationSettingsExceptionsRequest{}
)

func (g *GetChatNotificationSettingsExceptionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Scope == nil) {
		return false
	}
	if !(g.CompareSound == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatNotificationSettingsExceptionsRequest) String() string {
	if g == nil {
		return "GetChatNotificationSettingsExceptionsRequest(nil)"
	}
	type Alias GetChatNotificationSettingsExceptionsRequest
	return fmt.Sprintf("GetChatNotificationSettingsExceptionsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatNotificationSettingsExceptionsRequest) TypeID() uint32 {
	return GetChatNotificationSettingsExceptionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatNotificationSettingsExceptionsRequest) TypeName() string {
	return "getChatNotificationSettingsExceptions"
}

// TypeInfo returns info about TL type.
func (g *GetChatNotificationSettingsExceptionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatNotificationSettingsExceptions",
		ID:   GetChatNotificationSettingsExceptionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Scope",
			SchemaName: "scope",
		},
		{
			Name:       "CompareSound",
			SchemaName: "compare_sound",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatNotificationSettingsExceptionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatNotificationSettingsExceptions#bfe0e11 as nil")
	}
	b.PutID(GetChatNotificationSettingsExceptionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatNotificationSettingsExceptionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatNotificationSettingsExceptions#bfe0e11 as nil")
	}
	if g.Scope == nil {
		return fmt.Errorf("unable to encode getChatNotificationSettingsExceptions#bfe0e11: field scope is nil")
	}
	if err := g.Scope.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatNotificationSettingsExceptions#bfe0e11: field scope: %w", err)
	}
	b.PutBool(g.CompareSound)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatNotificationSettingsExceptionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatNotificationSettingsExceptions#bfe0e11 to nil")
	}
	if err := b.ConsumeID(GetChatNotificationSettingsExceptionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatNotificationSettingsExceptions#bfe0e11: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatNotificationSettingsExceptionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatNotificationSettingsExceptions#bfe0e11 to nil")
	}
	{
		value, err := DecodeNotificationSettingsScope(b)
		if err != nil {
			return fmt.Errorf("unable to decode getChatNotificationSettingsExceptions#bfe0e11: field scope: %w", err)
		}
		g.Scope = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getChatNotificationSettingsExceptions#bfe0e11: field compare_sound: %w", err)
		}
		g.CompareSound = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatNotificationSettingsExceptionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatNotificationSettingsExceptions#bfe0e11 as nil")
	}
	b.ObjStart()
	b.PutID("getChatNotificationSettingsExceptions")
	b.Comma()
	b.FieldStart("scope")
	if g.Scope == nil {
		return fmt.Errorf("unable to encode getChatNotificationSettingsExceptions#bfe0e11: field scope is nil")
	}
	if err := g.Scope.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatNotificationSettingsExceptions#bfe0e11: field scope: %w", err)
	}
	b.Comma()
	b.FieldStart("compare_sound")
	b.PutBool(g.CompareSound)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatNotificationSettingsExceptionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatNotificationSettingsExceptions#bfe0e11 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatNotificationSettingsExceptions"); err != nil {
				return fmt.Errorf("unable to decode getChatNotificationSettingsExceptions#bfe0e11: %w", err)
			}
		case "scope":
			value, err := DecodeTDLibJSONNotificationSettingsScope(b)
			if err != nil {
				return fmt.Errorf("unable to decode getChatNotificationSettingsExceptions#bfe0e11: field scope: %w", err)
			}
			g.Scope = value
		case "compare_sound":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getChatNotificationSettingsExceptions#bfe0e11: field compare_sound: %w", err)
			}
			g.CompareSound = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetScope returns value of Scope field.
func (g *GetChatNotificationSettingsExceptionsRequest) GetScope() (value NotificationSettingsScopeClass) {
	if g == nil {
		return
	}
	return g.Scope
}

// GetCompareSound returns value of CompareSound field.
func (g *GetChatNotificationSettingsExceptionsRequest) GetCompareSound() (value bool) {
	if g == nil {
		return
	}
	return g.CompareSound
}

// GetChatNotificationSettingsExceptions invokes method getChatNotificationSettingsExceptions#bfe0e11 returning error if any.
func (c *Client) GetChatNotificationSettingsExceptions(ctx context.Context, request *GetChatNotificationSettingsExceptionsRequest) (*Chats, error) {
	var result Chats

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
