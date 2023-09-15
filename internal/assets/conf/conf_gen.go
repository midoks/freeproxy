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

var _confAppConf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\x41\x6f\xdb\x3e\x0c\xc5\xef\xfa\x14\x0f\x3e\xfb\xef\xbf\xdd\x0e\x6d\xe7\xc0\xa7\x02\xdd\x8e\xc3\x5a\x60\x87\x62\x70\x34\x8b\x4e\xb4\x58\xa2\x21\x29\x69\xb2\x4f\x3f\x48\xb1\xbd\x64\xe8\x75\x47\x52\xe4\x8f\x4f\x4f\x94\x1c\xc7\xd6\x4a\x43\x68\x90\xf5\x8e\x68\x74\x7c\x3c\x65\xc2\xed\x6d\xbb\xf7\xe4\x62\xda\x31\x87\x73\xc6\xb0\x4a\x85\x8a\x0e\x99\x10\xaf\x03\x6f\xbe\x8b\x15\x5e\xb6\x84\x81\x37\xe8\xd9\x19\x19\x40\x3a\x6c\xc9\x21\xfb\xe9\xd9\x66\x60\x87\x2c\xd0\x31\x64\x62\x3a\x6e\xe6\x38\x62\xdb\x51\x86\x6d\x4c\x0d\xbc\xf1\x11\xd9\xc9\x6e\x4b\x33\x34\x05\x90\x4a\x8e\x81\x5c\xbe\x80\x0d\x19\x76\xa7\x2c\x47\xe6\x48\x69\x9f\xe5\x69\x88\x21\x93\xea\xb3\x42\x4c\x1d\x68\x70\x2e\x15\x2b\x3c\xf1\x9f\x46\xb0\x1d\x4e\x39\x3e\x3d\x42\xdb\x40\xee\x20\x07\x68\x0b\x4f\x1d\x5b\xe5\x0b\xb1\x24\x1b\x54\x37\xe5\xdc\x7b\x1e\x05\x69\xd5\xc5\xa8\x1c\x1d\x5b\x4b\x5d\xd0\x6c\xb1\x65\x1f\x20\x95\x72\xe4\x7d\x2d\x56\xf8\x0f\xa9\xa7\x86\xa5\xf0\xc6\x6e\xd7\x84\x6e\xcc\xe3\x79\x53\xdf\xdd\xde\x7f\xcc\x47\xe9\xfd\x1b\x3b\xd5\x18\xd9\x49\xc7\x36\x57\x3f\x9a\x32\x1f\x99\x87\xd6\xeb\x5f\xd4\x54\x65\x99\x6b\x35\x50\x1b\xb4\x21\xde\x87\xa6\x7a\x28\x13\x76\x1e\x5f\x63\x5d\xdd\xdc\x17\x65\x51\x16\x55\x5d\x55\x37\x55\xb5\x16\x49\x44\x14\x7e\x95\x17\xe2\xd5\x93\xf7\x9a\xed\x6c\xed\x14\x62\x74\x7c\xd0\xea\x7d\x77\x7b\x3d\xd0\x64\xee\xf9\xf6\x85\x98\xcb\x2f\xad\x4d\x2f\xc5\xb6\xd7\x9b\xbd\x93\xc9\x89\x9e\x1d\x1c\xf9\x31\x1a\x73\xa0\x65\x46\x3d\xab\x67\x77\xaa\xa1\x98\x3c\x2c\x07\x58\x22\x05\x69\x4f\x13\x04\x27\x0a\xa9\x30\x8e\xaf\x17\xa1\x31\x42\xdc\x96\x1c\x54\x6c\x0a\xac\x95\x0c\xf2\xff\xe9\xd4\xaf\xff\x95\xdf\xb3\xf6\x76\x12\xd7\xe0\x6a\xee\x72\x7d\xde\x69\x42\xfa\x47\x81\xe1\x03\x3b\x42\xb8\xb0\x59\x2b\xb2\x41\xf7\x9a\x5c\x21\xce\xc5\xf3\xa7\xd3\xed\xa0\x77\xd4\x6a\x23\xf5\x20\x56\xf8\xb6\xa5\xf4\x0c\x91\x42\x61\x06\x6b\x8b\xcf\x2f\x2f\x5f\x9e\xd3\xe6\x2e\x04\x4f\xdd\xde\x45\x46\x2f\x07\x4f\x93\x94\xf7\xb7\x3a\x3d\xc9\x2c\x26\xde\xa0\x10\x9b\xae\xbd\xd8\xf4\xdb\xbb\xb2\x9c\x08\x46\x1e\xb5\xd9\x1b\x0c\xba\x27\x44\x2f\xfe\xe6\xc8\x99\x54\x08\x23\x8f\x6d\xac\x4b\x9e\xa1\xc1\xc3\xdd\x87\x85\x73\x69\x4a\x6c\x7b\x7c\xfe\xfa\x84\xc0\x3b\xb2\x85\xe8\xbc\xeb\xdb\x6b\x23\xda\x98\x13\xe2\x77\x00\x00\x00\xff\xff\x1b\x85\xa5\x6f\x90\x04\x00\x00"

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
