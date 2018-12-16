package d7

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bobbykaz/aoc2018/utilities"
)

func Part1() {
	input := utilities.ReadFileIntoLines("input/input7.txt")
	nodeMap := make(map[string]*node, len(input))
	for i := 0; i < len(input); i++ {
		parent, child := process(input[i])
		pN, prs := nodeMap[parent]
		if !prs {
			pN = &node{NodeID: parent, Children: make([]*node, 0), Parents: make([]*node, 0), Done: false}
			nodeMap[parent] = pN
		}

		cN, prs := nodeMap[child]
		if !prs {
			cN = &node{NodeID: child, Children: make([]*node, 0), Parents: make([]*node, 0), Done: false}
			nodeMap[child] = cN
		}

		fmt.Println("Parent:", parent, "child:", child)
		pN.relate(cN)
	}
	fmt.Println("nodes made")
	nodesReady := make([]string, 0)
	nodesDone := make([]string, 0)
	for _, v := range nodeMap {
		if v.canBeDone() {
			nodesReady = append(nodesReady, v.NodeID)
		}
	}
	sort.Strings(nodesReady)
	for len(nodesReady) > 0 {
		node := nodeMap[nodesReady[0]]
		node.Done = true
		nodesReady = nodesReady[1:]
		for i := 0; i < len(node.Children); i++ {
			if node.Children[i].canBeDone() {
				nodesReady = append(nodesReady, node.Children[i].NodeID)
			}
		}
		sort.Strings(nodesReady)
		nodesDone = append(nodesDone, node.NodeID)
	}
	fmt.Println("Order:", nodesDone)
	stepTimes := make([]int, 0)
	for i := 0; i < len(nodesDone); i++ {
		name := nodesDone[i]
		char := name[0]
		stepTime := (int)(char-'A') + 1 + 60
		stepTimes = append(stepTimes, stepTime)
	}

	fmt.Println("times:", stepTimes)

	totalTime := 0
	workers := 5
	inprogress := make([]int, 0)
	for i := 0; i < workers; i++ {
		num, times := dequeue(stepTimes)
		stepTimes = times
		inprogress = append(inprogress, num)
	}

	for len(inprogress) > 0 {
		sort.Ints(inprogress)
		done, remaining := dequeue(inprogress)
		inprogress = remaining
		for i := 0; i < len(inprogress); i++ {
			inprogress[i] = inprogress[i] - done
		}

		totalTime += done

		if len(stepTimes) > 0 {
			next, times := dequeue(stepTimes)
			stepTimes = times
			inprogress = append(inprogress, next)
		}
	}

	fmt.Println("Done in:", totalTime)
}

func dequeue(arr []int) (int, []int) {
	return arr[0], arr[1:]
}

type node struct {
	NodeID   string
	Done     bool
	Parents  [](*node)
	Children [](*node)
}

//returns a node, and the next index of the list
func process(line string) (string, string) {
	//fmt.Println("1: ", line)
	pLine := strings.TrimPrefix(line, "Step ")
	pLine = strings.Trim(pLine, "\r")
	//fmt.Println("2: ", pLine)
	pLine = strings.TrimSuffix(pLine, " can begin.")
	pLine = strings.Trim(pLine, "\r")
	//fmt.Println("3: ", pLine)
	parts := strings.Split(pLine, " must be finished before step ")
	return strings.Trim(strings.TrimSpace(parts[0]), "\r"), strings.Trim(strings.TrimSpace(parts[1]), "\r")
}

func (n *node) relate(child *node) {
	n.Children = append(n.Children, child)
	child.Parents = append(child.Parents, n)
}

func (n *node) canBeDone() bool {
	//fmt.Println("Checking if done: ", n.NodeID)
	for i := 0; i < len(n.Parents); i++ {
		//fmt.Println("parent ", n.Parents[i].NodeID, " done: ", n.Parents[i].Done)
		if !((n.Parents[i]).Done) {
			return false
		}
	}

	return !n.Done
}

func traverseSumMetadata(n *node) int {

	if n.Children != nil {
		for i := 0; i < len(n.Children); i++ {
			traverseSumMetadata(n.Children[i])
		}
	}

	return -1
}
