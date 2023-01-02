// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// WebPage represents TL type `webPage#dd96962e`.
type WebPage struct {
	// Original URL of the link
	URL string
	// URL to display
	DisplayURL string
	// Type of the web page. Can be: article, photo, audio, video, document, profile, app, or
	// something else
	Type string
	// Short name of the site (e.g., Google Docs, App Store)
	SiteName string
	// Title of the content
	Title string
	// Describes a web page preview
	Description FormattedText
	// Image representing the content; may be null
	Photo Photo
	// URL to show in the embedded preview
	EmbedURL string
	// MIME type of the embedded preview, (e.g., text/html or video/mp4)
	EmbedType string
	// Width of the embedded preview
	EmbedWidth int32
	// Height of the embedded preview
	EmbedHeight int32
	// Duration of the content, in seconds
	Duration int32
	// Author of the content
	Author string
	// Preview of the content as an animation, if available; may be null
	Animation Animation
	// Preview of the content as an audio file, if available; may be null
	Audio Audio
	// Preview of the content as a document, if available; may be null
	Document Document
	// Preview of the content as a sticker for small WEBP files, if available; may be null
	Sticker Sticker
	// Preview of the content as a video, if available; may be null
	Video Video
	// Preview of the content as a video note, if available; may be null
	VideoNote VideoNote
	// Preview of the content as a voice note, if available; may be null
	VoiceNote VoiceNote
	// Version of web page instant view (currently, can be 1 or 2); 0 if none
	InstantViewVersion int32
}

// WebPageTypeID is TL type id of WebPage.
const WebPageTypeID = 0xdd96962e

// Ensuring interfaces in compile-time for WebPage.
var (
	_ bin.Encoder     = &WebPage{}
	_ bin.Decoder     = &WebPage{}
	_ bin.BareEncoder = &WebPage{}
	_ bin.BareDecoder = &WebPage{}
)

func (w *WebPage) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.URL == "") {
		return false
	}
	if !(w.DisplayURL == "") {
		return false
	}
	if !(w.Type == "") {
		return false
	}
	if !(w.SiteName == "") {
		return false
	}
	if !(w.Title == "") {
		return false
	}
	if !(w.Description.Zero()) {
		return false
	}
	if !(w.Photo.Zero()) {
		return false
	}
	if !(w.EmbedURL == "") {
		return false
	}
	if !(w.EmbedType == "") {
		return false
	}
	if !(w.EmbedWidth == 0) {
		return false
	}
	if !(w.EmbedHeight == 0) {
		return false
	}
	if !(w.Duration == 0) {
		return false
	}
	if !(w.Author == "") {
		return false
	}
	if !(w.Animation.Zero()) {
		return false
	}
	if !(w.Audio.Zero()) {
		return false
	}
	if !(w.Document.Zero()) {
		return false
	}
	if !(w.Sticker.Zero()) {
		return false
	}
	if !(w.Video.Zero()) {
		return false
	}
	if !(w.VideoNote.Zero()) {
		return false
	}
	if !(w.VoiceNote.Zero()) {
		return false
	}
	if !(w.InstantViewVersion == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebPage) String() string {
	if w == nil {
		return "WebPage(nil)"
	}
	type Alias WebPage
	return fmt.Sprintf("WebPage%+v", Alias(*w))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebPage) TypeID() uint32 {
	return WebPageTypeID
}

// TypeName returns name of type in TL schema.
func (*WebPage) TypeName() string {
	return "webPage"
}

