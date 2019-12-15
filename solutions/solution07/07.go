package solution07

import (
	"adventofcode/intcodevm"
	"fmt"
	"math"
	"strings"
)

const program = "3,8,1001,8,10,8,105,1,0,0,21,38,59,84,97,110,191,272,353,434,99999,3,9,1002,9,2,9,101,4,9,9,1002,9,2,9,4,9,99,3,9,102,5,9,9,1001,9,3,9,1002,9,5,9,101,5,9,9,4,9,99,3,9,102,5,9,9,101,5,9,9,1002,9,3,9,101,2,9,9,1002,9,4,9,4,9,99,3,9,101,3,9,9,1002,9,3,9,4,9,99,3,9,102,5,9,9,1001,9,3,9,4,9,99,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,99"
//const program = "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
func Run(){
	phases  := []int{5, 6, 7, 8, 9 }
	posssiblePhases := quicPerm(phases)
	fmt.Println(posssiblePhases)
	max := math.MinInt32
	for _, phase := range posssiblePhases{
		fmt.Println(strings.Repeat("-",40))
		fmt.Println("Running phase:", phase)
		output := test1(phase)
		fmt.Println("Done phase. Output is:", output)
		if output > max{
			max = output
		}
		fmt.Println("Current max is:", max)

	}
	fmt.Println(strings.Repeat("-",40))
	fmt.Println("Max thruster signal is:", max)
}
func part1(){
	phases  := []int{0, 1, 2, 3, 4 }
	posssiblePhases := quicPerm(phases)
	fmt.Println(posssiblePhases)
	max := math.MinInt32
	for _, phase := range posssiblePhases{
		fmt.Println(strings.Repeat("-",40))
		fmt.Println("Running phase:", phase)
		output := test2(phase)
		fmt.Println("Done phase. Output is:", output)
		if output > max{
			max = output
		}
		fmt.Println("Current max is:", max)

	}
	fmt.Println(strings.Repeat("-",40))
	fmt.Println("Max thruster signal is:", max)
}

func test1(phase []int) int {
	out1, _ := intcodevm.Run(program,[]int{phase[0],0})
	out2, _ := intcodevm.Run(program,[]int{phase[1],out1})
	out3, _ := intcodevm.Run(program,[]int{phase[2],out2})
	out4, _ := intcodevm.Run(program,[]int{phase[3],out3})
	out5, _ := intcodevm.Run(program,[]int{phase[4],out4})
	return out5
}

func test2(phase []int) int {
	out1, _ := intcodevm.Run(program,[]int{phase[0],0})
	out2, _ := intcodevm.Run(program,[]int{phase[1],out1})
	out3, _ := intcodevm.Run(program,[]int{phase[2],out2})
	out4, _ := intcodevm.Run(program,[]int{phase[3],out3})
	out5, _ := intcodevm.Run(program,[]int{phase[4],out4})
	return out5
}

func factorial(n int) int{
	if n == 0{
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
func quicPerm(a []int) (result [][]int){
	N := len(a)
	p := make([]int,N)
	i := 1
	result = make([][]int,factorial(N))
	result[0] = append(result[0],a...)

	cnt := 1
	for i < N {
		if p[i] < i {
			j := 0
			if i % 2 == 1{
				j = p[i]
			}
			x, y := a[j],a[i]
			a[i] = x
			a[j] = y
			p[i]++
			i = 1
			result[cnt] = append(result[cnt],a...)
			cnt++
		} else if p[i] == i{
			p[i] = 0
			i++
		}
	}

	return
}
