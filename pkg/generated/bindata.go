// Code generated by go-bindata.
// sources:
// assets/clusterrole.yaml
// assets/clusterrolebinding.yaml
// assets/deployment.yaml
// assets/role.yaml
// assets/rolebinding.yaml
// assets/serviceaccount.yaml
// DO NOT EDIT!

package generated

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

var _assetsClusterroleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\xb1\x6a\xf4\x30\x10\x84\x7b\x3f\xc5\xe2\xfa\xec\x9f\xbf\x0b\x6e\x53\xa4\x4f\x91\x26\xa4\xd8\x93\x87\xbb\xe5\x64\x49\xec\xae\x1c\xc8\xd3\x07\x3b\x82\x14\x39\x82\x21\xd5\x8a\xd5\xf0\x8d\x46\x73\x93\x34\x4f\xf4\x18\xab\x39\xf4\x39\x47\x74\x5c\xe4\x05\x6a\x92\xd3\x44\x7a\xe6\x30\x72\xf5\x6b\x56\xf9\x60\x97\x9c\xc6\xdb\x83\x8d\x92\xff\xad\xff\xbb\x05\xce\x33\x3b\x4f\x1d\x51\xe2\x05\x13\x2d\x9c\x24\xf2\x50\x34\xaf\xb2\x01\xa0\x83\xd6\x94\xa0\x9d\xd6\x08\xdb\x84\x03\x71\x91\x27\xcd\xb5\xd8\x44\xaf\x7d\xff\xd6\x11\x11\x29\x2c\x57\x0d\xd8\x77\x86\xa0\x70\x6b\x57\x2b\xf4\xbc\xaf\x83\x82\x1d\xfd\x89\xfa\x0b\x7c\x1b\x33\x22\x1c\xbb\xec\x00\xb5\x6c\x99\xcc\x91\x7c\xcd\xb1\x2e\xf8\xc1\x6f\xd4\x28\xb6\xcf\x77\xf6\x70\xdd\x0e\xdf\xbe\x7f\x33\x0c\x91\x65\x39\xec\x5a\xcb\xcc\xf7\xbd\xcc\xb3\xf2\x05\xad\x88\xbb\x1f\xf8\xa5\x08\x91\xcd\x0e\xe6\x3c\x98\x09\x2b\xd2\x6f\xcd\xb4\x57\x9f\xa8\x2f\x0d\xfb\x19\x00\x00\xff\xff\x03\x8c\xed\xd0\x60\x02\x00\x00")

func assetsClusterroleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsClusterroleYaml,
		"assets/clusterrole.yaml",
	)
}

func assetsClusterroleYaml() (*asset, error) {
	bytes, err := assetsClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsClusterrolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcf\x31\x4e\xc5\x30\x0c\x80\xe1\x3d\xa7\xf0\x05\x5a\xc4\x86\xb2\x01\x03\xfb\x43\x62\x77\x13\x17\x4c\x53\x3b\x72\x9c\x0e\x9c\x1e\x45\x95\x58\xd0\xeb\x6c\xeb\xb3\xff\x8d\x25\x47\x78\x2d\xbd\x39\xd9\x4d\x0b\xbd\xb0\x64\x96\xcf\x80\x95\x3f\xc8\x1a\xab\x44\xb0\x05\xd3\x8c\xdd\xbf\xd4\xf8\x07\x9d\x55\xe6\xed\xa9\xcd\xac\x0f\xc7\x63\xd8\xc9\x31\xa3\x63\x0c\x00\x82\x3b\x45\xd8\x51\xb8\xe0\x54\x4d\x0f\x1e\x00\xd9\x64\x5a\x28\xb4\xbe\x7c\x53\xf2\x36\x36\x27\x38\x2f\xbf\x93\x1d\x9c\xe8\x39\x25\xed\xe2\x01\xe0\x3e\xf2\x37\x6c\x15\x13\x45\xc8\xb4\x62\x2f\x1e\x86\x7d\xa3\x75\xa8\xff\x6a\x2e\x7f\xea\x72\xaa\x58\xf9\xcd\xb4\xd7\x8b\xd0\xf0\x1b\x00\x00\xff\xff\xf8\x20\x60\xe7\x28\x01\x00\x00")

func assetsClusterrolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsClusterrolebindingYaml,
		"assets/clusterrolebinding.yaml",
	)
}

