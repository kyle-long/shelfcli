package shelfcli

import (
    "log"
    "io"
    "github.com/quantumew/shelflib"
    "github.com/tomnomnom/linkheader"
)

type viewFactory struct {
    errorLogger *log.Logger
    logger *log.Logger
}

func NewViewFactory(logger *log.Logger, errorLogger *log.Logger) (*viewFactory) {
    factory := new(viewFactory)
    factory.logger = logger
    factory.errorLogger = errorLogger
    return factory
}

func (this *viewFactory) assignCommonProperties(v *view) {
    v.logger = this.logger
    v.errorLogger = this.errorLogger
}

func (this *viewFactory) NewArtifactView(response io.ReadCloser) (*ArtifactView) {
    v := new(ArtifactView)
    this.assignCommonProperties(&v.view)
    v.response = response
    return v
}

func (this *viewFactory) NewArtifactListView(link_list linkheader.Links) (*ArtifactListView) {
    v := new(ArtifactListView)
    this.assignCommonProperties(&v.view)
    v.link_list = link_list
    return v
}

func (this *viewFactory) NewErrorView(err shelflib.ShelfError) (*ErrorView) {
    v := new(ErrorView)
    this.assignCommonProperties(&v.view)
    v.AddError(err.Message)
    return v
}
