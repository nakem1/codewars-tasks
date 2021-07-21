// import (
//  "strings"
//)
//
//func DecodeMorse(morseCode string) (decoded string) {
//  for _, word := range strings.Split(strings.TrimSpace(morseCode), "   ") {
//    for _, char := range strings.Split(word, " "){
//      decoded += MORSE_CODE[char]
//    }
//    decoded += " "
//  }
//  return strings.TrimSpace(decoded)
//}
// this is best practice solution
package decodemorse

import "strings"

func DecodeMorse(morseCode string) string {
	var str []string
	var tmpI, cutStart, cutEnd int = 0, 0, len(morseCode)
	for _, val := range morseCode {
		if val != ' ' {
			break
		}
		cutStart++
	}
	for i := len(morseCode) - 1; i >= 0; i-- {
		if morseCode[i] != ' ' {
			break
		}
		cutEnd--
	}
	newCode := strings.ReplaceAll(morseCode[cutStart:cutEnd], "   ", "+")
	for i, val := range newCode {
		if val == ' ' || val == '+' {
			str = append(str, MORSE_CODE[newCode[tmpI:i]])
			tmpI = i + 1
			if val == '+' {
				str = append(str, " ")
			}
		}
	}
	str = append(str, MORSE_CODE[newCode[tmpI:]])
	return strings.Join(str, "")
}
