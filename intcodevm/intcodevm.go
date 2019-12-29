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

	OUTPUT_STOPNONE         = OutputMode(0)
	OUTPUT_STOPONFIRST      = OutputMode(1)
	OUTPUT_STOPONSECOND     = OutputMode(2)
	OUTPUT_STOPIFNEXTISHALT = OutputMode(3)
)

type Opcode struct {
	Code       int
	ParamCount int
	Length     int
	ParamMode1 int
	ParamMode2 int
	ParamMode3 int
}
type OutputMode int
type IntcodeVm struct {
	Cursor              int
	RelativeBase        int
	OutputMode          OutputMode
	WaitForInput        bool
	PrimaryMemory       []int
	ExtendedMemory      map[int]int
	Output              []int
	Halted              bool
}

func Run(program string, inputs []int, cursor int) (output int, pausedOn int, halted bool) {
	vm := IntcodeVm{Cursor: cursor, RelativeBase: 0, OutputMode: OUTPUT_STOPONFIRST, WaitForInput: true}
	vm.LoadProgram(program)
	vm.Continue(inputs)
	output, pausedOn, halted = vm.Output[0], vm.Cursor, vm.Halted
	return
}


func (vm *IntcodeVm) LoadProgram(program string) {
	vm.PrimaryMemory = vm.parseInput(program)
	vm.ExtendedMemory = make(map[int]int)
}

func (vm *IntcodeVm) RunProgram(program string, inputs []int) {
	vm.LoadProgram(program)
	vm.Continue(inputs)
	return
}

func (vm *IntcodeVm) Continue(inputs []int) {
	vm.Output = make([]int, 0)
	done := false
	jump := 0

	for !done {
		jump = 0
		opcode := evalOpcode(vm.readmemory(vm.Cursor))
		if opcode.Code == HALT {
			vm.Halted = true
			break
		}

		switch opcode.Code {
		case ADD, MULTIPLY:
			v1 := vm.evalParam(vm.Cursor+1, opcode.ParamMode1)
			v2 := vm.evalParam(vm.Cursor+2, opcode.ParamMode2)
			v3 := vm.evalStoreAddres(vm.Cursor+3, opcode.ParamMode3)
			result := addOrMultiply(opcode, v1, v2)
			vm.storetomemory(v3, result)
			jump = opcode.Length
		case JUMPTRUE, JUMPFALSE:
			v1 := vm.evalParam(vm.Cursor+1, opcode.ParamMode1)
			v2 := vm.evalParam(vm.Cursor+2, opcode.ParamMode2)

			shouldJump := shouldJump(opcode, v1)
			if shouldJump {
				jump = v2
				vm.Cursor = jump
				jump = 0
			} else {
				jump = opcode.Length
			}

		case LESSTHAN, EQUALS:
			v1 := vm.evalParam(vm.Cursor+1, opcode.ParamMode1)
			v2 := vm.evalParam(vm.Cursor+2, opcode.ParamMode2)
			v3 := vm.evalStoreAddres(vm.Cursor+3, opcode.ParamMode3)
			result := compare(opcode, v1, v2)
			vm.storetomemory(v3, result)
			jump = opcode.Length
		case INPUT:
			if vm.WaitForInput && len(inputs) == 0 {
				done = true
			} else {
				v1 := vm.evalStoreAddres(vm.Cursor+1, opcode.ParamMode1)
				result := inputs[0]
				vm.storetomemory(v1, result)
				jump = opcode.Length
				inputs = inputs[1:len(inputs)]
			}
		case OUTPUT:
			output := vm.evalParam(vm.Cursor+1, opcode.ParamMode1)
			vm.Output = append(vm.Output, output)
			if vm.OutputMode == OUTPUT_STOPONFIRST && len(vm.Output) == 1 {
				done = true
			} else if vm.OutputMode == OUTPUT_STOPONSECOND && len(vm.Output) == 2 {
				done = true
			} else if vm.OutputMode == OUTPUT_STOPIFNEXTISHALT && vm.readmemory(vm.Cursor+2) == HALT {
				done = true
			}
			jump = opcode.Length
		case RELBASE:
			v1 := vm.evalParam(vm.Cursor+1, opcode.ParamMode1)
			vm.RelativeBase += v1
			jump = opcode.Length
		}

		vm.Cursor = vm.Cursor + jump
	}
	return
}

func (vm *IntcodeVm) evalParam(pos, pmode int) (v int) {
	tmp := vm.readmemory(pos)
	if pmode == PMODEPOS {
		v = vm.readmemory(tmp)
	} else if pmode == PMODEVAL {
		v = tmp
	} else if pmode == PMODEREL {
		v = vm.readmemory(vm.RelativeBase + tmp)
	}
	return
}

func (vm *IntcodeVm) evalStoreAddres(pos, pmode int) (v int) {
	tmp := vm.readmemory(pos)
	if pmode == PMODEPOS {
		v = tmp
	} else if pmode == PMODEREL {
		v = vm.RelativeBase + tmp
	} else {
		panic(fmt.Sprintf("Parameter mode %d not supported.", pmode))
	}
	return
}

func (vm *IntcodeVm) readmemory(index int) (result int) {
	memoryLen := len(vm.PrimaryMemory)
	if index < memoryLen {
		result = vm.PrimaryMemory[index]
		return
	}

	offset := index - memoryLen

	result, _ = vm.ExtendedMemory[offset]
	return
}

func (vm *IntcodeVm) storetomemory(index, value int) {
	idx := index
	memoryLen := len(vm.PrimaryMemory)
	if idx < memoryLen {
		vm.PrimaryMemory[idx] = value
		return
	}

	offset := idx - memoryLen
	vm.ExtendedMemory[offset] = value
	return
}

func (vm *IntcodeVm) parseInput(problem string) (result []int) {
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
