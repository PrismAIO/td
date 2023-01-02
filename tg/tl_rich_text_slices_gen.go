//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// RichTextClassArray is adapter for slice of RichTextClass.
type RichTextClassArray []RichTextClass

// Sort sorts slice of RichTextClass.
func (s RichTextClassArray) Sort(less func(a, b RichTextClass) bool) RichTextClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RichTextClass.
func (s RichTextClassArray) SortStable(less func(a, b RichTextClass) bool) RichTextClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RichTextClass.
func (s RichTextClassArray) Retain(keep func(x RichTextClass) bool) RichTextClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s RichTextClassArray) First() (v RichTextClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RichTextClassArray) Last() (v RichTextClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RichTextClassArray) PopFirst() (v RichTextClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RichTextClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RichTextClassArray) Pop() (v RichTextClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsTextPlain returns copy with only TextPlain constructors.
func (s RichTextClassArray) AsTextPlain() (to TextPlainArray) {
	for _, elem := range s {
		value, ok := elem.(*TextPlain)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextBold returns copy with only TextBold constructors.
func (s RichTextClassArray) AsTextBold() (to TextBoldArray) {
	for _, elem := range s {
		value, ok := elem.(*TextBold)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextItalic returns copy with only TextItalic constructors.
func (s RichTextClassArray) AsTextItalic() (to TextItalicArray) {
	for _, elem := range s {
		value, ok := elem.(*TextItalic)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextUnderline returns copy with only TextUnderline constructors.
func (s RichTextClassArray) AsTextUnderline() (to TextUnderlineArray) {
	for _, elem := range s {
		value, ok := elem.(*TextUnderline)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextStrike returns copy with only TextStrike constructors.
func (s RichTextClassArray) AsTextStrike() (to TextStrikeArray) {
	for _, elem := range s {
		value, ok := elem.(*TextStrike)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextFixed returns copy with only TextFixed constructors.
func (s RichTextClassArray) AsTextFixed() (to TextFixedArray) {
	for _, elem := range s {
		value, ok := elem.(*TextFixed)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextURL returns copy with only TextURL constructors.
func (s RichTextClassArray) AsTextURL() (to TextURLArray) {
	for _, elem := range s {
		value, ok := elem.(*TextURL)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextEmail returns copy with only TextEmail constructors.
func (s RichTextClassArray) AsTextEmail() (to TextEmailArray) {
	for _, elem := range s {
		value, ok := elem.(*TextEmail)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextConcat returns copy with only TextConcat constructors.
func (s RichTextClassArray) AsTextConcat() (to TextConcatArray) {
	for _, elem := range s {
		value, ok := elem.(*TextConcat)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextSubscript returns copy with only TextSubscript constructors.
func (s RichTextClassArray) AsTextSubscript() (to TextSubscriptArray) {
	for _, elem := range s {
		value, ok := elem.(*TextSubscript)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextSuperscript returns copy with only TextSuperscript constructors.
func (s RichTextClassArray) AsTextSuperscript() (to TextSuperscriptArray) {
	for _, elem := range s {
		value, ok := elem.(*TextSuperscript)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextMarked returns copy with only TextMarked constructors.
func (s RichTextClassArray) AsTextMarked() (to TextMarkedArray) {
	for _, elem := range s {
		value, ok := elem.(*TextMarked)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextPhone returns copy with only TextPhone constructors.
func (s RichTextClassArray) AsTextPhone() (to TextPhoneArray) {
	for _, elem := range s {
		value, ok := elem.(*TextPhone)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextImage returns copy with only TextImage constructors.
func (s RichTextClassArray) AsTextImage() (to TextImageArray) {
	for _, elem := range s {
		value, ok := elem.(*TextImage)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsTextAnchor returns copy with only TextAnchor constructors.
func (s RichTextClassArray) AsTextAnchor() (to TextAnchorArray) {
	for _, elem := range s {
		value, ok := elem.(*TextAnchor)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// TextPlainArray is adapter for slice of TextPlain.
type TextPlainArray []TextPlain

// Sort sorts slice of TextPlain.
func (s TextPlainArray) Sort(less func(a, b TextPlain) bool) TextPlainArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextPlain.
func (s TextPlainArray) SortStable(less func(a, b TextPlain) bool) TextPlainArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextPlain.
func (s TextPlainArray) Retain(keep func(x TextPlain) bool) TextPlainArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextPlainArray) First() (v TextPlain, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextPlainArray) Last() (v TextPlain, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextPlainArray) PopFirst() (v TextPlain, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextPlain
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextPlainArray) Pop() (v TextPlain, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextBoldArray is adapter for slice of TextBold.
type TextBoldArray []TextBold

// Sort sorts slice of TextBold.
func (s TextBoldArray) Sort(less func(a, b TextBold) bool) TextBoldArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextBold.
func (s TextBoldArray) SortStable(less func(a, b TextBold) bool) TextBoldArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextBold.
func (s TextBoldArray) Retain(keep func(x TextBold) bool) TextBoldArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextBoldArray) First() (v TextBold, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextBoldArray) Last() (v TextBold, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextBoldArray) PopFirst() (v TextBold, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextBold
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextBoldArray) Pop() (v TextBold, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextItalicArray is adapter for slice of TextItalic.
type TextItalicArray []TextItalic

// Sort sorts slice of TextItalic.
func (s TextItalicArray) Sort(less func(a, b TextItalic) bool) TextItalicArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextItalic.
func (s TextItalicArray) SortStable(less func(a, b TextItalic) bool) TextItalicArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextItalic.
func (s TextItalicArray) Retain(keep func(x TextItalic) bool) TextItalicArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextItalicArray) First() (v TextItalic, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextItalicArray) Last() (v TextItalic, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextItalicArray) PopFirst() (v TextItalic, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextItalic
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextItalicArray) Pop() (v TextItalic, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextUnderlineArray is adapter for slice of TextUnderline.
type TextUnderlineArray []TextUnderline

// Sort sorts slice of TextUnderline.
func (s TextUnderlineArray) Sort(less func(a, b TextUnderline) bool) TextUnderlineArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextUnderline.
func (s TextUnderlineArray) SortStable(less func(a, b TextUnderline) bool) TextUnderlineArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextUnderline.
func (s TextUnderlineArray) Retain(keep func(x TextUnderline) bool) TextUnderlineArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextUnderlineArray) First() (v TextUnderline, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextUnderlineArray) Last() (v TextUnderline, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextUnderlineArray) PopFirst() (v TextUnderline, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextUnderline
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextUnderlineArray) Pop() (v TextUnderline, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextStrikeArray is adapter for slice of TextStrike.
type TextStrikeArray []TextStrike

// Sort sorts slice of TextStrike.
func (s TextStrikeArray) Sort(less func(a, b TextStrike) bool) TextStrikeArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextStrike.
func (s TextStrikeArray) SortStable(less func(a, b TextStrike) bool) TextStrikeArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextStrike.
func (s TextStrikeArray) Retain(keep func(x TextStrike) bool) TextStrikeArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextStrikeArray) First() (v TextStrike, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextStrikeArray) Last() (v TextStrike, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextStrikeArray) PopFirst() (v TextStrike, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextStrike
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextStrikeArray) Pop() (v TextStrike, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextFixedArray is adapter for slice of TextFixed.
type TextFixedArray []TextFixed

// Sort sorts slice of TextFixed.
func (s TextFixedArray) Sort(less func(a, b TextFixed) bool) TextFixedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextFixed.
func (s TextFixedArray) SortStable(less func(a, b TextFixed) bool) TextFixedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextFixed.
func (s TextFixedArray) Retain(keep func(x TextFixed) bool) TextFixedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextFixedArray) First() (v TextFixed, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextFixedArray) Last() (v TextFixed, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextFixedArray) PopFirst() (v TextFixed, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextFixed
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextFixedArray) Pop() (v TextFixed, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextURLArray is adapter for slice of TextURL.
type TextURLArray []TextURL

// Sort sorts slice of TextURL.
func (s TextURLArray) Sort(less func(a, b TextURL) bool) TextURLArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextURL.
func (s TextURLArray) SortStable(less func(a, b TextURL) bool) TextURLArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextURL.
func (s TextURLArray) Retain(keep func(x TextURL) bool) TextURLArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextURLArray) First() (v TextURL, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextURLArray) Last() (v TextURL, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextURLArray) PopFirst() (v TextURL, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextURL
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextURLArray) Pop() (v TextURL, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextEmailArray is adapter for slice of TextEmail.
type TextEmailArray []TextEmail

// Sort sorts slice of TextEmail.
func (s TextEmailArray) Sort(less func(a, b TextEmail) bool) TextEmailArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextEmail.
func (s TextEmailArray) SortStable(less func(a, b TextEmail) bool) TextEmailArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextEmail.
func (s TextEmailArray) Retain(keep func(x TextEmail) bool) TextEmailArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextEmailArray) First() (v TextEmail, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextEmailArray) Last() (v TextEmail, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextEmailArray) PopFirst() (v TextEmail, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextEmail
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextEmailArray) Pop() (v TextEmail, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextConcatArray is adapter for slice of TextConcat.
type TextConcatArray []TextConcat

// Sort sorts slice of TextConcat.
func (s TextConcatArray) Sort(less func(a, b TextConcat) bool) TextConcatArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextConcat.
func (s TextConcatArray) SortStable(less func(a, b TextConcat) bool) TextConcatArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextConcat.
func (s TextConcatArray) Retain(keep func(x TextConcat) bool) TextConcatArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextConcatArray) First() (v TextConcat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextConcatArray) Last() (v TextConcat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextConcatArray) PopFirst() (v TextConcat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextConcat
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextConcatArray) Pop() (v TextConcat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextSubscriptArray is adapter for slice of TextSubscript.
type TextSubscriptArray []TextSubscript

// Sort sorts slice of TextSubscript.
func (s TextSubscriptArray) Sort(less func(a, b TextSubscript) bool) TextSubscriptArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextSubscript.
func (s TextSubscriptArray) SortStable(less func(a, b TextSubscript) bool) TextSubscriptArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextSubscript.
func (s TextSubscriptArray) Retain(keep func(x TextSubscript) bool) TextSubscriptArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextSubscriptArray) First() (v TextSubscript, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextSubscriptArray) Last() (v TextSubscript, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextSubscriptArray) PopFirst() (v TextSubscript, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextSubscript
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextSubscriptArray) Pop() (v TextSubscript, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextSuperscriptArray is adapter for slice of TextSuperscript.
type TextSuperscriptArray []TextSuperscript

// Sort sorts slice of TextSuperscript.
func (s TextSuperscriptArray) Sort(less func(a, b TextSuperscript) bool) TextSuperscriptArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextSuperscript.
func (s TextSuperscriptArray) SortStable(less func(a, b TextSuperscript) bool) TextSuperscriptArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextSuperscript.
func (s TextSuperscriptArray) Retain(keep func(x TextSuperscript) bool) TextSuperscriptArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextSuperscriptArray) First() (v TextSuperscript, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextSuperscriptArray) Last() (v TextSuperscript, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextSuperscriptArray) PopFirst() (v TextSuperscript, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextSuperscript
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextSuperscriptArray) Pop() (v TextSuperscript, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextMarkedArray is adapter for slice of TextMarked.
type TextMarkedArray []TextMarked

// Sort sorts slice of TextMarked.
func (s TextMarkedArray) Sort(less func(a, b TextMarked) bool) TextMarkedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextMarked.
func (s TextMarkedArray) SortStable(less func(a, b TextMarked) bool) TextMarkedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextMarked.
func (s TextMarkedArray) Retain(keep func(x TextMarked) bool) TextMarkedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextMarkedArray) First() (v TextMarked, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextMarkedArray) Last() (v TextMarked, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextMarkedArray) PopFirst() (v TextMarked, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextMarked
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextMarkedArray) Pop() (v TextMarked, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextPhoneArray is adapter for slice of TextPhone.
type TextPhoneArray []TextPhone

// Sort sorts slice of TextPhone.
func (s TextPhoneArray) Sort(less func(a, b TextPhone) bool) TextPhoneArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextPhone.
func (s TextPhoneArray) SortStable(less func(a, b TextPhone) bool) TextPhoneArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextPhone.
func (s TextPhoneArray) Retain(keep func(x TextPhone) bool) TextPhoneArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextPhoneArray) First() (v TextPhone, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextPhoneArray) Last() (v TextPhone, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextPhoneArray) PopFirst() (v TextPhone, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextPhone
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextPhoneArray) Pop() (v TextPhone, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextImageArray is adapter for slice of TextImage.
type TextImageArray []TextImage

// Sort sorts slice of TextImage.
func (s TextImageArray) Sort(less func(a, b TextImage) bool) TextImageArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextImage.
func (s TextImageArray) SortStable(less func(a, b TextImage) bool) TextImageArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextImage.
func (s TextImageArray) Retain(keep func(x TextImage) bool) TextImageArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextImageArray) First() (v TextImage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextImageArray) Last() (v TextImage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextImageArray) PopFirst() (v TextImage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextImage
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextImageArray) Pop() (v TextImage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// TextAnchorArray is adapter for slice of TextAnchor.
type TextAnchorArray []TextAnchor

// Sort sorts slice of TextAnchor.
func (s TextAnchorArray) Sort(less func(a, b TextAnchor) bool) TextAnchorArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of TextAnchor.
func (s TextAnchorArray) SortStable(less func(a, b TextAnchor) bool) TextAnchorArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of TextAnchor.
func (s TextAnchorArray) Retain(keep func(x TextAnchor) bool) TextAnchorArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s TextAnchorArray) First() (v TextAnchor, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s TextAnchorArray) Last() (v TextAnchor, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *TextAnchorArray) PopFirst() (v TextAnchor, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero TextAnchor
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *TextAnchorArray) Pop() (v TextAnchor, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
