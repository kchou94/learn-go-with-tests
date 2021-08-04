package main

import "fmt"

const chinese = "Chinese"
const french = "French"
const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "你好, "
const frencheHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case chinese:
		prefix = chineseHelloPrefix
	case french:
		prefix = frencheHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Kyle", ""))
}
