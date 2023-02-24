## 错解

### <span style="color:orange">思路一</span>

先中序遍历二叉树，即 左子树 -> 根节点 -> 右子树，若二叉树为『对称二叉树』，比如：

![binarytree-01.png](/Users/sun/Workspace/GO/Just-Algorithm/resources/imgs/binarytree-01.png)

中序遍历的结果为：3，2，4，<strong><span style="color:orange">1</span></strong>，4，2，3
。将中序遍历的结果存入数组中，再设置双指针 l 和 r 分别指向数组的两端，若 l 和 r 所指元素不同，则不是对称二叉树。

### <span style="color:orange">忽略点</span>

当出现如下二叉树时，即使双指针指向的值都是相同的，该二叉树也不是『对称二叉树』。

![binarytree-02.png](/Users/sun/Workspace/GO/Just-Algorithm/resources/imgs/binarytree-02.png)

该二叉树前序遍历的结果为：2,2,1,2,2。

错误代码如下：

```go
var values []int

func isSymmetric(root *common.TreeNode) bool {
	values = []int{}
	midOrder(root)
	fmt.Printf("values: %v\n", values)

	l, r := 0, len(values)-1
	for l < r {
		if values[l] != values[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func midOrder(root *common.TreeNode) {
	if root == nil {
		return
	}

	midOrder(root.Left)
	values = append(values, root.Val)
	midOrder(root.Right)
}
```

<span style="color:#e6e600">🤔 涉及问题：</span>

* 最初是将 values 数组作为参数放入 midOrder 函数中的，但是前序遍历结束后，数组却为空，代码如下：

```go
func isSymmetric(root *common.TreeNode) bool {
	var values []int
	midOrder(root, values)
	fmt.Printf("values: %v\n", values)

	l, r := 0, len(values)-1
	for l < r {
		if values[l] != values[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func midOrder(root *common.TreeNode, values []int) {
	if root == nil {
		return
	}

	midOrder(root.Left, values)
	values = append(values, root.Val)
	midOrder(root.Right, values)
}

func main() {
	root := common.TreeNode{Val: 1}
	root.Left = &common.TreeNode{Val: 2}
	root.Right = &common.TreeNode{Val: 2}
	root.Left.Left = &common.TreeNode{Val: 3}
	root.Left.Right = &common.TreeNode{Val: 4}
	root.Right.Left = &common.TreeNode{Val: 4}
	root.Right.Right = &common.TreeNode{Val: 3}
	fmt.Println(isSymmetric(&root))
}
```

控制台输出结果为：

```
values: []
true
```

## 正解

使用『递归』求解。