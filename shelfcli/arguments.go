package shelfcli

import (
    "os"
    "path"
)

const ACTION_ARTIFACT string = "artifact"
const ACTION_METADATA string = "metadata"
const ACTION_SEARCH string = "search"

type arguments struct {
    Host string
    Token string
    RefName string
    RemotePath string
    RemoteUrl string
    Action string
    LocalPath string
    MetadataKey string
    MetadataValue string
    MetadataImmutable bool
    SearchData []string
    SearchSort []string
    SearchLimit int
    ErrorList []string

    raw_args map[string]interface{}
}

func NewArguments(raw_args map[string]interface{}) (*arguments) {
    a := new(arguments)
    a.raw_args = raw_args
    return a
}

func (this *arguments) Process() {
    this.Host = this.getValue("--host", "SHELF_HOST")
    this.RemotePath = this.raw_args["<remotePath>"].(string)
    this.RefName = this.raw_args["<refName>"].(string)
    this.RemoteUrl = path.Join(this.Host, this.RefName, "artifact", this.RemotePath)
    this.Token = this.getValue("--token", "SHELF_AUTH_TOKEN")
    this.SearchLimit = -1

    if this.getAnyArgValueDefault([]string{"a", "artifact"}, false).(bool) {
        this.Action = ACTION_ARTIFACT
        this.processArtifact()
    } else if this.getAnyArgValueDefault([]string{"m", "meta"}, false).(bool) {
        this.Action = ACTION_METADATA
        this.processMetadata()
    } else if this.getAnyArgValueDefault([]string{"s", "search"}, false).(bool) {
        this.Action = ACTION_SEARCH
        this.processSearch()
    }
}

func (this *arguments) processArtifact() {
    if create := this.getAnyArgValueDefault([]string{"-c", "--create"}, nil); create != nil {
        this.LocalPath = create.(string)
    }
}

func (this *arguments) processMetadata() {
    if key := this.getAnyArgValueDefault([]string{"-k", "--key"}, nil); key != nil {
        this.MetadataKey = key.(string)
    }

    if value := this.getAnyArgValueDefault([]string{"-v", "--value"}, nil); value != nil {
        this.MetadataValue = value.(string)
    }

    this.MetadataImmutable = this.getAnyArgValueDefault([]string{"--immutable"}, false).(bool)
}

func (this *arguments) processSearch() {
    if limit := this.getAnyArgValueDefault([]string{"-l", "--limit"}, nil); limit != nil {
        var ok bool
        this.SearchLimit, ok = limit.(int)
        if ! ok {
            this.ErrorList = append(this.ErrorList, "Limit could not be converted to an integer")
        }
    }

    compare := func(val interface{}) bool {
        stringList := val.([]string)
        if len(stringList) <= 0 {
            return false
        }

        return true
    }

    searchData := this.getAnyArgValue([]string{"-d", "--data"}, compare)
    if nil != searchData {
        this.SearchData = searchData.([]string)
    }

    searchSort := this.getAnyArgValue([]string{"-s", "--sort"}, compare)

    if nil != searchSort {
        this.SearchSort = searchSort.([]string)
    }
}

func (this *arguments) getValue(argName, envName string) (string) {
    var value string = ""
    if val, ok := this.raw_args[argName]; ok && val != nil {
        value = val.(string)
    } else if val := os.Getenv(envName); val != "" {
        value = val
    }

    return value
}

// Attempts to find any argument whose value is not already the defaultValue
// for that argument.  If all checks fail, the defaultValue is returned.
func (this *arguments) getAnyArgValueDefault(argNameList []string, defaultValue interface{}) (interface{}) {
    compare := func(val interface{}) bool { return val != defaultValue }
    result := this.getAnyArgValue(argNameList, compare)
    if nil == result {
        result = defaultValue
    }

    return result
}

type comparator func(interface{}) (bool)

func (this *arguments) getAnyArgValue(argNameList []string, c comparator) (interface{}) {
    for _, argName := range argNameList {
        if val, ok := this.raw_args[argName]; ok && c(val) {
            return val
        }
    }

    return nil
}
