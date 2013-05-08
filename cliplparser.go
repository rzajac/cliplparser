// CLI playlist parser
//
// Copyright 2013 Rafal Zajac rzajac<at>gmail<dot>com. All rights reserved.
// http://github.com/rzajac/cliplparser
//
// See also:
// http://github.com/rzajac/plparser
//
// Licensed under the MIT license

package main

import (
	"flag"
	"fmt"
	"github.com/rzajac/plparser"
	"os"
	"strconv"
	"strings"
)

var usage = `
Usage:

cliplparser [options] playlist_file | URL

Options:
  --json       Return JSON string instead of text.
  --timeout n  Set timeout for HTTP connections, where n is number of seconds.

Notes:
  URL must start with http:// or https://

`
var (
	echoJson    bool
	httpTimeout int
)

func init() {
	flag.BoolVar(&echoJson, "json", false, "Return JSON string instead of text.")
	flag.IntVar(&httpTimeout, "timeout", 5, "Set timeout in seconds for HTTP connections.")
}

func main() {

	flag.Parse()

	args := flag.Args()

	var err error
	var plr *plparser.PlaylistResp

	if len(args) != 1 {
		fmt.Println("Wrong number of argumants.")
		fmt.Print(usage)
		os.Exit(1)
	}

	if strings.HasPrefix(args[0], "http") {
		plr, err = plparser.NewPlaylistRespUrl(args[0], httpTimeout)
	} else {
		plr, err = plparser.NewPlaylistRespFile(args[0])
	}

	if err != nil {
		err = plparser.NewPlParserError(err.Error(), echoJson)
		fmt.Println(err)
		os.Exit(1)
	}

	if !(plr.StatusCode >= 200 && plr.StatusCode < 300) {
		msg := "We got " + strconv.FormatInt(int64(plr.StatusCode), 10) + " getting the URL"
		err = plparser.NewPlParserError(msg, echoJson)
		fmt.Println(err)
		os.Exit(1)
	}

	if !plr.IsPotentialPlaylist() {
		msg := "It's not a playlist: " + plr.ContentType
		err = plparser.NewPlParserError(msg, echoJson)
		fmt.Println(err)
		os.Exit(1)
	}

	pl := plparser.NewPlaylist(plr)

	pl.Parse()

	if echoJson {
		fmt.Println(pl.StreamsAsJson())
	} else {

		fmt.Printf("Playlist type: %s\n", pl.Type)

		for _, stream := range pl.Streams {

			if stream.Url != "" {
				fmt.Printf("Stream %d url: %s\n", stream.Index, stream.Url)
			} else {
				continue
			}

			if stream.Title != "" {
				fmt.Printf("Stream %d title: %s\n", stream.Index, stream.Title)

			}

			if stream.Author != "" {
				fmt.Printf("Stream %d author: %s\n", stream.Index, stream.Author)
			}

			if stream.Description != "" {
				fmt.Printf("Stream %d description: %s\n", stream.Index, stream.Description)
			}

			if stream.Logo != "" {
				fmt.Printf("Stream %d logo: %s\n", stream.Index, stream.Logo)
			}

			if stream.Copyright != "" {
				fmt.Printf("Stream %d copyright: %s\n", stream.Index, stream.Copyright)
			}

			if stream.MoreInfo != "" {
				fmt.Printf("Stream %d info: %s\n", stream.Index, stream.MoreInfo)
			}
		}
	}
}
