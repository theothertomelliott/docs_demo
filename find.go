package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var dir = flag.String("dir", ".", "Directory for search")
	flag.Parse()
	err := filepath.Walk(*dir, visit)
	if err != nil {
		log.Fatal(err)
	}
}

func visit(path string, f os.FileInfo, err error) error {
	match, err := filepath.Match(flag.Arg(0), filepath.Base(path))
	if err != nil {
		return err
	}
	if match {
		fmt.Println(path)
	}
	return nil
}
