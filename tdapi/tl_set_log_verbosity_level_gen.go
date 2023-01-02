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

// SetLogVerbosityLevelRequest represents TL type `setLogVerbosityLevel#edea07d2`.
type SetLogVerbosityLevelRequest struct {
	// New value of the verbosity level for logging. Value 0 corresponds to fatal errors,
	// value 1 corresponds to errors, value 2 corresponds to warnings and debug warnings,
	NewVerbosityLevel int32
}

// SetLogVerbosityLevelRequestTypeID is TL type id of SetLogVerbosityLevelRequest.
const SetLogVerbosityLevelRequestTypeID = 0xedea07d2

// Ensuring interfaces in compile-time for SetLogVerbosityLevelRequest.
var (
	_ bin.Encoder     = &SetLogVerbosityLevelRequest{}
	_ bin.Decoder     = &SetLogVerbosityLevelRequest{}
	_ bin.BareEncoder = &SetLogVerbosityLevelRequest{}
	_ bin.BareDecoder = &SetLogVerbosityLevelRequest{}
)

func (s *SetLogVerbosityLevelRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.NewVerbosityLevel == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetLogVerbosityLevelRequest) String() string {
	if s == nil {
		return "SetLogVerbosityLevelRequest(nil)"
	}
	type Alias SetLogVerbosityLevelRequest
	return fmt.Sprintf("SetLogVerbosityLevelRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetLogVerbosityLevelRequest) TypeID() uint32 {
	return SetLogVerbosityLevelRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetLogVerbosityLevelRequest) TypeName() string {
	return "setLogVerbosityLevel"
}

// TypeInfo returns info about TL type.
func (s *SetLogVerbosityLevelRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setLogVerbosityLevel",
		ID:   SetLogVerbosityLevelRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NewVerbosityLevel",
			SchemaName: "new_verbosity_level",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetLogVerbosityLevelRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setLogVerbosityLevel#edea07d2 as nil")
	}
	b.PutID(SetLogVerbosityLevelRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetLogVerbosityLevelRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setLogVerbosityLevel#edea07d2 as nil")
	}
	b.PutInt32(s.NewVerbosityLevel)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetLogVerbosityLevelRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setLogVerbosityLevel#edea07d2 to nil")
	}
	if err := b.ConsumeID(SetLogVerbosityLevelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setLogVerbosityLevel#edea07d2: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetLogVerbosityLevelRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setLogVerbosityLevel#edea07d2 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setLogVerbosityLevel#edea07d2: field new_verbosity_level: %w", err)
		}
		s.NewVerbosityLevel = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetLogVerbosityLevelRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setLogVerbosityLevel#edea07d2 as nil")
	}
	b.ObjStart()
	b.PutID("setLogVerbosityLevel")
	b.Comma()
	b.FieldStart("new_verbosity_level")
	b.PutInt32(s.NewVerbosityLevel)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetLogVerbosityLevelRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setLogVerbosityLevel#edea07d2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setLogVerbosityLevel"); err != nil {
				return fmt.Errorf("unable to decode setLogVerbosityLevel#edea07d2: %w", err)
			}
		case "new_verbosity_level":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode setLogVerbosityLevel#edea07d2: field new_verbosity_level: %w", err)
			}
			s.NewVerbosityLevel = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetNewVerbosityLevel returns value of NewVerbosityLevel field.
func (s *SetLogVerbosityLevelRequest) GetNewVerbosityLevel() (value int32) {
	if s == nil {
		return
	}
	return s.NewVerbosityLevel
}

// SetLogVerbosityLevel invokes method setLogVerbosityLevel#edea07d2 returning error if any.
func (c *Client) SetLogVerbosityLevel(ctx context.Context, newverbositylevel int32) error {
	var ok Ok

	request := &SetLogVerbosityLevelRequest{
		NewVerbosityLevel: newverbositylevel,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
