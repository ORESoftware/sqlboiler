// Code generated by go-bindata. DO NOT EDIT.
// sources:
// override/templates/17_upsert.go.tpl (5.749kB)
// override/templates/singleton/psql_upsert.go.tpl (1.317kB)
// override/templates_test/singleton/psql_main_test.go.tpl (4.979kB)
// override/templates_test/singleton/psql_suites_test.go.tpl (255B)
// override/templates_test/upsert.go.tpl (1.746kB)

package driver

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\xdf\x6f\xdb\x38\x12\x7e\x96\xfe\x8a\x69\x70\x68\xa4\x83\xa3\xdc\x73\x0a\x3f\xe4\x47\xdb\x0b\x7a\x4d\x7d\x4d\xb3\x05\xb6\x28\x02\x59\x1a\xd9\x44\x68\x52\xa5\xa8\x38\x5e\xad\xfe\xf7\xc5\x90\x94\x25\xd9\x72\xe2\x76\xdb\xdd\xee\x53\x2c\x72\xc8\xf9\xf8\xcd\x37\x9c\x61\xaa\xea\x08\xfe\x15\x73\x16\x17\x70\x32\x86\xe8\x94\x7e\x61\x11\x7d\x88\xa7\x1c\xc1\xfe\x89\xae\xe2\x05\xd6\xb5\x6f\x4c\x8b\x64\x8e\x8b\xd8\x4e\xd3\x82\xd6\x02\x7e\x87\xe8\xba\x9d\x35\x0b\x58\x06\xd1\x69\x9a\xbe\xe6\x72\x1a\x73\x38\xaa\x6b\xff\xf8\x18\x6e\xf2\x02\x95\x7e\x0d\xb1\xd6\xb8\xc8\x75\x01\xb1\x00\x26\x68\x6c\x04\xb1\x48\x21\x95\x68\xc6\xca\x3c\x8d\x35\x82\x54\xc0\x66\x42\x2a\x04\x29\x20\x91\x22\xe3\x2c\xd1\x91\x9f\x95\x22\x81\x40\xc2\xbf\xab\xca\xe2\x8f\x6e\xf2\x6b\x26\x66\x25\x8f\x55\x5d\x87\x8d\x97\xc0\x80\x10\x52\x43\x74\x25\xcf\xa5\xd0\xf8\xa0\xeb\x3a\xd1\x0f\xb4\x15\x7d\x44\x6e\x70\x04\x55\x85\x22\x25\x90\xce\xf3\x3b\x71\xee\xbc\xc1\x54\x4a\x3e\x5a\x3b\x3f\x97\xbc\x5c\x88\x02\x3e\x7d\x2e\xb4\x62\x62\x36\x72\x0b\xdc\xf8\xc8\x9d\xa6\x31\x9b\x4a\xc6\x23\xf7\x11\x02\x2a\x25\x15\x54\xbe\xa7\x50\x97\x4a\x80\x8c\x2c\x52\x0b\xb4\x0b\xd2\xac\x7b\x8d\xfa\xe2\x2c\x08\xab\x0a\x79\x81\x06\xf8\x08\x9a\x09\x67\xe9\xe6\x45\x5a\xd7\xa3\x2d\xe8\x5b\xa8\x1f\x07\x1b\xfa\xb5\xef\xaf\x89\xf0\x6d\x08\x29\x28\x9d\x30\xd2\xcf\x49\x2c\x58\xb2\x11\xd0\xc9\x9f\x8b\x28\x98\x3d\x0b\x1a\x33\x1c\xed\x1d\xe2\xc9\x4f\x17\xe3\xca\xf7\x58\x46\xa7\xa0\x14\xf9\xc9\x02\xfc\xc2\xe0\x7a\x36\x06\xc1\x38\x01\xf5\x72\xa2\x3d\x30\x2e\x3f\xaa\x38\x7f\xa9\x54\x80\x4a\x85\xa1\xef\xd5\x43\x62\xd8\x11\xfd\xa1\xe0\x43\x59\x30\x31\xa3\x6f\x7c\xc0\xa4\xd4\x52\x7d\x4d\x82\x77\xb6\xce\xbf\x4d\x19\x93\x6d\xca\x09\x88\xa5\xf7\xa5\x83\xd4\x21\x7e\x5b\x2e\xad\xb9\x1b\xea\xac\x1a\x0e\xc7\x5f\x24\xa3\x01\xb1\x77\xc5\x4d\xb8\xff\x56\xa9\xac\x83\xf7\x23\x64\x71\x8d\xd8\x63\x0a\x52\x99\x94\x0b\x14\x3a\xd6\x4c\x0a\xc8\xa4\x82\xb9\x5c\x82\x96\x90\x2b\x99\xa3\xe2\x2b\x28\x0b\xec\x9f\xd5\x78\xec\x1d\x77\x5f\x55\xfd\xc3\x45\xb5\xae\x3f\x2c\x03\x09\xe3\x36\xb8\xae\x1e\x99\xf9\x22\xba\xc2\x65\x70\x50\x55\xd1\xe4\x6e\x66\xcb\xff\x09\x08\x09\x55\xd5\x6b\x09\x88\xdf\x7b\x96\x62\x6a\x38\x2f\x0d\x3d\x07\x46\x0d\xbe\x47\xdd\x02\x45\x9e\x53\x2c\x0f\x34\x5b\x60\xa1\xe3\x45\x7e\x6b\xad\x6e\xe7\xc8\x73\x54\x07\x10\x41\x6d\xad\x5b\x51\xff\x57\xca\xbb\xc2\xc8\xa8\x27\xff\x54\x9e\x61\x26\x15\xda\x28\x18\xa3\xbd\x73\x61\x5b\xca\xed\x69\x09\xae\x41\x6b\xc8\xf7\x7d\x4f\xfc\x76\x81\x59\x5c\x72\x6d\x5a\xa2\x2f\x25\x2a\x86\x45\x74\x25\xc5\xaf\xa8\xa4\x9b\xba\x46\xd2\x81\x53\xc9\x85\x5c\x8a\x56\x27\x8e\xe9\x8f\x4c\xcf\x9d\xf1\x08\x64\xe8\xfb\xde\xf1\x31\x9c\x95\x8c\xa7\x90\xc4\xc9\x1c\xe1\x0e\x57\xc0\xc4\x11\x67\x02\xa1\x9c\x71\xc6\x57\x70\x04\x8b\x55\xf1\x85\xc3\x7d\x01\x39\xfd\xcd\x95\x9c\x72\x5c\x14\xbe\x37\x2d\x33\x02\x53\x68\xb5\x88\xc5\x8c\x23\x55\x87\xb3\x32\xcb\x50\x05\xa1\xa1\x69\x4b\x32\x74\xc8\x69\x99\x45\x1f\x15\xd3\x78\xb6\xd2\x18\x1c\xea\x43\x8a\x0d\x90\x34\x87\xa6\x33\x33\xed\x6f\x0e\x47\x34\x4c\xf1\xbd\x1d\x41\x42\x20\x54\x2c\x66\xb8\x25\xc6\xde\x86\xd7\x46\x97\x41\xb2\x7b\xc3\x4d\xd3\x42\xab\x44\x8a\xfb\xe8\x52\xcb\x38\xe8\xc9\x39\x7a\xc3\x44\x1a\x0e\x62\xe8\xdb\x9d\x4b\xfe\x7d\x61\xf4\xaf\x87\xdd\x30\xfa\x76\xdf\x02\x63\x7b\xcf\x8e\x08\x1f\xd9\x8b\x34\x74\x32\x06\x9a\x75\x13\xa1\xef\xb5\x22\x99\x94\x8d\x48\xa6\x65\x16\x9a\x34\x1b\x94\xac\x4d\xa9\x73\x92\xe5\xdb\x52\x47\xef\xff\x27\x93\x3b\xda\xc9\x08\x75\x64\xf5\x9a\x92\xa3\xa7\xd7\x7f\xba\xc3\xd5\xe7\xbd\x1d\xdd\x08\x6e\x5d\xf9\xde\x7d\xac\x4c\x8e\x9a\xfb\xc7\x37\x9a\x7e\xe6\x1c\x13\x01\x4d\x3b\xa9\x50\x13\x90\x3e\xe5\x97\x9d\x2f\xca\x4c\xdf\xf3\x76\x21\x38\xe5\xbc\xb9\x26\x1f\xb1\x1a\xc8\xe1\xfd\xac\x65\xa9\xbb\x0b\xda\x28\xd2\x67\xe8\x7b\x9e\x2b\x6e\x27\xe3\x0d\xf1\xde\x74\xbe\xbe\xcb\x11\x26\x8a\x2d\x62\xb5\x7a\x83\xab\x8e\x31\x11\x3d\x78\x5b\x3c\x7f\x0e\x1c\x85\x4b\xbc\x90\xca\xc2\x7f\x0c\xed\x4f\x57\x85\x52\x98\xb7\xa0\x96\xee\xfe\xdf\xac\x11\x54\xb6\x4a\x9e\x9a\x5b\x7a\x6a\xae\x3f\x47\x41\x62\x60\x01\x67\x85\xa9\x19\xa6\x68\x78\xcd\xad\x42\x04\x6d\xdc\x30\x16\x39\xa1\x6c\x26\xba\x38\xd7\x0b\xc7\xb0\x88\xef\x30\x68\x6b\x23\xad\xd8\x97\x23\xca\x6f\xda\x2b\x5f\xad\x9d\x8c\x76\x89\x7e\x7b\xb1\x39\x84\x67\xb3\x26\xa2\xba\xb1\x82\xb1\x3d\xb3\xd5\xfd\xff\x69\x68\x22\x0b\x3d\x53\x58\x04\x29\x8b\x39\xd2\xfe\x07\x55\xd5\x7d\x56\xd7\xf5\xc1\x50\xeb\xa6\x50\x37\xc3\x6d\x27\xd0\x94\x7a\x13\x57\xeb\xf7\x3e\xe6\x25\xbe\x8d\xf3\xdc\x1c\x9e\x32\xaa\xad\x61\x67\x4c\xa4\x6e\x6a\x17\x25\x1f\x56\x39\xee\x3c\xf2\x7a\xdb\xc6\xab\xd7\x54\xe8\x4e\x65\xed\x95\x56\x43\x88\x0b\x9b\x42\x1d\x92\x61\x13\x31\x03\x57\xa1\xfe\xd1\x60\xc9\x2f\x39\x1c\x80\xda\xc7\x6a\xc0\xd6\xb6\x7d\x31\x34\x9a\xeb\x18\x33\x0a\x53\x74\x29\x52\xa6\x30\xd1\x41\x33\xf0\x0b\x59\xbc\xcb\x02\x49\xa2\xb9\x8f\x79\xaf\x5b\x30\x93\xc5\x2b\x25\x17\xcd\x11\xcc\x86\xee\x2e\xed\x05\x29\xb4\x77\x9f\x45\x42\x4d\x1d\x13\x1a\x55\x16\x27\x58\xd9\x0e\xc8\x48\x7e\x83\xac\x0e\x91\xcd\xc2\xd6\xf9\x44\xab\xdd\xae\x3b\x7b\x34\x8d\x5a\xaf\x9d\x5d\x37\x5e\xa6\x6d\xbc\xc0\x69\x39\x7b\x2b\x53\xdb\x30\x64\x0b\x1d\xbd\xca\x15\x13\x9a\x8b\xa0\x9d\x37\x85\x49\x35\x0e\x8c\xf0\xc3\xa7\xad\x89\xb2\xd0\x35\x5f\xa6\x25\x69\x1c\xa7\x64\x35\x82\xa5\xb1\x33\xe5\x8d\xd6\x5e\x16\x66\x75\x90\xe8\x87\xf0\x85\xb5\xd9\xc6\xb4\xdc\x03\xc9\x72\xc8\x7f\xf3\x6e\xd9\x83\xef\x41\xbe\x3c\xab\x5d\xea\x35\x23\x93\xe6\xef\xe5\x32\xe8\xa0\xb0\xee\xa2\x28\x0a\xa3\xeb\x24\x36\xb9\x40\x41\xa3\x01\xb3\x65\x4b\xc0\xd0\x4e\xce\x55\x60\x3a\xdb\xaf\xd9\xd5\x1d\x6b\xad\xfd\xf1\x18\x8a\x2f\x3c\x7a\xa9\xd4\x95\x7c\x2f\x97\xb6\xb7\x70\x1e\x29\x29\x8e\x8f\xa1\xb9\x9f\xcc\x73\x4c\x1c\x6a\x27\x4c\x88\xc5\x4a\xcf\xe9\xdd\xb6\x9c\xa3\x00\x3d\x47\x85\x87\x05\xbd\x09\xec\x9d\xe4\x32\xa7\x6d\x2e\x87\x69\xba\x6d\xb2\xdc\x9c\x8f\x1e\x3e\xc3\x2c\x6d\x92\xb2\xbd\xee\x69\x4e\xfa\x14\xb4\xaf\x89\xc1\x57\x00\x55\x37\x7a\xd3\xd2\x83\xd6\x5c\xc9\x5f\x53\xe3\x9a\x37\xcf\x46\xcf\xb2\x5f\x13\xd4\x34\x5b\x7b\x98\x9b\xe6\x0a\xc6\xf6\xb8\x7b\x3b\x58\x37\x59\xde\x23\x2f\xad\xf5\x7f\x23\x53\x79\x9a\x69\x54\xdf\xf4\xca\x72\xef\xa8\x75\xd8\xdc\xa6\x82\xf1\xee\x0b\xab\xf6\xff\x08\x00\x00\xff\xff\xeb\xce\x96\x83\x75\x16\x00\x00")

