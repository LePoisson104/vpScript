package main

import (
	"bufio"
	"fmt"
	"os"
	"vpScript/frontend"
)

func readFile() string {
	file, err := os.Open("test.vp")
	if err != nil {
		fmt.Println("Error open file: ", err)
		return ""
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""

	for scanner.Scan(){
		line := scanner.Text()
		content += line + "; "
	}

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }

	return content
}

func main() {
	content := readFile()
	fmt.Println(frontend.Tokenize(content))
}