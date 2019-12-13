package solution05

import (
	"fmt"
	tm "github.com/buger/goterm"
	"log"
	"math"
	"strconv"
	"strings"
)

const problem ="3,225,1,225,6,6,1100,1,238,225,104,0,2,171,209,224,1001,224,-1040,224,4,224,102,8,223,223,1001,224,4,224,1,223,224,223,102,65,102,224,101,-3575,224,224,4,224,102,8,223,223,101,2,224,224,1,223,224,223,1102,9,82,224,1001,224,-738,224,4,224,102,8,223,223,1001,224,2,224,1,223,224,223,1101,52,13,224,1001,224,-65,224,4,224,1002,223,8,223,1001,224,6,224,1,223,224,223,1102,82,55,225,1001,213,67,224,1001,224,-126,224,4,224,102,8,223,223,1001,224,7,224,1,223,224,223,1,217,202,224,1001,224,-68,224,4,224,1002,223,8,223,1001,224,1,224,1,224,223,223,1002,176,17,224,101,-595,224,224,4,224,102,8,223,223,101,2,224,224,1,224,223,223,1102,20,92,225,1102,80,35,225,101,21,205,224,1001,224,-84,224,4,224,1002,223,8,223,1001,224,1,224,1,224,223,223,1101,91,45,225,1102,63,5,225,1101,52,58,225,1102,59,63,225,1101,23,14,225,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,1008,677,677,224,1002,223,2,223,1006,224,329,101,1,223,223,1108,226,677,224,1002,223,2,223,1006,224,344,101,1,223,223,7,677,226,224,102,2,223,223,1006,224,359,1001,223,1,223,8,677,226,224,102,2,223,223,1005,224,374,1001,223,1,223,1107,677,226,224,102,2,223,223,1006,224,389,1001,223,1,223,1008,226,226,224,1002,223,2,223,1005,224,404,1001,223,1,223,7,226,677,224,102,2,223,223,1005,224,419,1001,223,1,223,1007,677,677,224,102,2,223,223,1006,224,434,1001,223,1,223,107,226,226,224,1002,223,2,223,1005,224,449,1001,223,1,223,1008,677,226,224,102,2,223,223,1006,224,464,1001,223,1,223,1007,677,226,224,1002,223,2,223,1005,224,479,1001,223,1,223,108,677,677,224,1002,223,2,223,1006,224,494,1001,223,1,223,108,226,226,224,1002,223,2,223,1006,224,509,101,1,223,223,8,226,677,224,102,2,223,223,1006,224,524,101,1,223,223,107,677,226,224,1002,223,2,223,1005,224,539,1001,223,1,223,8,226,226,224,102,2,223,223,1005,224,554,101,1,223,223,1108,677,226,224,102,2,223,223,1006,224,569,101,1,223,223,108,677,226,224,102,2,223,223,1006,224,584,1001,223,1,223,7,677,677,224,1002,223,2,223,1005,224,599,101,1,223,223,1007,226,226,224,102,2,223,223,1005,224,614,1001,223,1,223,1107,226,677,224,102,2,223,223,1006,224,629,101,1,223,223,1107,226,226,224,102,2,223,223,1005,224,644,1001,223,1,223,1108,677,677,224,1002,223,2,223,1005,224,659,101,1,223,223,107,677,677,224,1002,223,2,223,1006,224,674,1001,223,1,223,4,223,99,226"
const ADD=1
const MULT=2
const STORE=3
const OUT=4
const JUMP_TRUE = 5
const JUMP_FALSE = 6
const LESSTHAN = 7
const EQUALS = 8

const HALT=99
const _input = 1

func Run(input int){
	intCodes := parseInput()
	//codeMap := printCodes(intCodes)
	fmt.Println()
	output := -1
	cursor := 0
	done  := false
	jump := 0
	for{
		jump = 0
		if done {
			break
		}
		opcode := evalOpcode(intCodes[cursor])
		switch opcode.code {
		case ADD,MULT:
			p1 := intCodes[cursor+1]
			p2 := intCodes[cursor+2]
			storeTo := intCodes[cursor+3]

			v1 := 0
			v2 := 0

			if opcode.paramMode1 == 0{
				v1 = intCodes[p1]
			} else{
				v1 = p1
			}
			if opcode.paramMode2 == 0{
				v2 = intCodes[p2]
			} else{
				v2 = p2
			}
			result := addOrMultiply(opcode,v1,v2)
			intCodes[storeTo] = result
			jump = opcode.length
		case JUMP_TRUE,JUMP_FALSE:
			p1 := intCodes[cursor+1]
			p2 := intCodes[cursor+2]

			if opcode.paramMode1 == 0 {
				p1 = intCodes[p1]
			}

			shouldJump := shouldJump(opcode,p1)
			if shouldJump {
				jump = p2
				if opcode.paramMode2 == 0 {
					jump = intCodes[p2]
				}
				cursor = jump
				jump = 0
			} else {
				jump = opcode.length
			}
		case LESSTHAN,EQUALS:
			p1 := intCodes[cursor+1]
			p2 := intCodes[cursor+2]
			storeTo := intCodes[cursor+3]

			v1 := 0
			v2 := 0

			if opcode.paramMode1 == 0{
				v1 = intCodes[p1]
			} else{
				v1 = p1
			}
			if opcode.paramMode2 == 0{
				v2 = intCodes[p2]
			} else{
				v2 = p2
			}
			result := compare(opcode,v1,v2)
			intCodes[storeTo] = result
			jump = opcode.length
		case STORE:
			p1 := intCodes[cursor+1]
			intCodes[p1] = input
			jump = opcode.length
		case OUT:
			if intCodes[cursor+2] == HALT {
				p1 := intCodes[cursor+1]
				if opcode.paramMode1 == 0{
					output = intCodes[p1]
				} else if opcode.paramMode1 == 1{
					output = p1
				}
				done = true
			}

			jump = opcode.length
		}
		cursor = cursor+jump
	}

	fmt.Println("Input", input,"Output is:",output)
}

