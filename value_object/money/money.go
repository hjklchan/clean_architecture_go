package money

import (
	"errors"
	"fmt"
	"math"
)

type Money struct {
	currency string
	amount   float64
}

var (
	ErrEmptyUnit     = errors.New("currency can't be empty")
	ErrNegativeValue = errors.New("amount can't be negative")
	ErrUnitMismatch  = errors.New("currency don't match to target's")
)

func NewMoney(currency string, amount float64) (Money, error) {
	if currency == "" {
		return Money{}, nil
	}
	if amount < 0 {
		return Money{}, ErrNegativeValue
	}

	return Money{
		currency: currency,
		amount:   amount,
	}, nil
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Amount() float64 {
	return m.amount
}

func (m Money) Text() string {
	return fmt.Sprintf("%s %f", m.Currency(), m.Amount())
}

func (m Money) Gap(target Money) (float64, error) {
	if m.Currency() != target.Currency() {
		return math.NaN(), ErrUnitMismatch
	}

	return math.Abs(target.Amount() - m.Amount()), nil
}

func (m Money) GapWithoutCurrency(target Money) float64 {
	return math.Abs(m.Amount() - target.Amount())
}

func (p Money) EqualTo(target Money) bool {
	return p.Currency() == target.Currency() && p.Amount() == target.Amount()
}
