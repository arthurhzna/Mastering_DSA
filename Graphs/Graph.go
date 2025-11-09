package main

import (
	"fmt"
	"strings"
)

type Graph struct {
	numberOfNodes int
	adjacentList  map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		numberOfNodes: 0,
		adjacentList:  make(map[string][]string),
	}
}

func (g *Graph) AddVertex(node string) {
	g.adjacentList[node] = []string{}
	g.numberOfNodes++
}

func (g *Graph) AddEdge(node1, node2 string) {
	if !contains(g.adjacentList[node1], node2) {
		g.adjacentList[node1] = append(g.adjacentList[node1], node2)
	}
	if !contains(g.adjacentList[node2], node1) {
		g.adjacentList[node2] = append(g.adjacentList[node2], node1)
	}
}

func (g *Graph) ShowConnections() {
	for node, connections := range g.adjacentList {
		var connectionStr strings.Builder
		for _, vertex := range connections {
			connectionStr.WriteString(vertex + " ")
		}
		fmt.Printf("%s-->%s\n", node, connectionStr.String())
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	myGraph := NewGraph()
	myGraph.AddVertex("0")
	myGraph.AddVertex("1")
	myGraph.AddVertex("2")
	myGraph.AddVertex("3")
	myGraph.AddVertex("4")
	myGraph.AddVertex("5")
	myGraph.AddVertex("6")
	myGraph.AddEdge("3", "1")
	myGraph.AddEdge("3", "4")
	myGraph.AddEdge("4", "2")
	myGraph.AddEdge("4", "5")
	myGraph.AddEdge("1", "2")
	myGraph.AddEdge("1", "0")
	myGraph.AddEdge("0", "2")
	myGraph.AddEdge("6", "5")
	myGraph.ShowConnections()
}

// Output:
// 0-->1 2
// 1-->3 2 0
// 2-->4 1 0
// 3-->1 4
// 4-->3 2 5
// 5-->4 6
// 6-->5
