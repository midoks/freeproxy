package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"github.com/urfave/cli"

	"github.com/midoks/freeproxy/internal/app/router"
	"github.com/midoks/freeproxy/internal/conf"
	"github.com/midoks/freeproxy/internal/structs"
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

	verbose := true

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

	quit := make(chan bool)
	proxiesChan := make(chan []structs.Proxy, 99999)
	Collect(db, quit, proxiesChan, []string{"http", "https", "socks4", "socks5"}, []string{}, []int{}, verbose)

	fmt.Println("ddd:", proxiesChan)

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%d", conf.Http.Port))
	return nil
}
