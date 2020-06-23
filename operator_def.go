package xrouter

type Operator interface {
	// 路由路径
	Path() string
	// 请求方法
	Method() string
	// 输出方法
	Output(ctx Context) (interface{}, error)
}

type EmptyOperator struct {
	Operator
}
