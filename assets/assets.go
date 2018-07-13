// Code generated by go-bindata.
// sources:
// plugins/codeamp/graphql/schema.graphql
// plugins/codeamp/graphql/static/index.html
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

var _pluginsCodeampGraphqlSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\xcd\x72\xdb\x38\x12\xbe\xf3\x29\xe0\xca\x45\xa9\xf2\x13\xe8\x18\x3b\xbb\xeb\xdd\x64\xc7\x63\x4d\x4e\x2e\x1f\x60\x1a\x92\x30\x21\x01\x06\x00\x35\x51\x4d\xe5\xdd\xa7\xf0\x43\xa0\x1b\x04\x28\xc9\xce\x54\xcd\xc5\x26\x9a\xc4\xd7\x3f\xe8\x6e\x74\xb7\x74\x4b\x3b\xaa\xc8\x6f\xbc\x67\x4d\x78\xfe\xef\xe6\x97\xff\x37\x8d\x6e\xf7\xac\xa7\xe4\xcf\x86\x90\x6f\x23\x53\xc7\x35\xf9\xd5\xfe\x6b\x08\xe9\x47\x43\x0d\x97\x62\x4d\x3e\x87\xa7\xe6\x47\xd3\xbc\x0b\xef\xcd\x71\x60\xfe\xd1\xed\x7d\x47\xbe\x68\xa6\x1a\x42\x46\xcd\xd4\x8a\xbf\xac\xc9\xdd\xed\xfb\xf5\x44\xf4\x6f\x75\x78\xad\x57\xef\xd7\xe4\xd1\x52\x9e\xae\xdc\xcb\x7b\x25\x7f\x67\xad\x69\x08\x19\xfc\x53\x00\xb8\x26\xba\x1b\x77\x6b\xb2\x31\x8a\x8b\xdd\x35\x11\xb4\x67\x69\xc5\xc4\x81\x2b\x29\x7a\x26\xcc\xdd\xed\x44\x7e\xbf\x06\x68\x11\x59\x27\x68\xbd\x0a\x0f\x1b\x46\x55\xbb\x8f\x9f\xfb\xe5\x9d\x18\x46\x73\x4d\x06\xaa\x68\xaf\xd7\xe4\x9e\xee\xb8\xa0\x46\x2a\x47\x4f\xd8\x9f\xb8\x36\x5e\xf4\x7f\x31\x6a\x46\xc5\x2c\x83\x6d\x78\x5c\x55\x77\x87\x8f\xd3\xee\x0d\x53\x07\xde\xba\xdd\x3a\x3c\xd6\x77\x87\x8f\x67\xbb\x89\x1e\x58\x0b\x20\x36\x76\xe9\x4c\xbc\x49\x84\x60\xe9\x07\xd6\x31\xaa\x1d\x43\x15\x1e\xeb\x0c\xc3\xc7\x89\xe1\xc7\x64\x71\x8b\x00\x0e\x20\x59\x15\x1c\x98\x15\x01\x6c\x79\x9a\x81\x90\x03\x55\x9c\x3e\x77\xc1\x00\xad\x62\x66\x51\x7f\xfb\x41\x92\x26\x9c\xc5\xc7\xef\x86\x09\xcd\xa5\x70\x76\xb0\x62\x4d\x04\xbd\xaa\xb9\xc8\x63\xdc\xf4\x54\xc6\x02\x0e\x93\x68\xce\xa6\xf9\x97\xd8\xb0\x80\x79\x32\x71\x86\xf0\x90\x51\x27\x11\x98\xea\xb9\x8e\xcc\xd3\xca\x6e\xb2\xb1\x7a\xe5\xc3\x2f\x06\xa3\x8b\xc0\x69\x15\x82\xf0\x46\x31\x6a\xd8\x24\x7a\x43\x48\xeb\x08\x41\xe8\xe9\x90\xa2\x1f\x67\x6e\xed\x23\x75\x78\xc1\x10\xa3\x23\x5c\x02\x11\xa4\x08\xea\x47\x29\x82\xe2\xab\x40\x8f\x0e\x96\xf9\x9b\xf7\x6d\x23\x07\x00\xa0\x8d\x1c\xa6\xed\x3e\x37\x5c\x65\x1b\x02\xcf\x10\x03\x91\x67\x08\x81\x55\xa0\xc7\x28\xca\x82\x0a\x6a\x9e\x48\x5e\xf3\x4b\x20\x6e\x59\xc7\x90\x14\x2f\x8e\x70\x09\x04\x56\x64\x72\x6a\xa4\x8d\x0d\xe8\x15\x88\xf6\x08\x60\x17\x19\xe6\xc6\xef\x8f\xda\x65\xb8\x48\xc5\x57\xe1\x62\x95\x27\x5c\xa4\xf7\xab\x70\x83\x1d\x40\x08\x47\x33\x80\x14\x02\x43\x7c\x0d\x73\xcb\x04\xfb\x11\xed\x8f\x66\xc0\xb0\xde\x0a\x6f\x81\x0d\x56\xc0\xb0\xde\x08\x6f\x81\x9d\x1b\x21\xe6\x4c\xe0\x14\x36\x2f\xae\x7c\xfe\x9c\xd2\x24\x4e\x9a\x15\xcd\x21\xd6\xe4\x08\xe7\x61\xcd\xd5\x8d\x58\x04\x9c\xfe\x79\x60\x93\x92\x79\x22\x0f\x87\x3d\x91\x57\xf1\x83\x35\x89\xc4\x68\xb8\x89\x80\x54\xcd\x11\xc3\x39\xbf\x02\x71\x52\x38\x47\x0c\x47\xfc\x0a\xc4\x5c\xeb\x3c\x59\x27\xcc\xfc\x22\x5a\xcf\xae\xab\x2c\x09\x2f\x1b\x23\x4f\xe9\x3f\x8d\x11\xb0\x51\xa0\x79\xeb\xfc\x4d\x0a\xd9\x52\x12\xde\x93\x51\x2f\x5b\x59\x82\xdb\x74\x35\xe2\xb5\xaf\x4a\x01\x61\x62\xf6\xe8\xeb\x83\xab\x70\x21\xe3\x8b\x10\x55\x3b\x33\x13\x16\x2a\x21\x48\x4b\x8a\x00\x62\x64\x0b\x88\x13\xef\x0f\x52\x7e\xed\xa9\xfa\x0a\xae\xe1\xe7\x40\xba\x47\x25\xb2\xbd\x06\x3f\x48\xd9\x31\x2a\x42\x7d\x10\xeb\xa6\x86\x0b\xc3\xd4\x96\xb6\x2c\xd1\x5c\x99\x30\xd0\x1d\x5b\x93\x3b\xe1\xea\xa8\x56\x8e\x36\x17\x85\x95\x60\xdf\xcd\xcd\xa8\xb4\x54\x53\xb5\x14\x50\x41\x2d\xe8\x0b\x0f\x40\x20\xbc\x1f\x3a\xe6\x74\x3a\x9f\x13\x13\x46\x71\xa6\x53\x3d\xf4\xb4\xc8\x3f\x15\x7f\x9e\x7d\x5a\xbf\x91\xbb\x07\x3a\xc5\x3c\x56\xde\x13\xf7\x48\x78\x33\x7b\x87\xb4\xcc\x1f\xf4\x0d\x9e\x3f\x20\xbc\x91\x7f\x40\x5a\xe6\x0f\xba\x1e\xcf\x1f\x10\xde\xc8\x3f\x20\x2d\xf3\x77\x5d\xa4\x63\x6c\x9f\x1c\x78\xf0\x7e\x8b\xd5\x53\xde\xa5\x0d\xa8\x74\xce\x62\xda\xe7\xd7\x97\xb5\x6b\x83\xb1\x6d\x91\x5d\x33\x0e\x79\xc9\x6b\x69\x3d\xd3\xda\xe9\x97\xf8\xda\x34\x03\xd7\x7b\xaa\xf7\x48\x2e\xaa\x98\x30\xff\xc9\xa8\x8a\x6d\xe1\xb2\x24\xe2\x54\x1d\x42\xd7\x3b\x43\xc4\x56\xf6\x3d\x15\x2f\x10\x1d\x36\xd0\x57\xb8\x63\x44\x05\x18\x38\x31\x20\xbf\x54\x36\x93\x3d\xda\x5e\xe4\xe9\x0a\xf7\x7f\xa8\x72\xb1\xef\xac\xac\x0b\x7a\xd9\xab\x61\xe8\xe4\xd1\x7e\xbe\x31\x8a\x1a\xb6\x3b\x26\x68\xa7\x35\x2c\x85\x9c\xe6\xb0\x71\xc4\xda\xe7\x6a\x7d\x65\x47\xb8\xe4\xfa\x96\x6d\xe9\xd8\x19\x90\x2a\xad\x7e\x9d\x54\x8b\x22\x4e\x13\x03\xe4\xa6\xb9\x6c\xa9\x86\x02\x79\x29\x93\x2f\x93\xe7\x40\xbb\x11\x1f\x43\x2b\xb1\xb5\x4a\xc7\xe9\xdd\xcb\x06\x40\xc9\xbe\x07\xa6\xa2\xc7\xa7\x8c\xb6\x74\x42\x33\x75\xb9\xde\x84\x4a\x0d\x5f\x28\x70\xcc\x80\x7c\xd0\x7a\xca\x89\x93\x68\x87\xf1\x81\x7d\x1b\x99\x36\x19\xf5\x13\xef\x39\xa2\xf5\xac\x97\xea\x58\xf8\xd8\xbf\x98\x7d\x6f\x6c\x94\x0b\xd7\xfb\xfe\x5b\xd1\x96\xdd\x33\xc5\xe5\xcb\xa9\x50\x9a\x3a\x46\x78\x85\x9d\x11\x4a\xd8\xf6\x54\x19\xbe\xa5\xce\x31\x7c\x5f\x4e\xc8\x9e\xd1\x97\x90\x3c\xe2\x98\xc7\x49\x49\x79\x57\xa2\x6b\x43\x0d\xc3\x69\x20\x1b\x15\xd4\x06\x05\x6e\xe7\xe7\x79\xf6\xb9\xe0\xa8\x61\x4e\x47\xf9\xfc\xc4\x61\xc2\x61\x9c\x97\x79\x90\x9a\x1b\xa9\x8e\x38\xa7\x84\x72\x3f\x52\x76\xdc\x7c\x51\x5d\x46\xb9\x57\xd2\xc8\x56\x22\xb2\xd2\xf4\x5e\xf1\x03\x35\xec\x7f\x38\x62\xec\x8b\xf1\xb9\xe3\x6d\x46\x8f\x23\x37\xbd\x97\x7f\xdc\xba\x8c\x62\xd5\x0c\xfe\xbb\x30\xc7\xcb\x26\x71\xed\xa8\x6c\x6e\x7e\xc8\x06\x13\xaf\x99\x92\x5d\x3c\xc7\xbb\x70\xee\xc5\xa0\x87\x14\x87\x51\x3b\x6e\x3e\x28\x2a\x5a\x74\xcd\xb4\x52\x18\x2e\x46\x39\x6a\x6f\x26\x94\x0a\x19\x2a\x55\xe7\xf5\xe8\x54\x7a\x02\xdb\xd6\xdc\x2a\x9b\xc6\xf9\xc4\x1d\x69\x27\x92\x85\xec\x07\x29\x18\xbe\x75\xf2\x44\x97\xe5\xd2\x45\xb7\x97\x62\xcb\x77\x29\x48\x17\xc2\x20\x35\x18\x30\x1e\x6a\x82\x97\x52\x44\xa9\xf1\xab\x24\x8b\x99\x5c\xa3\x36\xb2\xbf\xc9\xa8\xb3\x14\xf1\xb3\x22\x7f\x36\xaa\x44\xd9\xb0\xa6\x73\x3e\xb4\xcb\x74\xce\x2d\x06\x43\x61\xc3\x77\x22\xa4\xc0\x3c\x4d\x54\xde\xe5\xaa\xe7\x4e\x50\x33\x45\xc1\xd8\x5b\x2e\xb8\xde\xc3\x7b\x6e\xc1\x0d\x08\xb7\x31\xd7\xb8\xbf\x68\xc0\x19\x8d\xe1\x79\xd5\x53\x58\x31\xd7\xe1\x70\x2c\x07\x14\x3e\xc8\x34\xad\x5e\x8a\x5d\x24\xbb\xff\x01\xa3\xa4\x01\xf8\x69\xc3\xe9\x31\xcf\xdb\xb5\x18\x87\x0d\xa5\x75\x09\x04\x8e\xf2\x94\x03\x6e\x51\xf1\xde\x10\xd2\xf9\x5b\xfb\x4e\x18\xec\x7a\x10\x07\x0e\x81\xe7\x66\x06\x17\x6b\xb2\x09\x88\x41\x4c\x2c\x5a\xd0\x79\x81\x54\x2d\x7b\x60\xcf\x23\xef\x66\xfa\x4d\xf5\x0d\x14\x0a\xce\x67\xe7\x42\x15\x79\x5f\x56\x6e\xe7\x7b\x2b\xb5\x76\x90\xe3\x5e\x2a\xef\x87\x57\x4f\x85\x68\xa8\x6a\x5d\xac\xb1\x6f\x67\xc4\x09\xd9\x59\xe3\x46\x0a\x43\xb9\x60\xca\xc9\x50\x32\x4a\x14\xc6\x77\x79\x52\x99\xec\x5c\xb2\xa0\xf8\xd1\x84\xfd\x15\xce\x0e\x26\xd7\xa9\xa7\xdf\x37\xa3\xda\xe5\xb4\x2f\x82\x1e\x28\xef\x6c\xbd\x5d\xee\xd1\xdd\xb5\x53\x92\x3a\x8e\x9a\xe7\xc7\xf9\xcf\x29\x5a\x67\xcd\x05\x54\x24\x1f\x17\x9f\x56\xe4\x35\x7d\x50\xb5\xbf\xc9\x8c\x1a\x67\xb9\x73\x31\x4e\xb5\x3b\xb3\x6c\x5e\x69\x7f\x50\x12\xac\x3a\x79\xad\x6d\x29\xfe\x3c\x88\x0d\x8a\xa6\x8d\x17\x9b\xb3\x58\xae\x54\xc5\xcc\x2f\x7d\x6c\x84\xb2\xc4\x85\x5c\xbe\x28\xf3\x75\x31\x37\x5d\xc3\x9b\x3a\xa3\x9f\x57\x8a\x54\x94\x8a\xb3\x99\x34\x4b\x45\x22\x97\xc6\xac\x4e\x62\xdb\x4e\x65\xd9\x1c\xcd\x6c\xf0\xc6\x90\x9e\x2e\xe5\x97\xd8\xcd\x1c\x70\xa7\xa8\xa8\xfa\x4b\x25\xfa\xe6\xaf\x13\x83\xea\xa1\x9f\xc9\xc8\x9b\xa6\xca\x28\x59\xae\x74\xf1\x60\xd3\x55\xc4\xf4\xe6\xfb\x2b\x00\x00\xff\xff\x0f\xbb\xe3\x06\x05\x22\x00\x00")

func pluginsCodeampGraphqlSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsCodeampGraphqlSchemaGraphql,
		"plugins/codeamp/graphql/schema.graphql",
	)
}

func pluginsCodeampGraphqlSchemaGraphql() (*asset, error) {
	bytes, err := pluginsCodeampGraphqlSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/codeamp/graphql/schema.graphql", size: 8709, mode: os.FileMode(420), modTime: time.Unix(1531446405, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsCodeampGraphqlStaticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xeb\x6f\xdb\x36\x10\xff\xde\xbf\xe2\xe6\x6c\x90\x5d\xd8\x94\xd3\xf5\x01\xa8\x76\x86\xb6\x49\xbb\x16\x69\xd3\xc6\x19\x8a\x7d\x2b\x4d\x9e\x2c\xa6\x14\xa9\x1e\x29\x3b\x6a\x91\xff\x7d\xa0\x64\xc9\x8a\x91\x0c\xd9\xb0\x41\x1f\x4c\xde\xe3\x77\x0f\xde\xc3\xb3\x9f\x8e\xcf\x5e\x5d\xfc\xf9\xf1\x04\x32\x9f\xeb\xa3\x07\xb3\xe6\x07\x60\x96\x21\x97\xe1\x00\x30\x73\xbe\xd2\xd8\x9c\x01\x96\x56\x56\xf0\x63\x7b\x01\xc8\x50\xad\x32\x9f\xc0\xe1\x74\xfa\xcb\xf3\x8e\x9a\x73\x5a\x29\x93\xc0\x74\x47\xda\x28\xe9\xb3\x7d\x39\xbb\x46\x4a\xb5\xdd\x24\x90\x29\x29\xd1\xb4\x9c\xeb\xed\xef\xc1\x8a\x78\x91\xa9\x6f\xfa\x76\x8b\xeb\x6c\x5f\x81\x5d\x6e\xfc\xc4\xdb\xaf\x68\x7a\x1a\x4b\x2e\xbe\xae\xc8\x96\x46\x26\xa0\x95\x41\x4e\x93\x15\x71\xa9\xd0\xf8\xe1\x41\xfa\x2c\x7c\x63\x38\xc0\x47\xe1\x1b\xed\x9c\x5b\x5a\x92\x48\x93\xa5\xf5\xde\xe6\x09\x1c\x16\x57\xe0\xac\x56\x12\x0e\xe4\x34\x7c\x3b\xc9\xd4\x1a\x3f\x49\x79\xae\x74\x95\x80\xab\x9c\xc7\x7c\x0c\x13\x5e\x14\x1a\x27\xed\x35\x5a\x70\x03\xaf\x89\x1b\xa1\x9c\xb0\xd1\x18\x22\xb6\x78\xfd\x61\x71\xac\x5c\xa1\x79\x35\x39\xc7\x55\xa9\x39\x05\xfa\x02\x57\x16\xe1\x8f\xb7\xd1\x18\xea\x63\x47\xfa\xfc\x31\xb0\x7f\x47\xbd\x46\xaf\x04\x87\x0f\x58\x62\x34\x86\xac\x25\x8c\x21\x3a\x2d\x85\x92\x1c\xde\x10\x37\x32\xf0\x38\x29\xae\xc7\xe0\xb8\x71\x13\x87\xa4\xd2\x9d\xd3\x05\x97\x52\x99\x55\x02\xcf\x8a\x2b\x38\x7c\x5c\x5c\xc1\xd3\xe2\x6a\x2f\x26\xa7\xbe\x63\x52\x33\x6f\x26\x7a\x16\xf7\x6a\x62\xa6\x95\xf9\x0a\x84\x7a\x3e\xa8\xa9\x2e\x43\xf4\x03\xc8\x08\xd3\xf9\x20\xf3\xbe\x70\x49\x1c\x0b\x69\x2e\x1d\x13\xda\x96\x32\xd5\x9c\x90\x09\x9b\xc7\xfc\x92\x5f\xc5\x5a\x2d\x5d\xdc\xbe\x73\x3c\x65\x87\x53\xf6\xa8\xbb\x33\xe1\xdc\x00\xe2\xdb\x0a\x31\x7e\x08\x67\x6b\x24\x52\x12\x1d\x3c\x8c\xdb\x02\x68\x35\x27\xc2\x1a\xcf\x95\x41\x02\xb6\x0e\x69\x58\x6a\x9c\xa0\x54\xde\xd2\x2d\xc5\xf4\xf4\xe9\xdf\x87\xe8\x04\xa9\xc2\x83\x23\x71\xef\x90\x52\xf4\x22\x8b\x1f\xb1\x29\xfb\xb5\x39\xb3\x5c\x19\x76\xe9\x06\x47\xb3\xb8\x81\xfb\xf7\xd8\x84\x5c\xf8\xf8\xf0\x09\x7b\xc2\x1e\x37\x97\xff\x15\x7c\x22\x6d\xfe\x1f\x1a\xb8\xf3\xb1\xf7\xe1\x67\x71\x3b\x85\x66\x61\xec\x6c\x2d\x4a\xb5\x06\xa1\xb9\x73\xf3\x41\xd7\xed\x83\xa3\x77\x9f\x2f\xe0\xa2\x6e\xfc\x99\x32\x45\xe9\x41\xc9\x3e\x1f\x0a\xcd\x05\x66\x56\x4b\xa4\xf9\x60\x27\xbc\xb2\xe8\x20\x43\xc2\x60\x59\xaa\x75\xcf\x46\x00\x68\x5d\x1b\x1c\x9d\x5a\x1e\xda\x85\x31\xd6\x97\xeb\xa7\x62\xcd\x09\x1c\x72\x12\x19\xcc\x61\xa3\x8c\xb4\x1b\xa6\xad\xe0\x5e\x59\xc3\x1a\xc6\xf3\x4e\xb0\xe0\xc4\x73\xf4\x48\x0e\xe6\xf0\xe3\xba\x61\x48\x2b\xca\x1c\x8d\x67\x2b\xf4\x27\x1a\xc3\xf1\x65\xf5\x56\x0e\xa3\x2e\x8e\x68\xc4\xd6\x5c\x97\x08\x73\x08\xd0\x7a\xe1\x2d\xf1\x15\x06\x85\xb7\x1e\xf3\x61\xd4\x3a\x9c\x5c\x6e\xfc\x45\xa3\xf1\xfc\x41\x0d\x9e\x96\x46\x04\x57\xa0\x16\xf9\x74\xfa\x3a\xd4\x24\xd2\x70\x7b\xfd\x18\x1c\x72\xa3\xae\x37\x84\x35\xce\x43\x8b\x02\xf3\x7f\xe0\x5c\xdb\x48\x1a\x3d\x84\x07\xdc\x06\xd9\x35\x5d\xf4\x42\x08\x2c\x7c\x94\x40\x14\x66\xa4\x6a\x52\x14\x5f\x3a\x6b\xa2\xf1\x4e\xea\x95\x35\x1e\x8d\x9f\x5c\x54\x05\xde\x2a\xdb\xf6\x6b\x6b\x4f\xa5\x30\x6c\x1d\x1e\xf5\xec\xdd\xc8\x94\xbb\x3b\x53\xe3\x2e\xdc\xde\x16\xb8\x2d\x80\xfb\x86\x70\x9f\x20\x6e\x48\xbf\x28\x7d\x66\x49\x7d\xaf\xf9\x51\xb2\x4b\xff\x6f\xf0\xe5\x25\x72\x42\x82\x9f\x7f\xb4\xc4\xeb\x2f\x90\x80\x29\xb5\xee\x10\xae\xf7\xd7\x21\xa1\x2f\xc9\x40\x3d\x7e\x86\x51\xfc\xad\x44\xaa\xa2\x71\x2f\x92\x1c\x7d\x66\x65\x02\x51\x61\x9d\xef\xf9\xb2\x8d\x7a\xdc\x5b\x86\xb2\x4a\xe0\xdd\xe2\xec\x03\x73\x9e\x94\x59\xa9\xb4\xda\x2b\x9d\x9d\xb0\x20\x94\x68\xbc\xe2\xda\x25\x10\x29\x23\x74\x19\xb6\x51\xeb\xdc\x88\xf9\x0c\xcd\xb0\x2b\xc8\x21\xa1\x2b\xac\x71\xd8\x7f\xb4\xad\xeb\x2d\x8b\x79\xbc\xf2\xc3\xee\x5d\xee\xc6\x78\x69\x65\xd5\xc7\xf1\x54\xdd\x78\xb9\x2d\x6e\x1d\x49\xc1\xc9\xe1\x4d\xcd\xdd\xc3\x5f\x83\xe0\x5e\x64\x30\x44\x22\x4b\xa3\xdb\x40\xfa\x9a\x3d\xc5\xce\xc7\x86\xd6\xdc\xe3\x18\xce\xd1\x48\x24\x98\xbd\xa9\x2b\xef\xd3\x29\xc4\x47\xa0\x8c\xb7\xe0\x33\xac\x13\xcc\x5a\xc9\x05\x62\x4d\x3c\x3f\x79\x71\xfc\xfe\x04\x94\xa9\x6f\xde\x16\xa0\x71\x8d\x1a\x6c\x0a\x3e\x53\x0e\x72\x2b\x4b\x1d\x18\xa0\x91\x93\x81\xdc\x12\x02\x5f\xda\xd2\xb7\x48\x99\xdd\x40\x65\x4b\x10\xdc\x80\x28\x9d\xb7\xb9\xfa\x8e\xd0\x79\xb0\xac\xa0\x20\xbb\x56\x61\xb4\x81\x54\x69\x8a\x84\xc6\x43\xdd\xc6\x0e\x2c\xb5\x30\xe1\xbf\x42\xc8\x33\xd7\x20\x32\xa5\x25\x60\x33\x01\x5c\xe3\xf2\x79\xd8\x12\xc7\x67\xef\x19\xd5\x21\x0e\xb7\x19\xa8\xc9\x4c\x10\x72\x8f\xdb\x91\x31\x6c\x4d\xf7\xab\x30\x6d\x46\x51\xb2\x37\x9a\x76\x05\xd3\x9e\xee\x9c\x41\x6d\x33\x47\xa3\x5a\x72\x9b\xf9\x9b\xbb\xa4\x59\x21\xb3\xb8\xf9\x8b\xfb\x57\x00\x00\x00\xff\xff\x45\xab\x31\x54\xfa\x0a\x00\x00")

func pluginsCodeampGraphqlStaticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsCodeampGraphqlStaticIndexHtml,
		"plugins/codeamp/graphql/static/index.html",
	)
}

func pluginsCodeampGraphqlStaticIndexHtml() (*asset, error) {
	bytes, err := pluginsCodeampGraphqlStaticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/codeamp/graphql/static/index.html", size: 2810, mode: os.FileMode(420), modTime: time.Unix(1529614596, 0)}
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
	"plugins/codeamp/graphql/schema.graphql": pluginsCodeampGraphqlSchemaGraphql,
	"plugins/codeamp/graphql/static/index.html": pluginsCodeampGraphqlStaticIndexHtml,
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
	"plugins": &bintree{nil, map[string]*bintree{
		"codeamp": &bintree{nil, map[string]*bintree{
			"graphql": &bintree{nil, map[string]*bintree{
				"schema.graphql": &bintree{pluginsCodeampGraphqlSchemaGraphql, map[string]*bintree{}},
				"static": &bintree{nil, map[string]*bintree{
					"index.html": &bintree{pluginsCodeampGraphqlStaticIndexHtml, map[string]*bintree{}},
				}},
			}},
		}},
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

