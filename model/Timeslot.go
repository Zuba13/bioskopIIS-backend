package model

import (
	"time"

	"gorm.io/gorm"
)

type Timeslot struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

func (t *Timeslot) BeforeSave(tx *gorm.DB) error {
	err := t.ToCET()
	if err != nil {
		return err
	}
	return nil
}

func (t *Timeslot) ToCET() error {
	loc, err := time.LoadLocation("CET")
	if err != nil {
		return err
	}
	t.StartTime = t.StartTime.In(loc)
	t.EndTime = t.EndTime.In(loc)
	return nil
}

func NewTimeslot(startTime time.Time, duration int) Timeslot {
	minute := startTime.Minute()
	remainder := minute % 15
	if remainder > 0 {
		minute = minute - remainder + 15
	}
	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(), minute, 0, 0, startTime.Location())

	endTime := startTime.Add(time.Duration(duration) * time.Minute)

	minute = endTime.Minute()
	remainder = minute % 15
	if remainder > 0 {
		minute = minute - remainder + 15
	}
	endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), endTime.Hour(), minute, 0, 0, endTime.Location())

	return Timeslot{
		StartTime: startTime,
		EndTime:   endTime,
	}
}

func (t Timeslot) IsValid() bool {
	if t.StartTime.After(t.EndTime) {
		return false
	}
	if t.StartTime.Minute()%15 != 0 || t.EndTime.Minute()%15 != 0 {
		return false
	}
	return true
}

func (timeslotA Timeslot) Overlaps(timeslotB Timeslot, minsInBetween uint16) bool {
	gap := time.Duration(minsInBetween) * time.Minute
	return (timeslotB.StartTime.Before(timeslotA.EndTime.Add(gap)) && timeslotB.EndTime.After(timeslotA.StartTime.Add(gap)) && !timeslotB.EndTime.Equal(timeslotA.StartTime.Add(gap))) ||
		(timeslotA.StartTime.Before(timeslotB.EndTime.Add(gap)) && timeslotA.EndTime.After(timeslotB.StartTime.Add(gap)) && !timeslotA.EndTime.Equal(timeslotB.StartTime.Add(gap)))
}

func (t Timeslot) Equals(other Timeslot) bool {
	return t.StartTime.Equal(other.StartTime) && t.EndTime.Equal(other.EndTime)
}

func (t Timeslot) StartsAfter(other Timeslot) bool {
	return t.StartTime.After(other.StartTime)
}

func (t Timeslot) AddMinutes(minutes int16) Timeslot {
	startTime := t.StartTime.Add(time.Duration(minutes) * time.Minute)
	endTime := t.EndTime.Add(time.Duration(minutes) * time.Minute)
	return Timeslot{
		StartTime: startTime,
		EndTime:   endTime,
	}
}

func (t Timeslot) AddDays(days int) Timeslot {
	startTime := t.StartTime.AddDate(0, 0, days)
	endTime := t.EndTime.AddDate(0, 0, days)
	return Timeslot{
		StartTime: startTime,
		EndTime:   endTime,
	}
}
