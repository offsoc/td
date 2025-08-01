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

// StarsTransaction represents TL type `starsTransaction#13659eb0`.
// Represents a Telegram Stars transaction »¹.
//
// Links:
//  1. https://core.telegram.org/api/stars
//
// See https://core.telegram.org/constructor/starsTransaction for reference.
type StarsTransaction struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this transaction is a refund.
	Refund bool
	// The transaction is currently pending.
	Pending bool
	// This transaction has failed.
	Failed bool
	// This transaction was a gift from the user in peer.peer.
	Gift bool
	// This transaction is a paid reaction »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/reactions#paid-reactions
	Reaction bool
	// StargiftUpgrade field of StarsTransaction.
	StargiftUpgrade bool
	// BusinessTransfer field of StarsTransaction.
	BusinessTransfer bool
	// StargiftResale field of StarsTransaction.
	StargiftResale bool
	// Transaction ID.
	ID string
	// Amount field of StarsTransaction.
	Amount StarsAmountClass
	// Date of the transaction (unixtime).
	Date int
	// Source of the incoming transaction, or its recipient for outgoing transactions.
	Peer StarsTransactionPeerClass
	// For transactions with bots, title of the bought product.
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// For transactions with bots, description of the bought product.
	//
	// Use SetDescription and GetDescription helpers.
	Description string
	// For transactions with bots, photo of the bought product.
	//
	// Use SetPhoto and GetPhoto helpers.
	Photo WebDocumentClass
	// If neither pending nor failed are set, the transaction was completed successfully, and
	// this field will contain the point in time (Unix timestamp) when the withdrawal was
	// completed successfully.
	//
	// Use SetTransactionDate and GetTransactionDate helpers.
	TransactionDate int
	// If neither pending nor failed are set, the transaction was completed successfully, and
	// this field will contain a URL where the withdrawal transaction can be viewed.
	//
	// Use SetTransactionURL and GetTransactionURL helpers.
	TransactionURL string
	// Bot specified invoice payload (i.e. the payload passed to inputMediaInvoice¹ when
	// creating the invoice²).
	//
	// Links:
	//  1) https://core.telegram.org/constructor/inputMediaInvoice
	//  2) https://core.telegram.org/api/payments
	//
	// Use SetBotPayload and GetBotPayload helpers.
	BotPayload []byte
	// For paid media transactions »¹, message ID of the paid media posted to peer.peer
	// (can point to a deleted message; either way, extended_media will always contain the
	// bought media).
	//
	// Links:
	//  1) https://core.telegram.org/api/paid-media
	//
	// Use SetMsgID and GetMsgID helpers.
	MsgID int
	// The purchased paid media »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/paid-media
	//
	// Use SetExtendedMedia and GetExtendedMedia helpers.
	ExtendedMedia []MessageMediaClass
	// The number of seconds between consecutive Telegram Star debiting for Telegram Star
	// subscriptions »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/stars#star-subscriptions
	//
	// Use SetSubscriptionPeriod and GetSubscriptionPeriod helpers.
	SubscriptionPeriod int
	// ID of the message containing the messageMediaGiveaway¹, for incoming star giveaway
	// prizes².
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messageMediaGiveaway
	//  2) https://core.telegram.org/api/giveaways#star-giveaways
	//
	// Use SetGiveawayPostID and GetGiveawayPostID helpers.
	GiveawayPostID int
	// This transaction indicates a purchase or a sale (conversion back to Stars) of a gift
	// »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/stars
	//
	// Use SetStargift and GetStargift helpers.
	Stargift StarGiftClass
	// This transaction is payment for paid bot broadcasts¹.  Paid broadcasts are only
	// allowed if the allow_paid_floodskip parameter of messages.sendMessage² and other
	// message sending methods is set while trying to broadcast more than 30 messages per
	// second to bot users. The integer value returned by this flag indicates the number of
	// billed API calls.
	//
	// Links:
	//  1) https://core.telegram.org/bots/faq#how-can-i-message-all-of-my-bot-39s-subscribers-at-once
	//  2) https://core.telegram.org/method/messages.sendMessage
	//
	// Use SetFloodskipNumber and GetFloodskipNumber helpers.
	FloodskipNumber int
	// This transaction is the receival (or refund) of an affiliate commission¹ (i.e. this
	// is the transaction received by the peer that created the referral link², flag 17 is
	// for transactions made by users that imported the referral link).
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/referrals
	//  2) https://core.telegram.org/api/links#referral-links
	//
	// Use SetStarrefCommissionPermille and GetStarrefCommissionPermille helpers.
	StarrefCommissionPermille int
	// For transactions made by referred users¹, the peer that received the affiliate
	// commission.
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/referrals
	//
	// Use SetStarrefPeer and GetStarrefPeer helpers.
	StarrefPeer PeerClass
	// For transactions made by referred users¹, the amount of Telegram Stars received by
	// the affiliate, can be negative for refunds.
	//
	// Links:
	//  1) https://core.telegram.org/api/bots/referrals
	//
	// Use SetStarrefAmount and GetStarrefAmount helpers.
	StarrefAmount StarsAmountClass
	// PaidMessages field of StarsTransaction.
	//
	// Use SetPaidMessages and GetPaidMessages helpers.
	PaidMessages int
	// PremiumGiftMonths field of StarsTransaction.
	//
	// Use SetPremiumGiftMonths and GetPremiumGiftMonths helpers.
	PremiumGiftMonths int
	// AdsProceedsFromDate field of StarsTransaction.
	//
	// Use SetAdsProceedsFromDate and GetAdsProceedsFromDate helpers.
	AdsProceedsFromDate int
	// AdsProceedsToDate field of StarsTransaction.
	//
	// Use SetAdsProceedsToDate and GetAdsProceedsToDate helpers.
	AdsProceedsToDate int
}

