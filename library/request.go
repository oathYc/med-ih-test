package library

import (
	"ihtest/library/global"

	"git.medlinker.com/golang/xerror"
	"git.medlinker.com/golang/xerror/xcode"
	"git.medlinker.com/service/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Request struct {
	ctx *gin.Context
}

// 初始化Request
func NewRequest(ctx *gin.Context) *Request {
	return &Request{
		ctx: ctx,
	}
}

// 参数解析
func (r *Request) Bind(req interface{}) error {
	if err := r.ShouldBind(req); nil != err {
		global.Log.Errorf("[ Request ] bind %v req: %+v, err: %v", utils.TypeNameOf(req), req, err)
		return xerror.WithXCode(xcode.ParamInvalid)
	}

	return nil
}

func (r *Request) ShouldBind(req interface{}) error {
	bType := binding.Default(r.ctx.Request.Method, r.ctx.ContentType())
	if binding.JSON == bType {
		err := r.ctx.ShouldBindBodyWith(req, binding.JSON)
		if nil != err {
			return err
		}
	} else {
		err := r.ctx.ShouldBindWith(req, bType)
		if nil != err {
			return err
		}
	}

	return nil
}
