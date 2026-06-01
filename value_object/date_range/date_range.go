package date_range

import (
	"errors"
	"time"
)

type DateRange struct {
	start time.Time
	end   time.Time
}

var ErrInvalidDate = errors.New("invalid time")
var ErrMismatchTimezone = errors.New("timezones mismatch")
var ErrInvalidDateRange = errors.New("invalid date range params")

func NewDateRange(start, end time.Time) (DateRange, error) {
	if start.IsZero() || end.IsZero() {
		return DateRange{}, ErrInvalidDate
	}

	startZoneName, startOffset := start.Zone()
	endZoneName, endOffset := start.Zone()

	if (startZoneName != endZoneName) || (startOffset != endOffset) {
		return DateRange{}, ErrMismatchTimezone
	}

	// 检查是否是一个有效区间
	if start.Equal(end) {
		return DateRange{}, ErrInvalidDateRange
	}
	if start.After(end) || end.Before(start) {
		return DateRange{}, ErrInvalidDateRange
	}

	return DateRange{
		start: start,
		end:   end,
	}, nil
}

// WithinRange 检查目标区间 target 是否在区间内
func (dr DateRange) WithinRange(target DateRange) bool {
	var (
		l bool = target.start.After(dr.start) || target.start.Equal(dr.start)
		r bool = target.end.Before(dr.end) || target.end.Equal(dr.end)
	)

	return l && r
}

func (dr DateRange) Difference() time.Duration {
	return dr.end.Sub(dr.start)
}
