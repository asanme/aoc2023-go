package day1_extended

import (
	"aoc2023-go/src/util"
	"fmt"
	"log"
	"strconv"
)

var NUMBERS = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func TrebuchetExtended() {
	fileBytes, err := util.ReadFileBytes("sample1.txt")

	if err != nil {
		log.Fatal("An error occurred while reading the file.\n", err)
	}

	calibrationTotal := getCalibrationTotal(fileBytes)
	fmt.Println("The total calibration value is", calibrationTotal)
}

func getCalibrationTotal(fileBytes []byte) int {
	currentLine := ""
	calibrationTotal := 0

	for _, currentByte := range fileBytes {
		currentChar := string(currentByte)

		if currentByte == '\n' {
			calibrationTotal += getCalibrationDetails(currentLine)
			currentLine = ""
		} else {
			currentLine += currentChar
		}
	}

	calibrationTotal += getCalibrationDetails(currentLine)
	return calibrationTotal
}

func getCalibrationDetails(line string) int {
	var numericValues []int
	var wordMatches []string
	firstLetterMatches := false

	for index, currentByte := range line {
		currentChar := string(currentByte)
		value, err := strconv.Atoi(currentChar)
		if err == nil {
			numericValues = append(numericValues, value)
		}

		for key := range NUMBERS {
			firstCharKey := string(key[0])
			if currentChar == firstCharKey {
				wordMatches = append(wordMatches, key)
				firstLetterMatches = true
			}
		}

		if firstLetterMatches {
			for _, word := range wordMatches {
				lineLength := len(line)
				wordLength := len(word)
				totalWordLength := index + wordLength

				if lineLength > wordLength && totalWordLength <= lineLength {
					currentWord := line[index : index+wordLength]
					if currentWord == word {
						actualValue := NUMBERS[word]
						numericValues = append(numericValues, actualValue)
					}
				}
			}
		}

		wordMatches = nil
		firstLetterMatches = false
	}

	firstValue := strconv.Itoa(numericValues[0])
	lastValue := strconv.Itoa(numericValues[len(numericValues)-1])
	mergedValue, err := strconv.Atoi(firstValue + lastValue)
	if err != nil {
		log.Fatal("An error ocurred while converting the digits.\n", err)
	}

	return mergedValue
}
