package main

import (
	"bytes"
	"io"
	"os"

	_ "embed"
)

//go:embed tcl-scripts/printtargets.tcl
var printtargetsFileContents []byte

var printtargetsFilePath string

//go:embed tcl-scripts/reset.tcl
var resetFileContents []byte

var resetFilePath string

func initTmpTclFile(name string, contents []byte) string {
	f, err := os.CreateTemp("", name + "-*.tcl")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.Copy(f, bytes.NewReader(contents))
	if err != nil {
		panic(err)
	}

	return f.Name()
}

func initTmpFiles() {
	printtargetsFilePath = initTmpTclFile("printtargets.tcl", printtargetsFileContents)
	resetFilePath = initTmpTclFile("resert.tcl", resetFileContents)
}
