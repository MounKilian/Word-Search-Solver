package Correction

import (
	"bufio"
	"log"
	"os"
)

func diagonalUp(diag []rune) []string {
	word := []string{}
	lineStr := ""
	for _, i := range diag {
		lineStr += string(i)
	}
	readFile, err := os.Open("words.txt")

	if err != nil {
		log.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		diagonalup := fileScanner.Text()
		if len(diagonalup) == len(lineStr) && (lineStr[0:len(diagonalup)] == diagonalup) {
			word = append(word, lineStr[0:len(diagonalup)])
		} else {
			for i := 0; i < len(lineStr)-len(diagonalup)+1; i++ {
				for j := 0; j < len(diagonalup); j++ {
					if lineStr[i] == diagonalup[j] {
						if lineStr[i:i+len(diagonalup)] == diagonalup {
							word = append(word, lineStr[i:i+len(diagonalup)])
						}
					}
				}
			}
		}
	}
	return word
}
