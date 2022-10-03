package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogLogic {
	return &DelLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLogLogic) DelLog(in *pb.DelLogReq) (*pb.DelLogResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelLogResp{}, nil
}
