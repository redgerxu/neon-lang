package main

import (
	"fmt"
	"log"
	"os"

	"github.com/narutopig/neon-lang/lib"
)

func main() {
	// cli entry point
	if len(os.Args) < 2 {
		log.Panic("No input file found")
		os.Exit(1)
	}

	var fileName string = os.Args[1]

	c, err := os.ReadFile(fileName)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	content := string(c)

	tokens, err := lib.Parse(content)
	if err != nil {
		log.Panic(err)
		return
	}

	for _, t := range tokens {
		fmt.Println(t)
	}
}
