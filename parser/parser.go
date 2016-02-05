package parser

import (
	"fmt"
	"regexp"
	"strings"

	inst "github.com/uiureo/hack-assembler/instruction"
)

type Parser struct {
}

func (p *Parser) Parse(str string) {
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if len(trimmedLine) == 0 {
			continue
		}

		if trimmedLine[0:2] == "//" {
			continue
		}

		if trimmedLine[0:1] == "@" {
			inst := &inst.Instruction{CommandType: "a", Symbol: trimmedLine[1:]}
			fmt.Println(inst.Code())
		} else {
			inst := parseAssignment(trimmedLine)
			fmt.Println(inst.Code())
		}
	}
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
