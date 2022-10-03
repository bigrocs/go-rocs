package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSecretKeyByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSecretKeyByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSecretKeyByIdLogic {
	return &GetSecretKeyByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSecretKeyByIdLogic) GetSecretKeyById(in *pb.GetSecretKeyByIdReq) (*pb.GetSecretKeyByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetSecretKeyByIdResp{}, nil
}
