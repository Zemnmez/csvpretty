package main

import (
	"bytes"
	"flag"
	"io/ioutil"
)

func main() {
	if err := do(); err != nil {
		panic(err)
	}
}

var (
	from string
	to   string
	file string
)

func do() (err error) {
	flag.Parse()

	fbt, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	return ioutil.WriteFile(file, bytes.ReplaceAll(fbt, []byte(from), []byte(to)), 0700)
}

func init() {
	flag.StringVar(&from, "from", "", "string to find")
	flag.StringVar(&file, "file", "", "file to replace in")
	flag.StringVar(&to, "to", "", "string to replace with")
}
