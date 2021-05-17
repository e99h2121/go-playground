package main

import (
	"compress/gzip"
	"io"
	"os"
	"flag"
	"fmt"
	"strings"
)

func main() {
    flag.Parse()
    var name string = flag.Args()[0]
    var txt string = name[:strings.LastIndex(name, ".gz")] 

	dist, err := os.Create(txt)
	if err != nil {
		panic(err)
	}

	src, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	gr, err := gzip.NewReader(src)
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(dist, gr); err != nil {
		panic(err)
	}

	dist.Close()
	src.Close()
	gr.Close()

    if err := os.Remove(name); err != nil {
        fmt.Println(err)
    }
}