func templates17_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertGoTpl,
		"templates/17_upsert.go.tpl",
	)
}

func templates17_upsertGoTpl() (*asset, error) {
	bytes, err := templates17_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbf, 0x1, 0x23, 0xda, 0xa, 0xfe, 0x87, 0x7a, 0xf7, 0x5c, 0xaf, 0xa8, 0x77, 0x7a, 0xaf, 0x1e, 0xd9, 0x67, 0x90, 0xc7, 0x29, 0x65, 0x7d, 0x47, 0x1, 0x4d, 0x6, 0x82, 0xff, 0x29, 0x41, 0xf1}}
	return a, nil
}

var _templatesSingletonPsql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x5f\x6b\xdb\x3e\x14\x7d\x96\x3e\xc5\xad\xa0\xd4\x06\xe1\xfe\xfa\xfa\x83\x3c\xb4\xb1\xdb\x65\x04\xbb\x89\xed\x6d\x30\xf6\xe0\xd8\xd7\xa9\xc0\x91\x33\xfd\xc9\x56\xd6\x7c\xf7\x21\xc7\xae\xd3\xa6\xa3\x14\x82\x12\x74\xef\x39\x3a\xf7\xe8\x28\x97\x97\xb0\xb2\xa2\xa9\xf2\xad\x46\x65\x16\x16\xd5\xe3\x7d\xab\xcd\x5a\xa1\x3e\x14\x34\x14\x90\x2e\xe6\xa0\x4d\x61\x70\x83\xd2\x80\x36\x4a\xc8\x35\x58\xed\x56\xf3\x80\x60\x3b\x6c\x58\x98\x02\xb6\xaa\xdd\x89\x0a\xab\x80\xd6\x56\x96\xff\xa4\xf6\x2a\x51\x40\xa5\xc4\x0e\x95\x0e\x42\x51\x34\x58\x1a\x0e\xa6\x58\x35\x18\x17\x1b\xec\x8f\xe0\x60\xb7\x55\x61\x30\x91\xd3\x56\xd6\x8d\x28\x0d\xac\xda\xb6\xe1\xa0\xd0\x0c\x35\x0e\x65\x5f\xe3\xf0\xeb\x41\x18\x6c\x84\x36\xf0\xfd\xc7\x81\xc1\x1f\xc4\xfe\xa1\x64\xe8\x83\x89\xdb\xdc\x14\x72\xdd\x60\x30\xab\x50\x9a\x85\x6d\x0d\xa6\x8d\x28\xd1\xe9\x0a\xe6\x0b\x0e\xee\x7b\xb9\x18\xc9\x7d\x4a\x46\xf6\x8f\x10\x3c\xa3\x7c\x4a\x14\x7e\x0c\xab\xd0\xf8\x94\x92\x95\xad\xe1\xff\x63\xdc\x1d\x9a\x1b\x5b\xd7\xa8\x3c\x9f\x92\x0a\x6b\x54\x47\xc5\x7b\x3b\x14\x57\xb6\x76\xf0\xb2\x6d\xec\x46\x6a\x47\xc1\xc2\xe8\xf6\x3a\x9f\x67\xf0\xe5\x7a\x9e\x47\x29\xa3\x44\xd4\xd0\xa0\xf4\x46\x95\x70\x36\x81\xff\x9c\x5d\xcf\xb8\x09\xd4\x1b\x13\xa4\x5b\x25\xa4\xa9\x3d\xe6\x9d\x6b\xbf\xc7\x83\xfb\xcd\x38\x25\x84\x1c\x6c\xd6\xc1\xe7\x56\x1c\xb1\x71\x60\x1c\x98\x3f\x74\x0c\x0a\x9b\xa2\xc4\x87\xb6\xa9\x50\x75\x41\x08\x72\x8d\x33\x59\xe1\xef\xe3\x02\x7f\xa5\x8b\xc3\x15\x87\x2b\xdf\xa7\x64\x4f\x29\x71\x8a\x6e\x7b\x45\x94\x38\x87\xdc\x19\x6c\x16\xa7\xd1\x32\x83\x59\x9c\x25\x70\xae\xdd\x27\x89\x61\x9a\xc4\xb7\xf3\xd9\x34\x83\x4e\xe9\x73\xc6\xf8\x38\x22\xa7\xc4\x19\x25\x6a\x38\x3b\x09\xdc\xd3\x53\x27\xe4\xb0\xef\xc3\x64\x70\x67\x65\xeb\xe0\xab\x12\x06\xd3\x6e\x72\x8f\x85\x09\xc4\x49\xf6\x69\x16\xdf\x31\x27\x12\xb0\xd1\xf8\xb2\xf3\xe6\xd1\xa0\x77\xe1\x5d\xf8\x6f\xc0\x5f\xf8\x37\x26\xba\xb3\xef\xad\x7e\xe6\x43\x98\x40\x7e\x1f\x5e\x67\x11\xa4\x51\x06\xcc\x4d\x40\xea\x56\x81\xe0\xb0\x73\x97\xad\x0a\xb9\xc6\xfe\x95\x74\x42\xdc\x80\x62\xbc\xdf\x13\x65\xbc\x53\x46\xf6\x6e\xf9\xe9\x52\x59\xbd\x8c\xdd\x18\xd7\x93\xa4\xee\x3a\xe4\x6b\x91\x07\x92\x37\x4b\x0c\x26\x10\x7d\x9b\xce\xf3\x30\x0a\x03\xf6\x0e\x7a\x7f\xb8\xf4\x3e\xab\xee\x55\x8c\x53\x9c\x12\x2f\xa3\x2c\x5f\xc6\xb3\xf8\x0e\xd8\xbb\x4e\x77\x7f\x24\x83\xc9\xee\x0c\x85\xc6\x2a\x09\x0e\xd4\xf7\xfb\x74\x4f\xff\x06\x00\x00\xff\xff\x76\xcb\x6a\x7a\x25\x05\x00\x00")

