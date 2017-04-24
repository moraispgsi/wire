package flow

import (
    
    
)


type InvokableNode struct {
    Node
    invokeHandler *func()
    invokeConditions []*func() bool
}

func NewInvokableNode(graph *Graph) *InvokableNode {
    
    //Invokables are not unit nodes
    node := NewVoidNode(false)
    
    invNode := InvokableNode {
        node,
        nil,
        []*func () bool {},
    }
    
    graph.AddNode(&invNode)
    
    node.AddAcceptCondition(func (idNode1 int64, idNode2 int64) bool {
        
        var node2 interface{}
        node2, err := graph.GetNode(idNode2)
        
        if err != nil {
            return false
        }
        _, ok := node2.(*InvokerNode) //Type assertion
        
        return ok
        
    })
    

    return &invNode

}

func (invNode *InvokableNode) SetInvokeHandler(handler *func()) {
    invNode.invokeHandler = handler
}

func (invNode *InvokableNode) AddInvokeCondition(condition *func() bool) {
    invNode.invokeConditions = append(invNode.invokeConditions, condition)
}

func (invNode *InvokableNode) CanInvoke() bool {
    
    for _, condition := range invNode.invokeConditions {
        
        ok := (*condition)()
        if !ok {
            return false
        }
    }
    
    return true
    
}

func (invNode *InvokableNode) Invoke() error {
    
    if !invNode.CanInvoke() {
        return &invokeError { status: 20, message: "Invoke conditions were not met." }
    }
    
    if invNode.invokeHandler != nil {
        go (*invNode.invokeHandler)()
    }

    return nil
    
}