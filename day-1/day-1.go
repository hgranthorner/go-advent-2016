package day_1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Stridx(s string, i int) string {
	return string([]byte{s[i]})
}

type TupleSet struct {
	vals map[int]map[int]bool
}

func (s TupleSet) get(x int, y int) bool {
	return s.vals[x][y]
}

func (s *TupleSet) put(x int, y int) {
	if len(s.vals) == 0 {
		s.vals = make(map[int]map[int]bool)
	}
	if len(s.vals[x]) == 0 {
		s.vals[x] = map[int]bool{y: true}
	} else {
		s.vals[x][y] = true
	}
}

func Makeline(minx int, maxx int, miny int, maxy int) [][]int {
	res := make([][]int, 0)
	if minx != maxx {
		step := int((maxx - minx) / int(math.Abs(float64(maxx-minx))))
		if step > 0 {
			for x := minx + step; x <= maxx; x += step {
				res = append(res, []int{x, miny})
			}
		} else {
			for x := minx + step; x >= maxx; x += step {
				res = append(res, []int{x, miny})
			}
		}
	}

	if miny != maxy {
		step := int((maxy - miny) / int(math.Abs(float64(maxy-miny))))
		if step > 0 {
			for y := miny + step; y <= maxy; y += step {
				res = append(res, []int{minx, y})
			}
		} else {
			for y := miny + step; y >= maxy; y += step {
				res = append(res, []int{minx, y})
			}
		}
	}

	return res
}

type Data struct {
	dir       string
	x         int
	y         int
	locations TupleSet
}

func abs_plus(n1 int, n2 int) int {
	abs_n1 := math.Abs(float64(n1))
	abs_n2 := math.Abs(float64(n2))
	return int(abs_n1 + abs_n2)
}

func (d Data) get_distance() int {
	return abs_plus(d.x, d.y)
}

func (d *Data) move(s string) int {
	turn := Stridx(s, 0)
	distance, err := strconv.Atoi(strings.Trim(s, "\n")[1:])
	minx := d.x
	miny := d.y

	if err != nil {
		panic(fmt.Sprintf("Invalid input! %s. Length: %d", s, len(strings.Trim(s, " "))))
	}

	if turn == "R" {
		switch d.dir {
		case "N":
			d.dir = "E"
		case "E":
			d.dir = "S"
		case "S":
			d.dir = "W"
		case "W":
			d.dir = "N"
		}
	} else {
		switch d.dir {
		case "N":
			d.dir = "W"
		case "W":
			d.dir = "S"
		case "S":
			d.dir = "E"
		case "E":
			d.dir = "N"
		}
	}

	switch d.dir {
	case "E":
		d.x += distance
	case "W":
		d.x -= distance
	case "S":
		d.y += distance
	case "N":
		d.y -= distance
	}

	maxx := d.x
	maxy := d.y

	for _, pos := range Makeline(minx, maxx, miny, maxy) {
		if d.locations.get(pos[0], pos[1]) {
			return abs_plus(pos[0], pos[1])
		}

		d.locations.put(pos[0], pos[1])
	}

	return -1
}

func PuzzleOne(input string) int {
	d := Data{"N", 0, 0, TupleSet{}}
	d.locations.put(0, 0)
	split := strings.Split(input, ", ")
	for _, s := range split {
		d.move(s)
	}

	return d.get_distance()
}

func PuzzleTwo(input string) int {
	d := Data{"N", 0, 0, TupleSet{}}
	d.locations.put(0, 0)
	split := strings.Split(input, ", ")

	for _, s := range split {
		result := d.move(s)
		if result != -1 {
			return result
		}
	}

	fmt.Printf("Locations: %+v\n", d.locations)

	panic("Woops!")
}

func SayHi() {
	fmt.Println("Hi")
}
