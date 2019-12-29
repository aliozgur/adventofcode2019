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

func part1() {
	phases := []int{0, 1, 2, 3, 4}
	posssiblePhases := quicPerm(phases)
	max := math.MinInt32
	for _, phase := range posssiblePhases {
		utils.Println(strings.Repeat("-", 40))
		utils.Println("Running phase:", phase)
		output := evalPhaseForPart1(phase)
		utils.Println("Done phase. Output is:", output)
		if output > max {
			max = output
		}
		utils.Println("Current max is:", max)

	}
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Day 7 Part 1: Max thruster signal is:", max)
}

func evalPhaseForPart1(phase []int) int {
	vm1 := intcodevm.IntcodeVm{OutputMode: intcodevm.OUTPUT_STOPONFIRST, WaitForInput: true}
	vm2 := intcodevm.IntcodeVm{OutputMode: intcodevm.OUTPUT_STOPONFIRST, WaitForInput: true}
	vm3 := intcodevm.IntcodeVm{OutputMode: intcodevm.OUTPUT_STOPONFIRST, WaitForInput: true}
	vm4 := intcodevm.IntcodeVm{OutputMode: intcodevm.OUTPUT_STOPONFIRST, WaitForInput: true}
	vm5 := intcodevm.IntcodeVm{OutputMode: intcodevm.OUTPUT_STOPONFIRST, WaitForInput: true}

	vm1.RunProgram(program, []int{phase[0],0})
	vm2.RunProgram(program, []int{phase[1], vm1.Output[0]})
	vm3.RunProgram(program, []int{phase[2], vm2.Output[0]})
	vm4.RunProgram(program, []int{phase[3], vm3.Output[0]})
	vm5.RunProgram(program, []int{phase[4], vm4.Output[0]})

	return vm5.Output[0]
}

func part2() {
	//utils.DEBUG = true
	phases := []int{5, 6, 7, 8, 9}
	posssiblePhases := quicPerm(phases)
	max := math.MinInt32
	for _, phase := range posssiblePhases {
		utils.Println(strings.Repeat("-", 40))
		utils.Println("Running phase:", phase)
		output := evalPhaseForPart2(program, phase)
		utils.Println("Done phase. Output is:", output)
		if output > max {
			max = output
		}
		utils.Println("Current max is:", max)

	}
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Day 7 Part 2: Max thruster signal is:", max)
}

func evalPhaseForPart2(program string, phase []int) (output int) {

	isFirstRun := true
	vm1 := intcodevm.IntcodeVm{OutputMode: intcodevm.OUTPUT_STOPONFIRST, WaitForInput: true}
	vm2 := intcodevm.IntcodeVm{OutputMode: intcodevm.OUTPUT_STOPONFIRST, WaitForInput: true}
	vm3 := intcodevm.IntcodeVm{OutputMode: intcodevm.OUTPUT_STOPONFIRST, WaitForInput: true}
	vm4 := intcodevm.IntcodeVm{OutputMode: intcodevm.OUTPUT_STOPONFIRST, WaitForInput: true}
	vm5 := intcodevm.IntcodeVm{OutputMode: intcodevm.OUTPUT_STOPONFIRST, WaitForInput: true}

	var halted = false

	for !halted {
		if isFirstRun {
			vm1.RunProgram(program, []int{phase[0],0})
			vm2.RunProgram(program, []int{phase[1], vm1.Output[0]})
			vm3.RunProgram(program, []int{phase[2], vm2.Output[0]})
			vm4.RunProgram(program, []int{phase[3], vm3.Output[0]})
			vm5.RunProgram(program, []int{phase[4], vm4.Output[0]})
			halted = vm5.Halted
		} else {
			vm1.Continue([]int{vm5.Output[0]})
			if len(vm1.Output) > 0 {
				vm2.Continue([]int{vm1.Output[0]})
			}
			if len(vm2.Output) > 0 {
				vm3.Continue([]int{vm2.Output[0]})
			}
			if len(vm3.Output) > 0 {
				vm4.Continue([]int{vm3.Output[0]})
			}
			if len(vm4.Output) > 0 {
				vm5.Continue([]int{vm4.Output[0]})
			}
			halted = vm5.Halted
		}

		if halted {
			break
		}

		isFirstRun = false

		output = vm5.Output[0]
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
