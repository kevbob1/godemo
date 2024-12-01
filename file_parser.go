package main

import (
	"log/slog"
	"fmt"
)


func ParseFile(filePath string) (string, error) {

}

func main() {
	filePath := os.Args[1]
	content, err := ParseFile(filePath)
	if err != nil {
		slog.Error("Error parsing file", "error", err)
		os.Exit(1)
	}
}
