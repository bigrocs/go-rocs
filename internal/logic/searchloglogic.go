package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogLogic {
	return &SearchLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLogLogic) SearchLog(in *pb.SearchLogReq) (*pb.SearchLogResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchLogResp{}, nil
}
