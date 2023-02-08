package main

import (
	"cliapp/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Entrypoint")

	application := app.Generate()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
