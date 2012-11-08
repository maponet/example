package interno

import "fmt"

type database struct {
}

func (db database) Save() {
    fmt.Printf("\t\tinterno: execute interno.Save\n")
}

type Saver interface {
    Save()
}

func New() Saver {
    fmt.Printf("\t\tinterno: execute interno.New\n")
    var s database
    return s
}
