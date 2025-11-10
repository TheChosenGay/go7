package gee

type tri_node struct {
	pattern  string
	part     string
	isWild   bool
	children []*tri_node
}

func (n *tri_node) matchChild(part string) *tri_node {
	for _, child := range n.children {
		if child.part == part {
			return child
		}
	}
	return nil
}

func (n *tri_node) matchChildren(part string) []*tri_node {
	nodes := make([]*tri_node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *tri_node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		newChild := &tri_node{
			part:     part,
			isWild:   part[0] == '*' || part[0] == ':',
			children: make([]*tri_node, 0),
		}
		n.children = append(n.children, newChild)
		newChild.insert(pattern, parts, height+1)
	} else {
		child.insert(pattern, parts, height+1)
	}
}

func (n *tri_node) search(parts []string, height int) *tri_node {
	if len(parts) == height || n.part[0] == '*' {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
