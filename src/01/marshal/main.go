package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Foo struct {
	Bar string `xml:"id,attr"`
	Baz string `xml:"div>p"`
}

func main() {
	f := Foo{"Joe Junior", "Hello Shabado"}
	jb, _ := json.Marshal(f)
	fmt.Println(string(jb))
	json.Unmarshal(jb, &f)
	xb, _ := xml.Marshal(f)
	fmt.Println(string(xb))
}
