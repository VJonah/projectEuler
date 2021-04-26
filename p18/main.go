package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("test.txt")
	fmt.Println(input)
	//startNode := &Node{Value: input[0][0]}
}

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func (n *Node) String() string {
	fmt.Println(n.Value)
	if n.Left != nil {
		fmt.Println(n.Left.String())
	}
	if n.Right != nil {
		fmt.Println(n.Right.String())
	}
	s := strconv.Itoa(n.Value)
	return s
}

func findLargest(node *Node, largest int) int {
	fmt.Println(largest, node.Value)
	if node.Left == nil {
		if largest < node.Value {
			largest = node.Value
		}
		return largest
	}
	findLargest(node.Left, largest)
	findLargest(node.Right, largest)
	return largest
}

func parseInput(href string) [][]int {
	data, _ := ioutil.ReadFile(href)
	lines := strings.Split(string(data), "\n")
	pyramid := [][]int{}
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		row := []int{}
		for _, n := range numbers {
			val, _ := strconv.Atoi(n)
			row = append(row, val)
		}
		pyramid = append(pyramid, row)
	}
	return pyramid
}
