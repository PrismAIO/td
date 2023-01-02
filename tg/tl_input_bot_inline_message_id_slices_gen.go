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

// InputBotInlineMessageIDClassArray is adapter for slice of InputBotInlineMessageIDClass.
type InputBotInlineMessageIDClassArray []InputBotInlineMessageIDClass

// Sort sorts slice of InputBotInlineMessageIDClass.
func (s InputBotInlineMessageIDClassArray) Sort(less func(a, b InputBotInlineMessageIDClass) bool) InputBotInlineMessageIDClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputBotInlineMessageIDClass.
func (s InputBotInlineMessageIDClassArray) SortStable(less func(a, b InputBotInlineMessageIDClass) bool) InputBotInlineMessageIDClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputBotInlineMessageIDClass.
func (s InputBotInlineMessageIDClassArray) Retain(keep func(x InputBotInlineMessageIDClass) bool) InputBotInlineMessageIDClassArray {
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
func (s InputBotInlineMessageIDClassArray) First() (v InputBotInlineMessageIDClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputBotInlineMessageIDClassArray) Last() (v InputBotInlineMessageIDClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputBotInlineMessageIDClassArray) PopFirst() (v InputBotInlineMessageIDClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputBotInlineMessageIDClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputBotInlineMessageIDClassArray) Pop() (v InputBotInlineMessageIDClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputBotInlineMessageID returns copy with only InputBotInlineMessageID constructors.
func (s InputBotInlineMessageIDClassArray) AsInputBotInlineMessageID() (to InputBotInlineMessageIDArray) {
	for _, elem := range s {
		value, ok := elem.(*InputBotInlineMessageID)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputBotInlineMessageID64 returns copy with only InputBotInlineMessageID64 constructors.
func (s InputBotInlineMessageIDClassArray) AsInputBotInlineMessageID64() (to InputBotInlineMessageID64Array) {
	for _, elem := range s {
		value, ok := elem.(*InputBotInlineMessageID64)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputBotInlineMessageIDArray is adapter for slice of InputBotInlineMessageID.
type InputBotInlineMessageIDArray []InputBotInlineMessageID

// Sort sorts slice of InputBotInlineMessageID.
func (s InputBotInlineMessageIDArray) Sort(less func(a, b InputBotInlineMessageID) bool) InputBotInlineMessageIDArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputBotInlineMessageID.
func (s InputBotInlineMessageIDArray) SortStable(less func(a, b InputBotInlineMessageID) bool) InputBotInlineMessageIDArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputBotInlineMessageID.
func (s InputBotInlineMessageIDArray) Retain(keep func(x InputBotInlineMessageID) bool) InputBotInlineMessageIDArray {
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
func (s InputBotInlineMessageIDArray) First() (v InputBotInlineMessageID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputBotInlineMessageIDArray) Last() (v InputBotInlineMessageID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputBotInlineMessageIDArray) PopFirst() (v InputBotInlineMessageID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputBotInlineMessageID
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputBotInlineMessageIDArray) Pop() (v InputBotInlineMessageID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputBotInlineMessageID by ID.
func (s InputBotInlineMessageIDArray) SortByID() InputBotInlineMessageIDArray {
	return s.Sort(func(a, b InputBotInlineMessageID) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputBotInlineMessageID by ID.
func (s InputBotInlineMessageIDArray) SortStableByID() InputBotInlineMessageIDArray {
	return s.SortStable(func(a, b InputBotInlineMessageID) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputBotInlineMessageIDArray) FillMap(to map[int64]InputBotInlineMessageID) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputBotInlineMessageIDArray) ToMap() map[int64]InputBotInlineMessageID {
	r := make(map[int64]InputBotInlineMessageID, len(s))
	s.FillMap(r)
	return r
}

// InputBotInlineMessageID64Array is adapter for slice of InputBotInlineMessageID64.
type InputBotInlineMessageID64Array []InputBotInlineMessageID64

// Sort sorts slice of InputBotInlineMessageID64.
func (s InputBotInlineMessageID64Array) Sort(less func(a, b InputBotInlineMessageID64) bool) InputBotInlineMessageID64Array {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputBotInlineMessageID64.
func (s InputBotInlineMessageID64Array) SortStable(less func(a, b InputBotInlineMessageID64) bool) InputBotInlineMessageID64Array {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputBotInlineMessageID64.
func (s InputBotInlineMessageID64Array) Retain(keep func(x InputBotInlineMessageID64) bool) InputBotInlineMessageID64Array {
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
func (s InputBotInlineMessageID64Array) First() (v InputBotInlineMessageID64, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputBotInlineMessageID64Array) Last() (v InputBotInlineMessageID64, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputBotInlineMessageID64Array) PopFirst() (v InputBotInlineMessageID64, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputBotInlineMessageID64
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputBotInlineMessageID64Array) Pop() (v InputBotInlineMessageID64, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputBotInlineMessageID64 by ID.
func (s InputBotInlineMessageID64Array) SortByID() InputBotInlineMessageID64Array {
	return s.Sort(func(a, b InputBotInlineMessageID64) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputBotInlineMessageID64 by ID.
func (s InputBotInlineMessageID64Array) SortStableByID() InputBotInlineMessageID64Array {
	return s.SortStable(func(a, b InputBotInlineMessageID64) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputBotInlineMessageID64Array) FillMap(to map[int]InputBotInlineMessageID64) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputBotInlineMessageID64Array) ToMap() map[int]InputBotInlineMessageID64 {
	r := make(map[int]InputBotInlineMessageID64, len(s))
	s.FillMap(r)
	return r
}
