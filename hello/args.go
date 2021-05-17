package main
import (
    "flag"
    "fmt"
)
func main() {
    flag.Parse()
    var name string = flag.Args()[0]
    fmt.Println(name)
}
