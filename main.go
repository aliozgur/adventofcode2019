package main

import (
	"adventofcode/inputs/input10"
	"adventofcode/solutions/solution10"
	"flag"
)

func main(){

	boolPtr := flag.Bool("p", false, "Print on screen")
	numbPtr := flag.Int("ps", 0, "Slow down print progress")
	flag.Parse()

	options := solution10.AsteroidMapOptions {Print:*boolPtr, PrintSleepMiliseconds:*numbPtr}
	asteroidMap := solution10.NewAsteroidMap(input10.ASTEROIDMAP,&options)
	asteroidMap.Solve()

}