// StarsTransactionTypeID is TL type id of StarsTransaction.
const StarsTransactionTypeID = 0x13659eb0

// Ensuring interfaces in compile-time for StarsTransaction.
var (
	_ bin.Encoder     = &StarsTransaction{}
	_ bin.Decoder     = &StarsTransaction{}
	_ bin.BareEncoder = &StarsTransaction{}
	_ bin.BareDecoder = &StarsTransaction{}
)

func (s *StarsTransaction) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Refund == false) {
		return false
	}
	if !(s.Pending == false) {
		return false
	}
	if !(s.Failed == false) {
		return false
	}
	if !(s.Gift == false) {
		return false
	}
	if !(s.Reaction == false) {
		return false
	}
	if !(s.StargiftUpgrade == false) {
		return false
	}
	if !(s.BusinessTransfer == false) {
		return false
	}
	if !(s.StargiftResale == false) {
		return false
	}
	if !(s.ID == "") {
		return false
	}
	if !(s.Amount == nil) {
		return false
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.Title == "") {
		return false
	}
	if !(s.Description == "") {
		return false
	}
	if !(s.Photo == nil) {
		return false
	}
	if !(s.TransactionDate == 0) {
		return false
	}
	if !(s.TransactionURL == "") {
		return false
	}
	if !(s.BotPayload == nil) {
		return false
	}
	if !(s.MsgID == 0) {
		return false
	}
	if !(s.ExtendedMedia == nil) {
		return false
	}
	if !(s.SubscriptionPeriod == 0) {
		return false
	}
	if !(s.GiveawayPostID == 0) {
		return false
	}
	if !(s.Stargift == nil) {
		return false
	}
	if !(s.FloodskipNumber == 0) {
		return false
	}
	if !(s.StarrefCommissionPermille == 0) {
		return false
	}
	if !(s.StarrefPeer == nil) {
		return false
	}
	if !(s.StarrefAmount == nil) {
		return false
	}
	if !(s.PaidMessages == 0) {
		return false
	}
	if !(s.PremiumGiftMonths == 0) {
		return false
	}
	if !(s.AdsProceedsFromDate == 0) {
		return false
	}
	if !(s.AdsProceedsToDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarsTransaction) String() string {
	if s == nil {
		return "StarsTransaction(nil)"
	}
	type Alias StarsTransaction
	return fmt.Sprintf("StarsTransaction%+v", Alias(*s))
}

