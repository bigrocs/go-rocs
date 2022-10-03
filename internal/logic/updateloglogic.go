package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogLogic {
	return &UpdateLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogLogic) UpdateLog(in *pb.UpdateLogReq) (*pb.UpdateLogResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateLogResp{}, nil
}
