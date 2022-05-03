package krendel

import (
	"strconv"
)

// digits count the number of digits in a number.
func (v Int64) digits() byte {
	var count byte
	for ; v != 0; v = v / 10 {
		count++
	}
	return count
}

// SetString parses and sets an int32 value.
func (v *Int64) SetString(stringValue string) error {
	intValue, err := strconv.ParseInt(stringValue, 10, 32)
	if err != nil {
		return err
	}
	*v = Int64(intValue)
	return nil
}
