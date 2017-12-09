package tymee

import "time"

type Datetime struct {
	Year       uint16
	Month      uint8
	Day        uint8
	Hour       uint8
	Minute     uint8
	Second     uint8
	Nanosecond uint32
}

func Now() Datetime {
	t := time.Now()

	return Datetime{
		Year:       uint16(t.Year()),
		Month:      uint8(t.Month()),
		Day:        uint8(t.Day()),
		Hour:       uint8(t.Hour()),
		Minute:     uint8(t.Minute()),
		Second:     uint8(t.Second()),
		Nanosecond: uint32(t.Nanosecond()),
	}
}
