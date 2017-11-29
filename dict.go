package dict

import (
	"github.com/liuzl/cedar-go"
	"sync"
)

type Cedar struct {
	*cedar.Cedar
	*sync.RWMutex
}

func New() *Cedar {
	return &Cedar{cedar.New(), new(sync.RWMutex)}
}

func (self *Cedar) MultiSearch(text string) (ret []string) {
	self.RLock()
	defer self.RUnlock()
	uText := []rune(text)
	for idx, _ := range uText {
		sub := uText[idx:]
		for _, id := range self.PrefixMatch([]byte(string(sub)), 0) {
			k, _ := self.Key(id)
			ret = append(ret, string(k))
		}
	}
	return ret
}

func (da *Cedar) SafeInsert(key []byte, value int) error {
	da.Lock()
	defer da.Unlock()
	return da.Insert(key, value)
}

func (da *Cedar) SafeUpdate(key []byte, value int) error {
	da.Lock()
	defer da.Unlock()
	return da.Update(key, value)
}

func (da *Cedar) SafeDelete(key []byte) error {
	da.Lock()
	defer da.Unlock()
	return da.Delete(key)
}
