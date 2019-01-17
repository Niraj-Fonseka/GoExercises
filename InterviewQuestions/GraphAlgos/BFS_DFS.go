package main

func DepthFirstSearch(adjList map[int][]int) {

	var Q []int

	//insert Q : Q = append(Q , value)
	//remove Q : Q = Q[1:len(Q)]

	var visited []int

	for parent, children := range adjList {
		visited = append(visited, parent)

		Q = append(Q, parent)
		for _, child := range children {
			//check in visitedd
		}
	}

}

func main() {

	graph := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}

	adjList := make(map[int][]int)
	for row_index, row := range graph {

		var children []int
		for _, node := range row {
			children = append(children, node)
		}
		adjList[row_index] = children

	}

	DepthFirstSearch(adjList)

}
