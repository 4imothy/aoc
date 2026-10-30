package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Two(part uint64) {
	content, err := os.ReadFile("input.2")
	sum := uint64(0)
	check(err)
	ranges := strings.SplitSeq(strings.TrimSpace(string(content)), ",")
	for r := range ranges {
		low, high, _ := strings.Cut(r, "-")
		l, err := strconv.ParseUint(low, 10, 64)
		check(err)
		h, err := strconv.ParseUint(high, 10, 64)
		check(err)
		out:
		for num := l; num <= h; num += 1 {
			length := uint64(1)
			for x := num; x >= 10; x /= 10 {
				length += 1
			}
			digits := make([]uint64, length)
			for x := uint64(0); x < length; x += 1 {
				digits[x] = (num / uint64(math.Pow10(int(length - x - 1)))) % 10
			}
			if part == 1 {
				if length % 2 != 0 {
					continue out
				}
				for x := uint64(0); x < length / 2; x += 1 {
					if digits[x] != digits[x + length / 2] {
						continue out
					}
				}
				sum += num
			}
			if part == 2 {
				in:
				for x := uint64(1); x <= length / 2; x += 1 {
					if length % x != 0 {
						continue in
					}
					groups := length / x
					vals := make([][]uint64, groups)
					for y := uint64(0); y < groups; y += 1 {
						vals[y] = make([]uint64, x)
					}

					for y := uint64(0); y < groups; y += 1 {
						for z := uint64(0); z < x; z += 1 {
							vals[y][z] = digits[y * x + z]
						}
					}
					for y := uint64(0); y < groups - 1; y += 1 {
						for z := uint64(0); z < x; z += 1 {
							if vals[y][z] != vals[y+1][z] {
								continue in
							}
						}
					}
					sum += num
					continue out
				}
			}
		}
	}
	fmt.Printf("%v\n", sum)
}

