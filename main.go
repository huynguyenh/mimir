package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

// 1. get PR
// 2, get file changes
// 3. do stuff
// 4. comment to that PR
func main() {
	var branch string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "branch",
			Destination: &branch,
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println(branch)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
