package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigByIdLogic {
	return &GetConfigByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConfigByIdLogic) GetConfigById(in *pb.GetConfigByIdReq) (*pb.GetConfigByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetConfigByIdResp{}, nil
}
