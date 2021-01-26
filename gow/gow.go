package gow

import (
	"amireshoon/gow/cmd"
	"fmt"
)

// Run gow app
func Run() bool {
	o := false
	Block{
		Try: func() {
			cmd.Execute()
			o = true
		},
		Catch: func(e Exception) {
			fmt.Println("Gow cannot start working")
		},
	}.Do()
	return o
}
