package main

import (
	"fmt"
	"io/ioutil"
	"os"

	p "github.com/uiureo/hack-assembler/parser"
)

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1]
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		p := new(p.Parser)
		insts := p.Parse(string(data))
		fmt.Print(p.Generate(insts))
	}
}
