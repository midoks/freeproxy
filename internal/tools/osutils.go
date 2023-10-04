package tools

import (
	"encoding/json"
	"os"
	"os/user"
)

// IsFile returns true if given path exists as a file (i.e. not a directory).
func IsFile(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsDir returns true if given path is a directory, and returns false when it's
// a file or does not exist.
func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	if e != nil {
		return false
	}
	return f.IsDir()
}

// IsExist returns true if a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// CurrentUsername returns the username of the current user.
func CurrentUsername() string {
	username := os.Getenv("USER")
	if len(username) > 0 {
		return username
	}

	username = os.Getenv("USERNAME")
	if len(username) > 0 {
		return username
	}

	if user, err := user.Current(); err == nil {
		username = user.Username
	}
	return username
}

// https://stackoverflow.com/questions/22128282/how-to-check-string-is-in-json-format
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func StringContains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func IntContains(arr []int, str int) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
