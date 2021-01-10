package main

import (
	"regexp"

	"github.com/SapienceDev/advent-of-code/utils"
)

func missing(element string) bool {
	pid, _ := regexp.MatchString("pid", element)
	ecl, _ := regexp.MatchString("ecl", element)
	byr, _ := regexp.MatchString("byr", element)
	iyr, _ := regexp.MatchString("iyr", element)
	eyr, _ := regexp.MatchString("eyr", element)
	hgt, _ := regexp.MatchString("hgt", element)
	hcl, _ := regexp.MatchString("hcl", element)
	return !(pid && ecl && byr && iyr && eyr && hgt && hcl)
}

func invalid(element string) bool {
	return !(validateByr(element) && validateIyr(element) && validateEyr(element) && validateHgt(element) &&
		validateHcl(element) && validateEcl(element) && validatePid(element))
}

func validateByr(input string) bool {
	boolean, _ := regexp.MatchString("byr:(19[2-9][0-9]|200[0-2])", input)
	return boolean
}

func validateIyr(input string) bool {
	boolean, _ := regexp.MatchString("iyr:20(1[0-9]|20)", input)
	return boolean
}

func validateEyr(input string) bool {
	boolean, _ := regexp.MatchString("eyr:20(2[0-9]|30)", input)
	return boolean
}

func validateHgt(input string) bool {
	boolean, _ := regexp.MatchString("hgt:(([5-6][0-9]|7[0-6])in|1([5-8][0-9]|9[0-3])cm)", input)
	return boolean
}

func validateHcl(input string) bool {
	boolean, _ := regexp.MatchString("hcl:#[0-9a-f]{6}", input)
	return boolean
}

func validateEcl(input string) bool {
	boolean, _ := regexp.MatchString("ecl:(amb|blu|brn|gry|grn|hzl|oth)", input)
	return boolean
}

func validatePid(input string) bool {
	boolean, _ := regexp.MatchString("pid:[0-9]{9}\\b", input)
	return boolean
}

func main() {
	re := regexp.MustCompile("\\n\\s")
	input := utils.GetInput(2020, 4)
	arr := re.Split(input, -1)
	var valid int = 0
	for _, element := range arr {
		if !missing(element) && !invalid(element) {
			valid++
		}
	}
	println(valid)
}
