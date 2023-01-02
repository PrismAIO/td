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

// BotsResetBotCommandsRequest represents TL type `bots.resetBotCommands#3d8de0f9`.
// Clear bot commands for the specified bot scope and language code
//
// See https://core.telegram.org/method/bots.resetBotCommands for reference.
type BotsResetBotCommandsRequest struct {
	// Command scope
	Scope BotCommandScopeClass
	// Language code
	LangCode string
}

// BotsResetBotCommandsRequestTypeID is TL type id of BotsResetBotCommandsRequest.
const BotsResetBotCommandsRequestTypeID = 0x3d8de0f9

// Ensuring interfaces in compile-time for BotsResetBotCommandsRequest.
var (
	_ bin.Encoder     = &BotsResetBotCommandsRequest{}
	_ bin.Decoder     = &BotsResetBotCommandsRequest{}
	_ bin.BareEncoder = &BotsResetBotCommandsRequest{}
	_ bin.BareDecoder = &BotsResetBotCommandsRequest{}
)

func (r *BotsResetBotCommandsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Scope == nil) {
		return false
	}
	if !(r.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *BotsResetBotCommandsRequest) String() string {
	if r == nil {
		return "BotsResetBotCommandsRequest(nil)"
	}
	type Alias BotsResetBotCommandsRequest
	return fmt.Sprintf("BotsResetBotCommandsRequest%+v", Alias(*r))
}

// FillFrom fills BotsResetBotCommandsRequest from given interface.
func (r *BotsResetBotCommandsRequest) FillFrom(from interface {
	GetScope() (value BotCommandScopeClass)
	GetLangCode() (value string)
}) {
	r.Scope = from.GetScope()
	r.LangCode = from.GetLangCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsResetBotCommandsRequest) TypeID() uint32 {
	return BotsResetBotCommandsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsResetBotCommandsRequest) TypeName() string {
	return "bots.resetBotCommands"
}

// TypeInfo returns info about TL type.
func (r *BotsResetBotCommandsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.resetBotCommands",
		ID:   BotsResetBotCommandsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Scope",
			SchemaName: "scope",
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *BotsResetBotCommandsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode bots.resetBotCommands#3d8de0f9 as nil")
	}
	b.PutID(BotsResetBotCommandsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *BotsResetBotCommandsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode bots.resetBotCommands#3d8de0f9 as nil")
	}
	if r.Scope == nil {
		return fmt.Errorf("unable to encode bots.resetBotCommands#3d8de0f9: field scope is nil")
	}
	if err := r.Scope.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.resetBotCommands#3d8de0f9: field scope: %w", err)
	}
	b.PutString(r.LangCode)
	return nil
}

// Decode implements bin.Decoder.
func (r *BotsResetBotCommandsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode bots.resetBotCommands#3d8de0f9 to nil")
	}
	if err := b.ConsumeID(BotsResetBotCommandsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.resetBotCommands#3d8de0f9: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *BotsResetBotCommandsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode bots.resetBotCommands#3d8de0f9 to nil")
	}
	{
		value, err := DecodeBotCommandScope(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.resetBotCommands#3d8de0f9: field scope: %w", err)
		}
		r.Scope = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.resetBotCommands#3d8de0f9: field lang_code: %w", err)
		}
		r.LangCode = value
	}
	return nil
}

// GetScope returns value of Scope field.
func (r *BotsResetBotCommandsRequest) GetScope() (value BotCommandScopeClass) {
	if r == nil {
		return
	}
	return r.Scope
}

// GetLangCode returns value of LangCode field.
func (r *BotsResetBotCommandsRequest) GetLangCode() (value string) {
	if r == nil {
		return
	}
	return r.LangCode
}

// BotsResetBotCommands invokes method bots.resetBotCommands#3d8de0f9 returning error if any.
// Clear bot commands for the specified bot scope and language code
//
// Possible errors:
//
//	400 LANG_CODE_INVALID: The specified language code is invalid.
//
// See https://core.telegram.org/method/bots.resetBotCommands for reference.
// Can be used by bots.
func (c *Client) BotsResetBotCommands(ctx context.Context, request *BotsResetBotCommandsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
