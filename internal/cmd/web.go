package cmd

import (
	// "fmt"
	"github.com/urfave/cli"
	// "github.com/go-facegit/facegit-rpc/internal/app"
	// "github.com/go-facegit/facegit-rpc/internal/app/router"
	// "github.com/go-facegit/facegit-rpc/internal/conf"
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

	app.Start(conf.Rpc.HttpPort)
	return nil
}
