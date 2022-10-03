package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFrontPermitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFrontPermitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFrontPermitLogic {
	return &AddFrontPermitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------frontPermits-----------------------
func (l *AddFrontPermitLogic) AddFrontPermit(in *pb.AddFrontPermitReq) (*pb.AddFrontPermitResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddFrontPermitResp{}, nil
}
