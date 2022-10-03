package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigLogic {
	return &AddConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------configs-----------------------
func (l *AddConfigLogic) AddConfig(in *pb.AddConfigReq) (*pb.AddConfigResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddConfigResp{}, nil
}
