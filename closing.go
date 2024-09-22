package main

import (
	"fmt"
	"time"
)

type SomeType struct {
	msg string
}



/**
	message emitter function factory
*/
func printer( msgs *[]SomeType) func() string {
	msgIndex := 0
	return func() string {
		if msgIndex >= len(*msgs) {
			msgIndex = 0
		}
		msgst := (*msgs)[msgIndex]
		msgIndex++
		return msgst.msg
	}
}

func main() {
	msgs := []SomeType{{"1. Hello World"}, {"2. WTF"}, {"3. This is Go"}, {"4. This is fun"}, {"5. How?"}}
	thing := printer(&msgs)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(thing())
		}
	}()

	time.Sleep(time.Second * 10)
}