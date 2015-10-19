// Code generated by go-bindata.
// sources:
// migrations/20150927-core.sql
// migrations/20151002-mingb.sql
// migrations/20151006-work-unit-type.sql
// migrations/20151013-index.sql
// migrations/20151014-index.sql
// migrations/20151019-worker-mode.sql
// DO NOT EDIT!

package postgres

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _migrations20150927CoreSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x56\x5d\x73\xa2\x3c\x14\xbe\xe7\x57\x9c\x3b\xf5\x7d\xeb\x1f\xa8\x57\x29\xa4\x2d\xb3\x08\x6e\x80\xee\xb8\x37\x4c\x2a\xa9\xcd\xac\x06\x1a\x42\x3f\xfe\xfd\x06\x14\x41\x3e\x5c\x75\x66\x2f\x96\x8b\x0e\x4d\xce\xe7\x73\x9e\xf3\xe0\x74\x0a\xd3\xff\xa6\xb0\x4d\x62\x76\x0b\xd9\xdb\x66\x56\xfc\x99\xa6\x32\x89\xf3\x95\xba\x85\x34\xc9\xd4\x5a\xb2\xac\x30\x32\xb4\xed\xff\x5b\xbe\x96\x54\x31\x08\x53\xc3\x24\x18\x05\x18\x02\x74\xe7\x60\x10\x74\xcb\xb2\x94\xae\xd8\xd8\x80\xdd\xc3\x63\xf0\x31\xb1\x91\x03\x0b\x62\xcf\x11\x59\xc2\x37\xbc\xbc\xa9\x6e\x0b\x7b\x78\x42\xc4\x7c\x44\x04\x42\xd7\xfe\x1e\x62\x70\xbd\x00\xdc\xd0\x71\x8c\xc9\xcc\x38\x8e\xfe\x91\xc8\x5f\x51\x96\xb2\xd5\x05\xd1\xcb\x6a\x22\x6d\x67\xbb\x01\x7e\xc0\xa4\x0e\x0f\x3d\x0f\xc1\xf7\x98\x60\xd7\xc4\x7e\xa3\x15\x1e\x4f\xc0\x73\xc1\xc2\x0e\xd6\xa5\x98\xc8\x37\x91\x85\xfb\x5b\xa8\x82\x1f\x6e\x63\xaa\x28\xdc\x2d\x03\x8c\xba\x77\xa9\xe4\x89\xe4\xea\xab\x53\xda\xc1\xe2\x83\xf1\xf5\xab\x1a\xbe\x4f\x69\x9e\xb1\x18\xee\x3c\xcf\xc1\xc8\xed\xde\xaf\x12\xa1\xb8\xc8\x93\x3c\x3b\x61\x43\x45\xf4\xcc\xa2\x73\x4c\xb9\x50\x4c\xbe\xd3\x4d\x59\x10\x79\xd2\xa8\x77\x4c\x04\xfb\x54\xcd\x58\x81\x3d\xc7\x7e\x80\xe6\x0b\xf8\x61\x07\x8f\xe5\xbf\xf0\xd3\x73\x6b\xf8\xb6\xf4\x33\x92\xb9\x10\x5c\xac\x87\xfb\x2c\x8c\xa8\x52\x6c\x9b\xaa\x2c\x92\x4c\xe5\x52\xb0\xee\x44\x8f\x8b\x38\x90\x25\x3a\x3d\xa1\x96\x71\x2a\x59\x99\x66\x18\x05\xd3\x73\xfd\x80\x20\x9d\xbd\x26\x64\x94\x0b\xfe\x96\xb3\x5d\xaa\x1d\x8f\xc7\x4d\xf6\xdd\x94\x34\x99\x0c\x50\x5a\x3b\xab\x33\x29\x5d\x67\xbc\x82\xd2\xf5\xfe\xfc\x75\x4a\x5b\x5e\x58\xb4\xb7\x20\xd8\xb4\x7d\xdb\xeb\x81\x91\xae\x14\x7f\x67\xd5\x58\x1b\xed\x0c\x02\x5d\xc0\xd4\x07\x74\x13\x93\x93\x40\x33\xf9\xef\x08\x47\x4a\x25\x13\xf5\xea\xb7\xa6\xa8\x3b\x39\x0e\xee\xe3\x5e\x78\x87\x59\x5c\x48\xfd\xf0\x06\x9d\x9a\x31\xfb\x4c\xb9\x16\x7f\x9e\x88\xc1\xed\xee\x3a\x6d\x68\xa6\x67\x97\xc6\xc5\x47\xe3\x7c\xaf\x16\x03\x98\xbc\x66\xcf\x96\x0b\x0c\x15\xcb\x32\x45\x95\x56\x25\xe4\x03\x76\xc3\xf9\x78\x94\x32\x11\x6b\xe5\x19\xdd\xc0\xa8\x6c\x8b\xc5\xc5\xeb\x0b\x17\x3c\x7b\x2d\xde\xab\x42\x46\x2f\x94\x6f\x76\x97\x5a\x7e\xe4\x17\x7d\xde\xb0\x51\x87\x63\xfb\x2c\x97\xac\x72\xc9\xe9\x6b\x57\xb9\xd4\x8d\xd3\x24\xdb\xc3\x76\x56\x86\x3f\x71\xac\x1d\x7b\x0f\x66\x0b\xdb\x2a\xbe\xf6\xba\x47\xa1\x13\x40\x0d\x72\x97\x5d\xcd\x58\x52\x45\x8a\x6f\x2f\xa1\x87\x8e\x7b\xda\xa5\x87\xb3\x97\xe6\x18\xd8\xa3\x43\x77\x01\x09\x71\x49\x37\xe4\xe8\x6f\x62\x5b\xd5\xf7\x51\x90\x65\xf5\xab\x59\x4b\x04\xf5\xb7\x95\xc7\x7b\x9f\x7b\x8f\x60\xfb\xc1\x2d\x28\x03\xe3\x8e\x58\x4e\x9a\xd3\xaa\x78\xa7\x8f\xf7\xce\x5d\x61\xd0\x15\x36\x7f\xb8\x59\xc9\x87\xe8\x2f\x19\x2c\xe2\x2d\xce\xaf\x76\x66\x94\xf6\x47\x0b\x50\x9d\x75\x37\xef\xc8\x7a\x47\xb2\xce\x51\x99\xab\x7b\x5a\x08\xfc\xd1\xe9\x61\xef\x67\xc6\xef\x00\x00\x00\xff\xff\x3d\x99\xdb\xbf\xbf\x0a\x00\x00")

