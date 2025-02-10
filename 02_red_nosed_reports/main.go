package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const filePath = "input.txt"

func main() {
	reports, err := parse(filePath)

	if err != nil {
		log.Fatal(err)
	}

	safeReports := checkReports(reports)

	fmt.Printf("Safe reports: %d\n", safeReports)
}

func parse(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var nums []int
		line := scanner.Text()
		parts := strings.Fields(line)

		for _, p := range parts {
			n, err := strconv.Atoi(p)

			if err != nil {
				return nil, err
			}

			nums = append(nums, n)
		}

		reports = append(reports, nums)
	}

	return reports, nil
}

func checkReports(reports [][]int) int {
	var result int

	for _, r := range reports {
		if r[0] > r[1] {
			if !descending(r) {
				continue
			}

			if !acceptedInterval(r) {
				continue
			}

			result++
		} else if r[0] < r[1] {
			if !ascending(r) {
				continue
			}

			if !acceptedInterval(r) {
				continue
			}

			result++
		} else {
			continue
		}
	}

	return result
}

func ascending(r []int) bool {
	for i := 1; i < len(r)-1; i++ {
		if r[i] >= r[i+1] {
			return false
		}
	}

	return true
}

func descending(r []int) bool {
	for i := 1; i < len(r)-1; i++ {
		if r[i] <= r[i+1] {
			return false
		}
	}

	return true
}

func acceptedInterval(r []int) bool {
	for i := 0; i < len(r)-1; i++ {
		diff := (r[i] - r[i+1]) * -1

		if diff > 3 {
			return false
		}
	}

	return true
}
