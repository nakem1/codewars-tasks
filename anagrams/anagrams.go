package anagrams

// import "reflect"

// func Anagrams(word string, words []string) []string {
// 	var retSlc []string = make([]string, 0, len(words))
// 	mapWord := make(map[rune]int)
// 	for _, w := range word {
// 		mapWord[w]++
// 	}
// 	for _, w := range words {
// 		mapRunes := make(map[rune]int)
// 		for _, r := range w {
// 			mapRunes[r]++
// 		}
// 		if reflect.DeepEqual(mapRunes, mapWord) {
// 			retSlc = append(retSlc, w)
// 		}
// 	}
// 	if len(retSlc) == 0 {
// 		return nil
// 	} else {
// 		return retSlc
// 	}
// }

func Anagrams(word string, words []string) []string {
	var anagrams []string
	wordSum := sumChars(word, -1)
	for _, tw := range words {
		twsum := sumChars(tw, wordSum)
		if twsum >= 0 && twsum == wordSum {
			anagrams = append(anagrams, tw)
		}
	}
	return anagrams
}

func sumChars(word string, limit int) int {
	total := 0
	for _, c := range word {
		total += int(c)
		if limit >= 0 && total > limit {
			return -1
		}
	}
	return total
}
