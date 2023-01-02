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

// MessagesDialogsClassArray is adapter for slice of MessagesDialogsClass.
type MessagesDialogsClassArray []MessagesDialogsClass

// Sort sorts slice of MessagesDialogsClass.
func (s MessagesDialogsClassArray) Sort(less func(a, b MessagesDialogsClass) bool) MessagesDialogsClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesDialogsClass.
func (s MessagesDialogsClassArray) SortStable(less func(a, b MessagesDialogsClass) bool) MessagesDialogsClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesDialogsClass.
func (s MessagesDialogsClassArray) Retain(keep func(x MessagesDialogsClass) bool) MessagesDialogsClassArray {
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
func (s MessagesDialogsClassArray) First() (v MessagesDialogsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesDialogsClassArray) Last() (v MessagesDialogsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesDialogsClassArray) PopFirst() (v MessagesDialogsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesDialogsClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesDialogsClassArray) Pop() (v MessagesDialogsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesDialogs returns copy with only MessagesDialogs constructors.
func (s MessagesDialogsClassArray) AsMessagesDialogs() (to MessagesDialogsArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesDialogs)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessagesDialogsSlice returns copy with only MessagesDialogsSlice constructors.
func (s MessagesDialogsClassArray) AsMessagesDialogsSlice() (to MessagesDialogsSliceArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesDialogsSlice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessagesDialogsNotModified returns copy with only MessagesDialogsNotModified constructors.
func (s MessagesDialogsClassArray) AsMessagesDialogsNotModified() (to MessagesDialogsNotModifiedArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesDialogsNotModified)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s MessagesDialogsClassArray) AppendOnlyModified(to []ModifiedMessagesDialogs) []ModifiedMessagesDialogs {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s MessagesDialogsClassArray) AsModified() (to []ModifiedMessagesDialogs) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s MessagesDialogsClassArray) FirstAsModified() (v ModifiedMessagesDialogs, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s MessagesDialogsClassArray) LastAsModified() (v ModifiedMessagesDialogs, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *MessagesDialogsClassArray) PopFirstAsModified() (v ModifiedMessagesDialogs, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *MessagesDialogsClassArray) PopAsModified() (v ModifiedMessagesDialogs, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// MessagesDialogsArray is adapter for slice of MessagesDialogs.
type MessagesDialogsArray []MessagesDialogs

// Sort sorts slice of MessagesDialogs.
func (s MessagesDialogsArray) Sort(less func(a, b MessagesDialogs) bool) MessagesDialogsArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesDialogs.
func (s MessagesDialogsArray) SortStable(less func(a, b MessagesDialogs) bool) MessagesDialogsArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesDialogs.
func (s MessagesDialogsArray) Retain(keep func(x MessagesDialogs) bool) MessagesDialogsArray {
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
func (s MessagesDialogsArray) First() (v MessagesDialogs, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesDialogsArray) Last() (v MessagesDialogs, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesDialogsArray) PopFirst() (v MessagesDialogs, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesDialogs
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesDialogsArray) Pop() (v MessagesDialogs, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessagesDialogsSliceArray is adapter for slice of MessagesDialogsSlice.
type MessagesDialogsSliceArray []MessagesDialogsSlice

// Sort sorts slice of MessagesDialogsSlice.
func (s MessagesDialogsSliceArray) Sort(less func(a, b MessagesDialogsSlice) bool) MessagesDialogsSliceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesDialogsSlice.
func (s MessagesDialogsSliceArray) SortStable(less func(a, b MessagesDialogsSlice) bool) MessagesDialogsSliceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesDialogsSlice.
func (s MessagesDialogsSliceArray) Retain(keep func(x MessagesDialogsSlice) bool) MessagesDialogsSliceArray {
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
func (s MessagesDialogsSliceArray) First() (v MessagesDialogsSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesDialogsSliceArray) Last() (v MessagesDialogsSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesDialogsSliceArray) PopFirst() (v MessagesDialogsSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesDialogsSlice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesDialogsSliceArray) Pop() (v MessagesDialogsSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessagesDialogsNotModifiedArray is adapter for slice of MessagesDialogsNotModified.
type MessagesDialogsNotModifiedArray []MessagesDialogsNotModified

// Sort sorts slice of MessagesDialogsNotModified.
func (s MessagesDialogsNotModifiedArray) Sort(less func(a, b MessagesDialogsNotModified) bool) MessagesDialogsNotModifiedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesDialogsNotModified.
func (s MessagesDialogsNotModifiedArray) SortStable(less func(a, b MessagesDialogsNotModified) bool) MessagesDialogsNotModifiedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesDialogsNotModified.
func (s MessagesDialogsNotModifiedArray) Retain(keep func(x MessagesDialogsNotModified) bool) MessagesDialogsNotModifiedArray {
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
func (s MessagesDialogsNotModifiedArray) First() (v MessagesDialogsNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesDialogsNotModifiedArray) Last() (v MessagesDialogsNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesDialogsNotModifiedArray) PopFirst() (v MessagesDialogsNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesDialogsNotModified
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesDialogsNotModifiedArray) Pop() (v MessagesDialogsNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
