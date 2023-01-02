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

// ChatTheme represents TL type `chatTheme#f9406c39`.
type ChatTheme struct {
	// Theme name
	Name string
	// Theme settings for a light chat theme
	LightSettings ThemeSettings
	// Theme settings for a dark chat theme
	DarkSettings ThemeSettings
}

// ChatThemeTypeID is TL type id of ChatTheme.
const ChatThemeTypeID = 0xf9406c39

// Ensuring interfaces in compile-time for ChatTheme.
var (
	_ bin.Encoder     = &ChatTheme{}
	_ bin.Decoder     = &ChatTheme{}
	_ bin.BareEncoder = &ChatTheme{}
	_ bin.BareDecoder = &ChatTheme{}
)

func (c *ChatTheme) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Name == "") {
		return false
	}
	if !(c.LightSettings.Zero()) {
		return false
	}
	if !(c.DarkSettings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatTheme) String() string {
	if c == nil {
		return "ChatTheme(nil)"
	}
	type Alias ChatTheme
	return fmt.Sprintf("ChatTheme%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatTheme) TypeID() uint32 {
	return ChatThemeTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatTheme) TypeName() string {
	return "chatTheme"
}

// TypeInfo returns info about TL type.
func (c *ChatTheme) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatTheme",
		ID:   ChatThemeTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "LightSettings",
			SchemaName: "light_settings",
		},
		{
			Name:       "DarkSettings",
			SchemaName: "dark_settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatTheme) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTheme#f9406c39 as nil")
	}
	b.PutID(ChatThemeTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatTheme) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTheme#f9406c39 as nil")
	}
	b.PutString(c.Name)
	if err := c.LightSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatTheme#f9406c39: field light_settings: %w", err)
	}
	if err := c.DarkSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatTheme#f9406c39: field dark_settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatTheme) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTheme#f9406c39 to nil")
	}
	if err := b.ConsumeID(ChatThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode chatTheme#f9406c39: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatTheme) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTheme#f9406c39 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatTheme#f9406c39: field name: %w", err)
		}
		c.Name = value
	}
	{
		if err := c.LightSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatTheme#f9406c39: field light_settings: %w", err)
		}
	}
	{
		if err := c.DarkSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatTheme#f9406c39: field dark_settings: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatTheme) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTheme#f9406c39 as nil")
	}
	b.ObjStart()
	b.PutID("chatTheme")
	b.Comma()
	b.FieldStart("name")
	b.PutString(c.Name)
	b.Comma()
	b.FieldStart("light_settings")
	if err := c.LightSettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatTheme#f9406c39: field light_settings: %w", err)
	}
	b.Comma()
	b.FieldStart("dark_settings")
	if err := c.DarkSettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatTheme#f9406c39: field dark_settings: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatTheme) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTheme#f9406c39 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatTheme"); err != nil {
				return fmt.Errorf("unable to decode chatTheme#f9406c39: %w", err)
			}
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatTheme#f9406c39: field name: %w", err)
			}
			c.Name = value
		case "light_settings":
			if err := c.LightSettings.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chatTheme#f9406c39: field light_settings: %w", err)
			}
		case "dark_settings":
			if err := c.DarkSettings.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chatTheme#f9406c39: field dark_settings: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetName returns value of Name field.
func (c *ChatTheme) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

// GetLightSettings returns value of LightSettings field.
func (c *ChatTheme) GetLightSettings() (value ThemeSettings) {
	if c == nil {
		return
	}
	return c.LightSettings
}

// GetDarkSettings returns value of DarkSettings field.
func (c *ChatTheme) GetDarkSettings() (value ThemeSettings) {
	if c == nil {
		return
	}
	return c.DarkSettings
}
