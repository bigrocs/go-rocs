package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFrontPermitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFrontPermitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFrontPermitLogic {
	return &UpdateFrontPermitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFrontPermitLogic) UpdateFrontPermit(in *pb.UpdateFrontPermitReq) (*pb.UpdateFrontPermitResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateFrontPermitResp{}, nil
}
