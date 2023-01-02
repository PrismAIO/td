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

// AccountSavedRingtonesClassArray is adapter for slice of AccountSavedRingtonesClass.
type AccountSavedRingtonesClassArray []AccountSavedRingtonesClass

// Sort sorts slice of AccountSavedRingtonesClass.
func (s AccountSavedRingtonesClassArray) Sort(less func(a, b AccountSavedRingtonesClass) bool) AccountSavedRingtonesClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AccountSavedRingtonesClass.
func (s AccountSavedRingtonesClassArray) SortStable(less func(a, b AccountSavedRingtonesClass) bool) AccountSavedRingtonesClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AccountSavedRingtonesClass.
func (s AccountSavedRingtonesClassArray) Retain(keep func(x AccountSavedRingtonesClass) bool) AccountSavedRingtonesClassArray {
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
func (s AccountSavedRingtonesClassArray) First() (v AccountSavedRingtonesClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AccountSavedRingtonesClassArray) Last() (v AccountSavedRingtonesClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AccountSavedRingtonesClassArray) PopFirst() (v AccountSavedRingtonesClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AccountSavedRingtonesClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AccountSavedRingtonesClassArray) Pop() (v AccountSavedRingtonesClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsAccountSavedRingtones returns copy with only AccountSavedRingtones constructors.
func (s AccountSavedRingtonesClassArray) AsAccountSavedRingtones() (to AccountSavedRingtonesArray) {
	for _, elem := range s {
		value, ok := elem.(*AccountSavedRingtones)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s AccountSavedRingtonesClassArray) AppendOnlyModified(to []*AccountSavedRingtones) []*AccountSavedRingtones {
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
func (s AccountSavedRingtonesClassArray) AsModified() (to []*AccountSavedRingtones) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s AccountSavedRingtonesClassArray) FirstAsModified() (v *AccountSavedRingtones, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s AccountSavedRingtonesClassArray) LastAsModified() (v *AccountSavedRingtones, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *AccountSavedRingtonesClassArray) PopFirstAsModified() (v *AccountSavedRingtones, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *AccountSavedRingtonesClassArray) PopAsModified() (v *AccountSavedRingtones, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// AccountSavedRingtonesArray is adapter for slice of AccountSavedRingtones.
type AccountSavedRingtonesArray []AccountSavedRingtones

// Sort sorts slice of AccountSavedRingtones.
func (s AccountSavedRingtonesArray) Sort(less func(a, b AccountSavedRingtones) bool) AccountSavedRingtonesArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AccountSavedRingtones.
func (s AccountSavedRingtonesArray) SortStable(less func(a, b AccountSavedRingtones) bool) AccountSavedRingtonesArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AccountSavedRingtones.
func (s AccountSavedRingtonesArray) Retain(keep func(x AccountSavedRingtones) bool) AccountSavedRingtonesArray {
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
func (s AccountSavedRingtonesArray) First() (v AccountSavedRingtones, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AccountSavedRingtonesArray) Last() (v AccountSavedRingtones, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AccountSavedRingtonesArray) PopFirst() (v AccountSavedRingtones, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AccountSavedRingtones
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AccountSavedRingtonesArray) Pop() (v AccountSavedRingtones, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
