package krendel

type Int64 int64

func (v Int64) String() string {
	if v&^0b111 == 0 {
		return string([]byte{byte(v) + 48})
	}
	if v == -9223372036854775808 {
		return "-9,223,372,036,854,775,808"
	}
	// Counting the number of digits.
	count := Digits(v)

	count += (count - 1) / 3
	var minus bool
	if v < 0 {
		v = 0 - v
		minus = true
		count++
	}
	output := make([]byte, count)
	currentIndex := byte(len(output) - 1)

	var counter byte = 0
	for v > 9 {
		output[currentIndex] = byte(v%10) + 48
		v = v / 10
		currentIndex--
		if counter == 2 {
			counter = 0
			output[currentIndex] = ','
			currentIndex--
		} else {
			counter++
		}
	}

	output[currentIndex] = byte(v) + 48
	if minus {
		output[0] = '-'
	}
	return string(output)
}
