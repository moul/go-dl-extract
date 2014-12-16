package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/pkg/archive"
	_ "github.com/docker/docker/pkg/units" // Required for godep
)

var url string
var verbose bool
var dest string
var tmptar string
var is_bin_sh string
var excludes string

func init() {
	// flags
	flag.StringVar(&is_bin_sh, "c", "", "Is $0 /bin/sh ?")
	flag.Parse()

	flag.StringVar(&url, "url", "", "help message for flagname")
	flag.BoolVar(&verbose, "v", false, "Verbose")
	flag.StringVar(&dest, "dest", "/", "Destination path")
	flag.StringVar(&excludes, "excludes",
		"sys|etc/hosts|etc/resolv.conf|proc", "Excludes")
	tmptar = "/tmp.tar"

	flag.CommandLine.Parse(strings.Split(is_bin_sh, " "))
	if len(flag.Args()) > 0 && len(url) == 0 {
		url = flag.Args()[0]
	}

	// logging
	log.SetOutput(os.Stderr)
	//log.SetFormatter(log.TextFormatter)
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	log.WithFields(log.Fields{
		"verbose": verbose,
		"url":     url,
		"dest":    dest,
		"tail":    flag.Args(),
	}).Debug("Args")

}

// Usage in a Dockerfile:
//   RUN {URL}
//   RUN ADD {URL} /
//   RUN -v -url={URL}
func main() {
	// Download
	log.Debugf("Downloading '%s' to '%s'", url, tmptar)
	resp, err := http.Get(url)
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

	// Extract
	log.Debugf("Extracting '%s' to '%s'", tmptar, dest)
	f, err = os.Open(tmptar)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = archive.Untar(f, dest, &archive.TarOptions{
		Excludes: strings.Split(excludes, "|"),
	})
	if err != nil {
		panic(err)
	}
}
