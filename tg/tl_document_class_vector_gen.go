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

// DocumentClassVector is a box for Vector<Document>
type DocumentClassVector struct {
	// Elements of Vector<Document>
	Elems []DocumentClass
}

// DocumentClassVectorTypeID is TL type id of DocumentClassVector.
const DocumentClassVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for DocumentClassVector.
var (
	_ bin.Encoder     = &DocumentClassVector{}
	_ bin.Decoder     = &DocumentClassVector{}
	_ bin.BareEncoder = &DocumentClassVector{}
	_ bin.BareDecoder = &DocumentClassVector{}
)

func (vec *DocumentClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *DocumentClassVector) String() string {
	if vec == nil {
		return "DocumentClassVector(nil)"
	}
	type Alias DocumentClassVector
	return fmt.Sprintf("DocumentClassVector%+v", Alias(*vec))
}

// FillFrom fills DocumentClassVector from given interface.
func (vec *DocumentClassVector) FillFrom(from interface {
	GetElems() (value []DocumentClass)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DocumentClassVector) TypeID() uint32 {
	return DocumentClassVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*DocumentClassVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *DocumentClassVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   DocumentClassVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *DocumentClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<Document> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *DocumentClassVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<Document> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<Document>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<Document>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *DocumentClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<Document> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *DocumentClassVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<Document> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<Document>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]DocumentClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<Document>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *DocumentClassVector) GetElems() (value []DocumentClass) {
	if vec == nil {
		return
	}
	return vec.Elems
}

// MapElems returns field Elems wrapped in DocumentClassArray helper.
func (vec *DocumentClassVector) MapElems() (value DocumentClassArray) {
	return DocumentClassArray(vec.Elems)
}
