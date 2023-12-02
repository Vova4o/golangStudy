package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readTextFile() []string {

	//Pass the file name to the ReadFile() function from
	//the ioutil package to get the content of the file.

	content, error := os.ReadFile("c:\\russian_nouns.txt")

	// defer .Close()
	// Check whether the 'error' is nil or not. If it
	//is not nil, then print the error and exit.
	if error != nil {

		log.Fatal(error)
	}

	// convert the content into a string
	str := string(content)

	arr := strings.Split(str, "\n")
	//Print the string str
	return arr
}

func findaMatch(arr []string, letter string) []string {

	newArr := []string{}

	for _, val := range arr {
		if strings.Contains(val, letter) {
			newArr = append(newArr, val)
		}
	}

	return newArr
}

func shrinkByLen(arr []string, a int) []string {

	newArr := []string{}

	for _, val := range arr {
		if len(val) <= a {
			newArr = append(newArr, val)
		}
	}

	return newArr

}

// func letterOrNumber(s string)  {

// }

func main() {
	arr := readTextFile()

	fmt.Println(len(arr))
	// var letter string
	// fmt.Scanf("%s", &letter)

	// fmt.Println(string(letter))

	// fmt.Println(findaMatch(arr, letter))

	endProgram := false

	for !endProgram {
		var letter string
		fmt.Println("Если хотите закончить программу, напишите - конец")
		fmt.Println("Если вы хотите ограничить длину поиска, напишите - число")
		fmt.Println("Enter letter: ")
		fmt.Scanf("%s", &letter)
		if letter == "конец" {
			endProgram = true
		}
		if letter == "число" {
			var a int
			fmt.Println("Введите число: ")
			fmt.Scanf("%d", &a)
			arr = shrinkByLen(arr, a)
		}
		// arr = shrinkByLen(arr, len(letter))
		arr = findaMatch(arr, letter)
		fmt.Println(arr)
		fmt.Println(len(arr))
	}
}
