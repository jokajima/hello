package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    fmt.Printf("hello, world\n")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(){
        fmt.Println(scanner.Text())
    }
  
}
