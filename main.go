package main

import (
	"fmt"
	"github.com/SimpleWebsitebyGo/pagefile"
)

func main() {
	pg := pagefile.NewPage("abc.html", []byte("<html>hello, world !</html>"))
	pg.Save()
	var pg2 *pagefile.Page = new(pagefile.Page)
	pg2.Load("abc.html")
	fmt.Printf("%s\n", pg2.Body)
}
