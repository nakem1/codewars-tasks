package part

import (
	"reflect"
	"testing"
)

type testsUnique struct {
	num  int
	want [][]int
}

type testsPart struct {
	num  int
	want string
}

func TestUniqueNum(t *testing.T) {
	var check []testsUnique = []testsUnique{
		{2, [][]int{{2, 0}, {1, 1}}},
		{3, [][]int{{3, 0, 0}, {2, 1, 0}, {1, 1, 1}}},
		{4, [][]int{{4, 0, 0, 0}, {3, 1, 0, 0}, {2, 2, 0, 0}, {2, 1, 1, 0}, {1, 1, 1, 1}}},
	}
	for _, test := range check {
		if got := uniqueNum(test.num); !reflect.DeepEqual(got, test.want) {
			t.Errorf("uniqueNum(%#v) == %#v, want = %#v", test.num, got, test.want)
		}
	}
}

func TestPart(t *testing.T) {
	var check []testsPart = []testsPart{
		{1, "Range: 0 Average: 1.00 Median: 1.00"},
		{2, "Range: 1 Average: 1.50 Median: 1.50"},
		{3, "Range: 2 Average: 2.00 Median: 2.00"},
		{4, "Range: 3 Average: 2.50 Median: 2.50"},
		{5, "Range: 5 Average: 3.50 Median: 3.50"},
		{6, "Range: 8 Average: 4.75 Median: 4.50"},
	}
	for _, test := range check {
		if got := Part(test.num); got != test.want {
			t.Errorf("Part(%#v) == %#v, want %#v", test.num, got, test.want)
		} else {
			t.Logf("Part(%#v) \u001b[32mSUCCESS\u001B[0m", got)
		}
	}
}
