package solution11

import (
	in11 "adventofcode/inputs/input11"
	intvm "adventofcode/intcodevm"
	"adventofcode/utils"
	"fmt"
	tm "github.com/buger/goterm"
	"math"
	"time"
)

const (
	BLACKPANEL = Color(0)
	WHITEPANEL = Color(1)

	DIRLEFT  = Direction(-1)
	DIRRIGHT = Direction(1)
	DIRSTILL = Direction(0)

	FORWARD  = Movement(1)
	BACKWARD = Movement(-1)

	HEADINGUP    = Heading("U")
	HEADINGDOWN  = Heading("D")
	HEADINGLEFT  = Heading("L")
	HEADINGRIGHT = Heading("R")
)

type Heading string
type Color int
type Movement int
type Direction int

func (h Heading) ToInt() (result int) {
	switch h {
	case "R":
		result = 1
	case "L":
		result = -1
	case "U":
		result = 1
	case "D":
		result = -1
	default:
		result = 0
	}
	return
}

func (h Heading) Turn(direction Direction) (result Heading) {
	if direction == DIRLEFT {
		if h == HEADINGLEFT {
			result = HEADINGDOWN
		} else if h == HEADINGRIGHT {
			result = HEADINGUP
		} else if h == HEADINGDOWN {
			result = HEADINGRIGHT
		} else if h == HEADINGUP {
			result = HEADINGLEFT
		} else {
			panic(fmt.Sprintf("Unknown direction '%s'", h))
		}
	} else if direction == DIRRIGHT {
		if h == HEADINGLEFT {
			result = HEADINGUP
		} else if h == HEADINGRIGHT {
			result = HEADINGDOWN
		} else if h == HEADINGDOWN {
			result = HEADINGLEFT
		} else if h == HEADINGUP {
			result = HEADINGRIGHT
		} else {
			panic(fmt.Sprintf("Unknown heading '%s'", h))
		}
	} else if direction == DIRSTILL {
		result = h
	} else {
		panic(fmt.Sprintf("Unknown direction '%d'", direction))
	}
	return
}

func OutputToDirection(output int) (dir Direction) {
	switch output {
	case 0:
		dir = DIRLEFT
	case 1:
		dir = DIRRIGHT
	default:
		panic(fmt.Sprintf("Unsupported output %d. Can not convert to direction", output))
	}
	return
}

type Point struct {
	X     int
	Y     int
	Color Color
}

func (point Point) String() (result string) {
	return fmt.Sprintf("(%d,%d) Color: %d ", point.X, point.Y, point.Color)
}

func (point Point) Key() (result string) {
	return fmt.Sprintf("(%d,%d)", point.X, point.Y)
}
func (point Point) colorToChar() (result string) {
	if point.Color == WHITEPANEL {
		result = "◼"
	} else if point.Color == BLACKPANEL {
		result = " "
	} else {
		panic("Invalid color")
	}
	return
}

type Robot struct {
	UserOptions  *utils.UserOptions
	Heading       Heading
	VisitedPoints map[string]*Point
	CurrentPoint  *Point
	Brain         intvm.IntcodeVm
	maxX          int
	minX          int
	maxY          int
	minY          int
}


func SolveParts(){
	options := utils.UserOptions{Print:false}
	Solve(in11.Puzzle,BLACKPANEL,HEADINGUP,&options)

	options = utils.UserOptions{Print:true,PrintSleepMiliseconds:5}
	Solve(in11.Puzzle,WHITEPANEL,HEADINGUP,&options)
}

func Solve(program string, startColor Color, startHeading Heading, userOptions *utils.UserOptions) (robot *Robot) {
	startPoint := Point{X: 0, Y: 0, Color: startColor}
	robot = &Robot{CurrentPoint: &startPoint, Heading: startHeading, UserOptions:userOptions}
	robot.addPoint(&startPoint)

	robot.Brain = intvm.IntcodeVm{OutputMode: intvm.OUTPUT_STOPONSECOND}
	robot.Brain.LoadProgram(program)
	robot.execute()
	return
}

func (robot *Robot) execute() {
	input := []int{int(robot.CurrentPoint.Color)}
	robot.Brain.LoadProgram(in11.Puzzle)

	for {
		input = []int{int(robot.CurrentPoint.Color)}
		robot.Brain.Continue(input)
		if robot.Brain.Halted {
			break
		}

		color := Color(robot.Brain.Output[0])
		direction := OutputToDirection(robot.Brain.Output[1])
		robot.CurrentPoint.Color = color

		robot.move(direction, FORWARD)
	}
	fmt.Println("Visited Points:", len(robot.VisitedPoints))
	robot.printPanels()
}

func (robot *Robot) addPoint(point *Point) {
	if len(robot.VisitedPoints) == 0 {
		robot.VisitedPoints = make(map[string]*Point)
		robot.maxX = math.MinInt32
		robot.minX = math.MaxInt32
		robot.maxY = math.MinInt32
		robot.minY = math.MaxInt32
	}

	if point.X > robot.maxX {
		robot.maxX = point.X
	}
	if point.X < robot.minX {
		robot.minX = point.X
	}
	if point.Y > robot.maxY {
		robot.maxY = point.Y
	}
	if point.Y < robot.minY {
		robot.minY = point.Y
	}

	key := point.Key()
	robot.VisitedPoints[key] = point
}

func (robot Robot) String() (result string) {
	return fmt.Sprintf("%s, Heading:%s", robot.CurrentPoint.String(), robot.Heading)
}

func (robot *Robot) move(turn Direction, move Movement) {
	point := robot.CurrentPoint
	result := Point{X: point.X, Y: point.Y, Color: BLACKPANEL}

	switch robot.Heading {
	case HEADINGUP:
		if turn == DIRSTILL {
			result.Y = point.Y + int(move)
		} else {
			result.X = point.X + int(turn)*int(move)
		}
	case HEADINGDOWN:
		if turn == DIRSTILL {
			result.Y = point.Y + -1*int(move)
		} else {
			result.X = point.X + -1*int(turn)*int(move)
		}
	case HEADINGLEFT:
		if turn == DIRSTILL {
			result.X = point.X + -1*int(move)
		} else {
			result.Y = point.Y + int(turn)*int(move)
		}
	case HEADINGRIGHT:
		if turn == DIRSTILL {
			result.X = point.X + -1*int(move)
		} else {
			result.Y = point.Y + -1*int(turn)*int(move)
		}
	default:
		panic(fmt.Sprintf("Unknown turn '%d'", turn))
	}

	robot.Heading = robot.Heading.Turn(turn)

	key := result.Key()
	p, exist := robot.VisitedPoints[key]
	if exist {
		robot.CurrentPoint = p
	} else {
		robot.CurrentPoint = &result
		robot.addPoint(&result)
	}
	return
}

func (robot *Robot) printPanels() {

	if robot.UserOptions == nil || !robot.UserOptions.Print{
		return
	}
	sleepMiliseconds := robot.UserOptions.PrintSleepMiliseconds
	fmt.Println()
	fmt.Println()
	h := tm.Height()
	w := tm.Width()
	offsetX := robot.minX + w/2 - 1
	for _, p := range robot.VisitedPoints {
		// Translate to screen coordinates
		screenX := p.X + w/2
		screenY := h/4 - p.Y
		tm.MoveCursor(screenX-offsetX, screenY)
		tm.Println(tm.Color(p.colorToChar(), tm.GREEN))
		tm.Flush()
		if sleepMiliseconds > 0 {
			time.Sleep(time.Duration(sleepMiliseconds)* time.Millisecond)
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

}
