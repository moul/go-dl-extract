package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/units"
)

/* FIXME:
- .xz, .gz, .bz2 uncompression
- stream curl and untar: remove intermediate .tar
*/

func useless() {
	// Hacking godep, this package is needed for cross-compile on Linux
	units.HumanSize(42)
}

func main() {
	is_bin_sh := flag.String("c", "", "Is $0 /bin/sh ?")
	flag.Parse()

	fmt.Println("is_bin_sh", *is_bin_sh)

	url := flag.String("url", "", "help message for flagname")
	verbose := flag.Bool("v", false, "Verbose")
	dest := flag.String("dest", "/", "Destination path")
	tmptar := "tmp.tar"

	flag.CommandLine.Parse(strings.Split(*is_bin_sh, " "))

	fmt.Println("verbose:", *verbose)
	fmt.Println("url:", *url)
	fmt.Println("dest:", *dest)
	fmt.Println("tail:", flag.Args())
	return

	/*
		if *verbose && *is_bin_sh {
			fmt.Println("$0 is /bin/sh")
		}
		fmt.Println(os.Args)
	*/

	resp, err := http.Get(*url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	f, err := os.Create(tmptar)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}

	nf, err := os.Open(tmptar)
	defer nf.Close()
	archive.Untar(nf, *dest, nil)
}
