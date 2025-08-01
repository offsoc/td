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

// PaymentsGetResaleStarGiftsRequest represents TL type `payments.getResaleStarGifts#7a5fa236`.
//
// See https://core.telegram.org/method/payments.getResaleStarGifts for reference.
type PaymentsGetResaleStarGiftsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// SortByPrice field of PaymentsGetResaleStarGiftsRequest.
	SortByPrice bool
	// SortByNum field of PaymentsGetResaleStarGiftsRequest.
	SortByNum bool
	// AttributesHash field of PaymentsGetResaleStarGiftsRequest.
	//
	// Use SetAttributesHash and GetAttributesHash helpers.
	AttributesHash int64
	// GiftID field of PaymentsGetResaleStarGiftsRequest.
	GiftID int64
	// Attributes field of PaymentsGetResaleStarGiftsRequest.
	//
	// Use SetAttributes and GetAttributes helpers.
	Attributes []StarGiftAttributeIDClass
	// Offset field of PaymentsGetResaleStarGiftsRequest.
	Offset string
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// PaymentsGetResaleStarGiftsRequestTypeID is TL type id of PaymentsGetResaleStarGiftsRequest.
const PaymentsGetResaleStarGiftsRequestTypeID = 0x7a5fa236

// Ensuring interfaces in compile-time for PaymentsGetResaleStarGiftsRequest.
var (
	_ bin.Encoder     = &PaymentsGetResaleStarGiftsRequest{}
	_ bin.Decoder     = &PaymentsGetResaleStarGiftsRequest{}
	_ bin.BareEncoder = &PaymentsGetResaleStarGiftsRequest{}
	_ bin.BareDecoder = &PaymentsGetResaleStarGiftsRequest{}
)

func (g *PaymentsGetResaleStarGiftsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.SortByPrice == false) {
		return false
	}
	if !(g.SortByNum == false) {
		return false
	}
	if !(g.AttributesHash == 0) {
		return false
	}
	if !(g.GiftID == 0) {
		return false
	}
	if !(g.Attributes == nil) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PaymentsGetResaleStarGiftsRequest) String() string {
	if g == nil {
		return "PaymentsGetResaleStarGiftsRequest(nil)"
	}
	type Alias PaymentsGetResaleStarGiftsRequest
	return fmt.Sprintf("PaymentsGetResaleStarGiftsRequest%+v", Alias(*g))
}

