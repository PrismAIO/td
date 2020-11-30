// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// AccountGetWallPapersRequest represents TL type `account.getWallPapers#aabb1763`.
type AccountGetWallPapersRequest struct {
	// Hash field of AccountGetWallPapersRequest.
	Hash int
}

// AccountGetWallPapersRequestTypeID is TL type id of AccountGetWallPapersRequest.
const AccountGetWallPapersRequestTypeID = 0xaabb1763

// Encode implements bin.Encoder.
func (g *AccountGetWallPapersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getWallPapers#aabb1763 as nil")
	}
	b.PutID(AccountGetWallPapersRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetWallPapersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getWallPapers#aabb1763 to nil")
	}
	if err := b.ConsumeID(AccountGetWallPapersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getWallPapers#aabb1763: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.getWallPapers#aabb1763: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetWallPapersRequest.
var (
	_ bin.Encoder = &AccountGetWallPapersRequest{}
	_ bin.Decoder = &AccountGetWallPapersRequest{}
)