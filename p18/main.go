package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("triangles.txt")
	root := generateTree(input[1:], &Node{Value: input[0]}, 2)
	fmt.Println(root)
	//values := &[]int{}
	//getValues(root, values)
	//fmt.Println(findLargest(*values))
	//fmt.Printf("=%d\n", traverseForMaximum(root))
}

func getValues(node *Node, values *[]int) {
	if node != nil {
		*values = append(*values, node.Value)
		getValues(node.Left, values)
		getValues(node.Right, values)
	}
}

func findLargest(values []int) int {
	largest := 0
	for _, x := range values {
		if x > largest {
			largest = x
		}
	}
	return largest
}

func generateTree(pyramid []int, curNode *Node, numberOfNodes int) *Node {
	if len(pyramid) == 0 {
		return nil
	}
	left := &Node{Value: pyramid[0] + curNode.Value}
	right := &Node{Value: pyramid[1] + curNode.Value}
	curNode.Left = left
	curNode.Right = right
	if len(pyramid) <= numberOfNodes+1 {
		generateTree([]int{}, left, numberOfNodes+1)
	} else {
		generateTree(pyramid[numberOfNodes:], left, numberOfNodes+1)
	}
	if len(pyramid) <= numberOfNodes+2 {
		generateTree([]int{}, right, numberOfNodes+1)
	} else {
		generateTree(pyramid[numberOfNodes+1:], right, numberOfNodes+1)
	}
	return curNode
}

func traverseForMaximum(node *Node) int {
	fmt.Printf("%d+\n", node.Value)
	if node.Left == nil {
		return node.Value
	}
	if node.Left.Value == node.Right.Value {
		fmt.Println("Equals!")
	}
	if node.Left.Value > node.Right.Value {
		return node.Value + traverseForMaximum(node.Left)
	}
	return node.Value + traverseForMaximum(node.Right)
}

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func (n *Node) String() {
	if n != nil {
		fmt.Println(n.Value)
		n.Left.String()
		n.Right.String()
	}
}

func parseInput(href string) []int {
	data, _ := ioutil.ReadFile(href)
	lines := strings.Split(string(data), "\n")
	pyramid := []int{}
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		for _, n := range numbers {
			val, _ := strconv.Atoi(n)
			pyramid = append(pyramid, val)
		}
	}
	return pyramid
}
