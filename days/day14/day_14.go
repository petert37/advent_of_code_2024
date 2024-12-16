package day14

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessPart1(input string, width, height int) string {
	robots := parseInput(input)
	q0 := 0
	q1 := 0
	q2 := 0
	q3 := 0
	midX := width / 2
	midY := height / 2
	for _, robot := range robots {
		x := mod((robot.position.x + robot.velocity.x*100), width)
		y := mod((robot.position.y + robot.velocity.y*100), height)
		if x < midX {
			if y < midY {
				q0++
			} else if y > midY {
				q3++
			}
		} else if x > midX {
			if y < midY {
				q1++
			} else if y > midY {
				q2++
			}
		}
	}
	result := q0 * q1 * q2 * q3
	return strconv.Itoa(result)
}

func ProcessPart2(input string, width, height int) string {
	robots := parseInput(input)
	formations := make(map[string]robotRound)

	for i := 0; i < 1_000_000; i++ {

		key := ""
		for _, robot := range robots {
			key += fmt.Sprintf("%d,%d,", robot.position.x, robot.position.y)
		}

		if _, ok := formations[key]; !ok {
			robotsCopy := make([]robot, len(robots))
			for i, r := range robots {
				robotCopy := robot{point{r.position.x, r.position.y}, point{r.velocity.x, r.velocity.y}}
				robotsCopy[i] = robotCopy
			}
			formations[key] = robotRound{robotsCopy, i}
		} else {
			break
		}

		for i, robot := range robots {
			robots[i].position.x = mod(robot.position.x+robot.velocity.x, width)
			robots[i].position.y = mod(robot.position.y+robot.velocity.y, height)
		}
	}

	maxOrder := 0
	maxOrderRound := 0
	var maxOrderRobots []robot = nil

	for _, formation := range formations {
		order := orderFactor(&formation.robots, width, height)
		if order > maxOrder {
			maxOrder = order
			maxOrderRound = formation.round
			maxOrderRobots = formation.robots
		}
	}

	fmt.Printf("Round %v\n", maxOrderRound)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			found := false
			for _, robot := range maxOrderRobots {
				if robot.position.x == x && robot.position.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return strconv.Itoa(maxOrderRound)
}

type point struct {
	x, y int
}

type robot struct {
	position point
	velocity point
}

type robotRound struct {
	robots []robot
	round  int
}

func parseInput(input string) []robot {
	result := make([]robot, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		parts := strings.Split(trimmedLine, " ")
		p := strings.Split(parts[0], "=")[1]
		v := strings.Split(parts[1], "=")[1]
		positions := strings.Split(p, ",")
		velocities := strings.Split(v, ",")
		x, err := strconv.Atoi(positions[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(positions[1])
		if err != nil {
			panic(err)
		}
		vx, err := strconv.Atoi(velocities[0])
		if err != nil {
			panic(err)
		}
		vy, err := strconv.Atoi(velocities[1])
		if err != nil {
			panic(err)
		}

		position := point{x, y}
		velocity := point{vx, vy}
		result = append(result, robot{position, velocity})
	}
	return result
}

func mod(a, b int) int {
	return (a%b + b) % b
}

var directions []point = []point{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func orderFactor(robots *[]robot, width, height int) int {
	order := 0
	for _, robot := range *robots {
		for _, direction := range directions {
			x := robot.position.x + direction.x
			y := robot.position.y + direction.y
			if x >= 0 && x < width && y >= 0 && y < height {
				for _, r := range *robots {
					if r.position.x == x && r.position.y == y {
						order++
						break
					}
				}
			}
		}
	}
	return order
}
