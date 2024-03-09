package note

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `josn:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Note) Save() (interface{}, error) {
	// jsoniffy the data
	marshledData, err := json.Marshal(n)
	if err != nil {
		return Note{}, err
	}
	os.WriteFile(n.Title+".json", marshledData, 0644)
	return *n, nil
}
func (n *Note) Display() error {
	jsonData, err := json.Marshal(*n)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	return nil
}

// 1st step is to write  a constructor function this will create an object
// and to this oject and over thi object we can call methods
// ishme refrence ka darkar nhi hai quinki ye toh objec create karege
// as javascriot mein hota hai

// use of inteface
func New(title, content string) *Note {

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

}
