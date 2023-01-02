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

// InputThemeClassArray is adapter for slice of InputThemeClass.
type InputThemeClassArray []InputThemeClass

// Sort sorts slice of InputThemeClass.
func (s InputThemeClassArray) Sort(less func(a, b InputThemeClass) bool) InputThemeClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputThemeClass.
func (s InputThemeClassArray) SortStable(less func(a, b InputThemeClass) bool) InputThemeClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputThemeClass.
func (s InputThemeClassArray) Retain(keep func(x InputThemeClass) bool) InputThemeClassArray {
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
func (s InputThemeClassArray) First() (v InputThemeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputThemeClassArray) Last() (v InputThemeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputThemeClassArray) PopFirst() (v InputThemeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputThemeClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputThemeClassArray) Pop() (v InputThemeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputTheme returns copy with only InputTheme constructors.
func (s InputThemeClassArray) AsInputTheme() (to InputThemeArray) {
	for _, elem := range s {
		value, ok := elem.(*InputTheme)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputThemeSlug returns copy with only InputThemeSlug constructors.
func (s InputThemeClassArray) AsInputThemeSlug() (to InputThemeSlugArray) {
	for _, elem := range s {
		value, ok := elem.(*InputThemeSlug)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputThemeArray is adapter for slice of InputTheme.
type InputThemeArray []InputTheme

// Sort sorts slice of InputTheme.
func (s InputThemeArray) Sort(less func(a, b InputTheme) bool) InputThemeArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputTheme.
func (s InputThemeArray) SortStable(less func(a, b InputTheme) bool) InputThemeArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputTheme.
func (s InputThemeArray) Retain(keep func(x InputTheme) bool) InputThemeArray {
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
func (s InputThemeArray) First() (v InputTheme, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputThemeArray) Last() (v InputTheme, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputThemeArray) PopFirst() (v InputTheme, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputTheme
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputThemeArray) Pop() (v InputTheme, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputTheme by ID.
func (s InputThemeArray) SortByID() InputThemeArray {
	return s.Sort(func(a, b InputTheme) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputTheme by ID.
func (s InputThemeArray) SortStableByID() InputThemeArray {
	return s.SortStable(func(a, b InputTheme) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputThemeArray) FillMap(to map[int64]InputTheme) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputThemeArray) ToMap() map[int64]InputTheme {
	r := make(map[int64]InputTheme, len(s))
	s.FillMap(r)
	return r
}

// InputThemeSlugArray is adapter for slice of InputThemeSlug.
type InputThemeSlugArray []InputThemeSlug

// Sort sorts slice of InputThemeSlug.
func (s InputThemeSlugArray) Sort(less func(a, b InputThemeSlug) bool) InputThemeSlugArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputThemeSlug.
func (s InputThemeSlugArray) SortStable(less func(a, b InputThemeSlug) bool) InputThemeSlugArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputThemeSlug.
func (s InputThemeSlugArray) Retain(keep func(x InputThemeSlug) bool) InputThemeSlugArray {
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
func (s InputThemeSlugArray) First() (v InputThemeSlug, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputThemeSlugArray) Last() (v InputThemeSlug, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputThemeSlugArray) PopFirst() (v InputThemeSlug, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputThemeSlug
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputThemeSlugArray) Pop() (v InputThemeSlug, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
