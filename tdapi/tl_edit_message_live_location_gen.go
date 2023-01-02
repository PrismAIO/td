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

// EditMessageLiveLocationRequest represents TL type `editMessageLiveLocation#ff29a512`.
type EditMessageLiveLocationRequest struct {
	// The chat the message belongs to
	ChatID int64
	// Identifier of the message
	MessageID int64
	// The new message reply markup; pass null if none; for bots only
	ReplyMarkup ReplyMarkupClass
	// New location content of the message; pass null to stop sharing the live location
	Location Location
	// The new direction in which the location moves, in degrees; 1-360. Pass 0 if unknown
	Heading int32
	// The new maximum distance for proximity alerts, in meters (0-100000). Pass 0 if the
	// notification is disabled
	ProximityAlertRadius int32
}

// EditMessageLiveLocationRequestTypeID is TL type id of EditMessageLiveLocationRequest.
const EditMessageLiveLocationRequestTypeID = 0xff29a512

// Ensuring interfaces in compile-time for EditMessageLiveLocationRequest.
var (
	_ bin.Encoder     = &EditMessageLiveLocationRequest{}
	_ bin.Decoder     = &EditMessageLiveLocationRequest{}
	_ bin.BareEncoder = &EditMessageLiveLocationRequest{}
	_ bin.BareDecoder = &EditMessageLiveLocationRequest{}
)

func (e *EditMessageLiveLocationRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.MessageID == 0) {
		return false
	}
	if !(e.ReplyMarkup == nil) {
		return false
	}
	if !(e.Location.Zero()) {
		return false
	}
	if !(e.Heading == 0) {
		return false
	}
	if !(e.ProximityAlertRadius == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditMessageLiveLocationRequest) String() string {
	if e == nil {
		return "EditMessageLiveLocationRequest(nil)"
	}
	type Alias EditMessageLiveLocationRequest
	return fmt.Sprintf("EditMessageLiveLocationRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditMessageLiveLocationRequest) TypeID() uint32 {
	return EditMessageLiveLocationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditMessageLiveLocationRequest) TypeName() string {
	return "editMessageLiveLocation"
}

// TypeInfo returns info about TL type.
func (e *EditMessageLiveLocationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editMessageLiveLocation",
		ID:   EditMessageLiveLocationRequestTypeID,
	}
	if e == nil {
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
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
		{
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "Heading",
			SchemaName: "heading",
		},
		{
			Name:       "ProximityAlertRadius",
			SchemaName: "proximity_alert_radius",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditMessageLiveLocationRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageLiveLocation#ff29a512 as nil")
	}
	b.PutID(EditMessageLiveLocationRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditMessageLiveLocationRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageLiveLocation#ff29a512 as nil")
	}
	b.PutInt53(e.ChatID)
	b.PutInt53(e.MessageID)
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editMessageLiveLocation#ff29a512: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editMessageLiveLocation#ff29a512: field reply_markup: %w", err)
	}
	if err := e.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editMessageLiveLocation#ff29a512: field location: %w", err)
	}
	b.PutInt32(e.Heading)
	b.PutInt32(e.ProximityAlertRadius)
	return nil
}

// Decode implements bin.Decoder.
func (e *EditMessageLiveLocationRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageLiveLocation#ff29a512 to nil")
	}
	if err := b.ConsumeID(EditMessageLiveLocationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditMessageLiveLocationRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageLiveLocation#ff29a512 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field message_id: %w", err)
		}
		e.MessageID = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	{
		if err := e.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field location: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field heading: %w", err)
		}
		e.Heading = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field proximity_alert_radius: %w", err)
		}
		e.ProximityAlertRadius = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditMessageLiveLocationRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageLiveLocation#ff29a512 as nil")
	}
	b.ObjStart()
	b.PutID("editMessageLiveLocation")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(e.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(e.MessageID)
	b.Comma()
	b.FieldStart("reply_markup")
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editMessageLiveLocation#ff29a512: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editMessageLiveLocation#ff29a512: field reply_markup: %w", err)
	}
	b.Comma()
	b.FieldStart("location")
	if err := e.Location.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editMessageLiveLocation#ff29a512: field location: %w", err)
	}
	b.Comma()
	b.FieldStart("heading")
	b.PutInt32(e.Heading)
	b.Comma()
	b.FieldStart("proximity_alert_radius")
	b.PutInt32(e.ProximityAlertRadius)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditMessageLiveLocationRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageLiveLocation#ff29a512 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editMessageLiveLocation"); err != nil {
				return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field chat_id: %w", err)
			}
			e.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field message_id: %w", err)
			}
			e.MessageID = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field reply_markup: %w", err)
			}
			e.ReplyMarkup = value
		case "location":
			if err := e.Location.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field location: %w", err)
			}
		case "heading":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field heading: %w", err)
			}
			e.Heading = value
		case "proximity_alert_radius":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageLiveLocation#ff29a512: field proximity_alert_radius: %w", err)
			}
			e.ProximityAlertRadius = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (e *EditMessageLiveLocationRequest) GetChatID() (value int64) {
	if e == nil {
		return
	}
	return e.ChatID
}

// GetMessageID returns value of MessageID field.
func (e *EditMessageLiveLocationRequest) GetMessageID() (value int64) {
	if e == nil {
		return
	}
	return e.MessageID
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (e *EditMessageLiveLocationRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	if e == nil {
		return
	}
	return e.ReplyMarkup
}

// GetLocation returns value of Location field.
func (e *EditMessageLiveLocationRequest) GetLocation() (value Location) {
	if e == nil {
		return
	}
	return e.Location
}

// GetHeading returns value of Heading field.
func (e *EditMessageLiveLocationRequest) GetHeading() (value int32) {
	if e == nil {
		return
	}
	return e.Heading
}

// GetProximityAlertRadius returns value of ProximityAlertRadius field.
func (e *EditMessageLiveLocationRequest) GetProximityAlertRadius() (value int32) {
	if e == nil {
		return
	}
	return e.ProximityAlertRadius
}

// EditMessageLiveLocation invokes method editMessageLiveLocation#ff29a512 returning error if any.
func (c *Client) EditMessageLiveLocation(ctx context.Context, request *EditMessageLiveLocationRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
