package gow

import (
	"bytes"
	"errors"
	"fmt"
	"io"
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
	UpdateReadmeWithTodo(path, content) // Updating Readme with new todo
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

func UpdateReadmeWithTodo(path string, newTodo string) {
	c, _ := GetTodo(path)
	r, _ := LoadReadme(path)
	scanner := NewScanner(strings.NewReader(c), len(c))
	lCount := 0
	for {
		_, _, err := scanner.Line()
		if err != nil {
			break
		}
		lCount++
	}

	scanner = NewScanner(strings.NewReader(r), len(r))
	reReadme := ""
	rCounter := 0
	for {
		line, _, err := scanner.Line()
		if err != nil {
			break
		}

		if rCounter > lCount {
			reReadme = line + "\n" + reReadme
		} else {
			reReadme += ""
		}

		rCounter++
	}
	// os.Exit((1))
	reReadme += "\n" + newTodo
	FillReadme(reReadme, path)
}

func LoadReadme(path string) (string, error) {

openingFile:
	content, err := ioutil.ReadFile(path + "/README.md")
	if err != nil {
		err := ioutil.WriteFile(path+"README.md", []byte(""), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}
		goto openingFile
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text, err
}

type Scanner struct {
	r   io.ReaderAt
	pos int
	err error
	buf []byte
}

func NewScanner(r io.ReaderAt, pos int) *Scanner {
	return &Scanner{r: r, pos: pos}
}

func (s *Scanner) readMore() {
	if s.pos == 0 {
		s.err = io.EOF
		return
	}
	size := 1024
	if size > s.pos {
		size = s.pos
	}
	s.pos -= size
	buf2 := make([]byte, size, size+len(s.buf))

	// ReadAt attempts to read full buff!
	_, s.err = s.r.ReadAt(buf2, int64(s.pos))
	if s.err == nil {
		s.buf = append(buf2, s.buf...)
	}
}

func (s *Scanner) Line() (line string, start int, err error) {
	if s.err != nil {
		return "", 0, s.err
	}
	for {
		lineStart := bytes.LastIndexByte(s.buf, '\n')
		if lineStart >= 0 {
			// We have a complete line:
			var line string
			line, s.buf = string(dropCR(s.buf[lineStart+1:])), s.buf[:lineStart]
			return line, s.pos + lineStart + 1, nil
		}
		// Need more data:
		s.readMore()
		if s.err != nil {
			if s.err == io.EOF {
				if len(s.buf) > 0 {
					return string(dropCR(s.buf)), 0, nil
				}
			}
			return "", 0, s.err
		}
	}
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
