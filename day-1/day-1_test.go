package day_1

import (
	"fmt"
	"io/ioutil"
	"testing"
)

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

func TestTupleSet(t *testing.T) {
	ts := TupleSet{}
	result := ts.get(1, 1)
	if result {
		t.Errorf("1, 1 should not be true!")
	}

	ts.put(1, 1)
	result = ts.get(1, 1)
	if !result {
		t.Errorf("1, 1 should be true!")
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

	filename := "../inputs/day-1.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Missing input file %s!", filename)
	}

	input = string(content)
	result = PuzzleOne(input)
	if result != 332 {
		t.Errorf("%d does not equal 332!", result)
	}
}

func TestMakeline(t *testing.T) {
	minx := 5
	maxx := 5
	miny := 5
	maxy := 10
	result := Makeline(minx, maxx, miny, maxy)
	if result[0][0] != 5 || result[0][1] != 6 {
		t.Errorf("Going down doesn't work")
	}
	minx = 5
	maxx = 2
	miny = 5
	maxy = 5
	result = Makeline(minx, maxx, miny, maxy)
	if result[0][0] != 4 || result[0][1] != 5 {
		t.Errorf("Going down doesn't work")
	}
}

func TestPuzzleTwo(t *testing.T) {
	input := "R8, R4, R4, R8"
	result := PuzzleTwo(input)
	if result != 4 {
		t.Errorf("%d does not equal 4!", result)
	}

	filename := "../inputs/day-1.txt"
	read_input, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Missing input file %s!", filename)
	}
	fmt.Printf("%d\n", PuzzleTwo(string(read_input)))
}
