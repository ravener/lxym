package main

import (
	"flag"
	"fmt"
	"github.com/MichaelMure/go-term-markdown"
	"github.com/mattn/go-colorable"
	"github.com/ravener/lxym/pages"
	"golang.org/x/term"
	"os"
	"strings"
	"sort"
)

var (
	version = "dev"
	commit = "none"
)

// For Windows support.
var stdout = colorable.NewColorableStdout()

func getTermWidth() int {
	w, _, err := term.GetSize(0)

	if err == nil && w > 0 {
		return w
	}

	// Default
	return 100
}

var lineWidth = flag.Int("w", getTermWidth(), "Line Width. By default it tries to measure your terminal width.")
var leftPad = flag.Int("l", 0, "Left Pad")
var noPager = flag.Bool("p", false, "Disable pager. By default it will try to find a 'less' executable in PATH and pipes output to it.")
var ver = flag.Bool("v", false, "print version and exit.")

func main() {
	flag.Parse()

	if *ver {
		fmt.Printf("lxym version: %s (commit: %s)\n", version, commit)
		os.Exit(0)
	}

	pgs := []string{}

	for k, _ := range pages.Pages {
		pgs = append(pgs, k)
	}

	sort.Strings(pgs)

	if flag.NArg() < 1 {
		fmt.Println("Usage: lxym <topic>\nAvailable Topics:\n")

		for _, v := range pgs {
			fmt.Println("  - " + v)
		}

		os.Exit(1)
	}

	key := strings.ToLower(flag.Arg(0))

	src, ok := pages.Pages[key]

	if !ok {
		fmt.Printf("Topic '%s' does not exist.\nAvailable Topics:\n", key)
		for _, v := range pgs {
			fmt.Println("  - " + v)
		}

		os.Exit(1)
	}

	output := markdown.Render(src, *lineWidth, *leftPad)

	if *noPager {
		stdout.Write(output)
	} else if ok := Pager(output); !ok {
		// If something went wrong with the pager or it wasn't found
		// Just fallback to writing to stdout.
		stdout.Write(output)
	}
}
