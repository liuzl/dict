package dict

import (
	"testing"
)

func TestDict(t *testing.T) {
	trie := New()
	trie.Insert([]byte("How many"), 0)
	trie.Insert([]byte("How many loved"), 1)
	trie.Insert([]byte("How many loved your moments"), 2)
	trie.Insert([]byte("How many loved your moments of glad grace"), 3)
	trie.Insert([]byte("姑苏"), 4)
	trie.Insert([]byte("姑苏城外"), 5)
	trie.Insert([]byte("姑苏城外寒山寺"), 6)
	trie.SaveToFile("cedar.gob", "gob")
	d := New()
	d.LoadFromFile("cedar.gob", "gob")

	str := "《姑苏城外寒山寺》是一首挺好的诗,姑苏就是现在的苏州."
	ret := d.MultiSearch(str)
	t.Log(ret)

	matches, err := d.MultiMatch(str)
	if err != nil {
		t.Error(err)
	}
	t.Log(matches)
}
