package flow

import (
    "strconv"
    "io"
)

type OutputNode struct {
    Node
    getReader func(context *Context) io.Reader
}

func NewOutputNode(graph *Graph, getReader func(context *Context) io.Reader) *OutputNode {

    node := NewVoidNode(false)

    node = OutputNode {
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


func (node *OutputNode) GetReader(context *Context) io.Reader {
    return node.getReader(context)
}