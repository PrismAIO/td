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

// MessagesGetSearchResultsCalendarRequest represents TL type `messages.getSearchResultsCalendar#49f0bde9`.
// Returns information about the next messages of the specified type in the chat split by
// days.
// Returns the results in reverse chronological order.
// Can return partial results for the last returned day.
//
// See https://core.telegram.org/method/messages.getSearchResultsCalendar for reference.
type MessagesGetSearchResultsCalendarRequest struct {
	// Peer where to search
	Peer InputPeerClass
	// Message filter, inputMessagesFilterEmpty¹, inputMessagesFilterMyMentions² filters
	// are not supported by this method.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/inputMessagesFilterEmpty
	//  2) https://core.telegram.org/constructor/inputMessagesFilterMyMentions
	Filter MessagesFilterClass
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetDate int
}

// MessagesGetSearchResultsCalendarRequestTypeID is TL type id of MessagesGetSearchResultsCalendarRequest.
const MessagesGetSearchResultsCalendarRequestTypeID = 0x49f0bde9

// Ensuring interfaces in compile-time for MessagesGetSearchResultsCalendarRequest.
var (
	_ bin.Encoder     = &MessagesGetSearchResultsCalendarRequest{}
	_ bin.Decoder     = &MessagesGetSearchResultsCalendarRequest{}
	_ bin.BareEncoder = &MessagesGetSearchResultsCalendarRequest{}
	_ bin.BareDecoder = &MessagesGetSearchResultsCalendarRequest{}
)

func (g *MessagesGetSearchResultsCalendarRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.OffsetDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetSearchResultsCalendarRequest) String() string {
	if g == nil {
		return "MessagesGetSearchResultsCalendarRequest(nil)"
	}
	type Alias MessagesGetSearchResultsCalendarRequest
	return fmt.Sprintf("MessagesGetSearchResultsCalendarRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetSearchResultsCalendarRequest from given interface.
func (g *MessagesGetSearchResultsCalendarRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetFilter() (value MessagesFilterClass)
	GetOffsetID() (value int)
	GetOffsetDate() (value int)
}) {
	g.Peer = from.GetPeer()
	g.Filter = from.GetFilter()
	g.OffsetID = from.GetOffsetID()
	g.OffsetDate = from.GetOffsetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetSearchResultsCalendarRequest) TypeID() uint32 {
	return MessagesGetSearchResultsCalendarRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetSearchResultsCalendarRequest) TypeName() string {
	return "messages.getSearchResultsCalendar"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetSearchResultsCalendarRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getSearchResultsCalendar",
		ID:   MessagesGetSearchResultsCalendarRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "OffsetDate",
			SchemaName: "offset_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetSearchResultsCalendarRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSearchResultsCalendar#49f0bde9 as nil")
	}
	b.PutID(MessagesGetSearchResultsCalendarRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetSearchResultsCalendarRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSearchResultsCalendar#49f0bde9 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getSearchResultsCalendar#49f0bde9: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getSearchResultsCalendar#49f0bde9: field peer: %w", err)
	}
	if g.Filter == nil {
		return fmt.Errorf("unable to encode messages.getSearchResultsCalendar#49f0bde9: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getSearchResultsCalendar#49f0bde9: field filter: %w", err)
	}
	b.PutInt(g.OffsetID)
	b.PutInt(g.OffsetDate)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetSearchResultsCalendarRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSearchResultsCalendar#49f0bde9 to nil")
	}
	if err := b.ConsumeID(MessagesGetSearchResultsCalendarRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getSearchResultsCalendar#49f0bde9: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetSearchResultsCalendarRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSearchResultsCalendar#49f0bde9 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchResultsCalendar#49f0bde9: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := DecodeMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchResultsCalendar#49f0bde9: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchResultsCalendar#49f0bde9: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchResultsCalendar#49f0bde9: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetSearchResultsCalendarRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetFilter returns value of Filter field.
func (g *MessagesGetSearchResultsCalendarRequest) GetFilter() (value MessagesFilterClass) {
	if g == nil {
		return
	}
	return g.Filter
}

// GetOffsetID returns value of OffsetID field.
func (g *MessagesGetSearchResultsCalendarRequest) GetOffsetID() (value int) {
	if g == nil {
		return
	}
	return g.OffsetID
}

// GetOffsetDate returns value of OffsetDate field.
func (g *MessagesGetSearchResultsCalendarRequest) GetOffsetDate() (value int) {
	if g == nil {
		return
	}
	return g.OffsetDate
}

// MessagesGetSearchResultsCalendar invokes method messages.getSearchResultsCalendar#49f0bde9 returning error if any.
// Returns information about the next messages of the specified type in the chat split by
// days.
// Returns the results in reverse chronological order.
// Can return partial results for the last returned day.
//
// Possible errors:
//
//	400 FILTER_NOT_SUPPORTED: The specified filter cannot be used in this context.
//
// See https://core.telegram.org/method/messages.getSearchResultsCalendar for reference.
func (c *Client) MessagesGetSearchResultsCalendar(ctx context.Context, request *MessagesGetSearchResultsCalendarRequest) (*MessagesSearchResultsCalendar, error) {
	var result MessagesSearchResultsCalendar

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
