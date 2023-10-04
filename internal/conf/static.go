package conf

import (
	"embed"
)

var (
	// Application settings
	App struct {
		// ⚠️ WARNING: Should only be set by the main package (i.e. "facegit-web.go").
		Version string `ini:"-"`

		Name    string
		RunMode string
		RunUser string

		StaticFile embed.FS
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
		Port int `ini:"port"`
	}

	// Security settings
	Security struct {
		InstallLock bool
		SecretKey   string
	}
)
