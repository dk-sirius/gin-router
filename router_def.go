package xrouter

import "fmt"

type Router struct {
	parent   *Router
	operator []Operator
	children map[*Router]bool
}

func NewRouter(operators ...Operator) *Router {
	return &Router{
		operator: operators,
	}
}

func (router *Router) Register(r *Router) {
	if router.children == nil {
		router.children = map[*Router]bool{}
	}
	if r.parent != nil {
		panic(fmt.Errorf("router %v already registered to router %v", r, r.parent))
	}
	router.children[r] = true
	r.parent = router
}
