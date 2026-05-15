package value_object

import "time"

type PasswordExpiryAt time.Time

func NewPasswordExpiryAtFromTime(t time.Time) PasswordExpiryAt {
	return PasswordExpiryAt(t)
}

func (p PasswordExpiryAt) IsExpired() bool {
	return time.Now().After(time.Time(p))
}

func (p PasswordExpiryAt) gap() time.Duration {
	target := time.Time(p)
	now := time.Now()

	return target.Sub(now)
}

func (p PasswordExpiryAt) Days() int {
	return int(p.Hours() / 24)
}

func (p PasswordExpiryAt) Hours() int {
	return int(p.gap().Hours())
}
