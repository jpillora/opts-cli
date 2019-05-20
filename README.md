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

https://github.com/jpillora/opts-cli/releases

``` sh
$ ./opts-cli init -h

  Usage: opts-cli init [options] <owner> <name>

  Options:
  --src-control-host, -s  default github.com
  --package, -p
  --command, -c
  --directory, -d         output directory (default .)
  --help, -h              display help
  ```
  
  ``` sh
  $ ./opts-cli init -d my-cli ghuser clis-ftw
#init {Module:github.com/ghuser/clis-ftw Command:clis-ftw Name:clis-ftw Owner:ghuser}
#go.mod
#main.go
#internal/initopts/init.go
#.goreleaser.yml
#.gitignore
$ cd my-cli/
$ go build
$ ./clis-ftw 
Version: dev
Date: na
Commit: na

  Usage: clis-ftw [options] <command>

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
