package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("need day part\n")
		return
	}
	day, err := strconv.ParseUint(os.Args[1], 10, 64)
	check(err)
	part, err:= strconv.ParseUint(os.Args[2], 10, 64)
	check(err)
	if day == 0 || day > 12 {
		fmt.Printf("bad day\n")
		return
	}
	if part == 0 || part > 2 {
		fmt.Printf("bad part\n")
		return
	}
	fmt.Printf("day: %v | part: %v: ", day, part)

	solvers := []func(uint64){One, Two}
	solvers[day - 1](part)
}

