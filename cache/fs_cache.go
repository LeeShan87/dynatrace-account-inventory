package cache

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FSCache struct {
	dir string
}

func NewFSCache(dir string) *FSCache {
	return &FSCache{dir: dir}
}

func (f *FSCache) Get(key string) (string, bool) {
	data, err := ioutil.ReadFile(filepath.Join(f.dir, key))
	if err != nil {
		fmt.Println("Cache miss:", key)
		return "", false
	}
	fmt.Println("Cache hit:", key)
	return string(data), true
}

func (f *FSCache) Set(key string, value string) {
	err := ioutil.WriteFile(filepath.Join(f.dir, key), []byte(value), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
