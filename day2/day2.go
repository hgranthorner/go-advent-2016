package day2

import "fmt"

const UP = 'U'
const RIGHT = 'R'
const LEFT = 'L'
const DOWN = 'D'

// 1 2 3
// 4 5 6
// 7 8 9

var pad = map[rune](map[rune]rune){
	'1': {UP: '1', LEFT: '1', RIGHT: '2', DOWN: '4'},
	'2': {UP: '2', LEFT: '1', RIGHT: '3', DOWN: '5'},
	'3': {UP: '3', LEFT: '2', RIGHT: '3', DOWN: '6'},
	'4': {UP: '1', LEFT: '4', RIGHT: '5', DOWN: '7'},
	'5': {UP: '2', LEFT: '4', RIGHT: '6', DOWN: '8'},
	'6': {UP: '3', LEFT: '5', RIGHT: '6', DOWN: '9'},
	'7': {UP: '4', LEFT: '7', RIGHT: '8', DOWN: '7'},
	'8': {UP: '5', LEFT: '7', RIGHT: '9', DOWN: '8'},
	'9': {UP: '6', LEFT: '8', RIGHT: '9', DOWN: '9'},
}

func determinePassword(directions []string, pad map[rune](map[rune]rune)) string {
	result := ""
  num := '5'
	for _, dir := range directions {
		for _, c := range dir {
      num = pad[num][c]
		}
    result = result + fmt.Sprintf("%c", num)
	}

	return result
}

//     1 
//   2 3 4
// 5 6 7 8 9
//   A B C
//     D

var pad2 = map[rune](map[rune]rune){
	'1': {UP: '1', LEFT: '1', RIGHT: '1', DOWN: '3'},
	'2': {UP: '2', LEFT: '2', RIGHT: '3', DOWN: '6'},
	'3': {UP: '1', LEFT: '2', RIGHT: '4', DOWN: '7'},
	'4': {UP: '4', LEFT: '3', RIGHT: '4', DOWN: '8'},
	'5': {UP: '5', LEFT: '5', RIGHT: '6', DOWN: '5'},
	'6': {UP: '2', LEFT: '5', RIGHT: '7', DOWN: 'A'},
	'7': {UP: '3', LEFT: '6', RIGHT: '8', DOWN: 'B'},
	'8': {UP: '4', LEFT: '7', RIGHT: '9', DOWN: 'C'},
	'9': {UP: '9', LEFT: '8', RIGHT: '9', DOWN: '9'},
	'A': {UP: '6', LEFT: 'A', RIGHT: 'B', DOWN: 'A'},
	'B': {UP: '7', LEFT: 'A', RIGHT: 'C', DOWN: 'D'},
	'C': {UP: '8', LEFT: 'B', RIGHT: 'C', DOWN: 'C'},
	'D': {UP: 'B', LEFT: 'D', RIGHT: 'D', DOWN: 'D'},
}
