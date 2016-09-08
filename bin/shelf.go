//usr/bin/env go run $0 $@; exit;
package main

import (
    "fmt"
    "github.com/docopt/docopt-go"
    // "os"
)

func main() {
    doc := `Shelf Cli
        A command line client for interacting with the shelf api
        (https://github.com/kyle-long/pyshelf).

        Usage:
            shelf <refName> [--host=<host>] [--token=<token>] (a | artifact) [--create=<localPath>] <remotePath>
            shelf <refName> [--host=<host>] [--token=<token>] (m | meta) [--name=<name>] [--value=<value>] [--immutable] <remotePath>
            shelf <refName> [--host=<host>] [-token=<token>] (s | search) [--data=<searchData>]... [--limit=<limit>] [--sort=<sort>]... <remotePath>

        Sub Commands:
            a artifact                          Specifies you would like to act on an artifact.  This can include getting or
                                                creating a new artifact.  The output will be whatever that artifact is unless
                                                you specify the "create" option, and then it will only give you a success
                                                message.

            m meta                              Specifies you would like to act on metadata for an artifact.  The output
                                                will be a JSON document which defines the metadata.  If you define a name
                                                it will only show the results for that particular metadata.

            s search                            Specifies you would like to search for an artifact.  The output will be a
                                                newline delimited list of paths to those artifacts.
        Arguments:
            <refName>                           Defines which "space" you would like to act upon.  This corresponds (usually) to a
                                                S3 bucket in amazon.

            <remotePath>                        The path in shelf you would like to act upon.  This always corresponds to an
                                                artifacts path.

        Options:
            -h host, --host host                The host of the API you would like to interact with. For example
                                                https://api.shelf.com/

            -t token, --token token             The authentication token for interacting with shelf.

            -c localPath, --create              localPath When using the artifact sub command this indicates you will be
                                                uploading a new artifact.  The value of the option is the path locally
                                                that you wish to upload.

            --name name                         When using the "meta" sub command, this will make the command act on a single
                                                metadata property instead of all metadata.

            --value value                       When using the "meta" sub command, this will set the value of a particular
                                                metadata property.

            --immutable                         When using the "meta" sub command, this will make a property be immutable.
                                                This only affects updates or creates.

            -d serachData, --date searchData    Specifies a search criterial. For more information see
                                                https://github.com/kyle-long/pyshelf/blob/master/docs/api/search.md

            -l limit, --limit limit             Defines how many links should come back from a search.

            -s sort, --sort sort                Defines how to sort links coming back.  For more information see
                                                https://github.com/kyle-long/pyshelf/blob/master/docs/api/search.md

    `

    arguments, err := docopt.Parse(doc, nil, true, "shelfcli 0.1", false)
    fmt.Println(arguments)
    fmt.Println(err)
}
