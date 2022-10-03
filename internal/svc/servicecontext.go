package svc

import (
	"github.com/go-rocs/user-rpc/internal/config"
	"github.com/go-rocs/user-rpc/types/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUsersModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