// FillFrom fills StarsTransaction from given interface.
func (s *StarsTransaction) FillFrom(from interface {
	GetRefund() (value bool)
	GetPending() (value bool)
	GetFailed() (value bool)
	GetGift() (value bool)
	GetReaction() (value bool)
	GetStargiftUpgrade() (value bool)
	GetBusinessTransfer() (value bool)
	GetStargiftResale() (value bool)
	GetID() (value string)
	GetAmount() (value StarsAmountClass)
	GetDate() (value int)
	GetPeer() (value StarsTransactionPeerClass)
	GetTitle() (value string, ok bool)
	GetDescription() (value string, ok bool)
	GetPhoto() (value WebDocumentClass, ok bool)
	GetTransactionDate() (value int, ok bool)
	GetTransactionURL() (value string, ok bool)
	GetBotPayload() (value []byte, ok bool)
	GetMsgID() (value int, ok bool)
	GetExtendedMedia() (value []MessageMediaClass, ok bool)
	GetSubscriptionPeriod() (value int, ok bool)
	GetGiveawayPostID() (value int, ok bool)
	GetStargift() (value StarGiftClass, ok bool)
	GetFloodskipNumber() (value int, ok bool)
	GetStarrefCommissionPermille() (value int, ok bool)
	GetStarrefPeer() (value PeerClass, ok bool)
	GetStarrefAmount() (value StarsAmountClass, ok bool)
	GetPaidMessages() (value int, ok bool)
	GetPremiumGiftMonths() (value int, ok bool)
	GetAdsProceedsFromDate() (value int, ok bool)
	GetAdsProceedsToDate() (value int, ok bool)
}) {
	s.Refund = from.GetRefund()
	s.Pending = from.GetPending()
	s.Failed = from.GetFailed()
	s.Gift = from.GetGift()
	s.Reaction = from.GetReaction()
	s.StargiftUpgrade = from.GetStargiftUpgrade()
	s.BusinessTransfer = from.GetBusinessTransfer()
	s.StargiftResale = from.GetStargiftResale()
	s.ID = from.GetID()
	s.Amount = from.GetAmount()
	s.Date = from.GetDate()
	s.Peer = from.GetPeer()
	if val, ok := from.GetTitle(); ok {
		s.Title = val
	}

	if val, ok := from.GetDescription(); ok {
		s.Description = val
	}

	if val, ok := from.GetPhoto(); ok {
		s.Photo = val
	}

	if val, ok := from.GetTransactionDate(); ok {
		s.TransactionDate = val
	}

	if val, ok := from.GetTransactionURL(); ok {
		s.TransactionURL = val
	}

	if val, ok := from.GetBotPayload(); ok {
		s.BotPayload = val
	}

	if val, ok := from.GetMsgID(); ok {
		s.MsgID = val
	}

	if val, ok := from.GetExtendedMedia(); ok {
		s.ExtendedMedia = val
	}

	if val, ok := from.GetSubscriptionPeriod(); ok {
		s.SubscriptionPeriod = val
	}

	if val, ok := from.GetGiveawayPostID(); ok {
		s.GiveawayPostID = val
	}

	if val, ok := from.GetStargift(); ok {
		s.Stargift = val
	}

	if val, ok := from.GetFloodskipNumber(); ok {
		s.FloodskipNumber = val
	}

	if val, ok := from.GetStarrefCommissionPermille(); ok {
		s.StarrefCommissionPermille = val
	}

	if val, ok := from.GetStarrefPeer(); ok {
		s.StarrefPeer = val
	}

	if val, ok := from.GetStarrefAmount(); ok {
		s.StarrefAmount = val
	}

	if val, ok := from.GetPaidMessages(); ok {
		s.PaidMessages = val
	}

	if val, ok := from.GetPremiumGiftMonths(); ok {
		s.PremiumGiftMonths = val
	}

	if val, ok := from.GetAdsProceedsFromDate(); ok {
		s.AdsProceedsFromDate = val
	}

	if val, ok := from.GetAdsProceedsToDate(); ok {
		s.AdsProceedsToDate = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarsTransaction) TypeID() uint32 {
	return StarsTransactionTypeID
}

// TypeName returns name of type in TL schema.
func (*StarsTransaction) TypeName() string {
	return "starsTransaction"
}

// TypeInfo returns info about TL type.
func (s *StarsTransaction) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starsTransaction",
		ID:   StarsTransactionTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Refund",
			SchemaName: "refund",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "Pending",
			SchemaName: "pending",
			Null:       !s.Flags.Has(4),
		},
		{
			Name:       "Failed",
			SchemaName: "failed",
			Null:       !s.Flags.Has(6),
		},
		{
			Name:       "Gift",
			SchemaName: "gift",
			Null:       !s.Flags.Has(10),
		},
		{
			Name:       "Reaction",
			SchemaName: "reaction",
			Null:       !s.Flags.Has(11),
		},
		{
			Name:       "StargiftUpgrade",
			SchemaName: "stargift_upgrade",
			Null:       !s.Flags.Has(18),
		},
		{
			Name:       "BusinessTransfer",
			SchemaName: "business_transfer",
			Null:       !s.Flags.Has(21),
		},
		{
			Name:       "StargiftResale",
			SchemaName: "stargift_resale",
			Null:       !s.Flags.Has(22),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Amount",
			SchemaName: "amount",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Description",
			SchemaName: "description",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "TransactionDate",
			SchemaName: "transaction_date",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "TransactionURL",
			SchemaName: "transaction_url",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "BotPayload",
			SchemaName: "bot_payload",
			Null:       !s.Flags.Has(7),
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
			Null:       !s.Flags.Has(8),
		},
		{
			Name:       "ExtendedMedia",
			SchemaName: "extended_media",
			Null:       !s.Flags.Has(9),
		},
		{
			Name:       "SubscriptionPeriod",
			SchemaName: "subscription_period",
			Null:       !s.Flags.Has(12),
		},
		{
			Name:       "GiveawayPostID",
			SchemaName: "giveaway_post_id",
			Null:       !s.Flags.Has(13),
		},
		{
			Name:       "Stargift",
			SchemaName: "stargift",
			Null:       !s.Flags.Has(14),
		},
		{
			Name:       "FloodskipNumber",
			SchemaName: "floodskip_number",
			Null:       !s.Flags.Has(15),
		},
		{
			Name:       "StarrefCommissionPermille",
			SchemaName: "starref_commission_permille",
			Null:       !s.Flags.Has(16),
		},
		{
			Name:       "StarrefPeer",
			SchemaName: "starref_peer",
			Null:       !s.Flags.Has(17),
		},
		{
			Name:       "StarrefAmount",
			SchemaName: "starref_amount",
			Null:       !s.Flags.Has(17),
		},
		{
			Name:       "PaidMessages",
			SchemaName: "paid_messages",
			Null:       !s.Flags.Has(19),
		},
		{
			Name:       "PremiumGiftMonths",
			SchemaName: "premium_gift_months",
			Null:       !s.Flags.Has(20),
		},
		{
			Name:       "AdsProceedsFromDate",
			SchemaName: "ads_proceeds_from_date",
			Null:       !s.Flags.Has(23),
		},
		{
			Name:       "AdsProceedsToDate",
			SchemaName: "ads_proceeds_to_date",
			Null:       !s.Flags.Has(23),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StarsTransaction) SetFlags() {
	if !(s.Refund == false) {
		s.Flags.Set(3)
	}
	if !(s.Pending == false) {
		s.Flags.Set(4)
	}
	if !(s.Failed == false) {
		s.Flags.Set(6)
	}
	if !(s.Gift == false) {
		s.Flags.Set(10)
	}
	if !(s.Reaction == false) {
		s.Flags.Set(11)
	}
	if !(s.StargiftUpgrade == false) {
		s.Flags.Set(18)
	}
	if !(s.BusinessTransfer == false) {
		s.Flags.Set(21)
	}
	if !(s.StargiftResale == false) {
		s.Flags.Set(22)
	}
	if !(s.Title == "") {
		s.Flags.Set(0)
	}
	if !(s.Description == "") {
		s.Flags.Set(1)
	}
	if !(s.Photo == nil) {
		s.Flags.Set(2)
	}
	if !(s.TransactionDate == 0) {
		s.Flags.Set(5)
	}
	if !(s.TransactionURL == "") {
		s.Flags.Set(5)
	}
	if !(s.BotPayload == nil) {
		s.Flags.Set(7)
	}
	if !(s.MsgID == 0) {
		s.Flags.Set(8)
	}
	if !(s.ExtendedMedia == nil) {
		s.Flags.Set(9)
	}
	if !(s.SubscriptionPeriod == 0) {
		s.Flags.Set(12)
	}
	if !(s.GiveawayPostID == 0) {
		s.Flags.Set(13)
	}
	if !(s.Stargift == nil) {
		s.Flags.Set(14)
	}
	if !(s.FloodskipNumber == 0) {
		s.Flags.Set(15)
	}
	if !(s.StarrefCommissionPermille == 0) {
		s.Flags.Set(16)
	}
	if !(s.StarrefPeer == nil) {
		s.Flags.Set(17)
	}
	if !(s.StarrefAmount == nil) {
		s.Flags.Set(17)
	}
	if !(s.PaidMessages == 0) {
		s.Flags.Set(19)
	}
	if !(s.PremiumGiftMonths == 0) {
		s.Flags.Set(20)
	}
	if !(s.AdsProceedsFromDate == 0) {
		s.Flags.Set(23)
	}
	if !(s.AdsProceedsToDate == 0) {
		s.Flags.Set(23)
	}
}

