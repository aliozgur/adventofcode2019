package main

import (
	"adventofcode/inputs/input10"
	"adventofcode/solutions/solution10"
	"fmt"
)

func main(){

	// Best is 11,13 with 210 other asteroids detected
	asteroidMap := solution10.NewAsteroidMap(input10.ASTEROIDMAP3)
	fmt.Println("Number of asteroids:",len(asteroidMap.Asteroids))
	//fmt.Println("Asteroids:",asteroidMap)

	max,asteroid :=	asteroidMap.Solve()
	fmt.Println("Max",max)
	fmt.Println("Max Point",asteroid)

	/*

		p1 := solution10.Asteroid{X:9,Y:10}
		p2 := solution10.Asteroid{X:8,Y:10}

		p3 := solution10.Asteroid{X:7,Y:10}
		p4 := solution10.Asteroid{X:5,Y:10}

		lf := p1.LineFunc(p2)
		fmt.Println("Is on line:",lf(p3))
		fmt.Println("Is on line:",lf(p4))
*/
}
