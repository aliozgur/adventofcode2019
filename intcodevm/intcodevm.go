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
	INPUT     = 3
	OUTPUT    = 4
	JUMPTRUE  = 5
	JUMPFALSE = 6
	LESSTHAN  = 7
	EQUALS    = 8
	RELBASE   = 9
	HALT      = 99

	PMODEPOS = 0
	PMODEVAL = 1
	PMODEREL = 2
)

type Opcode struct {
	Code       int
	ParamCount int
	Length     int
	ParamMode1 int
	ParamMode2 int
	ParamMode3 int
}

type IntcodeVm struct {
	Cursor				int
	RelativeBase        int
	AuxiliaryMemorySize int
	StopOnFirstOutput   bool
	WaitForInput        bool
	StopIfNextOpIsHalt  bool
}

func Run(program string, inputs []int, cursor int) (output int, newProg string, pausedOn int, halted bool) {
	vm := IntcodeVm{Cursor:cursor,RelativeBase: 0, AuxiliaryMemorySize:0,StopOnFirstOutput: true, WaitForInput: true, StopIfNextOpIsHalt: false}
	output, newProg, pausedOn, halted = vm.RunProgram(program, inputs)
	return
}

func (vm *IntcodeVm) RunProgram(program string, inputs []int) (output int, newProg string, pausedOn int, halted bool) {
	cursor := vm.Cursor
	primaryMemory := parseInput(program)
	extendedMemory := make(map[int]int)

	output = -1
	done := false
	jump := 0
	relBase := vm.RelativeBase

	readmemory := func(index int) (result int) {
		memoryLen := len(primaryMemory)
		if index < memoryLen {
			result = primaryMemory[index]
			return
		}

		offset := index - memoryLen

		result,_ = extendedMemory[offset]
		return
	}

	storetomemory := func(index, value int) {
		idx := index
		memoryLen := len(primaryMemory)
		if idx < memoryLen {
			primaryMemory[idx] = value
			return
		}

		offset := idx - memoryLen
		extendedMemory[offset] = value
		return
	}

	evalParam := func(pos, pmode int) (v int) {
		tmp := readmemory(pos)
		if pmode == PMODEPOS {
			v = readmemory(tmp)
		} else if pmode == PMODEVAL {
			v = tmp
		} else if pmode == PMODEREL {
			v = readmemory(relBase + tmp)
		}
		return
	}

	evalStoreAddres := func(pos, pmode int) (v int) {
		tmp := readmemory(pos)
		if pmode == PMODEPOS {
			v = tmp
		} else if pmode == PMODEREL {
			v = relBase + tmp
		} else {
			panic(fmt.Sprintf("Parameter mode %d not supported.", pmode))
		}
		return
	}

	for !done {
		jump = 0
		opcode := evalOpcode(readmemory(cursor))
		if opcode.Code == HALT {
			halted = true
			break
		}

		switch opcode.Code {
		case ADD, MULTIPLY:
			v1 := evalParam(cursor+1, opcode.ParamMode1)
			v2 := evalParam(cursor+2, opcode.ParamMode2)
			v3 := evalStoreAddres(cursor+3, opcode.ParamMode3)
			result := addOrMultiply(opcode, v1, v2)
			storetomemory(v3, result)
			jump = opcode.Length
		case JUMPTRUE, JUMPFALSE:
			v1 := evalParam(cursor+1, opcode.ParamMode1)
			v2 := evalParam(cursor+2, opcode.ParamMode2)

			shouldJump := shouldJump(opcode, v1)
			if shouldJump {
				jump = v2
				cursor = jump
				jump = 0
			} else {
				jump = opcode.Length
			}

		case LESSTHAN, EQUALS:
			v1 := evalParam(cursor+1, opcode.ParamMode1)
			v2 := evalParam(cursor+2, opcode.ParamMode2)
			v3 := evalStoreAddres(cursor+3, opcode.ParamMode3)
			result := compare(opcode, v1, v2)
			storetomemory(v3, result)
			jump = opcode.Length
		case INPUT:
			if vm.WaitForInput && len(inputs) == 0 {
				done = true
			} else {
				v1 := evalStoreAddres(cursor+1, opcode.ParamMode1)
				result := inputs[0]
				storetomemory(v1, result)
				jump = opcode.Length
				inputs = inputs[1:len(inputs)]
			}
		case OUTPUT:
			output = evalParam(cursor+1, opcode.ParamMode1)
			if vm.StopOnFirstOutput {
				done = true
			}
			if vm.StopIfNextOpIsHalt && readmemory(cursor+2) == HALT {
				done = true
			}
			jump = opcode.Length
		case RELBASE:
			v1 := evalParam(cursor+1, opcode.ParamMode1)
			relBase += v1
			jump = opcode.Length
		}

		cursor = cursor + jump
	}

	pausedOn = cursor
	newProg = rebuildProgram(primaryMemory)
	return
}

