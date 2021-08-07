package stringmix

import (
	"sort"
	"strings"
)

type str struct {
	letter rune
	count1 int
	count2 int
}

type mainMap map[rune]*str

func isAlpha(ch rune) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	} else {
		return false
	}
}

func (m mainMap) addStr(s string, pos int) {
	for _, val := range s {
		if isAlpha(val) {
			if dig, found := m[val]; !found {
				if pos == 1 {
					m[val] = &str{val, 1, 0}
				} else {
					m[val] = &str{val, 0, 1}
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
}

func (m mainMap) deleteLoner() {
	for _, val := range m {
		if val.count1 <= 1 && val.count2 <= 1 {
			delete(m, val.letter)
		}
	}
}

func sprintResult(strNum rune, letter rune, count int) string {
	str := make([]rune, 0, count+3)
	str = append(str, strNum)
	str = append(str, ':')
	for i := 0; i < count; i++ {
		str = append(str, letter)
	}
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
	mainMap.deleteLoner()
	for _, val := range mainMap {
		retStr = append(retStr, handleResult(*val))
	}
	sort.Slice(retStr, func(i, j int) bool {
		if len(retStr[i]) < len(retStr[j]) {
			return false
		} else if len(retStr[i]) > len(retStr[j]) {
			return true
		} else {
			return retStr[i] < retStr[j]
		}
	})
	size := len(retStr) - 1
	if size == -1 {
		return ""
	}
	retStr[size] = retStr[size][:len(retStr[size])-1]
	return strings.Join(retStr, "")
}
