//usr/bin/env go run $0 $@; exit;
package main

import (
    "github.com/docopt/docopt-go"
)

func main() {
    doc := `Shelf Cli
        A command line client for interacting with the shelf api
        (https://github.com/kyle-long/pyshelf).

        Usage:
            shelfcli (a | artifact) [-p=<localPath>|--put=<localPath>] [remotePath]
            shelfcli (m | meta) [p|put] [--name=<name>] [--value=<value>] [--immutable] [path]
            shelfcli (s | search) [-d=<searchData>|--data=<searchData>]... [-l=<limit>|--limit=<limit>] [-s=<sort>|--sort=<sort>]... [path]
    `

    arguments, _ := docopt.Parse(doc, nil, true, "shelfcli 0.1", false)
}
