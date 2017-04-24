package main
/*
import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
)


//REST API HANDLERS

func SocketsGETHandler(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(sc.GetSocketsIds()); err != nil {
        panic(err)
    }
    
}

type socket_post_request struct {
    IsUnit bool `json:"is-unit"`
}

func SocketsPOSTHandler(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    
    decoder := json.NewDecoder(r.Body)
    t := new(socket_post_request) 
    
    err := decoder.Decode(&t)
    
    if err != nil {

        fmt.Fprintln(w, `{
            "status": "error",
            "message": "Unable to decode request."
        }`)

        return
    }
    
    socketId := sc.CreateSocket(t.IsUnit);

    fmt.Fprintln(w,`{
        "status": "success",
        "socket-id": ` + strconv.FormatInt(socketId, 10) + ` 
    }`)
    
}

func SocketsPUTHandler(w http.ResponseWriter, r *http.Request) {
    
}

type socket_delete_request struct {
    SocketId int64 `json:"socket-id"`
}

func SocketsDELETEHandler(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    
    decoder := json.NewDecoder(r.Body)
    t := new(socket_delete_request) 
    
    err := decoder.Decode(&t)
    
    if err != nil {

        fmt.Fprintln(w, `{
            "status": "error",
            "message": "Unable to decode request."
        }`)

        return
    }
    
    sc.RemoveSocket(t.SocketId);

    fmt.Fprintln(w,`{
        "status": "success"
    }`)
    
    
}

*/