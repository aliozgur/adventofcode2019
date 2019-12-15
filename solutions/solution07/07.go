package solution07

import (
	"adventofcode/intcodevm"
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

const program = "3,8,1001,8,10,8,105,1,0,0,21,38,59,84,97,110,191,272,353,434,99999,3,9,1002,9,2,9,101,4,9,9,1002,9,2,9,4,9,99,3,9,102,5,9,9,1001,9,3,9,1002,9,5,9,101,5,9,9,4,9,99,3,9,102,5,9,9,101,5,9,9,1002,9,3,9,101,2,9,9,1002,9,4,9,4,9,99,3,9,101,3,9,9,1002,9,3,9,4,9,99,3,9,102,5,9,9,1001,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,99"
// 63103596
//const program = "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"
// 18216
//const program  ="3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
// 139629729

func Run() {
	utils.DEBUG = false
	part1()
	part2()
	utils.DEBUG = true
}

func part1(){
	phases  := []int{0, 1, 2, 3, 4 }
	posssiblePhases := quicPerm(phases)
	//utils.Println(posssiblePhases)
	max := math.MinInt32
	for _, phase := range posssiblePhases{
		utils.Println(strings.Repeat("-",40))
		utils.Println("Running phase:", phase)
		output := evalPhaseForPart1(phase)
		utils.Println("Done phase. Output is:", output)
		if output > max{
			max = output
		}
		utils.Println("Current max is:", max)

	}
	fmt.Println(strings.Repeat("-",40))
	fmt.Println("Day 7 Part 1: Max thruster signal is:", max)
}

func evalPhaseForPart1(phase []int) int {
	out1, _, _, _ := intcodevm.Run(program,[]int{phase[0],0},0)
	out2, _, _, _ := intcodevm.Run(program,[]int{phase[1],out1},0)
	out3, _, _, _ := intcodevm.Run(program,[]int{phase[2],out2},0)
	out4, _, _, _ := intcodevm.Run(program,[]int{phase[3],out3},0)
	out5, _, _, _ := intcodevm.Run(program,[]int{phase[4],out4},0)
	return out5
}

func part2(){
	phases  := []int{5,6,7,8,9}
	posssiblePhases := quicPerm(phases)
	//utils.Println(posssiblePhases)
	max := math.MinInt32
	for _, phase := range posssiblePhases{
		utils.Println(strings.Repeat("-",40))
		utils.Println("Running phase:", phase)
		output := evalPhaseForPart2(program,phase)
		utils.Println("Done phase. Output is:", output)
		if output > max{
			max = output
		}
		utils.Println("Current max is:", max)

	}
	fmt.Println(strings.Repeat("-",40))
	fmt.Println("Day 7 Part 2: Max thruster signal is:", max)
}

func evalPhaseForPart2(program string, phase []int) (output int) {

	programStates := make([]string, 5)
	for i := range programStates {
		programStates[i] = program
	}

	isFirstRun := true

	var out1, out2, out3, out4, out5 int = 0, 0, 0, 0, 0
	var p1, p2, p3, p4, p5  = "", "", "", "", ""
	var c1, c2, c3, c4, c5 = 0, 0, 0, 0, 0
	var halted = false

	for !halted {
		if isFirstRun {
			out1, p1, c1, _ = intcodevm.Run(programStates[0], []int{phase[0], 0}, c1)
			out2, p2, c2, _ = intcodevm.Run(programStates[1], []int{phase[1],out1}, c2)
			out3, p3, c3, _ = intcodevm.Run(programStates[2], []int{phase[2],out2}, c3)
			out4, p4, c4, _ = intcodevm.Run(programStates[3], []int{phase[3],out3}, c4)
			out5, p5, c5, halted = intcodevm.Run(programStates[4], []int{phase[4],out4}, c5)
		} else {
			out1, p1, c1, _ = intcodevm.Run(programStates[0], []int{out5}, c1)
			out2, p2, c2, _ = intcodevm.Run(programStates[1], []int{out1}, c2)
			out3, p3, c3, _ = intcodevm.Run(programStates[2], []int{out2}, c3)
			out4, p4, c4, _ = intcodevm.Run(programStates[3], []int{out3}, c4)
			out5, p5, c5, halted = intcodevm.Run(programStates[4], []int{out4}, c5)
		}

		if halted{
			break
		}

		isFirstRun = false

		programStates[0] = p1
		programStates[1] = p2
		programStates[2] = p3
		programStates[3] = p4
		programStates[4] = p5

		output = out5
	}
	return
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

/*
// https://quickperm.org/
The Counting QuickPerm Algorithm:
   let a[] represent an arbitrary list of objects to permute
   let N equal the length of a[]
   create an integer array p[] of size N to control the iteration
   initialize p[0] to 0, p[1] to 0, p[2] to 0, ..., and p[N-1] to 0
   initialize index variable i to 1
   while (i < N) do {
      if (p[i] < i) then {
         if i is odd, then let j = p[i] otherwise let j = 0
         swap(a[j], a[i])
         increment p[i] by 1
         let i = 1 (reset i to 1)
      } // end if
      else { // (p[i] equals i)
         let p[i] = 0 (reset p[i] to 0)
         increment i by 1
      } // end else (p[i] equals i)
   } // end while (i < N)
*/
func quicPerm(a []int) (result [][]int) {
	N := len(a)
	p := make([]int, N)
	i := 1
	result = make([][]int, factorial(N))
	result[0] = append(result[0], a...)

	cnt := 1
	for i < N {
		if p[i] < i {
			j := 0
			if i%2 == 1 {
				j = p[i]
			}
			x, y := a[j], a[i]
			a[i] = x
			a[j] = y
			p[i]++
			i = 1
			result[cnt] = append(result[cnt], a...)
			cnt++
		} else if p[i] == i {
			p[i] = 0
			i++
		}
	}

	return
}
