package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const filePath string = "./list.txt"

func main() {
	list1, list2, err := readList(filePath)

	if err != nil {
		log.Fatal(err)
	}

	diff := calculateDiff(list1, list2)
	similarity := calculateSimilarity(list1, list2)

	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Similarity: %d\n", similarity)
}

func readList(path string) ([]int, []int, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, nil, fmt.Errorf("%w\n", err)
	}

	defer file.Close()

	var list1, list2 []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		n1, n2, err := parseNums(line)

		if err != nil {
			return nil, nil, fmt.Errorf("%w\n", err)
		}

		list1 = append(list1, n1)
		list2 = append(list2, n2)

	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("%w\n", err)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	return list1, list2, nil
}

func parseNums(line string) (int, int, error) {
	parts := strings.Fields(line)

	n1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("%w", err)
	}

	n2, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("%w", err)
	}

	return n1, n2, nil
}

func calculateDiff(list1 []int, list2 []int) int {
	var answer int

	for i, n := range list1 {
		diff := n - list2[i]

		if diff < 0 {
			diff *= -1
		}

		answer += diff
	}

	return answer
}

func calculateSimilarity(l1 []int, l2 []int) int {
	var similarity int

	for _, n1 := range l1 {
		if i, found := slices.BinarySearch(l2, n1); found {
			similarity += n1

			for j := i + 1; j < len(l2) && l2[j] == n1; j++ {
				similarity += n1
			}
		}
	}

	return similarity
}
