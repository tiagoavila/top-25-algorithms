package kahnstopologicalsort

import "top25algorithms/unweighteddirectedgraph"

// A Topological sort or Topological ordering of a directed graph is a linear ordering of its vertices such
// that for every directed edge uv from vertex u to vertex v, u comes before v in the ordering.
// Topological order is possible if and only if the graph has no directed cycles, i.e. if the graph is DAG.

// Use Cases and Situations for Kahn's Algorithm:
// Task Scheduling and Dependency Resolution:
// Kahn's Algorithm is ideal for scenarios where tasks need to be performed in a specific order, and some tasks cannot start until others have finished. This is common in build systems, where certain modules must be compiled before others, and in package managers, which need to install dependencies before the dependent packages.

// Course Scheduling:
// In educational institutions, Kahn's Algorithm can help in scheduling courses where some courses have prerequisites. The algorithm can determine an order in which courses should be taken to satisfy all prerequisite requirements.

// Project Management:
// In project management, especially within software development, tasks may depend on the completion of others. Kahn's Algorithm can help project managers visualize tasks' dependencies and plan the project's workflow to ensure efficient completion without deadlock.

// Data Processing Pipelines:
// When dealing with complex data processing tasks that involve multiple stages, Kahn's Algorithm can help organize these stages so that each processing step is completed in the right order, ensuring that all prerequisites for a given stage are met before it begins.

// Resolving Circular Dependencies:
// In systems design and software engineering, circular dependencies can lead to problems. Kahn's Algorithm can detect circular dependencies by identifying if the graph has cycles, which is a situation where the algorithm cannot find a topological ordering.

// Graph Databases and Ontologies:
// In graph databases or systems dealing with ontologies, Kahn's Algorithm can help order entities in a way that respects their hierarchical or dependency relationships, aiding in efficient data retrieval and reasoning.

func TopologicalSort(graph *unweighteddirectedgraph.UnweightedDirectedGraph) []string {
	sorted := make([]string, 0, len(graph.Vertices))
	inDegreeZeroVertices := graph.GetInDegreeZeroVertices()

	for len(inDegreeZeroVertices) > 0 {
		currentVertex := inDegreeZeroVertices[0]
		sorted = append(sorted, currentVertex)
		inDegreeZeroVertices = inDegreeZeroVertices[1:]

		for _, edge := range graph.Vertices[currentVertex].Edges {
			graph.DecreaseInDegree(edge.Dest)

			if graph.Vertices[edge.Dest].InDegree == 0 {
				inDegreeZeroVertices = append(inDegreeZeroVertices, edge.Dest)
			}
		}
	}

	return sorted
}
