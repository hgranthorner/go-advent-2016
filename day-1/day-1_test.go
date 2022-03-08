package day_1

import "testing"

func TestReturns1(t *testing.T) {
	if 1 != 1 {
		t.Errorf("%d does not equal 1", 1)
	}
}

func TestStridx(t *testing.T) {
	input := "abc"
	result := Stridx(input, 0)

	if result != "a" {
		t.Errorf("%s does not equal a!", result)
	}

	result = Stridx(input, 1)
	if result != "b" {
		t.Errorf("%s does not equal b!", result)
	}
}

func TestPuzzleOne(t *testing.T) {
	input := "R2, L3"
	result := PuzzleOne(input)
	if result != 5 {
		t.Errorf("%s does not equal 5! Got %d instead.", input, result)
	}

	input = "R2, R2, R2"
	result = PuzzleOne(input)
	if result != 2 {
		t.Errorf("%s does not equal 2! Got %d instead.", input, result)
	}

	input = "R5, L5, R5, R3"
	result = PuzzleOne(input)
	if result != 12 {
		t.Errorf("%s does not equal 12! Got %d instead.", input, result)
	}
}
