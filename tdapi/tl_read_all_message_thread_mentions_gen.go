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

// ReadAllMessageThreadMentionsRequest represents TL type `readAllMessageThreadMentions#4edd7555`.
type ReadAllMessageThreadMentionsRequest struct {
	// Chat identifier
	ChatID int64
	// Message thread identifier in which mentions are marked as read
	MessageThreadID int64
}

// ReadAllMessageThreadMentionsRequestTypeID is TL type id of ReadAllMessageThreadMentionsRequest.
const ReadAllMessageThreadMentionsRequestTypeID = 0x4edd7555

// Ensuring interfaces in compile-time for ReadAllMessageThreadMentionsRequest.
var (
	_ bin.Encoder     = &ReadAllMessageThreadMentionsRequest{}
	_ bin.Decoder     = &ReadAllMessageThreadMentionsRequest{}
	_ bin.BareEncoder = &ReadAllMessageThreadMentionsRequest{}
	_ bin.BareDecoder = &ReadAllMessageThreadMentionsRequest{}
)

func (r *ReadAllMessageThreadMentionsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}
	if !(r.MessageThreadID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReadAllMessageThreadMentionsRequest) String() string {
	if r == nil {
		return "ReadAllMessageThreadMentionsRequest(nil)"
	}
	type Alias ReadAllMessageThreadMentionsRequest
	return fmt.Sprintf("ReadAllMessageThreadMentionsRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReadAllMessageThreadMentionsRequest) TypeID() uint32 {
	return ReadAllMessageThreadMentionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReadAllMessageThreadMentionsRequest) TypeName() string {
	return "readAllMessageThreadMentions"
}

// TypeInfo returns info about TL type.
func (r *ReadAllMessageThreadMentionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "readAllMessageThreadMentions",
		ID:   ReadAllMessageThreadMentionsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReadAllMessageThreadMentionsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode readAllMessageThreadMentions#4edd7555 as nil")
	}
	b.PutID(ReadAllMessageThreadMentionsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReadAllMessageThreadMentionsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode readAllMessageThreadMentions#4edd7555 as nil")
	}
	b.PutInt53(r.ChatID)
	b.PutInt53(r.MessageThreadID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReadAllMessageThreadMentionsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode readAllMessageThreadMentions#4edd7555 to nil")
	}
	if err := b.ConsumeID(ReadAllMessageThreadMentionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode readAllMessageThreadMentions#4edd7555: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReadAllMessageThreadMentionsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode readAllMessageThreadMentions#4edd7555 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode readAllMessageThreadMentions#4edd7555: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode readAllMessageThreadMentions#4edd7555: field message_thread_id: %w", err)
		}
		r.MessageThreadID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReadAllMessageThreadMentionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode readAllMessageThreadMentions#4edd7555 as nil")
	}
	b.ObjStart()
	b.PutID("readAllMessageThreadMentions")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(r.MessageThreadID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReadAllMessageThreadMentionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode readAllMessageThreadMentions#4edd7555 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("readAllMessageThreadMentions"); err != nil {
				return fmt.Errorf("unable to decode readAllMessageThreadMentions#4edd7555: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode readAllMessageThreadMentions#4edd7555: field chat_id: %w", err)
			}
			r.ChatID = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode readAllMessageThreadMentions#4edd7555: field message_thread_id: %w", err)
			}
			r.MessageThreadID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *ReadAllMessageThreadMentionsRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (r *ReadAllMessageThreadMentionsRequest) GetMessageThreadID() (value int64) {
	if r == nil {
		return
	}
	return r.MessageThreadID
}

// ReadAllMessageThreadMentions invokes method readAllMessageThreadMentions#4edd7555 returning error if any.
func (c *Client) ReadAllMessageThreadMentions(ctx context.Context, request *ReadAllMessageThreadMentionsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
