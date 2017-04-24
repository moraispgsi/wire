package wire
/*
import (
    "net/http"
    "github.com/gorilla/websocket"
    "log"
    "sync"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type ClientNode struct {
    Node
    interfaceIRI string
    sendChannel chan []byte
    receiveChannel chan []byte
    isWSConnected bool
    wsConn *websocket.Conn
    wg sync.WaitGroup
    mutex *sync.Mutex
}

func NewClientNode(graph *Graph, interfaceIRI string) *ClientNode {
    
    node := NewVoidNode(true)
    var wg sync.WaitGroup
    var mutex = &sync.Mutex{}
    cn := ClientNode {
        node,
        interfaceIRI,
        nil,
        nil,
        false,
        nil,
        wg,
        mutex,
    }
    
    graph.AddNode(&cn)

    return &cn
    
}


func (cn *ClientNode) Upgrade(w http.ResponseWriter, r *http.Request) {

    var conn, err = upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("websocket: ", err)
        return
    }

    cn.mutex.Lock()
    
    if cn.isWSConnected {
        if cn.wsConn != nil {
            cn.wsConn.Close()
        }
    }
    
    cn.mutex.Unlock()

    cn.wg.Wait()
    
    cn.wsConn = conn
    cn.isWSConnected = true

    cn.sendChannel = make(chan []byte)
    cn.receiveChannel = make(chan []byte)
    
    cn.wg.Add(2)
    go cn.sendHandler()
    go cn.receiveHandler()

}

func (cn *ClientNode) sendHandler() {
    
    defer cn.wg.Done()
    for message := range cn.sendChannel {
        
        if err := cn.wsConn.WriteMessage(websocket.TextMessage, message); err != nil {
            break
        }

    }

}

func (cn *ClientNode) receiveHandler() {
    
    defer cn.wg.Done()
    
    for cn.wsConn != nil {
        
         _, message, err := cn.wsConn.ReadMessage()
        if ; err != nil {
            break
        }

        cn.receiveChannel <- message

    }
    
    cn.mutex.Lock()
    cn.isWSConnected = false
    cn.wsConn = nil
    close(cn.sendChannel)
    close(cn.receiveChannel)
    cn.mutex.Unlock()

}
*/