// TypeInfo returns info about TL type.
func (w *WebPage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webPage",
		ID:   WebPageTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "DisplayURL",
			SchemaName: "display_url",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "SiteName",
			SchemaName: "site_name",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "EmbedURL",
			SchemaName: "embed_url",
		},
		{
			Name:       "EmbedType",
			SchemaName: "embed_type",
		},
		{
			Name:       "EmbedWidth",
			SchemaName: "embed_width",
		},
		{
			Name:       "EmbedHeight",
			SchemaName: "embed_height",
		},
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
		{
			Name:       "Author",
			SchemaName: "author",
		},
		{
			Name:       "Animation",
			SchemaName: "animation",
		},
		{
			Name:       "Audio",
			SchemaName: "audio",
		},
		{
			Name:       "Document",
			SchemaName: "document",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
		{
			Name:       "Video",
			SchemaName: "video",
		},
		{
			Name:       "VideoNote",
			SchemaName: "video_note",
		},
		{
			Name:       "VoiceNote",
			SchemaName: "voice_note",
		},
		{
			Name:       "InstantViewVersion",
			SchemaName: "instant_view_version",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *WebPage) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPage#dd96962e as nil")
	}
	b.PutID(WebPageTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebPage) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPage#dd96962e as nil")
	}
	b.PutString(w.URL)
	b.PutString(w.DisplayURL)
	b.PutString(w.Type)
	b.PutString(w.SiteName)
	b.PutString(w.Title)
	if err := w.Description.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field description: %w", err)
	}
	if err := w.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field photo: %w", err)
	}
	b.PutString(w.EmbedURL)
	b.PutString(w.EmbedType)
	b.PutInt32(w.EmbedWidth)
	b.PutInt32(w.EmbedHeight)
	b.PutInt32(w.Duration)
	b.PutString(w.Author)
	if err := w.Animation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field animation: %w", err)
	}
	if err := w.Audio.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field audio: %w", err)
	}
	if err := w.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field document: %w", err)
	}
	if err := w.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field sticker: %w", err)
	}
	if err := w.Video.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field video: %w", err)
	}
	if err := w.VideoNote.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field video_note: %w", err)
	}
	if err := w.VoiceNote.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field voice_note: %w", err)
	}
	b.PutInt32(w.InstantViewVersion)
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPage) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPage#dd96962e to nil")
	}
	if err := b.ConsumeID(WebPageTypeID); err != nil {
		return fmt.Errorf("unable to decode webPage#dd96962e: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebPage) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPage#dd96962e to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field url: %w", err)
		}
		w.URL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field display_url: %w", err)
		}
		w.DisplayURL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field type: %w", err)
		}
		w.Type = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field site_name: %w", err)
		}
		w.SiteName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field title: %w", err)
		}
		w.Title = value
	}
	{
		if err := w.Description.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field description: %w", err)
		}
	}
	{
		if err := w.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field photo: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field embed_url: %w", err)
		}
		w.EmbedURL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field embed_type: %w", err)
		}
		w.EmbedType = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field embed_width: %w", err)
		}
		w.EmbedWidth = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field embed_height: %w", err)
		}
		w.EmbedHeight = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field duration: %w", err)
		}
		w.Duration = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field author: %w", err)
		}
		w.Author = value
	}
	{
		if err := w.Animation.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field animation: %w", err)
		}
	}
	{
		if err := w.Audio.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field audio: %w", err)
		}
	}
	{
		if err := w.Document.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field document: %w", err)
		}
	}
	{
		if err := w.Sticker.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field sticker: %w", err)
		}
	}
	{
		if err := w.Video.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field video: %w", err)
		}
	}
	{
		if err := w.VideoNote.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field video_note: %w", err)
		}
	}
	{
		if err := w.VoiceNote.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field voice_note: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode webPage#dd96962e: field instant_view_version: %w", err)
		}
		w.InstantViewVersion = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (w *WebPage) EncodeTDLibJSON(b tdjson.Encoder) error {
	if w == nil {
		return fmt.Errorf("can't encode webPage#dd96962e as nil")
	}
	b.ObjStart()
	b.PutID("webPage")
	b.Comma()
	b.FieldStart("url")
	b.PutString(w.URL)
	b.Comma()
	b.FieldStart("display_url")
	b.PutString(w.DisplayURL)
	b.Comma()
	b.FieldStart("type")
	b.PutString(w.Type)
	b.Comma()
	b.FieldStart("site_name")
	b.PutString(w.SiteName)
	b.Comma()
	b.FieldStart("title")
	b.PutString(w.Title)
	b.Comma()
	b.FieldStart("description")
	if err := w.Description.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field description: %w", err)
	}
	b.Comma()
	b.FieldStart("photo")
	if err := w.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field photo: %w", err)
	}
	b.Comma()
	b.FieldStart("embed_url")
	b.PutString(w.EmbedURL)
	b.Comma()
	b.FieldStart("embed_type")
	b.PutString(w.EmbedType)
	b.Comma()
	b.FieldStart("embed_width")
	b.PutInt32(w.EmbedWidth)
	b.Comma()
	b.FieldStart("embed_height")
	b.PutInt32(w.EmbedHeight)
	b.Comma()
	b.FieldStart("duration")
	b.PutInt32(w.Duration)
	b.Comma()
	b.FieldStart("author")
	b.PutString(w.Author)
	b.Comma()
	b.FieldStart("animation")
	if err := w.Animation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field animation: %w", err)
	}
	b.Comma()
	b.FieldStart("audio")
	if err := w.Audio.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field audio: %w", err)
	}
	b.Comma()
	b.FieldStart("document")
	if err := w.Document.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field document: %w", err)
	}
	b.Comma()
	b.FieldStart("sticker")
	if err := w.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field sticker: %w", err)
	}
	b.Comma()
	b.FieldStart("video")
	if err := w.Video.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field video: %w", err)
	}
	b.Comma()
	b.FieldStart("video_note")
	if err := w.VideoNote.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field video_note: %w", err)
	}
	b.Comma()
	b.FieldStart("voice_note")
	if err := w.VoiceNote.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode webPage#dd96962e: field voice_note: %w", err)
	}
	b.Comma()
	b.FieldStart("instant_view_version")
	b.PutInt32(w.InstantViewVersion)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (w *WebPage) DecodeTDLibJSON(b tdjson.Decoder) error {
	if w == nil {
		return fmt.Errorf("can't decode webPage#dd96962e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("webPage"); err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field url: %w", err)
			}
			w.URL = value
		case "display_url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field display_url: %w", err)
			}
			w.DisplayURL = value
		case "type":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field type: %w", err)
			}
			w.Type = value
		case "site_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field site_name: %w", err)
			}
			w.SiteName = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field title: %w", err)
			}
			w.Title = value
		case "description":
			if err := w.Description.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field description: %w", err)
			}
		case "photo":
			if err := w.Photo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field photo: %w", err)
			}
		case "embed_url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field embed_url: %w", err)
			}
			w.EmbedURL = value
		case "embed_type":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field embed_type: %w", err)
			}
			w.EmbedType = value
		case "embed_width":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field embed_width: %w", err)
			}
			w.EmbedWidth = value
		case "embed_height":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field embed_height: %w", err)
			}
			w.EmbedHeight = value
		case "duration":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field duration: %w", err)
			}
			w.Duration = value
		case "author":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field author: %w", err)
			}
			w.Author = value
		case "animation":
			if err := w.Animation.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field animation: %w", err)
			}
		case "audio":
			if err := w.Audio.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field audio: %w", err)
			}
		case "document":
			if err := w.Document.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field document: %w", err)
			}
		case "sticker":
			if err := w.Sticker.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field sticker: %w", err)
			}
		case "video":
			if err := w.Video.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field video: %w", err)
			}
		case "video_note":
			if err := w.VideoNote.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field video_note: %w", err)
			}
		case "voice_note":
			if err := w.VoiceNote.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field voice_note: %w", err)
			}
		case "instant_view_version":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode webPage#dd96962e: field instant_view_version: %w", err)
			}
			w.InstantViewVersion = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (w *WebPage) GetURL() (value string) {
	if w == nil {
		return
	}
	return w.URL
}

