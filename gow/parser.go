package gow

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gookit/color"
)

// Parse : This function parse the main storage
func Parse() {
	content := `name:this is name:desc:this is some:is_done:yes:name:this is name2:desc:this is some2:is_done:no`
	res := strings.Split(content, `:`)
	red := color.FgRed.Render
	green := color.FgGreen.Render
	line := ""
	for p, e := range res {
		if p%2 == 0 {
			continue
		}

		if res[p-1] == "is_done" {
			if e == "yes" {
				line += `[` + green(`yes`) + `]` + green(e) + "\n"
			} else {
				line += `[` + red(`no`) + `]` + e + "\n"
			}
		} else {
			line += res[p-1] + e + "\n"
		}

	}

	println(line)
}

// ParseTodo Get todo.md content or create if not exists
// Added a title with content of passed name
func ParseTodo(path string, des string, args ...string) error {
openingFile:
	_, err := GetTodo(path)
	if err != nil {
		arg := string(strings.Join(args[:], " "))
		baseContent := "# " + arg + " \n\n" + des
		err := ioutil.WriteFile(path+"/TODO.md", []byte(baseContent), 0755)
		if err != nil {
			return errors.New("Unable to write TODO.md")
		}
		goto openingFile
	}
	return nil
}

// GetTodo gets TODO.md file contents
func GetTodo(path string) (string, error) {
	content, err := ioutil.ReadFile(path + "/TODO.md")
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// FillTodo will fill TODO.md file
func FillTodo(content string, path string) error {
	err := ioutil.WriteFile(path+"/TODO.md", []byte(content), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
		return errors.New("Unable to write file")
	}
	return nil
}

// FillReadme will fill README.md file
func FillReadme(content string, path string) error {
	err := ioutil.WriteFile(path+"/README.md", []byte(content), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
		return errors.New("Unable to write file")
	}
	return nil
}

func _load() (string, error) {
	p, err := Dir()

	if err != nil {
		panic(err)
	}

openingFile:
	content, err := ioutil.ReadFile(p + "/works.gow")
	if err != nil {
		err := ioutil.WriteFile(p+"/works.gow", []byte(""), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
		goto openingFile
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text, err
}

func getParseKeys() (primary []string, secondary []string) {
	return []string{"works", "todos"}, []string{"work", "todo"}
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
