package partitions

func Partitions(num int) int {
	var retSlc [][]int
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
	return (len(retSlc))
}
