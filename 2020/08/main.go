package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/SapienceDev/advent-of-code/utils"
)

type program struct {
	Accumulator int
	Visited     []int
}

func main() {
	input := utils.GetInput(2020, 8)
	instructions := strings.Split(input, "\n")
	program := program{}
	program.visitInstruction(instructions, 0)
}

func (p *program) visitInstruction(instructions []string, pos int) {
	if contains(p.Visited, pos) {
		fmt.Printf("Accumulator: %d\n", p.Accumulator)
		fmt.Printf("Pos: %d\n", pos)
		return
	}
	p.Visited = append(p.Visited, pos)
	instruction := instructions[pos]

	re := regexp.MustCompile("(nop|acc|jmp)\\s((\\+|-)[0-9]+)")
	res := re.FindAllStringSubmatch(instruction, -1)
	action := res[0][1]
	adj, _ := strconv.Atoi(res[0][2])

	switch action {
	case "nop":
		p.visitInstruction(instructions, pos+1)
		break
	case "acc":
		p.Accumulator += adj
		p.visitInstruction(instructions, pos+1)
		break
	case "jmp":
		p.visitInstruction(instructions, pos+adj)
		break
	}
}

func contains(arr []int, num int) bool {
	for _, val := range arr {
		if val == num {
			return true
		}
	}
	return false
}
