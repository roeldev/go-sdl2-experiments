SDL2 toolkit for Go
===================
[![Latest release][latest-release-img]][latest-release-url]
[![Build status][build-status-img]][build-status-url]
[![Go Report Card][report-img]][report-url]
[![Documentation][doc-img]][doc-url]

[latest-release-img]: https://img.shields.io/github/release/roeldev/go-sdl2-kit.svg?label=latest

[latest-release-url]: https://github.com/go-pogo/sdlkit/releases

[build-status-img]: https://github.com/go-pogo/sdlkit/workflows/Go/badge.svg

[build-status-url]: https://github.com/go-pogo/sdlkit/actions?query=workflow%3ATest

[report-img]: https://goreportcard.com/badge/github.com/go-pogo/sdlkit

[report-url]: https://goreportcard.com/report/github.com/go-pogo/sdlkit

[doc-img]: https://godoc.org/github.com/go-pogo/sdlkit?status.svg

[doc-url]: https://pkg.go.dev/github.com/go-pogo/sdlkit


Package `sdlkit` contains basic building blocks when using `SDL2` with go. It
uses `github.com/veandco/go-sdl2` at its base and builds upon it. This package
is not a game or physics engine.

```sh
go get github.com/go-pogo/sdlkit
```

```go
import "github.com/go-pogo/sdlkit"
```

## SDL2 installation on Windows

- install a gcc compiler (for example `tdm-gcc`
  from https://jmeubank.github.io/tdm-gcc/download/)
- make sure it is added to your PATH environment variable, test this by
  running `gcc -v`
- run `go get -v github.com/veandco/go-sdl2/sdl@master`
- run `go mod vendor` to create/sync vendor directory
- run `git submodule add https://github.com/veandco/go-sdl2-libs.git .go-sdl2-libs`
- copy/symlink `.go-sdl2-libs`
  to `vendor/github.com/veandco/go-sdl12/.go-sdl2-libs`
- build/run your project with the `-tags static` build flag
- optionally add the `-v -x` build flags to output the files that are compiled

## Documentation
Additional detailed documentation is available at [pkg.go.dev][doc-url]

## Created with
<a href="https://www.jetbrains.com/?from=roeldev" target="_blank"><img src="https://pbs.twimg.com/profile_images/1206615658638856192/eiS7UWLo_400x400.jpg" width="35" /></a>

## License
Copyright © 2020 [Roel Schut](https://roelschut.nl). All rights reserved.

This project is governed by a BSD-style license that can be found in
the [LICENSE](LICENSE) file.
