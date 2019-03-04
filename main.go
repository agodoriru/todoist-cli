package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "todoist-cli"
	app.Usage = "make an explosive entrance"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		// fmt.Printf("Hello %q", c.Args().Get(0))



		if c.Args().Get(0) == "list"{
			fmt.Println("show list sitai")
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