func migrations20150927CoreSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20150927CoreSql,
		"migrations/20150927-core.sql",
	)
}

func migrations20150927CoreSql() (*asset, error) {
	bytes, err := migrations20150927CoreSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20150927-core.sql", size: 2751, mode: os.FileMode(420), modTime: time.Unix(1443447734, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations20151002MingbSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xcd\x41\xcb\xc2\x20\x00\x87\xf1\xbb\x9f\xe2\x7f\x7e\x5f\x84\xce\xdb\xc9\xa6\xc1\xc0\x74\x2c\x3d\x8f\xda\x64\x8c\x72\x9a\x2e\x46\xdf\xbe\x76\x0c\xea\xf2\x9c\x1e\xf8\x51\x0a\xfa\x47\xe1\xc3\xe0\x0a\xe4\xfb\xad\xdc\x42\x63\x0a\xc3\xa3\x5f\x0a\xc4\x90\x97\x31\xb9\xbc\x4d\xe4\xfd\xfe\xfb\x69\x4c\xe7\xc5\xc1\x46\xc2\xa4\x11\x2d\x0c\xdb\x4b\x81\x35\xa4\x6b\x97\xa3\xeb\xc1\x38\x47\xa5\xa5\x3d\x2a\xf8\x69\xee\xbc\xf3\x21\x3d\xbb\xf1\x02\xae\xed\x76\x36\xad\xa8\xea\x53\xad\x15\x94\x36\x50\x56\x4a\x70\x71\x60\x56\x1a\xec\x4a\xf2\x61\xf0\xb0\xce\x3f\x14\xde\xea\xe6\x2b\x53\x92\x57\x00\x00\x00\xff\xff\x98\xf6\x73\x37\xd2\x00\x00\x00")

func migrations20151002MingbSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151002MingbSql,
		"migrations/20151002-mingb.sql",
	)
}

