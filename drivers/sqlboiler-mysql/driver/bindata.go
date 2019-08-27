// Code generated by go-bindata. DO NOT EDIT.
// sources:
// override/templates/17_upsert.go.tpl (7.223kB)
// override/templates/singleton/mysql_upsert.go.tpl (1.13kB)
// override/templates_test/singleton/mysql_main_test.go.tpl (5.129kB)
// override/templates_test/singleton/mysql_suites_test.go.tpl (255B)
// override/templates_test/upsert.go.tpl (1.848kB)

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

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5f\x6f\xdb\xc8\x11\x7f\x96\x3e\xc5\x9c\x90\xbb\x23\x0b\x86\x49\x81\xa2\x0f\x3e\xf8\x21\xfe\x93\x9c\x1b\x3b\x67\x5b\x71\x0d\xd4\x30\x82\x35\x39\x92\x17\x5e\xed\x32\xcb\xa5\x6d\x95\xe5\x77\x2f\x66\x96\x14\x49\x59\x92\x95\x34\x2e\xee\xc9\x22\x77\x76\xe6\xb7\xf3\x9b\x7f\x4b\x97\xe5\x6b\x78\x25\x94\x14\x39\xec\xec\x42\xfc\x8e\x7e\x61\x1e\x7f\x16\x37\x0a\xc1\xff\x89\x3f\x89\x19\x56\xd5\x90\x45\xf3\xe4\x16\x67\xc2\x2f\xd3\x86\x56\x02\xfe\x03\xf1\xb8\x5d\xe5\x0d\x72\x02\xf1\xbb\x34\xfd\xa0\xcc\x8d\x50\xf0\xba\xaa\x86\x6f\xde\xc0\x45\x96\xa3\x75\x1f\x40\x38\x87\xb3\xcc\xe5\x20\x34\x48\x4d\xef\x22\x10\x3a\x85\xd4\x20\xbf\x2b\xb2\x54\x38\x04\x63\x41\x4e\xb5\xb1\x08\x46\x43\x62\xf4\x44\xc9\xc4\xc5\xc3\x49\xa1\x13\x08\x0c\xfc\xa5\x2c\x3d\xfe\xf8\x22\x1b\x4b\x3d\x2d\x94\xb0\x55\x15\x36\x56\x02\x06\xa1\x8d\x83\xf8\x93\xd9\x37\xda\xe1\xa3\xab\xaa\xc4\x3d\x92\x2a\x7a\x88\xeb\x97\x11\x94\x25\xea\x94\x40\xd6\x96\xf7\x8d\x2a\x66\x3a\x8f\x6a\x70\xf5\x23\xdc\x18\xa9\xe2\xfa\x21\x04\xb4\xd6\x58\x28\x87\x03\x8b\xae\xb0\x1a\x4c\xec\x0d\x7b\xbb\x5d\x9b\xbc\xef\x03\xba\x83\xbd\x20\x2c\x4b\x54\x39\x32\x8e\x08\x9a\x85\x5a\xb2\x5e\xd7\x69\x55\x45\x1b\x91\x84\xc3\x6a\x38\x5c\x80\x1e\x7a\x77\x93\x03\x3b\x2e\xa7\x9f\xa7\x42\xcb\x64\xc9\xf9\xa7\xff\x9b\xf7\x81\x75\xe6\xf4\x8e\x1d\xb0\x35\x1d\xa7\x2f\xcd\x47\x39\x1c\xc8\x09\x81\xa2\xe8\xfc\x7f\x92\xf1\x1b\x1b\xfd\x69\x17\xb4\x54\x84\x62\x90\x91\x8b\x02\xd6\x77\x69\x45\x76\x68\x6d\x80\xd6\x86\xe1\x70\x50\xad\x22\x6e\x0d\x53\xab\x88\x82\x22\x97\x7a\x4a\xcf\xf8\x88\x49\xe1\x8c\xfd\x96\xc4\xe9\xa8\xce\xbe\x8f\xc5\xd3\xa7\xfe\x24\x20\xde\x77\x87\x35\xa4\x8e\x57\x9f\x52\xdb\x8a\xd7\xaf\x3a\xbb\x9e\xf7\xf5\xf6\x94\xaf\x88\xb3\x6e\x5c\x11\x8c\x97\xa3\xf5\x5e\x58\x98\xcd\xc7\x67\xc7\x2b\x9d\x79\xa1\xe5\xd7\xa2\xb1\x0a\xbb\x70\x75\x9d\x3b\x2b\xf5\xb4\xe4\x3a\x6b\x85\x9e\x22\xbc\x92\x11\xbc\x4a\x8c\xea\x54\xda\x66\x03\x59\x18\x90\xa4\x9c\xb0\x48\xec\xf5\xd1\xdb\x51\x59\xf2\x1b\x5f\xb6\x47\x91\x97\x6b\x60\xd5\xbf\x2b\x46\xbb\x88\x85\x97\x88\xb2\x31\x62\x8f\x29\x48\x4d\x52\xcc\x50\x3b\xe1\xa4\xd1\x30\x31\x16\x6e\xcd\x03\x38\x03\x99\x35\x19\x5a\x35\x87\x22\xc7\x3e\x1d\x6c\xb1\xc7\xc8\xb6\x41\xfa\xe7\x8a\xd1\x45\x9b\x90\x13\x30\xb0\xdb\x86\x53\xdd\x36\x78\x3d\x8f\x3f\xe1\x43\x30\x2a\xcb\xf8\xf4\x6e\xea\xd9\xdb\x01\x6d\xa0\x2c\x7b\x8d\x98\xdc\x75\x2f\x53\x4c\xd9\x85\x05\x9f\x76\xc4\xf1\xe7\x99\x26\x22\x15\x51\x33\x72\x72\x86\xb9\x13\xb3\xec\x8b\x97\xfa\x72\x8b\x2a\x43\x3b\x82\x18\x2a\x2f\xdd\xe6\xc8\xef\xc6\xdc\xd5\x61\xd5\xcd\xa6\xd4\xec\xe1\xc4\x58\xf4\x4e\x65\xa1\xad\x53\xeb\x69\xf2\xb4\xa7\x25\xb8\x83\x36\x16\x87\x03\xfd\xef\x03\x9c\x88\x42\x39\x1e\x44\xbe\x16\x68\x25\xe6\xf1\x27\xa3\xff\x85\xd6\xd4\x4b\x63\x24\x5a\x6b\xd2\x0f\xcc\x83\x6e\x69\xaf\x3d\x7d\x29\xdd\x6d\x2d\x1c\x81\x09\x49\xad\x4f\x8c\x67\xb4\x6e\x99\xa7\xac\x93\x1d\xa4\x50\x07\x0b\xdd\x21\x31\xfa\x76\x1d\x9f\x89\xd0\xe4\x2c\x4f\x01\x3c\x48\x77\x0b\x02\x1c\x4f\x50\xee\x56\x38\xa8\xd7\x9b\xdc\xa1\x3c\x12\x50\xb0\x66\x48\xd8\x6e\xc3\xee\x9b\x37\xb0\x57\x48\x95\x42\x22\x92\x5b\x84\x3b\x9c\x83\xd4\xaf\x95\xd4\x08\xc5\x54\x49\x35\x87\xd7\x30\x9b\xe7\x5f\x15\xdc\xe7\x90\xd1\xdf\xcc\x9a\x1b\x85\xb3\x7c\x38\xb8\x29\x26\xe4\x82\xdc\xd9\x99\xd0\x53\x85\xd4\xe4\xf6\x8a\xc9\x04\x6d\x10\xf2\x6a\x7c\x69\xa5\xc3\x31\x17\xa1\x20\x77\x36\x31\xfa\x3e\x3e\x72\x46\x04\xbd\x38\x8f\x3f\x4a\x9d\x52\xb9\xa3\xe0\xfb\x12\x41\x42\x5a\x7d\xb9\xea\xcb\xed\x1b\x95\xb3\x4b\x96\x75\x27\x7c\x9a\xf6\xf5\xde\xdc\x61\xf0\x6b\xfc\xeb\x73\x30\xfa\x65\x60\x3d\x8c\xbe\xdc\xf7\xc0\x78\xaa\xb3\x13\x9d\x3f\x40\x57\x13\x92\x1b\x54\x11\xb7\x3b\xbb\x40\xab\xf5\x42\x38\x1c\xb4\xe4\x9d\x16\x0d\x79\x37\xc5\x24\xe4\x54\x5e\x99\x16\x3e\x6d\xf7\x29\x5c\x4e\x0a\x17\x9f\x1f\x9b\xe4\x8e\x34\x71\x00\x45\x3e\x8e\x52\x32\xf4\xfc\xfe\xab\x3b\x9c\x5f\x6f\x6d\xe8\x42\x2b\x6f\x6a\x38\xa0\x3e\x48\x75\x80\x73\xc2\x67\xcf\x4f\xb5\x61\x72\x40\x33\x7c\x5a\x74\x04\xa4\xcf\xde\x51\xe7\x89\xf2\x74\x38\x18\xac\x43\xf0\x4e\xa9\x26\x4b\x37\x48\xad\xa8\x13\xdb\x49\x9b\xc2\x75\x37\xb4\x01\x41\x8f\xe1\x70\x30\xa8\xfb\xe1\xce\xee\x52\x1e\x5c\x74\x9e\x7e\xc8\x11\x4e\xad\x9c\x09\x3b\xff\x88\xf3\x8e\x30\x39\xba\xa9\x4b\xde\x7e\xa7\x28\x3d\xdf\x65\x0a\xed\xeb\x91\x69\xca\xd4\x52\xcf\x89\x20\x31\x85\x4a\xb9\xea\xdf\x70\x09\xaa\x8f\xeb\x0b\x14\x28\x99\x73\x0f\xe2\x32\x45\xe6\xa0\x5b\x6a\xc6\x34\x4f\xcf\x32\x85\xd4\xfd\x03\x8b\x2e\x6a\x93\x80\x36\x71\x34\xc4\x54\x9d\xe7\xb0\xeb\xf5\xfb\x78\x3a\xa3\x57\x27\x54\x9b\x83\x54\x0a\x85\x89\x8b\x60\xb4\x04\x6d\xd4\x34\xe2\xa6\x03\xb7\x1a\x2d\x7a\x0d\xb0\x0b\x93\x99\x8b\xc7\x99\x95\xda\x4d\x98\x81\xd1\xf8\xf0\xf8\x70\xff\x33\xfc\x9c\xc3\xfb\xf3\x3f\x4e\xe8\xbc\xc7\x67\x55\xb5\xa4\xbb\x2c\xe3\xf3\xb3\xaa\x82\xcb\xdf\x0f\xcf\x0f\xe1\xe7\x7c\xc4\xbc\xf8\x41\x2d\x8f\xff\x61\xa4\x0e\xda\x53\x1e\xa5\xa8\xdd\x59\x61\x1c\x8e\x95\x4c\xb0\x41\x1c\x1f\x9f\x45\xd0\xfc\x3e\x3f\xe3\x40\x0f\x23\x18\x45\xa3\xb0\xd1\x56\x2b\xb8\xbc\x45\x8b\xfb\x4a\x14\x39\x32\x3f\x04\x68\xe4\x0f\x7c\xee\x7f\xbe\xed\x3a\x6e\x41\xbb\x3f\xec\xbd\x50\x05\x9e\x88\x2c\x93\x7a\x1a\x71\xc2\xb5\x0d\x6f\x4f\xea\xb4\x5e\x5a\xd7\x40\x3f\xcf\x33\x8c\xd6\x95\x81\x85\xda\xd6\xc3\xf5\x90\xd0\x69\xee\xbd\xee\x4e\x35\xac\x89\x47\x3a\x30\x09\xd6\xc1\xb8\xe0\xe6\xa5\xc1\x92\x5d\x32\xb8\x02\x6a\x1f\x2b\x83\xad\x7c\x8f\x65\x37\x72\xb1\xc6\x09\x53\x76\xa4\x53\x69\x31\xa1\xb8\xf5\x2f\xfe\x49\x12\x7f\x4c\x02\x43\xed\xe7\x5e\xa8\xde\x68\xc1\x8b\xf9\x7b\x6b\x66\xcd\x11\x58\x61\x5d\x6a\x7b\x24\x85\xbe\x34\x7a\x24\x39\x5c\x5d\x4b\xed\xd0\x4e\x44\x82\x65\xb5\x98\x31\x96\x9d\xd5\x71\x64\xb3\xb1\x35\x7e\xea\xec\x7a\xd3\x1d\x1d\xcd\xac\xd8\x1b\x90\x17\xb3\x1f\x4f\xae\x07\x78\x53\x4c\x4f\x4c\x8a\x6c\x8a\xb2\xe7\x3d\x67\x8f\xd2\x41\xbb\xce\x7d\xcb\x36\x06\x38\x7f\xc3\xe7\xa5\xc9\x65\x61\x3d\xff\xd1\xfc\xbd\x30\x9c\x92\x54\x04\x0f\x2c\xc7\xdd\x8f\xf6\x1e\xe5\xbc\x3b\x48\xdc\x63\xf8\x9b\x97\x79\x8a\xe9\x61\x0b\x24\x0f\xab\xec\x37\x97\x35\x2a\xb6\x89\xd0\xc7\x22\x77\xbe\xe7\x1c\x1d\x74\x6f\x5d\x4b\x2b\xf5\xed\x8b\xef\x5e\xab\x96\x56\xfb\xd6\x62\xce\x73\x69\x3d\x5c\xd3\x88\xcc\x57\x91\xa0\x83\xda\xc3\x8b\xe3\x38\x64\x2d\xad\x7f\xd6\x6d\xae\x2d\x04\x3c\x7f\x6f\x50\x54\x1f\xb4\xa7\x73\x35\xcc\x2f\x4d\x42\x7e\x1b\xc0\xa7\xdb\xbe\x1d\x5a\x73\x1d\x58\x91\xb2\xfd\x16\x46\x57\x6f\xba\x77\xfb\xea\xb8\xa9\x91\xd1\xfc\xb5\xdc\x31\xba\x09\xb0\x9a\x40\x4a\x4c\x45\x6f\x0f\x40\x6a\xf7\xf7\xbf\xf5\xc0\xd1\xa2\x9f\xcf\x4f\x44\x06\x57\xd7\x45\x2d\x42\xef\x9b\xf2\xcc\x63\x67\x3f\xa5\x37\xe4\xf4\xa2\x53\x4f\x8d\x33\xc0\x53\x54\x7d\x23\x7b\x16\xa9\x47\xd9\xf8\xde\x47\x49\xdc\x11\x4b\x69\xdc\x5b\xeb\xce\x43\x6b\xc7\x73\x9d\xbc\x17\x52\x35\x96\x5e\x25\x46\xf1\xf7\x5c\x1e\xc6\x52\x7c\x6c\x92\xe0\xf4\x23\xce\x17\x77\xf9\xb7\x2d\x65\x4b\x5f\x28\xf8\xe3\x19\x0f\x05\x0b\x4d\x3d\xd1\xcf\xd2\x29\x3f\x73\xd6\xd5\x7b\x49\x9a\x64\x4d\xec\x71\x78\xd9\xaa\x02\x1e\x50\x13\xa3\x62\xaa\xfc\x55\x15\xf8\x53\xfb\x93\xd5\x3c\x71\x5d\xfc\xe5\x97\xf5\x1e\xfe\x2b\xad\x2e\xaf\x5c\xbd\xbd\xa6\xb5\xcd\xad\xe4\x6a\xd4\xba\xa5\xaa\x46\xd7\xeb\xa9\xea\x5d\x69\x17\x31\xf2\x62\x1d\xae\x3b\x45\xfd\x80\x94\xb1\xe8\xac\xc4\x7b\x6c\x6e\x9f\xdc\x3f\xf2\x35\x29\x04\x74\xdc\x5e\xb8\x6f\xea\x82\xdb\x74\xd3\xa8\xcd\xaa\xf0\x05\xfa\x53\x33\x0d\x6e\xd1\xa2\xba\xc7\xf2\x75\xea\x05\xbb\xd5\x5a\x5c\x0f\xcf\xa0\xe9\xf4\xae\x15\x9e\xea\x14\x63\x56\x7f\x6e\x1e\x82\xbe\xc1\xa7\x9a\xe3\x71\x22\x78\x5e\xa3\xc1\xc2\x9b\xea\x9e\x7a\x85\xca\x15\x35\xfe\x5b\xd5\x37\xe5\xff\x07\x04\x70\x66\xb2\x82\x3f\x7d\xa5\xfe\x72\xb6\x39\x82\xa9\xe0\x75\x13\x78\xe7\xc9\xdd\x74\xbb\xcb\x6e\x73\xa9\xde\x42\x9c\x2f\xd1\xb0\xeb\x3d\xb5\xb5\x81\xc5\x65\x7a\xb0\xe1\xab\xdd\xe2\x1f\x50\xa9\x79\x37\x71\x68\xbf\xeb\x8b\x5d\x5d\xc0\x3a\x53\x07\x2b\xd5\xd4\x1e\xba\x5f\x8e\xff\x1b\x00\x00\xff\xff\x14\xbb\x48\x0d\x37\x1c\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xce, 0xf3, 0xf1, 0xdb, 0x35, 0x64, 0x5e, 0x48, 0x85, 0x22, 0xfa, 0x70, 0x67, 0xcb, 0x9d, 0x27, 0x99, 0x3b, 0xdb, 0xf5, 0x98, 0xe6, 0xe8, 0xaf, 0x59, 0x80, 0x44, 0xdc, 0xb0, 0xf1, 0xbf, 0x41}}
	return a, nil
}

