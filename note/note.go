package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"` //struct tag, used to attach metadata to struct fields, can be used by libraries (e.g., JSON) for processing data
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Note title %v, note content %v", note.Title, note.Content)
}

func (note Note) Save() error {

	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// convert data to JSON
	json, err := json.Marshal(note) // to be marshalbe, struct property's name first character should be capital

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) // 0644 set read and edit permisson for the owner of the file

}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
