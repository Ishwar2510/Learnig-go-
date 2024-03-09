package userinput

import "fmt"

func GetUserInput() (title string, content string) {
	fmt.Println("Enter title ")
	fmt.Scan(&title)
	fmt.Println("Enter content")
	fmt.Scan(&content)
	return title, content
}
func GetTodoInput() string {
	fmt.Println("Enter todo text ")
	var text string
	fmt.Scan(&text)
	return text
}
