// https://adventofcode.com/2020/day/9/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func contains(items []int, item int, excludeIdx int) bool {
	for idx, anItem := range items {
		if anItem == item && idx != excludeIdx {
			return true
		}
	}

	return false
}

func canProduce(preamble []int, target int) bool {
	for i := len(preamble) - 1; i >= 0; i-- {
		firstNum := preamble[i]

		if firstNum < target {
			secondNum := target - firstNum

			if contains(preamble, secondNum, i) {
				return true
			}
		}
	}

	return false
}

func findContiguous(preamble []int, target int) (int, int) {
	smallest := -1
	largest := -1

	for i := len(preamble) - 1; i >= 0; i-- {
		smallest = -1
		largest = -1
		tmpTarget := target

		number := preamble[i]
		if smallest == -1 || number < smallest {
			smallest = number
		}

		if number > largest {
			largest = number
		}

		tmpTarget -= number
		nextIdx := i

		for {
			nextIdx--
			if nextIdx < 0 {
				// Impossible to get the target
				break
			}

			nextNum := preamble[nextIdx]
			tmpTarget -= nextNum

			if tmpTarget >= 0 {
				if smallest == -1 || nextNum < smallest {
					smallest = nextNum
				}

				if nextNum > largest {
					largest = nextNum
				}
			}

			if tmpTarget == 0 {
				return smallest, largest
			} else if tmpTarget < 0 {
				break
			}
		}
	}

	return 0, 0
}

func getPreamble(data []int, size int, current int) []int {
	preamble := make([]int, size)
	copy(preamble, data[current-size:current])

	return preamble
}

func main() {
	fc, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		return
	}

	lines := strings.Split(string(fc), "\n")
	data := make([]int, len(lines))

	for idx, line := range lines {
		data[idx], _ = strconv.Atoi(line)
	}

	// Now let's process the rest
	preambleLength := 25
	for i := preambleLength; i < len(lines); i++ {
		target, _ := strconv.Atoi(lines[i])
		preamble := getPreamble(data, preambleLength, i)

		if !canProduce(preamble, target) {
			fmt.Println(target)

			// Part 2
			smallest, largest := findContiguous(data[0:i], target)
			fmt.Println(smallest, largest)
			fmt.Println(smallest + largest)
		}
	}
}
