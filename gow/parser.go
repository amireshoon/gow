package gow

import (
	"fmt"
	"io/ioutil"
)

// Parse : This function parse the main storage
func Parse() {

	// c, err := _load()

	// if err != nil {
	// 	panic(err)
	// }
	// content := []byte(c)
	// a, b, f, k := jsonparser.Get(content, "works")
	// gh := string(a[:len(a)])

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
