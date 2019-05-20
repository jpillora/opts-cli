package main

import (
	"fmt"

	"github.com/jpillora/opts"
	"github.com/jpillora/opts-cli/internal/initopts"
)

var (
	Version string = "dev"
	Date    string = "na"
	Commit  string = "na"
)

type root struct {
	ParsedOpts opts.ParsedOpts `opts:"mode=parsedOpts"`
}

func main() {
	opts.New(&root{}).
		EmbedGlobalFlagSet().
		Complete().
		Version(Version).
		Call(initopts.Register).
		Parse().
		RunFatal()
}

func (rt *root) Run() {
	fmt.Printf("Version: %s\nDate: %s\nCommit: %s\n", Version, Date, Commit)
	fmt.Println(rt.ParsedOpts.Help())
}