// Encode implements bin.Encoder.
func (s *StarsTransaction) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransaction#13659eb0 as nil")
	}
	b.PutID(StarsTransactionTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarsTransaction) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starsTransaction#13659eb0 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starsTransaction#13659eb0: field flags: %w", err)
	}
	b.PutString(s.ID)
	if s.Amount == nil {
		return fmt.Errorf("unable to encode starsTransaction#13659eb0: field amount is nil")
	}
	if err := s.Amount.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starsTransaction#13659eb0: field amount: %w", err)
	}
	b.PutInt(s.Date)
	if s.Peer == nil {
		return fmt.Errorf("unable to encode starsTransaction#13659eb0: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starsTransaction#13659eb0: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutString(s.Title)
	}
	if s.Flags.Has(1) {
		b.PutString(s.Description)
	}
	if s.Flags.Has(2) {
		if s.Photo == nil {
			return fmt.Errorf("unable to encode starsTransaction#13659eb0: field photo is nil")
		}
		if err := s.Photo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode starsTransaction#13659eb0: field photo: %w", err)
		}
	}
	if s.Flags.Has(5) {
		b.PutInt(s.TransactionDate)
	}
	if s.Flags.Has(5) {
		b.PutString(s.TransactionURL)
	}
	if s.Flags.Has(7) {
		b.PutBytes(s.BotPayload)
	}
	if s.Flags.Has(8) {
		b.PutInt(s.MsgID)
	}
	if s.Flags.Has(9) {
		b.PutVectorHeader(len(s.ExtendedMedia))
		for idx, v := range s.ExtendedMedia {
			if v == nil {
				return fmt.Errorf("unable to encode starsTransaction#13659eb0: field extended_media element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode starsTransaction#13659eb0: field extended_media element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(12) {
		b.PutInt(s.SubscriptionPeriod)
	}
	if s.Flags.Has(13) {
		b.PutInt(s.GiveawayPostID)
	}
	if s.Flags.Has(14) {
		if s.Stargift == nil {
			return fmt.Errorf("unable to encode starsTransaction#13659eb0: field stargift is nil")
		}
		if err := s.Stargift.Encode(b); err != nil {
			return fmt.Errorf("unable to encode starsTransaction#13659eb0: field stargift: %w", err)
		}
	}
	if s.Flags.Has(15) {
		b.PutInt(s.FloodskipNumber)
	}
	if s.Flags.Has(16) {
		b.PutInt(s.StarrefCommissionPermille)
	}
	if s.Flags.Has(17) {
		if s.StarrefPeer == nil {
			return fmt.Errorf("unable to encode starsTransaction#13659eb0: field starref_peer is nil")
		}
		if err := s.StarrefPeer.Encode(b); err != nil {
			return fmt.Errorf("unable to encode starsTransaction#13659eb0: field starref_peer: %w", err)
		}
	}
	if s.Flags.Has(17) {
		if s.StarrefAmount == nil {
			return fmt.Errorf("unable to encode starsTransaction#13659eb0: field starref_amount is nil")
		}
		if err := s.StarrefAmount.Encode(b); err != nil {
			return fmt.Errorf("unable to encode starsTransaction#13659eb0: field starref_amount: %w", err)
		}
	}
	if s.Flags.Has(19) {
		b.PutInt(s.PaidMessages)
	}
	if s.Flags.Has(20) {
		b.PutInt(s.PremiumGiftMonths)
	}
	if s.Flags.Has(23) {
		b.PutInt(s.AdsProceedsFromDate)
	}
	if s.Flags.Has(23) {
		b.PutInt(s.AdsProceedsToDate)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarsTransaction) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransaction#13659eb0 to nil")
	}
	if err := b.ConsumeID(StarsTransactionTypeID); err != nil {
		return fmt.Errorf("unable to decode starsTransaction#13659eb0: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarsTransaction) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starsTransaction#13659eb0 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field flags: %w", err)
		}
	}
	s.Refund = s.Flags.Has(3)
	s.Pending = s.Flags.Has(4)
	s.Failed = s.Flags.Has(6)
	s.Gift = s.Flags.Has(10)
	s.Reaction = s.Flags.Has(11)
	s.StargiftUpgrade = s.Flags.Has(18)
	s.BusinessTransfer = s.Flags.Has(21)
	s.StargiftResale = s.Flags.Has(22)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := DecodeStarsAmount(b)
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field amount: %w", err)
		}
		s.Amount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := DecodeStarsTransactionPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field title: %w", err)
		}
		s.Title = value
	}
	if s.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field description: %w", err)
		}
		s.Description = value
	}
	if s.Flags.Has(2) {
		value, err := DecodeWebDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field photo: %w", err)
		}
		s.Photo = value
	}
	if s.Flags.Has(5) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field transaction_date: %w", err)
		}
		s.TransactionDate = value
	}
	if s.Flags.Has(5) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field transaction_url: %w", err)
		}
		s.TransactionURL = value
	}
	if s.Flags.Has(7) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field bot_payload: %w", err)
		}
		s.BotPayload = value
	}
	if s.Flags.Has(8) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field msg_id: %w", err)
		}
		s.MsgID = value
	}
	if s.Flags.Has(9) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field extended_media: %w", err)
		}

		if headerLen > 0 {
			s.ExtendedMedia = make([]MessageMediaClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageMedia(b)
			if err != nil {
				return fmt.Errorf("unable to decode starsTransaction#13659eb0: field extended_media: %w", err)
			}
			s.ExtendedMedia = append(s.ExtendedMedia, value)
		}
	}
	if s.Flags.Has(12) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field subscription_period: %w", err)
		}
		s.SubscriptionPeriod = value
	}
	if s.Flags.Has(13) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field giveaway_post_id: %w", err)
		}
		s.GiveawayPostID = value
	}
	if s.Flags.Has(14) {
		value, err := DecodeStarGift(b)
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field stargift: %w", err)
		}
		s.Stargift = value
	}
	if s.Flags.Has(15) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field floodskip_number: %w", err)
		}
		s.FloodskipNumber = value
	}
	if s.Flags.Has(16) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field starref_commission_permille: %w", err)
		}
		s.StarrefCommissionPermille = value
	}
	if s.Flags.Has(17) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field starref_peer: %w", err)
		}
		s.StarrefPeer = value
	}
	if s.Flags.Has(17) {
		value, err := DecodeStarsAmount(b)
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field starref_amount: %w", err)
		}
		s.StarrefAmount = value
	}
	if s.Flags.Has(19) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field paid_messages: %w", err)
		}
		s.PaidMessages = value
	}
	if s.Flags.Has(20) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field premium_gift_months: %w", err)
		}
		s.PremiumGiftMonths = value
	}
	if s.Flags.Has(23) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field ads_proceeds_from_date: %w", err)
		}
		s.AdsProceedsFromDate = value
	}
	if s.Flags.Has(23) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starsTransaction#13659eb0: field ads_proceeds_to_date: %w", err)
		}
		s.AdsProceedsToDate = value
	}
	return nil
}

