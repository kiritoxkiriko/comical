package comical

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// node is a node in the of trie DefaultRouter
type node struct {
	// children store children nodes
	children []*node
	// pattern store top level to current level path, if empty means this node isn't the end of path
	pattern string
	// part current level DefaultRouter
	part string
	// isWild if true means this node is a wild card node
	isWild bool
}

// newNode creates a new node
func newNode() *node {
	return &node{
		children: make([]*node, 0),
	}
}

// matchOne matches one node
func (n *node) matchOne(val string) *node {
	for _, child := range n.children {
		if child.isWild || child.part == val {
			return child
		}
	}
	return nil
}

// matchMany matches many nodes
func (n *node) matchMany(val string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.isWild || child.part == val {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert inserts a new node
// Params:
// - pattern: full path pattern
// - parts: path parts, length of parts is the depth of the path
// - depth: the current depth of the path
func (n *node) insert(pattern string, parts []string, depth int) {
	// if current node depth equals path depth, then set the full pattern and return
	if depth == len(parts) {
		if n.pattern != "" {
			panic(fmt.Errorf("pattern %s already exists", pattern))
		}
		n.pattern = pattern
		return
	}

	// get current part and child
	part := parts[depth]
	child := n.matchOne(part)
	// no node match in this depth, add one
	if child == nil {
		child = newNode()
		child.part = part
		child.isWild = strings.HasPrefix(part, ":") || strings.HasPrefix(part, "*")
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, depth+1)
}

// search find matched node
func (n *node) search(parts []string, depth int) *node {
	// if current node depth equals path depth, then check whether there is pattern
	if depth == len(parts) || strings.HasPrefix(n.part, "*") {
		// not found
		if n.pattern == "" {
			return nil
		}
		// found it
		return n
	}

	part := parts[depth]
	children := n.matchMany(part)

	// travel all matched children
	for _, child := range children {
		res := child.search(parts, depth+1)
		if res != nil {
			return res
		}
	}
	// found nothing, return
	return nil
}

func (n *node) travel(depth int) {
	log.Printf("depth: %d part: %s\n", depth, n.part)
	if n.pattern != "" {
		log.Printf("pattern: %s\n", n.pattern)
	}
	for _, child := range n.children {
		child.travel(depth + 1)
	}
}

// TrieRouter is a router support path pattern
// e.g. /p/:lang/doc, /p/*hello, /p/python/doc
type TrieRouter struct {
	trieRoots map[string]*node
	handlers  map[string]HandlerFunc
}

// NewTrieRouter creates a new trie defaultRouter
func NewTrieRouter() *TrieRouter {
	return &TrieRouter{
		trieRoots: map[string]*node{},
		handlers:  map[string]HandlerFunc{},
	}
}

func (r *TrieRouter) parsePattern(pattern string) []string {
	// split patterns
	patterns := strings.Split(pattern, "/")

	parts := make([]string, 0, len(patterns))

	for _, part := range patterns {
		if part != "" {
			parts = append(parts, part)
			// has wild card, then tailing path are same part
			if strings.HasPrefix(part, "*") {
				break
			}
		}
	}
	return parts
}

// Handle handles http requests by Context
func (r *TrieRouter) Handle(c *Context) {
	//r.trieRoots[c.Method].travel(0)
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
		return
	}
	c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
}

// AddRoute adds a route to the trie router
func (r *TrieRouter) AddRoute(method, pattern string, handler HandlerFunc) {
	parts := r.parsePattern(pattern)
	key := method + "-" + pattern
	if _, ok := r.trieRoots[method]; !ok {
		r.trieRoots[method] = newNode()
	}
	r.trieRoots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// getRoute gets a route from the trie router
func (r *TrieRouter) getRoute(method, pattern string) (n *node, params map[string]string) {
	searchParts := r.parsePattern(pattern)

	trieRoot, ok := r.trieRoots[method]
	if !ok {
		return
	}

	n = trieRoot.search(searchParts, 0)
	if n != nil {
		params = map[string]string{}
		// current n path parts
		parts := r.parsePattern(n.pattern)
		for i, part := range parts {
			if strings.HasPrefix(part, ":") {
				params[part[1:]] = searchParts[i]
			}
			if strings.HasPrefix(part, "*") && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[i:], "/")
				break
			}
		}
		return
	}
	return
}
