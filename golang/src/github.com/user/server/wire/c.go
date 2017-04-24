package wire 

import (
    "errors"
    "math"
)

const (
    connectionCapacity int = 10    
)

type Connection struct {
    queue [connectionCapacity]interface{}
    locker sync.Locker
    inPort *InPort
    outPort *OutPort
}
func (conn *Connection) Enqueue(IP interface{}) error {
    
    conn.locker.Lock()
    
    if len(conn.queue) == connectionCapacity - 1 {
        conn.locker.Unlock()
        return errors.New("Queue is full.") 
    }

    queue := append(queue, 1)
    conn.locker.Unlock()
    
} 

func (conn *Connection) Dequeue() (interface{}, error) {
    
    conn.locker.Lock()
    if len(conn.queue) == 0 {
        conn.locker.Unlock()
        return nil, errors.New("Queue is empty.")
    }
    
    value := conn.queue[0]
    conn.queue = conn.queue[1:]
    conn.locker.Unlock()
    return value
}

type InPort struct {
    graph *Graph
    nodeId int64
}

func (inPort * InPort) Read() interface {} error {
    
}

type OutPort struct {
    graph *Graph
    nodeId int64
}

func (inPort * InPort) Send(interface {}) error {
    
}


type Component struct {
    inputs []interface{}
    
}