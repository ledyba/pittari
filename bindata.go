// Code generated by go-bindata.
// sources:
// assets/static/css/main.css
// assets/templates/create-error.html
// assets/templates/index.html
// assets/templates/main.html
// assets/templates/upload-error.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
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

var _staticCssMainCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4f\x6f\xdb\xb6\x1b\x3e\x5b\x9f\x82\x40\x50\xe0\xf7\x33\x62\x59\xb6\xd1\x2c\x90\x4f\x41\xe2\x14\x01\xb2\xa4\xb0\x73\x58\x4f\x01\x25\xbe\xb2\x88\x50\x24\x41\x51\x8e\xdb\x21\x87\xda\xa7\x01\xbd\x0e\x3b\x6d\xc0\x86\xdd\x86\x7d\x83\x7d\x1b\x7d\x91\x81\xa2\x64\x53\xb6\xd3\xa5\xc5\x26\xc0\x68\xf3\x4a\x7c\xf8\x3c\xcf\xfb\x87\xec\x77\xd1\x1b\xe0\xa0\x30\x43\xdd\xbe\x97\xea\x8c\xa1\xef\xbd\x4e\x86\xd5\x9c\xf2\x10\x05\x63\xaf\xd3\xef\x4a\x4c\x08\xe5\x73\xf3\x27\x9a\x4d\xae\x27\xe7\x77\xe8\xe6\xf6\x0e\x5d\x5c\xcd\xde\x5e\x9f\xbd\x9b\x5c\xa0\xf3\xdb\xe9\x74\x72\x7e\x77\xfd\x0e\x5d\xdd\xa0\xcb\xab\xe9\xe4\xf2\xf6\x3b\x03\xf7\xe4\x79\x91\x20\xef\x0d\x62\x84\xe3\x87\xb9\x12\x05\x27\x21\x3a\xba\xac\x9e\xb1\x79\x8f\xcd\xcb\x58\x30\xa1\x42\x74\x14\x54\x8f\x8d\x87\xa9\x58\x80\x72\xdf\x26\xc9\x70\x38\x1c\xda\xb7\x47\x98\xc7\xa9\x50\xcf\x2c\x4e\x07\xe6\x45\x22\xb8\xee\xe5\xf4\x03\x84\x68\xb9\xec\x31\xac\xe6\x30\xae\xa3\x8f\x40\xe7\xa9\x0e\x51\x24\x18\x19\x3b\x72\xfd\x93\x6f\x20\x43\x81\x5c\x5a\x98\xe1\x2e\xcc\x0b\x51\x4e\x47\x2e\xca\x68\x07\xe5\x25\x18\x03\x3f\x70\x20\x28\x97\x85\x3e\xce\x81\x41\xac\x8f\x35\x2c\x35\x56\x80\xf7\x5c\x4d\xaa\x67\xec\x75\x22\xa1\x08\xa8\xde\x23\x25\x3a\x0d\xd1\xc0\x80\x34\xb1\x5c\xbf\x67\x10\x22\x22\xb4\x06\xb2\x0d\xef\x7a\xd8\x98\x1a\x31\x1c\x3f\x58\x15\x95\xd7\x6d\xe4\xa1\x5c\x1a\x8a\xcd\xef\x8b\x77\x79\xf2\x3c\x8d\x23\x06\x2f\x55\x32\xfc\x1a\x25\xdb\x38\xc3\x32\x87\x10\x35\xff\x73\xdc\x7e\x5d\xfb\xac\xd3\x3d\x2a\xbb\x38\x5f\x41\x65\xa3\xa6\x1d\xa8\x36\x24\xff\xa5\xf6\x27\xcf\xeb\x77\x51\xf9\xf1\xe7\x72\xf5\xa9\xfc\xf8\x4b\xb9\xfa\x64\x9a\xd2\xcf\x25\xe5\xdc\xb6\xd6\xa6\xb3\x2b\x07\x3a\x52\xe4\x54\x53\xc1\x43\x84\xa3\x5c\xb0\x42\x1b\x93\x94\x2d\x50\x8b\x47\xe8\xe2\x28\xc3\x94\x9b\xc5\x35\xb3\xd3\xc0\xe6\xde\x9a\xd9\x63\x90\xe8\x10\xe1\x42\x8b\x6d\xac\x86\xa8\x83\x96\xed\x86\x67\x2d\xc3\x2a\xac\xb7\x88\x05\xd7\xc0\xb5\x43\xb1\xc6\x1d\x40\x36\xde\xc6\x6a\xdc\x2a\x58\xaf\x4c\x01\x13\xab\x6d\xeb\xea\x7e\x26\x76\x74\xbf\xb4\xc5\x1b\x03\x80\x17\x11\x3e\x64\xa0\x03\x94\x67\x98\xb1\x6d\x6e\xb4\x90\xb6\x61\x6a\xb5\x6e\x92\x0c\x66\x22\x84\xae\xc7\x1d\x03\x6c\x5a\x4f\xe8\x74\x9f\xa9\x8b\xf6\xbc\x7f\x7e\x26\x48\x51\xf7\x55\x6d\xf6\xa1\xad\x9b\xf4\x44\x42\x6b\x91\x6d\xdb\x60\x8b\xd0\xdb\xda\xb9\x21\xb2\xbf\x4f\xcf\x49\xd7\x3f\x11\xdc\xc7\xe9\x77\xcb\xf5\x8f\xe5\xea\xd7\x72\xf5\x7b\xb9\xfe\xa3\x5c\xfd\x56\xae\xd7\xe5\xfa\xa7\x72\xfd\x67\xb9\xfe\xab\x5c\xff\xd0\xed\xdb\x29\xe1\x27\x42\x65\xc7\x5a\xf9\x52\x09\x79\xac\x89\xcf\x71\x06\xe6\xdf\x05\x66\x45\x4b\x2b\x17\x1c\x9a\xea\xff\x76\x32\x9b\x9d\xbd\x99\xcc\xd0\xd9\xcd\x05\x9a\x4c\xa7\xb7\xd3\x59\xd5\x04\x19\xe4\x39\x9e\x1f\x18\x3d\xa3\xe4\x34\x89\x51\xa1\xd8\xff\x7c\xbf\x4f\x33\x3c\x87\xbc\x9f\x3f\x50\xde\xa7\xdc\x10\xc0\xa6\x3f\x7c\xc9\xe7\xff\x47\xa7\x72\x89\x5e\x07\xaf\xbc\x4e\x87\x8b\x9e\x02\x09\x58\x3b\xd5\x3d\x90\x4b\x94\x0b\x46\x09\x3a\x8a\x86\x64\xd0\x1a\x00\x41\x70\x42\x22\xec\x8e\xfb\x6a\x84\x9a\x04\x54\x47\xae\x9b\xf4\xcd\x6f\x14\xc8\x65\xe3\x3b\x28\x25\x54\x7e\x68\x76\x8c\x92\xd1\x41\x12\xaa\x1a\x14\x0d\x81\x38\x6e\x55\xc0\xe7\x09\x6c\xc2\xed\xcd\x8b\xea\xa6\xc0\x68\xae\x9b\x79\x64\x7d\xef\x38\xd7\x85\x9d\x25\x8c\xee\x52\x3e\xe0\x33\x2c\x63\x86\x77\x7d\x0e\x5e\x21\xd7\x64\x46\xb9\xa9\xcd\xba\xfd\x4f\xdc\xc2\xaa\x07\xc5\xa8\x39\x3b\x35\x69\x76\xb7\xe7\xa7\x5b\x28\x3b\xf6\xb4\xbe\xae\xce\xdc\xcf\x7f\xdc\xef\xa2\xb7\xa9\xd0\xc2\x14\x14\xcd\xe6\x7e\x82\x63\xb8\xaf\xa4\x54\x0d\x73\x68\x98\x56\x7d\x61\xe7\xa5\xe5\x19\x38\xdd\x24\x45\x7e\xff\xa8\xb0\x94\x75\xc7\x6d\x00\x14\x30\xac\xe9\x02\xbe\x60\xca\x3a\x90\x56\xb7\xbd\x2c\x6d\xee\x28\xed\x5e\x74\x4f\x6f\x57\x69\x7d\xfe\x1f\x98\xa5\x8f\x29\xad\x04\x1d\x54\xf9\xa1\x47\x39\x81\xa5\x29\x2b\x7b\x9b\x28\x54\x6e\x16\x65\x62\x01\x2d\x6a\x19\xd4\x43\xef\x5f\xe3\x95\x53\xb6\x00\xf5\x0c\xb1\x27\xcf\xb3\x66\xf8\x44\xc4\x45\x06\x5c\xdf\x9b\x61\x7d\x6f\x8a\xd8\x39\xd1\x06\xc3\xd6\xc5\xcb\x97\x26\xc9\xd5\x97\xad\x63\xcf\xc9\xdc\xbc\xa0\x04\xaa\xdc\x81\x6c\x4f\x4b\xf7\x6c\xdc\xc8\x3b\xe4\xd9\xb3\x22\x9f\xbc\xbf\x03\x00\x00\xff\xff\xb6\x11\xd2\xd0\xa4\x0b\x00\x00")

