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

// PageListItemText represents TL type `pageListItemText#b92fb6cd`.
// List item
//
// See https://core.telegram.org/constructor/pageListItemText for reference.
type PageListItemText struct {
	// Text
	Text RichTextClass
}

// PageListItemTextTypeID is TL type id of PageListItemText.
const PageListItemTextTypeID = 0xb92fb6cd

// construct implements constructor of PageListItemClass.
func (p PageListItemText) construct() PageListItemClass { return &p }

// Ensuring interfaces in compile-time for PageListItemText.
var (
	_ bin.Encoder     = &PageListItemText{}
	_ bin.Decoder     = &PageListItemText{}
	_ bin.BareEncoder = &PageListItemText{}
	_ bin.BareDecoder = &PageListItemText{}

	_ PageListItemClass = &PageListItemText{}
)

func (p *PageListItemText) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Text == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageListItemText) String() string {
	if p == nil {
		return "PageListItemText(nil)"
	}
	type Alias PageListItemText
	return fmt.Sprintf("PageListItemText%+v", Alias(*p))
}

// FillFrom fills PageListItemText from given interface.
func (p *PageListItemText) FillFrom(from interface {
	GetText() (value RichTextClass)
}) {
	p.Text = from.GetText()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PageListItemText) TypeID() uint32 {
	return PageListItemTextTypeID
}

// TypeName returns name of type in TL schema.
func (*PageListItemText) TypeName() string {
	return "pageListItemText"
}

// TypeInfo returns info about TL type.
func (p *PageListItemText) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pageListItemText",
		ID:   PageListItemTextTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PageListItemText) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageListItemText#b92fb6cd as nil")
	}
	b.PutID(PageListItemTextTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PageListItemText) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageListItemText#b92fb6cd as nil")
	}
	if p.Text == nil {
		return fmt.Errorf("unable to encode pageListItemText#b92fb6cd: field text is nil")
	}
	if err := p.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageListItemText#b92fb6cd: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageListItemText) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageListItemText#b92fb6cd to nil")
	}
	if err := b.ConsumeID(PageListItemTextTypeID); err != nil {
		return fmt.Errorf("unable to decode pageListItemText#b92fb6cd: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PageListItemText) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageListItemText#b92fb6cd to nil")
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode pageListItemText#b92fb6cd: field text: %w", err)
		}
		p.Text = value
	}
	return nil
}

// GetText returns value of Text field.
func (p *PageListItemText) GetText() (value RichTextClass) {
	if p == nil {
		return
	}
	return p.Text
}

// PageListItemBlocks represents TL type `pageListItemBlocks#25e073fc`.
// List item
//
// See https://core.telegram.org/constructor/pageListItemBlocks for reference.
type PageListItemBlocks struct {
	// Blocks
	Blocks []PageBlockClass
}

// PageListItemBlocksTypeID is TL type id of PageListItemBlocks.
const PageListItemBlocksTypeID = 0x25e073fc

// construct implements constructor of PageListItemClass.
func (p PageListItemBlocks) construct() PageListItemClass { return &p }

// Ensuring interfaces in compile-time for PageListItemBlocks.
var (
	_ bin.Encoder     = &PageListItemBlocks{}
	_ bin.Decoder     = &PageListItemBlocks{}
	_ bin.BareEncoder = &PageListItemBlocks{}
	_ bin.BareDecoder = &PageListItemBlocks{}

	_ PageListItemClass = &PageListItemBlocks{}
)

func (p *PageListItemBlocks) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Blocks == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageListItemBlocks) String() string {
	if p == nil {
		return "PageListItemBlocks(nil)"
	}
	type Alias PageListItemBlocks
	return fmt.Sprintf("PageListItemBlocks%+v", Alias(*p))
}

// FillFrom fills PageListItemBlocks from given interface.
func (p *PageListItemBlocks) FillFrom(from interface {
	GetBlocks() (value []PageBlockClass)
}) {
	p.Blocks = from.GetBlocks()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PageListItemBlocks) TypeID() uint32 {
	return PageListItemBlocksTypeID
}

// TypeName returns name of type in TL schema.
func (*PageListItemBlocks) TypeName() string {
	return "pageListItemBlocks"
}

// TypeInfo returns info about TL type.
func (p *PageListItemBlocks) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pageListItemBlocks",
		ID:   PageListItemBlocksTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Blocks",
			SchemaName: "blocks",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PageListItemBlocks) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageListItemBlocks#25e073fc as nil")
	}
	b.PutID(PageListItemBlocksTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PageListItemBlocks) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageListItemBlocks#25e073fc as nil")
	}
	b.PutVectorHeader(len(p.Blocks))
	for idx, v := range p.Blocks {
		if v == nil {
			return fmt.Errorf("unable to encode pageListItemBlocks#25e073fc: field blocks element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode pageListItemBlocks#25e073fc: field blocks element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageListItemBlocks) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageListItemBlocks#25e073fc to nil")
	}
	if err := b.ConsumeID(PageListItemBlocksTypeID); err != nil {
		return fmt.Errorf("unable to decode pageListItemBlocks#25e073fc: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PageListItemBlocks) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageListItemBlocks#25e073fc to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode pageListItemBlocks#25e073fc: field blocks: %w", err)
		}

		if headerLen > 0 {
			p.Blocks = make([]PageBlockClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePageBlock(b)
			if err != nil {
				return fmt.Errorf("unable to decode pageListItemBlocks#25e073fc: field blocks: %w", err)
			}
			p.Blocks = append(p.Blocks, value)
		}
	}
	return nil
}

// GetBlocks returns value of Blocks field.
func (p *PageListItemBlocks) GetBlocks() (value []PageBlockClass) {
	if p == nil {
		return
	}
	return p.Blocks
}

// MapBlocks returns field Blocks wrapped in PageBlockClassArray helper.
func (p *PageListItemBlocks) MapBlocks() (value PageBlockClassArray) {
	return PageBlockClassArray(p.Blocks)
}

// PageListItemClassName is schema name of PageListItemClass.
const PageListItemClassName = "PageListItem"

// PageListItemClass represents PageListItem generic type.
//
// See https://core.telegram.org/type/PageListItem for reference.
//
// Example:
//
//	g, err := tg.DecodePageListItem(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.PageListItemText: // pageListItemText#b92fb6cd
//	case *tg.PageListItemBlocks: // pageListItemBlocks#25e073fc
//	default: panic(v)
//	}
type PageListItemClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PageListItemClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodePageListItem implements binary de-serialization for PageListItemClass.
func DecodePageListItem(buf *bin.Buffer) (PageListItemClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PageListItemTextTypeID:
		// Decoding pageListItemText#b92fb6cd.
		v := PageListItemText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageListItemClass: %w", err)
		}
		return &v, nil
	case PageListItemBlocksTypeID:
		// Decoding pageListItemBlocks#25e073fc.
		v := PageListItemBlocks{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageListItemClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PageListItemClass: %w", bin.NewUnexpectedID(id))
	}
}

// PageListItem boxes the PageListItemClass providing a helper.
type PageListItemBox struct {
	PageListItem PageListItemClass
}

// Decode implements bin.Decoder for PageListItemBox.
func (b *PageListItemBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PageListItemBox to nil")
	}
	v, err := DecodePageListItem(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PageListItem = v
	return nil
}

// Encode implements bin.Encode for PageListItemBox.
func (b *PageListItemBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PageListItem == nil {
		return fmt.Errorf("unable to encode PageListItemClass as nil")
	}
	return b.PageListItem.Encode(buf)
}
