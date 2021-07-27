package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"ihtest/library/global"

	"ihtest/library/http/context"

	"github.com/gin-gonic/gin"
)

// 日志中间件
func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data string
		// 获取request data
		if c.Request.Method == http.MethodPost {
			body, err := c.GetRawData()
			if err != nil {
				global.Log.Errorf("[ LogMiddleware ] c.GetRawData err, %v", err)
				return
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			data = string(body)
		}

		begin := time.Now()
		c.Next()
		cost := time.Since(begin)

		errmsg := context.GetErrMsg(c)

		url := c.Request.URL
		remote := c.ClientIP()
		method := c.Request.Method
		traceId := context.GetTraceId(c)

		if len(errmsg) == 0 {
			global.Log.Infof("method:%s, url:%s, ts:%s, remote:%s, errmsg:{%s}, data:%s, traceId:%s", method, url, cost, remote, errmsg, data, traceId)
		} else {
			global.Log.Errorf("method:%s, url:%s, ts:%s, remote:%s, errmsg:{%s}, data:%s, traceId:%s", method, url, cost, remote, errmsg, data, traceId)
		}
	}
}
