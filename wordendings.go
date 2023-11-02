package main

import (
	"fmt"
	"strconv"
)

func lastkDigit(val int) int {
	valString := strconv.Itoa(val)
	lastdidgit := len(valString) - 1

	if len(valString) == 1 {
		return val
	}
	var lastDidgitVal string

	lastDidgitVal = string(valString[lastdidgit])
	lastDidgitNum, _ := strconv.Atoi(lastDidgitVal)

	return lastDidgitNum
}

func secondLastkDigit(val int) int {
	valString := strconv.Itoa(val)
	lastdidgit := len(valString) - 2

	if len(valString) == 1 {
		return val
	}
	var lastDidgitVal string

	lastDidgitVal = string(valString[lastdidgit])
	lastDidgitNum, _ := strconv.Atoi(lastDidgitVal)

	return lastDidgitNum
}

func GetWordWithEnding(root string, endings []string, count int) string {
	answer := root

	// “яблок”, [“о”, “а”, “”]
	// 0 к, 1 ко, 2ка - 4ка, 5 к - 20 к, 21 ко, 22 ка 24 ка, 25
	if secondLastkDigit(count) == 1 {
		answer += endings[2]
	} else {

		switch p := lastkDigit(count); {
		case p == 0:
			answer += endings[2]
		case p == 1:
			answer += endings[0]
		case p >= 2 && p <= 4:
			answer += endings[1]
		case p >= 5 && p <= 19:
			answer += endings[2]
		default:
			answer += endings[2]
		}
	}
	return answer

}

func main() {

	arr := []string{"o", "a", ""}
	fmt.Println(GetWordWithEnding("яблок", arr, 22))

}
