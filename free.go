package main

import (
	"os"

	"github.com/midoks/freeproxy/internal/cmd"
	"github.com/midoks/freeproxy/internal/conf"
	"github.com/midoks/freeproxy/internal/log"
	"github.com/urfave/cli"
)

func init() {
	conf.App.Version = "0.0.1"
}

func main() {

	app := cli.NewApp()
	app.Name = "FreeProxy"
	app.Usage = "free proxy service"
	app.Version = conf.App.Version
	app.Commands = []cli.Command{
		cmd.Run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Infof("Failed to start application: %v", err)
	}
}
