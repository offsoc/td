// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
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

// SuggestedPost represents TL type `suggestedPost#e8e37e5`.
//
// See https://core.telegram.org/constructor/suggestedPost for reference.
type SuggestedPost struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Accepted field of SuggestedPost.
	Accepted bool
	// Rejected field of SuggestedPost.
	Rejected bool
	// Price field of SuggestedPost.
	//
	// Use SetPrice and GetPrice helpers.
	Price StarsAmountClass
	// ScheduleDate field of SuggestedPost.
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
}

// SuggestedPostTypeID is TL type id of SuggestedPost.
const SuggestedPostTypeID = 0xe8e37e5

// Ensuring interfaces in compile-time for SuggestedPost.
var (
	_ bin.Encoder     = &SuggestedPost{}
	_ bin.Decoder     = &SuggestedPost{}
	_ bin.BareEncoder = &SuggestedPost{}
	_ bin.BareDecoder = &SuggestedPost{}
)

func (s *SuggestedPost) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Accepted == false) {
		return false
	}
	if !(s.Rejected == false) {
		return false
	}
	if !(s.Price == nil) {
		return false
	}
	if !(s.ScheduleDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedPost) String() string {
	if s == nil {
		return "SuggestedPost(nil)"
	}
	type Alias SuggestedPost
	return fmt.Sprintf("SuggestedPost%+v", Alias(*s))
}

// FillFrom fills SuggestedPost from given interface.
func (s *SuggestedPost) FillFrom(from interface {
	GetAccepted() (value bool)
	GetRejected() (value bool)
	GetPrice() (value StarsAmountClass, ok bool)
	GetScheduleDate() (value int, ok bool)
}) {
	s.Accepted = from.GetAccepted()
	s.Rejected = from.GetRejected()
	if val, ok := from.GetPrice(); ok {
		s.Price = val
	}

	if val, ok := from.GetScheduleDate(); ok {
		s.ScheduleDate = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedPost) TypeID() uint32 {
	return SuggestedPostTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedPost) TypeName() string {
	return "suggestedPost"
}

// TypeInfo returns info about TL type.
func (s *SuggestedPost) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedPost",
		ID:   SuggestedPostTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Accepted",
			SchemaName: "accepted",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Rejected",
			SchemaName: "rejected",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "Price",
			SchemaName: "price",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "ScheduleDate",
			SchemaName: "schedule_date",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *SuggestedPost) SetFlags() {
	if !(s.Accepted == false) {
		s.Flags.Set(1)
	}
	if !(s.Rejected == false) {
		s.Flags.Set(2)
	}
	if !(s.Price == nil) {
		s.Flags.Set(3)
	}
	if !(s.ScheduleDate == 0) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *SuggestedPost) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedPost#e8e37e5 as nil")
	}
	b.PutID(SuggestedPostTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedPost) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedPost#e8e37e5 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode suggestedPost#e8e37e5: field flags: %w", err)
	}
	if s.Flags.Has(3) {
		if s.Price == nil {
			return fmt.Errorf("unable to encode suggestedPost#e8e37e5: field price is nil")
		}
		if err := s.Price.Encode(b); err != nil {
			return fmt.Errorf("unable to encode suggestedPost#e8e37e5: field price: %w", err)
		}
	}
	if s.Flags.Has(0) {
		b.PutInt(s.ScheduleDate)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedPost) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedPost#e8e37e5 to nil")
	}
	if err := b.ConsumeID(SuggestedPostTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedPost#e8e37e5: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedPost) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedPost#e8e37e5 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode suggestedPost#e8e37e5: field flags: %w", err)
		}
	}
	s.Accepted = s.Flags.Has(1)
	s.Rejected = s.Flags.Has(2)
	if s.Flags.Has(3) {
		value, err := DecodeStarsAmount(b)
		if err != nil {
			return fmt.Errorf("unable to decode suggestedPost#e8e37e5: field price: %w", err)
		}
		s.Price = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode suggestedPost#e8e37e5: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	return nil
}

// SetAccepted sets value of Accepted conditional field.
func (s *SuggestedPost) SetAccepted(value bool) {
	if value {
		s.Flags.Set(1)
		s.Accepted = true
	} else {
		s.Flags.Unset(1)
		s.Accepted = false
	}
}

// GetAccepted returns value of Accepted conditional field.
func (s *SuggestedPost) GetAccepted() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// SetRejected sets value of Rejected conditional field.
func (s *SuggestedPost) SetRejected(value bool) {
	if value {
		s.Flags.Set(2)
		s.Rejected = true
	} else {
		s.Flags.Unset(2)
		s.Rejected = false
	}
}

// GetRejected returns value of Rejected conditional field.
func (s *SuggestedPost) GetRejected() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(2)
}

// SetPrice sets value of Price conditional field.
func (s *SuggestedPost) SetPrice(value StarsAmountClass) {
	s.Flags.Set(3)
	s.Price = value
}

// GetPrice returns value of Price conditional field and
// boolean which is true if field was set.
func (s *SuggestedPost) GetPrice() (value StarsAmountClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Price, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *SuggestedPost) SetScheduleDate(value int) {
	s.Flags.Set(0)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *SuggestedPost) GetScheduleDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ScheduleDate, true
}
