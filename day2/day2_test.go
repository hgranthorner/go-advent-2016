package day2

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestDeterminePasswordTestDeterminePassword(t *testing.T) {
	directions := []string{"ULL",
		"RRDDD",
		"LURDL",
		"UUUUD"}
  result := determinePassword(directions, pad)

  if result != "1985" {
    t.Errorf("Didn't determine the password")
  }
}

func TestPuzzleOne(t *testing.T) {
  content, err := ioutil.ReadFile("../inputs/day-2.txt")
  if err != nil {
    t.Errorf("Failed to read file!")
  }
  
  directionStr := string(content)
  directions := strings.Split(strings.TrimSpace(directionStr), "\n")
  result := determinePassword(directions, pad)
  t.Logf("%s", result)
}

func TestPuzzleTwo(t *testing.T) {
  content, err := ioutil.ReadFile("../inputs/day-2.txt")
  if err != nil {
    t.Errorf("Failed to read file!")
  }
  
  directionStr := string(content)
  directions := strings.Split(strings.TrimSpace(directionStr), "\n")
  result := determinePassword(directions, pad2)
  t.Logf("%s", result)
  t.Fail()
}
