package date_range

import (
	"errors"
	"time"
)

type DateRange struct {
	from time.Time
	to   time.Time
}

var ErrInvalidDate = errors.New("invalid time")
var ErrMismatchTimezone = errors.New("timezones mismatch")

func NewDateRange(from, to time.Time) (DateRange, error) {
	if from.IsZero() || to.IsZero() {
		return DateRange{}, ErrInvalidDate
	}

	fromZoneName, fromOffset := from.Zone()
	toZoneName, toOffset := from.Zone()

	if (fromZoneName != toZoneName) || (fromOffset == toOffset) {
		return DateRange{}, ErrMismatchTimezone
	}

	return DateRange{
		from: from,
		to:   to,
	}, nil
}

func (dr DateRange) Diff() time.Duration {
	return dr.from.Sub(dr.to).Abs()
}

// TODO
//
// 检查目标区间 target 是否在区间内
func (dr DateRange) WithinRange(target DateRange) bool {

	return true
}
