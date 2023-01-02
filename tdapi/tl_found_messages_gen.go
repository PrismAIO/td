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

// FoundMessages represents TL type `foundMessages#2515f708`.
type FoundMessages struct {
	// Approximate total number of messages found; -1 if unknown
	TotalCount int32
	// List of messages
	Messages []Message
	// The offset for the next request. If empty, there are no more results
	NextOffset string
}

// FoundMessagesTypeID is TL type id of FoundMessages.
const FoundMessagesTypeID = 0x2515f708

// Ensuring interfaces in compile-time for FoundMessages.
var (
	_ bin.Encoder     = &FoundMessages{}
	_ bin.Decoder     = &FoundMessages{}
	_ bin.BareEncoder = &FoundMessages{}
	_ bin.BareDecoder = &FoundMessages{}
)

func (f *FoundMessages) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.TotalCount == 0) {
		return false
	}
	if !(f.Messages == nil) {
		return false
	}
	if !(f.NextOffset == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FoundMessages) String() string {
	if f == nil {
		return "FoundMessages(nil)"
	}
	type Alias FoundMessages
	return fmt.Sprintf("FoundMessages%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FoundMessages) TypeID() uint32 {
	return FoundMessagesTypeID
}

// TypeName returns name of type in TL schema.
func (*FoundMessages) TypeName() string {
	return "foundMessages"
}

// TypeInfo returns info about TL type.
func (f *FoundMessages) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "foundMessages",
		ID:   FoundMessagesTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FoundMessages) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode foundMessages#2515f708 as nil")
	}
	b.PutID(FoundMessagesTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FoundMessages) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode foundMessages#2515f708 as nil")
	}
	b.PutInt32(f.TotalCount)
	b.PutInt(len(f.Messages))
	for idx, v := range f.Messages {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare foundMessages#2515f708: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutString(f.NextOffset)
	return nil
}

// Decode implements bin.Decoder.
func (f *FoundMessages) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode foundMessages#2515f708 to nil")
	}
	if err := b.ConsumeID(FoundMessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode foundMessages#2515f708: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FoundMessages) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode foundMessages#2515f708 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode foundMessages#2515f708: field total_count: %w", err)
		}
		f.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode foundMessages#2515f708: field messages: %w", err)
		}

		if headerLen > 0 {
			f.Messages = make([]Message, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Message
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare foundMessages#2515f708: field messages: %w", err)
			}
			f.Messages = append(f.Messages, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode foundMessages#2515f708: field next_offset: %w", err)
		}
		f.NextOffset = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *FoundMessages) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode foundMessages#2515f708 as nil")
	}
	b.ObjStart()
	b.PutID("foundMessages")
	b.Comma()
	b.FieldStart("total_count")
	b.PutInt32(f.TotalCount)
	b.Comma()
	b.FieldStart("messages")
	b.ArrStart()
	for idx, v := range f.Messages {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode foundMessages#2515f708: field messages element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("next_offset")
	b.PutString(f.NextOffset)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *FoundMessages) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode foundMessages#2515f708 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("foundMessages"); err != nil {
				return fmt.Errorf("unable to decode foundMessages#2515f708: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode foundMessages#2515f708: field total_count: %w", err)
			}
			f.TotalCount = value
		case "messages":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value Message
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode foundMessages#2515f708: field messages: %w", err)
				}
				f.Messages = append(f.Messages, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode foundMessages#2515f708: field messages: %w", err)
			}
		case "next_offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode foundMessages#2515f708: field next_offset: %w", err)
			}
			f.NextOffset = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (f *FoundMessages) GetTotalCount() (value int32) {
	if f == nil {
		return
	}
	return f.TotalCount
}

// GetMessages returns value of Messages field.
func (f *FoundMessages) GetMessages() (value []Message) {
	if f == nil {
		return
	}
	return f.Messages
}

// GetNextOffset returns value of NextOffset field.
func (f *FoundMessages) GetNextOffset() (value string) {
	if f == nil {
		return
	}
	return f.NextOffset
}
