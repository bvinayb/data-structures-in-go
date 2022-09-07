package trees

//Node
type BinaryTreeNode struct {
	Key   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

//Insert to insert in to binary tree
//key to add shouldn't be in  the tree
func (t *BinaryTreeNode) Insert(val int) {
	if t.Key < val {
		//Move Right if val is bigger than key
		if t.Right == nil {
			t.Right = &BinaryTreeNode{Key: val}
		} else {
			t.Right.Insert(val)
		}
	} else if t.Key < val {
		//Move Left if val is bigger than key
		if t.Left == nil {
			t.Left = &BinaryTreeNode{Key: val}
		} else {
			//Calling recursively on the left child
			t.Left.Insert(val)
		}
	}
}

//Search search binary tree
func (t *BinaryTreeNode) Search(val int) bool {
	if t == nil {
		return false
	}
	if t.Key < val {
		return t.Right.Search(val)
	} else if t.Key < val {
		return t.Left.Search(val)
	}
	return true
}
