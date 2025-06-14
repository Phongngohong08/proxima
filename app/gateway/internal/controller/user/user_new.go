package user

import (
	"proxima/app/gateway/api/user"
	"proxima/app/gateway/utility"
	v1 "proxima/app/user/api/account/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type ControllerV1 struct {
	AccountClient v1.AccountClient
}

func NewV1() user.IUserV1 {
	var conn = grpcx.Client.MustNewGrpcClientConn("user", grpcx.Client.ChainUnary(
		utility.GrpcClientTimeout,
	))

	return &ControllerV1{
		AccountClient: v1.NewAccountClient(conn),
	}
}