func visitCode(ri codeMap){
	tm.MoveCursor(ri.start,ri.line)
	tm.Print(tm.Color(ri.value,tm.GREEN))
	tm.Flush()
}

func printCodes(codes []int) (result map[int]codeMap){
	result = make(map[int]codeMap)
	maxWidth := 120
	lineWidth := 0
	lineCnt := 1;
	colCnt := 1;
	for index,code := range codes{
		strValue := strconv.Itoa(code)
		length := len(strValue)
		ri := codeMap{value:strValue, index:index, length: length, line:lineCnt, start:colCnt}

		if lineWidth + length > maxWidth {
			colCnt = 1
			lineCnt++
			lineWidth = 0
			ri.line = lineCnt
			ri.start = colCnt
		} else{
			colCnt += length
			lineWidth += length
		}


		result[index] = ri
	}

	tm.Clear()
	tm.MoveCursor(1,1)
	tm.Flush()
	for index,_ := range codes{
		ri := result[index]
		tm.MoveCursor(ri.start,ri.line)
		tm.Print(ri.value)
		tm.Flush()
	}

	return
}

type codeMap struct{
	value string
	index int
	length int
	line int
	start int
}


type opcode struct{
	code int
	paramCount int
	length int
	paramMode1 int
	paramMode2 int
	paramMode3 int
}

func addOrMultiply(op opcode, v1, v2 int) (result int){
	switch op.code {
	case ADD:
		result = v1 + v2
	case MULT:
		result = v1*v2
	}
	return
}


func shouldJump(op opcode, v1 int) (result bool){
	result = false

	switch op.code {
	case JUMP_TRUE:
		if v1 != 0{
			result = true
		}
	case JUMP_FALSE:
		if v1 == 0{
			result = true
		}
	}
	return
}


func compare(op opcode, v1,v2 int) (value int){
	switch op.code {
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


func evalOpcode(intOpcode int)(result opcode){

	op :=  extractDigit(intOpcode,1)
	switch op {
	case ADD,MULT:
		result = opcode{
			code: op,
			paramCount:2,
			length:4,
			paramMode1: extractDigit(intOpcode,3),
			paramMode2: extractDigit(intOpcode,4),
			paramMode3:-1,
		}
		case JUMP_TRUE,JUMP_FALSE:
		result = opcode{
			code: op,
			paramCount:2,
			length:3,
			paramMode1: extractDigit(intOpcode,3),
			paramMode2: extractDigit(intOpcode,4),
			paramMode3:-1,
		}
		case EQUALS,LESSTHAN:
		result = opcode{
			code: op,
			paramCount:2,
			length:4,
			paramMode1: extractDigit(intOpcode,3),
			paramMode2: extractDigit(intOpcode,4),
			paramMode3:-1,
		}
		case STORE:
		result = opcode{
			code: op,
			paramCount:1,
			length:2,
			paramMode1: extractDigit(intOpcode,3),
			paramMode2: -1,
			paramMode3: -1,
		}
	case OUT:
		result = opcode{
			code: op,
			paramCount:1,
			length:2,
			paramMode1:extractDigit(intOpcode,3),
			paramMode2: -1,
			paramMode3: -1,
		}
	case HALT:
		result = opcode{
			code: op,
			paramCount:0,
			length:1,
			paramMode1: -1,
			paramMode2: -1,
			paramMode3: -1,
		}
	default:
		panic("Eval: Invalid operation code!")
	}
	return
}

func extractDigit(value int, digit int) int{
	pow10 := int(math.Pow10(digit-1))
	return (value / (pow10)) % 10
}



func parseInput() (result []int){
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


