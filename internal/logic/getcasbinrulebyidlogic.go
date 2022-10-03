package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCasbinRuleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCasbinRuleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCasbinRuleByIdLogic {
	return &GetCasbinRuleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCasbinRuleByIdLogic) GetCasbinRuleById(in *pb.GetCasbinRuleByIdReq) (*pb.GetCasbinRuleByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCasbinRuleByIdResp{}, nil
}
