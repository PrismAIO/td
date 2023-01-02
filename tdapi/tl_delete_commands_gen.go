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

// DeleteCommandsRequest represents TL type `deleteCommands#3bc47c2a`.
type DeleteCommandsRequest struct {
	// The scope to which the commands are relevant; pass null to delete commands in the
	// default bot command scope
	Scope BotCommandScopeClass
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string
}

// DeleteCommandsRequestTypeID is TL type id of DeleteCommandsRequest.
const DeleteCommandsRequestTypeID = 0x3bc47c2a

// Ensuring interfaces in compile-time for DeleteCommandsRequest.
var (
	_ bin.Encoder     = &DeleteCommandsRequest{}
	_ bin.Decoder     = &DeleteCommandsRequest{}
	_ bin.BareEncoder = &DeleteCommandsRequest{}
	_ bin.BareDecoder = &DeleteCommandsRequest{}
)

func (d *DeleteCommandsRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Scope == nil) {
		return false
	}
	if !(d.LanguageCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteCommandsRequest) String() string {
	if d == nil {
		return "DeleteCommandsRequest(nil)"
	}
	type Alias DeleteCommandsRequest
	return fmt.Sprintf("DeleteCommandsRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteCommandsRequest) TypeID() uint32 {
	return DeleteCommandsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteCommandsRequest) TypeName() string {
	return "deleteCommands"
}

// TypeInfo returns info about TL type.
func (d *DeleteCommandsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteCommands",
		ID:   DeleteCommandsRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Scope",
			SchemaName: "scope",
		},
		{
			Name:       "LanguageCode",
			SchemaName: "language_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteCommandsRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteCommands#3bc47c2a as nil")
	}
	b.PutID(DeleteCommandsRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteCommandsRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteCommands#3bc47c2a as nil")
	}
	if d.Scope == nil {
		return fmt.Errorf("unable to encode deleteCommands#3bc47c2a: field scope is nil")
	}
	if err := d.Scope.Encode(b); err != nil {
		return fmt.Errorf("unable to encode deleteCommands#3bc47c2a: field scope: %w", err)
	}
	b.PutString(d.LanguageCode)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteCommandsRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteCommands#3bc47c2a to nil")
	}
	if err := b.ConsumeID(DeleteCommandsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteCommands#3bc47c2a: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteCommandsRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteCommands#3bc47c2a to nil")
	}
	{
		value, err := DecodeBotCommandScope(b)
		if err != nil {
			return fmt.Errorf("unable to decode deleteCommands#3bc47c2a: field scope: %w", err)
		}
		d.Scope = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode deleteCommands#3bc47c2a: field language_code: %w", err)
		}
		d.LanguageCode = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteCommandsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteCommands#3bc47c2a as nil")
	}
	b.ObjStart()
	b.PutID("deleteCommands")
	b.Comma()
	b.FieldStart("scope")
	if d.Scope == nil {
		return fmt.Errorf("unable to encode deleteCommands#3bc47c2a: field scope is nil")
	}
	if err := d.Scope.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode deleteCommands#3bc47c2a: field scope: %w", err)
	}
	b.Comma()
	b.FieldStart("language_code")
	b.PutString(d.LanguageCode)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteCommandsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteCommands#3bc47c2a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteCommands"); err != nil {
				return fmt.Errorf("unable to decode deleteCommands#3bc47c2a: %w", err)
			}
		case "scope":
			value, err := DecodeTDLibJSONBotCommandScope(b)
			if err != nil {
				return fmt.Errorf("unable to decode deleteCommands#3bc47c2a: field scope: %w", err)
			}
			d.Scope = value
		case "language_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode deleteCommands#3bc47c2a: field language_code: %w", err)
			}
			d.LanguageCode = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetScope returns value of Scope field.
func (d *DeleteCommandsRequest) GetScope() (value BotCommandScopeClass) {
	if d == nil {
		return
	}
	return d.Scope
}

// GetLanguageCode returns value of LanguageCode field.
func (d *DeleteCommandsRequest) GetLanguageCode() (value string) {
	if d == nil {
		return
	}
	return d.LanguageCode
}

// DeleteCommands invokes method deleteCommands#3bc47c2a returning error if any.
func (c *Client) DeleteCommands(ctx context.Context, request *DeleteCommandsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
