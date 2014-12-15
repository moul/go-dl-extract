package main

import (
	"flag"
	"fmt"
)

func main() {
	url := ""
	flag.StringVar(&url, "url", "", "help message for flagname")
	flag.Parse()
	fmt.Printf("url: %s\n", url)

	/*
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		f, err := os.Create("my.tar.gz")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			panic(err)
		}
	*/
}
