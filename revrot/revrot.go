//The input is a string str of digits. Cut the string into chunks
//(a chunk here is a substring of the initial string) of size sz
//(ignore the last chunk if its size is less than sz).
//
//If a chunk represents an integer such as the sum of the cubes of
//its digits is divisible by 2, reverse that chunk; otherwise rotate
//it to the left by one position. Put together these modified chunks and
//return the result as a string.
//
// if
// 1) sz is <= 0 or if str is empty return ""
// 2) sz is greater (>) than the length of str it is impossible
// to take a chunk of size sz hence return "".
package revrot

import (
	"strconv"
	"strings"
)

func rotate(val string) string {
	ret_str := val[1:]
	return (ret_str + string(val[0]))
}

func reverse(val string) string {
	runes := []rune(val)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return (string(runes))
}

func checkChunk(digit_str string) (int, error) {
	var sum int
	for _, val := range digit_str {
		digit, err := strconv.Atoi(string(val))
		if err != nil {
			return 0, err
		}
		sum += digit * digit * digit
	}
	if sum%2 == 0 {
		return 1, nil
	} else {
		return 0, nil
	}
}

func Revrot(s string, n int) string {
	str := make([]string, 0, n)
	lenght := len(s)
	if n > len(s) || n <= 0 {
		return ""
	}
	for i, j := n, 0; i <= lenght; i += n {
		str = append(str, s[j:i])
		j += n
	}
	for i, val := range str {
		flag, err := checkChunk(val)
		if err != nil {
			return ""
		}
		if flag == 0 {
			str[i] = rotate(val)
		} else {
			str[i] = reverse(val)
		}
	}
	return (strings.Join(str, ""))
}
