package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type File struct {
	Dir          string
	Name         string
	Absolutepath string
}

func NewFileCache(dir string, name string) *File {
	return &File{
		Dir:          dir,
		Name:         name,
		Absolutepath: dir + name,
	}
}

func (f *File) Read() ([]byte, error) {
	return ioutil.ReadFile(f.Absolutepath)
}

func (f *File) Write(data []byte) error {
	return ioutil.WriteFile(f.Absolutepath, data, 0644)
}

func (f *File) WriteJson(v interface{}) error {
	uasJson, err := json.Marshal(v)
	if err != nil {
		return nil
	}

	return f.Write(uasJson)
}

func (f *File) Remove() error {
	err := os.Remove(f.Absolutepath)
	if err != nil {
		return nil
	}
	return nil
}

func (f *File) IsExist() (bool, error) {
	_, err := os.Stat(f.Absolutepath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}

func GetTempDir() string {
	tempDir := os.TempDir()
	if exist := strings.HasSuffix(tempDir, "/"); exist == false {
		tempDir += "/"
	}
	return tempDir
}