var _templatesSingletonMysql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xcd\x8e\xd3\x30\x10\xc7\xcf\xf6\x53\x0c\x91\x56\x8d\x25\x2b\xcb\x5e\x91\x7a\xd8\xa5\x65\x15\x28\xfd\x2e\x08\x21\x0e\x6e\x3d\x6e\x2d\xa5\x4e\xf1\x47\xa1\x42\x7d\x77\xe4\x24\x6d\xb3\x4b\x41\x1c\xf6\x92\x8c\x3d\x33\x7f\xcf\x6f\x66\x6e\x6f\x61\x19\x74\x21\x17\x3b\x87\xd6\x4f\x02\xda\xc3\xc7\xc3\x6c\x32\xa8\x6f\x1d\x08\x88\x07\xe7\x85\xc7\x2d\x1a\x0f\xce\x5b\x6d\xd6\x10\x5c\xfc\xfa\x0d\x42\xa8\x12\x7b\xc2\x0b\xd8\xd9\x72\xaf\x25\xca\x8c\xaa\x60\x56\xd7\x75\x53\xa9\x05\x48\xab\xf7\x68\x5d\xd6\xd3\xa2\xc0\x95\xe7\xe0\xc5\xb2\xc0\xa1\xd8\x62\xa3\xcf\x21\xec\xa4\xf0\xc8\xe1\xc7\x46\x7b\x2c\xb4\xf3\xf0\xf5\x5b\xed\x63\xa7\x1a\x7e\x51\x72\xf1\x76\xe3\xed\x56\x98\x75\x81\x59\x2e\xd1\xf8\x49\x28\x3d\xce\x0a\xbd\xc2\xf8\x64\x36\x98\x70\x88\xff\xe9\xa4\xa5\xc9\x28\xb9\xbc\x7c\x5d\xe1\x8f\xe4\x73\x02\xa3\x94\x2c\x83\x82\x37\xed\xc4\x47\xf4\x0f\x41\x29\xb4\x29\xa3\x44\xa2\x42\xdb\x72\x8e\xc3\xc9\xb9\x0c\x2a\xa6\xef\x85\x85\x55\x59\x84\xad\x71\x0d\x14\x25\x5a\x41\x81\x26\xbd\xd4\x08\xaf\xba\xf0\x3a\xc2\x92\x53\x68\xb7\x09\x76\xd9\xfb\x52\xb7\x42\x39\x24\x3c\x61\x94\x1c\xe9\x59\xa6\x6e\x23\x83\xee\x49\x43\x6d\x7d\xf6\x6e\x67\xb5\xf1\x2a\xa5\x84\x44\x02\x1e\xff\x49\x3e\x9c\xf5\xa7\x73\xc8\x1f\x87\xa3\x69\x1f\xf2\xe1\x7c\x04\x37\x0e\xd2\x1b\xc7\xe0\xd3\xfd\x60\xd1\x9f\x55\x76\x52\x05\x9f\x7b\x50\x9d\x9a\xb2\x2a\xbb\x05\x5b\x88\x15\x6e\xca\x42\xa2\x75\x55\x13\x17\x0e\x73\x23\xf1\x67\xdb\xc1\x9f\xb1\x72\xb8\xe3\x70\xc7\xa2\x14\xa3\x84\x58\xf4\xc1\x1a\x58\x06\x95\xcd\x2a\xe2\xb4\xa1\x7b\x46\xd1\x40\x9c\x19\xfe\x52\x3c\x8c\x86\xd0\x5b\x8c\x07\xf9\xdb\xfb\x79\x1f\x3e\xf4\xbf\xc0\x62\xdc\x8b\x66\x45\xf5\x04\xaa\xc5\xf4\x62\x48\x71\xe2\xaa\xb4\xa0\x39\xec\xe3\xd6\x58\x61\xd6\xd8\x2c\x7a\x35\x1b\xad\x40\x5f\xa6\x1d\xa9\xb2\xcf\x56\x7b\x7c\x38\x78\x4c\x3b\xbc\x13\x5b\x72\xa4\x84\x7c\x8f\x8b\x29\x9f\x2e\xde\x3f\x36\x76\xcf\x68\x4b\xac\x69\x64\xad\x71\xcd\x93\x40\xb7\x69\x5a\x9a\xfc\x67\x66\x5d\x20\xeb\x34\xd3\xb9\x36\xb6\x23\xfd\x1d\x00\x00\xff\xff\x4c\x0d\x4e\x35\x6a\x04\x00\x00")

