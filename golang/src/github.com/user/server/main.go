package main

import (
    //"log"
    //"net/http"
    "github.com/user/server/wire"
    //"github.com/gorilla/websocket"
    "fmt"
    "reflect"

)
/*

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}*/


type Adder struct {
    Value1 <- chan int
    Value2 <- chan int
    Sum chan <- int
}

func (adder Adder) Process() {
    
    value1 := <-adder.Value1
    value2 := <-adder.Value2
    
    adder.Sum <- value1 + value2
    
}

type Multiplier struct {
    In <-chan int
    Out chan<- int
}

func (m Multiplier) Process() {
    m.Out <- (<- m.In * 5)
}


func main() {
    
    
    
    
    net := wire.NewNetwork()
    
    m1 := new(Multiplier)
    
    inst1 := net.AddComponentInstance(m1)
    inst2 := net.AddComponentInstance(new(Multiplier))
    
    net.Connect(inst2, "In", inst1, "Out")
    
    //net.Start()
    
    fmt.Println()
    /*
    for {
        
    }*/
    

    /*
    graph := wire.NewGraph()
    ptrGraph := &graph
    
    node1 := wire.NewNode(true)
    node2 := wire.NewNode(true)
    
    id1 := ptrGraph.AddNode(&node1)
    id2 := ptrGraph.AddNode(&node2)
    
    err := ptrGraph.Connect(id1, id2)
    if err == nil {
        fmt.Println(ptrGraph.HasConnection(id1, id2))
    }
    
    adder := Adder{nil,nil, nil, nil}
    
    fmt.Println(wire.GetComponentOutputs(&adder)["Sum"])
    
    */
    
    
    
    
    /*
    invokableNode := wire.NewInvokableNode(ptrGraph)
    invokerNode1 := wire.NewInvokerNode(ptrGraph)
    invokerNode2 := wire.NewInvokerNode(ptrGraph)
    invokerNode3 := wire.NewInvokerNode(ptrGraph)
    
    err := invokerNode1.ConnectTo(invokableNode)
    if err != nil {
        fmt.Println(err.Error())
    }
    //fmt.Println(invokerNode1.IsConnectedTo(invokableNode.GetId()))
    
    err = invokerNode2.ConnectTo(invokableNode)
    if err != nil {
        fmt.Println(err.Error())
    }
    //fmt.Println(invokerNode2.IsConnectedTo(invokableNode.GetId()))
    
    err = invokerNode3.ConnectTo(invokableNode)
    if err != nil {
        fmt.Println(err.Error())
    }
    //fmt.Println(invokerNode3.IsConnectedTo(invokableNode.GetId()))
    
    invokeHandler := func () {
        fmt.Println("Invoked!!")
    }
    
    invokableNode.SetInvokeHandler(&invokeHandler)
    
    err = invokerNode1.Invoke()
    
    if err != nil {
        fmt.Println(err.Error())
    }
    
    err = invokerNode2.Invoke()
    
    if err != nil {
        fmt.Println(err.Error())
    }
    
    for {
        
    }
    
    */
    
    
    /*
    graphNodeServerRoute := "/graph/node/server/"

    router := NewRouter()
    
    interfaceIRI := "std/user-base/login-interface"
    cn := wire.NewClientNode(&graph, interfaceIRI)
    nodeRoute := "std/user-base/login-interface"
    
    route := Route {
        interfaceIRI,
        "GET",
        graphNodeServerRoute + nodeRoute,
        func (w http.ResponseWriter, r *http.Request) {
            cn.Upgrade(w, r)
        },
    }

    var handler http.Handler
        handler = route.HandlerFunc
    
    router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
            
    ///
    
    sn := wire.NewServerNode(&graph, "wss://echo.websocket.org")
    
    
    sn.ConnectTo(cn.GetId())
    
    
    
    ///
            
    
    
    fmt.Println(route.Pattern)

    log.Fatal(http.ListenAndServe(":8080", router))
    
    */
    
}