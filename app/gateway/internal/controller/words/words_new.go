package words

import (
	"proxima/app/gateway/api/words"
	"proxima/app/gateway/utility"
	v1 "proxima/app/word/api/words/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type ControllerV1 struct {
	WordsClient v1.WordsClient
}

func NewV1() words.IWordsV1 {
	var conn = grpcx.Client.MustNewGrpcClientConn("word", grpcx.Client.ChainUnary(
		utility.GrpcClientTimeout,
	))

	return &ControllerV1{
		WordsClient: v1.NewWordsClient(conn),
	}
}
