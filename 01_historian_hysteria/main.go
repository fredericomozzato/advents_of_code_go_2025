package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const filePath string = "./list.txt"
const totalLines int = 1000

func main() {
	list1, list2, err := readList(filePath)

	if err != nil {
		log.Fatal(err)
	}

	// printSideBySide(list1, list2)

	diff := calculateDiff(list1, list2)
	similarity := calculateSimilarity(list1, list2)

	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Similarity score: %d\n", similarity)
}

func readList(path string) ([]int, []int, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	list1 := make([]int, 0, totalLines)
	list2 := make([]int, 0, totalLines)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		n1, n2, err := parseNums(line)

		if err != nil {
			return nil, nil, err
		}

		list1 = append(list1, n1)
		list2 = append(list2, n2)

	}

	return list1, list2, nil
}

func parseNums(line string) (int, int, error) {
	strings := strings.Split(line, "   ")

	n1, err := strconv.Atoi(strings[0])
	if err != nil {
		return 0, 0, err
	}

	n2, err := strconv.Atoi(strings[1])
	if err != nil {
		return 0, 0, err
	}

	return n1, n2, nil
}

func calculateDiff(list1 []int, list2 []int) int {
	var answer int

	sort.Ints(list1)
	sort.Ints(list2)

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

	start := time.Now()
	for _, i := range l1 {
		var count int

		for _, j := range l2 {
			if i == j {
				count++
			}
		}
		similarity += i * count
	}

	elapsed := time.Since(start)
	// around 570µs
	fmt.Println(elapsed)

	return similarity
}

func printSideBySide(list1 []int, list2 []int) {
	sort.Ints(list1)
	sort.Ints(list2)

	for i, n := range list1 {
		fmt.Printf("%d   %d\n", n, list2[i])
	}

}
