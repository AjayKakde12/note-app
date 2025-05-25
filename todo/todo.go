package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const FILES_DIR string = "files-todo"

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) ShowTodo() {
	fmt.Println(todo)
}

func (todo Todo) Save() error {
	fileName := fmt.Sprintf("%v/todo.json", FILES_DIR)

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	err = os.MkdirAll(FILES_DIR, 0644)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}
