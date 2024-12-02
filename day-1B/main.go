package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/rad0r/AOC-2024-in-go/pkg/filereader"
)

const FilePath = "../day-1A/puzzle-input.txt"

func splitListsSides(stringSlice []string) ([]int, []int, error) {
	var leftSideList, rightSideList []int

	for _, v := range stringSlice {
		parts := strings.Fields(v)
		intFirstNumber, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		intSecondNumber, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		leftSideList = append(leftSideList, intFirstNumber)
		rightSideList = append(rightSideList, intSecondNumber)
	}
	return leftSideList, rightSideList, nil
}

func findOccurrences(slice1, slice2 []int) map[int]int {
	occurrences := make(map[int]int)

	for _, num := range slice1 {
		count := 0
		for _, val := range slice2 {
			if val == num {
				count++
			}
		}
		if count > 0 {
			occurrences[num] = count
		}
	}

	return occurrences
}

func main() {
	stringSlice := filereader.ReadFileLineByLine(FilePath)

	var leftSideList, rightSideList []int
	leftSideList, rightSideList, err := splitListsSides(stringSlice)
	if err != nil {
		log.Fatal(err)
	}

	occuranceMap := findOccurrences(leftSideList, rightSideList)
	var sum int
	for num, val := range occuranceMap {
		sum += val * num
	}
	print(sum)
}
