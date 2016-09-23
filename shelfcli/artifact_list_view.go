package shelfcli

type ArtifactListView struct {
    view
    link_list []string
}

func (this *ArtifactListView) Render() {
    for link := range this.link_list {
        this.logger.Print(link)
    }
}
