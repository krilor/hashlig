package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

var file = flag.String("file", "", "File to compute hash for.")
var hash = flag.String("hash", "", "Hex hash value to compare against.")
var algorithm = flag.String("algorithm", "", "Hash algorithm. One of ( sha256 ).")

func isAlgorithm(input string) bool {
	switch strings.ToUpper(input) {
	case
		"SHA256":
		return true
	}
	return false
}

func isHash(input string) bool {
	matched, err := regexp.MatchString("[0-9a-f]{64}", strings.ToLower(input))
	if err != nil {
		return false
	}

	return matched
}

func isFile(input string) bool {
	if _, err := os.Stat(input); !os.IsNotExist(err) {
		return true
	}

	return false
}

func init() {
	flag.StringVar(file, "f", "", "Short version of file")
	flag.StringVar(hash, "h", "", "Short version of hash")
	flag.StringVar(algorithm, "a", "", "Short version of algorithm")
}

func main() {

	flag.Parse()

	var f, h, a, c string

	if *file != "" || *hash != "" || *algorithm != "" {
		f = *file
		h = strings.ToLower(*hash)
		a = strings.ToUpper(*algorithm)
	} else if len(os.Args) > 1 { // no flags but args
		args := os.Args[1:]
		for _, v := range args {
			if isAlgorithm(v) {
				a = v
			} else if isHash(v) {
				h = strings.ToLower(v)
			} else if isFile(v) {
				f = v
			}
		}
	}

	// Find latest file if no file specified
	if f == "" {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		var modTime time.Time
		for _, fi := range files {
			if fi.Mode().IsRegular() {
				if fi.ModTime().After(modTime) {
					modTime = fi.ModTime()
					f = fi.Name()
				}
			}
		}
	}

	// Default algorithm
	if a == "" {
		a = "SHA256"
	} else if !isHash(a) {
		fmt.Printf("Invalid algorithm %s", a)
		os.Exit(1)
	}

	fmt.Printf("File: %s,  Hash: %s, Algorithm: %s\n", f, h, a)

	if isFile(f) {

		osf, err := os.Open(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer osf.Close()

		hasher := sha256.New()
		if _, err := io.Copy(hasher, osf); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		c = hex.EncodeToString((hasher.Sum(nil)))

	} else {
		fmt.Print("Could not open file.")
		os.Exit(1)
	}

	if h != "" {
		if h != c {
			fmt.Printf("Input hash (%s) does not match file hash (%s)\n", h, c)
			os.Exit(1)
		} else {
			fmt.Printf("File hash matches input hash\n")
			os.Exit(0)
		}
	} else {
		fmt.Printf("Hash (%s) computed : %s\n", a, c)
		os.Exit(0)
	}
}
