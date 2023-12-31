package Correction

import (
	"bufio"
	"log"
	"os"
)

func Horizontal(line [10]rune) []string {
	word := []string{}
	lineStr := ""
	for _, i := range line {
		lineStr += string(i)
	}
	readFile, err := os.Open("words.txt")

	if err != nil {
		log.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lineWord := fileScanner.Text()
		if len(lineWord) == len(lineStr) && (lineStr[0:len(lineWord)] == lineWord) {
			word = append(word, lineStr[0:len(lineWord)])
		} else {
			for i := 0; i < len(lineStr)-len(lineWord)+1; i++ {
				for j := 0; j < len(lineWord); j++ {
					if lineStr[i] == lineWord[j] {
						if lineStr[i:i+len(lineWord)] == lineWord {
							word = append(word, lineStr[i:i+len(lineWord)])
						}
					}
				}
			}
		}
	}
	return word
}
