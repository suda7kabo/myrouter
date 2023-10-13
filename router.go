package myrouter

import (
	"net/http"
)

type Node struct {
	isRoot    bool
	character byte
	children  []*Node
	handlers  map[string]http.Handler
}

func newNode(character byte) *Node {
	return &Node{
		character: character,
		children:  []*Node{},
		handlers:  make(map[string]http.Handler),
	}
}

func (n *Node) nextChild(character byte) *Node {
	for _, child := range n.children {
		if child.character == character {
			return child
		}
	}
	return nil
}

// ///////////////////////////////
type Router struct {
	tree *Node
}

func NewRouter() *Router {
	return &Router{
		tree: &Node{
			isRoot:   true,
			handlers: nil,
		},
	}
}
func (r *Router) GET(endpoint string, handler http.Handler) {
	r.insert(http.MethodGet, endpoint, handler)
}
func (r *Router) insert(method, endpoint string, handler http.Handler) {
	currentNode := r.tree
	for i := 0; i < len(endpoint); i++ {
		target := endpoint[i]

		nextNode := currentNode.nextChild(target)
		if nextNode == nil {
			node := newNode(target)
			currentNode.children = append(currentNode.children, node)
			currentNode = node
			continue
		}
		currentNode = nextNode
	}
	currentNode.handlers[method] = handler
}
