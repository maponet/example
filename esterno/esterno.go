package esterno

import "example/medio"
import "fmt"

func Start() {

    fmt.Printf("\nesterno: call NewExtender\n")
    e := medio.NewExtender()
    
    fmt.Printf("\nesterno: call AddAttribute\n")
    e.AddAttribute("01")
    
    fmt.Printf("\nesterno: call SetAttribute\n")
    e.SetAttribute("01", "pippo")
    
    fmt.Printf("\nesterno: call Save\n")
    e.Save()
}
