package model

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Date struct {
	time.Time
}

func NewDate(t time.Time) Date {
	return Date{t}
}

func (d Date) After(other Date) bool {
	return d.Year() > other.Year() ||
		(d.Year() == other.Year() && d.Month() > other.Month()) ||
		(d.Year() == other.Year() && d.Month() == other.Month() && d.Day() > other.Day())
}

func (d Date) Before(other Date) bool {
	return d.Year() < other.Year() ||
		(d.Year() == other.Year() && d.Month() < other.Month()) ||
		(d.Year() == other.Year() && d.Month() == other.Month() && d.Day() < other.Day())
}

func (d Date) Equals(other Date) bool {
	return d.Year() == other.Year() && d.Month() == other.Month() && d.Day() == other.Day()
}

// Value implements the driver.Valuer interface
func (d Date) Value() (driver.Value, error) {
	return d.Format("2006-01-02"), nil
}

// Scan implements the sql.Scanner interface
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date{Time: time.Time{}}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = Date{Time: v}
		return nil
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		*d = Date{Time: t}
		return nil
	default:
		return errors.New("failed to scan Date")
	}
}
