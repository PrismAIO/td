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

// InputMediaClassArray is adapter for slice of InputMediaClass.
type InputMediaClassArray []InputMediaClass

// Sort sorts slice of InputMediaClass.
func (s InputMediaClassArray) Sort(less func(a, b InputMediaClass) bool) InputMediaClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaClass.
func (s InputMediaClassArray) SortStable(less func(a, b InputMediaClass) bool) InputMediaClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaClass.
func (s InputMediaClassArray) Retain(keep func(x InputMediaClass) bool) InputMediaClassArray {
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
func (s InputMediaClassArray) First() (v InputMediaClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaClassArray) Last() (v InputMediaClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaClassArray) PopFirst() (v InputMediaClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaClassArray) Pop() (v InputMediaClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputMediaUploadedPhoto returns copy with only InputMediaUploadedPhoto constructors.
func (s InputMediaClassArray) AsInputMediaUploadedPhoto() (to InputMediaUploadedPhotoArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaUploadedPhoto)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaPhoto returns copy with only InputMediaPhoto constructors.
func (s InputMediaClassArray) AsInputMediaPhoto() (to InputMediaPhotoArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaPhoto)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaGeoPoint returns copy with only InputMediaGeoPoint constructors.
func (s InputMediaClassArray) AsInputMediaGeoPoint() (to InputMediaGeoPointArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaGeoPoint)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaContact returns copy with only InputMediaContact constructors.
func (s InputMediaClassArray) AsInputMediaContact() (to InputMediaContactArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaContact)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaUploadedDocument returns copy with only InputMediaUploadedDocument constructors.
func (s InputMediaClassArray) AsInputMediaUploadedDocument() (to InputMediaUploadedDocumentArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaUploadedDocument)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaDocument returns copy with only InputMediaDocument constructors.
func (s InputMediaClassArray) AsInputMediaDocument() (to InputMediaDocumentArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaDocument)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaVenue returns copy with only InputMediaVenue constructors.
func (s InputMediaClassArray) AsInputMediaVenue() (to InputMediaVenueArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaVenue)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaPhotoExternal returns copy with only InputMediaPhotoExternal constructors.
func (s InputMediaClassArray) AsInputMediaPhotoExternal() (to InputMediaPhotoExternalArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaPhotoExternal)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaDocumentExternal returns copy with only InputMediaDocumentExternal constructors.
func (s InputMediaClassArray) AsInputMediaDocumentExternal() (to InputMediaDocumentExternalArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaDocumentExternal)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaGame returns copy with only InputMediaGame constructors.
func (s InputMediaClassArray) AsInputMediaGame() (to InputMediaGameArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaGame)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaInvoice returns copy with only InputMediaInvoice constructors.
func (s InputMediaClassArray) AsInputMediaInvoice() (to InputMediaInvoiceArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaInvoice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaGeoLive returns copy with only InputMediaGeoLive constructors.
func (s InputMediaClassArray) AsInputMediaGeoLive() (to InputMediaGeoLiveArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaGeoLive)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaPoll returns copy with only InputMediaPoll constructors.
func (s InputMediaClassArray) AsInputMediaPoll() (to InputMediaPollArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaPoll)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMediaDice returns copy with only InputMediaDice constructors.
func (s InputMediaClassArray) AsInputMediaDice() (to InputMediaDiceArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMediaDice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputMediaUploadedPhotoArray is adapter for slice of InputMediaUploadedPhoto.
type InputMediaUploadedPhotoArray []InputMediaUploadedPhoto

// Sort sorts slice of InputMediaUploadedPhoto.
func (s InputMediaUploadedPhotoArray) Sort(less func(a, b InputMediaUploadedPhoto) bool) InputMediaUploadedPhotoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaUploadedPhoto.
func (s InputMediaUploadedPhotoArray) SortStable(less func(a, b InputMediaUploadedPhoto) bool) InputMediaUploadedPhotoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaUploadedPhoto.
func (s InputMediaUploadedPhotoArray) Retain(keep func(x InputMediaUploadedPhoto) bool) InputMediaUploadedPhotoArray {
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
func (s InputMediaUploadedPhotoArray) First() (v InputMediaUploadedPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaUploadedPhotoArray) Last() (v InputMediaUploadedPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaUploadedPhotoArray) PopFirst() (v InputMediaUploadedPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaUploadedPhoto
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaUploadedPhotoArray) Pop() (v InputMediaUploadedPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaPhotoArray is adapter for slice of InputMediaPhoto.
type InputMediaPhotoArray []InputMediaPhoto

// Sort sorts slice of InputMediaPhoto.
func (s InputMediaPhotoArray) Sort(less func(a, b InputMediaPhoto) bool) InputMediaPhotoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaPhoto.
func (s InputMediaPhotoArray) SortStable(less func(a, b InputMediaPhoto) bool) InputMediaPhotoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaPhoto.
func (s InputMediaPhotoArray) Retain(keep func(x InputMediaPhoto) bool) InputMediaPhotoArray {
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
func (s InputMediaPhotoArray) First() (v InputMediaPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaPhotoArray) Last() (v InputMediaPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaPhotoArray) PopFirst() (v InputMediaPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaPhoto
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaPhotoArray) Pop() (v InputMediaPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaGeoPointArray is adapter for slice of InputMediaGeoPoint.
type InputMediaGeoPointArray []InputMediaGeoPoint

// Sort sorts slice of InputMediaGeoPoint.
func (s InputMediaGeoPointArray) Sort(less func(a, b InputMediaGeoPoint) bool) InputMediaGeoPointArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaGeoPoint.
func (s InputMediaGeoPointArray) SortStable(less func(a, b InputMediaGeoPoint) bool) InputMediaGeoPointArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaGeoPoint.
func (s InputMediaGeoPointArray) Retain(keep func(x InputMediaGeoPoint) bool) InputMediaGeoPointArray {
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
func (s InputMediaGeoPointArray) First() (v InputMediaGeoPoint, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaGeoPointArray) Last() (v InputMediaGeoPoint, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaGeoPointArray) PopFirst() (v InputMediaGeoPoint, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaGeoPoint
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaGeoPointArray) Pop() (v InputMediaGeoPoint, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaContactArray is adapter for slice of InputMediaContact.
type InputMediaContactArray []InputMediaContact

// Sort sorts slice of InputMediaContact.
func (s InputMediaContactArray) Sort(less func(a, b InputMediaContact) bool) InputMediaContactArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaContact.
func (s InputMediaContactArray) SortStable(less func(a, b InputMediaContact) bool) InputMediaContactArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaContact.
func (s InputMediaContactArray) Retain(keep func(x InputMediaContact) bool) InputMediaContactArray {
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
func (s InputMediaContactArray) First() (v InputMediaContact, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaContactArray) Last() (v InputMediaContact, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaContactArray) PopFirst() (v InputMediaContact, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaContact
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaContactArray) Pop() (v InputMediaContact, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaUploadedDocumentArray is adapter for slice of InputMediaUploadedDocument.
type InputMediaUploadedDocumentArray []InputMediaUploadedDocument

// Sort sorts slice of InputMediaUploadedDocument.
func (s InputMediaUploadedDocumentArray) Sort(less func(a, b InputMediaUploadedDocument) bool) InputMediaUploadedDocumentArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaUploadedDocument.
func (s InputMediaUploadedDocumentArray) SortStable(less func(a, b InputMediaUploadedDocument) bool) InputMediaUploadedDocumentArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaUploadedDocument.
func (s InputMediaUploadedDocumentArray) Retain(keep func(x InputMediaUploadedDocument) bool) InputMediaUploadedDocumentArray {
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
func (s InputMediaUploadedDocumentArray) First() (v InputMediaUploadedDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaUploadedDocumentArray) Last() (v InputMediaUploadedDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaUploadedDocumentArray) PopFirst() (v InputMediaUploadedDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaUploadedDocument
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaUploadedDocumentArray) Pop() (v InputMediaUploadedDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaDocumentArray is adapter for slice of InputMediaDocument.
type InputMediaDocumentArray []InputMediaDocument

// Sort sorts slice of InputMediaDocument.
func (s InputMediaDocumentArray) Sort(less func(a, b InputMediaDocument) bool) InputMediaDocumentArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaDocument.
func (s InputMediaDocumentArray) SortStable(less func(a, b InputMediaDocument) bool) InputMediaDocumentArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaDocument.
func (s InputMediaDocumentArray) Retain(keep func(x InputMediaDocument) bool) InputMediaDocumentArray {
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
func (s InputMediaDocumentArray) First() (v InputMediaDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaDocumentArray) Last() (v InputMediaDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaDocumentArray) PopFirst() (v InputMediaDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaDocument
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaDocumentArray) Pop() (v InputMediaDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaVenueArray is adapter for slice of InputMediaVenue.
type InputMediaVenueArray []InputMediaVenue

// Sort sorts slice of InputMediaVenue.
func (s InputMediaVenueArray) Sort(less func(a, b InputMediaVenue) bool) InputMediaVenueArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaVenue.
func (s InputMediaVenueArray) SortStable(less func(a, b InputMediaVenue) bool) InputMediaVenueArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaVenue.
func (s InputMediaVenueArray) Retain(keep func(x InputMediaVenue) bool) InputMediaVenueArray {
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
func (s InputMediaVenueArray) First() (v InputMediaVenue, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaVenueArray) Last() (v InputMediaVenue, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaVenueArray) PopFirst() (v InputMediaVenue, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaVenue
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaVenueArray) Pop() (v InputMediaVenue, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaPhotoExternalArray is adapter for slice of InputMediaPhotoExternal.
type InputMediaPhotoExternalArray []InputMediaPhotoExternal

// Sort sorts slice of InputMediaPhotoExternal.
func (s InputMediaPhotoExternalArray) Sort(less func(a, b InputMediaPhotoExternal) bool) InputMediaPhotoExternalArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaPhotoExternal.
func (s InputMediaPhotoExternalArray) SortStable(less func(a, b InputMediaPhotoExternal) bool) InputMediaPhotoExternalArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaPhotoExternal.
func (s InputMediaPhotoExternalArray) Retain(keep func(x InputMediaPhotoExternal) bool) InputMediaPhotoExternalArray {
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
func (s InputMediaPhotoExternalArray) First() (v InputMediaPhotoExternal, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaPhotoExternalArray) Last() (v InputMediaPhotoExternal, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaPhotoExternalArray) PopFirst() (v InputMediaPhotoExternal, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaPhotoExternal
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaPhotoExternalArray) Pop() (v InputMediaPhotoExternal, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaDocumentExternalArray is adapter for slice of InputMediaDocumentExternal.
type InputMediaDocumentExternalArray []InputMediaDocumentExternal

// Sort sorts slice of InputMediaDocumentExternal.
func (s InputMediaDocumentExternalArray) Sort(less func(a, b InputMediaDocumentExternal) bool) InputMediaDocumentExternalArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaDocumentExternal.
func (s InputMediaDocumentExternalArray) SortStable(less func(a, b InputMediaDocumentExternal) bool) InputMediaDocumentExternalArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaDocumentExternal.
func (s InputMediaDocumentExternalArray) Retain(keep func(x InputMediaDocumentExternal) bool) InputMediaDocumentExternalArray {
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
func (s InputMediaDocumentExternalArray) First() (v InputMediaDocumentExternal, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaDocumentExternalArray) Last() (v InputMediaDocumentExternal, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaDocumentExternalArray) PopFirst() (v InputMediaDocumentExternal, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaDocumentExternal
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaDocumentExternalArray) Pop() (v InputMediaDocumentExternal, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaGameArray is adapter for slice of InputMediaGame.
type InputMediaGameArray []InputMediaGame

// Sort sorts slice of InputMediaGame.
func (s InputMediaGameArray) Sort(less func(a, b InputMediaGame) bool) InputMediaGameArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaGame.
func (s InputMediaGameArray) SortStable(less func(a, b InputMediaGame) bool) InputMediaGameArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaGame.
func (s InputMediaGameArray) Retain(keep func(x InputMediaGame) bool) InputMediaGameArray {
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
func (s InputMediaGameArray) First() (v InputMediaGame, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaGameArray) Last() (v InputMediaGame, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaGameArray) PopFirst() (v InputMediaGame, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaGame
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaGameArray) Pop() (v InputMediaGame, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaInvoiceArray is adapter for slice of InputMediaInvoice.
type InputMediaInvoiceArray []InputMediaInvoice

// Sort sorts slice of InputMediaInvoice.
func (s InputMediaInvoiceArray) Sort(less func(a, b InputMediaInvoice) bool) InputMediaInvoiceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaInvoice.
func (s InputMediaInvoiceArray) SortStable(less func(a, b InputMediaInvoice) bool) InputMediaInvoiceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaInvoice.
func (s InputMediaInvoiceArray) Retain(keep func(x InputMediaInvoice) bool) InputMediaInvoiceArray {
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
func (s InputMediaInvoiceArray) First() (v InputMediaInvoice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaInvoiceArray) Last() (v InputMediaInvoice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaInvoiceArray) PopFirst() (v InputMediaInvoice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaInvoice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaInvoiceArray) Pop() (v InputMediaInvoice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaGeoLiveArray is adapter for slice of InputMediaGeoLive.
type InputMediaGeoLiveArray []InputMediaGeoLive

// Sort sorts slice of InputMediaGeoLive.
func (s InputMediaGeoLiveArray) Sort(less func(a, b InputMediaGeoLive) bool) InputMediaGeoLiveArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaGeoLive.
func (s InputMediaGeoLiveArray) SortStable(less func(a, b InputMediaGeoLive) bool) InputMediaGeoLiveArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaGeoLive.
func (s InputMediaGeoLiveArray) Retain(keep func(x InputMediaGeoLive) bool) InputMediaGeoLiveArray {
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
func (s InputMediaGeoLiveArray) First() (v InputMediaGeoLive, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaGeoLiveArray) Last() (v InputMediaGeoLive, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaGeoLiveArray) PopFirst() (v InputMediaGeoLive, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaGeoLive
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaGeoLiveArray) Pop() (v InputMediaGeoLive, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaPollArray is adapter for slice of InputMediaPoll.
type InputMediaPollArray []InputMediaPoll

// Sort sorts slice of InputMediaPoll.
func (s InputMediaPollArray) Sort(less func(a, b InputMediaPoll) bool) InputMediaPollArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaPoll.
func (s InputMediaPollArray) SortStable(less func(a, b InputMediaPoll) bool) InputMediaPollArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaPoll.
func (s InputMediaPollArray) Retain(keep func(x InputMediaPoll) bool) InputMediaPollArray {
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
func (s InputMediaPollArray) First() (v InputMediaPoll, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaPollArray) Last() (v InputMediaPoll, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaPollArray) PopFirst() (v InputMediaPoll, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaPoll
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaPollArray) Pop() (v InputMediaPoll, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputMediaDiceArray is adapter for slice of InputMediaDice.
type InputMediaDiceArray []InputMediaDice

// Sort sorts slice of InputMediaDice.
func (s InputMediaDiceArray) Sort(less func(a, b InputMediaDice) bool) InputMediaDiceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMediaDice.
func (s InputMediaDiceArray) SortStable(less func(a, b InputMediaDice) bool) InputMediaDiceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMediaDice.
func (s InputMediaDiceArray) Retain(keep func(x InputMediaDice) bool) InputMediaDiceArray {
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
func (s InputMediaDiceArray) First() (v InputMediaDice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMediaDiceArray) Last() (v InputMediaDice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMediaDiceArray) PopFirst() (v InputMediaDice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMediaDice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMediaDiceArray) Pop() (v InputMediaDice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
