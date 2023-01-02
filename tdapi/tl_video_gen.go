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

// Video represents TL type `video#31a460cc`.
type Video struct {
	// Duration of the video, in seconds; as defined by the sender
	Duration int32
	// Video width; as defined by the sender
	Width int32
	// Video height; as defined by the sender
	Height int32
	// Original name of the file; as defined by the sender
	FileName string
	// MIME type of the file; as defined by the sender
	MimeType string
	// True, if stickers were added to the video. The list of corresponding sticker sets can
	// be received using getAttachedStickerSets
	HasStickers bool
	// True, if the video is supposed to be streamed
	SupportsStreaming bool
	// Video minithumbnail; may be null
	Minithumbnail Minithumbnail
	// Video thumbnail in JPEG or MPEG4 format; as defined by the sender; may be null
	Thumbnail Thumbnail
	// File containing the video
	Video File
}

// VideoTypeID is TL type id of Video.
const VideoTypeID = 0x31a460cc

// Ensuring interfaces in compile-time for Video.
var (
	_ bin.Encoder     = &Video{}
	_ bin.Decoder     = &Video{}
	_ bin.BareEncoder = &Video{}
	_ bin.BareDecoder = &Video{}
)

func (v *Video) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Duration == 0) {
		return false
	}
	if !(v.Width == 0) {
		return false
	}
	if !(v.Height == 0) {
		return false
	}
	if !(v.FileName == "") {
		return false
	}
	if !(v.MimeType == "") {
		return false
	}
	if !(v.HasStickers == false) {
		return false
	}
	if !(v.SupportsStreaming == false) {
		return false
	}
	if !(v.Minithumbnail.Zero()) {
		return false
	}
	if !(v.Thumbnail.Zero()) {
		return false
	}
	if !(v.Video.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *Video) String() string {
	if v == nil {
		return "Video(nil)"
	}
	type Alias Video
	return fmt.Sprintf("Video%+v", Alias(*v))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Video) TypeID() uint32 {
	return VideoTypeID
}

// TypeName returns name of type in TL schema.
func (*Video) TypeName() string {
	return "video"
}

