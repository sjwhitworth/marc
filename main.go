package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	fpath := flag.String("f", "", "The filepath of the text file")
	length := flag.Int("length", 10, "The length of sentence in words")
	flag.Parse()

	if *fpath == "" {
		fmt.Println("[Error] I need a valid filepath in order to be able to spew junk.")
		return
	}

	m := NewMarkov()
	err := m.Seed(*fpath)
	if err != nil {
		fmt.Println("[Error] Could not open file, check your filepath is all good.")
	}

	for {
		fmt.Println(m.GenerateText(*length))
		time.Sleep(5 * time.Second)
	}
}
