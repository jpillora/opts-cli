package main

import (
	"github.com/jpillora/opts"
	"github.com/jpillora/opts-cli/internal/genmd"
	"github.com/jpillora/opts-cli/internal/initopts"
)

var (
	Version string = "dev"
	Date    string = "na"
	Commit  string = "na"
)

type root struct{}

func main() {
	// Create and config flag stuffer
	ro := opts.New(&root{}).Name("opts-cli").
		EmbedGlobalFlagSet().Complete().Version(Version)
	// Subcommand registration pattern
	initopts.Register(ro)
	genmd.Register(ro)
	// Parse command line and run command
	ro.Parse().RunFatal()
}