// SetRefund sets value of Refund conditional field.
func (s *StarsTransaction) SetRefund(value bool) {
	if value {
		s.Flags.Set(3)
		s.Refund = true
	} else {
		s.Flags.Unset(3)
		s.Refund = false
	}
}

// GetRefund returns value of Refund conditional field.
func (s *StarsTransaction) GetRefund() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(3)
}

// SetPending sets value of Pending conditional field.
func (s *StarsTransaction) SetPending(value bool) {
	if value {
		s.Flags.Set(4)
		s.Pending = true
	} else {
		s.Flags.Unset(4)
		s.Pending = false
	}
}

// GetPending returns value of Pending conditional field.
func (s *StarsTransaction) GetPending() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(4)
}

// SetFailed sets value of Failed conditional field.
func (s *StarsTransaction) SetFailed(value bool) {
	if value {
		s.Flags.Set(6)
		s.Failed = true
	} else {
		s.Flags.Unset(6)
		s.Failed = false
	}
}

// GetFailed returns value of Failed conditional field.
func (s *StarsTransaction) GetFailed() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(6)
}

// SetGift sets value of Gift conditional field.
func (s *StarsTransaction) SetGift(value bool) {
	if value {
		s.Flags.Set(10)
		s.Gift = true
	} else {
		s.Flags.Unset(10)
		s.Gift = false
	}
}

