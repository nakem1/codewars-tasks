package validparentheses

import "testing"

type tests struct {
	s1   string
	want bool
}

func TestValidParentheses(t *testing.T) {
	var check []tests = []tests{
		{"()", true},
		{")", false},
		{"()()()()()()())", false},
		{"()()()()()()()", true},
	}
	for _, test := range check {
		if got := ValidParentheses(test.s1); got != test.want {
			t.Errorf("ValidParentheses(%#v) == %#v, want = %#v", test.s1, got, test.want)
		}
	}
}
