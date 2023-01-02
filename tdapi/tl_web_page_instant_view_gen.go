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

// WebPageInstantView represents TL type `webPageInstantView#2c0ec99c`.
type WebPageInstantView struct {
	// Content of the web page
	PageBlocks []PageBlockClass
	// Number of the instant view views; 0 if unknown
	ViewCount int32
	// Version of the instant view; currently, can be 1 or 2
	Version int32
	// True, if the instant view must be shown from right to left
	IsRtl bool
	// True, if the instant view contains the full page. A network request might be needed to
	// get the full web page instant view
	IsFull bool
	// An internal link to be opened to leave feedback about the instant view
	FeedbackLink InternalLinkTypeClass
}

// WebPageInstantViewTypeID is TL type id of WebPageInstantView.
const WebPageInstantViewTypeID = 0x2c0ec99c

// Ensuring interfaces in compile-time for WebPageInstantView.
var (
	_ bin.Encoder     = &WebPageInstantView{}
	_ bin.Decoder     = &WebPageInstantView{}
	_ bin.BareEncoder = &WebPageInstantView{}
	_ bin.BareDecoder = &WebPageInstantView{}
)

func (w *WebPageInstantView) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.PageBlocks == nil) {
		return false
	}
	if !(w.ViewCount == 0) {
		return false
	}
	if !(w.Version == 0) {
		return false
	}
	if !(w.IsRtl == false) {
		return false
	}
	if !(w.IsFull == false) {
		return false
	}
	if !(w.FeedbackLink == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebPageInstantView) String() string {
	if w == nil {
		return "WebPageInstantView(nil)"
	}
	type Alias WebPageInstantView
	return fmt.Sprintf("WebPageInstantView%+v", Alias(*w))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebPageInstantView) TypeID() uint32 {
	return WebPageInstantViewTypeID
}

// TypeName returns name of type in TL schema.
func (*WebPageInstantView) TypeName() string {
	return "webPageInstantView"
}

// TypeInfo returns info about TL type.
func (w *WebPageInstantView) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webPageInstantView",
		ID:   WebPageInstantViewTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PageBlocks",
			SchemaName: "page_blocks",
		},
		{
			Name:       "ViewCount",
			SchemaName: "view_count",
		},
		{
			Name:       "Version",
			SchemaName: "version",
		},
		{
			Name:       "IsRtl",
			SchemaName: "is_rtl",
		},
		{
			Name:       "IsFull",
			SchemaName: "is_full",
		},
		{
			Name:       "FeedbackLink",
			SchemaName: "feedback_link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *WebPageInstantView) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageInstantView#2c0ec99c as nil")
	}
	b.PutID(WebPageInstantViewTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebPageInstantView) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageInstantView#2c0ec99c as nil")
	}
	b.PutInt(len(w.PageBlocks))
	for idx, v := range w.PageBlocks {
		if v == nil {
			return fmt.Errorf("unable to encode webPageInstantView#2c0ec99c: field page_blocks element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare webPageInstantView#2c0ec99c: field page_blocks element with index %d: %w", idx, err)
		}
	}
	b.PutInt32(w.ViewCount)
	b.PutInt32(w.Version)
	b.PutBool(w.IsRtl)
	b.PutBool(w.IsFull)
	if w.FeedbackLink == nil {
		return fmt.Errorf("unable to encode webPageInstantView#2c0ec99c: field feedback_link is nil")
	}
	if err := w.FeedbackLink.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPageInstantView#2c0ec99c: field feedback_link: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPageInstantView) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageInstantView#2c0ec99c to nil")
	}
	if err := b.ConsumeID(WebPageInstantViewTypeID); err != nil {
		return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebPageInstantView) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageInstantView#2c0ec99c to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field page_blocks: %w", err)
		}

		if headerLen > 0 {
			w.PageBlocks = make([]PageBlockClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePageBlock(b)
			if err != nil {
				return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field page_blocks: %w", err)
			}
			w.PageBlocks = append(w.PageBlocks, value)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field view_count: %w", err)
		}
		w.ViewCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field version: %w", err)
		}
		w.Version = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field is_rtl: %w", err)
		}
		w.IsRtl = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field is_full: %w", err)
		}
		w.IsFull = value
	}
	{
		value, err := DecodeInternalLinkType(b)
		if err != nil {
			return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field feedback_link: %w", err)
		}
		w.FeedbackLink = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (w *WebPageInstantView) EncodeTDLibJSON(b tdjson.Encoder) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageInstantView#2c0ec99c as nil")
	}
	b.ObjStart()
	b.PutID("webPageInstantView")
	b.Comma()
	b.FieldStart("page_blocks")
	b.ArrStart()
	for idx, v := range w.PageBlocks {
		if v == nil {
			return fmt.Errorf("unable to encode webPageInstantView#2c0ec99c: field page_blocks element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode webPageInstantView#2c0ec99c: field page_blocks element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("view_count")
	b.PutInt32(w.ViewCount)
	b.Comma()
	b.FieldStart("version")
	b.PutInt32(w.Version)
	b.Comma()
	b.FieldStart("is_rtl")
	b.PutBool(w.IsRtl)
	b.Comma()
	b.FieldStart("is_full")
	b.PutBool(w.IsFull)
	b.Comma()
	b.FieldStart("feedback_link")
	if w.FeedbackLink == nil {
		return fmt.Errorf("unable to encode webPageInstantView#2c0ec99c: field feedback_link is nil")
	}
	if err := w.FeedbackLink.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webPageInstantView#2c0ec99c: field feedback_link: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (w *WebPageInstantView) DecodeTDLibJSON(b tdjson.Decoder) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageInstantView#2c0ec99c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("webPageInstantView"); err != nil {
				return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: %w", err)
			}
		case "page_blocks":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONPageBlock(b)
				if err != nil {
					return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field page_blocks: %w", err)
				}
				w.PageBlocks = append(w.PageBlocks, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field page_blocks: %w", err)
			}
		case "view_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field view_count: %w", err)
			}
			w.ViewCount = value
		case "version":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field version: %w", err)
			}
			w.Version = value
		case "is_rtl":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field is_rtl: %w", err)
			}
			w.IsRtl = value
		case "is_full":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field is_full: %w", err)
			}
			w.IsFull = value
		case "feedback_link":
			value, err := DecodeTDLibJSONInternalLinkType(b)
			if err != nil {
				return fmt.Errorf("unable to decode webPageInstantView#2c0ec99c: field feedback_link: %w", err)
			}
			w.FeedbackLink = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPageBlocks returns value of PageBlocks field.
func (w *WebPageInstantView) GetPageBlocks() (value []PageBlockClass) {
	if w == nil {
		return
	}
	return w.PageBlocks
}

// GetViewCount returns value of ViewCount field.
func (w *WebPageInstantView) GetViewCount() (value int32) {
	if w == nil {
		return
	}
	return w.ViewCount
}

// GetVersion returns value of Version field.
func (w *WebPageInstantView) GetVersion() (value int32) {
	if w == nil {
		return
	}
	return w.Version
}

// GetIsRtl returns value of IsRtl field.
func (w *WebPageInstantView) GetIsRtl() (value bool) {
	if w == nil {
		return
	}
	return w.IsRtl
}

// GetIsFull returns value of IsFull field.
func (w *WebPageInstantView) GetIsFull() (value bool) {
	if w == nil {
		return
	}
	return w.IsFull
}

// GetFeedbackLink returns value of FeedbackLink field.
func (w *WebPageInstantView) GetFeedbackLink() (value InternalLinkTypeClass) {
	if w == nil {
		return
	}
	return w.FeedbackLink
}
