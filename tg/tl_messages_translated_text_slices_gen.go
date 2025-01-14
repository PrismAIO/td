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

// MessagesTranslatedTextClassArray is adapter for slice of MessagesTranslatedTextClass.
type MessagesTranslatedTextClassArray []MessagesTranslatedTextClass

// Sort sorts slice of MessagesTranslatedTextClass.
func (s MessagesTranslatedTextClassArray) Sort(less func(a, b MessagesTranslatedTextClass) bool) MessagesTranslatedTextClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesTranslatedTextClass.
func (s MessagesTranslatedTextClassArray) SortStable(less func(a, b MessagesTranslatedTextClass) bool) MessagesTranslatedTextClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesTranslatedTextClass.
func (s MessagesTranslatedTextClassArray) Retain(keep func(x MessagesTranslatedTextClass) bool) MessagesTranslatedTextClassArray {
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
func (s MessagesTranslatedTextClassArray) First() (v MessagesTranslatedTextClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesTranslatedTextClassArray) Last() (v MessagesTranslatedTextClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesTranslatedTextClassArray) PopFirst() (v MessagesTranslatedTextClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesTranslatedTextClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesTranslatedTextClassArray) Pop() (v MessagesTranslatedTextClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesTranslateResultText returns copy with only MessagesTranslateResultText constructors.
func (s MessagesTranslatedTextClassArray) AsMessagesTranslateResultText() (to MessagesTranslateResultTextArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesTranslateResultText)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// MessagesTranslateResultTextArray is adapter for slice of MessagesTranslateResultText.
type MessagesTranslateResultTextArray []MessagesTranslateResultText

// Sort sorts slice of MessagesTranslateResultText.
func (s MessagesTranslateResultTextArray) Sort(less func(a, b MessagesTranslateResultText) bool) MessagesTranslateResultTextArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesTranslateResultText.
func (s MessagesTranslateResultTextArray) SortStable(less func(a, b MessagesTranslateResultText) bool) MessagesTranslateResultTextArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesTranslateResultText.
func (s MessagesTranslateResultTextArray) Retain(keep func(x MessagesTranslateResultText) bool) MessagesTranslateResultTextArray {
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
func (s MessagesTranslateResultTextArray) First() (v MessagesTranslateResultText, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesTranslateResultTextArray) Last() (v MessagesTranslateResultText, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesTranslateResultTextArray) PopFirst() (v MessagesTranslateResultText, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesTranslateResultText
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesTranslateResultTextArray) Pop() (v MessagesTranslateResultText, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
