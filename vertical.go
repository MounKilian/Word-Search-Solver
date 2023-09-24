package Correction

import (
	"bufio"
	"log"
	"os"
)

func Vertical(column []rune) []string {
	word := []string{}
	lineStr := ""
	for _, i := range column {
		lineStr += string(i)
	}
	readFile, err := os.Open("words.txt")

	if err != nil {
		log.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		columnWord := fileScanner.Text()
		if len(columnWord) == len(lineStr) && (lineStr[0:len(columnWord)] == columnWord) {
			word = append(word, lineStr[0:len(columnWord)])
		} else {
			for i := 0; i < len(lineStr)-len(columnWord)+1; i++ {
				for j := 0; j < len(columnWord); j++ {
					if lineStr[i] == columnWord[j] {
						if lineStr[i:i+len(columnWord)] == columnWord {
							word = append(word, lineStr[i:i+len(columnWord)])
						}
					}
				}
			}
		}
	}
	return word
}
