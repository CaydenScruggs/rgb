package models

import (
	"fmt"
	"strings"
	"time"
)

// Use a custom type that unmarshals and validates in one step
type Date struct {
	time.Time
}

const expectedLayout = "2006-01-02" // Go's reference time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	t, err := time.Parse(expectedLayout, s)
	if err != nil {
		return fmt.Errorf("date must be in YYYY-MM-DD format, got: %s", s)
	}
	d.Time = t
	return nil
}
