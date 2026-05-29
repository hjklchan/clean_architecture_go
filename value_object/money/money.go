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
	ErrEmptyUnit         = errors.New("currency can't be empty")
	ErrNegativeValue     = errors.New("amount can't be negative")
	ErrCurrencyMismatch  = errors.New("currency don't match to target's")
	ErrInsufficientFunds = errors.New("insufficient funds")
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

// Deposit 充值金额
func (m Money) Deposit(value Money) (Money, error) {
	if m.currency != value.currency {
		return Money{}, ErrCurrencyMismatch
	}

	return Money{
		currency: m.currency,
		amount:   m.amount + value.amount,
	}, nil
}

// Withdraw 提现金额
func (m Money) Withdraw(value Money) (Money, error) {
	if m.Currency() != value.Currency() {
		return Money{}, ErrCurrencyMismatch
	}
	if m.Amount() < value.Amount() {
		return Money{}, ErrInsufficientFunds
	}

	return Money{
		currency: m.Currency(),
		amount:   m.Amount() - value.Amount(),
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
		return math.NaN(), ErrCurrencyMismatch
	}

	return math.Abs(target.Amount() - m.Amount()), nil
}

func (m Money) GapWithoutCurrency(target Money) float64 {
	return math.Abs(m.Amount() - target.Amount())
}

func (p Money) EqualTo(target Money) bool {
	return p.Currency() == target.Currency() && p.Amount() == target.Amount()
}
