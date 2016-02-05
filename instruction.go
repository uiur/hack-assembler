package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	CommandType, Symbol, Dest, Comp, Jump string
}

func (inst *Instruction) Code() string {
	if inst.CommandType == "a" {
		value, err := strconv.Atoi(inst.Symbol)

		if err != nil {
			os.Exit(1)
		}

		return fmt.Sprintf("0%015b", value)
	} else if inst.CommandType == "c" {
		comp := compToCode(inst.Comp)

		a := strings.Contains(inst.Dest, "A")
		d := strings.Contains(inst.Dest, "D")
		m := strings.Contains(inst.Dest, "M")
		dest := fmt.Sprintf("%b%b%b", btoi(a), btoi(d), btoi(m))

		return fmt.Sprintf("111%07s%s%03s", comp, dest, inst.Jump)
	}

	return ""
}

func compToCode(comp string) string {
	a := btoi(strings.Contains(comp, "M"))
	return fmt.Sprintf("%b%s", a, compToCCode(comp))
}

func compToCCode(comp string) string {
	switch comp {
	case "0":
		return "101010"
	case "1":
		return "111111"
	case "-1":
		return "1110010"
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
