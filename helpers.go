package krendel

// Digits count the number of digits in a number.
func Digits(n Int64) byte {
	var count byte
	for ; n != 0; n = n / 10 {
		count++
	}
	//log.Println("number of digits is", count)
	return count
}
