package main 

type Graph struct {

    AdjacencyList map[int][]int

}


func NewGraph() *Graph {

    return &Graph{
        AdjacencyList: make(map[int][]int),
    }
}

func(g *Graph) AddVertex(vertex int) {

    g.AdjacencyList[vertex] = []int{}
}

func (g *Graph) AddEdge(a int, b int) {

    g.AdjacencyList[a] = append(g.AdjacencyList[a] ,b)
    g.AdjacencyList[b] = append(g.AdjacencyList[b], a)

}

func (g * Graph) HasEdge(a int, b int) bool {

    for _, neighbor := range g.AdjacencyList[a] {

        if neighbour == b {
            return true         
        }
    } 
    return false    
}


