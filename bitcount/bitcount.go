//Write a function that takes an integer as input, and returns the number
//of bits that are equal to one in the binary representation of that
//number. You can guarantee that input is non-negative.
//
//Example: The binary representation of 1234 is 10011010010, so the
//function should return 5 in this case
package bitcount

import "fmt"

func CountBits(digit uint) (count int) {
	bin_str := fmt.Sprintf("%b", int64(int(digit)))
	for _, val := range bin_str {
		if (val == '1') {
			count++
		}
	}
	return
}