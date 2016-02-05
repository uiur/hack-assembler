package instruction

import "testing"

func TestCodeA(t *testing.T) {
	inst := &Instruction{CommandType: "a", Symbol: "100"}

	expected := "0000000001100100"
	actual := inst.Code(map[string]int{})

	if !isCode(actual) {
		t.Errorf("%s is not code", actual)
	}

	if actual != expected {
		t.Errorf("got %s, expected %s", actual, expected)
	}
}

func TestCodeASymbol(t *testing.T) {
	inst := &Instruction{CommandType: "a", Symbol: "R2"}

	expected := "0000000000000010"
	actual := inst.Code(map[string]int{})

	if !isCode(actual) {
		t.Errorf("%s is not code", actual)
	}

	if actual != expected {
		t.Errorf("got %s, expected %s", actual, expected)
	}
}

func TestCodeC(t *testing.T) {
	inst := &Instruction{CommandType: "c", Comp: "A", Dest: "D"} // D=A

	expected := "1110110000010000"
	actual := inst.Code(map[string]int{})

	if !isCode(actual) {
		t.Errorf("%s is not code", actual)
	}

	if actual != expected {
		t.Errorf("got %s, expected %s", actual, expected)
	}
}

func TestCodeJump(t *testing.T) {
	inst := &Instruction{CommandType: "c", Comp: "D", Jump: "JGT"} // D;JGT

	expected := "1110001100000001"
	actual := inst.Code(map[string]int{})

	if !isCode(actual) {
		t.Errorf("%s is not code", actual)
	}

	if actual != expected {
		t.Errorf("got %s, expected %s", actual, expected)
	}
}

func isCode(str string) bool {
	return len(str) == 16
}
