package tries

const AlphabetSize = 26

//Node
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

//Tries
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{
		root: &Node{},
	}
	return result
}

//Insert
func (t *Trie) Insert(w string) {
	lenOfWord := len(w)
	currentNode := t.root
	for i := 0; i < lenOfWord; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

//Search
func (t *Trie) Search(w string) bool {
	lenOfWord := len(w)
	currentNode := t.root
	for i := 0; i < lenOfWord; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}
