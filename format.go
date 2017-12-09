package tymee

import "errors"

const (
	_ = iota
	shortYearComponent
	longYearComponent
	monthComponent
	dayComponent
	hourComponent
	minuteComponent
	secondComponent
	nanosecondComponent
)

func nextComponent(format string) (prefix string, component int, suffix string, err error) {
	for i := 0; i < len(format); i++ {
		if format[i] == '%' {
			if i == len(format)-1 {
				return "", 0, "", errors.New("format ends with single %")
			}

			switch format[i+1] {
			case 'H':
				return format[0:i], hourComponent, format[i+2:], nil
			case 'M':
				return format[0:i], minuteComponent, format[i+2:], nil
			case 'S':
				return format[0:i], secondComponent, format[i+2:], nil
			default:
				return format[0:i], -1, format[i+1:], nil
			}
		}
	}
	return format, 0, "", nil
}

func appendInt(b []byte, x, width int, fill byte) []byte {
	u := uint(x)
	if x < 0 {
		b = append(b, '-')
		u = uint(-x)
	}

	// Assemble decimal in reverse order.
	var buf [20]byte
	i := len(buf)
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add padding.
	for w := len(buf) - i; w < width; w++ {
		b = append(b, fill)
	}

	return append(b, buf[i:]...)
}

func (d *Datetime) Format(format string) string {
	var b []byte

	for format != "" {
		prefix, component, suffix, err := nextComponent(format)
		if err != nil {
			panic(err)
		}
		if prefix != "" {
			b = append(b, prefix...)
		}
		if component == 0 {
			break
		}
		format = suffix

		switch component {
		case hourComponent:
			b = appendInt(b, int(d.Hour), 2, '0')
		}
	}

	return string(b)
}

func (d *Datetime) String() string {
	return d.Format("%Y-%m-%d %H:%M:%S.%f")
}
