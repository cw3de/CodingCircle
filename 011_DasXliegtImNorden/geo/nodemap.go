package geo

import (
	"errors"
	"fmt"
	"strings"
)

type NodeMap map[string]*Node

func ShowMap(nodeMap NodeMap) {
	fmt.Println("Node Map:")
	for _, node := range nodeMap {
		fmt.Printf("%s:\n", node.Name)
		for _, n := range node.NodesToTheSouth {
			fmt.Printf("\tNorth of: %s\n", n.Name)
		}
		for _, e := range node.NodesToTheWest {
			fmt.Printf("\tEast of: %s\n", e.Name)
		}
	}
}

func BuildMap(listOfDirections []string) (NodeMap, error) {

	nodes := make(NodeMap)

	for _, line := range listOfDirections {
		fields := strings.Split(line, " ")
		if len(fields) != 3 {
			return nil, errors.New("invalid fields")
		}
		src, dir, dst := fields[0], fields[1], fields[2]

		a, ok := nodes[src]
		if !ok {
			a = &Node{Name: src}
			nodes[src] = a
		}
		b, ok := nodes[dst]
		if !ok {
			b = &Node{Name: dst}
			nodes[dst] = b
		}

		switch dir {
		case "N":
			a.NorthOf(b)
		case "S":
			b.NorthOf(a)
		case "E":
			a.EastOf(b)
		case "W":
			b.EastOf(a)
		case "NE":
			a.NorthOf(b)
			a.EastOf(b)
		case "NW":
			a.NorthOf(b)
			b.EastOf(a)
		case "SE":
			b.NorthOf(a)
			a.EastOf(b)
		case "SW":
			b.NorthOf(a)
			b.EastOf(a)
		default:
			return nil, errors.New("invalid direction")
		}
	}
	return nodes, nil
}

func checkToTheSouth(base, node *Node) bool {
	for _, southNode := range node.NodesToTheSouth {
		if southNode.IsNorthOf(node) {
			fmt.Println(node.Name, "is north of", southNode.Name, "and vice versa")
			return false
		}
		if southNode.IsNorthOf(base) {
			fmt.Println("Base", base.Name, "is north of", southNode.Name, "and vice versa")
			return false
		}
		if !checkToTheSouth(base, southNode) {
			return false
		}
	}
	return true
}

func checkToTheWest(base, node *Node) bool {
	for _, westNode := range node.NodesToTheWest {
		if westNode.IsEastOf(node) {
			fmt.Println(node.Name, "is east of", westNode.Name, "and vice versa")
			return false
		}
		if westNode.IsEastOf(base) {
			fmt.Println("Base", base.Name, "is east of", westNode.Name, "and vice versa")
			return false
		}
		if !checkToTheWest(base, westNode) {
			return false
		}
	}
	return true
}

func CheckMap(nodeMap NodeMap) bool {
	for _, node := range nodeMap {
		if !checkToTheSouth(node, node) {
			return false
		}
		if !checkToTheWest(node, node) {
			return false
		}
	}
	return true
}
