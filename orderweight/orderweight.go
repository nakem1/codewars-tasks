package orderweight

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type human struct {
	weight        int
	privateWeight int
}

type sorting []human

func (p sorting) Less(i, j int) bool {
	ret := p[i].privateWeight < p[j].privateWeight
	if p[i].privateWeight == p[j].privateWeight {
		ret = strconv.Itoa(p[i].weight) < strconv.Itoa(p[j].weight)
	}
	return ret
}

func (p sorting) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p sorting) Len() int {
	return len(p)
}

func sumDig(dig int) int {
	var ret int
	for dig != 0 {
		tmpI := dig % 10
		dig /= 10
		ret += tmpI
	}
	return (ret)
}

func OrderWeight(str string) string {
	var ret []string
	var humans []human
	slc := strings.Fields(str)
	for _, val := range slc {
		var dig int
		var err error
		if dig, err = strconv.Atoi(val); err != nil {
			fmt.Printf("Error")
		}
		sum := sumDig(dig)
		humans = append(humans, human{dig, sum})
	}
	sort.Sort(sorting(humans))
	for _, val := range humans {
		ret = append(ret, strconv.Itoa(val.weight))
	}
	return (strings.Join(ret, " "))
}
