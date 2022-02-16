package trie

// http://www.mrsnippet.com/2015/02/11/invalid-recursive-type-struct/
type Trie struct {
	Tree  [26]*Trie
	isEnd bool
}

func Constructor() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	cur := t
	for _, runeWord := range word {
		index := runeWord - 'a'
		if cur.Tree[index] == nil {
			cur.Tree[index] = &Trie{}
		}
		cur = cur.Tree[index]
	}
	cur.isEnd = true
}

func (t *Trie) search(s string) (lastRune *Trie) {
	cur := t
	for _, runeWord := range s {
		index := runeWord - 'a'
		if cur.Tree[index] == nil {
			return nil
		}
		cur = cur.Tree[index]
	}
	return cur
}

func (t *Trie) Search(word string) bool {
	lastRune := t.search(word)
	return lastRune != nil && lastRune.isEnd == true
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.search(prefix) != nil
}
