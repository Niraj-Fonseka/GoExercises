package main

import "fmt"

type Graph struct {
	Nodes   []*Node
	AdjList map[Node][]*Node
}

type Node struct {
	Value int
}

func (g Graph) NewGraph() *Graph {
	adjlist := make(map[Node][]*Node)
	var nodes []*Node
	return &Graph{
		Nodes:   nodes,
		AdjList: adjlist,
	}
}

func (g Graph) NewNode(value int) {
	node := Node{Value: value}
	g.Nodes = append(g.Nodes, &node)
	var nodeList []*Node
	g.AdjList[node] = nodeList
}

func (g Graph) PrintGraph() {
	for node, children := range g.AdjList {
		fmt.Printf("Node : %d \n", node.Value)
		fmt.Println(children)
		fmt.Println()
	}
}

func (g Graph) GetNodeByValue(value int) *Node {
	for _, val := range g.Nodes {
		if val.Value == value {
			return val
		}
	}
	return nil
}

func (g Graph) AddRelationship(parent int, child int) {
	node := g.GetNodeByValue(parent)
	if node == nil {
		g.NewNode(parent)
	}

	newNode := g.GetNodeByValue(parent)

	g.NewNode(child)
	childNode := g.GetNodeByValue(child)

	list := append(g.AdjList[*childNode], childNode)

	fmt.Println(list)
}

func main() {
	fmt.Println("Creating Graph")
	graph := Graph{}.NewGraph()

	graph.NewNode(1)
	graph.NewNode(2)
	graph.NewNode(3)
	graph.PrintGraph()

	graph.AddRelationship(1, 3)

	graph.PrintGraph()

}