// GetGift returns value of Gift conditional field.
func (s *StarsTransaction) GetGift() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(10)
}

// SetReaction sets value of Reaction conditional field.
func (s *StarsTransaction) SetReaction(value bool) {
	if value {
		s.Flags.Set(11)
		s.Reaction = true
	} else {
		s.Flags.Unset(11)
		s.Reaction = false
	}
}

// GetReaction returns value of Reaction conditional field.
func (s *StarsTransaction) GetReaction() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(11)
}

// SetStargiftUpgrade sets value of StargiftUpgrade conditional field.
func (s *StarsTransaction) SetStargiftUpgrade(value bool) {
	if value {
		s.Flags.Set(18)
		s.StargiftUpgrade = true
	} else {
		s.Flags.Unset(18)
		s.StargiftUpgrade = false
	}
}

// GetStargiftUpgrade returns value of StargiftUpgrade conditional field.
func (s *StarsTransaction) GetStargiftUpgrade() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(18)
}

// SetBusinessTransfer sets value of BusinessTransfer conditional field.
func (s *StarsTransaction) SetBusinessTransfer(value bool) {
	if value {
		s.Flags.Set(21)
		s.BusinessTransfer = true
	} else {
		s.Flags.Unset(21)
		s.BusinessTransfer = false
	}
}

