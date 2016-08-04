//usr/bin/env go run $0 $@; exit;
package main

import (
    "fmt"
    "github.com/docopt/docopt-go"
)

func main() {
    doc := `Shelf Cli
        A command line client for interacting with the shelf api
        (https://github.com/kyle-long/pyshelf).

        Usage:
            shelfcli <refName> (a | artifact) [-c=<localPath> | --create=<localPath>] [<remotePath>]
            shelfcli <refName> (m | meta) [-v=<verb> | --verb=<verb>] [--name=<name>] [--value=<value>] [--immutable] [<remotePath>]
            shelfcli <refName> (s | search) [-d=<searchData> | --data=<searchData>]... [-l=<limit> | --limit=<limit>] [-s=<sort> |--sort=<sort>]... [<remotePath>]
    `

    arguments, err := docopt.Parse(doc, nil, true, "shelfcli 0.1", false)
    fmt.Println(arguments)
    fmt.Println(err)
}
