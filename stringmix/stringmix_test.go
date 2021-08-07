package stringmix

import (
	"testing"
)

type tests struct {
	s1   string
	s2   string
	want string
}

func TestMix(t *testing.T) {
	var check []tests = []tests{
		{"Are they here", "yes, they are here", "2:eeeee/2:yy/=:hh/=:rr"},
		{"uuuuuu", "uuuuuu", "=:uuuuuu"},
		{"looping is fun but dangerous", "less dangerous than coding",
			"1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg"},
		{"codewars", "codewars", ""},
	}
	for _, test := range check {
		if got := Mix(test.s1, test.s2); got != test.want {
			t.Errorf("Mix(%v, %v) == %v, want %v\n", test.s1, test.s2, got, test.want)
		}
	}
}
