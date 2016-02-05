package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1]
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		parse(string(data))
	}
}

func parse(str string) {
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
			inst := &Instruction{CommandType: "a", Symbol: trimmedLine[1:]}
			fmt.Println(inst.Code())
		} else {
			inst := parseAssignment(trimmedLine)
			fmt.Println(inst.Code())
		}
	}
}

func parseAssignment(line string) *Instruction {
	r := regexp.MustCompile(`^(.+)=(.+)$`)
	match := r.FindStringSubmatch(line)

	if len(match) == 0 {
		return nil
	}

	left := match[1]
	right := match[2]

	return &Instruction{CommandType: "c", Dest: left, Comp: right}
}
