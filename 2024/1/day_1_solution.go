package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	list1, list2, counts, err := parseInput(string(input))
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		return
	}

	slices.Sort(list1)
	slices.Sort(list2)

	part1 := calculatePart1(list1, list2)
	part2 := calculatePart2(list1, counts)

	fmt.Println(part1)
	fmt.Println(part2)
}

func parseInput(input string) ([]int, []int, map[int]int, error) {
	var list1, list2 []int
	counts := make(map[int]int)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		var n1, n2 int
		_, err := fmt.Sscanf(line, "%d   %d", &n1, &n2)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("invalid line format: %q", line)
		}
		list1 = append(list1, n1)
		list2 = append(list2, n2)
		counts[n2]++
	}

	return list1, list2, counts, nil
}

func calculatePart1(list1, list2 []int) int {
	total := 0
	for i := range list1 {
		total += abs(list2[i] - list1[i])
	}
	return total
}

func calculatePart2(list1 []int, counts map[int]int) int {
	total := 0
	for _, n1 := range list1 {
		total += n1 * counts[n1]
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
