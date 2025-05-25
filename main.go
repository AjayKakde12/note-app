package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"note-app.com/root/note"
	"note-app.com/root/todo"
)

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getTodoData() string {
	return getUserInput("Todo text: ")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func main() {
	// Notes handeling
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.ShowNote()

	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed.", err)
		return
	}

	fmt.Println("Saving the note succeeded")

	// Todo Handeling
	content = getTodoData()

	userTodo, err := todo.New(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userTodo.ShowTodo()

	err = userTodo.Save()
	if err != nil {
		fmt.Println("Saving the todo failed.", err)
		return
	}

	fmt.Println("Saving the todo succeeded")
}