// GetBusinessTransfer returns value of BusinessTransfer conditional field.
func (s *StarsTransaction) GetBusinessTransfer() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(21)
}

// SetStargiftResale sets value of StargiftResale conditional field.
func (s *StarsTransaction) SetStargiftResale(value bool) {
	if value {
		s.Flags.Set(22)
		s.StargiftResale = true
	} else {
		s.Flags.Unset(22)
		s.StargiftResale = false
	}
}

// GetStargiftResale returns value of StargiftResale conditional field.
func (s *StarsTransaction) GetStargiftResale() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(22)
}

// GetID returns value of ID field.
func (s *StarsTransaction) GetID() (value string) {
	if s == nil {
		return
	}
	return s.ID
}

// GetAmount returns value of Amount field.
func (s *StarsTransaction) GetAmount() (value StarsAmountClass) {
	if s == nil {
		return
	}
	return s.Amount
}

// GetDate returns value of Date field.
func (s *StarsTransaction) GetDate() (value int) {
	if s == nil {
		return
	}
	return s.Date
}

// GetPeer returns value of Peer field.
func (s *StarsTransaction) GetPeer() (value StarsTransactionPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// SetTitle sets value of Title conditional field.
func (s *StarsTransaction) SetTitle(value string) {
	s.Flags.Set(0)
	s.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetTitle() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Title, true
}

// SetDescription sets value of Description conditional field.
func (s *StarsTransaction) SetDescription(value string) {
	s.Flags.Set(1)
	s.Description = value
}

// GetDescription returns value of Description conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetDescription() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.Description, true
}

// SetPhoto sets value of Photo conditional field.
func (s *StarsTransaction) SetPhoto(value WebDocumentClass) {
	s.Flags.Set(2)
	s.Photo = value
}

// GetPhoto returns value of Photo conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetPhoto() (value WebDocumentClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.Photo, true
}

// SetTransactionDate sets value of TransactionDate conditional field.
func (s *StarsTransaction) SetTransactionDate(value int) {
	s.Flags.Set(5)
	s.TransactionDate = value
}

// GetTransactionDate returns value of TransactionDate conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetTransactionDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(5) {
		return value, false
	}
	return s.TransactionDate, true
}

// SetTransactionURL sets value of TransactionURL conditional field.
func (s *StarsTransaction) SetTransactionURL(value string) {
	s.Flags.Set(5)
	s.TransactionURL = value
}

// GetTransactionURL returns value of TransactionURL conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetTransactionURL() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(5) {
		return value, false
	}
	return s.TransactionURL, true
}

// SetBotPayload sets value of BotPayload conditional field.
func (s *StarsTransaction) SetBotPayload(value []byte) {
	s.Flags.Set(7)
	s.BotPayload = value
}

// GetBotPayload returns value of BotPayload conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetBotPayload() (value []byte, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(7) {
		return value, false
	}
	return s.BotPayload, true
}

// SetMsgID sets value of MsgID conditional field.
func (s *StarsTransaction) SetMsgID(value int) {
	s.Flags.Set(8)
	s.MsgID = value
}

// GetMsgID returns value of MsgID conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetMsgID() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(8) {
		return value, false
	}
	return s.MsgID, true
}

// SetExtendedMedia sets value of ExtendedMedia conditional field.
func (s *StarsTransaction) SetExtendedMedia(value []MessageMediaClass) {
	s.Flags.Set(9)
	s.ExtendedMedia = value
}

// GetExtendedMedia returns value of ExtendedMedia conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetExtendedMedia() (value []MessageMediaClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(9) {
		return value, false
	}
	return s.ExtendedMedia, true
}

// SetSubscriptionPeriod sets value of SubscriptionPeriod conditional field.
func (s *StarsTransaction) SetSubscriptionPeriod(value int) {
	s.Flags.Set(12)
	s.SubscriptionPeriod = value
}

