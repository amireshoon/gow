package gow

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

type work struct {
	id     string
	name   string
	desc   string
	isDone bool
}

var works []work

// Parse : This function parse the main storage
func Parse() {

	c, err := _load()

	if err != nil {
		panic(err)
	}
	prim, sec := getParseKeys()
	if len(c) >= 0 {
		fmt.Println(c)
		scanner := bufio.NewScanner(strings.NewReader(c))
		level := 0
		on := ""
		for scanner.Scan() {
			fmt.Println("Im in level ", level, " and on ", on)

			for _, v := range prim {
				if scanner.Text() == v+":" {
				detectLevel:
					if on != "" && on != v+":" {
						level--
						on = v + ":"
						goto detectLevel
					}
					level++
					on = v + ":"
				}
			}

			fmt.Println(scanner.Text())
			works = append(works, newWork("a", "a", "a", false))
		}
		fmt.Println(works, sec, level)
	}

}

func newWork(id string, name string, desc string, isDone bool) work {
	w := work{
		id:     id,
		name:   name,
		desc:   desc,
		isDone: isDone,
	}
	return w
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