// GetDisplayURL returns value of DisplayURL field.
func (w *WebPage) GetDisplayURL() (value string) {
	if w == nil {
		return
	}
	return w.DisplayURL
}

// GetType returns value of Type field.
func (w *WebPage) GetType() (value string) {
	if w == nil {
		return
	}
	return w.Type
}

// GetSiteName returns value of SiteName field.
func (w *WebPage) GetSiteName() (value string) {
	if w == nil {
		return
	}
	return w.SiteName
}

// GetTitle returns value of Title field.
func (w *WebPage) GetTitle() (value string) {
	if w == nil {
		return
	}
	return w.Title
}

// GetDescription returns value of Description field.
func (w *WebPage) GetDescription() (value FormattedText) {
	if w == nil {
		return
	}
	return w.Description
}

// GetPhoto returns value of Photo field.
func (w *WebPage) GetPhoto() (value Photo) {
	if w == nil {
		return
	}
	return w.Photo
}

// GetEmbedURL returns value of EmbedURL field.
func (w *WebPage) GetEmbedURL() (value string) {
	if w == nil {
		return
	}
	return w.EmbedURL
}

// GetEmbedType returns value of EmbedType field.
func (w *WebPage) GetEmbedType() (value string) {
	if w == nil {
		return
	}
	return w.EmbedType
}

// GetEmbedWidth returns value of EmbedWidth field.
func (w *WebPage) GetEmbedWidth() (value int32) {
	if w == nil {
		return
	}
	return w.EmbedWidth
}

// GetEmbedHeight returns value of EmbedHeight field.
func (w *WebPage) GetEmbedHeight() (value int32) {
	if w == nil {
		return
	}
	return w.EmbedHeight
}

// GetDuration returns value of Duration field.
func (w *WebPage) GetDuration() (value int32) {
	if w == nil {
		return
	}
	return w.Duration
}

// GetAuthor returns value of Author field.
func (w *WebPage) GetAuthor() (value string) {
	if w == nil {
		return
	}
	return w.Author
}

// GetAnimation returns value of Animation field.
func (w *WebPage) GetAnimation() (value Animation) {
	if w == nil {
		return
	}
	return w.Animation
}

// GetAudio returns value of Audio field.
func (w *WebPage) GetAudio() (value Audio) {
	if w == nil {
		return
	}
	return w.Audio
}

// GetDocument returns value of Document field.
func (w *WebPage) GetDocument() (value Document) {
	if w == nil {
		return
	}
	return w.Document
}

// GetSticker returns value of Sticker field.
func (w *WebPage) GetSticker() (value Sticker) {
	if w == nil {
		return
	}
	return w.Sticker
}

// GetVideo returns value of Video field.
func (w *WebPage) GetVideo() (value Video) {
	if w == nil {
		return
	}
	return w.Video
}

// GetVideoNote returns value of VideoNote field.
func (w *WebPage) GetVideoNote() (value VideoNote) {
	if w == nil {
		return
	}
	return w.VideoNote
}

// GetVoiceNote returns value of VoiceNote field.
func (w *WebPage) GetVoiceNote() (value VoiceNote) {
	if w == nil {
		return
	}
	return w.VoiceNote
}

// GetInstantViewVersion returns value of InstantViewVersion field.
func (w *WebPage) GetInstantViewVersion() (value int32) {
	if w == nil {
		return
	}
	return w.InstantViewVersion
}
