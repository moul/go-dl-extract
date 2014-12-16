package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/docker/docker/pkg/archive"
	_ "github.com/docker/docker/pkg/units" // Required for godep
)

// Usage in a Dockerfile:
//   RUN {URL}
//   RUN ADD {URL} /
//   RUN -v -url={URL}
func main() {
	// Parsing
	is_bin_sh := flag.String("c", "", "Is $0 /bin/sh ?")
	flag.Parse()

	url := flag.String("url", "", "help message for flagname")
	verbose := flag.Bool("v", false, "Verbose")
	dest := flag.String("dest", "/", "Destination path")
	tmptar := "tmp.tar"

	flag.CommandLine.Parse(strings.Split(*is_bin_sh, " "))
	if len(flag.Args()) > 0 && len(*url) == 0 {
		url = &flag.Args()[0]
	}

	if *verbose {
		fmt.Println("verbose:", *verbose)
		fmt.Println("url:", *url)
		fmt.Println("dest:", *dest)
		fmt.Println("tail:", flag.Args())
	}

	// Get
	resp, err := http.Get(*url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Create the tarball locally
	f, err := os.Create(tmptar)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Fill the tarball
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}

	// Open the tarball
	nf, err := os.Open(tmptar)
	defer nf.Close()

	// Extract the tarball
	archive.Untar(nf, *dest, nil)
}