func templatesSingletonPsql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonPsql_upsertGoTpl,
		"templates/singleton/psql_upsert.go.tpl",
	)
}

func templatesSingletonPsql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonPsql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/psql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcb, 0x25, 0xb2, 0xa1, 0x78, 0x77, 0xde, 0x3, 0x46, 0x95, 0xac, 0xcf, 0xb0, 0xa1, 0xf5, 0xf1, 0x14, 0x37, 0x11, 0x4d, 0x5f, 0x91, 0x90, 0x62, 0xc7, 0x92, 0xef, 0xee, 0x53, 0xcb, 0xb2, 0x5f}}
	return a, nil
}

var _templates_testSingletonPsql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x6d\x6f\xe3\xb8\x11\xfe\x2c\xfd\x8a\x39\x03\x39\x48\x5b\x87\x3e\xf4\xe5\x4b\x0e\xc6\x22\x71\x9c\x74\x71\xd9\x24\x6b\xbb\x3d\x14\xdd\xf6\x8e\x96\x46\x0e\x11\x89\x64\x48\x2a\x59\xf7\x90\xff\x5e\x0c\x29\xd9\xb2\x63\x25\xdb\x5e\x0b\xdc\xa7\x44\xe4\x33\xef\x0f\x87\x43\x3f\x72\x03\x66\xf5\xe5\xf6\xf2\xe2\x1e\xd7\x30\x06\x83\x2b\xfc\xa2\xd9\xc7\xda\xba\x89\xaa\xb4\x28\x31\xf9\x39\x79\x5f\xa5\xff\x3c\xbd\x5a\x4c\x67\xb0\x38\x3d\xbb\x9a\xc2\xcd\xf5\xd5\xdf\x80\xbd\xfb\x2c\x3f\xdb\xdf\x9d\x9e\x9f\xc3\xe4\xe6\x7a\xbe\x98\x9d\x7e\xb8\x5e\x00\x7b\xf7\x1e\x2e\x6e\x66\xd3\x0f\x97\xd7\xf0\xc3\x94\x50\xef\xbf\xff\x2c\x7f\x4e\xe3\xd8\xad\x35\x82\x5e\x2d\xd0\x3a\x34\x60\x9d\xa9\x33\x07\xbf\xc4\x51\xbe\x9c\x28\x29\xe1\x9d\x7d\x28\xd9\xf9\x59\x4c\x0b\xd7\xbc\x42\x20\x88\x90\xab\x38\xba\x53\xd6\x01\x6c\xbf\x6b\x8b\xa6\xfb\xad\xb9\xb5\xdd\x6f\x6b\xcb\x4a\xe5\xb8\xdd\x57\xc6\xcb\x0b\xe9\xe2\x38\xd2\xab\x5b\x6e\xed\x85\x28\x37\x80\x38\x72\x68\xdd\xf9\x99\xb7\xba\x51\x72\x2f\xf4\xfc\xd3\xd5\xa4\xca\x61\xa9\x54\x19\x3f\xc7\x71\x51\xcb\x0c\x84\x14\x2e\x49\x83\xdf\x1f\xb9\x90\x30\x86\x6f\xdb\xa0\x7e\x79\x26\xd8\x68\x04\x16\x5d\xad\x21\xaf\x2b\x6d\xc1\xdd\x21\xe4\xdc\xf1\x25\xb7\x08\x36\xbb\xc3\x8a\x03\x97\x39\x88\x8a\xfc\xb2\x20\x1c\x39\xa6\x80\x83\x43\x5a\xe2\x66\x0d\x86\xcb\x5c\x55\xe5\x9a\x74\xad\x50\xa2\xe1\x0e\x73\x20\x2f\x3b\xaa\x14\xb8\x3b\xee\xfc\xaa\x85\x8c\x4b\x58\x22\x98\x5a\x02\x5f\x71\x21\xad\x23\xc5\xb5\x15\x72\x45\x1e\xec\x2a\xb2\x0f\xe5\x52\x89\x12\x0d\xdc\xcc\x3e\x82\xe6\xd9\x3d\x5f\x21\x0b\xf1\x25\x1a\xde\xb5\xf1\xa4\x21\x90\x24\x05\x34\x46\x19\x0a\x9a\xe8\x82\xc6\x84\x85\x38\x8e\x1e\x85\x46\xc3\xe6\xe8\xce\xb1\xe0\x75\xe9\x92\x81\xa6\x3a\x86\x38\x07\x43\x18\xe8\x7a\x59\x8a\x6c\x90\xf6\x42\x29\x0b\x83\x21\xfc\xe9\x8f\x7f\xf8\x7d\x3f\xa8\x29\x29\x29\x34\xf8\x50\x0b\x83\x83\x94\x6a\xc9\x1a\xae\x8c\x21\x08\x5e\xa2\x9b\xfb\x02\x36\x72\xf9\x52\xf2\x8a\xb0\x91\x66\x9e\x46\x7d\x40\xda\x0c\x30\xcf\xae\x3e\x18\x6d\x06\x98\x27\x5d\x1f\x8c\x36\x1b\x18\x71\xaf\x03\xfb\x20\x77\xe2\xf6\x98\x96\xaf\x7d\xda\xda\xe0\x3d\xb8\x43\xd5\x3e\x3c\x41\xba\x81\x77\xa8\xdc\x11\x39\x53\xaa\x6c\x0d\xdc\x0b\xfa\x9b\x55\xb9\xcf\x2a\xd5\x77\x0c\x8f\xbc\xe4\xec\x0c\x57\x42\xfe\x95\x97\x22\xe7\x4e\x28\x99\xa4\xac\xf9\xc0\x24\x8e\x22\x0f\x09\xa6\xaf\x95\x9b\x56\xda\xad\x93\x90\x40\x2a\xfc\x36\x5f\xc3\x5e\x2c\xa5\xbd\xc5\x86\x12\x6c\xb0\xd7\xca\x25\xfe\x9f\xe9\x43\xcd\x4b\x9b\x84\x5c\x0e\xe1\xbb\x16\x1f\x12\xf8\x8a\xf2\xc0\x8d\x16\xde\x66\xa4\x1f\xdf\xe4\xb9\x15\xd8\xa4\x7d\x18\x47\x29\x9b\xdc\x61\x76\x9f\x50\x7a\x44\xe1\x4f\xc0\x37\x63\x90\xa2\xa4\x33\x11\x19\x74\xb5\x91\xb4\x1a\x47\xcf\x71\x1c\x8d\x46\x20\x0a\x90\xca\x9f\x4d\x3a\x81\xe7\x67\x40\x94\xc0\xdc\x4b\x97\x28\x93\x6e\x21\x53\x18\x8f\xe1\x3b\xaf\x69\x34\x82\x89\x41\xee\x10\x78\xd3\x04\xc4\xbf\x30\x87\x7c\x09\xe4\x3c\x8b\xa3\x7d\x06\x6c\x40\x6c\xee\xf8\xb2\xc4\xb0\xb1\x09\x3e\x0d\x0e\x35\x2e\x8f\x41\xb3\x8a\xdf\xe3\xed\x65\xdb\x02\x93\xf4\xfb\xb7\x82\x11\x05\x7c\xb3\xc3\x21\x02\x75\x14\xe6\x46\xe9\x85\x77\xe9\x80\xb2\x1d\x6d\xd1\xf3\xae\x64\xe6\x23\xfd\x6a\xd9\x38\x8a\xa8\xa3\x92\x0b\x27\x63\xc0\x2f\x98\xb1\x89\xaa\x2a\x2e\xf3\x64\xa0\x57\x3f\xd1\x1e\xf5\x87\xe3\xe3\xd0\x7c\x8e\x95\x2c\xd7\x83\x21\x74\x52\xd1\xca\xb3\xa9\x7c\x84\x31\x70\xad\x51\xe6\x89\xb2\xf4\x2d\x0c\xd1\x9b\xe0\x7a\x35\x95\x8f\x49\xca\x18\x23\x91\xe0\xe4\x61\xa3\xf6\xa1\xf4\x06\x3a\xa5\xec\x4a\x7c\xbd\x19\x4a\xfb\x10\x9e\xc8\x84\x50\xec\x56\x68\x4c\x3a\xee\xce\x5d\x4e\xa9\x39\x19\xc3\xb7\xcb\xb5\x43\xcb\xce\xea\xa2\xf0\xb7\x4d\xc7\x58\x3f\xa8\x13\xf7\xdc\xe5\xaa\xa6\x7e\xf4\xb4\xbb\x18\x2a\xb2\x63\x2e\xde\x89\x64\xee\x72\x7f\xd5\x49\x7c\xba\xf8\x01\xd7\xe7\x68\x9d\x51\x6b\x34\xc9\x66\x74\x18\x82\x49\xf7\x45\x82\xda\x3d\x17\xe3\x2e\x09\xb6\x3e\x70\xe3\x5e\xe7\x80\x32\x96\xfd\x68\xb8\x4e\xd0\x50\x7b\x29\xb8\x28\xe9\x4e\x54\x60\x49\x16\x1a\x06\x40\x16\xaa\x43\x9d\x6f\x97\x6f\x5d\xcf\x7e\xb5\x31\xfb\x50\xee\x59\x3a\x14\xd5\x8f\x5c\x1c\xb4\x53\x54\x8e\xdd\x1a\x21\x5d\x29\xc9\x40\xba\xbf\xb6\x53\x88\xa6\x4f\x25\x69\xfa\x95\x2e\x3e\x71\xe1\xa0\x50\xa6\x27\x25\x71\x14\xfd\x44\x0c\x60\x93\x52\x59\x4c\x52\x18\x8d\xe0\xb4\xa0\x91\xac\x3d\x5d\xc2\x42\xae\x24\x0e\x21\x23\x84\x1f\x60\x9e\x8c\x70\x08\x28\x73\x50\x85\x5f\xd0\x42\x63\x7c\x38\xbd\xff\x6d\xd4\x7b\x3c\xf9\x15\x71\xbf\xac\x8e\x8f\xbb\xd1\x21\xc5\x76\x9a\xdb\x9d\x76\x4c\x2d\x27\x55\x9e\x58\x22\xfb\xb0\xd5\xd0\x4c\x84\x43\xe0\x66\x65\x81\x31\x16\xbe\x3b\x33\x51\x76\xa0\x39\x34\xc2\x41\x2a\xb4\x92\xec\x3f\xeb\x08\xcd\x45\xe1\x9d\x49\x29\x91\xe1\x86\xc8\x3a\xa7\x31\x78\x62\xd9\x35\x3e\xcd\x90\xe7\x68\x1a\x74\x08\xd7\x86\xc3\x7e\xa8\x6d\xd8\xfe\x8e\x92\x75\xdb\x44\x50\xb1\x59\x0c\x95\x0e\xc2\x9b\x4b\xe5\x64\x0c\xb4\x3d\xab\xe5\x81\xa2\x77\xeb\xdb\x96\xca\xd4\x52\x0a\xb9\x3a\x19\x6c\x52\x1c\xb2\x94\xee\xe1\x83\xf1\x1d\x1a\xec\x6d\xef\xb3\x64\xff\xea\x7a\xb3\xe0\x4d\xc6\xe1\xef\xff\x08\xa9\x24\x9f\x1b\xa1\x76\xa9\x8d\x62\xae\xc9\x6e\x91\x0c\x6e\x2f\xff\x7c\x33\x5f\x8c\x8f\xac\x6f\xfd\x34\xb4\xf8\x91\x62\x0f\x73\x7b\x33\x5b\x8c\x8f\x72\x8f\xa1\x41\xe5\x10\xe6\x2f\xf3\xe9\xac\xd5\x43\x83\xd2\x41\x3d\xa7\xf3\xf9\xc5\x87\xab\x69\x8b\xdb\xbe\x5e\x08\xfd\xdc\x13\xd7\xfe\x25\xbf\xe5\xaa\xab\xf4\xb0\x2d\x9b\x50\xb5\x13\x25\x5b\x60\xa5\x3d\x6c\xe0\xe7\xf5\x55\x3b\xbc\xbe\x36\xe7\xf4\x1e\xc2\x70\x88\x41\x69\x1a\x17\xa1\x10\xa5\x9f\x41\xa9\x18\x14\xd8\x45\x13\x98\xf7\x62\x70\x64\x4f\x8e\xf2\x13\xad\xac\x5b\x19\xb4\x27\x9d\x8c\xb6\x59\xdb\x64\xa6\x33\x37\x91\x7b\x9d\xf3\xf0\x52\x6d\xab\xc8\x03\xc9\x76\x07\x53\x4a\x02\xa5\xaf\xb8\x73\xd4\xeb\x48\x3b\x4e\xfe\x86\x5c\xda\x0e\x1e\xff\x47\xb7\xba\xa4\x83\x31\xb8\x4a\x33\x3f\x63\xa6\x9b\xb3\x42\x4b\xcd\x6d\xd2\x43\xc8\xdd\x51\x6f\x4b\xc7\x46\x81\x66\x4d\xeb\xf5\x14\x0c\xe0\x7c\xf9\x62\xb6\x3a\xac\xbb\x3b\x80\xbe\xa1\x99\xa0\x5e\xef\xe0\xf8\x58\x14\xc7\xf8\x45\x58\x67\x0f\x99\x19\x8d\xc0\x21\x37\xb9\x7a\x92\xbe\xaf\xd7\x0e\x2d\x64\x25\x72\x59\x6b\x70\xdc\xde\x5b\x78\xba\x43\xe9\xaf\xc2\xf0\x00\x2f\x84\x14\xf6\xae\x6d\x6e\x87\xfc\x6c\x15\xf6\x3f\xa7\x77\xc6\x6a\xff\xab\x48\x9b\xd6\x37\xa6\xf4\xa8\xc5\x83\x47\xfc\xcf\xa7\xf6\x4e\x33\x55\x96\xcd\xb0\x52\x8f\xf4\xc6\xe8\x34\xa3\xbe\xba\x2b\x49\xf1\x26\xcd\x8f\x3b\xc3\x10\xa8\xff\xf9\x44\x14\x9b\x28\x0f\x04\xd6\x6e\x0d\x7d\x3c\xde\x81\xbd\x5c\x6d\x11\xcd\xb5\xf4\x50\xb2\x1b\x8d\x32\x19\xb4\x1d\x65\x30\x84\xdc\x88\x47\x34\xec\x76\xfe\xe9\xea\xac\x16\x65\xfe\xa9\x46\xb3\x6e\xae\x8c\xf6\xa5\x1a\xf8\xff\xf2\x38\xed\x1f\xb6\xe6\x3d\x98\xbe\xd6\x1a\xa5\x28\x87\x2f\xee\x9f\xdd\x58\x9e\xe3\x7f\x07\x00\x00\xff\xff\x7d\xe6\x96\x96\x73\x13\x00\x00")

