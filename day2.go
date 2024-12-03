package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	file, err := os.Open("in2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ans1 := 0
	ans2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		report := make([]int, len(numbers))
		for i, num := range numbers {
			n, _ := strconv.Atoi(num)
			report[i] = n
		}
		safe1 := day2A(report)
		safe2 := day2B(report)
		if safe1 {
			ans1++
		}
		if safe2 {
			ans2++
		}
	}

	fmt.Printf("%d\n", ans1)
	fmt.Printf("%d\n", ans2)
}

func day2A(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	incFlag := false
	if report[1] > report[0] {
		incFlag = true
	}

	for i := 1; i < len(report); i++ {
		if incFlag {
			if report[i] <= report[i-1] {
				return false
			}
		} else {
			if report[i] >= report[i-1] {
				return false
			}
		}
		diff := int(math.Abs(float64(report[i] - report[i-1])))
		if diff > 3 {
			return false
		}
	}

	return true
}

func day2B(report []int) bool {
	if safe := day2A(report); safe {
		return true
	}

	for i := 0; i < len(report); i++ {
		newReport := []int{}
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if safe := day2A(newReport); safe {
			return true
		}
	}

	return false
}
