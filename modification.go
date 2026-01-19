package xmlquery

func MoveChildNodes(source, destination *Node) {
	for childNodes := source.FirstChild; childNodes != nil; childNodes = childNodes.NextSibling {
		if destination.LastChild == nil {
			destination.FirstChild = childNodes
			destination.LastChild = childNodes
		} else {
			destination.LastChild.NextSibling = childNodes
			childNodes.PrevSibling = destination.LastChild
			destination.LastChild = childNodes
		}
		childNodes.level = destination.level + 1
		adaptSubLevel(childNodes, childNodes.level+1)
		childNodes.Parent = destination
	}
	source.FirstChild = nil
	source.LastChild = nil
}

func adaptSubLevel(node *Node, level int) {
	for childNode := node.FirstChild; childNode != nil; childNode = childNode.NextSibling {
		childNode.level = level
		adaptSubLevel(childNode, level+1)
	}
}

func RemoveWithCriterium(node *Node, xpath string, f func(node *Node) bool) {
	FindEach(node, xpath, func(i int, n *Node) {
		if f(n) {
			if n.PrevSibling == nil {
				n.FirstChild = n.NextSibling
			} else {
				n.PrevSibling.NextSibling = n.NextSibling
				if n.NextSibling != nil {
					n.NextSibling.PrevSibling = n.PrevSibling
				}
			}
		}
	})
}
