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

// MessagesRequestSimpleWebViewRequest represents TL type `messages.requestSimpleWebView#299bec8e`.
// Open a bot web app¹.
//
// Links:
//  1. https://core.telegram.org/api/bots/webapps
//
// See https://core.telegram.org/method/messages.requestSimpleWebView for reference.
type MessagesRequestSimpleWebViewRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Bot that owns the webapp
	Bot InputUserClass
	// Web app URL
	URL string
	// Theme parameters
	//
	// Use SetThemeParams and GetThemeParams helpers.
	ThemeParams DataJSON
	// Platform field of MessagesRequestSimpleWebViewRequest.
	Platform string
}

// MessagesRequestSimpleWebViewRequestTypeID is TL type id of MessagesRequestSimpleWebViewRequest.
const MessagesRequestSimpleWebViewRequestTypeID = 0x299bec8e

// Ensuring interfaces in compile-time for MessagesRequestSimpleWebViewRequest.
var (
	_ bin.Encoder     = &MessagesRequestSimpleWebViewRequest{}
	_ bin.Decoder     = &MessagesRequestSimpleWebViewRequest{}
	_ bin.BareEncoder = &MessagesRequestSimpleWebViewRequest{}
	_ bin.BareDecoder = &MessagesRequestSimpleWebViewRequest{}
)

func (r *MessagesRequestSimpleWebViewRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Bot == nil) {
		return false
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.ThemeParams.Zero()) {
		return false
	}
	if !(r.Platform == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesRequestSimpleWebViewRequest) String() string {
	if r == nil {
		return "MessagesRequestSimpleWebViewRequest(nil)"
	}
	type Alias MessagesRequestSimpleWebViewRequest
	return fmt.Sprintf("MessagesRequestSimpleWebViewRequest%+v", Alias(*r))
}

// FillFrom fills MessagesRequestSimpleWebViewRequest from given interface.
func (r *MessagesRequestSimpleWebViewRequest) FillFrom(from interface {
	GetBot() (value InputUserClass)
	GetURL() (value string)
	GetThemeParams() (value DataJSON, ok bool)
	GetPlatform() (value string)
}) {
	r.Bot = from.GetBot()
	r.URL = from.GetURL()
	if val, ok := from.GetThemeParams(); ok {
		r.ThemeParams = val
	}

	r.Platform = from.GetPlatform()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesRequestSimpleWebViewRequest) TypeID() uint32 {
	return MessagesRequestSimpleWebViewRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesRequestSimpleWebViewRequest) TypeName() string {
	return "messages.requestSimpleWebView"
}

// TypeInfo returns info about TL type.
func (r *MessagesRequestSimpleWebViewRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.requestSimpleWebView",
		ID:   MessagesRequestSimpleWebViewRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "ThemeParams",
			SchemaName: "theme_params",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "Platform",
			SchemaName: "platform",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *MessagesRequestSimpleWebViewRequest) SetFlags() {
	if !(r.ThemeParams.Zero()) {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *MessagesRequestSimpleWebViewRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestSimpleWebView#299bec8e as nil")
	}
	b.PutID(MessagesRequestSimpleWebViewRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesRequestSimpleWebViewRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestSimpleWebView#299bec8e as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestSimpleWebView#299bec8e: field flags: %w", err)
	}
	if r.Bot == nil {
		return fmt.Errorf("unable to encode messages.requestSimpleWebView#299bec8e: field bot is nil")
	}
	if err := r.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestSimpleWebView#299bec8e: field bot: %w", err)
	}
	b.PutString(r.URL)
	if r.Flags.Has(0) {
		if err := r.ThemeParams.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.requestSimpleWebView#299bec8e: field theme_params: %w", err)
		}
	}
	b.PutString(r.Platform)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRequestSimpleWebViewRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestSimpleWebView#299bec8e to nil")
	}
	if err := b.ConsumeID(MessagesRequestSimpleWebViewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.requestSimpleWebView#299bec8e: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesRequestSimpleWebViewRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestSimpleWebView#299bec8e to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.requestSimpleWebView#299bec8e: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestSimpleWebView#299bec8e: field bot: %w", err)
		}
		r.Bot = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestSimpleWebView#299bec8e: field url: %w", err)
		}
		r.URL = value
	}
	if r.Flags.Has(0) {
		if err := r.ThemeParams.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.requestSimpleWebView#299bec8e: field theme_params: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestSimpleWebView#299bec8e: field platform: %w", err)
		}
		r.Platform = value
	}
	return nil
}

// GetBot returns value of Bot field.
func (r *MessagesRequestSimpleWebViewRequest) GetBot() (value InputUserClass) {
	if r == nil {
		return
	}
	return r.Bot
}

// GetURL returns value of URL field.
func (r *MessagesRequestSimpleWebViewRequest) GetURL() (value string) {
	if r == nil {
		return
	}
	return r.URL
}

// SetThemeParams sets value of ThemeParams conditional field.
func (r *MessagesRequestSimpleWebViewRequest) SetThemeParams(value DataJSON) {
	r.Flags.Set(0)
	r.ThemeParams = value
}

// GetThemeParams returns value of ThemeParams conditional field and
// boolean which is true if field was set.
func (r *MessagesRequestSimpleWebViewRequest) GetThemeParams() (value DataJSON, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.ThemeParams, true
}

// GetPlatform returns value of Platform field.
func (r *MessagesRequestSimpleWebViewRequest) GetPlatform() (value string) {
	if r == nil {
		return
	}
	return r.Platform
}

// MessagesRequestSimpleWebView invokes method messages.requestSimpleWebView#299bec8e returning error if any.
// Open a bot web app¹.
//
// Links:
//  1. https://core.telegram.org/api/bots/webapps
//
// See https://core.telegram.org/method/messages.requestSimpleWebView for reference.
func (c *Client) MessagesRequestSimpleWebView(ctx context.Context, request *MessagesRequestSimpleWebViewRequest) (*SimpleWebViewResultURL, error) {
	var result SimpleWebViewResultURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
