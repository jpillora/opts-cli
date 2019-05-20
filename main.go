package main

import (
	"fmt"
	"os"

	"github.com/jpillora/opts"
	"github.com/jpillora/opts-cli/internal/initopts"
)

var (
	Version string = "dev"
	Date    string = "na"
	Commit  string = "na"
)

type root struct {
	parsedOpts opts.ParsedOpts
}

func main() {
	r := root{}
	r.parsedOpts = opts.New(&r).
		EmbedGlobalFlagSet().
		Complete().
		Version(Version).
		Call(initopts.Register).
		Parse()
	err := r.parsedOpts.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "run error %v\n", err)
		os.Exit(2)
	}
}

func (rt *root) Run() {
	fmt.Printf("Version: %s\nDate: %s\nCommit: %s\n", Version, Date, Commit)
	fmt.Println(rt.parsedOpts.Help())
}
