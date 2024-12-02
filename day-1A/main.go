package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filePath := "puzzle-input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var leftSideList []int
	var rightSideList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 2 {
			intFirstNumber, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error converting number:", err)
				return
			}
			intSecondNumber, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error converting number:", err)
				return
			}

			leftSideList = append(leftSideList, intFirstNumber)
			rightSideList = append(rightSideList, intSecondNumber)

			sort.Ints(leftSideList)
			sort.Ints(rightSideList)

		}
	}

	var distance int
	for i, v := range leftSideList {
		fmt.Println("index: ", i, "value left: ", v, "value right: ", rightSideList[i])
		if v == rightSideList[i] {
			continue
		}
		if v > rightSideList[i] {
			distance += v - rightSideList[i]
			continue
		}
		if v < rightSideList[i] {
			distance += rightSideList[i] - v
			continue
		}
	}
	fmt.Println("final distance: ", distance)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
