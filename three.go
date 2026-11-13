package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Three(part uint64) {
	file, err := os.Open("input.3")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := uint64(0)
	count := 2
	if part == 2 {
		count = 12
	}
	for scanner.Scan() {
		line := scanner.Text()
		cur := uint64(0)
		left := uint64(0)

		for i := 0; i < count; i += 1 {
			max := uint64(0)
			id := uint64(0)
			for j, c := range line[left:len(line) - ((count - i) - 1)]{
				num := uint64(c - '0')
				if num > max {
					max = num
					id = uint64(j)
				}
			}
			left += id + 1
			cur += max * uint64(math.Pow10((count - i) - 1))
		}
		total += cur
	}
	fmt.Println(total)
}
