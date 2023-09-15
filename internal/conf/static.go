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

	// web settings
	Rpc struct {
		HttpAddr    string `ini:"http_addr"`
		HttpPort    int    `ini:"http_port"`
		Domain      string
		AppDataPath string
	}

	// Repo settings
	Repo struct {
		RootPath string
	}

	// Other settings
	Other struct {
		ShowFooterBranding         bool
		ShowFooterTemplateLoadTime bool
	}
)
