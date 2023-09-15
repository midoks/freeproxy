package cmd

import (
	// "fmt"
	"github.com/urfave/cli"

	"github.com/midoks/freeproxy/internal/app"
	"github.com/midoks/freeproxy/internal/app/router"
	"github.com/midoks/freeproxy/internal/conf"
)

var Run = cli.Command{
	Name:        "web",
	Usage:       "This command starts all free proxy web service",
	Description: `Start Free Proxy Service`,
	Action:      WebRun,
	Flags: []cli.Flag{
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func WebRun(c *cli.Context) error {
	err := router.Init(c.String("config"))
	if err != nil {
		return err
	}

	app.Start(conf.Http.Port)
	return nil
}
