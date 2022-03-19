package krendel

import "log"

func getSliceInt64[I ~int64](num I) []uint16 {
	if num == 0 {
		return []uint16{0}
	}

	// Counting the number of digits.
	var count byte
	for n := num; n != 0; n = n / 10 {
		count++
	}
	log.Println("number of digits is", count)

	var output []uint16
	var index = byte(0)
	var size = (count-1)/3 + 1

	output = make([]uint16, size)
	index = size - 1

	log.Println("len", len(output))
	log.Println("last index", index)

	var n I

	count = 0
	for n = num; n != 0; {
		log.Println("==== loop ====")
		log.Println("index", index)
		log.Println("count", count)
		if num < 1000 {
			output[index] = uint16(num)
			break
		}
		n = n / 10
		log.Println("n", n)
		if count != 0 && (count+1)%3 == 0 {
			log.Println("> count%3", count%3)
			log.Println("num", num)
			log.Println("item in index", index, "is", num-(n*1000))

			output[index] = uint16(num - (n * 1000))
			// We finished processing all the digits.
			if index == 0 {
				break
			}
			num = num / 1000
			index--
		}

		count++
	}

	log.Println(output)
	return output
}
