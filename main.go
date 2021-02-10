package main

import (
	"extract/typless"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	apiKey := flag.String("key", "", "API Key provided by Typless service (Required)")
	inFile := flag.String("in", "", "Path to invoice file (Required)")
	outFile := flag.String("out", "", "Path to file where to save extracted invoice fields (Required)")
	template := flag.String("template", "", "Name of document type template registered in Typless service (Required)")
	flag.Parse()
	if *apiKey == "" || *inFile == "" || *outFile == "" || *template == "" {
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	resp, err := typless.Extract(*apiKey, *inFile, *template)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	err = ioutil.WriteFile(*outFile, resp, os.FileMode(0644))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
