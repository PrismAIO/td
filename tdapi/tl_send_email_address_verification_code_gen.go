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

// SendEmailAddressVerificationCodeRequest represents TL type `sendEmailAddressVerificationCode#f2ca537d`.
type SendEmailAddressVerificationCodeRequest struct {
	// Email address
	EmailAddress string
}

// SendEmailAddressVerificationCodeRequestTypeID is TL type id of SendEmailAddressVerificationCodeRequest.
const SendEmailAddressVerificationCodeRequestTypeID = 0xf2ca537d

// Ensuring interfaces in compile-time for SendEmailAddressVerificationCodeRequest.
var (
	_ bin.Encoder     = &SendEmailAddressVerificationCodeRequest{}
	_ bin.Decoder     = &SendEmailAddressVerificationCodeRequest{}
	_ bin.BareEncoder = &SendEmailAddressVerificationCodeRequest{}
	_ bin.BareDecoder = &SendEmailAddressVerificationCodeRequest{}
)

func (s *SendEmailAddressVerificationCodeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.EmailAddress == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendEmailAddressVerificationCodeRequest) String() string {
	if s == nil {
		return "SendEmailAddressVerificationCodeRequest(nil)"
	}
	type Alias SendEmailAddressVerificationCodeRequest
	return fmt.Sprintf("SendEmailAddressVerificationCodeRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendEmailAddressVerificationCodeRequest) TypeID() uint32 {
	return SendEmailAddressVerificationCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendEmailAddressVerificationCodeRequest) TypeName() string {
	return "sendEmailAddressVerificationCode"
}

// TypeInfo returns info about TL type.
func (s *SendEmailAddressVerificationCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendEmailAddressVerificationCode",
		ID:   SendEmailAddressVerificationCodeRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EmailAddress",
			SchemaName: "email_address",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendEmailAddressVerificationCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendEmailAddressVerificationCode#f2ca537d as nil")
	}
	b.PutID(SendEmailAddressVerificationCodeRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendEmailAddressVerificationCodeRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendEmailAddressVerificationCode#f2ca537d as nil")
	}
	b.PutString(s.EmailAddress)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendEmailAddressVerificationCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendEmailAddressVerificationCode#f2ca537d to nil")
	}
	if err := b.ConsumeID(SendEmailAddressVerificationCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendEmailAddressVerificationCode#f2ca537d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendEmailAddressVerificationCodeRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendEmailAddressVerificationCode#f2ca537d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sendEmailAddressVerificationCode#f2ca537d: field email_address: %w", err)
		}
		s.EmailAddress = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SendEmailAddressVerificationCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendEmailAddressVerificationCode#f2ca537d as nil")
	}
	b.ObjStart()
	b.PutID("sendEmailAddressVerificationCode")
	b.Comma()
	b.FieldStart("email_address")
	b.PutString(s.EmailAddress)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SendEmailAddressVerificationCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendEmailAddressVerificationCode#f2ca537d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sendEmailAddressVerificationCode"); err != nil {
				return fmt.Errorf("unable to decode sendEmailAddressVerificationCode#f2ca537d: %w", err)
			}
		case "email_address":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sendEmailAddressVerificationCode#f2ca537d: field email_address: %w", err)
			}
			s.EmailAddress = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEmailAddress returns value of EmailAddress field.
func (s *SendEmailAddressVerificationCodeRequest) GetEmailAddress() (value string) {
	if s == nil {
		return
	}
	return s.EmailAddress
}

// SendEmailAddressVerificationCode invokes method sendEmailAddressVerificationCode#f2ca537d returning error if any.
func (c *Client) SendEmailAddressVerificationCode(ctx context.Context, emailaddress string) (*EmailAddressAuthenticationCodeInfo, error) {
	var result EmailAddressAuthenticationCodeInfo

	request := &SendEmailAddressVerificationCodeRequest{
		EmailAddress: emailaddress,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
