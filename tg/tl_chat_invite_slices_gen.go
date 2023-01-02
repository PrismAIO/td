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

// ChatInviteClassArray is adapter for slice of ChatInviteClass.
type ChatInviteClassArray []ChatInviteClass

// Sort sorts slice of ChatInviteClass.
func (s ChatInviteClassArray) Sort(less func(a, b ChatInviteClass) bool) ChatInviteClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatInviteClass.
func (s ChatInviteClassArray) SortStable(less func(a, b ChatInviteClass) bool) ChatInviteClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatInviteClass.
func (s ChatInviteClassArray) Retain(keep func(x ChatInviteClass) bool) ChatInviteClassArray {
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
func (s ChatInviteClassArray) First() (v ChatInviteClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatInviteClassArray) Last() (v ChatInviteClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatInviteClassArray) PopFirst() (v ChatInviteClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatInviteClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatInviteClassArray) Pop() (v ChatInviteClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsChatInviteAlready returns copy with only ChatInviteAlready constructors.
func (s ChatInviteClassArray) AsChatInviteAlready() (to ChatInviteAlreadyArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatInviteAlready)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsChatInvite returns copy with only ChatInvite constructors.
func (s ChatInviteClassArray) AsChatInvite() (to ChatInviteArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatInvite)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsChatInvitePeek returns copy with only ChatInvitePeek constructors.
func (s ChatInviteClassArray) AsChatInvitePeek() (to ChatInvitePeekArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatInvitePeek)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ChatInviteAlreadyArray is adapter for slice of ChatInviteAlready.
type ChatInviteAlreadyArray []ChatInviteAlready

// Sort sorts slice of ChatInviteAlready.
func (s ChatInviteAlreadyArray) Sort(less func(a, b ChatInviteAlready) bool) ChatInviteAlreadyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatInviteAlready.
func (s ChatInviteAlreadyArray) SortStable(less func(a, b ChatInviteAlready) bool) ChatInviteAlreadyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatInviteAlready.
func (s ChatInviteAlreadyArray) Retain(keep func(x ChatInviteAlready) bool) ChatInviteAlreadyArray {
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
func (s ChatInviteAlreadyArray) First() (v ChatInviteAlready, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatInviteAlreadyArray) Last() (v ChatInviteAlready, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatInviteAlreadyArray) PopFirst() (v ChatInviteAlready, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatInviteAlready
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatInviteAlreadyArray) Pop() (v ChatInviteAlready, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ChatInviteArray is adapter for slice of ChatInvite.
type ChatInviteArray []ChatInvite

// Sort sorts slice of ChatInvite.
func (s ChatInviteArray) Sort(less func(a, b ChatInvite) bool) ChatInviteArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatInvite.
func (s ChatInviteArray) SortStable(less func(a, b ChatInvite) bool) ChatInviteArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatInvite.
func (s ChatInviteArray) Retain(keep func(x ChatInvite) bool) ChatInviteArray {
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
func (s ChatInviteArray) First() (v ChatInvite, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatInviteArray) Last() (v ChatInvite, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatInviteArray) PopFirst() (v ChatInvite, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatInvite
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatInviteArray) Pop() (v ChatInvite, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ChatInvitePeekArray is adapter for slice of ChatInvitePeek.
type ChatInvitePeekArray []ChatInvitePeek

// Sort sorts slice of ChatInvitePeek.
func (s ChatInvitePeekArray) Sort(less func(a, b ChatInvitePeek) bool) ChatInvitePeekArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatInvitePeek.
func (s ChatInvitePeekArray) SortStable(less func(a, b ChatInvitePeek) bool) ChatInvitePeekArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatInvitePeek.
func (s ChatInvitePeekArray) Retain(keep func(x ChatInvitePeek) bool) ChatInvitePeekArray {
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
func (s ChatInvitePeekArray) First() (v ChatInvitePeek, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatInvitePeekArray) Last() (v ChatInvitePeek, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatInvitePeekArray) PopFirst() (v ChatInvitePeek, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatInvitePeek
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatInvitePeekArray) Pop() (v ChatInvitePeek, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
