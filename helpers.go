package krendel

import (
	"github.com/kaatinga/assets"
	"log"
)

func getSliceInt64[I ~int64](number I) [][]byte {
	if number == 0 {
		return [][]byte{{48}}
	}
	// Counting the number of digits.
	var count byte
	for n := number; n != 0; n = n / 10 {
		count++
	}
	log.Println("number of digits is", count)

	var output [][]byte
	var index = byte(0)
	var size = (count-1)/3 + 1

	output = make([][]byte, size)
	index = size - 1

	log.Println("len", len(output))
	log.Println("last index", index)

	for number > 999 {
		output[index] = assets.Uint162Bytes(uint16(number % divider))
		switch len(output[index]) {
		case 1:
			output[index] = append([]byte{48, 48}, output[index]...)
		case 2:
			output[index] = append([]byte{48}, output[index]...)
		}
		number = number / divider
		index--
	}
	log.Println("last number", number)
	log.Println("index", index)
	output[index] = assets.Uint162Bytes(uint16(number % divider))
	log.Println(output)
	return output
}
