package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFrontPermitByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFrontPermitByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFrontPermitByIdLogic {
	return &GetFrontPermitByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFrontPermitByIdLogic) GetFrontPermitById(in *pb.GetFrontPermitByIdReq) (*pb.GetFrontPermitByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFrontPermitByIdResp{}, nil
}
