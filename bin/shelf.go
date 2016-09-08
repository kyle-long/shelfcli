//usr/bin/env go run $0 $@; exit;
package main

import (
    "fmt"
    "github.com/docopt/docopt-go"
    "os"
    "bufio"
    "strings"
    "unicode/utf8"
)

func main() {
    doc := `Shelf Cli
        A command line client for interacting with the shelf api
        (https://github.com/kyle-long/pyshelf).

        Usage:
            shelfcli <refName> (a | artifact) [-c=<localPath> | --create=<localPath>] [<remotePath>]
            shelfcli <refName> (m | meta) [-v=(PUT | POST) | --verb=(PUT | POST)] [--name=<name>] [--value=<value>] [--immutable] [<remotePath>]
            shelfcli <refName> (s | search) [-d=<searchData> | --data=<searchData>]... [-l=<limit> | --limit=<limit>] [-s=<sort> |--sort=<sort>]... [<remotePath>]

        Sub Commands:
            a artifact                      Specifies you would like to act on an artifact.  This can include getting or
                                            creating a new artifact.  The output will be whatever that artifact is unless
                                            you specify the "create" option, and then it will only give you a success
                                            message.

            m meta                          Specifies you would like to act on metadata for an artifact.  The output
                                            will be a JSON document which defines the metadata.  If you define a name
                                            it will only show the results for that particular metadata.

            s search                        Specifies you would like to search for an artifact.  The output will be a
                                            newline delimited list of paths to those artifacts.
        Arguments:
            <refName>                       Defines which "space" you would like to act upon.  This corresponds (usually) to a
                                            S3 bucket in amazon.

            <remotePath>                    The path in shelf you would like to act upon.  This always corresponds to an
                                            artifacts path.

        Options:
            -c localPath --create           localPath When using the artifact sub command this indicates you will be
                                            uploading a new artifact.  The value of the option is the path locally
                                            that you wish to upload.

            -v verb --verb verb             When using the "meta" sub command, this will specify how you would like to
                                            upload metadata.  If "POST" is given it will error if that particular metadata
                                            already exists.  If "PUT" is given it will simply update the metadata if it
                                            exists and create it if it does not.  Note: You will not be able to update
                                            metadata if it is set as immutable.

            --name name                     When using the "meta" sub command, this will make the command act on a single
                                            metadata property instead of all metadata.

            --value value                   When using the "meta" sub command, this will set the value of a particular
                                            metadata property.

            --immutable                     When using the "meta" sub command, this will make a property be immutable.
                                            This only affects updates or creates.

            -d serachData --date searchData Specifies a search criterial. For more information see
                                            https://github.com/kyle-long/pyshelf/blob/master/docs/api/search.md

            -l limit --limit limit          Defines how many links should come back from a search.

            -s sort --sort sort             Defines how to sort links coming back.  For more information see
                                            https://github.com/kyle-long/pyshelf/blob/master/docs/api/search.md

    `

    arguments, err := docopt.Parse(doc, nil, true, "shelfcli 0.1", false)
    //fmt.Println(arguments)
    fmt.Println(err)
    value, _ := find_remote_path(arguments)
    fmt.Println(value)
}

func find_remote_path(arguments map[string]interface{})(string, bool) {
    var success bool
    var value string
    var ok bool

    success = false

    if value, ok = arguments["<remotePath>"].(string); ! ok {
        stdinStat, err := os.Stdin.Stat()

        if err != nil {
            fmt.Println("Error while using Stat for stdin", err)
        } else if stdinStat.Mode() & os.ModeNamedPipe != 0 {
            reader := bufio.NewReader(os.Stdin)

            value, err = reader.ReadString('\n')
            if utf8.RuneCountInString(value) > 0 {
                value = strings.TrimSpace(value)
            } else {
                value = ""
            }


        } else {
            value = ""
        }
    }

    if value == "" {
        success = false
    }

    return value, success
}
