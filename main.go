package main 

import (  "fmt"
         "math/rand/v2"
)


func main() {
    

    graph := NewGraph()
    
    vertexCount := rand.IntN(6) + 5
    edgeCount := rand.IntN(6) + 4 


    for i := 0; i < vertexCount; i++{

        graph.AddVertex(i)
        
    }

    
    edgesAdded := 0
    
    for edgesAdded < edgeCount {

        a := rand.IntN(vertexCount)
        b := rand.IntN(vertexCount)
        
        
        if a != b && !graph.HasEdge(a,b) {
            graph.AddEdge(a,b)
            edgeAdded++
        }
    } 
 
    fmt.Println("Number of vertices: ", vertexCount)
    fmt.Println("Requested edges: ", edgeCount)


    fmt.Println(graph)

}
