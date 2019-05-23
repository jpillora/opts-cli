package init

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jpillora/opts"
)

type initOpts struct {
	Directory      string `opts:"mode=arg,help=output directory"`
	Force          bool   `opts:"help=Do not check if output dir is empty"`
	SrcControlHost string `opts:"help=Repo domain or host"`
	Owner          string `opts:"help=Repo owner. Defaults to $USERNAME"`
	Name           string `opts:"help=Project name or path. Defaults to current directory"`
}

func New() opts.Opts {
	in := &initOpts{
		SrcControlHost: "github.com",
		Directory:      ".",
	}
	if us, err := user.Current(); err == nil {
		in.Owner = us.Name
	}
	if di, err := os.Getwd(); err == nil {
		in.Name = path.Base(di)
	}
	return opts.New(in).Name("init")
}

func (in *initOpts) Run() error {
	if !in.Force {
		dir, err := os.OpenFile(in.Directory, os.O_APPEND, 0755)
		if err != nil {
			err = os.MkdirAll(in.Directory, 0755)
			if err != nil {
				return err
			}
		} else {
			names, err := dir.Readdirnames(1)
			if len(names) > 0 {
				return errors.New("output directory not empty")
			}
			if err != io.EOF {
				return err
			}
		}
	}
	cmdN := strings.Split(in.Name, "/")
	data := struct {
		Module  string
		Command string
		Name    string
		Owner   string
	}{
		Module:  in.SrcControlHost + "/" + in.Owner + "/" + in.Name,
		Command: cmdN[len(cmdN)-1],
		Name:    in.Name,
		Owner:   in.Owner,
	}
	fmt.Printf("#init %+v\n", data)
	for _, fi := range files {
		tmpl, err := template.New(fi.Path).Parse(fi.Tmpl)
		if err != nil {
			fmt.Printf("tmpl parse error : %v\n", err)
			continue
		}
		fmt.Printf("#%v\n", fi.Path)
		pa := filepath.Join(in.Directory, path.Dir(fi.Path))
		_ = os.MkdirAll(pa, 0755)
		pa = filepath.Join(pa, path.Base(fi.Path))
		ofi, err := os.OpenFile(pa, os.O_RDWR|os.O_CREATE, 0664)
		if err != nil {
			fmt.Printf("new file error: %v", err)
			continue
		}
		err = tmpl.Execute(ofi, data)
		if err != nil {
			fmt.Printf("tmpl exec error : %v\n", err)
			continue
		}
	}
	return nil
}

type file struct {
	Path string
	Tmpl string
}

var files = []file{
	{
		Path: "go.mod",
		Tmpl: `module {{.Module}}

go 1.12

require github.com/jpillora/opts v1.0.1
`,
	},
	{
		Path: "main.go",
		Tmpl: `package main

import (
	"fmt"

	"github.com/jpillora/opts"
	"{{.Module}}/internal/initopts"
)

var (
	version string = "dev"
	date    string = "na"
	commit  string = "na"
)

type root struct {
	help string
}

func main() {
	opts.New(&root{}).
		Name("{{.Name}}").
		EmbedGlobalFlagSet().
		Complete().
		Version(version).
		AddCommand(initopts.New()).
		Parse().
		RunFatal()
}

func (r *root) Run() {
	fmt.Printf("version: %s\ndate: %s\ncommit: %s\n", version, date, commit)
	fmt.Println(r.help)
}
`,
	},
	{
		Path: "internal/initopts/init.go",
		Tmpl: `package initopts

import (
	"fmt"

	"github.com/jpillora/opts"
)

type initOpts struct {
}

func New() opts.Opts {
	in := initOpts{	}
	return opts.New(&in).Name("init")
}

func (in *initOpts) Run() error {
	fmt.Printf("#init %+v\n", in)
	return nil
}

`,
	},
	{
		Path: ".goreleaser.yml",
		Tmpl: `# This is an example goreleaser.yaml file with some sane defaults.
# Make sure to check the documentation at http://goreleaser.com
project_name: {{.Name}}
release:
  github:
    owner: {{.Owner}} 
    name: {{.Name}}
  name_template: '{{"{{"}}.Tag}}'
  # disable: true

builds:
- 
  binary: {{.Command}}
  env:
  - CGO_ENABLED=0
  goos:
  - linux
  - darwin
  - windows
  goarch:
  - amd64
  - "386"
  ignore:
  - goos: darwin
    goarch: 386
  main: .
  ldflags:
  - -s -w -X main.version={{"{{"}}.Version}} -X main.commit={{"{{"}}.Commit}} -X main.date={{"{{"}}.Date}}

archive:
  replacements:
    386: i386
    amd64: x86_64
  format_overrides:
  - goos: windows
    format: zip
  files:
  - licence*
  - LICENCE*
  - license*
  - LICENSE*
  - readme*
  - README*
  - changelog*
  - CHANGELOG*

checksum:
  name_template: 'checksums.txt'

snapshot:
  name_template: "{{"{{"}} .Tag {{"}}"}}-next"

changelog:
  sort: asc
  filters:
    exclude:
    - '^docs:'
    - '^test:'
`,
	},
	{
		Path: ".gitignore",
		Tmpl: `{{.Command}}
dist/`,
	},
}
