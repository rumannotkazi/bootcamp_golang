package main

import (
	"fmt"
	"os"
)

func main() {
	var input_dir string
	fmt.Scanln(&input_dir)
	files, err := os.ReadDir(input_dir)
	if err != nil {
		panic("Wrong Directory")
	}
	for _, file := range files {
		fmt.Println(file.Type())
	}

}