func assetsClusterrolebindingYaml() (*asset, error) {
	bytes, err := assetsClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x3f\x6b\x33\x31\x0c\xc6\xf7\xfb\x14\x22\xfb\xe5\x0f\xef\xf2\x62\xe8\x50\xe8\x52\x28\x25\x53\x97\xd2\x41\xd5\x29\x89\x38\xd9\x32\xb6\x72\x90\x6f\x5f\x9c\xb4\x21\xa1\x4d\x34\xdd\xc1\xf3\xfb\x3d\xb2\x8d\x59\xde\xb8\x54\xb1\x14\x00\x73\xae\x8b\x69\xd5\x8d\x92\x86\x00\x4f\x9c\xd5\x0e\x91\x93\x77\x91\x1d\x07\x74\x0c\x1d\x40\xc2\xc8\x01\x22\x26\x51\xec\x73\xb1\x49\x1a\xcc\xa5\x03\x50\xfc\x64\xad\x2d\x04\xcd\xf5\x67\xaa\x66\xa6\x96\x28\x9c\x55\x08\x6b\x80\x55\x07\x50\x59\x99\xdc\xca\x89\x8d\xe8\xb4\x7b\xb9\x90\xdd\xd6\x01\x38\xc7\xac\xe8\xfc\x8d\x5e\x6c\xda\x46\xaf\x2c\xf7\x3c\x00\x3f\xab\x1d\xbf\xb9\x4c\x42\xfc\x48\x64\xfb\xe4\x37\x11\x00\xb2\xe4\x28\x89\xcb\xb9\xa4\xbf\x77\x45\xa7\x91\x88\x5b\x0e\x30\x18\x8d\x5c\xe6\x62\x8b\xf1\x7f\x25\xb5\xfd\x70\x0c\x0f\x5c\x16\xbf\xe1\x30\x2d\xe7\xff\xe6\xcb\x6b\xc7\x7a\xaf\xba\x36\x15\x3a\x04\x78\xde\xbc\x9a\xaf\x0b\xd7\xf6\x62\xe7\xf3\x96\x6d\x0d\xf0\x0e\xb3\xfe\xd2\xf5\x60\x99\x53\xdd\xc9\xc6\x5b\xf9\xf1\xc7\x91\xc6\xfe\xd4\x3a\x83\x8f\xee\x2b\x00\x00\xff\xff\x35\x20\xfd\x63\x17\x02\x00\x00")

func assetsDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDeploymentYaml,
		"assets/deployment.yaml",
	)
}

func assetsDeploymentYaml() (*asset, error) {
	bytes, err := assetsDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xcd\x31\x4e\xc5\x30\x0c\x06\xe0\x3d\xa7\xb0\x32\xbf\x16\xb1\xa1\x5c\x80\x9d\x81\x05\xbd\xc1\x2f\xb1\x5e\xad\xa6\x76\xe4\x38\x45\xe2\xf4\xa8\x11\x37\x60\xb3\xfc\xdb\xdf\xbf\xb3\x94\x04\x1f\x5a\x29\x60\xe3\x4f\xb2\xce\x2a\x09\xec\x81\x79\xc5\xe1\x9b\x1a\xff\xa0\xb3\xca\xba\xbf\xf5\x95\xf5\xe5\x7c\x0d\x07\x39\x16\x74\x4c\x01\x40\xf0\xa0\x04\x95\xb0\x90\x2d\x55\xf3\xce\xf2\x5c\x0e\x14\xae\xb8\x34\xd3\x93\x2f\x8f\x2c\xd8\xa8\xd4\xaf\x87\x05\xb0\xf1\xbb\xe9\x68\x3d\xc1\x57\x8c\xf7\x00\x00\x60\xd4\x75\x58\xa6\xb9\x23\x29\x4d\x59\xbc\xff\x85\x27\xd9\x63\x06\x4f\xf2\x78\x83\x38\x5a\x41\xa7\x6b\x6a\xe8\x79\x9b\x57\xff\x61\x2b\xf7\xe9\x7e\x4f\xed\x06\x31\x1b\x5d\x05\xf7\xf0\x1b\x00\x00\xff\xff\x06\xfa\x1e\x71\x1e\x01\x00\x00")

func assetsRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRoleYaml,
		"assets/role.yaml",
	)
}

