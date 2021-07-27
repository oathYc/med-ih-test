package boot

import (
	"context"
	"fmt"
	"ihtest/library/global"

	"git.medlinker.com/golang/xerror"
	"git.medlinker.com/service/grpcwrapper"
	"google.golang.org/grpc"
)

func makeClientWithConfig(addr, title string) (*grpc.ClientConn, error) {
	if "" == addr {
		return nil, xerror.New(fmt.Sprintf("[%s] GRPC 地址为空", title))

	}

	opt := grpcwrapper.NewClientOpt()

	gc, err := opt.DialXEContext(context.Background(), addr, grpc.WithBlock())
	if nil != err {
		return nil, xerror.Wrap(err, fmt.Sprintf("创建（%s) GRPC 客户端失败[%s]", title))
	}

	return gc, nil

}

func InitClient() error {
	var err error

	global.MedTestConn, err = makeClientWithConfig(global.Config.Service.MedTestAddr, "med_test")
	if nil != err {
		return err
	}

	return nil
}
