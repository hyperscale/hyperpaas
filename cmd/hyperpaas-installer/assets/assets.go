// Code generated by go-bindata.
// sources:
// static/config/docker-compose.yml
// static/index.html
// DO NOT EDIT!

package assets

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

var _staticConfigDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\x10\xc2\x02\xbe\xac\x65\x65\x93\x00\x01\x01\x1d\xbc\x8e\x37\x6b\x40\xb1\x03\xdb\x69\xd1\x93\x40\x8b\x63\x9b\x08\x45\x0a\x24\xe5\x0f\x14\xfd\xef\x05\xf5\x65\xc9\x52\x93\x26\xdb\xdc\x7a\x11\xc0\x99\x37\xc3\xd1\xcc\x7b\x24\xf7\xa0\x34\x93\x02\x23\xe7\xda\xbd\x71\x7a\x1a\xd4\x9e\x45\xa0\x71\x0f\x21\xa3\x08\x6c\xd8\x4b\xc8\x04\x33\x76\x8d\x10\x8b\xc9\x16\x70\xe9\xc0\x57\xee\x6d\x66\x8e\x64\x1c\x13\x41\x73\x0c\x42\x03\xe4\x68\x23\x15\x44\x52\x6c\xd8\xd6\x39\x5b\x07\x03\x92\xb0\xc6\x9a\xc2\x86\xa4\xdc\x80\x30\xea\x94\x48\x26\x8c\xf6\x77\xc6\x24\x5f\xed\x47\x37\x90\x19\xe4\x29\x87\xcc\x48\x0c\xd8\x42\xd0\x88\x52\x05\x5a\x63\x7c\xe7\xbd\x8d\xd6\x67\xf8\xcd\xcd\x35\x5a\x05\xcb\x66\x6d\x51\x0c\x2d\x83\x7b\xce\xe4\xb7\x8b\xca\x10\xd6\x3c\xde\x11\xce\x41\x6c\x3b\x12\x34\xdc\x97\xe9\x3a\xf6\x8b\x09\xe3\xfe\x97\xc9\xe3\x68\x1a\xb4\xbd\xb6\xaf\x64\x0b\x7e\x31\x81\xa1\x35\x0e\x49\x14\xc9\x54\x98\x36\x7a\x2e\x7e\x4a\x6d\x16\x29\xb7\x01\x69\x47\x6d\x52\xdc\x83\x9d\x9c\xbf\x21\x5c\x37\xfd\x54\x46\x2f\xa0\x3a\x4c\xae\x3e\x10\x15\xc7\x92\x76\xe1\x5d\x38\x26\x52\x03\x5d\x9f\x8a\xc9\xb6\x37\x2e\x80\x54\xc6\x84\x09\xff\xcb\xfd\xfc\x71\x34\x9d\x75\x21\x0e\xc4\x44\xbb\x86\x83\xcb\x6d\x00\x7b\xe0\xfe\xfd\xe4\xfb\xf3\x43\xc3\x15\x49\xa1\x53\xde\x61\x72\x41\xd0\x8c\x59\x7e\xbe\xc6\x77\xb7\x9e\xd7\x85\x4b\x14\x6c\xd8\xb1\xec\x6c\x8e\x10\x60\x0e\x52\xbd\xe8\x33\xb3\x77\xa7\x04\x54\x42\x88\xce\x2c\x14\x12\x2e\x4f\xa5\x97\x93\x35\xf0\x0a\x9b\xc9\xc2\xad\xf0\x2e\x13\x06\x94\x20\x1c\xa3\xbe\xed\x49\xbf\x80\x29\xd0\x86\x28\x13\x26\x92\xb3\xe8\x54\x0f\x16\x94\x99\x4c\x99\x52\x0c\x36\x84\xf1\x54\x41\xb9\x29\x08\xaa\x43\x29\xce\x65\xe5\xff\xd0\x3b\x8b\xf6\x55\xbd\x76\x65\xa8\x8b\xfd\x32\xad\x5d\xec\x25\x4f\x63\xa8\x75\xc2\x9c\x12\xc0\x68\xcd\x04\xad\x6a\xd6\x32\x55\x11\x60\x34\xdc\x13\x35\x54\xa9\x18\x96\x8c\x91\xd1\x4b\x05\x32\x44\x6d\xc1\xbc\x02\x4a\xa4\x32\xf5\x7d\x0a\xfc\x9d\x57\xa5\x48\xd2\x35\x67\x7a\x07\xb4\x61\xb5\x94\xc4\x68\x27\xb5\x69\xc5\xde\xdc\x5c\x77\x05\xd7\xcd\x17\xd1\xed\x23\xed\xb3\x38\xf6\x3e\x79\xbf\x8f\x91\xf9\x4f\x6d\xb9\x5c\x13\xfe\x8b\x24\xcd\x46\x5b\x8b\xe3\x2c\x66\xa6\xb6\x46\x28\x86\x58\xaa\x13\x46\x57\x9e\xf7\x58\x99\x15\xd8\x1b\x85\x58\x22\x77\x82\x6f\x2b\xec\xfb\x85\x90\xff\x2f\x27\x36\x8d\x3e\x0f\x92\x1c\x43\x62\x0c\xc4\x89\xd1\x18\x7d\x3b\xf3\xe3\xc0\x04\x95\x07\x8c\xae\xef\xbc\x12\x9d\x70\x12\x41\x0c\xc2\x34\xb6\xd3\x46\x11\x7b\x73\xd4\xeb\x1d\xa0\xbe\x90\x14\x5c\x25\x39\x20\xdf\x47\x31\x11\x64\x0b\xaa\x6c\x50\x9a\x50\x62\x20\xcc\x6f\xbc\x73\x5c\x42\x94\x3d\xf8\x39\xd3\x31\x46\x57\x97\x45\x5f\x79\xda\x2a\xb6\xe0\x4c\x5d\xb0\x35\xdd\x95\x44\x44\x64\x0b\xc2\xa0\x81\x6d\x27\x28\x34\x58\x4b\x69\x6c\xa1\xc9\x00\x8e\x09\x44\xc6\xbf\xea\x96\x69\x9e\x6a\x40\x89\x21\x78\x98\x2f\x86\x76\x91\x01\x40\xec\x99\x92\xa2\xde\x81\x01\x1a\xcf\x67\xcb\xe7\x20\x0c\xe6\xe3\x51\x10\x8e\xe7\xb3\x1f\xd3\x07\xff\x4f\x47\xbf\xb0\x24\xe4\x40\xf6\x10\x4a\x11\x66\x3c\x51\x69\x62\x1c\x6c\x79\xf2\xd5\xc9\xcb\xca\x57\x7f\x5d\xa6\xfa\x3e\x9d\xdd\x87\xd3\xd9\x6a\xb2\xf8\x31\x1a\x4f\x7c\x30\x3b\xef\x12\x32\x0e\xa6\x93\xd9\xaa\x0b\xd4\x64\xb4\x82\x84\xb3\x88\xe8\x73\x3b\x3f\x44\xe8\xf7\x0c\xbe\x6b\xee\xbf\xc0\xd9\x0f\x53\xa5\xda\xf2\x93\xa5\xf8\xda\x11\xd3\xb3\x69\xb6\x4c\x1b\x75\x6a\x10\xb6\x32\x7e\xeb\x66\x61\xe9\x2f\x78\x68\x8f\x7e\xce\xd6\xc3\xd2\xfc\xcf\x64\x5c\x4c\x1e\xa6\xcb\xd5\xe2\x8f\x70\xf4\xbc\xfa\xe9\x3b\x3b\x93\x10\xad\x0f\xd4\xe9\x06\x84\x3f\x57\x4f\xa3\xe5\xf2\xf7\xfb\xf0\x69\x64\xe1\xd9\xfd\xa2\x21\x52\x60\xf4\x90\x09\x0a\x47\xb7\x78\x70\xb8\xff\x3a\xd3\x62\x32\x0a\x1e\x7d\x67\x51\x94\x8a\x16\x40\x78\xdc\x11\xb5\x5c\xcd\x17\xa3\x87\x89\xef\x6c\x18\x07\x7d\xd2\x06\xba\x50\xc1\xfc\x21\x0c\x26\xbf\x4d\x02\xdf\xa1\xb0\x4e\x8b\xd7\xf1\xa7\x70\xbc\x7a\x0b\xb8\xf6\x46\xc5\xa8\x7f\xeb\x79\x5e\xdb\xb9\x51\x52\x18\x10\xb4\xf6\x34\xd5\x18\xf5\xb3\xc7\xee\x2b\x68\x95\x72\xc0\xa8\x6f\x1f\x98\xb8\xd1\xd9\x76\x0c\x08\xb2\xce\xc0\xdd\xd5\x15\xf7\x7f\x41\x3b\xbb\x75\xf9\x67\xff\xab\xf5\x6d\xb5\x16\xec\x6e\x4b\xad\x62\xf8\xdb\xa2\xae\x3b\x2b\x73\x8e\xa4\x8a\xed\x41\x61\x24\xf7\xa0\x38\xc9\x85\x4a\x8c\x21\xd1\x2e\x9f\xa9\x1d\x69\xaf\x4d\xd1\x37\x08\xda\xab\x1d\x10\xf5\x0b\xaa\x88\x2e\x37\xe5\x32\xfa\xf8\xbb\xa5\x79\xe4\xfc\x77\x99\x7b\xb5\x8e\xb7\x7a\x9d\x67\x81\x63\x19\x93\xb5\xe7\xef\x00\x00\x00\xff\xff\xe0\x77\x25\x38\x63\x0f\x00\x00")

func staticConfigDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_staticConfigDockerComposeYml,
		"static/config/docker-compose.yml",
	)
}

func staticConfigDockerComposeYml() (*asset, error) {
	bytes, err := staticConfigDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/config/docker-compose.yml", size: 3939, mode: os.FileMode(420), modTime: time.Unix(1523401393, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x6d\x6f\xdb\x36\x10\xfe\xde\x5f\x71\x23\x06\xd8\xc1\x22\xa9\xdd\x1b\xb6\xd6\x32\xb2\x2e\x29\xd6\xa1\x5b\x8a\xa5\xfb\xb4\x76\x03\x2d\x9e\x6c\xae\x14\x29\x90\x27\xa7\xee\xd0\xff\x3e\x90\x92\x1c\xbf\x50\x8e\xd3\xe5\xc3\xaa\x2f\xa2\xa8\xe7\x9e\xbb\x7b\x78\x7c\x91\x26\x9f\x09\x53\xd0\xaa\x46\x58\x50\xa5\xa6\x0f\x26\xfe\x06\x8a\xeb\x79\xce\x50\xb3\xe9\x03\x00\x80\xc9\x02\xb9\x68\x9b\xe1\xb1\x42\xe2\x50\x2c\xb8\x75\x48\x39\x6b\xa8\x4c\xbe\x63\x90\x6d\x00\x48\x92\xc2\xe9\x4f\xab\x1a\xed\x4b\xce\xaf\xe0\xb9\x76\xc4\x95\x42\x3b\xc9\xda\x57\x37\xd0\x19\x77\x08\x0b\x8b\x65\xce\xb2\x40\xb2\xe3\x46\xf3\x0a\x73\xb6\x94\x78\x5d\x1b\x4b\x0c\x0a\xa3\x09\x35\xe5\xec\x5a\x0a\x5a\xe4\x02\x97\xb2\xc0\x24\x3c\x9c\x82\xd4\x92\x24\x57\x89\x2b\xb8\xc2\xfc\xd1\x76\x50\x4a\xea\xb7\x60\x51\xe5\x4c\x16\x46\x33\xf0\x59\xe7\x4c\x56\x7c\x8e\xd9\xbb\xa4\xed\x6b\x03\x29\xf9\xd2\x3f\xa6\xb2\x30\x3b\x21\x05\x8e\x16\xb4\x20\xaa\xdd\xe3\x2c\x6b\x74\xfd\x76\x9e\x16\xa6\xca\xce\x0a\x65\xb3\x46\x9e\x3d\x4c\x1f\x3d\x4a\xbf\x4c\x6a\x4e\xc5\x22\x2b\x94\x4d\x1a\x99\x56\x52\xa7\x85\x73\xac\x0d\xc0\xd1\x4a\xa1\x5b\x20\x52\x24\xc4\x83\xf4\x3e\x2e\xb7\xef\x21\x74\x1f\xe9\xc4\x15\x56\xd6\x04\xce\x16\x1f\xef\xe5\x6f\xc7\xa6\x93\xac\x65\xea\x6a\x24\xbb\x29\x92\xc9\xcc\x88\xd5\x86\x47\x21\x97\x50\x28\xee\x5c\xce\x2a\x2e\x75\xe2\xc7\x90\x4b\x8d\x96\xdd\x80\x76\x81\x9e\x6d\x0f\xb0\x0b\x9a\x59\xae\x85\xd4\xf3\x08\xac\x4d\xb5\xe6\xba\xc7\x86\xc2\x63\xf1\xa2\xf4\xb8\x88\xa7\x4c\xc8\xe5\x4e\x84\x6d\xd7\x60\xd4\x5d\x75\x0e\x66\x38\x84\xe7\x16\xf9\x50\x12\x1b\x70\x6b\xae\xd9\x8e\xf7\x61\x66\x95\xbc\x73\xc9\xb7\x60\xca\xd2\x21\xf9\xf6\x57\x03\x1e\xa2\xe6\xdc\x8a\x5b\xe0\x31\x93\x64\x70\xd8\x62\x57\x37\x02\x9c\xa4\xd1\xb7\xbb\xda\x1f\x8d\xa3\x22\x9a\x29\x53\xbc\x3d\x32\xa0\x49\x69\x6c\x05\x15\xd2\xc2\x88\x9c\xd5\xc6\x11\x03\x29\x72\x26\xfb\x52\x39\x92\x67\x37\x10\x4f\x9b\xcc\xad\x69\xea\x3b\x10\x04\x12\xc5\x67\xa8\xa0\x34\x36\x67\xc2\xf8\xc9\xc3\xa6\xe7\xe1\x0e\xa6\x84\x95\x69\x2c\xf8\x5a\x9e\x64\x01\x77\x47\x6e\xa9\xeb\x86\xba\x65\x90\xf0\x5d\x97\x6b\xe7\x06\x6a\xc5\x0b\x5c\x18\x25\x70\xed\x3b\x25\x25\x18\xd4\x9c\x08\xad\xce\xd9\x9f\x7f\xf0\xe4\xfd\xc3\xe4\xfb\x37\xdd\xfd\x75\xf2\x3a\x7d\xf3\x05\x03\x27\xdf\x63\xce\xbe\xfe\x86\x75\x8b\x77\xcf\x98\xdd\x41\xbc\xc8\x3c\x3b\x88\xbf\x6f\xb1\xb1\xe2\x52\xb1\xe9\x85\xbf\xf9\x1e\x78\x81\x34\x72\x70\xa1\x0b\xbb\xaa\xe9\xbf\xeb\xdd\xf2\x07\xc1\xbb\xe6\x96\xde\x8e\x70\x89\x67\x9b\xaa\xef\x8a\xda\x59\x7d\x4a\x9a\x5a\x9c\x4b\x47\x76\x95\x34\xce\xcf\xa4\xdf\x1d\xda\x20\xed\x6f\x5d\xff\x7d\x56\xf1\xb6\xaf\x88\xb8\xfb\x8a\xf6\x26\x7f\xb5\x26\x9f\xa4\xb2\x35\x77\xee\xda\xf8\xa5\xfb\x65\xd7\xba\x67\x85\xd7\x0e\xb6\x55\xbe\xe9\x1e\x54\xf5\x06\x72\x77\x65\x8f\x83\x7a\x29\x8f\xd8\x20\x3e\x72\x1f\x29\x8d\xa1\xa3\x37\x80\x60\xed\x15\xaa\xad\x99\x5b\xf4\x67\xb2\x8e\xab\xef\x00\x65\x4c\xcd\x20\x1c\xd2\x72\x26\xa4\xab\x15\x5f\x3d\x06\x6d\x34\x3e\xb9\xcb\x2e\xd3\xf3\x4d\x27\xd9\xba\x79\xa4\x5e\xc7\x4b\x3b\x6b\x88\x8c\x06\x21\x1d\x9f\x29\x14\x21\x33\xa3\xbb\xed\xbb\x3f\x49\xbb\x66\x56\x49\x62\xbe\xe0\xaa\xcd\x2d\x73\x7d\x58\x23\x0d\x33\xd2\x49\x6d\x65\xc5\xed\x8a\x4d\x3b\xfb\x49\xd6\xd2\xdf\xcb\xd0\xdd\x02\x39\x34\x53\x07\x4c\x87\x4f\x81\x43\x8f\x9b\xc7\x62\x7f\x09\x53\x34\x15\x6a\x4a\xb9\x10\x17\x4b\xd4\xf4\x42\x3a\x42\x8d\x76\x3c\x3a\xbf\xfc\xe5\xc7\xf6\x08\xf8\xc2\x70\x81\x62\x74\x0a\xe3\x13\xc8\xa7\xf0\xcf\x96\x43\x7f\xee\x36\x0a\x53\x65\xe6\x11\x9b\x93\x27\xdb\x09\x2d\xb9\x85\xcf\xd1\x3b\x72\x90\x83\xc6\x6b\x08\x5e\xaf\x4c\x63\x0b\x1c\x8f\xb2\xf5\xd0\x64\x2d\xc8\x13\xec\xd9\x77\xa0\xa7\xa4\x21\xbf\xc9\x60\x8e\x74\xa1\xd0\x37\x9f\xae\x9e\x8b\xf1\x68\x5d\x04\x87\x38\xd0\x3e\xf3\xc7\xaa\x03\x34\x6b\x60\x94\xa6\xdd\x06\x9f\x49\x54\xe2\x10\x49\x0b\x8b\x32\x84\xcd\xf2\x56\x82\x80\x8a\xda\xaf\x27\xed\x01\xeb\x1e\xb3\x37\x1c\x1b\x52\xa6\x16\x2b\xb3\xc4\x1f\x88\xac\x9c\x35\x84\xe3\x51\x3f\xa5\xf6\xad\xda\xb1\x89\xd4\x4c\x85\xce\xf1\x39\xfa\x52\xe9\x9a\x91\x8a\xe9\x43\x0f\x2c\x90\xc3\xcf\x57\x97\xbf\xa6\xb5\xff\x62\xef\x8d\x52\xc1\x89\xef\x7a\xdd\x2b\xb6\x60\x3f\x3a\x6d\x79\x62\x68\x59\xc2\x38\xbc\x4c\xc3\x5f\x84\x3c\xcf\xa1\x1f\xce\xb4\x94\x5a\xba\x85\x4f\x2e\x12\x5d\x48\xb2\x17\x2d\x0d\xeb\x60\xda\x2d\x83\x90\xc3\xc8\x2f\x84\xa3\x27\x71\xab\x8f\x12\xb4\xbf\xb2\xec\xec\xd5\xe5\xf9\x25\x58\x14\xd2\x62\x41\x40\x06\x2a\xae\xf9\x1c\x2d\x34\x56\xed\xd9\x7c\x38\x94\x33\x2f\xfc\xb7\x4b\x9b\x75\xc9\xa5\xfa\x9f\xe6\xda\xfb\x42\x6b\x8d\x8d\x64\xb8\xf5\x34\x54\xbf\xed\x34\x8e\x14\x64\xbb\xf0\xfb\x7a\x6c\xab\x24\x5e\x8d\xad\x60\xb5\x0d\xf7\x73\x2c\x79\xa3\x68\x7c\x6b\xfd\x5d\x05\xee\xb0\xa7\x44\xb3\xdc\xd4\xc7\x21\xc5\xc4\x39\x85\x2d\xa1\xee\x30\x2e\xe1\xd3\x71\x14\xf1\x5a\x22\x15\x8b\xcd\x95\x74\x74\x3a\x30\xea\xed\xb7\xe4\x63\x18\xbd\xbc\xbc\x7a\x35\x3a\x8d\x62\x66\x46\xf8\x9d\x1f\xaf\xc1\x0b\x7c\xce\x89\x8f\xb7\x35\x3f\xd9\xb7\xfb\x70\x92\x16\xdc\x07\x61\x91\x3b\x5f\x81\x31\xc9\x37\xd5\x44\x6b\x3b\x6c\x44\x03\x38\xbe\x56\x76\xac\xb7\xaa\x65\xf3\xe5\xee\xdf\xa1\xf6\x97\xd0\x24\x6b\x7f\x33\xfe\x1b\x00\x00\xff\xff\x7a\xf2\xa5\x0a\x77\x14\x00\x00")

func staticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticIndexHtml,
		"static/index.html",
	)
}

func staticIndexHtml() (*asset, error) {
	bytes, err := staticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/index.html", size: 5239, mode: os.FileMode(420), modTime: time.Unix(1523401164, 0)}
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
	"static/config/docker-compose.yml": staticConfigDockerComposeYml,
	"static/index.html": staticIndexHtml,
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
	"static": &bintree{nil, map[string]*bintree{
		"config": &bintree{nil, map[string]*bintree{
			"docker-compose.yml": &bintree{staticConfigDockerComposeYml, map[string]*bintree{}},
		}},
		"index.html": &bintree{staticIndexHtml, map[string]*bintree{}},
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

