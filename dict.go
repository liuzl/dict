package dict

import (
	"fmt"
	"sync"

	"github.com/liuzl/cedar-go"
)

type Cedar struct {
	*cedar.Cedar
	*sync.RWMutex
}

type Pos struct {
	StartByte int `json:"start_byte"`
	EndByte   int `json:"end_byte"`
}

func (p *Pos) String() string {
	return fmt.Sprintf("[%d-%d]", p.StartByte, p.EndByte)
}

func New() *Cedar {
	return &Cedar{cedar.New(), new(sync.RWMutex)}
}

func (da *Cedar) MultiMatch(text string) (map[string][]*Pos, error) {
	da.RLock()
	defer da.RUnlock()
	ret := make(map[string][]*Pos)
	r := []rune(text)
	for i, _ := range r {
		start := len(string(r[:i]))
		sub := r[i:]
		for _, id := range da.PrefixMatch([]byte(string(sub)), 0) {
			k, err := da.Key(id)
			if err != nil {
				return nil, err
			}
			match := string(k)
			pos := &Pos{start, start + len(match)}
			if ret[match] == nil {
				ret[match] = []*Pos{pos}
			} else {
				ret[match] = append(ret[match], pos)
			}
		}
	}
	return ret, nil
}

func (da *Cedar) MultiSearch(text string) (ret []string) {
	da.RLock()
	defer da.RUnlock()
	uText := []rune(text)
	for idx, _ := range uText {
		sub := uText[idx:]
		for _, id := range da.PrefixMatch([]byte(string(sub)), 0) {
			k, _ := da.Key(id)
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
