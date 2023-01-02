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

// MessagesSetInlineBotResultsRequest represents TL type `messages.setInlineBotResults#eb5ea206`.
// Answer an inline query, for bots only
//
// See https://core.telegram.org/method/messages.setInlineBotResults for reference.
type MessagesSetInlineBotResultsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set this flag if the results are composed of media files
	Gallery bool
	// Set this flag if results may be cached on the server side only for the user that sent
	// the query. By default, results may be returned to any user who sends the same query
	Private bool
	// Unique identifier for the answered query
	QueryID int64
	// Vector of results for the inline query
	Results []InputBotInlineResultClass
	// The maximum amount of time in seconds that the result of the inline query may be
	// cached on the server. Defaults to 300.
	CacheTime int
	// Pass the offset that a client should send in the next query with the same text to
	// receive more results. Pass an empty string if there are no more results or if you
	// don't support pagination. Offset length can't exceed 64 bytes.
	//
	// Use SetNextOffset and GetNextOffset helpers.
	NextOffset string
	// If passed, clients will display a button with specified text that switches the user to
	// a private chat with the bot and sends the bot a start message with a certain parameter.
	//
	// Use SetSwitchPm and GetSwitchPm helpers.
	SwitchPm InlineBotSwitchPM
}

// MessagesSetInlineBotResultsRequestTypeID is TL type id of MessagesSetInlineBotResultsRequest.
const MessagesSetInlineBotResultsRequestTypeID = 0xeb5ea206

// Ensuring interfaces in compile-time for MessagesSetInlineBotResultsRequest.
var (
	_ bin.Encoder     = &MessagesSetInlineBotResultsRequest{}
	_ bin.Decoder     = &MessagesSetInlineBotResultsRequest{}
	_ bin.BareEncoder = &MessagesSetInlineBotResultsRequest{}
	_ bin.BareDecoder = &MessagesSetInlineBotResultsRequest{}
)

func (s *MessagesSetInlineBotResultsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Gallery == false) {
		return false
	}
	if !(s.Private == false) {
		return false
	}
	if !(s.QueryID == 0) {
		return false
	}
	if !(s.Results == nil) {
		return false
	}
	if !(s.CacheTime == 0) {
		return false
	}
	if !(s.NextOffset == "") {
		return false
	}
	if !(s.SwitchPm.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetInlineBotResultsRequest) String() string {
	if s == nil {
		return "MessagesSetInlineBotResultsRequest(nil)"
	}
	type Alias MessagesSetInlineBotResultsRequest
	return fmt.Sprintf("MessagesSetInlineBotResultsRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetInlineBotResultsRequest from given interface.
func (s *MessagesSetInlineBotResultsRequest) FillFrom(from interface {
	GetGallery() (value bool)
	GetPrivate() (value bool)
	GetQueryID() (value int64)
	GetResults() (value []InputBotInlineResultClass)
	GetCacheTime() (value int)
	GetNextOffset() (value string, ok bool)
	GetSwitchPm() (value InlineBotSwitchPM, ok bool)
}) {
	s.Gallery = from.GetGallery()
	s.Private = from.GetPrivate()
	s.QueryID = from.GetQueryID()
	s.Results = from.GetResults()
	s.CacheTime = from.GetCacheTime()
	if val, ok := from.GetNextOffset(); ok {
		s.NextOffset = val
	}

	if val, ok := from.GetSwitchPm(); ok {
		s.SwitchPm = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetInlineBotResultsRequest) TypeID() uint32 {
	return MessagesSetInlineBotResultsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetInlineBotResultsRequest) TypeName() string {
	return "messages.setInlineBotResults"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetInlineBotResultsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setInlineBotResults",
		ID:   MessagesSetInlineBotResultsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Gallery",
			SchemaName: "gallery",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Private",
			SchemaName: "private",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "Results",
			SchemaName: "results",
		},
		{
			Name:       "CacheTime",
			SchemaName: "cache_time",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "SwitchPm",
			SchemaName: "switch_pm",
			Null:       !s.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSetInlineBotResultsRequest) SetFlags() {
	if !(s.Gallery == false) {
		s.Flags.Set(0)
	}
	if !(s.Private == false) {
		s.Flags.Set(1)
	}
	if !(s.NextOffset == "") {
		s.Flags.Set(2)
	}
	if !(s.SwitchPm.Zero()) {
		s.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSetInlineBotResultsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setInlineBotResults#eb5ea206 as nil")
	}
	b.PutID(MessagesSetInlineBotResultsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetInlineBotResultsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setInlineBotResults#eb5ea206 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setInlineBotResults#eb5ea206: field flags: %w", err)
	}
	b.PutLong(s.QueryID)
	b.PutVectorHeader(len(s.Results))
	for idx, v := range s.Results {
		if v == nil {
			return fmt.Errorf("unable to encode messages.setInlineBotResults#eb5ea206: field results element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.setInlineBotResults#eb5ea206: field results element with index %d: %w", idx, err)
		}
	}
	b.PutInt(s.CacheTime)
	if s.Flags.Has(2) {
		b.PutString(s.NextOffset)
	}
	if s.Flags.Has(3) {
		if err := s.SwitchPm.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.setInlineBotResults#eb5ea206: field switch_pm: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSetInlineBotResultsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setInlineBotResults#eb5ea206 to nil")
	}
	if err := b.ConsumeID(MessagesSetInlineBotResultsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setInlineBotResults#eb5ea206: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetInlineBotResultsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setInlineBotResults#eb5ea206 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setInlineBotResults#eb5ea206: field flags: %w", err)
		}
	}
	s.Gallery = s.Flags.Has(0)
	s.Private = s.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setInlineBotResults#eb5ea206: field query_id: %w", err)
		}
		s.QueryID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setInlineBotResults#eb5ea206: field results: %w", err)
		}

		if headerLen > 0 {
			s.Results = make([]InputBotInlineResultClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputBotInlineResult(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.setInlineBotResults#eb5ea206: field results: %w", err)
			}
			s.Results = append(s.Results, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setInlineBotResults#eb5ea206: field cache_time: %w", err)
		}
		s.CacheTime = value
	}
	if s.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setInlineBotResults#eb5ea206: field next_offset: %w", err)
		}
		s.NextOffset = value
	}
	if s.Flags.Has(3) {
		if err := s.SwitchPm.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setInlineBotResults#eb5ea206: field switch_pm: %w", err)
		}
	}
	return nil
}

// SetGallery sets value of Gallery conditional field.
func (s *MessagesSetInlineBotResultsRequest) SetGallery(value bool) {
	if value {
		s.Flags.Set(0)
		s.Gallery = true
	} else {
		s.Flags.Unset(0)
		s.Gallery = false
	}
}

// GetGallery returns value of Gallery conditional field.
func (s *MessagesSetInlineBotResultsRequest) GetGallery() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// SetPrivate sets value of Private conditional field.
func (s *MessagesSetInlineBotResultsRequest) SetPrivate(value bool) {
	if value {
		s.Flags.Set(1)
		s.Private = true
	} else {
		s.Flags.Unset(1)
		s.Private = false
	}
}

// GetPrivate returns value of Private conditional field.
func (s *MessagesSetInlineBotResultsRequest) GetPrivate() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// GetQueryID returns value of QueryID field.
func (s *MessagesSetInlineBotResultsRequest) GetQueryID() (value int64) {
	if s == nil {
		return
	}
	return s.QueryID
}

// GetResults returns value of Results field.
func (s *MessagesSetInlineBotResultsRequest) GetResults() (value []InputBotInlineResultClass) {
	if s == nil {
		return
	}
	return s.Results
}

// GetCacheTime returns value of CacheTime field.
func (s *MessagesSetInlineBotResultsRequest) GetCacheTime() (value int) {
	if s == nil {
		return
	}
	return s.CacheTime
}

// SetNextOffset sets value of NextOffset conditional field.
func (s *MessagesSetInlineBotResultsRequest) SetNextOffset(value string) {
	s.Flags.Set(2)
	s.NextOffset = value
}

