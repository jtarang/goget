package main

import (
	"flag"
	"fmt"
	"github.com/jaytarang92/goget"
	"os"
)

// gogetter -url=http://google.com -dst=/Users/J_Hack92/Desktop/index.html -md5sum=1234567890123456789
// gogetter -url=http://google.com -dst=/Users/J_Hack92/Desktop/index.html -sha256=1234567890123456789
func main() {
	// arguments accepted
	md5Check := flag.String("md5sum", "", "If md5 is passed it will check it before saving file.")
	sha256chk := flag.String("sha256", "", "Same as md5 but using sha256")
	url := flag.String("url", "", "url to file")
	dest := flag.String("dst", "", "the output filename")
	flag.Parse()
	var hash string
	var hashtype string
	if fmt.Sprintf("%s", *md5Check) != "" && fmt.Sprintf("%s", *sha256chk) != "" {
		fmt.Println("\tPlease choose 1. Either sha256 or md5.\n")
		os.Exit(1)
	}
	//Sprintf formats md5 check t string
	if fmt.Sprintf("%s", *md5Check) == "" && fmt.Sprintf("%s", *sha256chk) == "" {
		////run gogetter to save the file no hash check
		gogetter.SaveIt(gogetter.GoTo(*url), *dest)
	} else {
		//run gogetter to save the file with hash check
		if fmt.Sprintf("%s", *md5Check) != "" {
			hash = fmt.Sprintf("%s", *md5Check)
			hashtype = "md5"
		} else {
			hash = fmt.Sprintf("%s", *sha256chk)
			hashtype = "sha256"
		}
		input := gogetter.GoTo(*url)
		gogetter.SaveIt(gogetter.HashCheck(input, hash, gogetter.Hash2str(input, hash, hashtype)), *dest)
	}
}
