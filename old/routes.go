package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    /*
    Route {
        "SocketsGETHandler",
        "GET",
        "/sockets",
        SocketsGETHandler,
    },
    Route {
        "SocketsPOSTHandler",
        "POST",
        "/sockets",
        SocketsPOSTHandler,
    },
    Route {
        "SocketsDELETEHandler",
        "DELETE",
        "/sockets",
        SocketsDELETEHandler,
    },
    Route {
        "SocketsDELETEHandler",
        "GET",
        "/sockets",
        SocketsDELETEHandler,
    },
    
    Route {
        "ConnectionGETHandler",
        "GET",
        "/sockets/connections",
        ConnectionGETHandler,
    },
    
    Route {
        "ConnectionPOSTHandler",
        "POST",
        "/sockets/connections",
        ConnectionPOSTHandler,
    },
    
    Route {
        "ConnectionDELETEHandler",
        "DELETE",
        "/sockets/connections",
        ConnectionDELETEHandler,
    },
    
    //WEBSOCKETS
    Route {
        "WebSocketHandler",
        "GET",
        "/websocket",
        WebSocketHandler,
    },
    
    Route{
        "ClientWebSocketHandler",
        "GET",
        "/websocketclient",
        ClientWebSocketHandler,
    },
    */
    
    
}