// GetSubscriptionPeriod returns value of SubscriptionPeriod conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetSubscriptionPeriod() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(12) {
		return value, false
	}
	return s.SubscriptionPeriod, true
}

// SetGiveawayPostID sets value of GiveawayPostID conditional field.
func (s *StarsTransaction) SetGiveawayPostID(value int) {
	s.Flags.Set(13)
	s.GiveawayPostID = value
}

// GetGiveawayPostID returns value of GiveawayPostID conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetGiveawayPostID() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(13) {
		return value, false
	}
	return s.GiveawayPostID, true
}

// SetStargift sets value of Stargift conditional field.
func (s *StarsTransaction) SetStargift(value StarGiftClass) {
	s.Flags.Set(14)
	s.Stargift = value
}

// GetStargift returns value of Stargift conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetStargift() (value StarGiftClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(14) {
		return value, false
	}
	return s.Stargift, true
}

// SetFloodskipNumber sets value of FloodskipNumber conditional field.
func (s *StarsTransaction) SetFloodskipNumber(value int) {
	s.Flags.Set(15)
	s.FloodskipNumber = value
}

// GetFloodskipNumber returns value of FloodskipNumber conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetFloodskipNumber() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(15) {
		return value, false
	}
	return s.FloodskipNumber, true
}

// SetStarrefCommissionPermille sets value of StarrefCommissionPermille conditional field.
func (s *StarsTransaction) SetStarrefCommissionPermille(value int) {
	s.Flags.Set(16)
	s.StarrefCommissionPermille = value
}

// GetStarrefCommissionPermille returns value of StarrefCommissionPermille conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetStarrefCommissionPermille() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(16) {
		return value, false
	}
	return s.StarrefCommissionPermille, true
}

// SetStarrefPeer sets value of StarrefPeer conditional field.
func (s *StarsTransaction) SetStarrefPeer(value PeerClass) {
	s.Flags.Set(17)
	s.StarrefPeer = value
}

// GetStarrefPeer returns value of StarrefPeer conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetStarrefPeer() (value PeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(17) {
		return value, false
	}
	return s.StarrefPeer, true
}

// SetStarrefAmount sets value of StarrefAmount conditional field.
func (s *StarsTransaction) SetStarrefAmount(value StarsAmountClass) {
	s.Flags.Set(17)
	s.StarrefAmount = value
}

// GetStarrefAmount returns value of StarrefAmount conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetStarrefAmount() (value StarsAmountClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(17) {
		return value, false
	}
	return s.StarrefAmount, true
}

// SetPaidMessages sets value of PaidMessages conditional field.
func (s *StarsTransaction) SetPaidMessages(value int) {
	s.Flags.Set(19)
	s.PaidMessages = value
}

// GetPaidMessages returns value of PaidMessages conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetPaidMessages() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(19) {
		return value, false
	}
	return s.PaidMessages, true
}

// SetPremiumGiftMonths sets value of PremiumGiftMonths conditional field.
func (s *StarsTransaction) SetPremiumGiftMonths(value int) {
	s.Flags.Set(20)
	s.PremiumGiftMonths = value
}

// GetPremiumGiftMonths returns value of PremiumGiftMonths conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetPremiumGiftMonths() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(20) {
		return value, false
	}
	return s.PremiumGiftMonths, true
}

// SetAdsProceedsFromDate sets value of AdsProceedsFromDate conditional field.
func (s *StarsTransaction) SetAdsProceedsFromDate(value int) {
	s.Flags.Set(23)
	s.AdsProceedsFromDate = value
}

// GetAdsProceedsFromDate returns value of AdsProceedsFromDate conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetAdsProceedsFromDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(23) {
		return value, false
	}
	return s.AdsProceedsFromDate, true
}

// SetAdsProceedsToDate sets value of AdsProceedsToDate conditional field.
func (s *StarsTransaction) SetAdsProceedsToDate(value int) {
	s.Flags.Set(23)
	s.AdsProceedsToDate = value
}

// GetAdsProceedsToDate returns value of AdsProceedsToDate conditional field and
// boolean which is true if field was set.
func (s *StarsTransaction) GetAdsProceedsToDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(23) {
		return value, false
	}
	return s.AdsProceedsToDate, true
}

// MapExtendedMedia returns field ExtendedMedia wrapped in MessageMediaClassArray helper.
func (s *StarsTransaction) MapExtendedMedia() (value MessageMediaClassArray, ok bool) {
	if !s.Flags.Has(9) {
		return value, false
	}
	return MessageMediaClassArray(s.ExtendedMedia), true
}
