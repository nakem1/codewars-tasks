//Write a simple parser that will parse and run Deadfish.
//
//Deadfish has 4 commands, each 1 character long:
//
//i increments the value (initially 0)
//d decrements the value
//s squares the value
//o outputs the value into the return array
//Invalid characters should be ignored.
package deadfish

func Parse(data string) []int{
	arr := []int{0}
	var point int = 0
	length := len(data)
	for i := 0; i < length; i++ {
		switch data[i] {
		case 'i':
			arr[point]++
		case 'd':
			arr[point]--
		case 's':
			arr[point] *= arr[point]
		case 'o':
			arr = append(arr, arr[point])
			point++
		}
	}
	return (arr[:len(arr) - 1])
}