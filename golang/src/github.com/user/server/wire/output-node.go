package wire

import (
    "reflect"    
)

type outputNode struct {
    *Node
    kind reflect.Kind
}

func newOutputNode(graph *Graph, kind reflect.Kind) *outputNode {

    node := NewNode(false)

    out := outputNode {
        &node,
        kind,
    }
    
    graph.AddNode(&node)
    
    out.AddAcceptCondition(func (idNode1 int64, idNode2 int64) bool {
        
        var node2 interface{}
        node2, err := graph.GetNode(idNode2)
        
        if err != nil {
            return false
        }
        inputNode, ok := node2.(*inputNode) //Type assertion
        
        return ok && inputNode.kind == out.kind
        
    })

    return &out

}
