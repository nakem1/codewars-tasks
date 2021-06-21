//Write a function that takes a string of braces, and determines if the
//order of the braces is valid. It should return true if the string is
//valid, and false if it's invalid.
//
//This Kata is similar to the Valid Parentheses Kata, but introduces
//new characters: brackets [], and curly braces {}. Thanks to @arnedag for the idea!
//
//All input strings will be nonempty, and will only consist of parentheses,
//brackets and curly braces: ()[]{}.
//
//What is considered Valid?
//
//A string of braces is considered valid if all braces are matched with
//the correct brace.
//
//"(){}[]"   =>  True
//"([{}])"   =>  True
//"(}"       =>  False
//"[(])"     =>  False
//"[({})](]" =>  False
package validbraces

import "errors"

type stack struct {
	stack  []rune
	lenght int
}

func (st *stack) Push(value rune) {
	st.stack = append(st.stack, value)
	st.lenght += 1
}

func (st *stack) Pop() error {
	if st.lenght <= 0 {
		err := errors.New("Pop with empty stack")
		return err
	}
	st.stack = st.stack[:len(st.stack)-1]
	st.lenght -= 1
	return (nil)
}

func (Brackets *stack) addCloseBracket(closeBracket rune) bool {
	var openBracket rune
	switch closeBracket {
	case ')':
		openBracket = '('
	case ']':
		openBracket = '['
	case '}':
		openBracket = '{'
	}
	if (Brackets.lenght >= 1) {
		Brackets.Push(closeBracket)
	} else {
		return (false)
	}
	if (Brackets.stack[len(Brackets.stack) - 2] == openBracket) {
		Brackets.Pop()
		Brackets.Pop()
	}
	return (true)
}

func ValidBraces(str string) bool {
	var Brackets stack
	for _, val := range str {
		switch val {
		case '(', '[', '{':
			Brackets.Push(val)
		case ')', ']', '}':
			if (!Brackets.addCloseBracket(val)) {
				return (false)
			}
		}
	}
	return (Brackets.lenght == 0)
}