func migrations20151002MingbSql() (*asset, error) {
	bytes, err := migrations20151002MingbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151002-mingb.sql", size: 210, mode: os.FileMode(420), modTime: time.Unix(1443812987, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations20151006WorkUnitTypeSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x90\x3d\x4f\xf3\x30\x14\x85\xf7\xfc\x8a\xb3\xf5\x7d\x01\x67\xae\xda\xc9\x6d\xc3\x87\x14\x0a\x4a\x1d\xa4\x4e\x55\x48\x2e\x49\x04\xb6\x83\x7d\x43\xe9\xbf\xc7\x6e\x91\x10\x88\x09\x0f\x1e\xae\xcf\xf1\xf3\xe8\x0a\x01\x71\x26\xa0\x6d\x43\x33\xf8\xd7\x97\x79\xbc\xc4\xe0\x6c\x33\xd6\x3c\xc3\x60\x3d\xb7\x8e\x7c\x0c\x25\x21\xab\xba\xde\xa3\xee\x2a\xd3\x86\x19\x77\x84\xbd\x75\xcf\xbb\xd1\xf4\x9c\x9a\x4a\x13\xf8\x30\x10\x9e\x9c\xd5\x78\x90\xc5\xf2\x5a\x16\x60\x8b\xc5\x56\x65\x32\x05\xa4\x39\x80\xde\x7b\xcf\xbd\x69\xe3\x67\xb1\x8b\xd8\xf5\xa8\x1c\xa1\xf2\x7e\xd4\xd4\xc4\xc6\x23\xa1\x54\x97\x62\x0a\x32\x75\x30\x6b\xd2\x24\xe6\xcf\x75\xdf\xba\x8a\xc3\xdb\x90\xc8\x5c\x65\x05\x94\x5c\xe4\xd9\x97\x03\x4e\xd3\xe5\x5d\x5e\xde\xae\x71\x14\xda\x64\x0a\x2b\xa9\x24\xd4\xf6\x3e\x3b\x99\x24\x38\x9e\x72\x73\xb3\xbe\x42\x6d\xcd\x1b\x39\xde\xb1\xfd\x17\xf3\x17\x98\x04\xf0\x74\xf2\x7f\xfe\x9d\xb8\xb2\x7b\xf3\x47\xe6\xe7\x22\x7e\xa5\xc6\x4d\xfd\xe4\x7e\x04\x00\x00\xff\xff\xd4\xb3\x9b\xee\x92\x01\x00\x00")

func migrations20151006WorkUnitTypeSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151006WorkUnitTypeSql,
		"migrations/20151006-work-unit-type.sql",
	)
}

func migrations20151006WorkUnitTypeSql() (*asset, error) {
	bytes, err := migrations20151006WorkUnitTypeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151006-work-unit-type.sql", size: 402, mode: os.FileMode(420), modTime: time.Unix(1444152140, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations20151013IndexSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xcc\xbd\x0e\x82\x30\x14\xc5\xf1\xbd\x4f\x71\x46\xbf\xfa\x02\x30\x11\xda\xc1\x05\x0c\x68\xe2\x46\x88\x34\x4d\xa3\xed\xad\xb7\x25\xc4\xb7\x57\x26\xe3\xe0\x72\x86\x93\x7f\x7e\x52\x42\xee\x24\x3c\x4d\xa6\x40\x7a\x3e\xca\x75\x64\x64\x9a\xe6\x5b\x2e\x10\x29\x65\xcb\x26\xad\x91\xf8\xb4\x7b\xef\x2c\x8f\xd9\xe0\x12\x45\xdd\xe9\xea\xac\x71\x6c\x94\xbe\x62\x21\xbe\x0f\x73\x70\x79\x20\x9e\x0c\xbb\x60\xd1\x36\xdf\x77\x13\xd9\x11\xbb\xfc\x82\xd2\x7d\x7d\x40\x18\xbd\x41\xd5\xd7\xdb\x52\xfc\xb0\x8a\x96\x20\x54\xd7\x9e\xfe\xb2\xa5\x78\x07\x00\x00\xff\xff\x3f\x3c\x19\x8a\xb3\x00\x00\x00")

func migrations20151013IndexSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151013IndexSql,
		"migrations/20151013-index.sql",
	)
}

