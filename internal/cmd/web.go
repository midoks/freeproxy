package cmd

import (
	"fmt"

	"github.com/oschwald/geoip2-golang"
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

	bytes, readFileError := conf.App.StaticFile.ReadFile("static/GeoLite2-Country.mmdb")

	if readFileError != nil {
		fmt.Sprintf("read geo error: %v", readFileError)
		return readFileError
	}
	db, dbErr := geoip2.FromBytes(bytes)
	if dbErr != nil {
		fmt.Sprintf("read geoip2 error: %v", dbErr)
		return dbErr
	}

	fmt.Println("ddd:", db)
	app.Start(conf.Http.Port)
	return nil
}
