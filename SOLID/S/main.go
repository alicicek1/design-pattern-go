package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	//
}

// separation of concerns

func (j *Journal) Save(fileName string) {
	_ = ioutil.WriteFile(fileName, []byte(j.String()), 0644)
}

func (j *Journal) Load(fileName string) {

}
func (j *Journal) LoadFromWeb(url *url.URL) {

}

//These three methods are only represents for Journal what if I want to save another type of file!

var lineSeparator = "\n"

func SaveToFile(j *Journal, fileName string) {
	_ = ioutil.WriteFile(fileName, []byte(strings.Join(j.entries, lineSeparator)), 0644)
}

// This above method no longer on the Journal anymore. It exists by itself.

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, fileName string) {
	_ = ioutil.WriteFile(fileName, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	SaveToFile(&j, "journal.txt")

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
