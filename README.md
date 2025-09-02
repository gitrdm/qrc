QR code generator for text terminals
======================================================================

  * Copyright (C) 2014-2017 SATOH Fumiyasu @ OSS Technology Corp., Japan
  * License: MIT License
  * Development home: <https://github.com/fumiyas/qrc>
  * Author's home: <https://fumiyas.github.io/>

What's this?
---------------------------------------------------------------------

This program generates QR codes in
[ASCII art](http://en.wikipedia.org/wiki/ASCII_art) or
[Sixel](http://en.wikipedia.org/wiki/Sixel) format for
text terminals, e.g., console, xterm (with `-ti 340` option to enable Sixel),
[mlterm](http://sourceforge.net/projects/mlterm/),
Windows command prompt and so on.

Use case
---------------------------------------------------------------------

You can transfer data to smartphones with a QR code reader application
from your terminal.

Usage
---------------------------------------------------------------------

`qrc` program takes a text from command-line argument or standard
input (if no command-line argument) and encodes it to a QR code.

```console
$ qrc --help
Usage: qrc [OPTIONS] [TEXT]

Options:
  -h, --help
    Show this help message
  -i, --invert
    Invert color
  -f, --format <aa|sixel>
    Output format. If omitted, auto-detects Sixel support on TTY using env hints
    (TERM contains "sixel" or "mlterm", XTERM_SIXEL=1, or QRC_SIXEL=1 to force).

Text examples:
  http://www.example.jp/
  MAILTO:foobar@example.jp
  WIFI:S:myssid;T:WPA;P:pass123;;
$ qrc https://fumiyas.github.io/
...
$ qrc 'WIFI:S:Our-ssid;T:WPA;P:secret;;'
...
```

You can get a QR code in Sixel graphics if the standard output is
a terminal and it supports Sixel.

![](qrc-demo.png)

Download
---------------------------------------------------------------------

Binary files are here for Linux, Mac OS X and Windows:

  * https://github.com/fumiyas/qrc/releases

Build from source codes
---------------------------------------------------------------------

If you have Go 1.18+ installed:

```console
$ git clone https://github.com/fumiyas/qrc
$ cd qrc
$ go build ./cmd/qrc
```

Reproducible builds (vendored dependencies)
---------------------------------------------------------------------

This project commits the `vendor/` directory and builds with `-mod=vendor`.
That means:

- You can build and test completely offline using only the pinned deps.
- Supply-chain surface is reduced; versions are locked by `go.mod`/`go.sum` and source is vendored.
- The Makefile targets (`build`, `test`, `cross`) already enforce vendor mode.

Updating dependencies (maintainers):

```console
$ make get         # runs: go mod tidy
$ make vendor      # runs: go mod vendor
$ git add go.mod go.sum vendor/
$ git commit -m "deps: update modules and refresh vendor"
```

If you prefer not to vendor in your fork, remove `-mod=vendor` in `Makefile`
and do not commit `vendor/`. However, the default in this repo is to keep
`vendor/` for predictable, offline builds.

TODO
----------------------------------------------------------------------

  * Add the following options:
    * `--format <aa|sixel>`
    * `--aa-color-scheme <ansi|windows>`
    * `--foreground-color R:G:B`
    * `--background-color R:G:B`
    * `--margin-color R:G:B`
    * `--margin-size N`
    * `--input-encoding E`
  * Timeout for tty.GetDeviceAttributes1()

Contributors
----------------------------------------------------------------------

  * Hayaki Saito (@saitoha)

Similar products
----------------------------------------------------------------------

  * Go
    * <https://godoc.org/github.com/GeertJohan/go.qrt>
  * JavaScript (Node)
    * <https://github.com/gtanner/qrcode-terminal>
  * Ruby
    * <https://gist.github.com/saitoha/10483508> (qrcode-sixel)
    * <https://gist.github.com/fumiyas/10490722> (qrcode-sixel)
    * <https://github.com/fumiyas/home-commands/blob/master/qrcode-aa>

Releases on GitHub
---------------------------------------------------------------------

This repo publishes release binaries via GitHub Actions when you push a tag
matching `v*` (e.g., `v1.2.3`). The workflow builds cross-platform binaries and
attaches them to the release with checksums.

Steps to release:

```console
$ git tag vX.Y.Z
$ git push origin vX.Y.Z
```

The release job will:

- Run `make get && make vendor && make cross`
- Upload `dist/qrc-<os>-<arch>` and `dist/SHASUMS256.txt` to the GitHub Release

