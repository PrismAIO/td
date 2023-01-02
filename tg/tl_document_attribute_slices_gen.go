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

// DocumentAttributeClassArray is adapter for slice of DocumentAttributeClass.
type DocumentAttributeClassArray []DocumentAttributeClass

// Sort sorts slice of DocumentAttributeClass.
func (s DocumentAttributeClassArray) Sort(less func(a, b DocumentAttributeClass) bool) DocumentAttributeClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DocumentAttributeClass.
func (s DocumentAttributeClassArray) SortStable(less func(a, b DocumentAttributeClass) bool) DocumentAttributeClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DocumentAttributeClass.
func (s DocumentAttributeClassArray) Retain(keep func(x DocumentAttributeClass) bool) DocumentAttributeClassArray {
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
func (s DocumentAttributeClassArray) First() (v DocumentAttributeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentAttributeClassArray) Last() (v DocumentAttributeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentAttributeClassArray) PopFirst() (v DocumentAttributeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DocumentAttributeClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DocumentAttributeClassArray) Pop() (v DocumentAttributeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsDocumentAttributeImageSize returns copy with only DocumentAttributeImageSize constructors.
func (s DocumentAttributeClassArray) AsDocumentAttributeImageSize() (to DocumentAttributeImageSizeArray) {
	for _, elem := range s {
		value, ok := elem.(*DocumentAttributeImageSize)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsDocumentAttributeSticker returns copy with only DocumentAttributeSticker constructors.
func (s DocumentAttributeClassArray) AsDocumentAttributeSticker() (to DocumentAttributeStickerArray) {
	for _, elem := range s {
		value, ok := elem.(*DocumentAttributeSticker)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsDocumentAttributeVideo returns copy with only DocumentAttributeVideo constructors.
func (s DocumentAttributeClassArray) AsDocumentAttributeVideo() (to DocumentAttributeVideoArray) {
	for _, elem := range s {
		value, ok := elem.(*DocumentAttributeVideo)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsDocumentAttributeAudio returns copy with only DocumentAttributeAudio constructors.
func (s DocumentAttributeClassArray) AsDocumentAttributeAudio() (to DocumentAttributeAudioArray) {
	for _, elem := range s {
		value, ok := elem.(*DocumentAttributeAudio)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsDocumentAttributeFilename returns copy with only DocumentAttributeFilename constructors.
func (s DocumentAttributeClassArray) AsDocumentAttributeFilename() (to DocumentAttributeFilenameArray) {
	for _, elem := range s {
		value, ok := elem.(*DocumentAttributeFilename)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsDocumentAttributeCustomEmoji returns copy with only DocumentAttributeCustomEmoji constructors.
func (s DocumentAttributeClassArray) AsDocumentAttributeCustomEmoji() (to DocumentAttributeCustomEmojiArray) {
	for _, elem := range s {
		value, ok := elem.(*DocumentAttributeCustomEmoji)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// DocumentAttributeImageSizeArray is adapter for slice of DocumentAttributeImageSize.
type DocumentAttributeImageSizeArray []DocumentAttributeImageSize

// Sort sorts slice of DocumentAttributeImageSize.
func (s DocumentAttributeImageSizeArray) Sort(less func(a, b DocumentAttributeImageSize) bool) DocumentAttributeImageSizeArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DocumentAttributeImageSize.
func (s DocumentAttributeImageSizeArray) SortStable(less func(a, b DocumentAttributeImageSize) bool) DocumentAttributeImageSizeArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DocumentAttributeImageSize.
func (s DocumentAttributeImageSizeArray) Retain(keep func(x DocumentAttributeImageSize) bool) DocumentAttributeImageSizeArray {
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
func (s DocumentAttributeImageSizeArray) First() (v DocumentAttributeImageSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentAttributeImageSizeArray) Last() (v DocumentAttributeImageSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentAttributeImageSizeArray) PopFirst() (v DocumentAttributeImageSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DocumentAttributeImageSize
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DocumentAttributeImageSizeArray) Pop() (v DocumentAttributeImageSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// DocumentAttributeStickerArray is adapter for slice of DocumentAttributeSticker.
type DocumentAttributeStickerArray []DocumentAttributeSticker

// Sort sorts slice of DocumentAttributeSticker.
func (s DocumentAttributeStickerArray) Sort(less func(a, b DocumentAttributeSticker) bool) DocumentAttributeStickerArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DocumentAttributeSticker.
func (s DocumentAttributeStickerArray) SortStable(less func(a, b DocumentAttributeSticker) bool) DocumentAttributeStickerArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DocumentAttributeSticker.
func (s DocumentAttributeStickerArray) Retain(keep func(x DocumentAttributeSticker) bool) DocumentAttributeStickerArray {
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
func (s DocumentAttributeStickerArray) First() (v DocumentAttributeSticker, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentAttributeStickerArray) Last() (v DocumentAttributeSticker, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentAttributeStickerArray) PopFirst() (v DocumentAttributeSticker, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DocumentAttributeSticker
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DocumentAttributeStickerArray) Pop() (v DocumentAttributeSticker, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// DocumentAttributeVideoArray is adapter for slice of DocumentAttributeVideo.
type DocumentAttributeVideoArray []DocumentAttributeVideo

// Sort sorts slice of DocumentAttributeVideo.
func (s DocumentAttributeVideoArray) Sort(less func(a, b DocumentAttributeVideo) bool) DocumentAttributeVideoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DocumentAttributeVideo.
func (s DocumentAttributeVideoArray) SortStable(less func(a, b DocumentAttributeVideo) bool) DocumentAttributeVideoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DocumentAttributeVideo.
func (s DocumentAttributeVideoArray) Retain(keep func(x DocumentAttributeVideo) bool) DocumentAttributeVideoArray {
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
func (s DocumentAttributeVideoArray) First() (v DocumentAttributeVideo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentAttributeVideoArray) Last() (v DocumentAttributeVideo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentAttributeVideoArray) PopFirst() (v DocumentAttributeVideo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DocumentAttributeVideo
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DocumentAttributeVideoArray) Pop() (v DocumentAttributeVideo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// DocumentAttributeAudioArray is adapter for slice of DocumentAttributeAudio.
type DocumentAttributeAudioArray []DocumentAttributeAudio

// Sort sorts slice of DocumentAttributeAudio.
func (s DocumentAttributeAudioArray) Sort(less func(a, b DocumentAttributeAudio) bool) DocumentAttributeAudioArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DocumentAttributeAudio.
func (s DocumentAttributeAudioArray) SortStable(less func(a, b DocumentAttributeAudio) bool) DocumentAttributeAudioArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DocumentAttributeAudio.
func (s DocumentAttributeAudioArray) Retain(keep func(x DocumentAttributeAudio) bool) DocumentAttributeAudioArray {
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
func (s DocumentAttributeAudioArray) First() (v DocumentAttributeAudio, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentAttributeAudioArray) Last() (v DocumentAttributeAudio, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentAttributeAudioArray) PopFirst() (v DocumentAttributeAudio, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DocumentAttributeAudio
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DocumentAttributeAudioArray) Pop() (v DocumentAttributeAudio, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// DocumentAttributeFilenameArray is adapter for slice of DocumentAttributeFilename.
type DocumentAttributeFilenameArray []DocumentAttributeFilename

// Sort sorts slice of DocumentAttributeFilename.
func (s DocumentAttributeFilenameArray) Sort(less func(a, b DocumentAttributeFilename) bool) DocumentAttributeFilenameArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DocumentAttributeFilename.
func (s DocumentAttributeFilenameArray) SortStable(less func(a, b DocumentAttributeFilename) bool) DocumentAttributeFilenameArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DocumentAttributeFilename.
func (s DocumentAttributeFilenameArray) Retain(keep func(x DocumentAttributeFilename) bool) DocumentAttributeFilenameArray {
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
func (s DocumentAttributeFilenameArray) First() (v DocumentAttributeFilename, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentAttributeFilenameArray) Last() (v DocumentAttributeFilename, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentAttributeFilenameArray) PopFirst() (v DocumentAttributeFilename, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DocumentAttributeFilename
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DocumentAttributeFilenameArray) Pop() (v DocumentAttributeFilename, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// DocumentAttributeCustomEmojiArray is adapter for slice of DocumentAttributeCustomEmoji.
type DocumentAttributeCustomEmojiArray []DocumentAttributeCustomEmoji

// Sort sorts slice of DocumentAttributeCustomEmoji.
func (s DocumentAttributeCustomEmojiArray) Sort(less func(a, b DocumentAttributeCustomEmoji) bool) DocumentAttributeCustomEmojiArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DocumentAttributeCustomEmoji.
func (s DocumentAttributeCustomEmojiArray) SortStable(less func(a, b DocumentAttributeCustomEmoji) bool) DocumentAttributeCustomEmojiArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DocumentAttributeCustomEmoji.
func (s DocumentAttributeCustomEmojiArray) Retain(keep func(x DocumentAttributeCustomEmoji) bool) DocumentAttributeCustomEmojiArray {
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
func (s DocumentAttributeCustomEmojiArray) First() (v DocumentAttributeCustomEmoji, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DocumentAttributeCustomEmojiArray) Last() (v DocumentAttributeCustomEmoji, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DocumentAttributeCustomEmojiArray) PopFirst() (v DocumentAttributeCustomEmoji, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DocumentAttributeCustomEmoji
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DocumentAttributeCustomEmojiArray) Pop() (v DocumentAttributeCustomEmoji, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
