package money

type Amount int64

func FromMilli(v int64) Amount {
	return Amount(v)
}

func FromCenti(v int64) Amount {
	return Amount(v * 10)
}

func (a Amount) AsMilli() int64 {
	return int64(a)
}

func (a Amount) AsCenti() int64 {
	return int64(a / 10)
}

func (a Amount) AsDeci() int64 {
	return int64(a / 100)
}

func (a Amount) AsInt() int64 {
	return int64(a / 1000)
}

func (a *Amount) FromMilli(v int64) {
	*a = FromMilli(v)
}

func (a *Amount) FromCenti(v int64) {
	*a = FromCenti(v)
}

type Formatter struct {
	Thousand rune
	Decimal  rune
}

func toRune(v int64) rune {
	return '0' + rune(v)
}

func (f Formatter) Sprint(a Amount) string {
	v := a.AsCenti()

	// Get abs for converting to ascii
	negative := v < 0
	if negative {
		v = v * -1
	}

	// Convert chars
	values := make([]rune, 0, 10)
	for v > 0 {
		c := toRune(v % 10)
		v = v / 10
		values = append(values, c)
	}

	// Calculate buffer length
	lv := len(values)
	bufferLen := 1 + (lv/3)*4 + lv%3
	if bufferLen < 5 {
		bufferLen = 5
	}

	buffer := make([]rune, bufferLen)

	// Prefill negative
	buffer[0] = '-'

	// Prefill the minor values
	buffer[bufferLen-1] = '0'
	buffer[bufferLen-2] = '0'
	buffer[bufferLen-3] = f.Decimal
	buffer[bufferLen-4] = '0'

	pos := bufferLen - 1
	for i, ch := range values {
		// Jump decimal
		if i == 2 {
			pos = pos - 1
		}

		// Add thousand
		if i > 3 && i%3 == 2 {
			buffer[pos] = f.Thousand
			pos = pos - 1
		}

		// Insert char
		buffer[pos] = ch
		pos = pos - 1
	}

	// Remove negative if positive
	if !negative {
		buffer = buffer[1:]
	}
	return string(buffer)
}
