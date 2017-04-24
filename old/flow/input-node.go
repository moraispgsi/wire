package flow

import (
    "strconv"
    "io"
)

type InputNode struct {
    Node
}

func NewInputNode(graph *Graph, getReader func(context *Context) io.Reader) *InputNode {

    node := NewVoidNode(false)

    node = InputNode {
        node,
    }
    
    graph.AddNode(&node)
    
    node.AddAcceptCondition(func (idNode1 int64, idNode2 int64) bool {
        
        var node2 interface{}
        node2, err := graph.GetNode(idNode2)
        
        if err != nil {
            return false
        }
        _, ok := node2.(*InputNode) //Type assertion
        
        return ok
        
    })
    
    node.getReader = getReader

    return &node

}


func () GetValue() interface{} {
    return node.getReader(context)
}