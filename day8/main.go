package main

import (
	"fmt"
	"github.com/cfagiani/aoc2018/util"
	"strconv"
	"strings"
)

type TreeNode struct {
	metadata []int
	children []TreeNode
}

func main() {
	inputString := util.ReadFileAsString("input/day8.input")
	tree := buildTree(strings.Split(inputString, " "))
	part1(tree)
	part2(tree)
}

func part1(root TreeNode) {
	sum := sumMetadata(root)
	fmt.Printf("Sum of metadata is %d\n", sum)
}

func part2(root TreeNode) {
	val := getValue(root)
	fmt.Printf("Value of root is %d\n", val)
}

func getValue(node TreeNode) int {
	if len(node.children) == 0 {
		return sumMetadata(node)
	} else {
		sum := 0
		for i := 0; i < len(node.metadata); i++ {
			idx := node.metadata[i]
			if idx > len(node.children) {
				continue
			} else {
				sum += getValue(node.children[idx-1])
			}

		}
		return sum
	}

}

func sumMetadata(node TreeNode) int {
	sum := 0
	for i := 0; i < len(node.children); i++ {
		sum += sumMetadata(node.children[i])
	}
	for i := 0; i < len(node.metadata); i++ {
		sum += node.metadata[i]
	}
	return sum
}

func buildTree(input []string) TreeNode {
	root, _ := buildNode(input, 0)
	return root
}

func buildNode(input []string, start int) (TreeNode, int) {
	node := TreeNode{metadata: []int{}, children: []TreeNode{}}
	pos := start
	numChildren, _ := strconv.Atoi(input[pos])
	pos++
	numMeta, _ := strconv.Atoi(input[pos])
	pos++
	for i := 0; i < numChildren; i++ {
		child, advanceBy := buildNode(input, pos)
		pos += advanceBy
		node.children = append(node.children, child)
	}
	for i := 0; i < numMeta; i++ {
		val, _ := strconv.Atoi(input[pos+i])
		node.metadata = append(node.metadata, val)
	}
	pos += numMeta
	return node, pos - start
}
