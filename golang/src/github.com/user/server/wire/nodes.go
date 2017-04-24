package wire

import (
	"strconv"
	"fmt"
	"errors"
)

//Connection Error
type ConnectionError struct {
	error string
	idNode1 int64
	idNode2 int64
}

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("Node connection error: " + e.error)
}


type Connection struct{
	id		int64
	graph	*Graph
	idNode1	int64
	idNode2	int64
}

func (c *Connection) Disconnect() {
	c.graph.Disconnect(c.idNode1, c.idNode2)
}



//Node
type Node struct {
	//package protected
	id						int64
	graph					*Graph
	isUnit              	bool
	connections         	map[int64]int64//key is node 'to' id; value is connection id
	acceptConditions    	[]func (int64, int64) bool
	connectionActions   	[]func (int64, int64)
	disconnectionActions	[]func (int64, int64)
}

func NewNode(isUnit bool) Node {
	
	return Node{
		id: -1,
		graph: nil,
		isUnit: isUnit,
		connections: make(map[int64]int64),
		acceptConditions: []func (int64, int64) bool {},
		connectionActions: []func (int64, int64) {},
		disconnectionActions: []func (int64, int64) {},
	}
	
}

func (node *Node) addConnection(idNode int64, connId int64) {
	node.connections[idNode] = connId
}

func (node *Node) getConnectionId(idNode int64) (int64, bool) {
	
	connId, ok := node.connections[idNode]
	return connId, ok
}

func (node *Node) deleteConnection(idNode int64) {
	
	delete(node.connections, idNode)
	
}

func (node *Node) hasOwner() bool {
	return node.id != -1
}

func (node *Node) setOwner(graph *Graph, id int64) {
	node.id = id
	node.graph = graph
	node.connections = make(map[int64]int64)
}

func (node *Node) accept(to int64) bool {
 
    for _,acceptFunc := range node.acceptConditions {
        
        if !acceptFunc(node.id, to) {
            return false
        }
            
    }
    
    return true
    
}

func (node *Node) requestConnection(to int64) bool {

    return node.accept(to)
}

func (node *Node) runConnectionActions(to int64) {
	
	for _,connectionActionFunc := range node.connectionActions {
        connectionActionFunc(node.id, to)
    }
	
} 

func (node *Node) runDisconnectionActions(to int64) {
	
	for _,disconnectionActionFunc := range node.disconnectionActions {
        disconnectionActionFunc(node.id, to)
    }
	
} 



func (node *Node) AddAcceptCondition(handler func (int64, int64) bool) {
	
	node.acceptConditions = append(node.acceptConditions, handler)
	
}

func (node *Node) AddConnectionAction(handler func (int64, int64)) {
	node.connectionActions = append(node.connectionActions, handler)
}

func (node *Node) AddDisconnectionAction(handler func (int64, int64)) {
	node.disconnectionActions = append(node.disconnectionActions, handler)
}

func (node *Node) GetId() int64 {
	return node.id
}

func (node *Node) GetGraph() (*Graph, error) {
	
	if !node.hasOwner() {
		return nil, errors.New("This node doesnt have a graph owner")
	}
	
	return node.graph, nil
}

func (node *Node) GetConnected() []int64 {
	
	if !node.hasOwner() {
		return []int64 {}
	}
	
	connectedIds := []int64 {}
	
	for id := range node.connections {
		connectedIds = append(connectedIds, id)
	}
	
	return connectedIds
	
}	









type GraphError struct {
	error string
}

func (e *GraphError) Error() string {
	return fmt.Sprintf("Container error: " + e.error)
}

type Graph struct {
	nodesMap map[int64]*Node
	connections map[int64]*Connection
	nodesCurrentId int64 //needs to be atomic
	connectionsCurrentId int64 //needs to be atomic
}

func NewGraph() Graph {
	var graph Graph
	graph.nodesMap = make(map[int64]*Node)
	graph.connections = make(map[int64]*Connection)
	graph.nodesCurrentId = 1
	graph.connectionsCurrentId = 1
	return graph
}

