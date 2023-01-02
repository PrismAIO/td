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

// MessagesCheckHistoryImportRequest represents TL type `messages.checkHistoryImport#43fe19f3`.
// Obtains information about a chat export file, generated by a foreign chat app, click
// here for more info about imported chats »¹.
//
// Links:
//  1. https://core.telegram.org/api/import
//
// See https://core.telegram.org/method/messages.checkHistoryImport for reference.
type MessagesCheckHistoryImportRequest struct {
	// Beginning of the message file; up to 100 lines.
	ImportHead string
}

// MessagesCheckHistoryImportRequestTypeID is TL type id of MessagesCheckHistoryImportRequest.
const MessagesCheckHistoryImportRequestTypeID = 0x43fe19f3

// Ensuring interfaces in compile-time for MessagesCheckHistoryImportRequest.
var (
	_ bin.Encoder     = &MessagesCheckHistoryImportRequest{}
	_ bin.Decoder     = &MessagesCheckHistoryImportRequest{}
	_ bin.BareEncoder = &MessagesCheckHistoryImportRequest{}
	_ bin.BareDecoder = &MessagesCheckHistoryImportRequest{}
)

func (c *MessagesCheckHistoryImportRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ImportHead == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesCheckHistoryImportRequest) String() string {
	if c == nil {
		return "MessagesCheckHistoryImportRequest(nil)"
	}
	type Alias MessagesCheckHistoryImportRequest
	return fmt.Sprintf("MessagesCheckHistoryImportRequest%+v", Alias(*c))
}

// FillFrom fills MessagesCheckHistoryImportRequest from given interface.
func (c *MessagesCheckHistoryImportRequest) FillFrom(from interface {
	GetImportHead() (value string)
}) {
	c.ImportHead = from.GetImportHead()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesCheckHistoryImportRequest) TypeID() uint32 {
	return MessagesCheckHistoryImportRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesCheckHistoryImportRequest) TypeName() string {
	return "messages.checkHistoryImport"
}

// TypeInfo returns info about TL type.
func (c *MessagesCheckHistoryImportRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.checkHistoryImport",
		ID:   MessagesCheckHistoryImportRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ImportHead",
			SchemaName: "import_head",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *MessagesCheckHistoryImportRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.checkHistoryImport#43fe19f3 as nil")
	}
	b.PutID(MessagesCheckHistoryImportRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *MessagesCheckHistoryImportRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.checkHistoryImport#43fe19f3 as nil")
	}
	b.PutString(c.ImportHead)
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesCheckHistoryImportRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.checkHistoryImport#43fe19f3 to nil")
	}
	if err := b.ConsumeID(MessagesCheckHistoryImportRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.checkHistoryImport#43fe19f3: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *MessagesCheckHistoryImportRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.checkHistoryImport#43fe19f3 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.checkHistoryImport#43fe19f3: field import_head: %w", err)
		}
		c.ImportHead = value
	}
	return nil
}

// GetImportHead returns value of ImportHead field.
func (c *MessagesCheckHistoryImportRequest) GetImportHead() (value string) {
	if c == nil {
		return
	}
	return c.ImportHead
}

// MessagesCheckHistoryImport invokes method messages.checkHistoryImport#43fe19f3 returning error if any.
// Obtains information about a chat export file, generated by a foreign chat app, click
// here for more info about imported chats »¹.
//
// Links:
//  1. https://core.telegram.org/api/import
//
// Possible errors:
//
//	400 IMPORT_FORMAT_UNRECOGNIZED: The specified chat export file was exported from an unsupported chat app.
//
// See https://core.telegram.org/method/messages.checkHistoryImport for reference.
func (c *Client) MessagesCheckHistoryImport(ctx context.Context, importhead string) (*MessagesHistoryImportParsed, error) {
	var result MessagesHistoryImportParsed

	request := &MessagesCheckHistoryImportRequest{
		ImportHead: importhead,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
