package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	l1, l2 := []int{}, []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			a, _ := strconv.Atoi(numbers[0])
			b, _ := strconv.Atoi(numbers[1])
			l1 = append(l1, a)
			l2 = append(l2, b)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	dist := part1(l1, l2)
	fmt.Printf("%d\n", dist)

	similarity := part2(l1, l2)
	fmt.Printf("%d\n", similarity)
}

func part1(l1 []int, l2 []int) int {
	slices.Sort(l1)
	slices.Sort(l2)

	dist := 0
	for i := 0; i < len(l1); i++ {
		dist += int(math.Abs(float64(l1[i]) - float64(l2[i])))
	}

	return dist
}

// This time, you'll need to figure out exactly how often each number from the left list appears in the right list.
// Calculate a total similarity score by adding up each number in the left list after multiplying it by the number of times that number appears in the right list.
func part2(l1 []int, l2 []int) int {
	similarity := 0
	m := make(map[int]int)

	for _, num := range l2 {
		m[num]++
	}
	for _, num := range l1 {
		if count, ok := m[num]; ok {
			similarity += num * count
		}
	}

	return similarity
}
