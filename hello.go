package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    fmt.Printf("hello, world\n")
    fmt.Print("Press 'Enter' to continue...")
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}
