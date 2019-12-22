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
	// (23,19)
	/*
	var x1,y1,x2,y2 float64 = 1,1,2,2

	degree := math.Mod(math.Atan2(-1*(y2-y1), x2-x1),2 * math.Pi)
	fmt.Println("Degree",degree)
	 */
}
