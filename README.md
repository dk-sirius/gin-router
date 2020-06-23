# gin-router
http router base gonic-gin framework

# how to use

## config gin-router engine
### init gin engine
engine:=gin-router.InitDefaultEngine()
### register your router
engine.RegisterRoutes(your router)
### start gin-router
engine.run()

## define router
### define Root Router and define your router's path group
var RootRouter = gin-router.NewRouter(gin-router.NewGroup("/hello"))
### make a new file to define your interface logic and implement interface
func init() {
	RootRouter.Register(xrouter.NewRouter(Hello{}))
}

type Hello struct {
}

func (h Hello) Path() string {
	return ""
}

func (h Hello) Method() string {
	return http.MethodGet
}

func (h Hello) Output(ctx xrouter.Context) (interface{}, error) {
	return "hello,world", nil
}



