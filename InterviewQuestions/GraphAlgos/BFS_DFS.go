package main

import "fmt"
import "time"

func DepthFirstSearch(adjList map[int][]int, root int, find int) {

	var Queue []int
	var visited []int
	fmt.Println("Root : ", root)
	Queue = append(Queue, root)

	fmt.Println("Stating Loop")
	for len(Queue) != 0 {
		fmt.Println("Before removing the element : ", Queue)
		dequeuedElement := Queue[0]
		fmt.Printf("Removed element : %d \n", dequeuedElement)
		Queue = Queue[1:len(Queue)]
		fmt.Println("New Queue after removed : ", Queue)

		visited := append(visited, dequeuedElement)
		fmt.Printf("Visiting element : %d \n", dequeuedElement)
		fmt.Println("List of visted : ", visited)

		fmt.Println("Iterating over the children :  ", adjList[dequeuedElement])
		for _, child := range adjList[dequeuedElement] {
			fmt.Printf("Child : %d \n", child)
			fmt.Println("Before checking for visited ")
			fmt.Println("Visisted ", visited)
			isVisited := false

			//check if the node is previously visited
			for _, visitedNode := range visited {
				if child == visitedNode {
					isVisited = true
				}
			}

			if !isVisited {
				visited = append(visited, child)
				Queue = append(Queue, child)
			}
		}
		fmt.Println()
		time.Sleep(5 * time.Second)
	}

	// for parent, children := range adjList {
	// 	visited = append(visited, parent)

	// 	Q = append(Q, parent)
	// 	for _, child := range children {
	// 		//check in visitedd
	// 	}
	// }

}

func main() {

	graph := [][]int{{1, 2, 3}, {}, {4, 5}, {}, {}, {6, 7}}

	adjList := make(map[int][]int)
	for row_index, row := range graph {

		var children []int
		for _, node := range row {
			children = append(children, node)
		}
		adjList[row_index] = children

	}

	fmt.Println(adjList)
	DepthFirstSearch(adjList, 0)

}