func templatesSingletonMysql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonMysql_upsertGoTpl,
		"templates/singleton/mysql_upsert.go.tpl",
	)
}

func templatesSingletonMysql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonMysql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/mysql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3e, 0xcc, 0xa0, 0x4c, 0x62, 0x37, 0xa8, 0x14, 0xaa, 0x37, 0x45, 0xfd, 0x34, 0x76, 0xe0, 0x86, 0x6, 0x53, 0x4, 0xc, 0xf3, 0x3, 0xab, 0xc9, 0xcc, 0x8f, 0xb8, 0xe0, 0xcd, 0xca, 0x43, 0x5}}
	return a, nil
}

var _templates_testSingletonMysql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5b\x6f\xe3\x36\x16\x7e\x96\x7e\xc5\xa9\x81\xe9\x4a\x53\x85\x19\xa0\xc0\x3e\xa4\x10\x82\x89\xe3\x14\x41\x27\x37\x3b\xbb\xc5\xa2\x29\x5a\x46\xa2\x13\x22\x12\xa9\x90\x54\x1c\x6f\x90\xff\xbe\x38\xa4\x2e\x94\x63\x79\x33\xbb\xf3\xd8\xa7\x44\xe2\xc7\x73\xf9\xce\xe1\xe1\x27\x3f\x51\x05\xea\xee\xf9\x6c\xbd\xb8\xfa\xf2\xc0\xd6\x90\x82\x62\x77\xec\xb9\x22\x67\xb5\x36\x53\x59\x56\xbc\x60\xd1\x9f\xd1\x61\x19\x47\x51\x72\x23\xe2\xc3\x1b\xfd\xc3\xf4\xe2\x7c\x71\x3d\xff\x7c\x7a\x7e\x4d\x3e\x1e\x9e\x5c\xcc\x67\xa7\x3f\x9f\xc3\x2f\xb3\x7f\x91\x8f\x87\x37\x22\xfe\xe1\xcf\x38\x0c\xcd\xba\x62\x50\xae\xf5\x63\x71\xcd\xb4\x61\x0a\xb4\x51\x75\x66\xe0\x25\x0c\xf2\xdb\xa9\x14\x02\x3e\xea\xc7\x82\x1c\x1f\x85\xf8\xe2\x9c\x96\x0c\x10\xc2\xc5\x5d\x18\xdc\x4b\x6d\x00\xfa\xe7\x5a\x33\xe5\x3f\x57\x54\x6b\xff\x59\xeb\xa2\x94\x39\xeb\xd7\xa5\xb2\xfb\xb9\x30\x61\x18\xc8\xca\x70\x29\x4e\x78\xd1\x01\xc2\xc0\x30\x6d\x8e\x8f\xac\xd7\xce\xc8\x03\xaf\x16\x57\x5f\xa6\x65\x0e\xb7\x52\x16\xe1\x6b\x18\x2e\x6b\x91\x01\x17\xdc\x44\xb1\x8b\xfb\x8c\x72\x01\x29\x7c\xef\xe5\xf5\xf2\xda\x21\xa3\x12\x3e\x7a\x2b\x31\x68\x66\xea\x2a\x8a\x81\x29\x25\x15\x5a\x40\xae\x99\x52\xee\x45\x18\x06\x4f\xbc\x62\x8a\x2c\x98\x39\x66\x4b\x5a\x17\x26\x9a\xd8\xfd\xa4\x49\x68\x92\xc0\xc4\xa8\x9a\x4d\xe2\x71\x28\xe6\x3a\x49\xe0\xc7\x1f\x3f\xfd\x3d\x0e\xc3\xa0\x24\x0d\x99\x29\xb8\x1d\x3f\x33\xb3\xb0\x19\xb6\x1b\xf2\x5b\x41\x4b\x6b\xb2\x24\x96\xe8\x51\x24\xae\x3a\x9c\x2d\xc0\x28\x0e\x57\x1d\xce\x16\x66\x14\x87\xab\x0d\x0e\x0b\xe4\xe1\x4e\xc5\x30\x1f\x0b\x6a\xab\x3a\x6a\xaf\x65\xc9\xa2\xbd\x8a\x8e\x6e\x40\x8c\x9f\xbe\x57\x72\x6f\xcf\x91\x94\x45\xe7\xe2\x81\x57\xfa\xb1\xc8\xca\x7c\x82\xec\x62\xed\x52\x78\xa2\x05\x25\x47\xec\x8e\x8b\x7f\xd2\x82\xe7\x14\xdb\x2b\x8a\x49\xf3\xc0\xa2\x30\x08\x2c\xc4\x39\x3f\x97\x66\x56\x56\x66\x1d\x39\x1a\x13\x18\xb0\x96\x8c\x82\x91\xfd\x0e\xec\x4a\xd1\x81\xcf\xa5\x89\xec\x3f\xb3\xc7\x9a\x16\x3a\x72\x8c\x26\xf0\xa9\xdb\xe0\x68\xdc\x61\xde\xb5\x49\x87\x6f\x69\x19\xdf\xd0\xb0\xdd\xed\xe8\xd8\x4f\xc2\x20\x26\xd3\x7b\x96\x3d\x44\xc8\x11\x5f\xda\x16\xff\x2e\x05\xc1\x0b\x6c\xfa\x40\x31\x53\x2b\x81\x6f\xc3\xe0\x35\x0c\x83\xfd\x7d\x98\x2a\x46\x0d\x03\x0a\x8a\x8a\x5c\x96\xfc\xdf\x2c\x87\xfc\x16\x30\x06\x62\x4d\x14\x4c\x44\x7e\x51\x63\x48\x53\xf8\x64\xcd\x6d\xd4\xba\xb3\x40\x16\x86\xde\x16\xcc\x2d\x74\x19\xc6\xce\x67\x13\x55\x0a\x25\x29\xe9\x03\xbb\xe8\x66\x42\x14\xff\x34\x1e\xaf\x54\x9a\xfc\xaa\x68\x15\x31\x85\x85\xcb\x64\x5d\xe4\xe2\x6f\x06\xd0\x04\xb8\xb9\x02\x4b\x5e\xd8\x76\x6a\xbc\x7c\x37\x68\x2b\x34\xe7\xb9\xce\x95\xac\xae\x6d\xf0\x5b\xdc\x0e\x78\x0a\x5e\x87\x3b\x33\x4b\xd8\xbb\xf7\x86\x41\x90\xd7\x65\x85\x21\x1c\xa4\xc0\x9e\x59\x46\xa6\xb2\x2c\xa9\xc8\x9b\xce\xc6\xd5\x49\x82\x21\xb9\x71\xa2\x1d\x17\x09\x4c\xf6\xf6\x84\xdc\xcb\xa9\xa1\x6e\xb9\x25\x31\x70\x11\x8c\x5b\x1c\xb3\x86\xa6\x6e\xa9\x66\x76\xdd\x2b\x28\xc6\xa8\x12\x58\xa1\x39\x2e\xc9\x25\xaf\x58\x14\xf7\x71\x2f\x4c\x8e\x39\x1e\xa4\xf0\xfd\xed\xda\x30\x4d\x8e\xea\xe5\xd2\x8e\x5b\x2f\x94\x71\x50\x6f\x88\x2c\x4c\x2e\x6b\x1c\x37\xab\xe1\x4b\x47\xed\xc0\x5d\xe8\x1b\x47\x8c\x1d\xf7\x82\xad\x4e\x7e\x61\xeb\x63\xa6\x8d\x92\x6b\xa6\x22\xef\xba\x4c\x40\xc5\x9b\x9b\x9c\xe1\x8d\x20\x43\xbf\x9e\x7d\x14\x54\x99\xdd\xe5\xdc\x68\xc1\x25\xe5\x05\xcb\xc1\x48\xd0\xb8\x17\xba\x62\x42\xe6\xaa\x81\xad\x38\x6c\x1e\x3f\xb6\x6f\xe2\x6e\xc3\xd5\xb6\xc4\x7e\xa5\x7c\xab\xa3\x65\x69\xc8\xa5\xe2\xc2\x14\x02\x3d\xc4\x9b\xef\x06\xd5\x68\x66\x50\x14\xc7\xef\x8c\x71\x45\xb9\x81\xa5\x54\xa3\xac\x84\x41\xf0\x07\x36\x02\x99\x16\x52\xb3\x28\x86\xfd\x7d\xf8\xbc\x44\x75\xd2\x9e\x16\xae\x21\x97\x82\x25\x90\x21\x02\xcc\x3d\x83\x95\xe2\x86\x01\x13\x39\xc8\xa5\x7d\x51\xf1\x8a\x85\xdb\x19\xfe\x5f\xf3\xde\x68\x96\xff\x33\xf3\xcd\x5e\xc0\xc4\x1b\x23\x82\x17\x3b\xf4\x8a\x2e\xce\x64\xce\x22\x4f\x4c\xc5\xcd\x5f\x4c\x43\xaf\xb8\xc9\xee\xc1\xae\xbe\x84\x41\x46\x35\x6b\xf4\xc9\x41\x3f\x35\x27\xf3\xd9\xd5\x3f\x4e\xe7\xb3\xe3\x49\x8b\x58\xd2\x42\x0f\x21\xc7\xa7\x8b\xcf\x47\x5f\x2c\xa4\x19\x18\xfe\xea\xe5\x7c\x76\x32\x9b\x3b\x0b\x3b\xc4\xd5\x70\xd4\x78\x61\x36\x76\x90\xde\x45\x85\xfc\x2e\x23\x1c\x43\x0d\x7c\x0f\xe7\x75\xfa\x41\xdb\x71\xd4\x4b\xc3\x78\xdc\xd1\xe6\x7d\xd1\xcb\x39\x53\x56\x09\x34\x03\x88\xcb\xda\xf0\x82\x5c\xb3\xb2\xb2\xb0\x09\x8a\x37\x67\xbf\xbd\x21\x76\xdd\x8c\xa3\x95\x75\x9d\xb1\xf5\xb2\xd1\xd7\xd3\x4b\x74\x6d\x09\x0e\x83\x3f\x92\xa6\x1d\xa5\xc6\x93\x6e\x1a\x0d\xe1\x1c\x4b\x4d\x4e\x35\xde\xe6\xcf\x5c\x1b\xdb\x82\xee\x6e\xb2\x36\x52\xc0\x2a\x86\xc1\x2b\xb0\x42\x33\xf8\x8a\x38\xed\x8d\x08\x42\x1a\x9c\x0f\x06\xca\x4e\x33\x62\x80\x58\x81\x93\xaa\xe9\x70\xcb\xd5\xe4\xb7\xac\xe0\x4c\x98\xdf\x11\xd2\x2f\x2f\x9b\x55\xdc\x9c\x7e\xd0\x37\xc2\x16\xa7\x09\xfe\x2d\x0c\xb5\x4d\xfa\x21\x6f\x60\xf8\xb4\x15\x86\x02\xab\xb7\x86\x4f\xb1\x27\x2d\x50\x8c\xc6\x98\xa3\x13\x15\x5b\xbc\x50\xad\x57\x52\xe5\xbd\x09\xbb\x05\x53\xdb\x82\xd6\xba\xd8\xc3\x83\xd1\xa3\xbb\xc3\xd4\x2a\xa5\xd8\xb9\x77\x94\x0f\x7d\x76\xfc\x54\x4a\x1a\x99\xc9\x22\x35\x59\xb5\x8b\xc6\x6e\xc0\xfd\xc5\xe4\x57\x30\xe9\x1f\x78\x6c\xfa\xb2\x22\x56\x2b\xc6\xfd\x7c\xc4\x77\xcd\xe5\x30\x3e\x11\x86\x62\xac\x9f\x07\x38\x7a\xf1\x3c\xfa\x93\xa7\x39\xbf\xad\x0a\x82\x0f\xfa\xa7\x37\x4a\xa8\x75\x5e\x12\x55\x8b\x69\x99\x47\xfa\xb1\x68\x75\xf6\x64\x47\x1c\xbe\x9c\xdc\x1d\x05\x22\xfb\x18\xf0\x80\xe3\x1c\xd0\xdf\x34\x1a\xc3\xa8\xca\xe5\x4a\xf8\xb1\xf0\xa5\xd5\x90\xf6\x7b\xff\xed\x3c\x69\x97\x3a\xc6\xff\xab\x88\x3e\xf8\x7a\x15\xed\x5d\x7e\x52\x93\x39\x2b\xe5\x13\xb6\xd2\xbb\x46\x7f\x4b\x00\x0a\xc1\xa4\xbd\x55\x9b\xab\x26\x01\xaa\xee\x34\x10\x42\xda\x9b\xb2\xcb\xda\x2e\xa4\x40\xab\x8a\x89\x3c\xfa\xed\x77\x07\x78\xd9\x94\xc7\xaf\xce\x04\x21\x04\x1b\x30\xdb\xa2\xac\x1b\x8f\x1e\x0e\x61\x9d\x30\x75\x76\x35\x39\x67\xab\x39\xa3\x39\x53\x2e\x52\xb4\xa6\x9d\xe8\xdd\x26\x9f\xf5\xb8\xb2\xce\x7c\xb9\xec\x4c\x74\x2f\xdd\xdd\xe2\x36\x87\x5e\x3d\x70\x79\x5e\x8b\xb7\xa5\xf0\xf5\x4d\x7b\xa1\xa9\x5a\x08\x2e\xee\x0e\x26\x1d\x9b\x2e\xb7\x78\x08\x77\xae\x7d\x15\xb4\xb1\xba\xa1\x91\x36\xbf\x30\xdf\x23\x76\x32\x29\xb0\x55\xa3\xe6\x67\xa8\xc4\x95\x2f\x1e\xef\xda\x8d\xa6\x4d\xac\x79\xeb\x6e\xf8\xb3\x4e\xd0\x23\x1a\xce\x1e\x0b\x72\x51\x31\xd1\x7f\x28\xe5\x8a\x3f\x31\x45\xec\x47\xc4\x51\xcd\x8b\xfc\xaa\x66\x6a\xdd\x24\xd4\xfe\x4e\xe0\xc6\xe4\xf0\x74\xb6\xd3\xbc\x1d\xd7\xcd\x78\xf4\x86\xe2\xb0\x06\x3d\x11\xc9\x1b\x76\x86\x89\xbc\x86\xff\x09\x00\x00\xff\xff\xdd\x6c\xf4\x1c\x09\x14\x00\x00")