// TypeInfo returns info about TL type.
func (v *Video) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "video",
		ID:   VideoTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
		{
			Name:       "Width",
			SchemaName: "width",
		},
		{
			Name:       "Height",
			SchemaName: "height",
		},
		{
			Name:       "FileName",
			SchemaName: "file_name",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
		{
			Name:       "HasStickers",
			SchemaName: "has_stickers",
		},
		{
			Name:       "SupportsStreaming",
			SchemaName: "supports_streaming",
		},
		{
			Name:       "Minithumbnail",
			SchemaName: "minithumbnail",
		},
		{
			Name:       "Thumbnail",
			SchemaName: "thumbnail",
		},
		{
			Name:       "Video",
			SchemaName: "video",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *Video) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode video#31a460cc as nil")
	}
	b.PutID(VideoTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *Video) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode video#31a460cc as nil")
	}
	b.PutInt32(v.Duration)
	b.PutInt32(v.Width)
	b.PutInt32(v.Height)
	b.PutString(v.FileName)
	b.PutString(v.MimeType)
	b.PutBool(v.HasStickers)
	b.PutBool(v.SupportsStreaming)
	if err := v.Minithumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode video#31a460cc: field minithumbnail: %w", err)
	}
	if err := v.Thumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode video#31a460cc: field thumbnail: %w", err)
	}
	if err := v.Video.Encode(b); err != nil {
		return fmt.Errorf("unable to encode video#31a460cc: field video: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *Video) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode video#31a460cc to nil")
	}
	if err := b.ConsumeID(VideoTypeID); err != nil {
		return fmt.Errorf("unable to decode video#31a460cc: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *Video) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode video#31a460cc to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode video#31a460cc: field duration: %w", err)
		}
		v.Duration = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode video#31a460cc: field width: %w", err)
		}
		v.Width = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode video#31a460cc: field height: %w", err)
		}
		v.Height = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode video#31a460cc: field file_name: %w", err)
		}
		v.FileName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode video#31a460cc: field mime_type: %w", err)
		}
		v.MimeType = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode video#31a460cc: field has_stickers: %w", err)
		}
		v.HasStickers = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode video#31a460cc: field supports_streaming: %w", err)
		}
		v.SupportsStreaming = value
	}
	{
		if err := v.Minithumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode video#31a460cc: field minithumbnail: %w", err)
		}
	}
	{
		if err := v.Thumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode video#31a460cc: field thumbnail: %w", err)
		}
	}
	{
		if err := v.Video.Decode(b); err != nil {
			return fmt.Errorf("unable to decode video#31a460cc: field video: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (v *Video) EncodeTDLibJSON(b tdjson.Encoder) error {
	if v == nil {
		return fmt.Errorf("can't encode video#31a460cc as nil")
	}
	b.ObjStart()
	b.PutID("video")
	b.Comma()
	b.FieldStart("duration")
	b.PutInt32(v.Duration)
	b.Comma()
	b.FieldStart("width")
	b.PutInt32(v.Width)
	b.Comma()
	b.FieldStart("height")
	b.PutInt32(v.Height)
	b.Comma()
	b.FieldStart("file_name")
	b.PutString(v.FileName)
	b.Comma()
	b.FieldStart("mime_type")
	b.PutString(v.MimeType)
	b.Comma()
	b.FieldStart("has_stickers")
	b.PutBool(v.HasStickers)
	b.Comma()
	b.FieldStart("supports_streaming")
	b.PutBool(v.SupportsStreaming)
	b.Comma()
	b.FieldStart("minithumbnail")
	if err := v.Minithumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode video#31a460cc: field minithumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("thumbnail")
	if err := v.Thumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode video#31a460cc: field thumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("video")
	if err := v.Video.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode video#31a460cc: field video: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (v *Video) DecodeTDLibJSON(b tdjson.Decoder) error {
	if v == nil {
		return fmt.Errorf("can't decode video#31a460cc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("video"); err != nil {
				return fmt.Errorf("unable to decode video#31a460cc: %w", err)
			}
		case "duration":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode video#31a460cc: field duration: %w", err)
			}
			v.Duration = value
		case "width":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode video#31a460cc: field width: %w", err)
			}
			v.Width = value
		case "height":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode video#31a460cc: field height: %w", err)
			}
			v.Height = value
		case "file_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode video#31a460cc: field file_name: %w", err)
			}
			v.FileName = value
		case "mime_type":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode video#31a460cc: field mime_type: %w", err)
			}
			v.MimeType = value
		case "has_stickers":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode video#31a460cc: field has_stickers: %w", err)
			}
			v.HasStickers = value
		case "supports_streaming":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode video#31a460cc: field supports_streaming: %w", err)
			}
			v.SupportsStreaming = value
		case "minithumbnail":
			if err := v.Minithumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode video#31a460cc: field minithumbnail: %w", err)
			}
		case "thumbnail":
			if err := v.Thumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode video#31a460cc: field thumbnail: %w", err)
			}
		case "video":
			if err := v.Video.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode video#31a460cc: field video: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDuration returns value of Duration field.
func (v *Video) GetDuration() (value int32) {
	if v == nil {
		return
	}
	return v.Duration
}

// GetWidth returns value of Width field.
func (v *Video) GetWidth() (value int32) {
	if v == nil {
		return
	}
	return v.Width
}

// GetHeight returns value of Height field.
func (v *Video) GetHeight() (value int32) {
	if v == nil {
		return
	}
	return v.Height
}

// GetFileName returns value of FileName field.
func (v *Video) GetFileName() (value string) {
	if v == nil {
		return
	}
	return v.FileName
}

// GetMimeType returns value of MimeType field.
func (v *Video) GetMimeType() (value string) {
	if v == nil {
		return
	}
	return v.MimeType
}

// GetHasStickers returns value of HasStickers field.
func (v *Video) GetHasStickers() (value bool) {
	if v == nil {
		return
	}
	return v.HasStickers
}

// GetSupportsStreaming returns value of SupportsStreaming field.
func (v *Video) GetSupportsStreaming() (value bool) {
	if v == nil {
		return
	}
	return v.SupportsStreaming
}

// GetMinithumbnail returns value of Minithumbnail field.
func (v *Video) GetMinithumbnail() (value Minithumbnail) {
	if v == nil {
		return
	}
	return v.Minithumbnail
}

// GetThumbnail returns value of Thumbnail field.
func (v *Video) GetThumbnail() (value Thumbnail) {
	if v == nil {
		return
	}
	return v.Thumbnail
}

// GetVideo returns value of Video field.
func (v *Video) GetVideo() (value File) {
	if v == nil {
		return
	}
	return v.Video
}
