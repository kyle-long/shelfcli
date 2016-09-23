package shelfcli

import (
    "io"
    "os"
)

type ArtifactView struct {
    view
    response io.ReadCloser
}

func (this *ArtifactView) Render() {
    io.Copy(os.Stdout, this.response)
}
