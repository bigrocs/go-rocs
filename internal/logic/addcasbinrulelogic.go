package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCasbinRuleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCasbinRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCasbinRuleLogic {
	return &AddCasbinRuleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------casbinRule-----------------------
func (l *AddCasbinRuleLogic) AddCasbinRule(in *pb.AddCasbinRuleReq) (*pb.AddCasbinRuleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddCasbinRuleResp{}, nil
}
