package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"` //struct tag, used to attach metadata to struct fields, can be used by libraries (e.g., JSON) for processing data
}

func (todo Todo) Display() {

	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {

	fileName := "todo.json"

	// convert data to JSON
	json, err := json.Marshal(todo) // to be marshalbe, struct property's name first character should be capital

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) // 0644 set read and edit permisson for the owner of the file

}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("Invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}
