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

// ChannelsEditLocationRequest represents TL type `channels.editLocation#58e63f6d`.
type ChannelsEditLocationRequest struct {
	// Channel field of ChannelsEditLocationRequest.
	Channel InputChannelClass
	// GeoPoint field of ChannelsEditLocationRequest.
	GeoPoint InputGeoPointClass
	// Address field of ChannelsEditLocationRequest.
	Address string
}

// ChannelsEditLocationRequestTypeID is TL type id of ChannelsEditLocationRequest.
const ChannelsEditLocationRequestTypeID = 0x58e63f6d

// Encode implements bin.Encoder.
func (e *ChannelsEditLocationRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editLocation#58e63f6d as nil")
	}
	b.PutID(ChannelsEditLocationRequestTypeID)
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field channel: %w", err)
	}
	if e.GeoPoint == nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field geo_point is nil")
	}
	if err := e.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field geo_point: %w", err)
	}
	b.PutString(e.Address)
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditLocationRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editLocation#58e63f6d to nil")
	}
	if err := b.ConsumeID(ChannelsEditLocationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: field geo_point: %w", err)
		}
		e.GeoPoint = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: field address: %w", err)
		}
		e.Address = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsEditLocationRequest.
var (
	_ bin.Encoder = &ChannelsEditLocationRequest{}
	_ bin.Decoder = &ChannelsEditLocationRequest{}
)