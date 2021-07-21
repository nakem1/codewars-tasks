package mexicanwave

import "unicode"

func returnStr(words string, i int) string {
	var tmp []rune = []rune(words)
	tmp[i] = unicode.ToUpper(tmp[i])
	return (string(tmp))
}

func Wave(words string) []string {
	var size int = len(words)
	ret := make([]string, 0, size)

	for i := 0; i < size; i++ {
		if words[i] != ' ' {
			ret = append(ret, returnStr(words, i))
		}
	}
	return (ret)
}
