package xrouter

import "github.com/gin-gonic/gin"

type Context struct {
	// 上下文
	*gin.Context
}

const (
	CONTENT_TYPE_Disposition = "application/octet-stream"
	CONTENT_TYPE_JSON        = "application/json"
)

const (
	HEADER_DISPOSITION = "content-disposition"
)
