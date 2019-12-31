package solution12

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"strings"
)

const puzzle = `
<x=3, y=-6, z=6>
<x=10, y=7, z=-9>
<x=-3, y=-7, z=9>
<x=-8, y=0, z=4>
`

func readInput(puzzle string) (result []Moon) {
	var lines = strings.NewReader(puzzle)
	scanner := bufio.NewScanner(lines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		moon := Moon{}
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &moon.Position[0], &moon.Position[1], &moon.Position[2])
		result = append(result, moon)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Can not scan inputs!")
	}
	return
}

type Point int

type Moon struct {
	Name     string
	Gravity  [3]Point
	Position [3]Point
	Velocity [3]Point
}

func (moon *Moon) String() string {
	return fmt.Sprintf("%s pos=<x=%d,y=%d,z=%d>, vel=<x=%d,y=%d,z=%d>",
		moon.Name,
		moon.Position[0],
		moon.Position[1],
		moon.Position[2],
		moon.Velocity[0],
		moon.Velocity[1],
		moon.Velocity[2],
	)
}

func Part1() {
	moons := readInput(puzzle)
	for step := 1; step <= 1000; step++ {
		for i := 0; i < len(moons); i++ {
			for j := i; j < len(moons); j++ {
				if j == i {
					continue
				}
				for k := 0; k < 3; k++ {
					m1 := &moons[i]
					m2 := &moons[j]
					g1, g2 := comparePoints(m1.Position[k], m2.Position[k])
					m1.Gravity[k] += g1
					m2.Gravity[k] += g2
				}
			}
			process(&moons[i])
		}
	}

	totalEnergy := 0
	for i := 0; i < len(moons); i++ {
		totalEnergy += calculateEnergy(moons[i])
	}

	fmt.Println("Total Energy:", totalEnergy)
}

func Part2() {
	moons := readInput(puzzle)
	initialMoons := readInput(puzzle)
	var initialX, initialY, initialZ string
	for i:=0;i<len(moons);i++{
		initialX += fmt.Sprint(moons[i].Position[0])
		initialY += fmt.Sprint(moons[i].Position[1])
		initialZ += fmt.Sprint(moons[i].Position[2])
	}

	var freqX, freqY, freqZ int
	for step := 2; ;step++ {
		var stateX, stateY, stateZ string
		for i := 0; i < len(moons); i++ {
			for j := i; j < len(moons); j++ {
				if j == i {
					continue
				}
				for k := 0; k < 3; k++ {
					m1 := &moons[i]
					m2 := &moons[j]
					g1, g2 := comparePoints(m1.Position[k], m2.Position[k])
					m1.Gravity[k] += g1
					m2.Gravity[k] += g2
				}
			}
			process(&moons[i])
			stateX += fmt.Sprint(moons[i].Position[0])
			stateY += fmt.Sprint(moons[i].Position[1])
			stateZ += fmt.Sprint(moons[i].Position[2])
		}

		if stateX == initialX && freqX == 0 {
			freqX = step
		}
		if stateY == initialY && freqY == 0 {
			freqY = step
		}
		if stateZ == initialZ && freqZ == 0 {
			freqZ = step
		}

		if freqX != 0 && freqY != 0 && freqZ != 0 {
			break
		}
	}

	maxFreq := freqX
	if freqY > maxFreq {
		maxFreq = freqY
	}
	if freqZ > maxFreq {
		maxFreq = freqZ
	}
	var mcmFreqs int = maxFreq

	for mcmFreqs%freqX != 0 || mcmFreqs%freqY != 0 || mcmFreqs%freqZ != 0 {
		mcmFreqs += maxFreq
	}
	fmt.Println("Part 2 (Steps until initial position):", mcmFreqs)

}

func printMoons(moons []Moon, step int) {
	if !utils.DEBUG {
		return
	}
	fmt.Println("Step: ", step)
	for i := 0; i < len(moons); i++ {
		fmt.Println(moons[i].String())
	}
}
func comparePoints(p1 Point, p2 Point) (change1, change2 Point) {
	if p1 < p2 {
		change1 = Point(1)
		change2 = Point(-1)
	} else if p2 < p1 {
		change1 = Point(-1)
		change2 = Point(1)
	} else {
		change1 = Point(0)
		change2 = Point(0)
	}
	return
}

func applyGravity(m *Moon) {
	for i := 0; i < 3; i++ {
		m.Velocity[i] += m.Gravity[i]
		m.Gravity[i] = 0
	}
}

func applyVelocity(m *Moon) {
	for i := 0; i < 3; i++ {
		m.Position[i] += m.Velocity[i]
	}
}

func calculateEnergy(m Moon) (result int) {
	pot := 0
	kin := 0

	for i := 0; i < 3; i++ {
		pot += int(math.Abs(float64(m.Position[i])))
		kin += int(math.Abs(float64(m.Velocity[i])))
	}
	result = pot * kin
	return
}

func process(m *Moon) {
	applyGravity(m)
	applyVelocity(m)
}

func compareStates(initial []Moon, current []Moon) (result bool) {
	result = true
	exit := false
	for i := 0; i < len(initial); i++ {
		for k := 0; k < 3; k++ {
			v1, v2 := comparePoints(initial[i].Velocity[k], current[i].Velocity[k])
			p1, p2 := comparePoints(initial[i].Position[k], current[i].Position[k])
			if v1 != 0 || v2 != 0 || p1 != 0 || p2 != 0 {
				result = false
				exit = true
				break
			}
		}
		if exit {
			break
		}
	}
	return result
}
