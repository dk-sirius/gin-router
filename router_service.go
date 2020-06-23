package xrouter

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 路配置
type RouterService struct {
	engine *gin.Engine
}

func InitDefaultEngine() *RouterService {
	return &RouterService{engine: gin.Default()}
}

func InitEngine(engine *gin.Engine) *RouterService {
	return &RouterService{engine: engine}
}

func (e *RouterService) AddLoadHtmlDir(dir string) {
	e.engine.LoadHTMLGlob(dir)
}

// 注册路由
func (e *RouterService) RegisterRoutes(router *Router) {
	if router != nil {
		for _, op := range router.operator {
			if !e.isGroup(op) {
				if op.Path() != "" || op.Method() != "" {
					h := switchToHandleFunc(op.Output)
					e.engine.Handle(op.Method(), "path", h)
				} else {
					fmt.Println("Path and Method not allowed empty")
				}
			} else {
				for child, _ := range router.children {
					e.traverse(op.Path(), child)
				}
			}
		}
	}
}

func (e *RouterService) traverse(path string, child *Router) {
	if child != nil {
		for _, op := range child.operator {
			if op.Path() != "" {
				path = path + op.Path()
			}
			if !e.isGroup(op) {
				if op.Method() != "" {
					handleFunc := switchToHandleFunc(op.Output)
					e.engine.Handle(op.Method(), path, handleFunc)
				} else {
					fmt.Println("Path and Method not allowed empty")
				}
			} else {
				for child, _ := range child.children {
					e.traverse(path, child)
				}
			}
		}
	}
}

func (e *RouterService) isGroup(ob interface{}) bool {
	_, ok := ob.(GroupOperator)
	_, ok1 := ob.(*GroupOperator)
	return ok || ok1
}

func (e *RouterService) Run(addr ...string) error {
	return e.engine.Run(addr...)
}

func switchToHandleFunc(output func(ctx Context) (resp interface{}, err error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := output(Context{ctx})
		if err == nil {
			if resp == nil {
				ctx.JSON(http.StatusNoContent, resp)
				return
			}
			// 返回文件
			if file := retAttachFile(resp); file != nil {
				ctx.Header(HEADER_DISPOSITION, file.GetHeaderValue())
				ctx.Data(http.StatusOK, CONTENT_TYPE_Disposition, file.data)
			} else {
				ctx.JSON(http.StatusOK, resp)
			}
		} else {
			ctx.JSON(500, err.Error())
		}
	}
}

func retAttachFile(resp interface{}) *AttachFile {
	t := reflect.TypeOf(resp)
	if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
		if att, ok := resp.(*AttachFile); ok {
			return att
		}
	} else if t.Kind() != reflect.Ptr && t.Kind() == reflect.Struct {
		if att, ok := resp.(AttachFile); ok {
			return &att
		}
	}
	return nil
}
