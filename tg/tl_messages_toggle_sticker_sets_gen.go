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

// MessagesToggleStickerSetsRequest represents TL type `messages.toggleStickerSets#b5052fea`.
// Apply changes to multiple stickersets
//
// See https://core.telegram.org/method/messages.toggleStickerSets for reference.
type MessagesToggleStickerSetsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Uninstall the specified stickersets
	Uninstall bool
	// Archive the specified stickersets
	Archive bool
	// Unarchive the specified stickersets
	Unarchive bool
	// Stickersets to act upon
	Stickersets []InputStickerSetClass
}

// MessagesToggleStickerSetsRequestTypeID is TL type id of MessagesToggleStickerSetsRequest.
const MessagesToggleStickerSetsRequestTypeID = 0xb5052fea

// Ensuring interfaces in compile-time for MessagesToggleStickerSetsRequest.
var (
	_ bin.Encoder     = &MessagesToggleStickerSetsRequest{}
	_ bin.Decoder     = &MessagesToggleStickerSetsRequest{}
	_ bin.BareEncoder = &MessagesToggleStickerSetsRequest{}
	_ bin.BareDecoder = &MessagesToggleStickerSetsRequest{}
)

func (t *MessagesToggleStickerSetsRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Uninstall == false) {
		return false
	}
	if !(t.Archive == false) {
		return false
	}
	if !(t.Unarchive == false) {
		return false
	}
	if !(t.Stickersets == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesToggleStickerSetsRequest) String() string {
	if t == nil {
		return "MessagesToggleStickerSetsRequest(nil)"
	}
	type Alias MessagesToggleStickerSetsRequest
	return fmt.Sprintf("MessagesToggleStickerSetsRequest%+v", Alias(*t))
}

// FillFrom fills MessagesToggleStickerSetsRequest from given interface.
func (t *MessagesToggleStickerSetsRequest) FillFrom(from interface {
	GetUninstall() (value bool)
	GetArchive() (value bool)
	GetUnarchive() (value bool)
	GetStickersets() (value []InputStickerSetClass)
}) {
	t.Uninstall = from.GetUninstall()
	t.Archive = from.GetArchive()
	t.Unarchive = from.GetUnarchive()
	t.Stickersets = from.GetStickersets()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesToggleStickerSetsRequest) TypeID() uint32 {
	return MessagesToggleStickerSetsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesToggleStickerSetsRequest) TypeName() string {
	return "messages.toggleStickerSets"
}

// TypeInfo returns info about TL type.
func (t *MessagesToggleStickerSetsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.toggleStickerSets",
		ID:   MessagesToggleStickerSetsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Uninstall",
			SchemaName: "uninstall",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "Archive",
			SchemaName: "archive",
			Null:       !t.Flags.Has(1),
		},
		{
			Name:       "Unarchive",
			SchemaName: "unarchive",
			Null:       !t.Flags.Has(2),
		},
		{
			Name:       "Stickersets",
			SchemaName: "stickersets",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (t *MessagesToggleStickerSetsRequest) SetFlags() {
	if !(t.Uninstall == false) {
		t.Flags.Set(0)
	}
	if !(t.Archive == false) {
		t.Flags.Set(1)
	}
	if !(t.Unarchive == false) {
		t.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (t *MessagesToggleStickerSetsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.toggleStickerSets#b5052fea as nil")
	}
	b.PutID(MessagesToggleStickerSetsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesToggleStickerSetsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.toggleStickerSets#b5052fea as nil")
	}
	t.SetFlags()
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.toggleStickerSets#b5052fea: field flags: %w", err)
	}
	b.PutVectorHeader(len(t.Stickersets))
	for idx, v := range t.Stickersets {
		if v == nil {
			return fmt.Errorf("unable to encode messages.toggleStickerSets#b5052fea: field stickersets element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.toggleStickerSets#b5052fea: field stickersets element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesToggleStickerSetsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.toggleStickerSets#b5052fea to nil")
	}
	if err := b.ConsumeID(MessagesToggleStickerSetsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.toggleStickerSets#b5052fea: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesToggleStickerSetsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.toggleStickerSets#b5052fea to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.toggleStickerSets#b5052fea: field flags: %w", err)
		}
	}
	t.Uninstall = t.Flags.Has(0)
	t.Archive = t.Flags.Has(1)
	t.Unarchive = t.Flags.Has(2)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.toggleStickerSets#b5052fea: field stickersets: %w", err)
		}

		if headerLen > 0 {
			t.Stickersets = make([]InputStickerSetClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputStickerSet(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.toggleStickerSets#b5052fea: field stickersets: %w", err)
			}
			t.Stickersets = append(t.Stickersets, value)
		}
	}
	return nil
}

// SetUninstall sets value of Uninstall conditional field.
func (t *MessagesToggleStickerSetsRequest) SetUninstall(value bool) {
	if value {
		t.Flags.Set(0)
		t.Uninstall = true
	} else {
		t.Flags.Unset(0)
		t.Uninstall = false
	}
}

// GetUninstall returns value of Uninstall conditional field.
func (t *MessagesToggleStickerSetsRequest) GetUninstall() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(0)
}

// SetArchive sets value of Archive conditional field.
func (t *MessagesToggleStickerSetsRequest) SetArchive(value bool) {
	if value {
		t.Flags.Set(1)
		t.Archive = true
	} else {
		t.Flags.Unset(1)
		t.Archive = false
	}
}

// GetArchive returns value of Archive conditional field.
func (t *MessagesToggleStickerSetsRequest) GetArchive() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(1)
}

// SetUnarchive sets value of Unarchive conditional field.
func (t *MessagesToggleStickerSetsRequest) SetUnarchive(value bool) {
	if value {
		t.Flags.Set(2)
		t.Unarchive = true
	} else {
		t.Flags.Unset(2)
		t.Unarchive = false
	}
}

// GetUnarchive returns value of Unarchive conditional field.
func (t *MessagesToggleStickerSetsRequest) GetUnarchive() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(2)
}

// GetStickersets returns value of Stickersets field.
func (t *MessagesToggleStickerSetsRequest) GetStickersets() (value []InputStickerSetClass) {
	if t == nil {
		return
	}
	return t.Stickersets
}

// MapStickersets returns field Stickersets wrapped in InputStickerSetClassArray helper.
func (t *MessagesToggleStickerSetsRequest) MapStickersets() (value InputStickerSetClassArray) {
	return InputStickerSetClassArray(t.Stickersets)
}

// MessagesToggleStickerSets invokes method messages.toggleStickerSets#b5052fea returning error if any.
// Apply changes to multiple stickersets
//
// See https://core.telegram.org/method/messages.toggleStickerSets for reference.
func (c *Client) MessagesToggleStickerSets(ctx context.Context, request *MessagesToggleStickerSetsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
