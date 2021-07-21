package pyramid

func retIntArray(n int) []int {
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = 1
	}
	return (tmp)
}

func Pyramid(n int) [][]int {
	arr := make([][]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = retIntArray(i)
	}
	return (arr)
}
