package wire

import (
    //"reflect"    
    "errors"
)

type componentInstance struct {
    processHandler func ()
    inputNodes map[string] *inputNode
    outputNodes map[string] *outputNode
}

type Network struct {
    graph Graph
    componentCurrentId int64
    components map[int64] componentInstance
}

func NewNetwork() Network{
    return Network{
        NewGraph(),
        0,
        make(map[int64] componentInstance),
    }
}

func (net *Network) AddComponentInstance(c Component) int64 {
    
    mapInputs := GetComponentInputs(c)
    mapOutputs := GetComponentOutputs(c)
    
    inputNodes := make(map[string] *inputNode)
    outputNodes := make(map[string] *outputNode)
    
    for key, value := range mapInputs {
        inputNodes[key] = newInputNode(&net.graph, value)
    }
    for key, value := range mapOutputs {
        outputNodes[key] = newOutputNode(&net.graph, value)
    }

    componentId := net.componentCurrentId
    net.componentCurrentId ++
    
    net.components[componentId] = componentInstance {
        c.Process,
        inputNodes,
        outputNodes,
    }
    
    return componentId
    
}

func (net *Network) Connect(compInId int64, in string, compOutId int64, out string) error {
    
    if _,ok := net.components[compInId]; !ok {
        return errors.New("Component id for the in port was not found.")
    }
    
    if _,ok := net.components[compOutId]; !ok {
        return errors.New("Component id for the out port was not found.")
    }
    
    inNode := net.components[compInId].inputNodes[in]
    outNode := net.components[compOutId].outputNodes[out]
    
    if inNode == nil {
        return errors.New("In port was not found.")
    }
    
    if outNode == nil {
        return errors.New("Out port was not found.")
    }
    
    return net.graph.Connect(inNode.GetId(), outNode.GetId()) 
    
}

func (net *Network) Start(){
    
    for _, inst := range net.components {
        
        go inst.processHandler()
        
    }
    
}