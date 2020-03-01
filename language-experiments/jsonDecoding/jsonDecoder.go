package jsonDecoding

import (
	"encoding/json"
	"fmt"
	"strings"
)

const blob = `[
	{"Title":"cdugga", "Skill": "full stack developer"},
	{"Title": "gholly", "Skill": "full stack teacher"}
]`


type Item struct {
	Title string
	Skill string
}

func TestJson(){
	var items []*Item

	json.NewDecoder(strings.NewReader(blob)).Decode(&items)
	for _, item := range items {
		fmt.Printf("Title: %v Skill: %v\n", item.Title, item.Skill)
	}

}