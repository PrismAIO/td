// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// EchoVectorRequest represents TL type `echoVector#d4785939`.
//
// See https://localhost:80/doc/method/echoVector for reference.
type EchoVectorRequest struct {
	// Ids field of EchoVectorRequest.
	Ids []int
}

// EchoVectorRequestTypeID is TL type id of EchoVectorRequest.
const EchoVectorRequestTypeID = 0xd4785939

// Encode implements bin.Encoder.
func (e *EchoVectorRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode echoVector#d4785939 as nil")
	}
	b.PutID(EchoVectorRequestTypeID)
	b.PutVectorHeader(len(e.Ids))
	for _, v := range e.Ids {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EchoVectorRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode echoVector#d4785939 to nil")
	}
	if err := b.ConsumeID(EchoVectorRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode echoVector#d4785939: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode echoVector#d4785939: field ids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode echoVector#d4785939: field ids: %w", err)
			}
			e.Ids = append(e.Ids, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for EchoVectorRequest.
var (
	_ bin.Encoder = &EchoVectorRequest{}
	_ bin.Decoder = &EchoVectorRequest{}
)

// EchoVector invokes method echoVector#d4785939 returning error if any.
//
// See https://localhost:80/doc/method/echoVector for reference.
func (c *Client) EchoVector(ctx context.Context, request *EchoVectorRequest) ([]int, error) {
	var result IntVector
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}