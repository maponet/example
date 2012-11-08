package interno

import "fmt"

type Storer interface {
    ToStoreFormat() String
}

func Store(object Storer) {
    fmt.Printf("\t\tinterno: execute interno.Store on object: %v\n", object)
    actualDbSaving(object.ToStoreFormat())
}

func actualDbSaving(objectInStoreFormat String) {
	fmt.Printf("\t\tinterno: execute interno.actualDbSaving with string: %s\n", objectInStoreFormat)