func staticCssMainCssBytes() ([]byte, error) {
	return bindataRead(
		_staticCssMainCss,
		"static/css/main.css",
	)
}

func staticCssMainCss() (*asset, error) {
	bytes, err := staticCssMainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/css/main.css", size: 2980, mode: os.FileMode(420), modTime: time.Unix(1552212315, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCreateErrorHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xca\x4f\xa9\x54\xaa\xad\xe5\xb2\xc9\x30\xb4\x7b\xb2\x77\xce\xb3\x8e\x09\x8f\x9b\x56\x3c\x6e\x5e\xf9\xb8\x79\x8f\x8d\x7e\x86\xa1\x1d\x97\x4d\x4a\x66\x99\x42\x72\x4e\x62\x71\xb1\xad\x52\x6e\x6a\x71\x71\x62\x7a\xaa\x92\x5d\x75\xb5\x9e\x2f\x84\x5d\x5b\x6b\xa3\x9f\x92\x59\x66\xc7\x65\xa3\x9f\x96\x5f\x94\x6b\xc7\x55\x5d\x9d\x9a\x97\x52\x5b\xcb\x05\x08\x00\x00\xff\xff\x0a\x85\x3b\xea\x63\x00\x00\x00")

func templatesCreateErrorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesCreateErrorHtml,
		"templates/create-error.html",
	)
}

