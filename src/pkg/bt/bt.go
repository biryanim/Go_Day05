package bt

import "container/list"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func countToysOnSubtree(root *TreeNode, count *int) {
	if root == nil {
		return
	}
	if root.HasToy {
		*count++
	}
	if root.Right != nil {
		countToysOnSubtree(root.Right, count)
	}
	if root.Left != nil {
		countToysOnSubtree(root.Left, count)
	}
}

func areToysBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left, right := 0, 0
	if root.Right != nil {
		countToysOnSubtree(root.Right, &right)
	}
	if root.Left != nil {
		countToysOnSubtree(root.Left, &left)
	}
	if left != right {
		return false
	}
	return true
}

func unrollGarland(root *TreeNode) []bool {
	if root == nil {
		return nil
	}
	var (
		res   []bool
		level = 1
	)

	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		listSize := l.Len()
		curLevelList := list.New()
		for i := 0; i < listSize; i++ {
			node := l.Remove(l.Front()).(*TreeNode)
			curLevelList.PushBack(node)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}

		for curLevelList.Len() > 0 {
			var node *TreeNode
			if level%2 == 0 {
				node = curLevelList.Remove(curLevelList.Front()).(*TreeNode)
			} else {
				node = curLevelList.Remove(curLevelList.Back()).(*TreeNode)
			}
			res = append(res, node.HasToy)
		}
		level++
	}
	return res
}
