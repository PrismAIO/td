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

// ReportMessageReactionsRequest represents TL type `reportMessageReactions#36c88827`.
type ReportMessageReactionsRequest struct {
	// Chat identifier
	ChatID int64
	// Message identifier
	MessageID int64
	// Identifier of the sender, which added the reaction
	SenderID MessageSenderClass
}

// ReportMessageReactionsRequestTypeID is TL type id of ReportMessageReactionsRequest.
const ReportMessageReactionsRequestTypeID = 0x36c88827

// Ensuring interfaces in compile-time for ReportMessageReactionsRequest.
var (
	_ bin.Encoder     = &ReportMessageReactionsRequest{}
	_ bin.Decoder     = &ReportMessageReactionsRequest{}
	_ bin.BareEncoder = &ReportMessageReactionsRequest{}
	_ bin.BareDecoder = &ReportMessageReactionsRequest{}
)

func (r *ReportMessageReactionsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}
	if !(r.MessageID == 0) {
		return false
	}
	if !(r.SenderID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportMessageReactionsRequest) String() string {
	if r == nil {
		return "ReportMessageReactionsRequest(nil)"
	}
	type Alias ReportMessageReactionsRequest
	return fmt.Sprintf("ReportMessageReactionsRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportMessageReactionsRequest) TypeID() uint32 {
	return ReportMessageReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportMessageReactionsRequest) TypeName() string {
	return "reportMessageReactions"
}

// TypeInfo returns info about TL type.
func (r *ReportMessageReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportMessageReactions",
		ID:   ReportMessageReactionsRequestTypeID,
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
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "SenderID",
			SchemaName: "sender_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportMessageReactionsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportMessageReactions#36c88827 as nil")
	}
	b.PutID(ReportMessageReactionsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportMessageReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportMessageReactions#36c88827 as nil")
	}
	b.PutInt53(r.ChatID)
	b.PutInt53(r.MessageID)
	if r.SenderID == nil {
		return fmt.Errorf("unable to encode reportMessageReactions#36c88827: field sender_id is nil")
	}
	if err := r.SenderID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reportMessageReactions#36c88827: field sender_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportMessageReactionsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportMessageReactions#36c88827 to nil")
	}
	if err := b.ConsumeID(ReportMessageReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reportMessageReactions#36c88827: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportMessageReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportMessageReactions#36c88827 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode reportMessageReactions#36c88827: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode reportMessageReactions#36c88827: field message_id: %w", err)
		}
		r.MessageID = value
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode reportMessageReactions#36c88827: field sender_id: %w", err)
		}
		r.SenderID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportMessageReactionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportMessageReactions#36c88827 as nil")
	}
	b.ObjStart()
	b.PutID("reportMessageReactions")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(r.MessageID)
	b.Comma()
	b.FieldStart("sender_id")
	if r.SenderID == nil {
		return fmt.Errorf("unable to encode reportMessageReactions#36c88827: field sender_id is nil")
	}
	if err := r.SenderID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reportMessageReactions#36c88827: field sender_id: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportMessageReactionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportMessageReactions#36c88827 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportMessageReactions"); err != nil {
				return fmt.Errorf("unable to decode reportMessageReactions#36c88827: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode reportMessageReactions#36c88827: field chat_id: %w", err)
			}
			r.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode reportMessageReactions#36c88827: field message_id: %w", err)
			}
			r.MessageID = value
		case "sender_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode reportMessageReactions#36c88827: field sender_id: %w", err)
			}
			r.SenderID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *ReportMessageReactionsRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// GetMessageID returns value of MessageID field.
func (r *ReportMessageReactionsRequest) GetMessageID() (value int64) {
	if r == nil {
		return
	}
	return r.MessageID
}

// GetSenderID returns value of SenderID field.
func (r *ReportMessageReactionsRequest) GetSenderID() (value MessageSenderClass) {
	if r == nil {
		return
	}
	return r.SenderID
}

// ReportMessageReactions invokes method reportMessageReactions#36c88827 returning error if any.
func (c *Client) ReportMessageReactions(ctx context.Context, request *ReportMessageReactionsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}