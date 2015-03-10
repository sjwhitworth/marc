package main

import (
	"fmt"
	"time"
)

func main() {
	m := NewMarkov()
	err := m.Seed("/Users/stephenwhitworth/Desktop/test.txt")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println(m.GenerateText(100))
		time.Sleep(1 * time.Second)
	}
}
