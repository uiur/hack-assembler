package instruction

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	CommandType, Symbol, Dest, Comp, Jump string
}

func (inst *Instruction) Code(symbols map[string]int) string {
	if inst.CommandType == "a" {
		return fmt.Sprintf("0%015b", symbolToInt(inst.Symbol, symbols))
	} else if inst.CommandType == "c" {
		comp := compToCode(inst.Comp)

		a := strings.Contains(inst.Dest, "A")
		d := strings.Contains(inst.Dest, "D")
		m := strings.Contains(inst.Dest, "M")
		dest := fmt.Sprintf("%b%b%b", btoi(a), btoi(d), btoi(m))

		return fmt.Sprintf("111%07s%s%03s", comp, dest, jumpToCode(inst.Jump))
	}

	return ""
}

func symbolToInt(str string, symbols map[string]int) int {
	value, err := strconv.Atoi(str)

	if err == nil {
		return value
	}

	r := regexp.MustCompile(`^R(\d+)$`)
	matches := r.FindStringSubmatch(str)
	if len(matches) > 0 {
		number, _ := strconv.Atoi(matches[1])
		return number
	}

	switch str {
	case "SP":
		return 0
	case "LCL":
		return 1
	case "ARG":
		return 2
	case "THIS":
		return 3
	case "THAT":
		return 4
	case "SCREEN":
		return 16384
	case "KBD":
		return 24576
	}

	return symbols[str]
}

func compToCode(comp string) string {
	a := btoi(strings.Contains(comp, "M"))
	return fmt.Sprintf("%b%s", a, compToCCode(comp))
}

var jumpTable = map[string]string{
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

func jumpToCode(jump string) string {
	return jumpTable[jump]
}

func compToCCode(comp string) string {
	switch comp {
	case "0":
		return "101010"
	case "1":
		return "111111"
	case "-1":
		return "111010"
	case "D":
		return "001100"
	case "A", "M":
		return "110000"
	case "!D":
		return "001101"
	case "!A", "!M":
		return "110001"
	case "-D":
		return "001111"
	case "-A", "-M":
		return "110011"
	case "D+1":
		return "011111"
	case "A+1", "M+1":
		return "110111"
	case "D-1":
		return "001110"
	case "A-1", "M-1":
		return "110010"
	case "D+A", "D+M":
		return "000010"
	case "D-A", "D-M":
		return "010011"
	case "A-D", "M-D":
		return "000111"
	case "D&A", "D&M":
		return "000000"
	case "D|A", "D|M":
		return "010101"
	default:
		return ""
	}

}

func btoi(b bool) int {
	if b {
		return 1
	}

	return 0
}
