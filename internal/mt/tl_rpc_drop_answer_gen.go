// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// RPCAnswerUnknown represents TL type `rpc_answer_unknown#5e2ad36e`.
type RPCAnswerUnknown struct {
}

// RPCAnswerUnknownTypeID is TL type id of RPCAnswerUnknown.
const RPCAnswerUnknownTypeID = 0x5e2ad36e

// construct implements constructor of RPCDropAnswerClass.
func (r RPCAnswerUnknown) construct() RPCDropAnswerClass { return &r }

// Ensuring interfaces in compile-time for RPCAnswerUnknown.
var (
	_ bin.Encoder     = &RPCAnswerUnknown{}
	_ bin.Decoder     = &RPCAnswerUnknown{}
	_ bin.BareEncoder = &RPCAnswerUnknown{}
	_ bin.BareDecoder = &RPCAnswerUnknown{}

	_ RPCDropAnswerClass = &RPCAnswerUnknown{}
)

func (r *RPCAnswerUnknown) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *RPCAnswerUnknown) String() string {
	if r == nil {
		return "RPCAnswerUnknown(nil)"
	}
	type Alias RPCAnswerUnknown
	return fmt.Sprintf("RPCAnswerUnknown%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RPCAnswerUnknown) TypeID() uint32 {
	return RPCAnswerUnknownTypeID
}

// TypeName returns name of type in TL schema.
func (*RPCAnswerUnknown) TypeName() string {
	return "rpc_answer_unknown"
}

// TypeInfo returns info about TL type.
func (r *RPCAnswerUnknown) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "rpc_answer_unknown",
		ID:   RPCAnswerUnknownTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *RPCAnswerUnknown) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_answer_unknown#5e2ad36e as nil")
	}
	b.PutID(RPCAnswerUnknownTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RPCAnswerUnknown) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_answer_unknown#5e2ad36e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RPCAnswerUnknown) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_answer_unknown#5e2ad36e to nil")
	}
	if err := b.ConsumeID(RPCAnswerUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode rpc_answer_unknown#5e2ad36e: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RPCAnswerUnknown) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_answer_unknown#5e2ad36e to nil")
	}
	return nil
}

// RPCAnswerDroppedRunning represents TL type `rpc_answer_dropped_running#cd78e586`.
type RPCAnswerDroppedRunning struct {
}

// RPCAnswerDroppedRunningTypeID is TL type id of RPCAnswerDroppedRunning.
const RPCAnswerDroppedRunningTypeID = 0xcd78e586

// construct implements constructor of RPCDropAnswerClass.
func (r RPCAnswerDroppedRunning) construct() RPCDropAnswerClass { return &r }

// Ensuring interfaces in compile-time for RPCAnswerDroppedRunning.
var (
	_ bin.Encoder     = &RPCAnswerDroppedRunning{}
	_ bin.Decoder     = &RPCAnswerDroppedRunning{}
	_ bin.BareEncoder = &RPCAnswerDroppedRunning{}
	_ bin.BareDecoder = &RPCAnswerDroppedRunning{}

	_ RPCDropAnswerClass = &RPCAnswerDroppedRunning{}
)

func (r *RPCAnswerDroppedRunning) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *RPCAnswerDroppedRunning) String() string {
	if r == nil {
		return "RPCAnswerDroppedRunning(nil)"
	}
	type Alias RPCAnswerDroppedRunning
	return fmt.Sprintf("RPCAnswerDroppedRunning%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RPCAnswerDroppedRunning) TypeID() uint32 {
	return RPCAnswerDroppedRunningTypeID
}

// TypeName returns name of type in TL schema.
func (*RPCAnswerDroppedRunning) TypeName() string {
	return "rpc_answer_dropped_running"
}

// TypeInfo returns info about TL type.
func (r *RPCAnswerDroppedRunning) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "rpc_answer_dropped_running",
		ID:   RPCAnswerDroppedRunningTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *RPCAnswerDroppedRunning) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_answer_dropped_running#cd78e586 as nil")
	}
	b.PutID(RPCAnswerDroppedRunningTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RPCAnswerDroppedRunning) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_answer_dropped_running#cd78e586 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RPCAnswerDroppedRunning) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_answer_dropped_running#cd78e586 to nil")
	}
	if err := b.ConsumeID(RPCAnswerDroppedRunningTypeID); err != nil {
		return fmt.Errorf("unable to decode rpc_answer_dropped_running#cd78e586: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RPCAnswerDroppedRunning) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_answer_dropped_running#cd78e586 to nil")
	}
	return nil
}

// RPCAnswerDropped represents TL type `rpc_answer_dropped#a43ad8b7`.
type RPCAnswerDropped struct {
	// MsgID field of RPCAnswerDropped.
	MsgID int64
	// SeqNo field of RPCAnswerDropped.
	SeqNo int
	// Bytes field of RPCAnswerDropped.
	Bytes int
}

// RPCAnswerDroppedTypeID is TL type id of RPCAnswerDropped.
const RPCAnswerDroppedTypeID = 0xa43ad8b7

// construct implements constructor of RPCDropAnswerClass.
func (r RPCAnswerDropped) construct() RPCDropAnswerClass { return &r }

