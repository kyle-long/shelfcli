package shelfcli

import (
    "log"
)
type manager struct {
    Logger log.Logger
}

func NewManager(logger log.Logger) (*manager) {
    m := new(manager)
    m.Logger = logger
    return m
}

func (this *manager) Run(args *arguments) {
    switch args.Action {
        case ACTION_ARTIFACT:
            this.runArtifact(args)

        case ACTION_METADATA:
            this.runMetadata(args)

        case ACTION_SEARCH:
            this.runSearch(args)
    }
}

func (this *manager) runArtifact(args *arguments) {
    if args.LocalPath == "" {
        // this.whatever.ListArtifact(args.RemotePath)
        // if there is only one link
        // this.whatever.GetArtifact(args.RemotePath)
        // output to stdout
    } else {
        // Attempt to open args.LocalPath and get a stream
        // this.whatever.CreateArtifact(args.RemotePath, stream)
    }
}

func (this *manager) runMetadata(args *arguments) {
    if args.MetadataValue == "" {
        // item := map[string]interface{}{
        //     "key": args.MetadataKey,
        //     "value": args.MetadataValue,
        //     "immutable": args.MetadataImmutable,
        // }
        // this.whatever.UpdateMetadataItem(args.RemotePath, item)
    } else {
        if args.MetadataKey == "" {
            // this.whatever.GetMetadata(args.RemotePath)
        } else {
            // this.whatever.GetMetadataItem(args.RemotePath, args.MetadataKey)
        }
    }
}

func (this *manager) runSearch(args *arguments) {
    // this.whatever.Search(args.RemotePath, args.SearchData, args.SearchSort, args.Limit)
}
