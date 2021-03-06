// Code generated by go-bindata.
// sources:
// app/schema/service.json
// app/schema/stack.json
// app/docs/swagger.json
// DO NOT EDIT!

package asset

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

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemaServiceJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x31\x4f\x03\x31\x0c\x85\xf7\xfb\x15\x96\xd5\x8d\x86\x54\x82\x85\xfc\x06\x06\xf6\xea\xa8\xac\xc6\xbd\x0b\xf4\x92\xe0\x98\x01\x50\xff\x3b\x4a\x52\x55\xb7\x77\xc9\xf0\xf2\xbd\xf7\x6c\xff\x0d\x00\xb8\x29\xc7\x99\x17\x42\x07\x38\xab\x66\x67\xed\x47\x49\xd1\x74\xf5\x31\xc9\x64\xbd\xd0\x49\xcd\xee\xd9\x5e\xc9\x6d\xb5\x65\x49\x99\x45\x03\x17\x74\x50\x83\x00\xb0\x28\x1d\x3f\x0f\xc1\xdf\x14\x00\xd4\x9f\xcc\x35\xbb\xa8\x84\x38\x35\x6f\xd3\x33\xa9\xb2\xc4\xfa\xf5\xbe\x27\xf3\x3b\xd6\x67\x67\x5e\x0e\xe3\xc3\x06\x1b\x74\xe9\x2c\x46\x5a\xf8\xee\xc4\x1b\xb5\x84\xf8\xca\x71\xd2\x19\x1d\x3c\xf5\x9e\xe1\xda\x85\xc2\x5f\xdf\x41\xb8\xce\xbf\xef\xb5\xdb\xd5\x52\x63\x63\xc8\xfb\xa0\x21\x45\x3a\xbf\xad\x2f\x70\xa2\x73\xe1\xe1\x32\xfc\x07\x00\x00\xff\xff\x67\xf5\x00\x07\x53\x01\x00\x00")

func schemaServiceJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaServiceJson,
		"schema/service.json",
	)
}

func schemaServiceJson() (*asset, error) {
	bytes, err := schemaServiceJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/service.json", size: 339, mode: os.FileMode(420), modTime: time.Unix(1548716490, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaStackJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\xc1\x4a\xc4\x30\x10\x86\xef\x7d\x8a\x61\xd8\x83\xba\xc6\x2e\xe8\xc5\x5c\x7c\x01\x11\xc1\xe3\x52\x61\xd8\xce\xb6\x23\xdb\x34\x4e\x46\x41\x65\xdf\x5d\x92\x14\xed\x2d\xf3\xcd\xff\xe5\x4f\x7e\x1a\x00\xdc\xa4\xc3\xc8\x13\xa1\x07\x1c\xcd\xa2\x6f\xdb\xb7\x34\x07\x57\xe9\xcd\xac\x43\xdb\x2b\x1d\xcd\xed\xee\xda\x25\x79\x9d\xb5\xa8\x73\x64\x35\xe1\x84\x1e\xf2\x45\x00\xf8\x44\x13\xff\x4d\x00\x68\x5f\x31\xcf\x98\x4c\x25\x0c\xc5\x2b\x3c\x92\x19\x6b\xc8\xab\xd7\x3d\xb9\xef\x6e\x7b\xf1\xe0\x5d\x3e\xed\xdc\x7d\xb7\xbd\xbc\xda\xfc\x67\x27\x09\x8f\x1c\x06\x1b\xd1\xc3\x6d\x81\xe7\xba\xc3\x17\xd6\x4f\x39\xac\xfa\x57\x8d\x12\x8c\x07\x56\xac\x42\xb3\x48\xa8\xfc\xfe\x21\xca\x3d\x7a\xd8\xd7\xd7\x76\x85\x53\xdf\x8b\xc9\x1c\xe8\xf4\xbc\xfe\xd6\x91\x4e\x89\x9b\x73\xf3\x1b\x00\x00\xff\xff\x3a\xee\x42\x2c\x28\x01\x00\x00")

func schemaStackJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaStackJson,
		"schema/stack.json",
	)
}

func schemaStackJson() (*asset, error) {
	bytes, err := schemaStackJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/stack.json", size: 296, mode: os.FileMode(420), modTime: time.Unix(1548716490, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xdf\x6f\xdb\xb6\x13\x7f\xf7\x5f\x41\xe8\xdb\x87\xef\xb6\xca\x72\xd2\x60\x68\xfc\xb2\x05\x4d\xb1\x04\xd8\x82\x62\xe9\xf6\x92\x79\x01\x2d\x9d\x25\x36\x92\xc8\x90\x94\x1b\xd7\xd0\xff\x3e\x90\xfa\x61\x52\x92\x1d\xc9\xb0\xdb\x60\xd8\x9b\x4c\xde\x1d\xef\x3e\xfc\xf0\x78\x47\xaf\x47\x08\x39\x3e\x4d\x45\x96\x80\x70\xa6\xe8\x6e\x84\x10\x42\x0e\x66\x2c\x26\x3e\x96\x84\xa6\xde\x27\x41\x53\x67\x84\xd0\xec\xb5\x92\x65\x9c\x06\x99\xdf\x4f\x56\xf8\x11\x58\x66\x23\x29\x99\x30\xe6\x3f\xe3\x30\x04\xee\x4c\x91\x73\x3a\x9e\x38\x7a\x8c\xa4\x0b\xea\x4c\xd1\xba\x50\x90\x44\xc6\xa0\xe6\xaf\x56\x0c\xb8\x1f\xd3\x2c\x40\xb7\xc0\x97\xc0\xd1\xc5\x87\xeb\xb1\x56\x29\x02\x90\xd8\x97\xb5\x1e\x42\x4e\x8a\x13\xad\x78\xf1\x04\x31\x7a\x2f\xfd\x08\x96\xc0\xf9\xaa\xd4\x40\xc8\x81\x04\x93\x58\x49\xe0\x27\x88\x7f\x86\x5a\x62\x3c\x27\x5f\x1c\x2d\x94\x97\xd6\x63\xe2\x43\x2a\xa0\xcb\xfa\x6f\xd7\x1f\x37\x16\x33\xae\xed\xa9\x20\xa7\x9e\x47\x19\xa4\x82\x66\xdc\x87\x31\xe5\xa1\x57\x1a\x11\x9e\x52\xb1\xcc\x2f\x81\x0b\x42\x53\xa5\x7a\x32\x9e\x8c\x27\x6a\x56\xcf\x39\x11\x15\x2a\x26\x27\xa6\x3e\x8e\xf5\x0f\x3d\x3c\xc7\x02\x3e\x60\x19\xa9\x29\x6f\x79\x52\x0c\x32\x2c\x23\xb1\x01\xce\x5b\x9e\x78\x01\xb0\x98\xae\xbc\xb5\xb1\x41\xf7\x24\xc8\xbd\x88\xd2\x07\xe1\xad\x19\xa7\x4b\x12\x00\xcf\xcd\xc0\x42\x30\x51\x44\xc8\x09\x40\xf8\x9c\x30\x59\x7a\xf8\x91\x13\xb5\x65\x28\x85\xcf\xa8\xb0\x9f\x40\x2a\xd1\x82\x72\xd4\x5c\x07\x2d\x38\x4d\x90\xb1\xcc\xeb\x8d\x55\x89\xc3\x0d\x2d\xca\xb1\x4b\x6d\xce\xa9\x87\x66\x86\x3c\x65\xc0\xb5\xe1\xeb\x40\x79\xc1\xa8\x90\x85\xf8\x15\xa5\x0f\x57\x38\x0d\x62\xe0\xa6\x7d\x0e\x82\x51\x05\xb7\x0a\x26\x2f\xc7\x73\x0b\x76\x85\x10\x2c\x21\x95\x62\x48\xfc\xbf\x80\x44\x32\x02\xa4\x35\x51\x4c\xca\x3d\xd9\x11\xd6\x7b\x25\xd9\x2b\xaa\x10\xa4\x16\x16\xcf\x05\x64\x9a\x3f\x9d\x4c\x1a\x43\x08\x39\xaf\x38\x2c\x94\xc5\xff\x79\xb5\x9e\xd7\xf0\xa3\x82\xc3\xfc\x6a\x03\x94\xd2\x00\xf6\xc2\x47\x2b\xf6\xc2\xe7\x86\x06\xd0\x17\x1e\x25\x7b\x1c\x74\x6c\x2f\xfa\x81\x23\x80\x2f\x89\xbf\x1f\x3e\x95\x6e\x2f\x88\x6e\x0b\xe1\xbe\x28\x95\xe2\x5d\x40\x31\xcc\x71\x02\x12\x78\x73\x8d\x06\x44\x0c\x4b\x09\x5c\x7b\xfc\xf7\xdd\xc4\x3d\xc7\xee\x62\xb6\x7e\x9b\xbb\xf5\xf7\x59\xee\x9e\xd5\x3f\xde\xe4\xee\xdd\xdb\x73\x3c\x9f\x59\x23\xd5\xf7\xc9\x69\xfe\xca\xf0\xa1\x88\x72\xc5\x74\xfe\x14\x92\x93\x34\x6c\xce\x56\xd9\xf5\x56\x62\xff\xe1\xfa\xb2\x39\x4d\xb4\x5f\x8f\x19\xf0\x55\xf7\x96\xcd\x0e\x4a\x8d\x16\xfa\x9d\xec\xa8\x6f\x01\x56\x24\xed\xed\x14\x78\xc7\x01\xcb\x9a\x01\x7f\xa5\xbf\xc3\x63\x06\x42\x4e\x51\xb9\x50\x31\x5f\x8e\x1e\x92\x1a\xca\xb3\x52\x7e\xd8\x19\x3a\x39\x24\x50\x5b\x8f\x91\xb7\x26\xc1\xa0\xbb\x48\x9d\x25\x5c\xe1\x88\xe6\x2b\x44\x82\x67\xe1\x52\x8c\x1a\x78\x8e\x8e\x92\x6f\x0e\x4d\xaa\x3f\x58\xb0\x93\x54\x07\xe5\x51\xf6\x02\xa0\xd9\x4d\x23\x4f\x48\x3c\xfc\x62\xd7\x4a\xba\x96\xf9\x37\x93\xaa\x8d\x9c\xf2\x7f\xbf\x5b\x4c\x6b\xf6\xbb\xc3\x06\x81\xa4\xed\x1e\x07\x23\xdb\x8f\x03\xe5\x72\x6d\xf4\x60\x08\xe8\x44\xbd\x07\x04\x7d\xf3\x74\x0f\x08\xb6\x90\x64\xdf\x1c\xad\x94\x8f\x71\x98\x94\xf0\x37\xa3\xc9\x2e\x8c\x3a\x8b\xc3\x67\xe9\x54\x26\x21\xff\x61\x60\x7d\x38\x18\xb3\x1d\x25\xe2\x37\x4b\x43\x65\x1f\xec\xc5\x58\x82\x85\x53\xef\x74\x54\x68\xa2\xaa\xa1\x7e\x06\xb5\x3f\x4b\xb1\x9e\xb8\x95\xe2\xba\x8d\x3a\x0a\x6c\x2d\x7f\x76\xc0\x56\x3f\x11\x04\xb0\x20\x29\x51\x8e\x1a\xcd\x7f\xb5\x01\x06\x84\x55\xc1\x4d\xe7\x9f\xc0\xdf\x10\x6a\xf3\xc0\x52\xea\x20\x0e\x8c\x83\x50\x5d\x28\xc2\x55\xf1\x30\xde\xc8\x73\x78\xcc\x08\x87\xc0\x02\xb3\x28\xd8\x47\x0d\x14\x1d\xc6\x15\x8a\x92\x34\x30\x71\x7c\x9d\x34\x83\x7b\x2c\x9b\x58\xed\x68\x0b\x9c\x05\xe5\x89\xd6\x70\x54\xa5\xe3\x4a\x92\x80\x2d\xf0\xe4\x86\xd4\xad\x5a\x87\x22\x31\x07\x17\x46\xc7\x9b\x1b\xbb\x15\x51\x21\x5b\x3b\x55\xad\x8e\x39\xc7\x2b\xdb\x36\x91\x90\x34\xe5\xdb\xfe\x9a\x3b\xb7\xdd\xb5\x2b\xbd\x76\xa7\x5b\x24\x18\x82\xc8\x51\x9b\x34\xdb\xe7\xeb\xcb\x6e\x87\xcb\x79\xcb\xe5\xc6\xe1\xac\x88\xa5\x44\xc7\xf6\x12\xbb\x82\x4b\x48\xfa\x2b\xa4\xa1\x7e\xe5\x7a\xb3\x3d\x6c\xec\x7e\x99\xb8\xe7\xb3\x1f\xfe\xff\xd3\xd4\xad\x7f\x7c\xf7\xfd\xae\x58\x6e\x0c\xb2\xda\xd1\xe8\xd4\x7b\xff\x62\x37\xa1\x6a\x87\x3b\x7d\xcf\x74\x03\x70\xdc\x43\x55\x34\x19\xf6\xa1\x6a\xd5\x4c\x5a\x83\x61\xff\x01\x87\x5a\x29\x24\x32\xca\xe6\x63\x9f\x26\x5e\xb4\x62\xc0\x85\x8f\x63\x28\x3e\xf5\x7b\xae\x17\x60\x89\xe7\x58\x80\x07\xa9\x24\x72\x65\x3f\x8e\x16\x57\xdb\xa0\x44\xa6\xef\x4f\x3b\x8d\xa9\xa1\x67\x93\xd8\xcd\xa0\x24\x76\xd3\x83\xf9\xda\x93\xaf\xcc\xfb\x4e\x72\xdc\xb6\x8b\x11\xcb\x11\x92\x4a\x08\xad\xfb\xcc\xe2\x06\x49\xe5\x8f\x67\x5f\x79\xcf\xdf\xd1\x84\x51\x31\xf0\x0e\x33\x14\x3b\x18\x80\xca\x29\x83\x09\x2f\xf8\x8a\x7a\xb1\x69\x68\x5b\x06\xfa\x2f\x7b\x0e\xcf\x9e\x86\x0b\x9b\xbf\x81\xf6\x3f\xa2\xdb\x17\x6f\x15\x97\x47\x3b\xc4\x75\x69\xda\xae\x88\xbb\x0a\x53\xfd\x1f\x21\xb6\x4f\xdf\xa6\x3a\x36\xca\xdb\x66\x5b\x91\xef\xbe\x27\x06\xd8\x35\xbb\xa7\x0e\xab\x1d\x99\x68\xa0\xf1\xca\x42\xbb\x86\x1f\xe5\xa3\x7f\x02\x00\x00\xff\xff\x19\xcf\x2f\xa1\x82\x1d\x00\x00")

func docsSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_docsSwaggerJson,
		"docs/swagger.json",
	)
}

func docsSwaggerJson() (*asset, error) {
	bytes, err := docsSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/swagger.json", size: 7554, mode: os.FileMode(420), modTime: time.Unix(1548716490, 0)}
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
	"schema/service.json": schemaServiceJson,
	"schema/stack.json": schemaStackJson,
	"docs/swagger.json": docsSwaggerJson,
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
	"docs": &bintree{nil, map[string]*bintree{
		"swagger.json": &bintree{docsSwaggerJson, map[string]*bintree{}},
	}},
	"schema": &bintree{nil, map[string]*bintree{
		"service.json": &bintree{schemaServiceJson, map[string]*bintree{}},
		"stack.json": &bintree{schemaStackJson, map[string]*bintree{}},
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

