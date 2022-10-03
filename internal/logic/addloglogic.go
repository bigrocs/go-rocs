package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogLogic {
	return &AddLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------logs-----------------------
func (l *AddLogLogic) AddLog(in *pb.AddLogReq) (*pb.AddLogResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddLogResp{}, nil
}