// FillFrom fills PaymentsGetResaleStarGiftsRequest from given interface.
func (g *PaymentsGetResaleStarGiftsRequest) FillFrom(from interface {
	GetSortByPrice() (value bool)
	GetSortByNum() (value bool)
	GetAttributesHash() (value int64, ok bool)
	GetGiftID() (value int64)
	GetAttributes() (value []StarGiftAttributeIDClass, ok bool)
	GetOffset() (value string)
	GetLimit() (value int)
}) {
	g.SortByPrice = from.GetSortByPrice()
	g.SortByNum = from.GetSortByNum()
	if val, ok := from.GetAttributesHash(); ok {
		g.AttributesHash = val
	}

	g.GiftID = from.GetGiftID()
	if val, ok := from.GetAttributes(); ok {
		g.Attributes = val
	}

	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetResaleStarGiftsRequest) TypeID() uint32 {
	return PaymentsGetResaleStarGiftsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetResaleStarGiftsRequest) TypeName() string {
	return "payments.getResaleStarGifts"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetResaleStarGiftsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getResaleStarGifts",
		ID:   PaymentsGetResaleStarGiftsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SortByPrice",
			SchemaName: "sort_by_price",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "SortByNum",
			SchemaName: "sort_by_num",
			Null:       !g.Flags.Has(2),
		},
		{
			Name:       "AttributesHash",
			SchemaName: "attributes_hash",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "GiftID",
			SchemaName: "gift_id",
		},
		{
			Name:       "Attributes",
			SchemaName: "attributes",
			Null:       !g.Flags.Has(3),
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *PaymentsGetResaleStarGiftsRequest) SetFlags() {
	if !(g.SortByPrice == false) {
		g.Flags.Set(1)
	}
	if !(g.SortByNum == false) {
		g.Flags.Set(2)
	}
	if !(g.AttributesHash == 0) {
		g.Flags.Set(0)
	}
	if !(g.Attributes == nil) {
		g.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (g *PaymentsGetResaleStarGiftsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getResaleStarGifts#7a5fa236 as nil")
	}
	b.PutID(PaymentsGetResaleStarGiftsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PaymentsGetResaleStarGiftsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getResaleStarGifts#7a5fa236 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.getResaleStarGifts#7a5fa236: field flags: %w", err)
	}
	if g.Flags.Has(0) {
		b.PutLong(g.AttributesHash)
	}
	b.PutLong(g.GiftID)
	if g.Flags.Has(3) {
		b.PutVectorHeader(len(g.Attributes))
		for idx, v := range g.Attributes {
			if v == nil {
				return fmt.Errorf("unable to encode payments.getResaleStarGifts#7a5fa236: field attributes element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode payments.getResaleStarGifts#7a5fa236: field attributes element with index %d: %w", idx, err)
			}
		}
	}
	b.PutString(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetResaleStarGiftsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getResaleStarGifts#7a5fa236 to nil")
	}
	if err := b.ConsumeID(PaymentsGetResaleStarGiftsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getResaleStarGifts#7a5fa236: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PaymentsGetResaleStarGiftsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getResaleStarGifts#7a5fa236 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.getResaleStarGifts#7a5fa236: field flags: %w", err)
		}
	}
	g.SortByPrice = g.Flags.Has(1)
	g.SortByNum = g.Flags.Has(2)
	if g.Flags.Has(0) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getResaleStarGifts#7a5fa236: field attributes_hash: %w", err)
		}
		g.AttributesHash = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getResaleStarGifts#7a5fa236: field gift_id: %w", err)
		}
		g.GiftID = value
	}
	if g.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getResaleStarGifts#7a5fa236: field attributes: %w", err)
		}

		if headerLen > 0 {
			g.Attributes = make([]StarGiftAttributeIDClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStarGiftAttributeID(b)
			if err != nil {
				return fmt.Errorf("unable to decode payments.getResaleStarGifts#7a5fa236: field attributes: %w", err)
			}
			g.Attributes = append(g.Attributes, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getResaleStarGifts#7a5fa236: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.getResaleStarGifts#7a5fa236: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// SetSortByPrice sets value of SortByPrice conditional field.
func (g *PaymentsGetResaleStarGiftsRequest) SetSortByPrice(value bool) {
	if value {
		g.Flags.Set(1)
		g.SortByPrice = true
	} else {
		g.Flags.Unset(1)
		g.SortByPrice = false
	}
}

// GetSortByPrice returns value of SortByPrice conditional field.
func (g *PaymentsGetResaleStarGiftsRequest) GetSortByPrice() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(1)
}

// SetSortByNum sets value of SortByNum conditional field.
func (g *PaymentsGetResaleStarGiftsRequest) SetSortByNum(value bool) {
	if value {
		g.Flags.Set(2)
		g.SortByNum = true
	} else {
		g.Flags.Unset(2)
		g.SortByNum = false
	}
}

// GetSortByNum returns value of SortByNum conditional field.
func (g *PaymentsGetResaleStarGiftsRequest) GetSortByNum() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(2)
}

// SetAttributesHash sets value of AttributesHash conditional field.
func (g *PaymentsGetResaleStarGiftsRequest) SetAttributesHash(value int64) {
	g.Flags.Set(0)
	g.AttributesHash = value
}

// GetAttributesHash returns value of AttributesHash conditional field and
// boolean which is true if field was set.
func (g *PaymentsGetResaleStarGiftsRequest) GetAttributesHash() (value int64, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.AttributesHash, true
}

// GetGiftID returns value of GiftID field.
func (g *PaymentsGetResaleStarGiftsRequest) GetGiftID() (value int64) {
	if g == nil {
		return
	}
	return g.GiftID
}

// SetAttributes sets value of Attributes conditional field.
func (g *PaymentsGetResaleStarGiftsRequest) SetAttributes(value []StarGiftAttributeIDClass) {
	g.Flags.Set(3)
	g.Attributes = value
}

// GetAttributes returns value of Attributes conditional field and
// boolean which is true if field was set.
func (g *PaymentsGetResaleStarGiftsRequest) GetAttributes() (value []StarGiftAttributeIDClass, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(3) {
		return value, false
	}
	return g.Attributes, true
}

// GetOffset returns value of Offset field.
func (g *PaymentsGetResaleStarGiftsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *PaymentsGetResaleStarGiftsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// MapAttributes returns field Attributes wrapped in StarGiftAttributeIDClassArray helper.
func (g *PaymentsGetResaleStarGiftsRequest) MapAttributes() (value StarGiftAttributeIDClassArray, ok bool) {
	if !g.Flags.Has(3) {
		return value, false
	}
	return StarGiftAttributeIDClassArray(g.Attributes), true
}

// PaymentsGetResaleStarGifts invokes method payments.getResaleStarGifts#7a5fa236 returning error if any.
//
// See https://core.telegram.org/method/payments.getResaleStarGifts for reference.
// Can be used by bots.
func (c *Client) PaymentsGetResaleStarGifts(ctx context.Context, request *PaymentsGetResaleStarGiftsRequest) (*PaymentsResaleStarGifts, error) {
	var result PaymentsResaleStarGifts

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
