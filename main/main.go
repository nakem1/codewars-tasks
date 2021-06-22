package main

import (
	"fmt"

	"github.com/nakem1/codewars_tasks/asmtranslate"
)

func main() {
	// val := os.Args[1]
	// val1, err := strconv.Atoi(os.Args[2])
	// if err != nil {
	// panic(err)
	// }
	str := []string{"mov a 1", "mov b 1", "mov c 0", "mov d 26", "jnz c 2",
		"jnz 1 5", "mov c 7", "inc d", "dec c", "jnz c -2", "mov c a", "inc a",
		"dec b", "jnz b -2", "mov b c", "dec d", "jnz d -6", "mov c 18",
		"mov d 11", "inc a", "dec d", "jnz d -2", "dec c", "jnz c -5"}
	fmt.Printf("%#v\n", asmtranslate.SimpleAssembler(str))
}
