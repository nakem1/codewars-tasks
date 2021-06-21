package main

import (
	"fmt"

	"github.com/nakem1/codewars_tasks/validbraces"
)

func main() {
	// val := os.Args[1]
	// val1, err := strconv.Atoi(os.Args[2])
	// if err != nil {
		// panic(err)
	// }
	fmt.Printf("%v\n", validbraces.ValidBraces("(({{[[]]}}))"))
}
