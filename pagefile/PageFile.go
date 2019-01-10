package pagefile

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type Any interface{}
type Page struct {
	Title string
	Body  []byte
}

type Pages []*Page

//factory method
func NewPage(title string, body []byte) *Page {
	//TODO check parameter
	return &Page{title, body}
}

// save this page's body into file
// Title is the filename
func (pg *Page) Save() {
	outFile, outError := os.OpenFile(pg.Title, os.O_WRONLY|os.O_CREATE, 0666)
	if outError != nil {
		fmt.Printf("An error occurred with file openning or creation (%s)\n", pg.Title)
		return
	}
	defer outFile.Close()

	outWriter := bufio.NewWriter(outFile)
	n, ok := outWriter.Write(pg.Body)
	if ok != nil {
		fmt.Printf("write %d bytes to file but error occurred", n)
		return
	}
	err := outWriter.Flush()
	if err != nil {
		fmt.Println("flush the conntent to disk error")
		return
	}
}

// load the file  "title" into page
func (pg *Page) Load(title string) {
	tempBytes, inError := ioutil.ReadFile(title)
	if inError != nil {
		fmt.Printf("An error occurred with file openning (%s)\n", title)
		tempBytes = nil
	}

	pg.Body = tempBytes //two copy ?
	pg.Title = title
}
