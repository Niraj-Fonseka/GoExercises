package main

import "fmt"

func DepthFirstSearchDetecCycle(adjList map[int][]int, root int, find int) {

	var Queue []int
	visited := make(map[int]bool)

	Queue = append(Queue, root)
	visited[root] = true

	fmt.Println("Stating Loop")
	for len(Queue) != 0 {
		dequeuedElement := Queue[0]
		Queue = Queue[1:len(Queue)]

		visited[dequeuedElement] = true

		fmt.Printf("Checking %d 's children \n", dequeuedElement)
		fmt.Printf("---> Children : %v \n", adjList[dequeuedElement])

		for _, child := range adjList[dequeuedElement] {
			isVisited := false

			if visited[child] {
				//detect cycle

			}

			if child == find {
				fmt.Println("Found : %d \n", find)
				break
			}

			//check if the node is previously visited
			for visitedNode, _ := range visited {

				if child == visitedNode {
					isVisited = true
				}
			}

			if !isVisited {
				visited[child] = true
				Queue = append(Queue, child)
			}
		}
	}
	fmt.Println("Visited : ")
	fmt.Println(visited)
}

func DepthFirstSearch(adjList map[int][]int, root int, find int) {

	var Queue []int
	visited := make(map[int]bool)

	Queue = append(Queue, root)
	visited[root] = true

	fmt.Println("Stating Loop")
	for len(Queue) != 0 {
		dequeuedElement := Queue[0]
		Queue = Queue[1:len(Queue)]

		visited[dequeuedElement] = true

		fmt.Printf("Checking %d 's children \n", dequeuedElement)
		fmt.Println("Queue before : %v \n", Queue)
		for _, child := range adjList[dequeuedElement] {
			isVisited := false

			// if child == find {
			// 	fmt.Println("Found : %d \n", find)
			// 	break
			// }

			//check if the node is previously visited
			for visitedNode, _ := range visited {
				if child == visitedNode {
					isVisited = true
				}
			}

			if !isVisited {
				visited[child] = true
				fmt.Println("visiting : ", child)
				Queue = append(Queue, child)

			}
		}
	}
	fmt.Println("Visited : ")
	fmt.Println(visited)
}

func DepthFirstSearchStack(adjList map[int][]int, root int, find int) {

	var Stack []int
	visited := make(map[int]bool)

	Stack = append(Stack, root)
	visited[root] = true

	fmt.Println("Stating Loop")
	for len(Stack) != 0 {
		dequeuedElement := Stack[len(Stack)-1]
		Stack = Stack[0 : len(Stack)-1]

		visited[dequeuedElement] = true

		fmt.Printf("Checking %d 's children \n", dequeuedElement)
		fmt.Println("Queue before : %v \n", Stack)
		for _, child := range adjList[dequeuedElement] {
			isVisited := false

			// if child == find {
			// 	fmt.Println("Found : %d \n", find)
			// 	break
			// }

			//check if the node is previously visited
			for visitedNode, _ := range visited {
				if child == visitedNode {
					isVisited = true
				}
			}

			if !isVisited {
				visited[child] = true
				fmt.Println("visiting : ", child)
				Stack = append(Stack, child)

			}
		}
	}
	fmt.Println("Visited : ")
	fmt.Println(visited)
}

func main() {

	graph := [][]int{{1, 2, 3}, {}, {4, 5}, {}, {0}, {6, 7}}

	adjList := make(map[int][]int)
	for row_index, row := range graph {

		var children []int
		for _, node := range row {
			children = append(children, node)
		}
		adjList[row_index] = children

	}

	fmt.Println(adjList)
	DepthFirstSearchStack(adjList, 0, 5)

}
