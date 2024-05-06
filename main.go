package main

import (
	"command-line-app/app"
	"log"
	"os"
)

func main() {
	application := app.Gerar()
	application.Run(os.Args)
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
