package shelfcli

import (
    "github.com/tomnomnom/linkheader"
)

type ArtifactListView struct {
    view
    link_list linkheader.Links
}

func (this *ArtifactListView) Render() {
    if len(this.link_list) < 1 {
        this.errorLogger.Print("Not Found")
    } else {
        for _, link := range this.link_list {
            this.logger.Print(link.URL)
        }
    }
}
