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
	for i := 0; i < len(instructions); i++ {
		instructions := strings.Split(input, "\n")
		program := program{}
		if strings.HasPrefix(instructions[i], "nop") || strings.HasPrefix(instructions[i], "jmp") {
			if strings.HasPrefix(instructions[i], "nop") {
				instructions[i] = strings.Replace(instructions[i], "nop", "jmp", -1)
			} else {
				instructions[i] = strings.Replace(instructions[i], "jmp", "nop", -1)
			}
		}
		done := program.visitInstruction(instructions, 0)
		if done {
			break
		}
	}
}

func (p *program) visitInstruction(instructions []string, pos int) bool {
	if pos >= len(instructions) {
		fmt.Printf("acc: %d pos %d\n", p.Accumulator, pos)
		return true
	}
	if contains(p.Visited, pos) {
		// Part 1 answer: fmt.Printf("acc: %d pos: %d\n", p.Accumulator, pos)
		return false
	}
	p.Visited = append(p.Visited, pos)
	instruction := instructions[pos]

	re := regexp.MustCompile("(nop|acc|jmp)\\s((\\+|-)[0-9]+)")
	res := re.FindAllStringSubmatch(instruction, -1)
	action := res[0][1]
	adj, _ := strconv.Atoi(res[0][2])

	switch action {
	case "nop":
		done := p.visitInstruction(instructions, pos+1)
		if done {
			return true
		}
		break
	case "acc":
		p.Accumulator += adj
		done := p.visitInstruction(instructions, pos+1)
		if done {
			return true
		}
		break
	case "jmp":
		done := p.visitInstruction(instructions, pos+adj)
		if done {
			return true
		}
		break
	}
	return false
}

func contains(arr []int, num int) bool {
	for _, val := range arr {
		if val == num {
			return true
		}
	}
	return false
}
