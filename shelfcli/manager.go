package shelfcli

import (
    "log"
    "github.com/quantumew/shelflib"
)

type manager struct {
    Logger *log.Logger
    args *arguments
    lib *shelflib.ShelfLib
    viewFactory *viewFactory
}

func NewManager(logger *log.Logger, errorLogger *log.Logger, args *arguments) (*manager) {
    m := new(manager)
    m.lib = shelflib.New(args.Token, errorLogger)
    m.Logger = logger
    m.args = args
    m.viewFactory = NewViewFactory(logger, errorLogger)
    return m
}

func (this *manager) Run() (View){
    var view View
    switch this.args.Action {
        case ACTION_ARTIFACT:
            view = this.runArtifact()

        case ACTION_METADATA:
            this.runMetadata()

        case ACTION_SEARCH:
            this.runSearch()
    }

    return view
}

func (this *manager) runArtifact() (View) {
    var view View
    if this.args.LocalPath == "" {
        artifactLinkList, err := this.lib.ListArtifact(this.args.RemoteUrl)
        if err != nil {
            view = this.handleError(err)
        } else {
            artifactLinkList := artifactLinkList.FilterByRel("item")
            if len(artifactLinkList) == 0 {
                // TODO: Actually handle this properly.
                err := shelflib.NewShelfError("Not Found", "NOT_FOUND")
                view = this.viewFactory.NewErrorView(*err)
            } else if len(artifactLinkList) == 1 {
                response, err := this.lib.GetArtifact(this.args.RemoteUrl)

                if err != nil {
                    view = this.handleError(err)
                } else {
                    view = this.viewFactory.NewArtifactView(response)
                }
            } else {
                view = this.viewFactory.NewArtifactListView(artifactLinkList)
            }
        }
    } else {
        // Attempt to open args.LocalPath and get a stream
        // this.whatever.CreateArtifact(args.RemotePath, stream)
    }

    return view
}

func (this *manager) runMetadata() {
    if this.args.MetadataValue == "" {
        // this.whatever.UpdateMetadataItem(args.RemotePath, item)
    } else {
        if this.args.MetadataKey == "" {
            // this.whatever.GetMetadata(args.RemotePath)
        } else {
            // this.whatever.GetMetadataItem(args.RemotePath, args.MetadataKey)
        }
    }
}

func (this *manager) runSearch() {
    // this.whatever.Search(args.RemotePath, args.SearchData, args.SearchSort, args.Limit)
}

func (this *manager) handleError(err error) (*ErrorView) {
    // view := this.viewFactory.NewErrorView(err)
    // return view
    fakeErr := shelflib.NewShelfError(err.Error(), "BAD")
    return this.viewFactory.NewErrorView(*fakeErr)
}
