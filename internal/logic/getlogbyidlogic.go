package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogByIdLogic {
	return &GetLogByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogByIdLogic) GetLogById(in *pb.GetLogByIdReq) (*pb.GetLogByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetLogByIdResp{}, nil
}
