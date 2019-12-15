package intcodevm

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	ADD       = 1
	MULTIPLY  = 2
	STORE     = 3
	OUT       = 4
	JUMPTRUE  = 5
	JUMPFALSE = 6
	LESSTHAN  = 7
	EQUALS    = 8
	HALT      = 99
)

type Opcode struct{
	Code       int
	ParamCount int
	Length     int
	ParamMode1 int
	ParamMode2 int
	ParamMode3 int
}


func rebuildProgram(opcodes []int) (program string){

	program = ""
	for i,code := range opcodes{
		if i == len(opcodes)-1{
			program+=strconv.Itoa(code)
		} else{
			program+=strconv.Itoa(code) + ","
		}
	}
	return
}


func Run(program string, inputs[] int)(output int, newProg string){
	cursor := 0
	intCodes := parseInput(program)
	output = -1
	done  := false
	jump := 0

	for{
		jump = 0
		if done {
			break
		}
		opcode := evalOpcode(intCodes[cursor])
		if opcode.Code == HALT{
			break
		}
		switch opcode.Code {
		case ADD, MULTIPLY:
			p1 := intCodes[cursor+1]
			p2 := intCodes[cursor+2]
			storeTo := intCodes[cursor+3]
			v1 := 0
			v2 := 0

			if opcode.ParamMode1 == 0{
				v1 = intCodes[p1]
			} else{
				v1 = p1
			}
			if opcode.ParamMode2 == 0{
				v2 = intCodes[p2]
			} else{
				v2 = p2
			}
			result := addOrMultiply(opcode,v1,v2)
			intCodes[storeTo] = result
			jump = opcode.Length
		case JUMPTRUE, JUMPFALSE:
			p1 := intCodes[cursor+1]
			p2 := intCodes[cursor+2]
			if opcode.ParamMode1 == 0 {
				p1 = intCodes[p1]
			}
			shouldJump := shouldJump(opcode,p1)
			if shouldJump {
				jump = p2
				if opcode.ParamMode2 == 0 {
					jump = intCodes[p2]
				}
				cursor = jump
				jump = 0
			} else {
				jump = opcode.Length
			}

		case LESSTHAN,EQUALS:
			p1 := intCodes[cursor+1]
			p2 := intCodes[cursor+2]
			storeTo := intCodes[cursor+3]

			v1 := 0
			v2 := 0

			if opcode.ParamMode1 == 0{
				v1 = intCodes[p1]
			} else{
				v1 = p1
			}
			if opcode.ParamMode2 == 0{
				v2 = intCodes[p2]
			} else{
				v2 = p2
			}
			result := compare(opcode,v1,v2)
			intCodes[storeTo] = result
			jump = opcode.Length
		case STORE:
			p1 := intCodes[cursor+1]
			intCodes[p1] = inputs[0] // pop input
			if len(inputs) > 1 {
				inputs = inputs[1:len(inputs)] // dequeue
			}
			jump = opcode.Length
		case OUT:
			p1 := intCodes[cursor+1]
			if opcode.ParamMode1 == 0 {
				output = intCodes[p1]
			} else if opcode.ParamMode1 == 1 {
				output = p1
			}
			done = true
			jump = opcode.Length
		}
		cursor = cursor+jump
	}

	newProg = rebuildProgram(intCodes)
	return
}

func parseInput(problem string) (result []int){
	var values = strings.Split(problem,",")
	for _,value := range values {
		if value == "" {
			continue
		}
		v,e := strconv.Atoi(value)
		if e == nil {
			result = append(result, v)
		} else{
			log.Println("Can not convert value ",value)
		}
	}
	return
}

func addOrMultiply(op Opcode, v1, v2 int) (result int){
	switch op.Code {
	case ADD:
		result = v1 + v2
	case MULTIPLY:
		result = v1*v2
	}
	return
}

func shouldJump(op Opcode, v1 int) (result bool){
	result = false

	switch op.Code {
	case JUMPTRUE:
		if v1 != 0{
			result = true
		}
	case JUMPFALSE:
		if v1 == 0{
			result = true
		}
	}
	return
}

func compare(op Opcode, v1,v2 int) (value int){
	switch op.Code {
	case LESSTHAN:
		if v1 < v2{
			value = 1
		} else{
			value = 0
		}
	case EQUALS:
		if v1 == v2{
			value = 1
		} else{
			value = 0
		}
	}
	return
}

func evalOpcode(intOpcode int)(result Opcode){

	op :=  extractDigit(intOpcode,1)
	switch op {
	case ADD, MULTIPLY:
		result = Opcode{
			Code:       op,
			ParamCount: 2,
			Length:     4,
			ParamMode1: extractDigit(intOpcode,3),
			ParamMode2: extractDigit(intOpcode,4),
			ParamMode3: -1,
		}
	case JUMPTRUE, JUMPFALSE:
		result = Opcode{
			Code:       op,
			ParamCount: 2,
			Length:     3,
			ParamMode1: extractDigit(intOpcode,3),
			ParamMode2: extractDigit(intOpcode,4),
			ParamMode3: -1,
		}
	case EQUALS,LESSTHAN:
		result = Opcode{
			Code:       op,
			ParamCount: 2,
			Length:     4,
			ParamMode1: extractDigit(intOpcode,3),
			ParamMode2: extractDigit(intOpcode,4),
			ParamMode3: -1,
		}
	case STORE:
		result = Opcode{
			Code:       op,
			ParamCount: 1,
			Length:     2,
			ParamMode1: extractDigit(intOpcode,3),
			ParamMode2: -1,
			ParamMode3: -1,
		}
	case OUT:
		result = Opcode{
			Code:       op,
			ParamCount: 1,
			Length:     2,
			ParamMode1: extractDigit(intOpcode,3),
			ParamMode2: -1,
			ParamMode3: -1,
		}
	case HALT:
		result = Opcode{
			Code:       op,
			ParamCount: 0,
			Length:     1,
			ParamMode1: -1,
			ParamMode2: -1,
			ParamMode3: -1,
		}
	default:
		panic(fmt.Sprintf("Eval: Invalid operation Code. %d",op))
	}
	return
}

func extractDigit(value int, digit int) int{
	pow10 := int(math.Pow10(digit-1))
	return (value / (pow10)) % 10
}