func migrations20151013IndexSql() (*asset, error) {
	bytes, err := migrations20151013IndexSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151013-index.sql", size: 179, mode: os.FileMode(420), modTime: time.Unix(1444782480, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations20151014IndexSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\xd0\xd5\xd2\x55\xc8\xcd\x4f\x49\xb5\x52\x28\x2e\xcc\xb1\x06\x11\xba\x05\x45\xf9\x29\xa5\xc9\x25\x56\x0a\x05\xf9\xc5\x25\xe9\x45\xa9\xc5\x20\x45\x5c\x40\xb5\xda\xb9\x99\xe9\x45\x89\x25\xa9\x0a\xa1\x05\x5c\xce\x41\xae\x8e\x21\xae\x0a\x9e\x7e\x2e\xae\x11\x0a\xe5\xf9\x45\xd9\xf1\xa5\x79\x99\x25\xf1\xc5\x05\xa9\xc9\x0a\xfe\x7e\x08\x11\x0d\x30\x0b\x24\x1c\x9f\x99\xa2\x69\x8d\x4f\x5f\x7c\x62\x49\x49\x6a\x6e\x41\x09\x6e\xfd\x3a\x0a\x89\xc9\x25\x99\x65\xa9\x30\x95\x10\x23\x51\x9c\xe6\x92\x5f\x9e\xc7\xe5\x12\xe4\x1f\x80\xd7\x0a\x6b\xdc\x4a\xac\xb9\x00\x01\x00\x00\xff\xff\x75\x53\xe1\x4c\x16\x01\x00\x00")

func migrations20151014IndexSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151014IndexSql,
		"migrations/20151014-index.sql",
	)
}

func migrations20151014IndexSql() (*asset, error) {
	bytes, err := migrations20151014IndexSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151014-index.sql", size: 278, mode: os.FileMode(420), modTime: time.Unix(1444848828, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _migrations20151019WorkerModeSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\xcd\x41\x4b\xc4\x30\x10\x05\xe0\x7b\x7f\xc5\xbb\x2d\x28\x59\x3c\x6f\x4f\x71\x37\xac\x0b\xb5\x4a\x37\x15\x3c\x96\x26\xb6\x41\x9b\xd4\x99\x91\xf5\xe7\xbb\x71\x2d\xe8\x55\x08\x61\x78\xbc\xf9\x46\x29\xa8\x2b\x85\x29\x39\xbf\x01\xbf\xbf\x95\xf9\x53\x33\x25\xf7\xd1\xcb\x06\x73\x62\x19\xc8\x73\x2e\x15\x2a\x3f\xd8\x31\x30\xfa\xb1\x8b\xc3\x39\x96\xd1\xe3\x94\xe8\xd5\xd3\x37\x81\x17\x4a\x13\x42\x14\x48\x02\x0b\x85\x38\xac\x01\xf3\x19\x58\xce\xe3\xef\x26\x67\xaa\x23\x0f\x17\xb8\xef\xc8\x79\xb7\xfe\xf1\xaf\xa7\x30\x50\x27\x1e\xed\x5c\xe8\xca\x9a\x06\x56\xdf\x56\x66\x59\xbe\x44\xdb\x87\xaa\xbd\xaf\x2f\x37\x8f\xc6\x62\xa7\xad\x86\x7d\x7e\x34\x78\xd2\xcd\xf6\x4e\x37\x68\x8f\x87\x7a\x8f\xd5\xaa\x2c\xfe\xa0\xbb\x74\x8a\xff\x61\x0f\xb5\x35\x7b\xb3\xb0\x37\x65\xf1\x15\x00\x00\xff\xff\x85\x69\xfb\x23\x39\x01\x00\x00")

func migrations20151019WorkerModeSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20151019WorkerModeSql,
		"migrations/20151019-worker-mode.sql",
	)
}

func migrations20151019WorkerModeSql() (*asset, error) {
	bytes, err := migrations20151019WorkerModeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20151019-worker-mode.sql", size: 313, mode: os.FileMode(420), modTime: time.Unix(1445281179, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"migrations/20150927-core.sql": migrations20150927CoreSql,
	"migrations/20151002-mingb.sql": migrations20151002MingbSql,
	"migrations/20151006-work-unit-type.sql": migrations20151006WorkUnitTypeSql,
	"migrations/20151013-index.sql": migrations20151013IndexSql,
	"migrations/20151014-index.sql": migrations20151014IndexSql,
	"migrations/20151019-worker-mode.sql": migrations20151019WorkerModeSql,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"migrations": &bintree{nil, map[string]*bintree{
		"20150927-core.sql": &bintree{migrations20150927CoreSql, map[string]*bintree{
		}},
		"20151002-mingb.sql": &bintree{migrations20151002MingbSql, map[string]*bintree{
		}},
		"20151006-work-unit-type.sql": &bintree{migrations20151006WorkUnitTypeSql, map[string]*bintree{
		}},
		"20151013-index.sql": &bintree{migrations20151013IndexSql, map[string]*bintree{
		}},
		"20151014-index.sql": &bintree{migrations20151014IndexSql, map[string]*bintree{
		}},
		"20151019-worker-mode.sql": &bintree{migrations20151019WorkerModeSql, map[string]*bintree{
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