func (graph * Graph) AddNode(node *Node) int64 {
	id := graph.nodesCurrentId
	node.setOwner(graph, id)
	graph.nodesMap[id] = node
	graph.nodesCurrentId++
	return id
}

func (graph * Graph) HasNode(id int64) bool {

	_, ok := graph.nodesMap[id];
	return ok

}

func (graph * Graph) GetNode(id int64) (*Node, error) {
	
	if _, ok := graph.nodesMap[id]; !ok {
        return nil, &GraphError {error: "The node id "+ strconv.FormatInt(id, 10) +" does not exists."}
    }
	
	return graph.nodesMap[id], nil
	
}

func (graph *Graph) GetIds() []int64 {
	
	nodeIds := make([]int64, len(graph.nodesMap))

    i := 0
    for k := range graph.nodesMap {
        nodeIds[i] = k
        i++
    }
    
    return nodeIds
    
}

func (graph *Graph) RemoveNode(id int64) {
	
	delete(graph.nodesMap, id)
	
	//TODO set no owner
	//TODO remove connections and run disconnect handlers
}

func (graph *Graph) HasConnection(idNode1 int64, idNode2 int64) bool {
	
	if !graph.HasNode(idNode1) || !graph.HasNode(idNode2) {
		return false
	}
	
	node1,_ := graph.GetNode(idNode1)
	node2,_ := graph.GetNode(idNode2)
	
	_, ok1 := node1.getConnectionId(idNode2)
	_, ok2 := node2.getConnectionId(idNode1)
	
	if ok1 && ok2 {
		return true
	}
	
	return false

}

func (graph *Graph) Connect(idNode1 int64, idNode2 int64) error  {
	
	if !graph.HasNode(idNode1) || !graph.HasNode(idNode2) {
		return &ConnectionError{error: "Both nodes must exist in the graph"}
	}
	
	if graph.HasConnection(idNode1, idNode2) {
		return &ConnectionError {
			error: "Connection already exists",
			idNode1: idNode1,
			idNode2: idNode2,
		}
	}
	
	node1,_ := graph.GetNode(idNode1)
	if !node1.requestConnection(idNode2) {
		return &ConnectionError {
			error: "Node with Id" + strconv.FormatInt(idNode1, 10) + " conditions were not met.",
			idNode1: idNode1,
			idNode2: idNode2,
		}
	}
	
	node2,_ := graph.GetNode(idNode2)
	if !node2.requestConnection(idNode1) {
		return &ConnectionError { 
			error: "Node with Id" + strconv.FormatInt(idNode2, 10) + " conditions were not met.",
			idNode1: idNode1,
			idNode2: idNode2,
		}
	}
	
	connectionId := graph.connectionsCurrentId 
	graph.connectionsCurrentId ++
	connection := Connection {
		id:connectionId, 
		graph: graph,  
		idNode1: idNode1,
		idNode2: idNode2,
	}
	graph.connections[connectionId] = &connection
	
	node1.addConnection(idNode2, connectionId)
	node2.addConnection(idNode1, connectionId)
	
	node1.runConnectionActions(idNode2)
	node2.runConnectionActions(idNode1)
	
	return nil
}

func (graph *Graph) Accepts(idNode1 int64, idNode2 int64) bool {
	
	if graph.HasConnection(idNode1, idNode2) {
		return false
	}
	
	node1, err := graph.GetNode(idNode1)
	
	if err != nil {
		return false
	}
	
	node2, err := graph.GetNode(idNode2)
	
	if err != nil {
		return false
	}
	
	return node1.accept(idNode2) && node2.accept(idNode1)

}

func (graph *Graph) Disconnect(idNode1 int64, idNode2 int64) {
	
	if !graph.HasNode(idNode1) || !graph.HasNode(idNode2) {
		return
	}
	
	//TODO: test if both are connected
	
	node1,_ := graph.GetNode(idNode1)
	node2,_ := graph.GetNode(idNode2)
	connectionId,_ := node1.getConnectionId(idNode2)
	node1.deleteConnection(idNode2)
	node2.deleteConnection(idNode1)
	delete(graph.connections, connectionId)
	
	go node1.runDisconnectionActions(idNode2)
	go node2.runDisconnectionActions(idNode1)
	
}