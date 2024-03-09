package main

import (
	"fmt"

	"example.com/struct/note"
	"example.com/struct/todo"
	"example.com/struct/userinput"
)

type Saver interface {
	Save() (interface{}, error) // this are exact same type as returned by the functions
}
type Displayer interface {
	Display() error
}

type SaveAndDisplay interface {
	Saver
	Displayer
}

func main() {
	// for saving notes
	title, content := userinput.GetUserInput()
	newNote := note.New(title, content)
	// if data, err := newNote.Save(); err != nil {
	// 	fmt.Println("error", err)
	// } else {
	// 	fmt.Println(data)
	// }
	// if err := newNote.Display(); err != nil {
	// 	fmt.Println("error", err)
	// } else {
	// 	fmt.Println("displaying data")
	// }

	// for saving todo

	text := userinput.GetTodoInput()
	newTodo := todo.New(text)
	// if data, err := newTodo.Save(); err != nil {
	// 	fmt.Println("error", err)
	// } else {
	// 	fmt.Println(data)
	// }

	// if err := newTodo.Display(); err != nil {
	// 	fmt.Println("error", err)
	// } else {
	// 	fmt.Println("displaying data")
	// }

	// now we learn the concept of interface and why it necessary

	// in above two struct we can see thar we need to call save method and dispaly methd individually for both of them
	// now we can make on function which can call these method

	// the above two calls can be made even more simpler by using interface

	// MySave(newNote)
	// MySave(newTodo)
	// MyDisplay(newNote)
	// MyDisplay(newTodo)

	// nowe we can more refinr it by calling one fucntion which will save and display both instaed of above four calls

	Both(newNote)
	Both(newTodo)

}

// this says the data struct will definately have all the methods mentioned in the Saver interface
func MySave(data Saver) {
	if data, err := data.Save(); err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(data)
	}
}
func MyDisplay(data Displayer) {
	if err := data.Display(); err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("displaying data")
	}
}

func Both(data SaveAndDisplay) {
	if data, err := data.Save(); err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(data)
	}
	if err := data.Display(); err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("displaying data")
	}
}
