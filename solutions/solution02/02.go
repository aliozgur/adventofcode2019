package solution02

import (
	"adventofcode/inputs/input02"
	"fmt"
	"time"
)
var opcodes []int
func init(){
	fmt.Println("Continue on:",time.Now())
}
func Run(){

	// 100*noun + verb
	result := -1

	for noun :=0; noun <=99; noun++{
		if result >= 0 {
			break
		}
		for verb :=0; verb <=99; verb++{
			opcodes = input02.ReadInput()
			opcodes[1] = noun // noun
			opcodes[2] = verb // verb
			process(0)
			if opcodes[0] == 19690720 {
				result = (100* noun) + verb
				break
			}
		}
	}

	if result >= 0 {
		fmt.Println("result:",result)
	}else{
		fmt.Println("Failed...")
	}
}

func process(opcodeaddress int){
	var opcode = opcodes[opcodeaddress]
	switch opcode {
	case 1:
		vpos1 := opcodes[opcodes[opcodeaddress+1]]
		vpos2 := opcodes[opcodes[opcodeaddress+2]]
		spos  := opcodes[opcodeaddress+3]
		opcodes[spos] = vpos1+vpos2
		process(opcodeaddress +4)
	case 2:
		vpos1 := opcodes[opcodes[opcodeaddress+1]]
		vpos2 := opcodes[opcodes[opcodeaddress+2]]
		spos  := opcodes[opcodeaddress+3]
		opcodes[spos] = vpos1*vpos2
		process(opcodeaddress +4)
	case 99:
		return
	default:
		panic("Unknown operation code!")
	}
}



