//Write a function that takes in a string of one or more words, \
//and returns the same string, but with all five or more letter words \
//reversed (like the name of this kata).
//
//1) Strings passed in will consist of only letters and spaces.
//2) Spaces will be included only when more than one word is present.
//
//spinWords("Hey fellow warriors") => "Hey wollef sroirraw"
//spinWords("This is a test") => "This is a test"
//spinWords("This is another test") => "This is rehtona test"
package spinwords

import "strings"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SpinWords(str string) string {
	str_slc := strings.Fields(str)
	ret_str := make([]string, len(str_slc))
	for i, val := range str_slc {
		if len(val) < 5 {
			ret_str[i] = val
		} else {
			ret_str[i] = Reverse(val)
		}
	}
	return (strings.Join(ret_str, " "))
}
