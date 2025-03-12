package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving the note: ", err)
		return
	}

	fmt.Println("Note saved")

}

func getNoteData() (string, string) {
	title := getUserIntpu("Note title: \n")

	content := getUserIntpu("Note Content: \n")

	return title, content
}

func getUserIntpu(promt string) string {
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
