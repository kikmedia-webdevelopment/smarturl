package main

import (
	"log"

	"github.com/juliankoehn/mchurl/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
