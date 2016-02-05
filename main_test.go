package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"
)

func TestMain(t *testing.T) {
	matches, err := filepath.Glob("fixtures/*.asm")

	if err != nil {
		t.Error(err)
	}

	for _, file := range matches {
		r := regexp.MustCompile(`fixtures/(.+)\.asm`)
		submatch := r.FindStringSubmatch(file)
		name := submatch[1]
		hackFile := fmt.Sprintf("fixtures/%s.hack", name)

		expected, err := ioutil.ReadFile(hackFile)

		if err != nil {
			t.Error(err)
		}

		actual, err := exec.Command("go", "run", "main.go", "instruction.go", file).Output()

		if err != nil {
			t.Error(err)
		}

		if string(actual) != string(expected) {
			t.Errorf("expect:\n%s\nactual:\n%s", string(expected), string(actual))
		}
	}
}
