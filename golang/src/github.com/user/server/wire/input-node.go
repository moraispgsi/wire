package wire

import (
    "reflect"    
)


type inputNode struct {
    *Node
    kind reflect.Kind
}

func newInputNode(graph *Graph, kind reflect.Kind) *inputNode {

    node := NewNode(false)

    in := inputNode {
        &node,
        kind,
    }
    
    graph.AddNode(&node)
    
    in.AddAcceptCondition(func (idNode1 int64, idNode2 int64) bool {
        
        var node2 interface{}
        node2, err := graph.GetNode(idNode2)
        
        if err != nil {
            return false
        }
        outputNode, ok := node2.(*outputNode) //Type assertion
        
        return ok && outputNode.kind == in.kind
        
    })
    
    /*in.AddConnectionAction(func (idNode1 int64, idNode2 int64){
        
        var node2 interface{}
        node2, err := graph.GetNode(idNode2)
        if err != nil {
            return 
        }
        
        outputNode, ok := node2.(*outputNode)
        
        
    })*/

    return &in

}
