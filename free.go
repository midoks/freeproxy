package main

import (
	"embed"
	"os"

	"github.com/urfave/cli"

	"github.com/midoks/freeproxy/internal/cmd"
	"github.com/midoks/freeproxy/internal/conf"
	"github.com/midoks/freeproxy/internal/log"
)

//go:embed static/*
var assetFS embed.FS

func init() {
	conf.App.Version = "0.0.1"
	conf.App.StaticFile = assetFS
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
