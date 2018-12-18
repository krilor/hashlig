package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	fpath := flag.String("file", "", "File to compute hash for")
	flag.Parse()
	fmt.Printf("%s\n", *fpath)

	f, err := os.Open(*fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", h.Sum(nil))
}
