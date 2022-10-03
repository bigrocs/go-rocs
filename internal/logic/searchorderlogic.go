package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOrderLogic {
	return &SearchOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOrderLogic) SearchOrder(in *pb.SearchOrderReq) (*pb.SearchOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchOrderResp{}, nil
}
