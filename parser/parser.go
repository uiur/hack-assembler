package parser

import (
	"regexp"
	"strings"

	inst "github.com/uiureo/hack-assembler/instruction"
)

type Parser struct{}

func (p *Parser) Parse(str string) []*inst.Instruction {
	var lines []string
	for _, line := range strings.Split(str, "\n") {
		trimmedLine := removeComment(strings.TrimSpace(line))

		if len(trimmedLine) == 0 {
			continue
		}

		lines = append(lines, trimmedLine)
	}

	var insts []*inst.Instruction
	for _, line := range lines {
		inst := parseAInstruction(line)
		if inst != nil {
			insts = append(insts, inst)
			continue
		}

		inst = parseAssignment(line)
		if inst != nil {
			insts = append(insts, inst)
			continue
		}

		inst = parseJump(line)
		if inst != nil {
			insts = append(insts, inst)
			continue
		}

		inst = parseLabel(line)
		if inst != nil {
			insts = append(insts, inst)
			continue
		}
	}

	return insts
}

func (p *Parser) findSymbols(insts []*inst.Instruction) map[string]int {
	symbols := map[string]int{}

	total := 0
	for _, inst := range insts {
		if inst.CommandType == "l" {
			symbols[inst.Symbol] = total
			continue
		}
		total++
	}

	totalVar := 0
	for _, inst := range insts {
		if inst.CommandType == "a" && symbolIsVariable(inst.Symbol) {
			if symbols[inst.Symbol] == 0 {
				symbols[inst.Symbol] = 16 + totalVar
				totalVar++
			}
		}
	}

	return symbols
}

func symbolIsVariable(symbol string) bool {
	return regexp.MustCompile(`^[a-z_\.\$:][0-9a-z_\.\$:]+$`).MatchString(symbol)
}

func (p *Parser) Generate(insts []*inst.Instruction) string {
	symbols := p.findSymbols(insts)

	var str string

	for _, inst := range insts {
		code := inst.Code(symbols)

		if len(code) > 0 {
			str += code + "\n"
		}
	}

	return str
}

func isComment(line string) bool {
	return line[0:2] == "//"
}

func removeComment(line string) string {
	return regexp.MustCompile(`\s*//.+$`).ReplaceAllString(line, "")
}

func parseJump(line string) *inst.Instruction {
	matches := regexp.MustCompile(`^(\S+);(\S+)$`).FindStringSubmatch(line)

	if len(matches) == 0 {
		return nil
	}

	comp := matches[1]
	jump := matches[2]

	return &inst.Instruction{CommandType: "c", Comp: comp, Jump: jump}
}

func parseLabel(line string) *inst.Instruction {
	matches := regexp.MustCompile(`^\((\S+)\)$`).FindStringSubmatch(line)

	if len(matches) > 0 {
		symbol := matches[1]
		return &inst.Instruction{CommandType: "l", Symbol: symbol}
	}

	return nil
}

func parseAInstruction(line string) *inst.Instruction {
	if line[0:1] == "@" {
		return &inst.Instruction{CommandType: "a", Symbol: line[1:]}
	}

	return nil
}

func parseAssignment(line string) *inst.Instruction {
	r := regexp.MustCompile(`^(.+)=(.+)$`)
	match := r.FindStringSubmatch(line)

	if len(match) == 0 {
		return nil
	}

	left := match[1]
	right := match[2]

	return &inst.Instruction{CommandType: "c", Dest: left, Comp: right}
}
