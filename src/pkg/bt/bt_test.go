package bt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalancedTree(t *testing.T) {
	testCases := []struct {
		name           string
		tree           *TreeNode
		expectedResult bool
	}{
		{
			name: "balancedTree#1",
			tree: &TreeNode{
				HasToy: false,
				Right: &TreeNode{
					HasToy: true,
				},
				Left: &TreeNode{
					HasToy: false,
					Right: &TreeNode{
						HasToy: true,
					},
					Left: &TreeNode{
						HasToy: false,
					},
				},
			},
			expectedResult: true,
		},
		{
			name: "balancedTree#2",
			tree: &TreeNode{
				HasToy: true,
				Right: &TreeNode{
					HasToy: false,
					Right: &TreeNode{
						HasToy: true,
					},
					Left: &TreeNode{
						HasToy: true,
					},
				},
				Left: &TreeNode{
					HasToy: true,
					Right: &TreeNode{
						HasToy: false,
					},
					Left: &TreeNode{
						HasToy: true,
					},
				},
			},
			expectedResult: true,
		},
		{
			name: "notBalancedTree#1",
			tree: &TreeNode{
				HasToy: true,
				Right: &TreeNode{
					HasToy: false,
				},
				Left: &TreeNode{
					HasToy: true,
				},
			},
			expectedResult: false,
		},
		{
			name: "notBalancedTree#2",
			tree: &TreeNode{
				HasToy: false,
				Right: &TreeNode{
					HasToy: false,
					Right: &TreeNode{
						HasToy: true,
					},
				},
				Left: &TreeNode{
					HasToy: true,
					Right: &TreeNode{
						HasToy: true,
					},
				},
			},
			expectedResult: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, areToysBalanced(tc.tree), tc.expectedResult)
		})
	}
}

func TestUnrollGarland(t *testing.T) {
	testCases := []struct {
		name           string
		tree           *TreeNode
		expectedResult []bool
	}{
		{
			name: "unrollGarland#1",
			tree: &TreeNode{
				HasToy: true,
				Right: &TreeNode{
					HasToy: false,
					Right: &TreeNode{
						HasToy: true,
					},
					Left: &TreeNode{
						HasToy: true,
					},
				},
				Left: &TreeNode{
					HasToy: true,
					Right: &TreeNode{
						HasToy: false,
					},
					Left: &TreeNode{
						HasToy: true,
					},
				},
			},
			expectedResult: []bool{true, true, false, true, true, false, true},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, unrollGarland(tc.tree))
		})
	}
}
