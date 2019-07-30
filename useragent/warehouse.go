package useragent

import (
	"math/rand"
	"sync"
)

type useragent struct {
	data map[string][]string
	lock sync.Mutex
}

var (
	UA = useragent{data: make(map[string][]string)}
)

func (u *useragent) Get(key string) []string {
	return u.data[key]
}

func (u *useragent) GetAll() map[string][]string {
	return u.data
}

func (u *useragent) GetRandom(key string) string {
	browser := u.Get(key)
	l := len(browser)
	if l < 1 {
		return ""
	}
	n := rand.Intn(l)
	return browser[n]
}

func (u *useragent) GetAllRandom() string {
	browsers := u.GetAll()
	data := []string{}
	for _, uas := range browsers {
		data = append(data, uas...)
	}
	l := len(data)
	if l < 1 {
		return ""
	}
	n := rand.Intn(l)
	return data[n]
}

func (u *useragent) Set(key, value string) {
	u.lock.Lock()
	defer u.lock.Unlock()
	u.data[key] = append(u.data[key], value)
}

func (u *useragent) SetData(data map[string][]string) {
	u.data = data
}
