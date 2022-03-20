package krendel

type Int64 int64

// Separator can be carefully used to set separator character.
var Separator byte = ' '

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
	if v < 0 {
		v = 0 - v
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
			output[currentIndex] = Separator
			currentIndex--
		} else {
			counter++
		}
	}

	output[currentIndex] = byte(v) + 48
	if currentIndex == 1 {
		output[0] = '-'
	}
	return string(output)
}
