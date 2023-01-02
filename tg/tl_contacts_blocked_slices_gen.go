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

// ContactsBlockedClassArray is adapter for slice of ContactsBlockedClass.
type ContactsBlockedClassArray []ContactsBlockedClass

// Sort sorts slice of ContactsBlockedClass.
func (s ContactsBlockedClassArray) Sort(less func(a, b ContactsBlockedClass) bool) ContactsBlockedClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ContactsBlockedClass.
func (s ContactsBlockedClassArray) SortStable(less func(a, b ContactsBlockedClass) bool) ContactsBlockedClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ContactsBlockedClass.
func (s ContactsBlockedClassArray) Retain(keep func(x ContactsBlockedClass) bool) ContactsBlockedClassArray {
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
func (s ContactsBlockedClassArray) First() (v ContactsBlockedClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ContactsBlockedClassArray) Last() (v ContactsBlockedClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ContactsBlockedClassArray) PopFirst() (v ContactsBlockedClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ContactsBlockedClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ContactsBlockedClassArray) Pop() (v ContactsBlockedClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsContactsBlocked returns copy with only ContactsBlocked constructors.
func (s ContactsBlockedClassArray) AsContactsBlocked() (to ContactsBlockedArray) {
	for _, elem := range s {
		value, ok := elem.(*ContactsBlocked)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsContactsBlockedSlice returns copy with only ContactsBlockedSlice constructors.
func (s ContactsBlockedClassArray) AsContactsBlockedSlice() (to ContactsBlockedSliceArray) {
	for _, elem := range s {
		value, ok := elem.(*ContactsBlockedSlice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ContactsBlockedArray is adapter for slice of ContactsBlocked.
type ContactsBlockedArray []ContactsBlocked

// Sort sorts slice of ContactsBlocked.
func (s ContactsBlockedArray) Sort(less func(a, b ContactsBlocked) bool) ContactsBlockedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ContactsBlocked.
func (s ContactsBlockedArray) SortStable(less func(a, b ContactsBlocked) bool) ContactsBlockedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ContactsBlocked.
func (s ContactsBlockedArray) Retain(keep func(x ContactsBlocked) bool) ContactsBlockedArray {
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
func (s ContactsBlockedArray) First() (v ContactsBlocked, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ContactsBlockedArray) Last() (v ContactsBlocked, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ContactsBlockedArray) PopFirst() (v ContactsBlocked, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ContactsBlocked
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ContactsBlockedArray) Pop() (v ContactsBlocked, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ContactsBlockedSliceArray is adapter for slice of ContactsBlockedSlice.
type ContactsBlockedSliceArray []ContactsBlockedSlice

// Sort sorts slice of ContactsBlockedSlice.
func (s ContactsBlockedSliceArray) Sort(less func(a, b ContactsBlockedSlice) bool) ContactsBlockedSliceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ContactsBlockedSlice.
func (s ContactsBlockedSliceArray) SortStable(less func(a, b ContactsBlockedSlice) bool) ContactsBlockedSliceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ContactsBlockedSlice.
func (s ContactsBlockedSliceArray) Retain(keep func(x ContactsBlockedSlice) bool) ContactsBlockedSliceArray {
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
func (s ContactsBlockedSliceArray) First() (v ContactsBlockedSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ContactsBlockedSliceArray) Last() (v ContactsBlockedSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ContactsBlockedSliceArray) PopFirst() (v ContactsBlockedSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ContactsBlockedSlice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ContactsBlockedSliceArray) Pop() (v ContactsBlockedSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
