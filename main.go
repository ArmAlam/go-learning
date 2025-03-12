package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// if the interface have only one method, then interface name should be that method name + 'er' at the end of the interfac name
type saver interface {
	Save() error
}

func main() {

	title, content := getNoteData()
	todoText := getUserInput("Todo Text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	todo.Display()

	err = saveData(todo)

	if err != nil {
		return
	}

	userNote.Display()

	err = saveData(userNote)

	if err != nil {
		return
	}
}

// saver interface works with both todo and note struct as they both have Save() method with same signature
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving the note: ", err)
		return err
	}

	fmt.Println("Note saved")

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: \n")

	content := getUserInput("Note Content: \n")

	return title, content
}

func getUserInput(promt string) string {
	fmt.Print(promt)

	// bufio is used to take long input (input that contains spaces)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n') // when new line encountered, stop taking input

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n") // remove new line character
	text = strings.TrimSuffix(text, "\r") // for windows

	return text
}
