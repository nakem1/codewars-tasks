package stringmix

import (
	"fmt"
	"strings"
)

type str struct {
	letter   rune
	count1   int
	count2   int
	maxCount int
}

type max struct {
	maxDig  int
	maxChar rune
}

type mainMap map[rune]*str

func (m mainMap) addStr(s string, pos int) {
	for _, val := range s {
		if dig, found := m[val]; !found {
			if pos == 1 {
				m[val] = &str{val, 1, 0, -1}
			} else {
				m[val] = &str{val, 0, 1, -1}
			}
		} else {
			if pos == 1 {
				dig.count1++
			} else {
				dig.count2++
			}
		}
	}
}

func (m mainMap) findMax() (str, error) {
	var maxTmp int = -1
	var max str = str{'+', -1, -1, -1}
	for key, val := range m {
		if val.count1 > val.count2 {
			maxTmp = val.count1
		} else {
			maxTmp = val.count2
		}
		if maxTmp == max.maxCount && key > max.letter || maxTmp > max.maxCount {
			max = *val
		}
	}
	if max.maxCount == -1 {
		err := fmt.Errorf("maxdig not found")
		return str{}, err
	}
	delete(m, max.letter)
	return max, nil
}

func (m mainMap) setMaxCount() {
	for _, val := range m {
		if val.count1 > val.count2 {
			val.maxCount = val.count1
		} else {
			val.maxCount = val.count2
		}
	}
}

func sprintResult(strNum rune, letter rune, count int) string {
	str := make([]rune, count+3)
	str = append(str, strNum)
	str = append(str, ':')
	for i := 0; i < count; i++ {
		str = append(str, letter)
	}
	fmt.Printf("/")
	str = append(str, '/')
	return string(str)
}

func handleResult(res str) string {
	if res.count1 > res.count2 {
		return sprintResult('1', res.letter, res.count1)
	} else if res.count2 > res.count1 {
		return sprintResult('2', res.letter, res.count2)
	} else {
		return sprintResult('=', res.letter, res.count1)
	}
}

func Mix(s1, s2 string) string {
	var retStr []string
	mainMap := make(mainMap)
	mainMap.addStr(s1, 1)
	mainMap.addStr(s2, 2)
	mainMap.setMaxCount()
	for {
		res, err := mainMap.findMax()
		if err != nil {
			break
		} else {
			retStr = append(retStr, handleResult(res))
		}
	}
	return strings.Join(retStr, "")
}
