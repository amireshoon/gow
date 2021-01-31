package markdown

import (
	"amireshoon/gow/gow"
	"bufio"
	"fmt"
	"strings"
)

// AddTitle add or change existing title
func AddTitle(title string) {

}

// AddTodo adds new todo to TODO.md file
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

// CheckTodo checks existing todo
func CheckTodo(index int, path string) {
	c, err := gow.GetTodo(path)

	if err != nil {
		fmt.Println("Could not read TODO.md file")
	}

	// fmt.Println(c)
	reGeneratedTodo := ``
	scanner := bufio.NewScanner(strings.NewReader(c))
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, `- [ ] `) {
			if counter == index {
				reGeneratedTodo += strings.Replace(line, `- [ ] `, `- [x] `, -1) + "\n"
			} else {
				reGeneratedTodo += line + "\n"
			}
			counter++
		} else {
			reGeneratedTodo += line + "\n"
		}
	}
	err = gow.FillTodo(reGeneratedTodo, path)
	if err != nil {
		fmt.Println("Could not write to file")
	}
}

// HasTodo returns true if it's already initilized or TODO.md already exists
func HasTodo(path string) bool {
	_, err := gow.GetTodo(path)
	if err != nil {
		return false
	}
	return true
}