// Ensuring interfaces in compile-time for RPCAnswerDropped.
var (
	_ bin.Encoder     = &RPCAnswerDropped{}
	_ bin.Decoder     = &RPCAnswerDropped{}
	_ bin.BareEncoder = &RPCAnswerDropped{}
	_ bin.BareDecoder = &RPCAnswerDropped{}

	_ RPCDropAnswerClass = &RPCAnswerDropped{}
)

func (r *RPCAnswerDropped) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.MsgID == 0) {
		return false
	}
	if !(r.SeqNo == 0) {
		return false
	}
	if !(r.Bytes == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RPCAnswerDropped) String() string {
	if r == nil {
		return "RPCAnswerDropped(nil)"
	}
	type Alias RPCAnswerDropped
	return fmt.Sprintf("RPCAnswerDropped%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RPCAnswerDropped) TypeID() uint32 {
	return RPCAnswerDroppedTypeID
}

// TypeName returns name of type in TL schema.
func (*RPCAnswerDropped) TypeName() string {
	return "rpc_answer_dropped"
}

// TypeInfo returns info about TL type.
func (r *RPCAnswerDropped) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "rpc_answer_dropped",
		ID:   RPCAnswerDroppedTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "SeqNo",
			SchemaName: "seq_no",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RPCAnswerDropped) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_answer_dropped#a43ad8b7 as nil")
	}
	b.PutID(RPCAnswerDroppedTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RPCAnswerDropped) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_answer_dropped#a43ad8b7 as nil")
	}
	b.PutLong(r.MsgID)
	b.PutInt(r.SeqNo)
	b.PutInt(r.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (r *RPCAnswerDropped) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_answer_dropped#a43ad8b7 to nil")
	}
	if err := b.ConsumeID(RPCAnswerDroppedTypeID); err != nil {
		return fmt.Errorf("unable to decode rpc_answer_dropped#a43ad8b7: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RPCAnswerDropped) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_answer_dropped#a43ad8b7 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_answer_dropped#a43ad8b7: field msg_id: %w", err)
		}
		r.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_answer_dropped#a43ad8b7: field seq_no: %w", err)
		}
		r.SeqNo = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_answer_dropped#a43ad8b7: field bytes: %w", err)
		}
		r.Bytes = value
	}
	return nil
}

// GetMsgID returns value of MsgID field.
func (r *RPCAnswerDropped) GetMsgID() (value int64) {
	if r == nil {
		return
	}
	return r.MsgID
}

// GetSeqNo returns value of SeqNo field.
func (r *RPCAnswerDropped) GetSeqNo() (value int) {
	if r == nil {
		return
	}
	return r.SeqNo
}

// GetBytes returns value of Bytes field.
func (r *RPCAnswerDropped) GetBytes() (value int) {
	if r == nil {
		return
	}
	return r.Bytes
}

// RPCDropAnswerClassName is schema name of RPCDropAnswerClass.
const RPCDropAnswerClassName = "RpcDropAnswer"

// RPCDropAnswerClass represents RpcDropAnswer generic type.
//
// Example:
//
//	g, err := mt.DecodeRPCDropAnswer(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *mt.RPCAnswerUnknown: // rpc_answer_unknown#5e2ad36e
//	case *mt.RPCAnswerDroppedRunning: // rpc_answer_dropped_running#cd78e586
//	case *mt.RPCAnswerDropped: // rpc_answer_dropped#a43ad8b7
//	default: panic(v)
//	}
type RPCDropAnswerClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() RPCDropAnswerClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeRPCDropAnswer implements binary de-serialization for RPCDropAnswerClass.
func DecodeRPCDropAnswer(buf *bin.Buffer) (RPCDropAnswerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case RPCAnswerUnknownTypeID:
		// Decoding rpc_answer_unknown#5e2ad36e.
		v := RPCAnswerUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RPCDropAnswerClass: %w", err)
		}
		return &v, nil
	case RPCAnswerDroppedRunningTypeID:
		// Decoding rpc_answer_dropped_running#cd78e586.
		v := RPCAnswerDroppedRunning{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RPCDropAnswerClass: %w", err)
		}
		return &v, nil
	case RPCAnswerDroppedTypeID:
		// Decoding rpc_answer_dropped#a43ad8b7.
		v := RPCAnswerDropped{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RPCDropAnswerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode RPCDropAnswerClass: %w", bin.NewUnexpectedID(id))
	}
}

// RPCDropAnswer boxes the RPCDropAnswerClass providing a helper.
type RPCDropAnswerBox struct {
	RpcDropAnswer RPCDropAnswerClass
}

// Decode implements bin.Decoder for RPCDropAnswerBox.
func (b *RPCDropAnswerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode RPCDropAnswerBox to nil")
	}
	v, err := DecodeRPCDropAnswer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RpcDropAnswer = v
	return nil
}

// Encode implements bin.Encode for RPCDropAnswerBox.
func (b *RPCDropAnswerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.RpcDropAnswer == nil {
		return fmt.Errorf("unable to encode RPCDropAnswerClass as nil")
	}
	return b.RpcDropAnswer.Encode(buf)
}