func templates_testSingletonPsql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_main_testGoTpl,
		"templates_test/singleton/psql_main_test.go.tpl",
	)
}

func templates_testSingletonPsql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6d, 0xc1, 0x99, 0x24, 0x98, 0xcb, 0x39, 0x6d, 0x51, 0xc0, 0x1d, 0x6c, 0xd8, 0x44, 0xfe, 0x54, 0xcc, 0x3, 0x28, 0x8d, 0xc8, 0x86, 0xcd, 0x6b, 0xf6, 0xf3, 0x33, 0x91, 0x7f, 0x4c, 0x2b, 0x94}}
	return a, nil
}

var _templates_testSingletonPsql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonPsql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_suites_testGoTpl,
		"templates_test/singleton/psql_suites_test.go.tpl",
	)
}

func templates_testSingletonPsql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x10, 0xc4, 0x71, 0xaf, 0xd9, 0x16, 0x41, 0x8b, 0x4b, 0xfc, 0xe8, 0xba, 0xfd, 0xfa, 0x4d, 0x2c, 0x1, 0xd1, 0x0, 0xe1, 0xb0, 0x78, 0xee, 0x7f, 0xd0, 0x65, 0xf3, 0xa1, 0x43, 0xba, 0x3c, 0xe7}}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\xcd\x6e\xdb\x3c\x10\x3c\x8b\x4f\xb1\x9f\xf1\xb5\xa0\x0a\x85\x41\xaf\x29\x7c\x70\x7e\x0e\x41\x51\xc3\x88\xe5\x73\xc1\x48\x2b\x87\x30\x4d\x0a\xe4\xaa\xb6\x2b\xf0\xdd\x0b\x52\x4e\xe2\xfc\x15\x46\xd1\xa2\xe8\xc1\x96\x48\xcc\xce\xec\xce\xee\xaa\xef\x4f\xe0\x7f\xa9\x95\xf4\x70\x36\x06\x31\x89\x6f\xe8\x45\x29\x6f\x35\xc2\xf0\x10\x53\xb9\xc6\x10\x58\xd3\x99\x0a\x08\x3d\xf5\xfd\x10\x21\x16\xed\x4c\x77\x4e\xea\x10\x16\xad\x47\x47\x9c\xe0\x43\x04\x28\xb3\x14\x65\x0e\x3d\xcb\x48\xcc\xa4\x93\x5a\xa3\xe6\x39\x63\x99\x6a\x40\xa3\xe1\x0f\x04\x97\x76\x63\xe6\xca\x2c\x3b\x2d\x5d\x08\x13\xad\x2f\xac\xee\xd6\xc6\xe7\x30\x1e\xff\x0c\x39\x73\x6a\x2d\xdd\xee\x33\xee\x1e\x02\x7a\x96\x65\x24\xe6\x2b\xd5\xf2\x51\xfc\x6f\x95\x59\x02\xa5\x32\x36\x8a\xee\xc0\x1a\xbd\x83\x76\x88\x83\x15\xee\xa0\x1a\x22\x47\x39\xcb\x02\x63\x99\x47\xac\xa3\x05\x4e\x9a\xda\xae\xd5\x77\x14\x53\xdc\xcc\x11\x6b\x9e\xb3\xec\x9b\x74\x80\x2e\xfd\xac\x63\xd9\xe9\x29\x4c\x88\x70\xdd\x12\xd0\x1d\xc2\xf5\x74\x7e\x75\x53\x82\x57\x35\x82\x6d\x40\x1a\x58\xcc\xe2\x0d\xcb\x6c\x64\x3c\xb0\xeb\xb1\x82\x3e\x24\x37\x22\xe9\xa1\xe6\x9c\x5c\x57\x11\x8f\xc9\x14\xf0\xde\x16\xf0\x86\x01\x97\xe7\xe5\xae\x45\x5f\x00\xb9\x0e\xf3\x4f\x89\xe7\xbf\x31\x18\xa5\xf7\x46\x5c\xc5\x4c\x1b\x3e\x5a\x98\x64\x01\xd9\x47\x91\xd7\x13\x02\x9f\xa4\xcf\xe0\x9d\x1f\x15\x91\x6f\xef\x4b\xdf\xab\x06\x8c\x25\x10\x53\x7b\x61\x0d\xe1\x96\x42\xa8\x68\x1b\x2b\xab\x86\xb3\x38\x97\xd5\x6a\xe9\x6c\x67\x6a\x9e\xf7\x3d\x9a\x3a\x04\x96\x0d\x90\x2f\x9d\xa7\x72\xcb\x13\xcb\x21\xc3\x8b\x8b\x5b\xab\xb4\x38\xc7\xa5\x32\x89\x43\x7b\x3c\xbc\x2b\xb7\xbc\xa2\x6d\x11\x0b\xbc\x57\x38\x0a\x94\xb3\xac\xc6\x06\x1d\xc4\xe1\xe5\x39\xf4\xf0\x15\xc6\x40\x5b\x71\x63\xb5\xbe\x95\xd5\x8a\xe7\x10\x62\x87\x1f\x7a\x61\xc5\x7e\x96\xdf\x2a\x3c\xf6\x04\x4d\x0d\x27\x21\x40\x3c\x35\x52\x7b\x4c\xa2\x05\xa4\x5c\xae\x4d\x83\x8e\xe7\x4f\x4f\xc7\xf5\xa8\x4b\xd2\xaf\x37\xe8\x45\x67\x2a\xdb\x19\x4a\x17\xcf\xa6\xec\x7e\x29\x79\x2e\x2e\x22\xe6\xc8\x52\x1e\x5d\x78\x99\x25\xbf\x97\x8d\x90\x24\x1c\x41\x1f\x9f\x40\x46\x1b\x69\x08\xac\x41\x70\x58\x59\x57\x17\xb0\xb4\x74\x36\x2a\x06\xfc\x3e\xe9\x67\xab\xb3\x98\x5d\x4e\xca\xab\xd7\x56\xe7\x77\x2c\xc7\xbe\x35\xc7\x7e\x44\x84\x10\x7f\x74\x95\x7e\x7d\xc6\xe2\x96\xff\xe5\x11\xfb\x47\x26\x2c\xb0\x1f\x01\x00\x00\xff\xff\x53\x0f\x25\xbd\xd2\x06\x00\x00")

