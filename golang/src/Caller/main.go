package main

import (
    "os/exec"
    "fmt"
    "log"
    "strconv"
)

func main() {

	cmd := exec.Command("/home/ubuntu/workspace/golang/bin/Component")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	
	
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	
	stdin.Write([]byte(strconv.Itoa(55) + "\n"))
	
    
    b1 := make([]byte, 200)
    stdout.Read(b1)

	fmt.Print(string(b1))
	
	
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	
	
}