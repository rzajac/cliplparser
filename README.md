# Command line playlist parser in Go

Command line playlist parser / detector in pure Go.

# Supported playlist formats

* PLS
* ASX
* ASF
* M3U

# Installation / compilation

1. Clone this repo
2. Run `go get github.com/rzajac/plparser`
3. Run `go build cliplparser.go`
4. Enjoy

# Updating

Run

    go get -u github.com/rzajac/plparser
    go get -u github.com/rzajac/cliplparser

Compile again.

# Benchmarking

    go test -bench=".*" github.com/rzajac/plparser

# Usage

    $ ./cliplparser
    Wrong number of argumants.

    Usage:

    cliplparser [options] playlist_file | URL

    Options:
      --json       Return JSON string instead of text.
      --timeout n  Set timeout for HTTP connections, where n is number of seconds.

    Notes:
      URL must start with http:// or https://

# See

Go and see the the library this tool is using [http://github.com/rzajac/plparser](http://github.com/rzajac/plparser)

# License

Licensed under the MIT license

