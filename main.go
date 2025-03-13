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

type dispalyer interface {
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

// embeding existing interfaces to create new interface
type outputtable interface {
	saver
	Display()
}

func main() {

	printAnything("Hello")
	printAnything(1)
	printAnything(1.0)
	add(1, 2)

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

	err = outputData(todo)

	// no error handling here, as program execution stops here
	outputData(userNote)

}

// this function can take any type of data
func printAnything(value interface{}) {
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer ", value)
	// case string:
	// 	fmt.Println("String ", value)
	// default:
	// 	fmt.Println("Unknown ", value)
	// }
	// alternatvie way to do the same thing

	intVal, ok := value.(int) // check if integer

	if ok {
		fmt.Println("Integer ", intVal)
		return
	}

	floatVal, ok := value.(float64) // check if integer

	if ok {
		fmt.Println("Float ", floatVal)
		return
	}

	stringVal, ok := value.(string) // check if integer

	if ok {
		fmt.Println("String ", stringVal)
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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

// genecis, can receive int, float64, string
func add[T int | float64 | string](a, b T) T {
	return a + b
}