func assetsRoleYaml() (*asset, error) {
	bytes, err := assetsRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8f\x31\x8e\xc2\x30\x10\x45\x7b\x9f\x62\x2e\xe0\xac\xb6\x5b\xb9\xdb\x6d\xb6\x0f\x12\xfd\xc4\x1e\xc2\x10\x67\xc6\x1a\xdb\x29\x38\x3d\x8a\x10\xa2\x41\x48\x1c\xe0\xbd\xf7\xff\xc2\x92\x02\x8c\x9a\xe9\x8f\x25\xb1\xcc\x0e\x0b\x1f\xc9\x2a\xab\x04\xb0\x09\xe3\x80\xbd\x9d\xd5\xf8\x8a\x8d\x55\x86\xe5\xa7\x0e\xac\x5f\xdb\xb7\x5b\xa9\x61\xc2\x86\xc1\x01\x08\xae\x14\x20\x13\x26\x32\x9f\x35\x2e\x2c\xb3\x5f\x51\x38\xa3\x2f\xa6\x1b\xef\x3e\x32\x57\xfb\x74\xa1\xd8\xea\xce\x78\xb8\xc7\x0f\x64\x1b\x47\xfa\x8d\x51\xbb\x34\x07\xf0\xd0\xbd\xe0\x4d\x33\x8d\x74\xda\xf1\xe7\xf2\x0f\xfa\x00\x58\xf8\xdf\xb4\x97\x37\xe7\xdc\x2d\x00\x00\xff\xff\xc3\x78\x1f\x5e\x15\x01\x00\x00")

func assetsRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRolebindingYaml,
		"assets/rolebinding.yaml",
	)
}

func assetsRolebindingYaml() (*asset, error) {
	bytes, err := assetsRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xcb\xb1\x0d\x84\x30\x0c\x05\xd0\xde\x53\x78\x81\x2b\xae\x75\xc7\x0c\x48\xf4\x5f\x89\x8b\x2f\x88\x13\x99\x90\xf9\x11\x03\x3c\x0c\x1e\x9e\x37\x7b\x98\xae\xbf\x9c\x8c\x6a\xba\x7b\x2e\x16\xdf\x4a\xe9\x4f\x4c\x69\x3e\x51\x31\x61\xa2\x1a\x68\x6e\xda\x10\xbc\xf0\x1b\xd9\x17\x3f\xeb\x29\x6f\x00\x00\x00\xff\xff\x05\xe2\x35\x12\x49\x00\x00\x00")

func assetsServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsServiceaccountYaml,
		"assets/serviceaccount.yaml",
	)
}

func assetsServiceaccountYaml() (*asset, error) {
	bytes, err := assetsServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"assets/clusterrole.yaml": assetsClusterroleYaml,
	"assets/clusterrolebinding.yaml": assetsClusterrolebindingYaml,
	"assets/deployment.yaml": assetsDeploymentYaml,
	"assets/role.yaml": assetsRoleYaml,
	"assets/rolebinding.yaml": assetsRolebindingYaml,
	"assets/serviceaccount.yaml": assetsServiceaccountYaml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"clusterrole.yaml": &bintree{assetsClusterroleYaml, map[string]*bintree{}},
		"clusterrolebinding.yaml": &bintree{assetsClusterrolebindingYaml, map[string]*bintree{}},
		"deployment.yaml": &bintree{assetsDeploymentYaml, map[string]*bintree{}},
		"role.yaml": &bintree{assetsRoleYaml, map[string]*bintree{}},
		"rolebinding.yaml": &bintree{assetsRolebindingYaml, map[string]*bintree{}},
		"serviceaccount.yaml": &bintree{assetsServiceaccountYaml, map[string]*bintree{}},
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

