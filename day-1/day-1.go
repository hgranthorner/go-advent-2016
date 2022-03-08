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

type Data struct {
	dir string
	x   int
	y   int
}

func (d Data) get_distance() int {
	abs_x := math.Abs(float64(d.x))
	abs_y := math.Abs(float64(d.y))
	return int(abs_x + abs_y)
}

func (d *Data) move(s string) {
	turn := Stridx(s, 0)
	distance, err := strconv.Atoi(strings.Trim(s, "\n")[1:])

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
}

func PuzzleOne(input string) int {
	d := Data{"N", 0, 0}
	split := strings.Split(input, ", ")
	for _, s := range split {
		d.move(s)
	}

	return d.get_distance()
}

func SayHi() {
	fmt.Println("Hi")
}
