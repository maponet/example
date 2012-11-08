package medio

import "example/interno"
import "fmt"

type Extender interface {
    AddAttribute(name string)
    SetAttribute(name, value string)
    SetPath(p string)
    Save()
    interno.Storer
}

type object struct {
    path string
    attributes map[string]string
}

func (o object) AddAttribute(name string) {
    o.attributes[name] = ""
    fmt.Printf("\tmedio: execute AddAttribute\n")
}

func (o object) SetAttribute(name, value string) {
    o.attributes[name] = value
    fmt.Printf("\tmedio: execute SetAttribute\n")
}

func (o object) SetPath(path string) {
    o.path = path
    fmt.Printf("\tmedio: execute SetPath\n")
}

func (o object) Save() {
    fmt.Printf("\tmedio: execute Save\n")
    
    fmt.Printf("\tmedio: call interno.Store\n")
    interno.Store(o)
}

func (o object) ToStoreFormat() string {
    fmt.Printf("\tmedio: execute ToStoreFormat\n")
    return "Hello!"
}

func NewExtender() Extender {
    fmt.Printf("\tmedio: execute NewExtender\n")
    var o object
    o.attributes = make(map[string]string, 0)
    return o
}
