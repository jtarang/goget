[![GoDoc](https://godoc.org/github.com/jaytarang92/goget?status.svg)](https://godoc.org/github.com/jaytarang92/goget)

# goget

goget lets you download files from the internet with checksum verification for md5 and sha256.

If md5/sha256 don't match it will not save it to local drive. It does the verifcation before placing it in local storage.

I will be adding more crpyto support. :)

**Think wget but cross platform w/ verification

##Example
```go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jaytarang92/goget"
)

// gogetter -url=http://www.7-zip.org/a/7z1602-src.7z -dst=7zip.7z -md5=8523200928a577cd1747d8575c4ecacf
// gogetter -url=https://storage.googleapis.com/golang/go1.7.src.tar.gz -dst=go.src.tar.gz -sha256=72680c16ba0891fcf2ccf46d0f809e4ecf47bbf889f5d884ccb54c5e9a17e1c0
func main() {
	// arguments accepted
	proxy := flag.String("proxy", "", "Proxy i.e http://company.proxy.com:8080")
	md5Check := flag.String("md5", "", "If md5 is passed it will check it before saving file.")
	sha256chk := flag.String("sha256", "", "Same as md5 but using sha256")
	url := flag.String("url", "", "url to file")
	dest := flag.String("dst", "", "the output filename")
	flag.Parse()
	// make variables for future use
	var (
		hash     string
		hashtype string
		pproxy   string
	)

	if *url == "" || *dest == "" {
		fmt.Println("  Insufficient amount of arguments! Please check if dst and url were passed!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// if md5 and sha256 are passed then exit
	if fmt.Sprintf("%s", *md5Check) != "" && fmt.Sprintf("%s", *sha256chk) != "" {
		fmt.Println("\tPlease choose 1. Either sha256 or md5.")
		os.Exit(1)
	}

	//if no proxy is defined then bypass it
	if *proxy == "" {
		pproxy = "proxy is undefined"
	} else {
		pproxy = fmt.Sprintf("%s", *proxy)
	}

	//Sprintf formats %x to a %s
	if fmt.Sprintf("%s", *md5Check) == "" && fmt.Sprintf("%s", *sha256chk) == "" {
		//just downloads
		fmt.Printf("\tProxy:%s\n", pproxy)
		gogetter.SaveIt(gogetter.GoTo(*url, *proxy), *dest)
	} else {
		//checks the hash passed and runs gogetter
		if fmt.Sprintf("%s", *md5Check) != "" {
			hash = fmt.Sprintf("%s", *md5Check)
			hashtype = "md5"
			if len(hash) != 32 {
				fmt.Println("\tPlease pass a valid md5 value!")
				os.Exit(1)
			}
		} else if fmt.Sprintf("%s", *sha256chk) != "" {
			hash = fmt.Sprintf("%s", *sha256chk)
			hashtype = "sha256"
			if len(hash) != 64 {
				fmt.Println("\tPlease pass a valid sha256 value!")
				os.Exit(1)
			}
		}
		input := gogetter.GoTo(*url, *proxy)
		gogetter.SaveIt(gogetter.HashCheck(input, hash, gogetter.Hash2str(input, hash, hashtype)), *dest)
	}
}
```
