package main

import "fmt"

const chinese = "Chinese"
const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "你好, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == chinese {
		return chineseHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Kyle", ""))
}