func templates_testUpsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertGoTpl,
		"templates_test/upsert.go.tpl",
	)
}

func templates_testUpsertGoTpl() (*asset, error) {
	bytes, err := templates_testUpsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfb, 0xde, 0xdf, 0x7e, 0xe3, 0x82, 0x69, 0xda, 0x51, 0xad, 0xf5, 0x62, 0xf6, 0x3, 0xfb, 0x16, 0x16, 0xcf, 0x5a, 0xc7, 0x97, 0x5c, 0x53, 0x26, 0x8a, 0x91, 0x81, 0x86, 0x3b, 0x2c, 0x22, 0x58}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"templates/17_upsert.go.tpl": templates17_upsertGoTpl,

	"templates/singleton/psql_upsert.go.tpl": templatesSingletonPsql_upsertGoTpl,

	"templates_test/singleton/psql_main_test.go.tpl": templates_testSingletonPsql_main_testGoTpl,

	"templates_test/singleton/psql_suites_test.go.tpl": templates_testSingletonPsql_suites_testGoTpl,

	"templates_test/upsert.go.tpl": templates_testUpsertGoTpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.go.tpl": &bintree{templates17_upsertGoTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_upsert.go.tpl": &bintree{templatesSingletonPsql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_main_test.go.tpl":   &bintree{templates_testSingletonPsql_main_testGoTpl, map[string]*bintree{}},
			"psql_suites_test.go.tpl": &bintree{templates_testSingletonPsql_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
