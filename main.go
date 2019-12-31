package main

import (
	"adventofcode/solutions/solution13"
	"flag"
)

func main() {

	partPtr := flag.Int("p", 0, "Part number")
	flag.Parse()
	if *partPtr == 0 || *partPtr == 1 {
		solution13.Part1()
	} else if *partPtr == 2 {
		solution13.Part2()
	}
}