func templates_testSingletonMysql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_main_testGoTpl,
		"templates_test/singleton/mysql_main_test.go.tpl",
	)
}

func templates_testSingletonMysql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0xa, 0x66, 0xf8, 0x27, 0x50, 0x1f, 0xc8, 0x5, 0xd5, 0x88, 0xc1, 0x79, 0xe3, 0x28, 0xff, 0xb6, 0x1b, 0xf9, 0xfa, 0x1, 0x4b, 0xf0, 0xeb, 0x10, 0x5, 0xf8, 0xdc, 0xfe, 0xba, 0x82, 0x80}}
	return a, nil
}

var _templates_testSingletonMysql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonMysql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_suites_testGoTpl,
		"templates_test/singleton/mysql_suites_test.go.tpl",
	)
}

func templates_testSingletonMysql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x10, 0xc4, 0x71, 0xaf, 0xd9, 0x16, 0x41, 0x8b, 0x4b, 0xfc, 0xe8, 0xba, 0xfd, 0xfa, 0x4d, 0x2c, 0x1, 0xd1, 0x0, 0xe1, 0xb0, 0x78, 0xee, 0x7f, 0xd0, 0x65, 0xf3, 0xa1, 0x43, 0xba, 0x3c, 0xe7}}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x62\xd6\xd8\x5d\x50\x0b\x85\xd9\x5e\x53\xf8\xe0\x7c\x1c\x82\xb6\x86\x1b\xcb\xe7\x82\x91\x46\x0e\x61\x9a\x54\xc9\x51\x6d\x57\xe0\x7f\x2f\x28\xd9\x8e\x13\x3b\x6d\x0e\xed\x21\x07\x7f\x70\xf0\x66\xde\x7b\x33\x1c\xb6\xed\x19\xfc\x2d\xb5\x92\x1e\x2e\x86\x20\x46\xf1\x1f\x7a\x91\xcb\x7b\x8d\xd0\xff\x88\xb1\x5c\x62\x08\xac\x6a\x4c\x01\x84\x9e\xda\xb6\xcf\x10\xb3\x7a\xa2\x1b\x27\x75\x08\xb3\xda\xa3\x23\x4e\xf0\x5f\x04\x28\x33\x17\x79\x0a\x2d\x4b\x48\x4c\xa4\x93\x5a\xa3\xe6\x29\x63\x89\xaa\x40\xa3\xe1\xfb\x02\xd7\x76\x65\xa6\xca\xcc\x1b\x2d\x5d\x08\x23\xad\xaf\xac\x6e\x96\xc6\xa7\x30\x1c\xfe\x0c\x39\x71\x6a\x29\xdd\xe6\x03\x6e\xf6\x09\x2d\x4b\x12\x12\xd3\x85\xaa\xf9\x20\x7e\xd7\xca\xcc\x81\x3a\x1b\x2b\x45\x0f\x60\x8d\xde\x40\xdd\xe7\xc1\x02\x37\x50\xf4\x99\x83\x94\x25\x61\xaf\x6c\xb9\x99\x7e\xfe\x78\xe0\xef\x91\x72\x66\xd4\xd7\x06\x0f\xf5\xfd\xff\x4b\x4e\x63\xa1\xe9\xd2\x76\x64\x40\x16\x0a\x6b\x2a\xad\x0a\x02\x6b\x7a\x6e\x96\x78\xc4\x32\xb6\xdf\x49\x53\xda\xa5\xfa\x8e\x62\x8c\xab\x29\x62\xc9\x53\x96\x7c\x93\x0e\xd0\x75\x1f\xeb\x58\x72\x7e\x0e\x23\x22\x5c\xd6\x04\xf4\x80\x70\x3b\x9e\xde\xdc\xe5\xe0\x55\x89\x60\x2b\x90\x06\x66\x93\x18\x61\x89\x8d\x15\x4f\x5a\x69\x7b\xbf\xb1\xe8\x21\xe7\x94\x5c\x53\x10\x8f\x62\x32\xf8\xd7\x66\xf0\x42\xf3\xaf\x2f\xf3\x4d\x8d\x3e\x83\x4a\x6a\x8f\xe9\xfb\xae\xd0\x5f\x43\x30\x4a\x6f\x3b\x72\x13\xa5\x56\x7c\x30\x33\x5d\x2f\xc8\x3e\xb2\x9c\x56\x04\xbe\xe3\xbe\x80\x7f\xfc\x20\x8b\xf5\xb6\x8d\x69\x5b\x55\x81\xb1\x04\x62\x6c\xaf\xac\x21\x5c\x53\x08\x05\xad\xa3\xb5\xa2\x3f\x8b\x4b\x59\x2c\xe6\xce\x36\xa6\xe4\x69\xdb\xa2\x29\x43\x60\x49\x0f\xf9\xd4\x78\xca\xd7\xbc\xab\x72\x58\xe1\x28\x70\x6f\x95\x16\x97\x38\x57\xa6\xab\xa1\x3d\x1e\xc6\xf2\x35\x2f\x68\x9d\x45\x83\x3b\x86\x57\x81\x52\x96\x94\x58\xa1\x83\xb8\x39\x3c\x85\x16\xbe\xc0\x10\x68\x2d\xee\xac\xd6\xf7\xb2\x58\xf0\x14\x42\x1c\xf1\x7e\x18\x56\x6c\x17\xe9\x25\xe3\x71\x28\x68\x4a\x38\x0b\x01\xe2\xa9\xe3\xbf\x35\x15\x3a\x9e\x3e\x3d\xbd\x6e\x2e\x4d\x47\x77\x7a\x28\x47\xd3\x28\x6c\x63\xa8\x0b\x3c\xbb\x5a\xbb\x57\x80\xa7\xe2\x2a\x62\x5e\x29\xff\xd1\xf9\xb1\x4a\xbe\xa3\x8d\x90\x8e\x38\x82\xde\x3d\x81\x0c\x56\xd2\xc4\x35\x42\x70\x58\x58\x57\x66\x30\xb7\x74\x31\xc8\x7a\xfc\x56\xf4\xb3\x7d\x99\x4d\xae\x47\xf9\xcd\xa9\x7d\xf9\x6d\x1b\xf1\x22\xec\xe8\xd5\x12\x42\xfc\xd1\xf5\x79\x7b\xf7\xea\x8d\x5c\xab\xc0\x7e\x04\x00\x00\xff\xff\x63\xfd\x7b\x9f\x38\x07\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1f, 0xec, 0xa5, 0xcb, 0xb0, 0x71, 0x0, 0x4c, 0x49, 0xa2, 0xc9, 0x7f, 0x22, 0x1e, 0x72, 0x79, 0x7e, 0xfb, 0x4c, 0xfc, 0xf6, 0xab, 0x29, 0xd3, 0xef, 0x4b, 0xb4, 0x26, 0x7d, 0xd4, 0x7b, 0xfc}}
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

	"templates/singleton/mysql_upsert.go.tpl": templatesSingletonMysql_upsertGoTpl,

	"templates_test/singleton/mysql_main_test.go.tpl": templates_testSingletonMysql_main_testGoTpl,

	"templates_test/singleton/mysql_suites_test.go.tpl": templates_testSingletonMysql_suites_testGoTpl,

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
			"mysql_upsert.go.tpl": &bintree{templatesSingletonMysql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_main_test.go.tpl":   &bintree{templates_testSingletonMysql_main_testGoTpl, map[string]*bintree{}},
			"mysql_suites_test.go.tpl": &bintree{templates_testSingletonMysql_suites_testGoTpl, map[string]*bintree{}},
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
