// Code generated for package conf by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../conf/app.conf
package conf

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _confAppConf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\xc1\x6e\xdb\x30\x0c\xbd\xeb\x2b\x08\x9f\x3d\xcf\x4e\x87\xb6\x73\xa0\x53\x81\x6e\xc7\x61\x2d\xb0\x43\x51\x38\x9a\x45\xc7\x5a\x2c\xd1\x90\x98\x34\xd9\xd7\x0f\x52\x6c\x2f\x19\x7a\xdd\x91\xd4\xe3\x7b\x4f\x4f\x94\x1a\xc7\xc6\x29\x8b\x20\x21\xeb\x3c\xe2\xe8\xe9\x78\xca\x84\xdf\xbb\x66\x1f\xd0\xc7\xb6\x27\xe2\x73\xc7\x92\x4e\x40\x8d\x87\x4c\x88\x97\x9e\x79\x7c\x15\x23\x79\x06\x09\xab\xb2\x2c\x85\x78\x19\x68\xfb\x2a\xd6\xf0\xdc\x23\x0c\xb4\x85\x8e\xbc\x55\x0c\x68\xb8\x47\x0f\xd9\xaf\x40\x2e\x03\xf2\x90\x31\x1e\x39\x13\xd3\xb1\x9c\xeb\x28\xd5\x8c\x8a\xfb\xd8\x1a\x68\x1b\xa2\x4c\xab\xda\x1e\x67\xd2\x54\x80\xd2\x6a\x64\xf4\xf9\x42\x6c\xd1\x92\x3f\x65\x39\x64\x1e\xb5\x09\x59\x9e\x44\x2c\xda\x84\xcf\x0a\x31\x4d\x80\x84\x33\x54\xac\xe1\x91\xfe\x0e\x02\xb9\xe1\x94\xc3\x97\x07\x30\x8e\xd1\x1f\xd4\x00\xc6\x41\xc0\x96\x9c\x0e\x85\x58\x9a\x12\xaa\x55\x39\xcf\x9e\xa5\x40\x39\x7d\x21\x95\x43\x4b\xce\x61\xcb\x86\x1c\xf4\x14\x18\x94\xd6\x1e\x43\xa8\xc5\x1a\x3e\x40\x9a\xa9\xc1\x21\xbf\x91\xdf\x49\x6e\xc7\x3c\x9e\xcb\xfa\xf6\xe6\xee\x73\x3e\xaa\x10\xde\xc8\x6b\x69\x55\xab\x3c\xb9\x5c\xff\x94\x65\x3e\x12\x0d\x4d\x30\xbf\x51\x56\x65\x99\x1b\x3d\x60\xc3\xc6\x22\xed\x59\x56\xf7\x65\xa2\x9d\xe5\x6b\xd8\x54\xab\xbb\xa2\x2c\xca\xa2\xaa\xab\x6a\x55\x55\x1b\x91\x4c\x44\xe3\x57\x7d\x21\x5e\x02\x86\x60\xc8\xcd\xd1\x4e\x25\x8c\x9e\x0e\x46\xbf\x9f\x6e\x67\x06\x9c\xc2\x3d\xdf\xbe\x10\x33\xfc\x32\xda\xf4\x52\xe4\x3a\xb3\xdd\x7b\x95\x92\xe8\xc8\x83\xc7\x30\xc6\x60\x0e\xb8\x68\xd4\xb3\x7b\xf2\xa7\x1a\x34\x61\x00\x47\x0c\x0e\x51\x83\x72\xa7\x89\x04\x4e\xc8\x09\x18\xe5\xeb\xc5\x68\xac\x20\x6e\x4b\x0e\x58\x6c\x0b\xd8\x68\xc5\xea\xe3\x74\x1a\x36\xff\x2b\xef\xd9\x7b\x33\x99\x93\x70\xa5\xbb\x5c\x9f\x76\x06\x21\xfd\x2d\x26\x08\x4c\x1e\x81\x2f\x62\x36\x1a\x1d\x9b\xce\xa0\x2f\xc4\x19\x3c\x7f\x44\xd3\x0c\x66\x87\x8d\xb1\xca\x0c\x62\x0d\x3f\x7a\x4c\xcf\x10\x59\x90\x67\x62\xe3\xe0\xeb\xf3\xf3\xb7\xa7\xb4\xb9\x0b\x43\xc0\x76\xef\x23\x47\xa7\x86\x80\x93\x95\xf7\xb7\x3a\x3d\xc9\x6c\x26\xde\xa0\x10\xdb\xb6\xb9\xd8\xf4\x9b\xdb\xb2\x9c\x18\xac\x3a\x1a\xbb\xb7\x30\x98\x0e\x21\x66\xf1\x2f\x8f\x9a\x99\x0a\x61\xd5\xb1\x89\xb8\x94\x19\x48\xb8\xbf\xfd\xb4\xf0\x5c\x86\x12\xc7\x1e\x9e\xbe\x3f\x02\xd3\x0e\x5d\x21\xda\xe0\xbb\xe6\x3a\x88\x26\xf6\x84\xf8\x13\x00\x00\xff\xff\x11\x95\x0b\x6b\xa4\x04\x00\x00"

func confAppConfBytes() ([]byte, error) {
	return bindataRead(
		_confAppConf,
		"conf/app.conf",
	)
}

func confAppConf() (*asset, error) {
	bytes, err := confAppConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/app.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"conf/app.conf": confAppConf,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"conf": &bintree{nil, map[string]*bintree{
		"app.conf": &bintree{confAppConf, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
