package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCasbinRuleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCasbinRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCasbinRuleLogic {
	return &SearchCasbinRuleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCasbinRuleLogic) SearchCasbinRule(in *pb.SearchCasbinRuleReq) (*pb.SearchCasbinRuleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchCasbinRuleResp{}, nil
}
