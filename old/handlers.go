package main
/*
import (
    "net/http"
    "github.com/gorilla/websocket"
    "fmt"
    "time"
    "log"
)

func ClientWebSocketHandler (w http.ResponseWriter, r *http.Request) {
    go func() {
      
        // Dial with Gorilla package. The x/net/websocket package has issues.
        fmt.Println("HERE")
        c, _, err := websocket.DefaultDialer.Dial("wss://wire-moraispgsi.c9users.io/websocket", nil)
        if err != nil {
            log.Fatal("dial:", err)
            fmt.Println("client: " + string(err.Error()))
        }
        // Clean up on exit from this goroutine
        defer c.Close()
        // Loop reading messages. Send each message to the channel.
        for {
            _, m, err := c.ReadMessage()
            if err != nil {
                log.Fatal("read:", err)
                return
            }
            fmt.Println("client: " + string(m))
        }
    }()
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
    
    var conn, err = upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    go func(conn *websocket.Conn) {
		for {
		    
			_, p, err := conn.ReadMessage()
			if err != nil {
				conn.Close()
				return
			}
			
			if string(p) == "create-socket" {
			    sc.CreateSocket(true)
			    
			    json := []int64{}

                for _, id := range sc.GetSocketsIds() {
                    
                    json = append(json, id)
                    
                }
    
    		    conn.WriteJSON(json)
    			}
			
			fmt.Println("Mensagem recebida:" + string(p))
		}
	}(conn)
    
    
    
	go func(conn *websocket.Conn) {
		ch := time.Tick(5 * time.Second)

		for range ch {
			
			if conn == nil {
			    return
			}
			
			json := []int64{}

            for _, id := range sc.GetSocketsIds() {
                
                json = append(json, id)
                
            }

		    conn.WriteJSON(json)
		}
	}(conn)
	
	//
	
	// Create channel to receive messages from all connections
        //messages := make(chan []byte)

// Run a goroutine for each URL that you want to dial.
    
    
    	
	
	
	
}
*/