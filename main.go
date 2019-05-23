package main

import (
	"fmt"
	"os"

	"github.com/jpillora/opts"
	initPkg "github.com/jpillora/opts-cli/internal/init"
)

var (
	version = "dev"
	date    = "na"
	commit  = "na"
)

type root struct {
	help string
}

func main() {
	r := root{}
	o := opts.New(&r).
		EmbedGlobalFlagSet().
		Complete().
		Version(version).
		AddCommand(initPkg.New()).
		Parse()
	r.help = o.Help()
	err := o.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "run error %v\n", err)
		os.Exit(2)
	}
}

func (r *root) Run() {
	fmt.Printf("version: %s\ndate: %s\ncommit: %s\n", version, date, commit)
	fmt.Println(r.help)
}
