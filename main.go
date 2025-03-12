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

}

func getNoteData() (string, string) {
	title := getUserIntpu("Note title")

	content := getUserIntpu("Note Contetn")

	return title, content
}

func getUserIntpu(promt string) string {
	fmt.Print(promt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
