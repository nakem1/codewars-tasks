//Return the number (count) of vowels in the given string.
//We will consider a, e, i, o, u as vowels for this Kata (but not y).
//The input string will only consist of lower case letters and/or spaces.
package vowelcount

import "strings"

func GetCount(str string) (count int) {
	const needle string = "aeiou"
	for _, val := range str {
		if (strings.ContainsRune(needle, val)) {
			count++
		}
	}
	return (count)
}