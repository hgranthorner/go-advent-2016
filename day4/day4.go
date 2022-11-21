package day4

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)
const A = int('a')

type Room struct {
  Encrypted string
  SectorId int
  Checksum string
}

type charCount struct{
  rune rune
  count uint
}

func (r Room) IsValid() bool {
  var chars [26]charCount
  for _, c := range r.Encrypted {
    if c == '-' {
      continue
    }

    chars[int(c - 'a')].rune = c
    chars[int(c - 'a')].count = chars[int(c - 'a')].count + 1
  }

  sort.SliceStable(chars[:], func(i, j int) bool { return chars[i].count > chars[j].count})

  result := true
  for i, cc := range chars[:len(r.Checksum)] {
    if r.Checksum[i] != byte(cc.rune) {
      result = false
    }
  }

  return result
}

var regex = regexp.MustCompile(`(.*)-(\d+)\[(\w+)\]`)

func parseInput(content string) ([]Room, error) {
  lines := strings.Split(strings.TrimSpace(content), "\n")
  rooms := make([]Room, len(lines))
  for i, line := range lines {
    groups := regex.FindStringSubmatch(line)
    id, err := strconv.Atoi(groups[2])
    if err != nil {
      fmt.Printf("Failed to parse sector id '%s' in group '%v'", groups[2], groups)
      return nil, err
    }
    rooms[i] = Room{Encrypted: groups[1], SectorId: id, Checksum: groups[3]}
  }

  return rooms, nil
}
