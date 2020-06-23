package xrouter

// 路由组
type GroupOperator struct {
	EmptyOperator
	path string
}

func NewGroup(path string) *GroupOperator {
	return &GroupOperator{path: path}
}

func (e GroupOperator) Path() string {
	return e.path
}