func templatesCreateErrorHtml() (*asset, error) {
	bytes, err := templatesCreateErrorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/create-error.html", size: 99, mode: os.FileMode(420), modTime: time.Unix(1552212315, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x56\xdd\x4f\x1b\xc7\x16\x7f\x86\xbf\x62\x35\x2f\xfb\x84\x17\x1c\x20\xf7\x72\xd7\x2b\xd9\xf7\x2a\x37\x52\x55\x29\x52\x1f\xf2\x58\xad\xbd\x83\xbd\xcd\x7e\x69\x3d\x0e\xa1\x08\xc9\x33\x0b\x8d\xf9\x70\xa9\x80\x7c\x90\xd0\x12\xda\x08\xbb\x40\x70\x42\x12\x15\x82\x81\x3f\xe6\x78\x6d\xf2\xc4\xbf\x50\xcd\xac\x0d\x0e\x10\xf7\xc9\xa3\x9d\x33\xe7\xfc\x7e\xe7\xe3\x77\x3c\x35\x65\xe0\x71\xd3\xc1\x12\x4a\xbb\xc6\x24\x9a\x9e\xee\x57\x73\x71\xad\xb5\x7a\x14\x06\x4b\xc0\x96\x81\x7e\x00\xfa\x07\xd0\x0d\x60\xf3\xc0\x3e\x02\x7b\x0d\xec\x13\xd0\x4a\x58\x7e\x1b\x96\xfe\x02\xfa\x0c\x58\x09\xe8\x4f\xaa\x92\x8b\x6b\xfd\xaa\xd7\xf5\x90\x1d\x41\xf0\x1e\x02\x0a\xc1\x26\x04\x75\x08\x4a\x10\xec\x84\xe5\xe7\x8d\xe3\x32\xd0\x4a\x73\xf1\x71\xb8\xf7\x82\x3f\xe7\x9e\x2f\xdc\xee\x74\xdc\xae\x01\x5b\x10\x57\x14\xe8\x9e\x67\x8c\x43\xf0\x04\xd8\x26\x37\x0b\x76\x80\x2d\x37\x8e\xd7\x05\xaa\x2d\xa0\x4b\xc0\x16\xb9\x31\xfb\x28\xa2\xac\x00\x3b\x04\x5a\xe1\x1e\x8a\x4c\x55\x3c\x01\x0a\xd8\x7b\x81\x65\x05\x82\x05\x60\x33\xe1\x9b\xad\xe6\xab\x4d\xa0\x7b\xe1\xec\xbb\xd6\x6a\x15\x82\x67\x10\x6c\x73\x03\x76\xca\x9f\x32\x06\x45\x7a\x09\xea\xcb\x1c\x74\x01\xdf\xba\xe0\xda\x41\x5d\x01\x5a\x06\x7a\xd2\x1d\x7c\xdc\xf5\x6d\xc9\xc6\x24\xe7\x1a\x09\xd9\x73\xf3\x44\x96\xf4\x0c\x31\x5d\x27\x21\xc7\x94\x82\x67\xb9\xba\x21\x4b\xd8\xc9\x90\x49\x0f\x27\x90\x5d\xb0\x88\xe9\xe9\x3e\x51\xf8\xbb\x01\x43\x27\x3a\xd2\xfa\xfb\x54\xa2\xa7\x2d\x2c\x99\x46\x42\x36\x9d\x71\x57\xd6\xfa\xfb\xfa\x54\xe2\xf3\x9f\x3e\x95\x18\xed\xa4\x9f\xd7\x4b\x43\x83\xdf\xa6\x04\x80\xca\x79\x7d\x4e\x55\x88\x71\x61\xa2\x9a\x8e\x57\x20\x92\x08\x23\x8f\x9b\x16\x96\x25\x47\xb7\x71\x42\x36\x6d\x3d\x8b\x65\x29\x6f\xfe\x88\x13\x72\x7c\x50\x96\x14\xad\xf3\x50\x55\xa2\x18\x5d\xa1\xa4\x3c\x99\xb4\x70\x02\x3d\xd4\x7d\x62\x66\x74\x6b\x40\xb7\xcc\xac\x33\x26\x11\xd7\xfb\x0f\xd2\xae\x25\x64\xad\x5d\x19\x91\xc9\x2f\xf0\x34\xab\xdb\xdd\x90\x90\x53\xb0\xd3\xd8\x47\xdd\xa0\xee\x9b\x06\xc9\x75\x90\x8d\xc8\x52\x9e\x60\x2f\x81\x06\x63\x83\x43\x48\x52\xb4\x8c\x2d\x3d\x92\x5a\x07\x5b\xff\xe4\xe5\x2e\x36\xb3\x39\xd2\xc3\x8d\x9a\xf6\xb5\xf3\xfa\x7c\xf8\x76\xa9\xf9\xe4\x6d\x8b\x1d\x02\x63\x8d\xe3\x53\xa0\xa5\xa8\x92\x3d\x92\xa1\xb5\x56\xab\xad\x0f\x6b\x37\x13\xe4\xbf\x7d\x6a\x1e\x5b\x38\x43\x22\x3c\xc8\xd3\x3d\xec\xa3\xf6\x8d\xeb\xf1\x2e\x90\x1e\xea\x56\x01\x27\x50\x72\x08\x69\xc9\x21\x55\x89\xbe\xde\x6c\x12\x47\x5a\x32\xde\xdb\xe4\x16\xd2\x92\xb7\x7a\x9b\x0c\x23\x29\x02\x85\x8d\x04\xea\x9c\x90\x96\x1c\xee\xfd\x6c\x04\x69\xc9\x91\xde\x26\xa3\x48\x4b\x8e\xf6\x34\x49\x0d\x21\x2d\xd5\x9b\x65\x2a\x8e\xb4\x54\x6f\x96\xa9\x5b\x48\x4b\xf5\x66\x99\x1a\x46\x5a\xaa\x37\xa3\xd4\x08\xd2\x52\xbd\x19\xa5\x46\x91\x96\xba\xca\x48\x89\x72\x16\xd5\xb9\x47\x6f\x7c\x7d\xfa\xf2\x85\xb4\x6d\x12\xb9\x1d\x45\xce\x63\xc7\xb8\x61\xea\x54\x45\x0c\xbd\xd6\xaf\x0a\x2d\xd0\x84\x2c\x0b\x51\x7d\x2a\xa4\xea\x35\x04\xbb\xed\xf3\xa5\x1a\x47\xb2\xb9\x03\xb4\xd6\x25\xc8\x1b\xcd\xd2\x2f\x40\x9f\x70\x9d\xa4\x1b\xf7\xfe\x77\x87\xdf\xea\x52\xce\xc7\xe3\x09\x94\x23\xc4\x1b\x53\x94\x89\x89\x89\x98\xe7\x9b\x0e\x31\x9d\x6c\xcc\xc1\xb1\x1f\x3c\x05\x7d\x2d\xd8\x1e\x04\x65\x08\x02\x21\xe8\x1d\xe1\x0c\x4a\xaa\xa2\x6b\x5c\xbf\x6b\x27\xe1\xe9\x7a\x24\x8f\x40\x67\xa2\x11\x3a\xaf\x97\xe2\x83\x43\xc3\xe1\xe1\x87\x78\x73\xbd\xd4\x5a\x3a\x09\xd7\xab\xe7\xf5\x39\x28\xb2\xc6\xd1\x53\xee\xb0\x5b\x9e\x85\x00\x5f\x75\xc3\x59\x2d\x70\x61\xe6\xfb\x66\x51\x78\x7d\x09\x6c\x05\xe8\xa2\x90\xea\x25\x71\x3b\x7f\xf9\xbd\x4b\xfa\x83\x22\xb0\x2d\xe1\xfe\x8d\x58\x10\x73\xd1\xd2\xe1\x79\x60\xcb\xad\xb5\xa3\xcf\x8b\xfb\x9d\x38\xf3\xe1\xa7\x67\xd7\x92\x59\x85\x22\xfd\xca\xb6\xaa\xf4\x5e\x00\x50\x5c\x0c\x7f\xde\x08\x6b\x07\xe1\xe3\x4f\xe1\xfc\x4b\x28\x96\xb9\xe5\xe9\x2c\xd0\x17\x9f\xe9\x81\x80\x5f\x69\x1c\x2c\xf0\xd2\xd0\x99\x0b\xc8\xb9\xb8\x16\x2e\xb1\xb3\x62\xd0\x25\x2b\xbc\x94\x9c\x8c\x9a\x27\xbe\xeb\x64\xb5\x2f\xca\x97\x1f\x53\x94\x47\xce\xc0\x40\xc1\x99\x1c\xfc\xf7\xed\x4c\x2c\x63\xfa\x19\x0b\xc7\xec\xbc\x92\x23\xb6\xa5\x64\x0a\x84\x60\xdb\x53\x32\xae\x6d\x3e\xc0\xdf\xf3\x73\x8c\x5f\xa0\x68\x2b\x6e\x00\x7b\x27\x4a\x20\xd6\x27\xab\x89\xfd\xba\x13\xd5\x97\xd7\x54\x55\xda\x31\x81\xd6\x78\x26\xaa\xdb\xc3\x19\x9b\xeb\xee\x48\xec\x76\xc6\x8e\x36\xed\x79\xbd\xf4\xdf\x7f\x8d\x8a\xf3\x5e\x33\x98\x0d\x5f\xbd\x8b\xaa\x2b\x08\x45\x8c\x66\xcb\x67\xfb\xbb\x11\x11\xa0\x2b\x97\x01\xdb\xfb\xba\xd6\x9a\xd9\x6c\x9c\xfe\x7a\x56\xad\x5f\xec\xee\x1b\xcc\xf8\xde\x3f\x8d\xf6\x7e\x6b\x75\x03\xe8\x73\xa0\x1b\x8d\x83\x72\x58\xfa\xb3\xf5\x72\x3e\x6a\xbe\xa8\x94\x67\xfb\xbb\x8d\xa3\x23\x60\xcb\x67\xfb\xaf\x3a\x5d\xd8\xee\x0b\x01\xe7\x9e\x3b\x81\x7d\x43\x4a\x4f\x5e\x4c\x49\x27\xda\xef\xed\x9e\xa6\x35\xd5\xc0\x96\x76\x75\x4e\xf2\x8e\xee\x8d\xfb\xba\x8d\x27\x5c\xff\x41\x2c\xe3\xda\x0a\xd2\xbe\x73\x74\x6f\x4c\x4a\x4a\x77\xf5\xfc\x03\x6c\x59\xd2\x7d\x9c\x96\xee\x74\x6c\xa2\xb1\xa8\x34\x2b\x0b\xad\xda\x5c\x67\x08\x2f\x27\x43\x24\xe9\x5a\x20\x5e\xd1\xac\x6b\xe9\x4e\x36\xe6\xfa\x59\x05\x69\xff\x77\xcf\xaa\xc5\xb3\xed\xdf\xda\x43\xd6\xaa\x1c\x35\x5f\x97\x05\xd5\x93\x76\x53\x76\xff\xd7\x39\x16\x95\x3c\x14\xe5\x8d\x1a\xbe\x76\xdd\xbd\x49\x72\x85\xb4\x60\x60\x61\x63\x32\xad\x2b\xf7\x4c\x42\x74\xdf\x1c\xf8\xa6\xe0\x28\x48\x8b\xee\xf9\x30\xcc\xee\x7e\x7e\xba\x70\x65\xa0\x05\x8e\x76\xc8\xa9\x29\xec\x18\xd3\xd3\xfd\x7f\x07\x00\x00\xff\xff\xd6\xd7\x87\x93\x48\x0a\x00\x00")

func templatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexHtml,
		"templates/index.html",
	)
}

func templatesIndexHtml() (*asset, error) {
	bytes, err := templatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.html", size: 2632, mode: os.FileMode(420), modTime: time.Unix(1552212449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMainHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcf\x8a\xd4\x40\x10\xc6\xef\xf3\x14\x65\x9d\xf4\x30\xd3\x3b\x8b\xe0\x22\x9d\x5c\x7c\x05\x5f\xa0\x37\xa9\x6c\x5a\x7b\x92\x31\x5d\x2e\x0c\x61\x0e\x49\x2e\x51\x10\x11\x61\x3c\x0a\x2a\xa2\x07\xc5\xbf\x08\xb2\xe8\x23\xf8\x10\x4d\x58\xf7\xb4\xaf\x20\x9d\xde\x65\xc6\x8b\x78\x4a\xf2\x15\xf5\xd5\xaf\xbe\x54\x5d\xa7\x94\xe9\x82\x00\x0f\x95\x25\x5c\xaf\x27\xf2\x4a\x5a\x26\xbc\x5a\x12\xe4\xbc\x30\xf1\x44\xfa\x07\x18\x55\x1c\x45\x78\x47\xa1\x17\x48\xa5\xf1\x04\x40\x2e\x88\x15\x24\xb9\xaa\x2c\x71\x84\xf7\x39\x9b\x1e\xa0\x2f\xd4\xb5\xce\x80\xee\xc1\xec\xb6\x66\x43\x80\xde\x16\x40\xb2\xff\x8a\x5d\xf3\xc5\x35\x2f\x5d\xf3\xdc\xb5\x0f\x87\x47\x1f\x86\xfe\x9b\x6b\x1e\xbb\xf6\xa9\x14\xa1\x3e\xf6\x93\xb1\xb4\xdb\x54\xd7\xc1\x6b\xbd\x86\x29\xfc\x9f\x43\x91\x06\x03\xa3\x8b\xbb\x50\x91\x89\xd0\xf2\xca\x90\xcd\x89\x18\x21\xaf\x28\x8b\x70\x26\x2c\x2b\xd6\x89\x48\xac\x15\x0b\xa5\x8b\x59\x62\x2d\x82\x08\x16\x4c\x8b\xa5\x51\x4c\x80\x7e\x63\x84\x99\x4f\x47\x84\xed\xe5\x61\x99\xae\xc6\x10\x52\x7d\x0c\x3a\x8d\xd0\xb7\x8f\xdb\xef\x68\x49\x59\x30\x15\x7c\x21\x03\xc8\x7c\x1e\x4b\x05\x89\x51\xd6\x46\x38\xd2\x6e\x51\xf0\x1f\xd1\xa8\x58\x8a\x7c\x7e\x69\xb3\x8b\xe6\x39\x02\xda\x38\x40\xa4\xfa\xf8\x2f\xac\xac\x2c\x99\x2a\x8c\xaf\xde\xba\xb6\xbf\x37\xbf\x3e\xdd\xdf\x9b\x1f\x80\x54\x17\x53\x73\xe6\xa5\xbd\x29\xc4\x0d\x5d\xce\xca\xea\x48\x60\xfc\xab\x3f\x3f\xe9\x5d\xf7\xcc\xb5\x5f\x5d\xfb\xea\xfc\xe4\x81\x6b\xde\xff\xee\xfb\xe1\xc9\xe7\xb3\xcd\x8b\xe1\x67\x77\xfa\xf6\xfb\xd9\x66\x24\x02\x01\xc3\xeb\xcd\xe9\xa7\x1f\xae\x79\xe3\xba\x77\xae\xfd\xe8\xba\xce\x75\xfd\x96\x21\xbc\x48\x11\xb2\x92\x22\x1c\xd4\xe5\xbf\xd9\x9e\xde\x18\xef\xa8\x84\xca\x9f\x00\x00\x00\xff\xff\xeb\x02\x7a\xdb\x95\x02\x00\x00")

func templatesMainHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMainHtml,
		"templates/main.html",
	)
}

func templatesMainHtml() (*asset, error) {
	bytes, err := templatesMainHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/main.html", size: 661, mode: os.FileMode(420), modTime: time.Unix(1552212315, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUploadErrorHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xc1\x8a\xd5\x30\x14\x86\xf7\x7d\x8a\xc3\xd9\x3b\x61\xf6\x69\xdf\xc0\x87\x48\x6f\x72\xef\x0d\xa4\x49\x68\xd2\x0b\x97\x52\x98\x36\x30\x20\x3e\x80\x0c\x08\xa2\x0b\xd1\x85\xae\x45\x06\x7c\x98\xcc\x20\xbe\x85\x24\x99\xea\x2d\x22\xcc\x2e\xc9\x7f\xfe\xf3\x9f\x7e\xa7\xe3\xc8\xc5\x5e\x6a\x01\xd8\x1a\x7e\xc6\x69\xaa\xe8\xf1\xba\x79\xbc\xbd\xfb\xf9\xf6\x5d\x9c\xbf\xc6\xe5\x43\x0c\x21\x86\x37\x31\x7c\x89\xe1\x3e\x86\x57\x71\xf9\x14\xc3\xe7\x18\xee\x29\x39\x5e\x37\x15\xe5\xf2\x04\x3b\xc5\x9c\xab\xb1\x13\xce\xb1\x83\xc0\x66\x1c\xaf\x5e\x96\xf3\x34\x51\xc2\xe5\xa9\xc9\x5d\xe3\xb2\xc4\xf9\xf6\xe1\xdb\xcd\xe3\xf7\x8f\xbf\x6e\xe6\x87\x1f\xef\xe3\x7c\x17\x97\xd7\x4f\x9d\xf6\xa6\xef\x80\xed\xbc\x34\xba\xc6\x2b\x32\x58\x65\x18\x47\xe8\x84\x3f\x1a\x5e\xa3\x35\xce\x23\x08\xbd\xf3\x67\x2b\x6a\xec\x06\xe5\xa5\x65\xbd\x27\xc9\xf7\x82\x33\xcf\xb0\xa9\x00\x2e\x07\xe2\x92\x29\x73\xc8\xcf\x00\xd4\xb3\x56\x89\x55\x4a\x26\x04\x37\x74\x1d\xeb\xcf\x35\x96\x30\xc8\xaf\xa5\x3c\x19\x12\x92\xf5\x96\xee\xfd\xea\xb6\xbd\xb1\xf8\x57\x49\x1a\x87\x13\x53\xf2\xa0\x6b\xf4\xc6\xe2\x5a\xa8\x59\x27\x36\x85\x00\x54\xb1\x56\xa8\x14\x55\x63\x99\xba\xd0\xa6\x24\x0b\x9b\xae\xc4\xf3\xe7\xa4\x9c\x98\x1a\xfe\x89\x91\xda\x0e\x1e\x0a\xac\xbd\x54\x02\x41\xf2\xa7\x44\x48\x63\xad\x67\xf2\xff\x44\x4a\x7c\xff\x87\x06\xb9\xc0\x41\x49\x86\x99\x79\x97\xfd\x6e\xc1\xb7\x83\xf7\x46\xbb\x95\xbc\xb3\x4c\x6f\x25\x6c\x36\xf3\xb9\xa1\xed\xa4\x47\xc8\x1f\x52\x63\xf9\x39\xb0\xa1\x24\x39\x2f\x52\x68\x5e\x76\x53\x8d\xa3\xd0\x7c\x9a\xaa\xdf\x01\x00\x00\xff\xff\xde\xec\x8e\x94\xbe\x02\x00\x00")

func templatesUploadErrorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesUploadErrorHtml,
		"templates/upload-error.html",
	)
}

func templatesUploadErrorHtml() (*asset, error) {
	bytes, err := templatesUploadErrorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/upload-error.html", size: 702, mode: os.FileMode(420), modTime: time.Unix(1552212315, 0)}
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
	"static/css/main.css":         staticCssMainCss,
	"templates/create-error.html": templatesCreateErrorHtml,
	"templates/index.html":        templatesIndexHtml,
	"templates/main.html":         templatesMainHtml,
	"templates/upload-error.html": templatesUploadErrorHtml,
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
		"css": &bintree{nil, map[string]*bintree{
			"main.css": &bintree{staticCssMainCss, map[string]*bintree{}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"create-error.html": &bintree{templatesCreateErrorHtml, map[string]*bintree{}},
		"index.html":        &bintree{templatesIndexHtml, map[string]*bintree{}},
		"main.html":         &bintree{templatesMainHtml, map[string]*bintree{}},
		"upload-error.html": &bintree{templatesUploadErrorHtml, map[string]*bintree{}},
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

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
