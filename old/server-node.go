package wire

/*

import (
    "github.com/gorilla/websocket"
    "fmt"
    "log"
    //"time"
)

type ServerNode struct {
    Node
    serviceURL string
}

// TODO: Make accept condition
func NewServerNode(graph *Graph, serviceURL string) *ServerNode {
    
    node := NewVoidNode(true)
    
    sn := ServerNode {
        &node,
        serviceURL,
    }
    
    graph.AddNode(&sn)
    
    node.AddAcceptCondition(func (idNode1 int64, idNode2 int64) bool {
        
        var node2 interface{}
        node2, _ = graph.GetNode(idNode2)
        _, ok := node2.(*ClientNode) //Type assertion
        
        return ok
        
    })
    
    node.AddConnectionAction(func (idNode1 int64, idNode2 int64) {
        
        conn, _, err := websocket.DefaultDialer.Dial(sn.serviceURL, nil)
        if err != nil {
            log.Fatal("dial:", err)
            fmt.Println("client: " + string(err.Error()))
        }
        // Clean up on exit from this goroutine
        fmt.Println("client: Connected")
        // Loop reading messages. Send each message to the channel.
        
        
        go func () {
            
            node2, _ := graph.GetNode(idNode2)
            
            cn, _ := node2.(*ClientNode)
            
            defer conn.Close()
            for {
                fmt.Println("client: readstart")
                _, m, err := conn.ReadMessage()
                if err != nil {
                    log.Fatal("read:", err)
                    return
                }
                fmt.Println("client: " + string(m))
                
                cn.sendChannel <- m
                
            }
            
        }()
        
        go func() {
            
            node2, err := graph.GetNode(idNode2)
            if err != nil {
                fmt.Println("Error: "+ err.Error())
                return
            }
            
            cn, ok := node2.(*ClientNode)
            
            if !ok {
                fmt.Println("client: Cast problem")
                return
            }
            
            fmt.Println("client: waiting for receiveChannel")
            
            for cn.receiveChannel == nil {//BAD
                
            }
            
            for message := range cn.receiveChannel {
                fmt.Println("client: send message" + string(message))
                if err := cn.wsConn.WriteMessage(websocket.TextMessage, message); err != nil {
                    break
                }
                
                fmt.Println("Message sent: " + string(message))
        
            }
            
            fmt.Println("client: close receiving")
            
            
        }()
        
        
        
    })
    
    
    return &sn

}


*/