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

// PageListItemClassArray is adapter for slice of PageListItemClass.
type PageListItemClassArray []PageListItemClass

// Sort sorts slice of PageListItemClass.
func (s PageListItemClassArray) Sort(less func(a, b PageListItemClass) bool) PageListItemClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PageListItemClass.
func (s PageListItemClassArray) SortStable(less func(a, b PageListItemClass) bool) PageListItemClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PageListItemClass.
func (s PageListItemClassArray) Retain(keep func(x PageListItemClass) bool) PageListItemClassArray {
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
func (s PageListItemClassArray) First() (v PageListItemClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PageListItemClassArray) Last() (v PageListItemClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PageListItemClassArray) PopFirst() (v PageListItemClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PageListItemClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PageListItemClassArray) Pop() (v PageListItemClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsPageListItemText returns copy with only PageListItemText constructors.
func (s PageListItemClassArray) AsPageListItemText() (to PageListItemTextArray) {
	for _, elem := range s {
		value, ok := elem.(*PageListItemText)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPageListItemBlocks returns copy with only PageListItemBlocks constructors.
func (s PageListItemClassArray) AsPageListItemBlocks() (to PageListItemBlocksArray) {
	for _, elem := range s {
		value, ok := elem.(*PageListItemBlocks)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// PageListItemTextArray is adapter for slice of PageListItemText.
type PageListItemTextArray []PageListItemText

// Sort sorts slice of PageListItemText.
func (s PageListItemTextArray) Sort(less func(a, b PageListItemText) bool) PageListItemTextArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PageListItemText.
func (s PageListItemTextArray) SortStable(less func(a, b PageListItemText) bool) PageListItemTextArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PageListItemText.
func (s PageListItemTextArray) Retain(keep func(x PageListItemText) bool) PageListItemTextArray {
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
func (s PageListItemTextArray) First() (v PageListItemText, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PageListItemTextArray) Last() (v PageListItemText, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PageListItemTextArray) PopFirst() (v PageListItemText, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PageListItemText
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PageListItemTextArray) Pop() (v PageListItemText, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PageListItemBlocksArray is adapter for slice of PageListItemBlocks.
type PageListItemBlocksArray []PageListItemBlocks

// Sort sorts slice of PageListItemBlocks.
func (s PageListItemBlocksArray) Sort(less func(a, b PageListItemBlocks) bool) PageListItemBlocksArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PageListItemBlocks.
func (s PageListItemBlocksArray) SortStable(less func(a, b PageListItemBlocks) bool) PageListItemBlocksArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PageListItemBlocks.
func (s PageListItemBlocksArray) Retain(keep func(x PageListItemBlocks) bool) PageListItemBlocksArray {
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
func (s PageListItemBlocksArray) First() (v PageListItemBlocks, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PageListItemBlocksArray) Last() (v PageListItemBlocks, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PageListItemBlocksArray) PopFirst() (v PageListItemBlocks, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PageListItemBlocks
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PageListItemBlocksArray) Pop() (v PageListItemBlocks, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
