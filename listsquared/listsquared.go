package listsquared

import "math"

func getDivisors(n int) (retSlc []int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			retSlc = append(retSlc, i)
		}
	}
	return
}

func getSumSquares(divisors []int) float64 {
	var retSum float64
	for _, val := range divisors {
		retSum += float64(val * val)
	}
	return retSum
}

func ListSquared(m, n int) [][]int {
	var retSlc [][]int = make([][]int, 0)
	for i := m; i < n; i++ {
		divisors := getDivisors(i)
		if sumSquare := math.Sqrt(getSumSquares(divisors)); sumSquare == math.Round(sumSquare) {
			retSlc = append(retSlc, []int{i, int(sumSquare * sumSquare)})
		}
	}
	return retSlc
}
