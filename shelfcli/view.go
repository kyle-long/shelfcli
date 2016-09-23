package shelfcli

import (
    "log"
)

type View interface {
    Render()
    AddError(message string)
}

type view struct {
    logger *log.Logger
    errorList []string
    errorLogger *log.Logger
}

func (this *view) AddError(message string) {
    this.errorList = append(this.errorList, message)
}
