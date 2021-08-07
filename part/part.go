package part

import (
	"fmt"
	"sort"
)

func median(arr []int) float32 {
	if len(arr) == 1 {
		return float32(arr[0])
	}
	if len(arr)%2 == 1 {
		return (float32(arr[(len(arr) / 2)]))
	} else {
		var floatTmp float32 = float32(arr[len(arr)/2]+
			arr[(len(arr)/2)-1]) / 2
		return (floatTmp)
	}
}

func sumArr(arr []int) float32 {
	var sum int

	for _, value := range arr {
		sum += value
	}
	return float32(sum)
}

func mean(arr []int) float32 {
	return (sumArr(arr) / float32(len(arr)))
}

func rangeDig(arr []int) int {
	return arr[len(arr)-1] - arr[0]
}

func uniqueNum(num int) (retSlc [][]int) {
	var k int = 0
	var piece []int = make([]int, num)
	piece[k] = num

	for {
		tmp := make([]int, len(piece))
		copy(tmp, piece)
		retSlc = append(retSlc, tmp)
		rem_val := 0
		for k >= 0 && piece[k] == 1 {
			rem_val += piece[k]
			k--
		}
		if k < 0 {
			return
		}
		piece[k]--
		rem_val++
		for rem_val > piece[k] {
			piece[k+1] = piece[k]
			rem_val -= piece[k]
			k++
		}
		piece[k+1] = rem_val
		k++
	}
}

func multiSlc(arr []int) int {
	var sum int = 1

	for _, val := range arr {
		if val == 0 {
			continue
		}
		sum *= val
	}
	return sum
}

func sortDeleteDup(partN [][]int) (retSlc []int) {
	mapDup := make(map[int]bool)
	for _, val := range partN {
		mapDup[multiSlc(val)] = true
	}
	for key, _ := range mapDup {
		retSlc = append(retSlc, key)
	}
	sort.Ints(retSlc)
	return
}

func Part(num int) string {
	partN := uniqueNum(num)
	slc := sortDeleteDup(partN)
	rangeN := rangeDig(slc)
	meanN := mean(slc)
	medianN := median(slc)
	return fmt.Sprintf("Range: %d Average: %.2f Median: %.2f", rangeN, meanN, medianN)
}
