// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// InvokeAfterMsgRequest represents TL type `invokeAfterMsg#cb9f372d`.
// Invokes a query after successfull completion of one of the previous queries.
//
// See https://core.telegram.org/constructor/invokeAfterMsg for reference.
type InvokeAfterMsgRequest struct {
	// Message identifier on which a current query depends
	MsgID int64
	// The query itself
	Query bin.Object
}

// InvokeAfterMsgRequestTypeID is TL type id of InvokeAfterMsgRequest.
const InvokeAfterMsgRequestTypeID = 0xcb9f372d

func (i *InvokeAfterMsgRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.MsgID == 0) {
		return false
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeAfterMsgRequest) String() string {
	if i == nil {
		return "InvokeAfterMsgRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InvokeAfterMsgRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(i.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tQuery: ")
	sb.WriteString(fmt.Sprint(i.Query))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InvokeAfterMsgRequest) TypeID() uint32 {
	return InvokeAfterMsgRequestTypeID
}

// Encode implements bin.Encoder.
func (i *InvokeAfterMsgRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeAfterMsg#cb9f372d as nil")
	}
	b.PutID(InvokeAfterMsgRequestTypeID)
	b.PutLong(i.MsgID)
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeAfterMsg#cb9f372d: field query: %w", err)
	}
	return nil
}

// GetMsgID returns value of MsgID field.
func (i *InvokeAfterMsgRequest) GetMsgID() (value int64) {
	return i.MsgID
}

// GetQuery returns value of Query field.
func (i *InvokeAfterMsgRequest) GetQuery() (value bin.Object) {
	return i.Query
}

// Decode implements bin.Decoder.
func (i *InvokeAfterMsgRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeAfterMsg#cb9f372d to nil")
	}
	if err := b.ConsumeID(InvokeAfterMsgRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode invokeAfterMsg#cb9f372d: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode invokeAfterMsg#cb9f372d: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeAfterMsg#cb9f372d: field query: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InvokeAfterMsgRequest.
var (
	_ bin.Encoder = &InvokeAfterMsgRequest{}
	_ bin.Decoder = &InvokeAfterMsgRequest{}
)