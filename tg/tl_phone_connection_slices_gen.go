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

// PhoneConnectionClassArray is adapter for slice of PhoneConnectionClass.
type PhoneConnectionClassArray []PhoneConnectionClass

// Sort sorts slice of PhoneConnectionClass.
func (s PhoneConnectionClassArray) Sort(less func(a, b PhoneConnectionClass) bool) PhoneConnectionClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneConnectionClass.
func (s PhoneConnectionClassArray) SortStable(less func(a, b PhoneConnectionClass) bool) PhoneConnectionClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneConnectionClass.
func (s PhoneConnectionClassArray) Retain(keep func(x PhoneConnectionClass) bool) PhoneConnectionClassArray {
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
func (s PhoneConnectionClassArray) First() (v PhoneConnectionClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneConnectionClassArray) Last() (v PhoneConnectionClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneConnectionClassArray) PopFirst() (v PhoneConnectionClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneConnectionClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneConnectionClassArray) Pop() (v PhoneConnectionClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhoneConnectionClass by ID.
func (s PhoneConnectionClassArray) SortByID() PhoneConnectionClassArray {
	return s.Sort(func(a, b PhoneConnectionClass) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhoneConnectionClass by ID.
func (s PhoneConnectionClassArray) SortStableByID() PhoneConnectionClassArray {
	return s.SortStable(func(a, b PhoneConnectionClass) bool {
		return a.GetID() < b.GetID()
	})
}

// FillPhoneConnectionMap fills only PhoneConnection constructors to given map.
func (s PhoneConnectionClassArray) FillPhoneConnectionMap(to map[int64]*PhoneConnection) {
	for _, elem := range s {
		value, ok := elem.(*PhoneConnection)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// PhoneConnectionToMap collects only PhoneConnection constructors to map.
func (s PhoneConnectionClassArray) PhoneConnectionToMap() map[int64]*PhoneConnection {
	r := make(map[int64]*PhoneConnection, len(s))
	s.FillPhoneConnectionMap(r)
	return r
}

// AsPhoneConnection returns copy with only PhoneConnection constructors.
func (s PhoneConnectionClassArray) AsPhoneConnection() (to PhoneConnectionArray) {
	for _, elem := range s {
		value, ok := elem.(*PhoneConnection)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillPhoneConnectionWebrtcMap fills only PhoneConnectionWebrtc constructors to given map.
func (s PhoneConnectionClassArray) FillPhoneConnectionWebrtcMap(to map[int64]*PhoneConnectionWebrtc) {
	for _, elem := range s {
		value, ok := elem.(*PhoneConnectionWebrtc)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// PhoneConnectionWebrtcToMap collects only PhoneConnectionWebrtc constructors to map.
func (s PhoneConnectionClassArray) PhoneConnectionWebrtcToMap() map[int64]*PhoneConnectionWebrtc {
	r := make(map[int64]*PhoneConnectionWebrtc, len(s))
	s.FillPhoneConnectionWebrtcMap(r)
	return r
}

// AsPhoneConnectionWebrtc returns copy with only PhoneConnectionWebrtc constructors.
func (s PhoneConnectionClassArray) AsPhoneConnectionWebrtc() (to PhoneConnectionWebrtcArray) {
	for _, elem := range s {
		value, ok := elem.(*PhoneConnectionWebrtc)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// PhoneConnectionArray is adapter for slice of PhoneConnection.
type PhoneConnectionArray []PhoneConnection

// Sort sorts slice of PhoneConnection.
func (s PhoneConnectionArray) Sort(less func(a, b PhoneConnection) bool) PhoneConnectionArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneConnection.
func (s PhoneConnectionArray) SortStable(less func(a, b PhoneConnection) bool) PhoneConnectionArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneConnection.
func (s PhoneConnectionArray) Retain(keep func(x PhoneConnection) bool) PhoneConnectionArray {
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
func (s PhoneConnectionArray) First() (v PhoneConnection, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneConnectionArray) Last() (v PhoneConnection, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneConnectionArray) PopFirst() (v PhoneConnection, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneConnection
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneConnectionArray) Pop() (v PhoneConnection, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhoneConnection by ID.
func (s PhoneConnectionArray) SortByID() PhoneConnectionArray {
	return s.Sort(func(a, b PhoneConnection) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhoneConnection by ID.
func (s PhoneConnectionArray) SortStableByID() PhoneConnectionArray {
	return s.SortStable(func(a, b PhoneConnection) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s PhoneConnectionArray) FillMap(to map[int64]PhoneConnection) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s PhoneConnectionArray) ToMap() map[int64]PhoneConnection {
	r := make(map[int64]PhoneConnection, len(s))
	s.FillMap(r)
	return r
}

// PhoneConnectionWebrtcArray is adapter for slice of PhoneConnectionWebrtc.
type PhoneConnectionWebrtcArray []PhoneConnectionWebrtc

// Sort sorts slice of PhoneConnectionWebrtc.
func (s PhoneConnectionWebrtcArray) Sort(less func(a, b PhoneConnectionWebrtc) bool) PhoneConnectionWebrtcArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneConnectionWebrtc.
func (s PhoneConnectionWebrtcArray) SortStable(less func(a, b PhoneConnectionWebrtc) bool) PhoneConnectionWebrtcArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneConnectionWebrtc.
func (s PhoneConnectionWebrtcArray) Retain(keep func(x PhoneConnectionWebrtc) bool) PhoneConnectionWebrtcArray {
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
func (s PhoneConnectionWebrtcArray) First() (v PhoneConnectionWebrtc, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneConnectionWebrtcArray) Last() (v PhoneConnectionWebrtc, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneConnectionWebrtcArray) PopFirst() (v PhoneConnectionWebrtc, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneConnectionWebrtc
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneConnectionWebrtcArray) Pop() (v PhoneConnectionWebrtc, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhoneConnectionWebrtc by ID.
func (s PhoneConnectionWebrtcArray) SortByID() PhoneConnectionWebrtcArray {
	return s.Sort(func(a, b PhoneConnectionWebrtc) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhoneConnectionWebrtc by ID.
func (s PhoneConnectionWebrtcArray) SortStableByID() PhoneConnectionWebrtcArray {
	return s.SortStable(func(a, b PhoneConnectionWebrtc) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s PhoneConnectionWebrtcArray) FillMap(to map[int64]PhoneConnectionWebrtc) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s PhoneConnectionWebrtcArray) ToMap() map[int64]PhoneConnectionWebrtc {
	r := make(map[int64]PhoneConnectionWebrtc, len(s))
	s.FillMap(r)
	return r
}
