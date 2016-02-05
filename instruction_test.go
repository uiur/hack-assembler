package main

import "testing"

func TestCodeA(t *testing.T) {
	inst := &Instruction{CommandType: "a", Symbol: "100"}

	expected := "0000000001100100"
	actual := inst.Code()
	if actual != expected {
		t.Errorf("got %s, expected %s", actual, expected)
	}
}

func TestCodeC(t *testing.T) {
	inst := &Instruction{CommandType: "c", Comp: "A", Dest: "D"} // D=A

	expected := "1110110000010000"
	actual := inst.Code()

	if actual != expected {
		t.Errorf("got %s, expected %s", actual, expected)
	}
}
