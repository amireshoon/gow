package markdown

import (
	"amireshoon/gow/gow"
	"fmt"
)

func AddTitle(title string) {

}

func AddTodo(desc string, path string) {
	c, err := gow.GetTodo(path)

	if err != nil {
		fmt.Println("Could not read TODO.md file")
	}

	c += `
- [ ] ` + desc + ``
	err = gow.FillTodo(c, path)
	if err != nil {
		fmt.Println("Could not write to file")
	}
}

func CheckTodo(name string, path string) {
	c, err := gow.GetTodo(path)

	if err != nil {
		fmt.Println("Could not read TODO.md file")
	}

	fmt.Println(c)
}

// HasTodo returns true if it's already initilized or TODO.md already exists
func HasTodo(path string) bool {
	_, err := gow.GetTodo(path)
	if err != nil {
		return false
	}
	return true
}
