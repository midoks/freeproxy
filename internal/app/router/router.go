package router

import (
	// "fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/ini.v1"

	"github.com/midoks/freeproxy/internal/conf"
	"github.com/midoks/freeproxy/internal/log"
	"github.com/midoks/freeproxy/internal/tools"
	"github.com/midoks/freeproxy/internal/tools/debug"
)

func autoMakeCustomConf(customConf string) error {

	_, err := os.Getwd()
	if err != nil {
		return err
	}

	if tools.IsExist(customConf) {
		return nil
	}

	// Auto make custom conf file
	cfg := ini.Empty()
	if tools.IsFile(customConf) {
		// Keeps custom settings if there is already something.
		if err := cfg.Append(customConf); err != nil {
			return err
		}
	}

	cfg.Section("").Key("app_name").SetValue("freeproxy")
	cfg.Section("").Key("run_mode").SetValue("prod")

	cfg.Section("http").Key("port").SetValue("2000")
	cfg.Section("session").Key("provider").SetValue("file")

	os.MkdirAll(filepath.Dir(customConf), os.ModePerm)
	if err := cfg.SaveTo(customConf); err != nil {
		return err
	}

	return nil
}

func Init(customConf string) error {
	var err error

	if strings.EqualFold(customConf, "") {
		customConf = filepath.Join(conf.CustomDir(), "conf", "app.conf")
	} else {
		customConf, err = filepath.Abs(customConf)
		if err != nil {
			return errors.Wrap(err, "custom conf file get absolute path")
		}
	}

	err = autoMakeCustomConf(customConf)
	if err != nil {
		return err
	}

	conf.Init(customConf)
	log.Init()

	if strings.EqualFold(conf.App.RunMode, "dev") {
		go debug.Pprof()
	}

	return nil
}
