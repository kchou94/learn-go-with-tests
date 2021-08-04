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

	if language == chinese {
		return chineseHelloPrefix + name
	}
	if language == french {
		return frencheHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Kyle", ""))
}
