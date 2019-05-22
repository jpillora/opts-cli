<p align="center">
<img width="443" alt="logo" src="https://user-images.githubusercontent.com/633843/57529538-84a22780-7378-11e9-9235-312633dc125e.png"><br>
<b>A Go (golang) package for building productive command-line interfaces</b><br><br>
<a href="https://godoc.org/github.com/jpillora/opts#Opts" rel="nofollow">
	<img src="https://camo.githubusercontent.com/42566bdba17f1a0c86c1a1de859d6ab70bde1457/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f6a70696c6c6f72612f6f7074733f7374617475732e737667" alt="GoDoc" data-canonical-src="https://godoc.org/github.com/jpillora/opts?status.svg" style="max-width:100%;">
</a>
<a href="https://circleci.com/gh/jpillora/opts" rel="nofollow">
	<img src="https://camo.githubusercontent.com/34202387888c6b05f640653a29bb1e204f5a9e19/68747470733a2f2f636972636c6563692e636f6d2f67682f6a70696c6c6f72612f6f7074732e7376673f7374796c653d736869656c6426636972636c652d746f6b656e3d36396566396336616330643863656263623335346262383563333737656365666637376266623162" alt="CircleCI" data-canonical-src="https://circleci.com/gh/jpillora/opts.svg?style=shield&amp;circle-token=69ef9c6ac0d8cebcb354bb85c377eceff77bfb1b" style="max-width:100%;">
</a>
</p>

---

A CLI for a library; because creating command-line interfaces should be really simple:

## Install
Download binary from https://github.com/jpillora/opts-cli/releases

or

Build from source `go get github.com/jpillora/opts-cli`

`opts-cli init -h`
<!--tmpl,code=plain:go run main.go init -h -->
``` plain 

  Usage: opts-cli init [options] <directory>

  output directory

  Options:
  --force, -f             Do not check if output dir is empty
  --src-control-host, -s  Repo domain or host (default github.com)
  --owner, -o             Repo owner. Defaults to $USERNAME (default garym)
  --name, -n              Project name or path. Defaults to current directory (default opts-cli)
  --help, -h              display help

```
<!--/tmpl-->

## Using

`opts-cli init --owner jpillora --name opts-golangsyd talk`
<!--tmpl,code=plain:go run main.go init --owner jpillora --name opts-golangsyd talk -->
``` plain 
#init {Module:github.com/jpillora/opts-golangsyd Command:opts-golangsyd Name:opts-golangsyd Owner:jpillora}
#go.mod
#main.go
#internal/initopts/init.go
#.goreleaser.yml
#.gitignore
```
<!--/tmpl-->

```
cd talk
go build
./opts-golangsyd -h
```
<!--tmpl,code=plain:cd talk; go run main.go -h -->
``` plain 

  Usage: opts-golangsyd [options] <command>

  Options:
  --version, -v    display version
  --help, -h       display help

  Completion options:
  --install, -i    install bash-completion
  --uninstall, -u  uninstall bash-completion

  Commands:
  · init

  Version:
    dev

```
<!--/tmpl-->

## CircleCI

The CircleCI config with run goreleaser is the commit is tagged.

GITHUB_TOKEN: When setting up circleci an env-var with the personal access token is needed. https://circleci.com/gh/jplillora/opts-cli/edit#env-vars

