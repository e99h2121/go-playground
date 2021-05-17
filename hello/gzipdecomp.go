package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
    flag.Parse()
    var name string = flag.Args()[0]
    
	dist, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer dist.Close()

	src, err := os.Open("test.txt.gz")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	gr, err := gzip.NewReader(src)
	if err != nil {
		panic(err)
	}
	defer gr.Close()

	if _, err := io.Copy(dist, gr); err != nil {
		panic(err)
	}
}
