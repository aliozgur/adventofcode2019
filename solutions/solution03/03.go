package solution03

import (
	problem "adventofcode/inputs/input03"
	"fmt"
	"golang.org/x/tools/container/intsets"
)


const empty = ""
const left ="←"
const right = "→"
const up = "↑"
const down = "↓"
const intersect="x"

func Run(){
	var m1,_ = problem.ParseWire(problem.Wire1)
	var m2,_ = problem.ParseWire(problem.Wire2)

	max := 12000
	capacity := max*2 + 1
	center := max

	fmt.Println("Start...")

	board1,steps1 := Wire(m1,capacity,center)
	board2,steps2 := Wire(m2,capacity,center)

	fmt.Println("wiring done")

	minDist := intsets.MaxInt
	minSteps := intsets.MaxInt


	fmt.Println("Comparing")
	for i := range board1{
		for j := range board1[i]{
			if i == center && j == center{
				continue
			}

			if board1[i][j] != empty &&  board2[i][j] != empty{
				d := abs(i-center)  + abs(j-center)
				if d < minDist{
					minDist = d
				}

				key := fmt.Sprintf("%d.%d",i,j)
				s := steps1[key]+steps2[key]
				if s < minSteps{
					minSteps = s
				}
			}
		}
	}

	fmt.Println("Min Manhattan Distance:",minDist)
	fmt.Println("Min Total Steps:",minSteps)
}

type intersection struct{
	x int
	y int
	steps int
}

func abs(value int) int{
	if value < 0{
		return value * -1
	}
	return value
}


func Wire(w1 []problem.Movement,capacity int, center int) (result [][]string, steps map[string]int)  {
	result = make([][]string,capacity)

	// Initialize columns
	for i := range result {
		result[i] = make([]string,capacity)
	}

	result[center][center] = "O"

	steps = wire(w1,center,result)
	return
}

func wire(wiring []problem.Movement,center int,result [][]string) map[string]int{
	row := center
	col := center
	steps := 0
	intersection := make(map[string]int)

	for _,m := range wiring {
		switch m.Pos {
		case "U":
			for i:=row-1;i >= row-m.Steps;i-- {
				steps++
				intersection[fmt.Sprintf("%d.%d",i,col)] = steps
				result[i][col] = up
			}
			row = row-m.Steps
		case "D":
			for i:=row+1;i <= row+m.Steps;i++ {
				steps++
				intersection[fmt.Sprintf("%d.%d",i,col)] = steps
				result[i][col] = down

			}
			row = row+m.Steps
		case "L":
			for i:=col-1;i >= col - m.Steps;i-- {
				steps++
				intersection[fmt.Sprintf("%d.%d",row,i)] = steps
				result[row][i] = left
			}

			col = col - m.Steps
		case "R":
			for i:=col+1;i <= col + m.Steps;i++ {
				steps++
				intersection[fmt.Sprintf("%d.%d",row,i)] = steps
				result[row][i] = right
			}
			col = col + m.Steps
		default:
			panic("Unknown direction!!!")
		}
	}
	return intersection
}

