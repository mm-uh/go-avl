package avl

type Node struct {
	leftChild  *Node
	rigthChild *Node
	height     int
	key        Key
	Value      interface{}
}

type Key interface {
	Less(other interface{}) (bool, error)
}

func NewNode(key Key, value interface{}) *Node {
	return &Node{
		leftChild:  nil,
		rigthChild: nil,
		height:     1,
		key:        key,
		value:      value,
	}
}

func (n *Node) isLeaf() bool {
	return n.rigthChild != nil && n.leftChild != nil
}

func (n *Node) hasRChild() bool {
	return n.rigthChild != nil
}

func (n *Node) hasLChild() bool {
	return n.leftChild != nil
}

func getHeight(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}
