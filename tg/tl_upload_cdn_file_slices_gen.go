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

// UploadCDNFileClassArray is adapter for slice of UploadCDNFileClass.
type UploadCDNFileClassArray []UploadCDNFileClass

// Sort sorts slice of UploadCDNFileClass.
func (s UploadCDNFileClassArray) Sort(less func(a, b UploadCDNFileClass) bool) UploadCDNFileClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UploadCDNFileClass.
func (s UploadCDNFileClassArray) SortStable(less func(a, b UploadCDNFileClass) bool) UploadCDNFileClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UploadCDNFileClass.
func (s UploadCDNFileClassArray) Retain(keep func(x UploadCDNFileClass) bool) UploadCDNFileClassArray {
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
func (s UploadCDNFileClassArray) First() (v UploadCDNFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UploadCDNFileClassArray) Last() (v UploadCDNFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UploadCDNFileClassArray) PopFirst() (v UploadCDNFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UploadCDNFileClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UploadCDNFileClassArray) Pop() (v UploadCDNFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsUploadCDNFileReuploadNeeded returns copy with only UploadCDNFileReuploadNeeded constructors.
func (s UploadCDNFileClassArray) AsUploadCDNFileReuploadNeeded() (to UploadCDNFileReuploadNeededArray) {
	for _, elem := range s {
		value, ok := elem.(*UploadCDNFileReuploadNeeded)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUploadCDNFile returns copy with only UploadCDNFile constructors.
func (s UploadCDNFileClassArray) AsUploadCDNFile() (to UploadCDNFileArray) {
	for _, elem := range s {
		value, ok := elem.(*UploadCDNFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// UploadCDNFileReuploadNeededArray is adapter for slice of UploadCDNFileReuploadNeeded.
type UploadCDNFileReuploadNeededArray []UploadCDNFileReuploadNeeded

// Sort sorts slice of UploadCDNFileReuploadNeeded.
func (s UploadCDNFileReuploadNeededArray) Sort(less func(a, b UploadCDNFileReuploadNeeded) bool) UploadCDNFileReuploadNeededArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UploadCDNFileReuploadNeeded.
func (s UploadCDNFileReuploadNeededArray) SortStable(less func(a, b UploadCDNFileReuploadNeeded) bool) UploadCDNFileReuploadNeededArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UploadCDNFileReuploadNeeded.
func (s UploadCDNFileReuploadNeededArray) Retain(keep func(x UploadCDNFileReuploadNeeded) bool) UploadCDNFileReuploadNeededArray {
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
func (s UploadCDNFileReuploadNeededArray) First() (v UploadCDNFileReuploadNeeded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UploadCDNFileReuploadNeededArray) Last() (v UploadCDNFileReuploadNeeded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UploadCDNFileReuploadNeededArray) PopFirst() (v UploadCDNFileReuploadNeeded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UploadCDNFileReuploadNeeded
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UploadCDNFileReuploadNeededArray) Pop() (v UploadCDNFileReuploadNeeded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UploadCDNFileArray is adapter for slice of UploadCDNFile.
type UploadCDNFileArray []UploadCDNFile

// Sort sorts slice of UploadCDNFile.
func (s UploadCDNFileArray) Sort(less func(a, b UploadCDNFile) bool) UploadCDNFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UploadCDNFile.
func (s UploadCDNFileArray) SortStable(less func(a, b UploadCDNFile) bool) UploadCDNFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UploadCDNFile.
func (s UploadCDNFileArray) Retain(keep func(x UploadCDNFile) bool) UploadCDNFileArray {
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
func (s UploadCDNFileArray) First() (v UploadCDNFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UploadCDNFileArray) Last() (v UploadCDNFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UploadCDNFileArray) PopFirst() (v UploadCDNFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UploadCDNFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UploadCDNFileArray) Pop() (v UploadCDNFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
