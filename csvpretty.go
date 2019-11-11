package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {

	var err error
	if err = do(); err == nil {
		return
	}

	switch err.(type) {
	case errUsage:
		flag.Usage()
	}

	panic(err)
}

type byteReplacer struct {
	out  io.Writer
	from byte
	to   string
}

func (br byteReplacer) Write(b []byte) (n int, err error) {
	n, err = br.out.Write(bytes.Replace(b, []byte{br.from}, []byte(br.to), -1))
	return len(b), err
}

var input string
var output string
var overwrite bool
var debug bool

func init() {
	flag.StringVar(&input, "input", "", "input file")
	flag.StringVar(&output, "output", "", "output file")
	flag.BoolVar(&overwrite, "w", false, "overwrite input with output")
	flag.BoolVar(&debug, "debug", false, "print debug info")
}

const holder = '\x01'

type errUsage string

func (e errUsage) Error() string { return string(e) }

var missingInput errUsage = "missing input"

func do() (err error) {
	flag.Parse()

	if input == "" {
		return missingInput
	}

	if overwrite {
		output = input
	}

	var bt []byte
	if bt, err = ioutil.ReadFile(input); err != nil {
		return
	}

	var records [][]string
	if records, err = csv.NewReader(bytes.NewReader(bt)).ReadAll(); err != nil {
		return
	}

	var buf bytes.Buffer
	var tabFlags uint = 0
	var padChr byte = ' '
	if debug {
		tabFlags |= tabwriter.Debug
		padChr = '-'
	}
	tabWriter := tabwriter.NewWriter(&buf, 0, 1, 3, padChr, tabFlags)

	var tabReplacer = byteReplacer{
		out:  tabWriter,
		from: holder,
		to:   ",\t",
	}

	// tabwriter wants a \t at the end of each row too.
	var lineReplacer = byteReplacer{
		out:  tabReplacer,
		from: '\n',
		to:   "\t\n",
	}

	csvWriter := csv.NewWriter(lineReplacer)
	csvWriter.Comma = holder

	for _, row := range records {
		for i, f := range row {
			// strip any parsed spaces (which i'd like to be ignored)
			row[i] = strings.TrimSpace(f)
		}
		if err = csvWriter.Write(row); err != nil {
			return
		}
	}

	csvWriter.Flush()
	if _, err = tabWriter.Write([]byte("\t")); err != nil {
		return
	}

	if err = tabWriter.Flush(); err != nil {
		return
	}

	// trim extra space around each line
	var trimmed bytes.Buffer
	for _, line := range bytes.Split(buf.Bytes(), []byte("\n")) {
		if _, err = trimmed.Write(append(bytes.TrimSpace(line), []byte("\n")...)); err != nil {
			return
		}
	}

	var outfile io.Writer = os.Stdout

	if output != "" && output != "-" {
		var f *os.File
		if f, err = os.OpenFile(output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0700); err != nil {
			return
		}

		defer f.Close()
		outfile = f
	}

	if _, err = io.Copy(outfile, &trimmed); err != nil {
		return
	}

	return

}
