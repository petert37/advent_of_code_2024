package day24

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func ProcessPart1(input string) string {
	wires, gates := parseInput(input)
	outputWires := make([]string, 0)
	for wire := range wires {
		if strings.HasPrefix(wire, "z") {
			outputWires = append(outputWires, wire)
		}
	}
	slices.Sort(outputWires)
	output := 0
	for i, wire := range outputWires {
		wireValue := getWireValue(wire, &wires, &gates)
		output = output | wireValue<<i
	}
	return strconv.Itoa(output)
}

func ProcessPart2(input string) string {
	wires, gates := parseInput(input)
	visualize(&wires, &gates)
	fmt.Println("Open ./adder.gv in a graphviz viewer (https://dreampuf.github.io/GraphvizOnline/?engine=fdp) to see the circuit diagram")
	fmt.Println("The circuit diagram should be made up of full adders, liket the one on the sample image")
	fmt.Println("Find the 4 adders that are wrong and fix them")
	return "Find the solution manually"
}

type gate struct {
	input1   string
	input2   string
	output   string
	gateType string
}

func parseInput(input string) (wires map[string]int, gates []gate) {
	wires = make(map[string]int)
	gates = make([]gate, 0)
	lines := strings.Split(input, "\n")
	isGates := false
	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		if len(trimmedLine) == 0 {
			isGates = true
			continue
		}
		if isGates {
			parts := strings.Split(trimmedLine, " ")
			g := gate{
				input1:   parts[0],
				input2:   parts[2],
				output:   parts[4],
				gateType: parts[1],
			}
			gates = append(gates, g)
			if _, ok := wires[g.input1]; !ok {
				wires[g.input1] = -1
			}
			if _, ok := wires[g.input2]; !ok {
				wires[g.input2] = -1
			}
			if _, ok := wires[g.output]; !ok {
				wires[g.output] = -1
			}
		} else {
			parts := strings.Split(trimmedLine, ": ")
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			wires[parts[0]] = value
		}
	}
	return
}

func getWireValue(wire string, wires *map[string]int, gates *[]gate) int {
	if value, ok := (*wires)[wire]; ok && value != -1 {
		return value
	}
	var gate *gate = nil
	for _, g := range *gates {
		if g.output == wire {
			gate = &g
			break
		}
	}
	if gate != nil {
		gateInput1 := getWireValue(gate.input1, wires, gates)
		gateInput2 := getWireValue(gate.input2, wires, gates)
		output := -1
		if gateInput1 != -1 && gateInput2 != -1 {
			switch gate.gateType {
			case "AND":
				result := gateInput1 == 1 && gateInput2 == 1
				if result {
					output = 1
				} else {
					output = 0
				}
			case "OR":
				result := gateInput1 == 1 || gateInput2 == 1
				if result {
					output = 1
				} else {
					output = 0
				}
			case "XOR":
				result := gateInput1 != gateInput2
				if result {
					output = 1
				} else {
					output = 0
				}
			}
		}
		if output != -1 {
			(*wires)[wire] = output
			return output
		}
	}
	return -1
}

func visualize(wires *map[string]int, gates *[]gate) {
	g := graph.New(graph.StringHash, graph.Directed())
	for wire := range *wires {
		if isInputWire(wire) {
			g.AddVertex(wire, graph.VertexAttribute("shape", "circle"))
		} else if isOutputWire(wire) {
			g.AddVertex(wire, graph.VertexAttribute("shape", "square"))
		}
	}
	for _, gate := range *gates {
		hash := gateHash(&gate)
		shape := "box"
		switch gate.gateType {
		case "AND":
			shape = "invhouse"
		case "OR":
			shape = "invtriangle"
		case "XOR":
			shape = "diamond"
		}
		g.AddVertex(hash, graph.VertexAttribute("shape", shape), graph.VertexAttribute("label", gate.gateType))
	}
	for wire := range *wires {
		addWire(wire, gates, &g)
	}
	file, _ := os.Create("./adder.gv")
	_ = draw.DOT(g, file)
}

func addWire(wire string, gates *[]gate, g *graph.Graph[string, string]) {
	start := ""
	if isInputWire(wire) {
		start = wire
	} else {
		for _, gate := range *gates {
			if gate.output == wire {
				start = gateHash(&gate)
				break
			}
		}
	}
	if isOutputWire(wire) {
		(*g).AddEdge(start, wire, graph.EdgeAttribute("label", wire))
	} else {
		for _, gate := range *gates {
			if gate.input1 == wire {
				(*g).AddEdge(start, gateHash(&gate), graph.EdgeAttribute("label", wire))
			}
			if gate.input2 == wire {
				(*g).AddEdge(start, gateHash(&gate), graph.EdgeAttribute("label", wire))
			}
		}
	}
}

func gateHash(gate *gate) string {
	return gate.input1 + " " + gate.gateType + " " + gate.input2 + " -> " + gate.output
}

func isInputWire(wire string) bool {
	return strings.HasPrefix(wire, "x") || strings.HasPrefix(wire, "y")
}

func isOutputWire(wire string) bool {
	return strings.HasPrefix(wire, "z")
}
