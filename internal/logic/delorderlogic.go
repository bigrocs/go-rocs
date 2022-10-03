package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOrderLogic {
	return &DelOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOrderLogic) DelOrder(in *pb.DelOrderReq) (*pb.DelOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelOrderResp{}, nil
}
