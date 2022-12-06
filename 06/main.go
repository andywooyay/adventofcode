package main

import (
	"io/ioutil"
	"log"
)

func main() {
	b, _ := ioutil.ReadFile("06/test.txt")
	log.Println("Start of packet:", findFirst(string(b), 4))
	log.Println("Start of message:", findFirst(string(b), 14))
}

func findFirst(b string, size int) int {
	for i := size; i < len(b); i += 1 {
		if allUniqueChars(b[i-size : i]) {
			return i
		}
	}

	return 0
}

func allUniqueChars(s string) bool {
	uniq := make(map[rune]struct{})
	for _, v := range s {
		if _, ok := uniq[v]; ok {
			return false
		}
		uniq[v] = struct{}{}
	}
	return true
}
