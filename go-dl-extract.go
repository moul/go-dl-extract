package main

import (
	"crypto/md5"
	"flag"
	"fmt"
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
var is_bin_sh string
var excludes string
var wanted_checksum string

func init() {
	// flags
	flag.StringVar(&is_bin_sh, "c", "", "Is $0 /bin/sh ?")
	// When go-dl-extract is used as a scratch image replacement,
	// the arguments will looks like $0 -c "..."
	flag.Parse()

	flag.StringVar(&url, "url", "", "The URL of the tarball")
	flag.BoolVar(&verbose, "v", false, "Increase verbosity")
	flag.StringVar(&dest, "dest", "/", "Destination path")
	flag.StringVar(&excludes, "Excluded files (separated by a pipe)",
		"sys|etc/hosts|etc/resolv.conf|proc|etc/hostname", "Excludes")
	flag.StringVar(&wanted_checksum, "md5", "",
		"If set, will raise an error if the checksums differ")

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
	log.Debugf("Downloading '%s'", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Checksum
	h := md5.New()
	t := io.TeeReader(resp.Body, h)

	// Extracting
	err = archive.Untar(t, dest, &archive.TarOptions{
		Excludes: strings.Split(excludes, "|"),
	})
	if err != nil {
		panic(err)
	}

	if len(wanted_checksum) > 0 {
		if wanted_checksum == fmt.Sprintf("%x", h.Sum(nil)) {
			fmt.Printf("MD5 checksum: %x (matches)\n", h.Sum(nil))
		} else {
			fmt.Printf("MD5 checksums differ, it is %x and it should be %s\n", h.Sum(nil), wanted_checksum)
			os.Exit(-1)
		}
	} else {
		fmt.Printf("MD5 checksum: %x", h.Sum(nil))
	}
}
