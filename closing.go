package main

import "fmt"

/**
	message emitter function factory
*/
func printer(msgs []string) func() string {
	msgIndex := 0
	return func() string {
		if msgIndex >= len(msgs) {
			msgIndex = 0
		}
		msg := msgs[msgIndex]
		msgIndex++
		return msg
	}
}

func main() {
	thing := printer([]string{"1. Hello World", "2. WTF", "3. This is Go", "4. This is fun", "5. How?"})

	for i := 0; i < 10; i++ {
		fmt.Println(thing())
	}
}