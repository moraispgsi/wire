package main

import fbpserver "fbp/server/fbp-server"

func main() {
	server := fbpserver.New()
	server.Start()
}
