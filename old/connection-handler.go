package main
/*
import (
    "encoding/json"
    "fmt"
    "net/http"
)

//REST API HANDLERS

type connection_get_request struct {
    SocketId int64 `json:"socket-id"`
}

func ConnectionGETHandler(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    
    decoder := json.NewDecoder(r.Body)
    t := new(connection_get_request) 
    
    err := decoder.Decode(&t)
    
    if err != nil {

        fmt.Fprintln(w, `{
            "status": "error",
            "message": "Unable to decode request."
        }`)

        return 
    }
    
    //
    socket, _  := sc.GetSocket(t.SocketId)
    connectedIds:= socket.GetConnected()
    
    if err := json.NewEncoder(w).Encode(connectedIds); err != nil {
        panic(err)
    }
    
    
}

type connection_post_request struct {
    SocketId1 int64 `json:"socket-id-endpoint-1"`
    SocketId2 int64 `json:"socket-id-endpoint-2"`
}

func ConnectionPOSTHandler(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    
    decoder := json.NewDecoder(r.Body)
    t := new(connection_post_request) 
    
    err := decoder.Decode(&t)
    
    if err != nil {

        fmt.Fprintln(w, `{
            "status": "error",
            "message": "Unable to decode request."
        }`)

        return 
    }
    
    sc.CreateConnection(t.SocketId1, t.SocketId2)
    
    fmt.Fprintln(w,`{
        "status": "success"
    }`)
    
}

type connection_delete_request struct {
    SocketId1 int64 `json:"socket-id-endpoint-1"`
    SocketId2 int64 `json:"socket-id-endpoint-2"`
}

func ConnectionDELETEHandler(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    
    decoder := json.NewDecoder(r.Body)
    t := new(connection_delete_request) 
    
    err := decoder.Decode(&t)
    
    if err != nil {

        fmt.Fprintln(w, `{
            "status": "error",
            "message": "Unable to decode request."
        }`)

        return 
        
    }
    
    sc.DeleteConnection(t.SocketId1, t.SocketId2)
    
    fmt.Fprintln(w,`{
        "status": "success"
    }`)
    
}

*/