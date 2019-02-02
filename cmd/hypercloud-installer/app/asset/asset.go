// Code generated by go-bindata.
// sources:
// app/static/config/docker-compose.yml
// app/static/index.html
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

var _staticConfigDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\xcd\x6e\xe3\x38\x0c\xbe\xe7\x29\x04\x63\x80\x5c\x26\x8e\x3b\xed\x60\x0b\x01\x3e\x64\xd2\x4c\x27\x80\x9b\x0c\x92\x74\x17\x7b\x32\x14\x8b\x49\x84\xca\x92\x21\xc9\x49\x83\xc5\xbe\xfb\x42\xfe\x77\xed\x6d\xa7\xed\xf6\xb6\x17\x03\x22\x3f\x52\x34\xf9\x91\xe2\x11\x94\x66\x52\x60\xe4\x5c\xba\x57\xce\x40\x83\x3a\xb2\x08\x34\x1e\x20\x64\x14\x81\x1d\x7b\x08\x99\x60\xc6\x9e\x11\x62\x31\xd9\x03\x2e\x15\xf8\xc2\xfd\x2d\x13\x47\x32\x8e\x89\xa0\x39\x06\xa1\x11\x72\xb4\x91\x0a\x22\x29\x76\x6c\xef\xd4\xd2\xd1\x88\x24\xac\x75\xa6\xb0\x23\x29\x37\x20\x8c\x3a\x27\x92\x09\xa3\xfd\x83\x31\xc9\x67\xfb\xd1\x2d\x64\x06\xf9\x99\x43\x16\x24\x06\x6c\x21\x68\x42\xa9\x02\xad\x31\xbe\xf6\x5e\x46\xeb\x1a\x7e\x75\x75\x89\x36\xc1\xba\x1d\x5b\x14\x43\x47\xe0\xd6\x9e\xfc\x6e\x50\x19\xc2\x70\x3d\x3d\x10\xce\x41\xec\xfb\xec\x63\xc2\xb8\xff\x69\x76\x37\x99\x07\x5d\xad\xcd\x13\xd9\x83\x5f\x64\x74\x6c\x85\x63\x12\x45\x32\x15\xa6\x8b\x5e\x8a\x1f\x52\x9b\x55\xca\xad\x41\xda\x73\x99\x14\x37\x60\x2b\xe1\xef\x08\xd7\x6d\x3d\x95\xd1\x03\xa8\x1e\x91\xab\x4f\x44\xc5\xb1\xa4\x7d\x78\x17\x1e\x13\xa9\x81\x6e\xcf\x45\xa5\xba\x17\x17\x40\x2a\x63\xc2\x84\xff\xe9\x66\x79\x37\x99\x2f\xfa\x10\x27\x62\xa2\x43\x4b\xc1\xe5\x3e\x80\x23\x70\xff\x66\xf6\xed\xfe\xb6\xa5\x8a\xa4\xd0\x29\xef\x11\xb9\x20\x68\xc6\x14\x3f\x3f\xe3\xeb\xaf\x9e\xd7\x87\x4b\x14\xec\xd8\x63\x99\xd9\x1c\x21\xc0\x9c\xa4\x7a\xd0\x35\x53\x0f\xe7\x04\x54\xc4\x65\x4a\x33\x11\x85\x84\xcb\x73\xa9\xe6\x64\x0b\xbc\x02\x67\x3c\x77\x6b\x03\x97\x09\x03\x4a\x10\x8e\xd1\xd0\x66\x65\x58\xe0\x14\x68\x43\x94\x09\x13\xc9\x59\x74\x6e\x5a\x0b\xca\x4c\xd6\x6b\x52\x8c\x76\x84\xf1\x54\x41\x79\x2b\x08\xaa\x43\x29\xea\xc0\xf2\xbf\x18\xd4\x6d\xf8\x6c\x07\xf6\x79\x68\xb6\xef\x53\xb7\xf6\x70\x94\x3c\x8d\xa1\x91\x0b\x73\x4e\x00\xa3\x2d\x13\xb4\x8a\x59\xcb\x54\x45\x80\xd1\xf8\x48\xd4\x58\xa5\x62\x5c\x72\x46\x46\x0f\x15\xc8\x10\xb5\x07\xf3\x0c\x28\x91\xca\x34\xef\x29\xf0\xd7\x5e\xe5\x22\x49\xb7\x9c\xe9\x03\xd0\x96\xd4\x92\x12\xa3\x83\xd4\xa6\x63\x7b\x75\x75\xd9\x67\xdc\x14\x3f\xb1\xee\x0e\xa9\x8f\x62\xd9\xeb\x1a\xfc\x95\x9c\xcc\xff\x6a\xcf\xe5\x96\xf0\xf7\xd2\x34\x2b\x6e\xc3\x90\xb3\x98\x99\xc6\x19\xa1\x18\x62\xa9\xce\x18\x5d\x78\xde\x5d\x25\x56\x60\x5f\x09\x62\xa9\xdc\x0b\xfe\x5a\x61\x5f\xdf\x0a\xf9\x0f\x73\x62\xdd\xe8\xba\x94\xe4\x31\x24\xc6\x40\x9c\x18\x8d\xd1\x97\x9a\x21\x27\x26\xa8\x3c\x61\x74\x79\xed\x95\xe8\x84\x93\x08\x62\x10\xa6\x75\x9d\x36\x8a\xd8\xd7\xa0\x19\xef\x08\x0d\x85\xa4\xe0\x2a\xc9\x01\xf9\x3e\x8a\x89\x20\x7b\x50\x65\x82\xd2\x84\x12\x03\x61\xfe\x8a\xd5\x76\x09\x51\x76\xd8\x73\xa6\x63\x8c\x2e\x9e\x06\x7d\xe1\x69\xdb\xb3\x05\x6b\x9a\x2d\xdb\xe8\xbc\x92\x8a\x88\xec\x41\x18\x34\xb2\xe9\x04\x85\x46\x5b\x29\x8d\x0d\x34\x19\xc1\x63\x02\x91\xf1\x2f\xfa\x1b\x35\x77\x35\xa2\xc4\x10\x3c\xce\x0f\x63\x7b\xc8\x00\x20\x8e\x4c\x49\xd1\xcc\xc0\x08\x4d\x97\x8b\xf5\x7d\x10\x06\xcb\xe9\x24\x08\xa7\xcb\xc5\xf7\xf9\xad\xff\x97\xa3\x1f\x58\x12\x72\x20\x47\x08\xa5\x08\x33\x9e\xa8\x34\x31\x0e\xb6\x3c\xf9\xec\xe4\x61\xe5\xa7\xbf\x9f\xba\xfa\x36\x5f\xdc\x84\xf3\xc5\x66\xb6\xfa\x3e\x99\xce\x7c\x30\x07\xef\x29\x64\x1a\xcc\x67\x8b\x4d\x1f\xa8\x4d\x69\x05\x09\x67\x11\xd1\x75\x3a\xdf\xc6\xe8\xd7\x54\xbe\xaf\xf0\xef\x20\xed\x9b\xb9\x52\x5d\xf9\xc1\xbd\xf8\xec\x94\x19\x58\x3f\x7b\xa6\x8d\x3a\xb7\x28\x5b\x09\xbf\xf4\xf3\xb0\xd4\x17\x4c\xb4\xe3\x9f\xb3\xed\xb8\x14\xff\x3b\x1d\x57\xb3\xdb\xf9\x7a\xb3\xfa\x33\x9c\xdc\x6f\x7e\xf8\xce\xc1\x24\x44\xeb\x13\x75\xfa\x01\xe1\x8f\xcd\xcf\xc9\x7a\xfd\xc7\x4d\xf8\x73\x62\xe1\xd9\x1b\xa3\x21\x52\x60\xf4\x98\x09\x0a\x8f\x6e\xb1\x76\xb8\xbf\xec\x69\x35\x9b\x04\x77\xbe\xb3\x2a\x42\x45\x2b\x20\x3c\xee\xb1\x5a\x6f\x96\xab\xc9\xed\xcc\x77\x76\x8c\x83\x3e\x6b\x03\x7d\xa8\x60\x79\x1b\x06\xb3\xdf\x67\x81\xef\x50\xd8\xa6\xc5\xce\xfb\x31\x2c\xaf\x16\x02\xd7\x3e\xab\x18\x0d\xbf\x7a\x9e\xd7\x55\xee\x94\x14\x06\x04\x6d\x2c\xb0\x1a\xa3\x61\xb6\xc3\x3e\x83\x56\x29\x07\x8c\x86\x76\xcf\xc4\xad\xd4\x76\x6d\x40\x90\x6d\x06\xee\x8f\xae\x58\x02\x0a\xe2\xd9\xab\xab\x5f\xfb\xbf\x61\x5f\x6e\xd8\x82\xdf\xdd\x66\xab\x38\xfe\x0b\x7d\xdd\xd4\xd6\xf2\x1c\x4b\x15\x3b\x82\xc2\x48\x1e\x41\x71\x92\x37\x2b\x31\x86\x44\x87\xbc\xac\xb6\xaa\x83\x2e\x4d\x5f\x22\xe9\xa0\x31\x25\x9a\xef\x54\x61\x5e\xde\xca\x65\xf4\x8e\xfd\xa5\x3d\x78\xfe\x43\xd7\x83\x46\xda\x3b\x09\xc7\x6f\x48\x87\xc5\xd8\xd1\x81\x91\x5b\xcd\xc7\x1a\x5f\x4d\xb1\x6e\x71\xff\x09\x00\x00\xff\xff\x71\x29\x4d\xc0\x98\x0f\x00\x00")

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

	info := bindataFileInfo{name: "static/config/docker-compose.yml", size: 3992, mode: os.FileMode(420), modTime: time.Unix(1548716490, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x5b\x73\xdb\xb6\x12\x7e\xcf\xaf\xd8\x83\x39\x13\xca\x73\x4c\x32\x39\xbd\x4c\x9b\x88\x1a\x37\xb5\x33\x4d\x27\xad\x33\x75\xfa\xd4\xa4\x1d\x88\x58\x49\x68\x40\x80\x03\x2c\xe5\x28\x9d\xfc\xf7\x0e\x00\x52\xd6\x85\x94\xe5\xd4\x0f\x0d\x5f\x08\x82\xdf\x7e\xbb\xfb\x61\x71\x21\xc7\xff\x11\xa6\xa4\x55\x8d\xb0\xa0\x4a\x4d\x1e\x8c\xfd\x0d\x14\xd7\xf3\x82\xa1\x66\x93\x07\x00\x00\xe3\x05\x72\x11\x9b\xe1\xb1\x42\xe2\x50\x2e\xb8\x75\x48\x05\x6b\x68\x96\x7e\xc3\x20\xdf\x00\x90\x24\x85\x93\x1f\x56\x35\xda\x52\x99\x46\xc0\x0b\xed\x88\x2b\x85\x76\x9c\xc7\x77\x37\xd8\x29\x77\x08\x0b\x8b\xb3\x82\xe5\x81\x65\xc7\x8f\xe6\x15\x16\x6c\x29\xf1\xba\x36\x96\x18\x94\x46\x13\x6a\x2a\xd8\xb5\x14\xb4\x28\x04\x2e\x65\x89\x69\x78\x38\x05\xa9\x25\x49\xae\x52\x57\x72\x85\xc5\xe3\xed\xa8\x94\xd4\xef\xc0\xa2\x2a\x98\x2c\x8d\x66\xe0\xd3\x2e\x98\xac\xf8\x1c\xf3\xf7\x69\xec\x8b\x81\xcc\xf8\xd2\x3f\x66\xb2\x34\x3b\x21\x05\x8e\x08\x5a\x10\xd5\xee\x49\x9e\x37\xba\x7e\x37\xcf\x4a\x53\xe5\x67\xa5\xb2\x79\x23\xcf\x1e\x65\x8f\x1f\x67\xff\x4f\x6b\x4e\xe5\x22\x2f\x95\x4d\x1b\x99\x55\x52\x67\xa5\x73\x2c\x06\xe0\x68\xa5\xd0\x2d\x10\xa9\x27\xc4\x83\xf4\x3e\x2e\xb7\xef\x21\x74\x1f\xe9\xc4\x95\x56\xd6\x04\xce\x96\x9f\xee\xe5\x4f\xc7\x26\xe3\x3c\x32\xb5\x45\x92\xdf\x54\xc9\x78\x6a\xc4\x6a\xc3\xa3\x90\x4b\x28\x15\x77\xae\x60\x15\x97\x3a\xf5\x63\xc8\xa5\x46\xcb\x6e\x40\xbb\x40\xcf\xb6\x07\xd8\x05\x4d\x2d\xd7\x42\xea\x79\x0f\x2c\xa6\x5a\x73\xdd\x61\x43\xe1\xb1\x81\xaa\xf4\xc0\x1e\x57\xb9\x90\xcb\x9d\x10\x63\xd7\x60\xd8\x6d\x79\x0e\xa6\x38\x84\xe7\x16\xf9\x50\x16\x1b\x70\x6b\xae\xd9\x8e\xf7\x61\x66\x95\xbe\x77\xe9\xd7\x60\x66\x33\x87\xe4\xdb\x5f\x0c\x78\xe8\x35\xe7\x56\xdc\x02\xef\x33\x49\x07\xc7\xad\xef\x6a\x47\x80\x93\x34\xfa\x76\x57\xfb\xa3\x71\x54\x44\x53\x65\xca\x77\x47\x06\x34\x9e\x19\x5b\x41\x85\xb4\x30\xa2\x60\xb5\x71\xc4\x40\x8a\x82\xc9\xae\x54\x8e\xe4\xd9\x0d\xc4\xd3\xa6\x73\x6b\x9a\xfa\x0e\x04\x81\x44\xf1\x29\x2a\x98\x19\x5b\x30\x61\xfc\xec\x61\x93\xf3\x70\x07\x33\x83\x95\x69\x2c\xbc\xe2\xfc\x6a\x9c\x07\xdc\x1d\xb9\xa5\xae\x1b\x6a\xd7\x41\xc2\xf7\x6d\xae\xad\x1b\xa8\x15\x2f\x71\x61\x94\xc0\xb5\xef\x8c\x94\x60\x50\x73\x22\xb4\xba\x60\xbf\xff\xc6\xd3\x0f\x8f\xd2\x6f\xdf\xb6\xf7\x37\xe9\x9b\xec\xed\xff\x18\x38\xf9\x01\x0b\xf6\xe5\x57\xac\x5d\xbd\x3b\xc6\xfc\x0e\xe2\xf5\xcc\xb3\x83\xf8\xfb\x16\x1b\x2b\x2e\x15\x9b\x5c\xf8\x9b\xef\x81\x97\x48\x89\x83\x0b\x5d\xda\x55\x4d\xff\x5c\xef\xc8\x1f\x04\x6f\x9b\x5b\x7a\x3b\xc2\x25\x9e\x6d\xaa\xbe\x2b\x6a\x6b\xf5\x39\x69\x6a\x71\x2e\x1d\xd9\x55\xda\x38\x3f\x93\x7e\x75\x68\x83\xb4\xbf\xb4\xfd\xf7\x59\xc5\xdb\xbe\x7a\xc4\xdd\x57\xb4\x33\xf9\x23\x9a\x7c\x96\xca\xd6\xdc\xb9\x6b\xe3\x97\xee\x57\x6d\xeb\x9e\x15\x5e\x3b\xd8\x56\xf9\xa6\x7b\x50\xd5\x1b\xc8\xdd\x95\x3d\x0e\xea\xa5\x3c\x62\x83\xf8\xc4\x7d\x64\x66\x0c\x1d\xbd\x01\x04\x6b\xaf\x50\x6d\xcd\xdc\xa2\x3f\x94\xb5\x5c\x5d\x07\x28\x63\x6a\x06\xe1\x94\x56\x30\x21\x5d\xad\xf8\xea\x09\x68\xa3\xf1\xe9\x5d\x76\x99\x8e\x6f\x32\xce\xd7\xcd\x23\xf5\x3a\x5e\xda\x69\x43\x64\x34\x08\xe9\xf8\x54\xa1\x08\x99\x19\xdd\x6e\xdf\xdd\x51\xda\x35\xd3\x4a\x12\xf3\x05\x57\x6d\x6e\x99\xeb\xd3\x1a\x69\x98\x92\x4e\x6b\x2b\x2b\x6e\x57\x6c\xd2\xda\x8f\xf3\x48\x7f\x2f\x43\x77\x0b\xe4\xd0\x4c\x1d\x30\x1d\x3e\x05\x0e\x3d\x6e\x9e\x8b\xfd\x25\x4c\xd9\x54\xa8\x29\xe3\x42\x5c\x2c\x51\xd3\x4b\xe9\x08\x35\xda\x51\x72\x7e\xf9\xd3\xf7\xf1\x08\xf8\xd2\x70\x81\x22\x39\x85\xd1\x09\x14\x13\xf8\x6b\xcb\xa1\x3f\x78\x1b\x85\x99\x32\xf3\x1e\x9b\x93\xa7\xdb\x09\x2d\xb9\x85\xff\xa2\x77\xe4\xa0\x00\x8d\xd7\x10\xbc\x5e\x99\xc6\x96\x38\x4a\xf2\xf5\xd0\xe4\x11\xe4\x09\xf6\xec\x5b\xd0\x33\xd2\x50\xdc\x64\x30\x47\xba\x50\xe8\x9b\xcf\x56\x2f\xc4\x28\x59\x17\xc1\x21\x0e\xb4\xcf\xfd\xb1\xea\x00\xcd\x1a\xd8\x4b\x13\xb7\xc1\xe7\x12\x95\x38\x44\x12\x61\xbd\x0c\x61\xb3\xbc\x95\x20\xa0\x7a\xed\xd7\x93\xf6\x80\x75\x87\xd9\x1b\x8e\x0d\x29\x33\x8b\x95\x59\xe2\x77\x44\x56\x4e\x1b\xc2\x51\xd2\x4d\xa9\x7d\xab\x38\x36\x3d\x35\x53\xa1\x73\x7c\x8e\xbe\x54\xda\x66\x4f\xc5\x74\xa1\x07\x16\x28\xe0\xc7\xab\xcb\x9f\xb3\xda\x7f\xb3\x77\x46\x99\xe0\xc4\x77\xbd\xee\x15\x5b\xb0\x4f\x4e\x23\x4f\x1f\x5a\xce\x60\x14\x5e\x66\xe1\x3f\x42\x51\x14\xd0\x0d\x67\x02\x0f\x1f\x46\xc3\x8c\x97\xfe\x8c\x1f\xdf\xce\xa4\x96\x6e\xe1\x53\xee\x89\x39\xa4\xde\x49\x99\x85\xd5\x31\x6b\x17\x47\x28\x20\xf1\xcb\x63\xf2\xb4\xdf\xea\x93\x64\xee\xae\x3c\x3f\x7b\x7d\x79\x7e\x09\x16\x85\xb4\x58\x12\x90\x81\x8a\x6b\x3e\x47\x0b\x8d\x55\x7b\x36\x1f\x0f\x29\xb1\x95\x2d\x97\xea\x5f\x9a\x6b\xe7\x0b\xad\x35\xb6\x27\xc3\xad\xa7\xa1\xaa\x8e\x93\xbb\xa7\x4c\xe3\x76\xe0\xab\x34\xd6\x4e\x7f\x8d\x46\xc1\x6a\x1b\xee\xe7\x38\xe3\x8d\xa2\xd1\xad\x55\x79\x15\xb8\xc3\x4e\xd3\x9b\xe5\xa6\x3e\x0e\xa9\x4f\x9c\x53\xd8\x12\xea\x0e\xe3\x12\x3e\x28\x93\x1e\xaf\x33\xa4\x72\xb1\xb9\xbe\x26\xa7\x03\xa3\x1e\xbf\x30\x9f\x40\xf2\xea\xf2\xea\x75\x72\xda\x8b\x99\x1a\xe1\xcf\x03\x78\x0d\x5e\xe0\x73\x4e\x7c\xb4\xad\xf9\xc9\xbe\xdd\xc7\x93\xac\xe4\x3e\x08\x8b\xdc\xf9\x0a\xec\x93\x7c\x53\x4d\xb4\xb6\xc5\xf6\x68\x00\xc7\xd7\xca\x8e\xf5\x56\xb5\x6c\xbe\xdc\xfd\x69\x14\xff\x14\x8d\xf3\xf8\xfb\xf1\xef\x00\x00\x00\xff\xff\xee\x72\x4e\x60\x8f\x14\x00\x00")

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

	info := bindataFileInfo{name: "static/index.html", size: 5263, mode: os.FileMode(420), modTime: time.Unix(1548716490, 0)}
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

