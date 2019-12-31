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

func readInput(puzzle string) (result []moon) {
	var lines = strings.NewReader(puzzle)
	scanner := bufio.NewScanner(lines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		moon := moon{}
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &moon.position[0], &moon.position[1], &moon.position[2])
		result = append(result, moon)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Can not scan inputs!")
	}
	return
}

type point int

type moon struct {
	tempGravity [3]point
	position    [3]point
	velocity    [3]point
}

func (moon *moon) String() string {
	return fmt.Sprintf("pos=<x=%d,y=%d,z=%d>, vel=<x=%d,y=%d,z=%d>",
		moon.position[0],
		moon.position[1],
		moon.position[2],
		moon.velocity[0],
		moon.velocity[1],
		moon.velocity[2],
	)
}

func Part1() {
	fmt.Println("Part 1 starting...")
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
					g1, g2 := comparePoints(m1.position[k], m2.position[k])
					m1.tempGravity[k] += g1
					m2.tempGravity[k] += g2
				}
			}
			process(&moons[i])
		}
		printMoons(moons, step)
	}

	totalEnergy := 0
	for i := 0; i < len(moons); i++ {
		totalEnergy += calculateEnergy(moons[i])
	}
	fmt.Println("Part 1 DONE (Total Energy):", totalEnergy)
}

func Part2() {
	fmt.Println("Part 2 starting...")
	moons := readInput(puzzle)
	//initialMoons := readInput(puzzle)
	var initialX, initialY, initialZ string
	for i := 0; i < len(moons); i++ {
		initialX += fmt.Sprint(moons[i].position[0])
		initialY += fmt.Sprint(moons[i].position[1])
		initialZ += fmt.Sprint(moons[i].position[2])
	}

	var freqX, freqY, freqZ int

	for step := 2; ; step++ {
		var stateX, stateY, stateZ string
		for i := 0; i < len(moons); i++ {
			for j := i; j < len(moons); j++ {
				if j == i {
					continue
				}
				for k := 0; k < 3; k++ {
					m1 := &moons[i]
					m2 := &moons[j]
					g1, g2 := comparePoints(m1.position[k], m2.position[k])
					m1.tempGravity[k] += g1
					m2.tempGravity[k] += g2
				}
			}
			process(&moons[i])
			stateX += fmt.Sprint(moons[i].position[0])
			stateY += fmt.Sprint(moons[i].position[1])
			stateZ += fmt.Sprint(moons[i].position[2])
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

		// All dimensions are aligned at least once
		if freqX != 0 && freqY != 0 && freqZ != 0 {
			fmt.Println("Part 2: All dimension are aligned at leas once after ",step, " steps")
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

	// Now lets brute force and increment by max freq until all dimensions are in line
	fmt.Println("Part 2: Trying to find the step when all dimensions will be inline...")
	mcmFreqs := maxFreq
	for mcmFreqs%freqX != 0 || mcmFreqs%freqY != 0 || mcmFreqs%freqZ != 0 {
		mcmFreqs += maxFreq
	}
	fmt.Println("Part 2 DONE (Steps until initial position):", mcmFreqs)

}

func printMoons(moons []moon, step int) {
	if !utils.DEBUG {
		return
	}
	fmt.Println("Step: ", step)
	for i := 0; i < len(moons); i++ {
		fmt.Println(moons[i].String())
	}
}

func comparePoints(p1 point, p2 point) (change1, change2 point) {
	if p1 < p2 {
		change1 = point(1)
		change2 = point(-1)
	} else if p2 < p1 {
		change1 = point(-1)
		change2 = point(1)
	} else {
		change1 = point(0)
		change2 = point(0)
	}
	return
}

func applyGravity(m *moon) {
	for i := 0; i < 3; i++ {
		m.velocity[i] += m.tempGravity[i]
		m.tempGravity[i] = 0
	}
}

func applyVelocity(m *moon) {
	for i := 0; i < 3; i++ {
		m.position[i] += m.velocity[i]
	}
}

func calculateEnergy(m moon) (result int) {
	pot := 0
	kin := 0

	for i := 0; i < 3; i++ {
		pot += int(math.Abs(float64(m.position[i])))
		kin += int(math.Abs(float64(m.velocity[i])))
	}
	result = pot * kin
	return
}

func process(m *moon) {
	applyGravity(m)
	applyVelocity(m)
}

func compareStates(initial []moon, current []moon) (result bool) {
	result = true
	exit := false
	for i := 0; i < len(initial); i++ {
		for k := 0; k < 3; k++ {
			v1, v2 := comparePoints(initial[i].velocity[k], current[i].velocity[k])
			p1, p2 := comparePoints(initial[i].position[k], current[i].position[k])
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
