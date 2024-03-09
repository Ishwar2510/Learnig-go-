package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Todo      string    `json:"Todo"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Todo) Save() (interface{}, error) {
	// jsoniffy the data
	marshledData, err := json.Marshal(n)
	if err != nil {
		return Todo{}, err
	}
	os.WriteFile(n.Todo+".json", marshledData, 0644)
	return *n, nil
}
func (n *Todo) Display() error {
	jsonData, err := json.Marshal(*n)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	return nil
}
func New(todo string) *Todo {

	return &Todo{
		Todo:      todo,
		CreatedAt: time.Now(),
	}

}
