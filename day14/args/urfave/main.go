package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the lonelines!"
	app.Action = func(c *cli.Context) error {
		//fmt.Println("Hello friend!")
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args()[0]
		}
		fmt.Println("Hello Friend cmd : ", cmd)
		return nil
	}
	app.Run(os.Args)
}
