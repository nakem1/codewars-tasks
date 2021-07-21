//This is the first part of this kata series. Second part is here.
//
//We want to create a simple interpreter of assembler which will support
//the following instructions:
//
//mov x y - copies y (either a constant value or the content of a register) into register x
//inc x - increases the content of the register x by one
//dec x - decreases the content of the register x by one
//jnz x y - jumps to an instruction y steps away (positive means forward,
//negative means backward, y can be a register or a constant), but only if
//x (a constant or a register) is not zero
//Register names are alphabetical (letters only). Constants are always
//integers (positive or negative).
//
//Note: the jnz instruction moves relative to itself. For example, an offset
//of -1 would continue at the previous instruction, while an offset of 2 would skip over the next instruction.
//
//The function will take an input list with the sequence of the program
//instructions and will execute them. The program ends when there are no more
//instructions to execute, then it returns a dictionary with the contents of the registers.
//
//Also, every inc/dec/jnz on a register will always be preceeded by a mov on
//the register first, so you don't need to worry about uninitialized registers.
package asmtranslate

import (
	"strconv"
	"strings"
)

type asm map[string]int

func (asm asm) mov(dest string, src string) {
	if val, err := strconv.Atoi(src); err != nil {
		asm[dest] = asm[src]
	} else {
		asm[dest] = val
	}
}

func (asm asm) inc(x string) {
	asm[x] += 1
}

func (asm asm) dec(x string) {
	asm[x] -= 1
}

func handleCommand(registers asm, split []string, val string) {
	switch split[0] {
	case "mov":
		registers.mov(split[1], split[2])
	case "inc":
		registers.inc(split[1])
	case "dec":
		registers.dec(split[1])
	}
}

func checkVariable(registers asm, str string) bool {
	if val, err := strconv.Atoi(str); err != nil {
		if registers[str] != 0 {
			return true
		}
	} else {
		if val != 0 {
			return true
		}
	}
	return (false)
}

func SimpleAssembler(program []string) map[string]int {
	registers := make(asm)
	size := len(program)
	for i := 0; i < size; i++ {
		split := strings.Fields(program[i])
		if split[0] == "jnz" && checkVariable(registers, split[1]) {
			step, _ := strconv.Atoi(split[2])
			i += step
			if i >= size {
				break
			}
			split = strings.Fields(program[i])
		}
		handleCommand(registers, split, program[i])
	}
	return (map[string]int(registers))
}
