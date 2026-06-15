package utils

import (
	"fmt"
	"log"
	"time"
)

const (
	DurationUnitDay   DurationUnit = "day"
	DurationUnitMonth DurationUnit = "month"
	DurationUnitYear  DurationUnit = "year"
)

type DurationUnit string

type Duration struct {
	Amount int          `json:"amount"`
	Unit   DurationUnit `json:"unit"`
}

func (d *Duration) AddToTime(t time.Time) (*time.Time, error) {
	var newTime time.Time

	switch d.Unit {
	case DurationUnitDay:
		newTime = t.AddDate(0, 0, d.Amount)
	case DurationUnitMonth:
		newTime = t.AddDate(0, d.Amount, 0)
	case DurationUnitYear:
		newTime = t.AddDate(d.Amount, 0, 0)
	default:
		return nil, fmt.Errorf("invalid duration unit")
	}

	return &newTime, nil
}

func (d *Duration) Abbreviation() string {
	switch d.Unit {
	case DurationUnitDay:
		return fmt.Sprintf("%dd", d.Amount)
	case DurationUnitMonth:
		return fmt.Sprintf("%dm", d.Amount)
	case DurationUnitYear:
		return fmt.Sprintf("%dy", d.Amount)
	}

	log.Panicln("unsupported duration unit for abbreviation")
	return ""
}
