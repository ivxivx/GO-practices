package json

import (
	"fmt"
	"testing"

	gj "github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

type EventType string

const (
	EventTypePayment EventType = "payment"
	EventTypePayout  EventType = "payout"
)

type PaymentMethodType string

const (
	PaymentMethodTypeCard        PaymentMethodType = "card"
	PaymentMethodTypeBankAccount PaymentMethodType = "bank_account"
)

type PaymentPayload struct {
	PaymentMethodType PaymentMethodType `json:"payment_method_type"`
	Amount            string            `json:"amount"`
	Currency          string            `json:"currency"`
	Source            string            `json:"source"`
	Details           any               `json:"-"` // can be CardPaymentDetails or BankPaymentDetails
	RawDetails        gj.RawMessage     `json:"details,omitempty"`
}

type PayoutPayload struct {
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Destination string `json:"destination"`
}

type CardPaymentDetails struct {
	CardNumber string `json:"card_number"`
}

type BankPaymentDetails struct {
	BankAccountNumber string `json:"bank_account_number"`
}

func (p *PaymentPayload) MarshalJSON() ([]byte, error) {
	details, errM := gj.Marshal(p.Details)
	if errM != nil {
		return nil, errM
	}

	type Alias PaymentPayload

	payload := &Alias{
		PaymentMethodType: p.PaymentMethodType,
		Amount:            p.Amount,
		Currency:          p.Currency,
		Source:            p.Source,
		Details:           p.Details,
		RawDetails:        details,
	}

	marshaledPayload, errM := gj.Marshal(payload)

	return marshaledPayload, errM
}

func (p *PaymentPayload) UnmarshalJSON(data []byte) error {
	type Alias PaymentPayload

	var tmp Alias

	if errW := gj.Unmarshal(data, &tmp); errW != nil {
		return UnmarshalJsonTypeError{
			Data: string(data),
			Err:  errW,
		}
	}

	switch tmp.PaymentMethodType {
	case PaymentMethodTypeCard:
		var details CardPaymentDetails
		if errD := gj.Unmarshal(tmp.RawDetails, &details); errD != nil {
			return UnmarshalJsonTypeError{
				Data: string(data),
				Err:  errD,
			}
		}

		p.PaymentMethodType = tmp.PaymentMethodType
		p.Amount = tmp.Amount
		p.Currency = tmp.Currency
		p.Source = tmp.Source
		p.Details = &details

		return nil
	case PaymentMethodTypeBankAccount:
		var details BankPaymentDetails
		if errD := gj.Unmarshal(tmp.RawDetails, &details); errD != nil {
			return UnmarshalJsonTypeError{
				Data: string(data),
				Err:  errD,
			}
		}

		p.PaymentMethodType = tmp.PaymentMethodType
		p.Amount = tmp.Amount
		p.Currency = tmp.Currency
		p.Source = tmp.Source
		p.Details = &details

		return nil
	default:
	}

	return fmt.Errorf("unknown payment method type: %s", tmp.PaymentMethodType)
}

type Event struct {
	Type       EventType     `json:"type,omitempty"`
	Payload    any           `json:"-"` // can be PaymentPayload or PayoutPayload
	RawPayload gj.RawMessage `json:"payload,omitempty"`
}

func (p *Event) MarshalJSON() ([]byte, error) {
	payload, errM := gj.Marshal(p.Payload)

	if errM != nil {
		return nil, errM
	}

	type Alias Event

	event := &Alias{
		Type:       p.Type,
		Payload:    p.Payload,
		RawPayload: payload,
	}

	marshaledEvent, errM := gj.Marshal(event)

	return marshaledEvent, errM
}

func (p *Event) UnmarshalJSON(data []byte) error {
	type Alias Event

	var tmp Alias

	if errW := gj.Unmarshal(data, &tmp); errW != nil {
		return UnmarshalJsonTypeError{
			Data: string(data),
			Err:  errW,
		}
	}

	switch tmp.Type {
	case EventTypePayment:
		var paymentPayload PaymentPayload
		if errD := gj.Unmarshal(tmp.RawPayload, &paymentPayload); errD != nil {
			return UnmarshalJsonTypeError{
				Data: string(data),
				Err:  errD,
			}
		}

		p.Type = tmp.Type
		p.Payload = &paymentPayload

		return nil
	case EventTypePayout:
		var payoutPayload PayoutPayload
		if errD := gj.Unmarshal(tmp.RawPayload, &payoutPayload); errD != nil {
			return UnmarshalJsonTypeError{
				Data: string(data),
				Err:  errD,
			}
		}

		p.Type = tmp.Type
		p.Payload = &payoutPayload

		return nil
	default:
	}

	return fmt.Errorf("unknown event type: %s", tmp.Type)
}

func Test_Nested_Polymorphism(t *testing.T) {
	t.Parallel()

	cardPaymentEvent := &Event{
		Type: EventTypePayment,
		Payload: &PaymentPayload{
			PaymentMethodType: PaymentMethodTypeCard,
			Amount:            "1.23",
			Currency:          "ZAR",
			Source:            "card1",
			Details: &CardPaymentDetails{
				CardNumber: "1234",
			},
		},
	}

	bankAccountPaymentEvent := &Event{
		Type: EventTypePayment,
		Payload: &PaymentPayload{
			PaymentMethodType: PaymentMethodTypeBankAccount,
			Amount:            "3.21",
			Currency:          "VUV",
			Source:            "account1",
			Details: &BankPaymentDetails{
				BankAccountNumber: "4321",
			},
		},
	}

	payoutEvent := &Event{
		Type: EventTypePayout,
		Payload: &PayoutPayload{
			Amount:      "0.01",
			Currency:    "USD",
			Destination: "dest1",
		},
	}

	marshaledCardPaymentEvent, err := gj.Marshal(cardPaymentEvent)
	if err != nil {
		t.Fatal(err)
	}

	var unmarshaledCardPaymentEvent Event
	if errU := gj.Unmarshal(marshaledCardPaymentEvent, &unmarshaledCardPaymentEvent); errU != nil {
		t.Fatal(errU)
	}

	assert.Equal(t, cardPaymentEvent, &unmarshaledCardPaymentEvent)

	marshaledBankAccountPaymentEvent, err := gj.Marshal(bankAccountPaymentEvent)
	if err != nil {
		t.Fatal(err)
	}

	var unmarshaledBankAccountPaymentEvent Event
	if errU := gj.Unmarshal(marshaledBankAccountPaymentEvent, &unmarshaledBankAccountPaymentEvent); errU != nil {
		t.Fatal(errU)
	}

	assert.Equal(t, bankAccountPaymentEvent, &unmarshaledBankAccountPaymentEvent)

	marshaledPayoutEvent, err := gj.Marshal(payoutEvent)
	if err != nil {
		t.Fatal(err)
	}

	var unmarshaledPayoutEvent Event
	if errU := gj.Unmarshal(marshaledPayoutEvent, &unmarshaledPayoutEvent); errU != nil {
		t.Fatal(errU)
	}

	assert.Equal(t, payoutEvent, &unmarshaledPayoutEvent)
}
