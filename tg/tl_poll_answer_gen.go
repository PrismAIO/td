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

// PollAnswer represents TL type `pollAnswer#6ca9c2e9`.
// A possible answer of a poll
//
// See https://core.telegram.org/constructor/pollAnswer for reference.
type PollAnswer struct {
	// Textual representation of the answer
	Text string
	// The param that has to be passed to messages.sendVote¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.sendVote
	Option []byte
}

// PollAnswerTypeID is TL type id of PollAnswer.
const PollAnswerTypeID = 0x6ca9c2e9

// Ensuring interfaces in compile-time for PollAnswer.
var (
	_ bin.Encoder     = &PollAnswer{}
	_ bin.Decoder     = &PollAnswer{}
	_ bin.BareEncoder = &PollAnswer{}
	_ bin.BareDecoder = &PollAnswer{}
)

func (p *PollAnswer) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Text == "") {
		return false
	}
	if !(p.Option == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PollAnswer) String() string {
	if p == nil {
		return "PollAnswer(nil)"
	}
	type Alias PollAnswer
	return fmt.Sprintf("PollAnswer%+v", Alias(*p))
}

// FillFrom fills PollAnswer from given interface.
func (p *PollAnswer) FillFrom(from interface {
	GetText() (value string)
	GetOption() (value []byte)
}) {
	p.Text = from.GetText()
	p.Option = from.GetOption()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PollAnswer) TypeID() uint32 {
	return PollAnswerTypeID
}

// TypeName returns name of type in TL schema.
func (*PollAnswer) TypeName() string {
	return "pollAnswer"
}

// TypeInfo returns info about TL type.
func (p *PollAnswer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pollAnswer",
		ID:   PollAnswerTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "Option",
			SchemaName: "option",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PollAnswer) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollAnswer#6ca9c2e9 as nil")
	}
	b.PutID(PollAnswerTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PollAnswer) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollAnswer#6ca9c2e9 as nil")
	}
	b.PutString(p.Text)
	b.PutBytes(p.Option)
	return nil
}

// Decode implements bin.Decoder.
func (p *PollAnswer) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollAnswer#6ca9c2e9 to nil")
	}
	if err := b.ConsumeID(PollAnswerTypeID); err != nil {
		return fmt.Errorf("unable to decode pollAnswer#6ca9c2e9: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PollAnswer) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollAnswer#6ca9c2e9 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pollAnswer#6ca9c2e9: field text: %w", err)
		}
		p.Text = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode pollAnswer#6ca9c2e9: field option: %w", err)
		}
		p.Option = value
	}
	return nil
}

// GetText returns value of Text field.
func (p *PollAnswer) GetText() (value string) {
	if p == nil {
		return
	}
	return p.Text
}

// GetOption returns value of Option field.
func (p *PollAnswer) GetOption() (value []byte) {
	if p == nil {
		return
	}
	return p.Option
}
