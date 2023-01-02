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

// InlineBotSwitchPM represents TL type `inlineBotSwitchPM#3c20629f`.
// The bot requested the user to message them in private
//
// See https://core.telegram.org/constructor/inlineBotSwitchPM for reference.
type InlineBotSwitchPM struct {
	// Text for the button that switches the user to a private chat with the bot and sends
	// the bot a start message with the parameter start_parameter (can be empty)
	Text string
	// The parameter for the /start parameter
	StartParam string
}

// InlineBotSwitchPMTypeID is TL type id of InlineBotSwitchPM.
const InlineBotSwitchPMTypeID = 0x3c20629f

// Ensuring interfaces in compile-time for InlineBotSwitchPM.
var (
	_ bin.Encoder     = &InlineBotSwitchPM{}
	_ bin.Decoder     = &InlineBotSwitchPM{}
	_ bin.BareEncoder = &InlineBotSwitchPM{}
	_ bin.BareDecoder = &InlineBotSwitchPM{}
)

func (i *InlineBotSwitchPM) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Text == "") {
		return false
	}
	if !(i.StartParam == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineBotSwitchPM) String() string {
	if i == nil {
		return "InlineBotSwitchPM(nil)"
	}
	type Alias InlineBotSwitchPM
	return fmt.Sprintf("InlineBotSwitchPM%+v", Alias(*i))
}

// FillFrom fills InlineBotSwitchPM from given interface.
func (i *InlineBotSwitchPM) FillFrom(from interface {
	GetText() (value string)
	GetStartParam() (value string)
}) {
	i.Text = from.GetText()
	i.StartParam = from.GetStartParam()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineBotSwitchPM) TypeID() uint32 {
	return InlineBotSwitchPMTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineBotSwitchPM) TypeName() string {
	return "inlineBotSwitchPM"
}

// TypeInfo returns info about TL type.
func (i *InlineBotSwitchPM) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineBotSwitchPM",
		ID:   InlineBotSwitchPMTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "StartParam",
			SchemaName: "start_param",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineBotSwitchPM) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineBotSwitchPM#3c20629f as nil")
	}
	b.PutID(InlineBotSwitchPMTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineBotSwitchPM) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineBotSwitchPM#3c20629f as nil")
	}
	b.PutString(i.Text)
	b.PutString(i.StartParam)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineBotSwitchPM) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineBotSwitchPM#3c20629f to nil")
	}
	if err := b.ConsumeID(InlineBotSwitchPMTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineBotSwitchPM#3c20629f: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineBotSwitchPM) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineBotSwitchPM#3c20629f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineBotSwitchPM#3c20629f: field text: %w", err)
		}
		i.Text = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineBotSwitchPM#3c20629f: field start_param: %w", err)
		}
		i.StartParam = value
	}
	return nil
}

// GetText returns value of Text field.
func (i *InlineBotSwitchPM) GetText() (value string) {
	if i == nil {
		return
	}
	return i.Text
}

// GetStartParam returns value of StartParam field.
func (i *InlineBotSwitchPM) GetStartParam() (value string) {
	if i == nil {
		return
	}
	return i.StartParam
}
