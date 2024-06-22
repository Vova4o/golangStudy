package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// readTextFile reads the content of a text file and returns it as array of strings.
func readTextFile() []string {
	// Pass the file name to the ReadFile() function from
	// the ioutil package to get the content of the file.

	content, error := os.ReadFile("russian_nouns.txt")

	if error != nil {
		log.Fatal(error)
	}

	// convert the content into a string
	str := string(content)

	arr := strings.Split(str, "\n")
	// Print the string str
	return arr
}

// findaMatch finds a match in an array of strings and returns a new array of strings.
func findMatch(arr []string, letter string) []string {
	newArr := []string{}

	for _, val := range arr {
		if strings.Contains(val, letter) {
			newArr = append(newArr, val)
		}
	}

	return newArr
}

// shrinkByLen returns only those strings from an array of strings that are shorter than a given length.
func shrinkByLen(arr []string, a int) []string {
	newArr := []string{}

	for _, val := range arr {
		// fmt.Println("Len(val)", len(val))
		// fmt.Println("a", a)
		if len(val) <= a {
			newArr = append(newArr, val)
		}
	}

	return newArr
}

func main() {
	arr := readTextFile()

	fmt.Println(len(arr))

	endProgram := false

	for !endProgram {
		var letter string
		fmt.Println("Если хотите закончить программу, напишите - конец")
		fmt.Println("Если вы хотите ограничить длину поиска, напишите - число")
		fmt.Println("Введите букву: ")
		fmt.Scanf("%s\n", &letter)
		if letter == "конец" {
			endProgram = true
		}
		if letter == "число" {
			var a int
			fmt.Println("Введите число: ")
			fmt.Scanf("%d\n", &a)
			arr = shrinkByLen(arr, a)
		}
		// arr = shrinkByLen(arr, len(letter))
		arr = findMatch(arr, letter)
		fmt.Println(arr)
		fmt.Println(len(arr))
	}
}
