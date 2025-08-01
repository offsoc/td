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

// PhoneDeleteConferenceCallParticipantsRequest represents TL type `phone.deleteConferenceCallParticipants#8ca60525`.
//
// See https://core.telegram.org/method/phone.deleteConferenceCallParticipants for reference.
type PhoneDeleteConferenceCallParticipantsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// OnlyLeft field of PhoneDeleteConferenceCallParticipantsRequest.
	OnlyLeft bool
	// Kick field of PhoneDeleteConferenceCallParticipantsRequest.
	Kick bool
	// Call field of PhoneDeleteConferenceCallParticipantsRequest.
	Call InputGroupCallClass
	// IDs field of PhoneDeleteConferenceCallParticipantsRequest.
	IDs []int64
	// Block field of PhoneDeleteConferenceCallParticipantsRequest.
	Block []byte
}

// PhoneDeleteConferenceCallParticipantsRequestTypeID is TL type id of PhoneDeleteConferenceCallParticipantsRequest.
const PhoneDeleteConferenceCallParticipantsRequestTypeID = 0x8ca60525

// Ensuring interfaces in compile-time for PhoneDeleteConferenceCallParticipantsRequest.
var (
	_ bin.Encoder     = &PhoneDeleteConferenceCallParticipantsRequest{}
	_ bin.Decoder     = &PhoneDeleteConferenceCallParticipantsRequest{}
	_ bin.BareEncoder = &PhoneDeleteConferenceCallParticipantsRequest{}
	_ bin.BareDecoder = &PhoneDeleteConferenceCallParticipantsRequest{}
)

func (d *PhoneDeleteConferenceCallParticipantsRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.OnlyLeft == false) {
		return false
	}
	if !(d.Kick == false) {
		return false
	}
	if !(d.Call == nil) {
		return false
	}
	if !(d.IDs == nil) {
		return false
	}
	if !(d.Block == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *PhoneDeleteConferenceCallParticipantsRequest) String() string {
	if d == nil {
		return "PhoneDeleteConferenceCallParticipantsRequest(nil)"
	}
	type Alias PhoneDeleteConferenceCallParticipantsRequest
	return fmt.Sprintf("PhoneDeleteConferenceCallParticipantsRequest%+v", Alias(*d))
}

// FillFrom fills PhoneDeleteConferenceCallParticipantsRequest from given interface.
func (d *PhoneDeleteConferenceCallParticipantsRequest) FillFrom(from interface {
	GetOnlyLeft() (value bool)
	GetKick() (value bool)
	GetCall() (value InputGroupCallClass)
	GetIDs() (value []int64)
	GetBlock() (value []byte)
}) {
	d.OnlyLeft = from.GetOnlyLeft()
	d.Kick = from.GetKick()
	d.Call = from.GetCall()
	d.IDs = from.GetIDs()
	d.Block = from.GetBlock()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneDeleteConferenceCallParticipantsRequest) TypeID() uint32 {
	return PhoneDeleteConferenceCallParticipantsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneDeleteConferenceCallParticipantsRequest) TypeName() string {
	return "phone.deleteConferenceCallParticipants"
}

// TypeInfo returns info about TL type.
func (d *PhoneDeleteConferenceCallParticipantsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.deleteConferenceCallParticipants",
		ID:   PhoneDeleteConferenceCallParticipantsRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "OnlyLeft",
			SchemaName: "only_left",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "Kick",
			SchemaName: "kick",
			Null:       !d.Flags.Has(1),
		},
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "IDs",
			SchemaName: "ids",
		},
		{
			Name:       "Block",
			SchemaName: "block",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (d *PhoneDeleteConferenceCallParticipantsRequest) SetFlags() {
	if !(d.OnlyLeft == false) {
		d.Flags.Set(0)
	}
	if !(d.Kick == false) {
		d.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (d *PhoneDeleteConferenceCallParticipantsRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode phone.deleteConferenceCallParticipants#8ca60525 as nil")
	}
	b.PutID(PhoneDeleteConferenceCallParticipantsRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *PhoneDeleteConferenceCallParticipantsRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode phone.deleteConferenceCallParticipants#8ca60525 as nil")
	}
	d.SetFlags()
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.deleteConferenceCallParticipants#8ca60525: field flags: %w", err)
	}
	if d.Call == nil {
		return fmt.Errorf("unable to encode phone.deleteConferenceCallParticipants#8ca60525: field call is nil")
	}
	if err := d.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.deleteConferenceCallParticipants#8ca60525: field call: %w", err)
	}
	b.PutVectorHeader(len(d.IDs))
	for _, v := range d.IDs {
		b.PutLong(v)
	}
	b.PutBytes(d.Block)
	return nil
}

