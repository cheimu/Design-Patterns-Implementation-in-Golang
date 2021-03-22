package factory

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash int = iota
	DebitCard
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitPM), nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized", m)
	}
}

type CashPM struct{}
type DebitPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using cash\n", amount)
}

func (d *DebitPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using debit card\n", amount)
}
