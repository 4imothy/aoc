package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func One(part uint64) {
	pass := 0
	val := 50
	file, err := os.Open("input.1")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		sign := 1
		if line[0] == 'L' {
			sign = -1
		}

		magnitude, err := strconv.ParseUint(line[1:], 10, 64)
		check(err)
		pos := val > 0
		val += sign * int(magnitude)

		if part == 2 {
			if val > 0 {
				pass += val / 100
			} else {
				pass += -1 * val / 100
				if pos {
					pass += 1
				}
			}
		}

		val = ((val % 100) + 100) % 100

		if part == 1 {
			if val == 0 {
				pass += 1
			}
		}
	}
	fmt.Printf("%v\n", pass)
}

