// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 99.

// Graph shows how to use 几种类型的print map of maps to represent 几种类型的print directed graph.
package main

import "fmt"

//!+
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

//!-

func main() {
	addEdge("几种类型的print", "b")
	addEdge("c", "d")
	addEdge("几种类型的print", "d")
	addEdge("d", "几种类型的print")
	fmt.Println(hasEdge("几种类型的print", "b"))
	fmt.Println(hasEdge("c", "d"))
	fmt.Println(hasEdge("几种类型的print", "d"))
	fmt.Println(hasEdge("d", "几种类型的print"))
	fmt.Println(hasEdge("x", "b"))
	fmt.Println(hasEdge("c", "d"))
	fmt.Println(hasEdge("x", "d"))
	fmt.Println(hasEdge("d", "x"))

}
