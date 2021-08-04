package main

import (
	"fmt"
)

const chinese = "Chinese"
const french = "French"
const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "你好, "
const frencheHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case chinese:
		prefix = chineseHelloPrefix
	case french:
		prefix = frencheHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Kyle", ""))
}
