package main

import (
	"fmt"
	"log/slog"
	"os"
)

func ParseFile(filePath string) (string, error) {
	return "", nil
}

func main() {
	filePath := os.Args[1]
	fmt.Println(filePath)
	content, err := ParseFile(filePath)
	if err != nil {
		slog.Error("Error parsing file", "error", err)
		os.Exit(1)
	}
	fmt.Println(content)
}
