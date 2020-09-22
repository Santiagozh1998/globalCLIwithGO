package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Santiagozh1998/create-project/javascript"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	if len(os.Args) < 3 {
		fmt.Println("Something")
		os.Exit(2)
	}

	language := strings.ToLower(os.Args[1])

	if language == "javascript" || language == "js" {

		javascript.FrontendOrBackendProject(dir, os.Args)
	}
}
