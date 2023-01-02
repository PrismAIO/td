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

// SuggestUserProfilePhotoRequest represents TL type `suggestUserProfilePhoto#9561f463`.
type SuggestUserProfilePhotoRequest struct {
	// User identifier
	UserID int64
	// Profile photo to suggest; inputChatPhotoPrevious isn't supported in this function
	Photo InputChatPhotoClass
}

// SuggestUserProfilePhotoRequestTypeID is TL type id of SuggestUserProfilePhotoRequest.
const SuggestUserProfilePhotoRequestTypeID = 0x9561f463

// Ensuring interfaces in compile-time for SuggestUserProfilePhotoRequest.
var (
	_ bin.Encoder     = &SuggestUserProfilePhotoRequest{}
	_ bin.Decoder     = &SuggestUserProfilePhotoRequest{}
	_ bin.BareEncoder = &SuggestUserProfilePhotoRequest{}
	_ bin.BareDecoder = &SuggestUserProfilePhotoRequest{}
)

func (s *SuggestUserProfilePhotoRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.Photo == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestUserProfilePhotoRequest) String() string {
	if s == nil {
		return "SuggestUserProfilePhotoRequest(nil)"
	}
	type Alias SuggestUserProfilePhotoRequest
	return fmt.Sprintf("SuggestUserProfilePhotoRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestUserProfilePhotoRequest) TypeID() uint32 {
	return SuggestUserProfilePhotoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestUserProfilePhotoRequest) TypeName() string {
	return "suggestUserProfilePhoto"
}

// TypeInfo returns info about TL type.
func (s *SuggestUserProfilePhotoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestUserProfilePhoto",
		ID:   SuggestUserProfilePhotoRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestUserProfilePhotoRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestUserProfilePhoto#9561f463 as nil")
	}
	b.PutID(SuggestUserProfilePhotoRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestUserProfilePhotoRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestUserProfilePhoto#9561f463 as nil")
	}
	b.PutInt53(s.UserID)
	if s.Photo == nil {
		return fmt.Errorf("unable to encode suggestUserProfilePhoto#9561f463: field photo is nil")
	}
	if err := s.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode suggestUserProfilePhoto#9561f463: field photo: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestUserProfilePhotoRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestUserProfilePhoto#9561f463 to nil")
	}
	if err := b.ConsumeID(SuggestUserProfilePhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestUserProfilePhoto#9561f463: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestUserProfilePhotoRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestUserProfilePhoto#9561f463 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode suggestUserProfilePhoto#9561f463: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := DecodeInputChatPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode suggestUserProfilePhoto#9561f463: field photo: %w", err)
		}
		s.Photo = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestUserProfilePhotoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestUserProfilePhoto#9561f463 as nil")
	}
	b.ObjStart()
	b.PutID("suggestUserProfilePhoto")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(s.UserID)
	b.Comma()
	b.FieldStart("photo")
	if s.Photo == nil {
		return fmt.Errorf("unable to encode suggestUserProfilePhoto#9561f463: field photo is nil")
	}
	if err := s.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode suggestUserProfilePhoto#9561f463: field photo: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestUserProfilePhotoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestUserProfilePhoto#9561f463 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestUserProfilePhoto"); err != nil {
				return fmt.Errorf("unable to decode suggestUserProfilePhoto#9561f463: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode suggestUserProfilePhoto#9561f463: field user_id: %w", err)
			}
			s.UserID = value
		case "photo":
			value, err := DecodeTDLibJSONInputChatPhoto(b)
			if err != nil {
				return fmt.Errorf("unable to decode suggestUserProfilePhoto#9561f463: field photo: %w", err)
			}
			s.Photo = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (s *SuggestUserProfilePhotoRequest) GetUserID() (value int64) {
	if s == nil {
		return
	}
	return s.UserID
}

// GetPhoto returns value of Photo field.
func (s *SuggestUserProfilePhotoRequest) GetPhoto() (value InputChatPhotoClass) {
	if s == nil {
		return
	}
	return s.Photo
}

// SuggestUserProfilePhoto invokes method suggestUserProfilePhoto#9561f463 returning error if any.
func (c *Client) SuggestUserProfilePhoto(ctx context.Context, request *SuggestUserProfilePhotoRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
