package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("EXECUTABLE")
    fmt.Println(len(os.Args), os.Args)
}