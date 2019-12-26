package solution01

import (
	"adventofcode/inputs/input01"
	"fmt"
	"math"
	"time"
)


func calculateFuelRecursive(mass int64) int64{
	fuel := calculateFuel(mass)
	if fuel <= 0 {
		return 0
	}
	return fuel + calculateFuelRecursive(fuel)
}

func calculateFuel(mass int64) int64{
	return int64(math.Floor(float64(mass) / 3) - 2)
}

func calculateFuelSumRecursive(masses []int64) int64{
	var sum = int64(0)
	for _,mass := range masses{
		sum += calculateFuelRecursive(mass)
	}
	return sum
}

func calculateFuelSum(masses []int64) int64{
	var sum = int64(0)
	for _,mass := range masses{
		sum += calculateFuel(mass)
	}
	return sum
}

func Run(){
	fmt.Println("Continue on:",time.Now())
	var masses = input01.ReadInput()
	var part01 = calculateFuelSum(masses)
	var part02 = calculateFuelSumRecursive(masses)

	fmt.Println("Total Fuel Part 01_01:",part01)
	fmt.Println("Total Fuel Part 01_02:",part02)
}
