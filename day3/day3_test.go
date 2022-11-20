package day3

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestPuzzleOne(t *testing.T) {
	content, err := ioutil.ReadFile("../inputs/day-3.txt")
	if err != nil {
    t.Error("Failed to read file!")
	}

  strs := strings.Split(strings.TrimSpace(string(content)), "\n")
  tris := make([]Triangle, len(strs))
  regex := regexp.MustCompile(`\d+`)
  for i, s := range strs {
    sides := regex.FindAllString(strings.TrimSpace(s), -1)
    if len(sides) != 3 {
      t.Errorf("%v", sides)
    }
    s1, err := strconv.Atoi(sides[0])
    s2, err := strconv.Atoi(sides[1])
    s3, err := strconv.Atoi(sides[2])
    if err != nil {
      t.Errorf("Failed to parse %s", s)
    }
    tris[i] = Triangle{s1, s2, s3}
  }

  count := 0
  for _, t := range tris {
    if validateTriangle(t) {
      count += 1
    }
  }

  if count != 869 {
    t.Errorf("Got %d, should have been 869", count)
  }
}

func TestPuzzleTwo(t *testing.T) {
	content, err := ioutil.ReadFile("../inputs/day-3.txt")
	if err != nil {
    t.Error("Failed to read file!")
	}

  strs := strings.Split(strings.TrimSpace(string(content)), "\n")
  triStrs := make([][]string, len(strs))
  regex := regexp.MustCompile(`\d+`)
  for i, s := range strs {
    triStrs[i] = regex.FindAllString(strings.TrimSpace(s), -1)
  }

  nums := make([]int, len(triStrs) * 3)
  for i, s := range triStrs {
    nums[i], err = strconv.Atoi(s[0])
    if err != nil {
      t.Errorf("Failed to convert '%s' to int", s[0])
    }
    nums[i + len(triStrs)], err = strconv.Atoi(s[1])
    if err != nil {
      t.Errorf("Failed to convert '%s' to int", s[1])
    }
    nums[i + len(triStrs) * 2], err = strconv.Atoi(s[2])
    if err != nil {
      t.Errorf("Failed to convert '%s' to int", s[2])
    }
  }

  count := 0
  for i := range triStrs {
    n := i * 3
    s1 := nums[n]
    s2 := nums[n + 1]
    s3 := nums[n + 2]
    tri := Triangle{s1, s2, s3}
    if validateTriangle(tri) {
      count = count + 1
    }
  }

  if count != 1544 {
    t.Errorf("%d does not equal 1544", count)
  }
}
