package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// Creating a interface
type Saver interface {
	Save() error
}

// type Displayer interface {
// 	Display()
// }

// Embedded Interface
type Outputable interface {
	Saver
	Display()
}

// Using the interface
func saveData(data Saver) error {
	if err := data.Save(); err != nil {
		fmt.Println("Saving the todo failed.")
		return err
	}
	fmt.Println("Saving the todo succeded!")
	return nil
}

func outputData(data Outputable) error {
	data.Display()
	if err := saveData(data); err != nil {
		return err
	}
	return nil
}

func main() {

	printSomething(1)
	printSomething(1.4)
	printSomething("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)

}

// "Any Value Allowed" type
// & Working with type switches
// & Extracting type information from values
func printSomething(value interface{}) {

	if intVal, ok := value.(int); ok {
		fmt.Println("Integer:", intVal)
	} else if floatVal, ok := value.(float64); ok {
		fmt.Println("Float 64:", floatVal)
	} else if stringVal, ok := value.(string); ok {
		fmt.Println("String:", stringVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float 64:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
