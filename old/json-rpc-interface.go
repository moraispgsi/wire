package wire
/*
import (
    "strings"
)

type Parameter struct {
    identifier string
    paramType string
}

type RemoteProcedure struct {
    identifier string
    params map[string]Parameter
    sendChannel chan []byte
    result chan []byte
    error chan []byte
}

func NewRemoteProcedure(identifier, params []Parameter) RemoteProcedure {
    
    var rp = RemoteProcedure {
        identifier,
        make(map[string]Parameter),
        make(chan []byte),
        make(chan []byte),
        make(chan []byte),
    }
    
    for param := range inputParams {
        rp.params[param.identifier] = param
    }
    
    
    
}

func (rp *RemoteProcedure) Call(id int, params ...Stringer) {
    
    var jsonRpc string
    var paramsStr string
    
    for param := range rp.params {
        paramsStr += `"` + param.String() + `", `
    }
    paramsStr = strings.TrimSuffix(paramsStr, ",")
    
    var jsonRpc = `{"method": "` + rp.identifier + `", "params": [` + paramsStr + `], "id": `+ string(id) +`}`
    
    rp.sendChannel <- []byte(jsonRpc)
    
}


type JSONRPCInterface struct {
    remoteProcedures map[string]*RemoteProcedure
    rpcCallChannel chan []byte
    rpcResponseChannel chan []byte
}

func NewJSONRPCInterface (rps []*RemoteProcedure) Interface {
    
    var I Interface    
    
    for rp := range rps {
        
        I.remoteProcedures[rp.identifier] = rp
        
        go I.handleCall(rp)

    }

    return I
    
}

func (I *JSONRPCInterface) handleCall(rp *remoteProcedures) {

    I.rpcCallChannel <- rp.sendChannel
    
}


//Exposing the interface to the graph
func (I *JSONRPCInterface) expose(graph *Graph) {
    //DO exposing, creating sockets 
    
}*/