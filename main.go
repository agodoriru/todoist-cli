package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "todoist-cli"
	app.Usage = "make an explosive entrance"
	app.Version = "0.0.1"

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