// Decode implements bin.Decoder.
func (d *PhoneDeleteConferenceCallParticipantsRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode phone.deleteConferenceCallParticipants#8ca60525 to nil")
	}
	if err := b.ConsumeID(PhoneDeleteConferenceCallParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.deleteConferenceCallParticipants#8ca60525: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *PhoneDeleteConferenceCallParticipantsRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode phone.deleteConferenceCallParticipants#8ca60525 to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.deleteConferenceCallParticipants#8ca60525: field flags: %w", err)
		}
	}
	d.OnlyLeft = d.Flags.Has(0)
	d.Kick = d.Flags.Has(1)
	{
		value, err := DecodeInputGroupCall(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.deleteConferenceCallParticipants#8ca60525: field call: %w", err)
		}
		d.Call = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.deleteConferenceCallParticipants#8ca60525: field ids: %w", err)
		}

		if headerLen > 0 {
			d.IDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode phone.deleteConferenceCallParticipants#8ca60525: field ids: %w", err)
			}
			d.IDs = append(d.IDs, value)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phone.deleteConferenceCallParticipants#8ca60525: field block: %w", err)
		}
		d.Block = value
	}
	return nil
}

// SetOnlyLeft sets value of OnlyLeft conditional field.
func (d *PhoneDeleteConferenceCallParticipantsRequest) SetOnlyLeft(value bool) {
	if value {
		d.Flags.Set(0)
		d.OnlyLeft = true
	} else {
		d.Flags.Unset(0)
		d.OnlyLeft = false
	}
}

// GetOnlyLeft returns value of OnlyLeft conditional field.
func (d *PhoneDeleteConferenceCallParticipantsRequest) GetOnlyLeft() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(0)
}

// SetKick sets value of Kick conditional field.
func (d *PhoneDeleteConferenceCallParticipantsRequest) SetKick(value bool) {
	if value {
		d.Flags.Set(1)
		d.Kick = true
	} else {
		d.Flags.Unset(1)
		d.Kick = false
	}
}

// GetKick returns value of Kick conditional field.
func (d *PhoneDeleteConferenceCallParticipantsRequest) GetKick() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(1)
}

// GetCall returns value of Call field.
func (d *PhoneDeleteConferenceCallParticipantsRequest) GetCall() (value InputGroupCallClass) {
	if d == nil {
		return
	}
	return d.Call
}

// GetIDs returns value of IDs field.
func (d *PhoneDeleteConferenceCallParticipantsRequest) GetIDs() (value []int64) {
	if d == nil {
		return
	}
	return d.IDs
}

// GetBlock returns value of Block field.
func (d *PhoneDeleteConferenceCallParticipantsRequest) GetBlock() (value []byte) {
	if d == nil {
		return
	}
	return d.Block
}

// PhoneDeleteConferenceCallParticipants invokes method phone.deleteConferenceCallParticipants#8ca60525 returning error if any.
//
// See https://core.telegram.org/method/phone.deleteConferenceCallParticipants for reference.
// Can be used by bots.
func (c *Client) PhoneDeleteConferenceCallParticipants(ctx context.Context, request *PhoneDeleteConferenceCallParticipantsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
