package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFrontPermitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelFrontPermitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFrontPermitLogic {
	return &DelFrontPermitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelFrontPermitLogic) DelFrontPermit(in *pb.DelFrontPermitReq) (*pb.DelFrontPermitResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelFrontPermitResp{}, nil
}
