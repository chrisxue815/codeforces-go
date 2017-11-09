package main

import (
	"fmt"
)

func main() {
	var visited [26]int
	var line string
	fmt.Scanln(&line)

	for i := 0; i < len(line); i++ {
		visited[line[i]-'a'] = 1
	}

	count := 0
	for i := 0; i < 26; i++ {
		count += visited[i]
	}

	if count&1 == 1 {
		fmt.Print("IGNORE HIM!")
	} else {
		fmt.Print("CHAT WITH HER!")
	}
}
