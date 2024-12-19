package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func remove[T comparable](l []T, item T) []T {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func removeAll[T comparable](l []T, item T) []T {
	for {
		initalLength := len(l)
		l = remove(l, item)
		if len(l) == initalLength {
			break
		}
	}

	return l
}


func getItems() []string {
	cmd := exec.Command("xsdb", printtargetsFilePath)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		if errb.Len() != 0 {
			fmt.Fprintf(os.Stderr, "%v\n", errb.String())
		}
		fmt.Fprintf(os.Stderr, "error running xsdb: %v\n", err)
		os.Exit(1)
	}

	items := strings.Split(outb.String(), "\n")

	items = removeAll(items, "")

	return items
}

func resetTarget(target int) {
	cmd := exec.Command("xsdb", resetFilePath, strconv.Itoa(target))
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		if errb.Len() != 0 {
			fmt.Fprintf(os.Stderr, "%v\n", errb.String())
		}
		fmt.Fprintf(os.Stderr, "error running xsdb: %v\n", err)
		os.Exit(2)
	}
}

func main() {
	initTmpFiles()

	targetString := pick(getItems())

	if targetString == "" {
		return
	}

	var target int
	_, err := fmt.Sscanf(targetString, "%d", &target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not parse target number from selection: %v", err)
		os.Exit(3)
	}

	resetTarget(target)
}
