package day4

import (
	"io/ioutil"
	"testing"
)

func TestParseInput(t *testing.T) {
  input := "aaaaa-bbb-z-y-x-123[abxyz]"
  rooms, err := parseInput(input)
  if err != nil {
    t.Errorf("Failed to parse first test case!")
  }

  if len(rooms) != 1 {
    t.Errorf("Incorrect parsing: %v", rooms)
  }
}

func TestIsValid(t *testing.T) {
  input := "aaaaa-bbb-z-y-x-123[abxyz]"
  rooms, err := parseInput(input)
  if err != nil {
    t.Errorf("Failed to parse first test case!")
  }

  room := rooms[0]

  if !room.IsValid() {
    t.Error("Failed to successfully validate room!")
  }

  input = "a-b-c-d-e-f-g-h-987[abcde]"

  rooms, err = parseInput(input)
  if err != nil {
    t.Errorf("Failed to parse first test case!")
  }

  room = rooms[0]

  if !room.IsValid() {
    t.Error("Failed to successfully validate room!")
  }

  input = "not-a-real-room-404[oarel]"
  rooms, err = parseInput(input)
  if err != nil {
    t.Errorf("Failed to parse first test case!")
  }

  room = rooms[0]

  if !room.IsValid() {
    t.Error("Failed to successfully validate room!")
  }

  input = "totally-real-room-200[decoy]"
  rooms, err = parseInput(input)
  if err != nil {
    t.Errorf("Failed to parse first test case!")
  }

  room = rooms[0]

  if room.IsValid() {
    t.Error("Failed to successfully invalidate room!")
  }
}

func TestSolveFirstPuzzle(t *testing.T) {
  content, err := ioutil.ReadFile("../inputs/day-4.txt")
  if err != nil {
    t.Errorf("Failed to read file!")
  }

  rooms, err := parseInput(string(content))
  if err != nil {
    t.Errorf("Failed to parse file contents!")
  }

  count := 0
  for _, room := range rooms {
    if room.IsValid() {
      count = count + room.SectorId
    }
  }

  if count != 185371 {
    t.Errorf("Failed to calculate correct number: got '%d'", count)
  }
}


func TestSolveSecondPuzzle(t *testing.T) {
  t.Errorf("Unimplemented")
}
