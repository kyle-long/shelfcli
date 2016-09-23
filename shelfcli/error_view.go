package shelfcli

type ErrorView struct {
    view
}

func (this *ErrorView) Render() {
    for _, err := range this.errorList {
        this.errorLogger.Print(err)
    }
}
