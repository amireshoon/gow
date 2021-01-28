package markdown

import (
	"amireshoon/gow/gow"
)

func AddTitle(title string) {

}

func AddTodo(desc string) {

}

func CheckTodo() {

}

// HasTodo returns true if it's already initilized or TODO.md already exists
func HasTodo(path string) bool {
	_, err := gow.GetTodo(path)
	if err != nil {
		return false
	}
	return true
}
