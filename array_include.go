package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(dat), "\n")
	target, _ := strconv.Atoi(lines[0])
	numbers := getNumbers(lines[1], target)
	sort.Ints(numbers)
	result := hasElement(numbers, target)
	writeResult(result)
}

func getNumbers(line string, target int) []int {
	numbers_data := strings.Split(line, " ")
	capacity := len(numbers_data)
	numbers := make([]int, 0, capacity)

	for _, number_string := range numbers_data {
		number, _ := strconv.Atoi(number_string)
		if number <= target {
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func hasElement(numbers []int, target int) string {
	size := len(numbers)
	first := 0
	last := size - 1

	for first < last {
		sum := numbers[first] + numbers[last]

		if sum == target {
			return "1"
		} else if target > sum {
			first = first + 1
		} else {
			last = last - 1
		}
	}

	return "0"
}

func writeResult(result string) error {
	d1 := []byte(result)
	return ioutil.WriteFile("output.txt", d1, 0644)
}
