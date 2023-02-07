// Given an unweighted, undirected graph which represents a metro map as follows

// vertices are stations
// edges are the path between two stations
// Given a start station and ending station, determine the minimum number of stops that must be made to get to the destination.

// Input: A Graph (unweighted/undirected), a starting Vertex, and an ending Vertex
// Output: Integer

// Input: The graph represented below, Vertex A, Vertex F
//  {
//     'A': ['B','C','D']
//     'B': ['A','E']
//     'C': ['A','F']
//     'D': ['A','G']
//     'E': ['G','B']
//     'F': ['C','G']
//     'G': ['D','E','F']
// }
// Output: 2 Stops (From A stop at C, and then stop at F)
/*
Constraints
Time Complexity: O(V + E) where V is the number of Vertices and E is the number of Edges
Auxiliary Space Complexity: O(V)

			'A': ['B','C','D']
				'B' -> ['A','E']
						'A' -> skip
						'E' -> ['G','B']
								'G' -> ['D','E','F']
								'B' -> skip

				'C' -> ['A','F']
						'A' -> skip
						'F' -> end

				'D' -> ['A','G']
						'A' -> skip
						'G' -> found

			'F': ['C','G']

Pseudocode
declare a map[string][]string to hold all routes
iterate the starting route
	foreach start route vertex
		increment hop counter
		if endvertex equals vertex or endroute contains vertex
			record new mincount (if lower)
			return mincount


*/

package main

import (
	"fmt"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

var depth int

func countEdgesToVertex(routes map[string][]string, visited map[string]int, current string, start string, end string, edges *int) {

	depth++
	vertices := routes[current]
	//visited[current] += 1

	fmt.Printf("[%s] %s %v\n\t", start, current, vertices)

	for _, vertex := range vertices {
		if vertex == end {
			*edges += 1
			fmt.Printf("\t%s -> %s %d\n", current, vertex, depth)
			return
		}

		_, ok := visited[vertex]
		if true == ok {
			continue
		}
		if current != vertex {
			visited[vertex] += 1
			countEdgesToVertex(routes, visited, vertex, start, end, edges)
		}
	}
}

func shortestRoute(routes map[string][]string, start string, end string) int {
	var shortestRoute int = 0
	depth = -1
	var visitedRoutes map[string]int
	visitedRoutes = make(map[string]int, 0)

	countEdgesToVertex(routes, visitedRoutes, start, start, end, &shortestRoute)

	fmt.Println()
	// for key, val := range visitedRoutes {
	// 	fmt.Printf("%s -> %v, ", key, val)
	// }
	// fmt.Println()
	// fmt.Printf("Depth -> %d\n", depth)

	return shortestRoute
}

func main() {
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C", "D"}
	graph["B"] = []string{"A", "E"}
	graph["C"] = []string{"A", "F"}
	graph["D"] = []string{"A", "G"}
	graph["E"] = []string{"G", "B"}
	graph["F"] = []string{"C", "G"}
	graph["G"] = []string{"D", "E", "F"}

	shortest := shortestRoute(graph, "A", "F")
	fmt.Println(shortest)

}