// GetNextOffset returns value of NextOffset conditional field and
// boolean which is true if field was set.
func (s *MessagesSetInlineBotResultsRequest) GetNextOffset() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.NextOffset, true
}

// SetSwitchPm sets value of SwitchPm conditional field.
func (s *MessagesSetInlineBotResultsRequest) SetSwitchPm(value InlineBotSwitchPM) {
	s.Flags.Set(3)
	s.SwitchPm = value
}

// GetSwitchPm returns value of SwitchPm conditional field and
// boolean which is true if field was set.
func (s *MessagesSetInlineBotResultsRequest) GetSwitchPm() (value InlineBotSwitchPM, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.SwitchPm, true
}

// MapResults returns field Results wrapped in InputBotInlineResultClassArray helper.
func (s *MessagesSetInlineBotResultsRequest) MapResults() (value InputBotInlineResultClassArray) {
	return InputBotInlineResultClassArray(s.Results)
}

// MessagesSetInlineBotResults invokes method messages.setInlineBotResults#eb5ea206 returning error if any.
// Answer an inline query, for bots only
//
// Possible errors:
//
//	400 ARTICLE_TITLE_EMPTY: The title of the article is empty.
//	400 AUDIO_CONTENT_URL_EMPTY: The remote URL specified in the content field is empty.
//	400 AUDIO_TITLE_EMPTY: An empty audio title was provided.
//	400 BUTTON_DATA_INVALID: The data of one or more of the buttons you provided is invalid.
//	400 BUTTON_TYPE_INVALID: The type of one or more of the buttons you provided is invalid.
//	400 BUTTON_URL_INVALID: Button URL invalid.
//	400 DOCUMENT_INVALID: The specified document is invalid.
//	400 FILE_CONTENT_TYPE_INVALID: File content-type is invalid.
//	400 FILE_TITLE_EMPTY: An empty file title was specified.
//	400 GIF_CONTENT_TYPE_INVALID: GIF content-type invalid.
//	400 MESSAGE_EMPTY: The provided message is empty.
//	400 MESSAGE_TOO_LONG: The provided message is too long.
//	400 NEXT_OFFSET_INVALID: The specified offset is longer than 64 bytes.
//	400 PHOTO_CONTENT_TYPE_INVALID: Photo mime-type invalid.
//	400 PHOTO_CONTENT_URL_EMPTY: Photo URL invalid.
//	400 PHOTO_INVALID: Photo invalid.
//	400 PHOTO_THUMB_URL_EMPTY: Photo thumbnail URL is empty.
//	400 QUERY_ID_INVALID: The query ID is invalid.
//	400 REPLY_MARKUP_INVALID: The provided reply markup is invalid.
//	400 RESULTS_TOO_MUCH: Too many results were provided.
//	400 RESULT_ID_DUPLICATE: You provided a duplicate result ID.
//	400 RESULT_ID_INVALID: One of the specified result IDs is invalid.
//	400 RESULT_TYPE_INVALID: Result type invalid.
//	400 SEND_MESSAGE_MEDIA_INVALID: Invalid media provided.
//	400 SEND_MESSAGE_TYPE_INVALID: The message type is invalid.
//	400 START_PARAM_EMPTY: The start parameter is empty.
//	400 START_PARAM_INVALID: Start parameter invalid.
//	400 STICKER_DOCUMENT_INVALID: The specified sticker document is invalid.
//	400 URL_INVALID: Invalid URL provided.
//	403 USER_BOT_INVALID: This method can only be called by a bot.
//	400 VIDEO_TITLE_EMPTY: The specified video title is empty.
//	400 WEBDOCUMENT_INVALID: Invalid webdocument URL provided.
//	400 WEBDOCUMENT_MIME_INVALID: Invalid webdocument mime type provided.
//	400 WEBDOCUMENT_SIZE_TOO_BIG: Webdocument is too big!
//	400 WEBDOCUMENT_URL_INVALID: The specified webdocument URL is invalid.
//
// See https://core.telegram.org/method/messages.setInlineBotResults for reference.
// Can be used by bots.
func (c *Client) MessagesSetInlineBotResults(ctx context.Context, request *MessagesSetInlineBotResultsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
