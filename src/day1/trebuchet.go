package day1

import (
	"aoc2023-go/src/util"
	"fmt"
	"log"
	"strconv"
)

func Trebuchet() {
	fileBytes, err := util.ReadFileBytes("sample1.txt")
	if err != nil {
		log.Fatal("An error ocurred while reading the file.\n", err)
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
	fullDigit := ""
	firstDigitFound := false
	secondDigitFound := false

	// Iterate from the start
	for i := 0; i < len(line) && !firstDigitFound; i++ {
		currentValue := string(line[i])
		_, err := strconv.Atoi(currentValue)

		if err == nil {
			firstDigitFound = true
			fullDigit += currentValue
		}
	}

	// Iterate from the end
	for i := len(line) - 1; i >= 0 && !secondDigitFound; i-- {
		currentValue := string(line[i])
		_, err := strconv.Atoi(currentValue)

		if err == nil {
			secondDigitFound = true
			fullDigit += currentValue
		}
	}

	resultDigit, err := strconv.Atoi(fullDigit)
	if err != nil {
		log.Fatal("An error ocurred converting the resulting digit.\n", err)
	}

	return resultDigit
}
