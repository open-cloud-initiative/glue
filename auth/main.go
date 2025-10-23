package main

import (
	"log"
	"os"

	"github.com/open-cloud-initiative/glue/auth/cmd"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	if err := cmd.Init(); err != nil {
		log.Fatal(err)
	}
}
