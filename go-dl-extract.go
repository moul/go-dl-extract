package main

import (
	"flag"
	"io"
	"net/http"
	"os"

	"github.com/docker/docker/pkg/archive"
)

/* FIXME:
- verbose flag
- .xz, .gz, .bz2 uncompression
*/

func main() {
	url := flag.String("url", "", "help message for flagname")
	flag.Parse()

	resp, err := http.Get(*url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	f, err := os.Create("tmp.tar")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}

	nf, err := os.Open("tmp.tar")
	defer nf.Close()
	archive.Untar(nf, "tmp", nil)
}
