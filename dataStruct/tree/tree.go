package tree

type Node struct {
	Id       int
	Name     string
	ParentId int
	Child    []*Node
}

func BuildTree(data []*Node) []*Node {
	var nodeMap = make(map[int]*Node)
	for _, datum := range data {
		nodeMap[datum.Id] = datum
	}

	var root []*Node
	for _, datum := range data {
		if node, ok := nodeMap[datum.ParentId]; ok {
			node.Child = append(node.Child, datum)
		} else {
			root = append(root, datum)
		}
	}
	return root
}
