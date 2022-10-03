package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelConfigLogic {
	return &DelConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelConfigLogic) DelConfig(in *pb.DelConfigReq) (*pb.DelConfigResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelConfigResp{}, nil
}
