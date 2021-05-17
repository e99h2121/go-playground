package main
import (
    "flag"
    "fmt"
    "strings"
)
func main() {
    flag.Parse()
    var name string = flag.Args()[0]
    fmt.Println(name)
    fmt.Println(name[:strings.LastIndex(name, ".gz")]) 
}