func parseInput(problem string) (result []int) {
	var values = strings.Split(problem, ",")
	for _, value := range values {
		if value == "" {
			continue
		}
		v, e := strconv.Atoi(value)
		if e == nil {
			result = append(result, v)
		} else {
			log.Println("Can not convert value ", value)
		}
	}
	return
}

func evalOpcode(intOpcode int) (result Opcode) {
	if intOpcode == HALT {
		result = Opcode{
			Code:       intOpcode,
			ParamCount: 0,
			Length:     1,
			ParamMode1: extractDigit(intOpcode, 3),
			ParamMode2: extractDigit(intOpcode, 4),
			ParamMode3: extractDigit(intOpcode, 5),
		}
		return
	}

	op := extractDigit(intOpcode, 1)
	switch op {
	case ADD, MULTIPLY:
		result = Opcode{
			Code:       op,
			ParamCount: 2,
			Length:     4,
		}
	case JUMPTRUE, JUMPFALSE:
		result = Opcode{
			Code:       op,
			ParamCount: 2,
			Length:     3,
		}
	case EQUALS, LESSTHAN:
		result = Opcode{
			Code:       op,
			ParamCount: 2,
			Length:     4,
		}
	case INPUT:
		result = Opcode{
			Code:       op,
			ParamCount: 1,
			Length:     2,
		}
	case OUTPUT:
		result = Opcode{
			Code:       op,
			ParamCount: 1,
			Length:     2,
		}
	case RELBASE:
		result = Opcode{
			Code:       op,
			ParamCount: 1,
			Length:     2,
		}

	default:
		panic(fmt.Sprintf("Eval: Invalid operation Code. %d", op))
	}

	result.ParamMode1 = extractDigit(intOpcode, 3)
	result.ParamMode2 = extractDigit(intOpcode, 4)
	result.ParamMode3 = extractDigit(intOpcode, 5)

	return
}

func extractDigit(value int, digit int) int {
	pow10 := int(math.Pow10(digit - 1))
	return (value / (pow10)) % 10
}

/*
This function rebuilds the program string from the primary memory values
Please note: This function discards the bits stored in the extended memory map
 */
func rebuildProgram(opcodes []int) (program string) {
	program = ""
	for i, code := range opcodes {
		if i == len(opcodes)-1 {
			program += strconv.Itoa(code)
		} else {
			program += strconv.Itoa(code) + ","
		}
	}

	return
}

func addOrMultiply(op Opcode, v1, v2 int) (result int) {
	switch op.Code {
	case ADD:
		result = v1 + v2
	case MULTIPLY:
		result = v1 * v2
	}
	return
}

func shouldJump(op Opcode, v1 int) (result bool) {
	result = false

	switch op.Code {
	case JUMPTRUE:
		if v1 != 0 {
			result = true
		}
	case JUMPFALSE:
		if v1 == 0 {
			result = true
		}
	}
	return
}

func compare(op Opcode, v1, v2 int) (value int) {
	switch op.Code {
	case LESSTHAN:
		if v1 < v2 {
			value = 1
		} else {
			value = 0
		}
	case EQUALS:
		if v1 == v2 {
			value = 1
		} else {
			value = 0
		}
	}
	return
}
