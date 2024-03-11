package gox

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// Date masks a [time.Time] to only marshal and unmarshal the date portion. For
// strings it chooses the ISO/RFC format of "2006-01-02"
type Date struct {
	value time.Time
}

func (d Date) Time() *time.Time {
	return &d.value
}

func (d Date) String() string {
	return d.value.Format(time.DateOnly)
}

// Parse accepts a string that is RFC3339 date-only such as "2006-01-02" and
// parses it into this value. If an error occurs it is returned.
func (d *Date) Parse(str string) (err error) {
	d.value, err = time.Parse(time.DateOnly, str)
	return
}

func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) UnmarshalText(src []byte) error {
	return d.Parse(string(src))
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Date) UnmarshalJSON(src []byte) error {
	if len(src) == 0 {
		d.value = TimeZero
		return nil
	} else if src[0] == '"' {
		return d.Parse(string(src[1 : len(src)-1]))
	}
	return errors.New("unknown format for JSON unmarshaling of Date")
}

func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

func (d *Date) Scan(src any) error {
	if src == nil {
		d.value = TimeZero
		return nil
	}

	if str, ok := src.(string); ok {
		return d.Parse(str)
	} else if byts, ok := src.([]byte); ok {
		return d.Parse(string(byts))
	} else if tm, ok := src.(time.Time); ok {
		d.value = tm
		return nil
	}

	return fmt.Errorf("failed to scan %T as Date", src)
}

func (d Date) Year() int {
	return d.value.Year()
}

func (d Date) Month() int {
	return int(d.value.Month())
}

func (d Date) Day() int {
	return d.value.Day()
}

func (d Date) Add(dur time.Duration) Date {
	return Date{d.value.Add(dur)}
}

func (d Date) Equal(other Date) bool {
	return d.Year() == other.Year() &&
		d.Month() == other.Month() &&
		d.Day() == other.Day()
}

// Compare returns -1 if THIS date is before the given other, 1 if THIS date is
// AFTER the given other, and 0 if they are equal.
func (d Date) Compare(other Date) (dif int) {
	if y := d.Year() - other.Year(); y != 0 {
		return Ternary(y < 0, -1, 1)
	}
	if m := d.Month() - other.Month(); m != 0 {
		return Ternary(m < 0, -1, 1)
	}
	if d := d.Day() - other.Day(); d != 0 {
		return Ternary(d < 0, -1, 1)
	}
	return 0
}

// Before returns true if THIS date is before the given other.
func (d Date) Before(other Date) bool {
	return d.Compare(other) == -1
}

// After returns true if THIS date is after the given other.
func (d Date) After(other Date) bool {
	return d.Compare(other) == 1
}

// DurationBetween may or may not work. It subtracts the other date from this
// date and truncates the resulting duration to 24 hours.
func (d Date) DurationBetween(other Date) time.Duration {
	dur := d.value.Sub(other.value)
	return dur.Truncate(time.Hour * 24)
}

// Age returns the number of years since the given date, probably.
func (d Date) Age() int {
	return int(time.Since(d.value).Truncate(time.Hour*24)) / (24 * 365)
}

// IsZero returns true if this Date is a zero-value.
func (d Date) IsZero() bool {
	return d.value.Equal(TimeZero)
}

// DateFromTime returns a new [Date] object using the given time
func DateFromTime(t time.Time) Date {
	return Date{t}
}

// DateNow returns a new [Date] object set to now via [time.Now].
func DateNow() Date {
	return Date{time.Now()}
}

// ParseDate accepts a string that is RFC3339 date-only such as "2006-01-02" and
// parses it into a new [Date] object.
func ParseDate(str string) (Date, error) {
	t, err := time.Parse(time.DateOnly, str)
	if err != nil {
		return Date{}, err
	}
	return Date{t}, nil
}
