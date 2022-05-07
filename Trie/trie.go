package trie

const chars_val = 26

type Node struct {
	to       [chars_val]*Node
	terminal bool
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) AddString(word string) {
	current := n
	for _, val := range word {
		val -= 'a'
		if current.to[val] == nil {
			current.to[val] = NewNode()
		}
		current = current.to[val]
	}
	current.terminal = true
}

func (n *Node) Find(word string) bool {
	current := n
	for _, val := range word {
		val -= 'a'
		if current.to[val] == nil {
			return false
		}
		current = current.to[val]
	}
	return current.terminal
}
