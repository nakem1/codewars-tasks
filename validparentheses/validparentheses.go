package validparentheses

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
	if Brackets.lenght >= 1 {
		Brackets.Push(closeBracket)
	} else {
		return (false)
	}
	if Brackets.stack[len(Brackets.stack)-2] == '(' {
		Brackets.Pop()
		Brackets.Pop()
	}
	return (true)
}

func ValidParentheses(str string) bool {
	var Brackets stack
	for _, val := range str {
		switch val {
		case '(':
			Brackets.Push(val)
		case ')':
			if !Brackets.addCloseBracket(val) {
				return (false)
			}
		}
	}
	return (Brackets.lenght == 0)
}
