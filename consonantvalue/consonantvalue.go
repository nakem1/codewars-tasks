//Given a lowercase string that has alphabetic characters only and no spaces,
//return the highest value of consonant substrings. Consonants are any
//letters of the alphabet except "aeiou".
//
//We shall assign the following values: a = 1, b = 2, c = 3, .... z = 26.
package consonantvalue

func checkConsonant(val byte) bool {
	if val > 97 && val <= 122 {
		switch val {
		case 'a', 'e', 'i', 'o', 'u':
			return false
		default:
			return true
		}
	}
	return false
}

func Solve(str string) int {
	var max, sum, size int = 0, 0, len(str)
	for i := 0; i < size; i++ {
		if checkConsonant(str[i]) {
			sum += int(str[i]) - 96
		} else {
			if max < sum {
				max = sum
			}
			sum = 0
		}
	}
	if max < sum {
		max = sum
	}
	return (max)
}
