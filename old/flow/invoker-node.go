package flow

import (
    "strconv"
)

type InvokerNode struct {
    Node
}

func NewInvokerNode(graph *Graph) *InvokerNode {
    
    //Invoker are unit nodes
    node := NewVoidNode(true)
    
    invNode := InvokerNode {
        node,
    }
    
    graph.AddNode(&invNode)
    
    node.AddAcceptCondition(func (idNode1 int64, idNode2 int64) bool {
        
        var node2 interface{}
        node2, err := graph.GetNode(idNode2)
        
        if err != nil {
            return false
        }
        _, ok := node2.(*InvokableNode) //Type assertion
        
        return ok
        
    })
    
    return &invNode

}

type invokeError struct {
    status int
    message string
}

func (err *invokeError) Error() string {
    return "Error status: " + strconv.Itoa(err.status) + " - " + err.message
}

func (invNode *InvokerNode) Invoke() error{
    
    graph := invNode.GetGraph()
    //Only one possible connection since this is a unit node
    for _,id := range invNode.GetConnected() {

        var node2 interface{}
        node2, err := graph.GetNode(id)
        
        if err != nil {
            return &invokeError{status: 10, message: "Something went wrong in the graph" }
        }
        invokableNode, ok := node2.(*InvokableNode) //Type assertion
        
        if !ok {
            return &invokeError{status: 11, message:  "Something went wrong with the cast, maybe the accepting conditions" }
        }
        
        go invokableNode.Invoke()

    }
    
    return nil
    
}