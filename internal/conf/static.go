package conf

import (
// "net/url"
// "os"
)

var (
	// Application settings
	App struct {
		// ⚠️ WARNING: Should only be set by the main package (i.e. "facegit-web.go").
		Version string `ini:"-"`

		Name    string
		RunMode string
	}

	// log
	Log struct {
		Format   string
		RootPath string
	}

	// Cache settings
	Cache struct {
		Adapter  string
		Interval int
		Host     string
	}

	// http settings
	Http struct {
		Port int `ini:"http_port"`
	}
)
