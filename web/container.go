//+build wireinject

package web

import (
	"github.com/google/wire"
	"gocho/dao"
	"gocho/services"
)

func CreateUserHandler() *UserHandler {
	panic(wire.Build(dao.GetDbConfig,
		dao.NewSqlxDatabase,
		dao.NewUserDao,
		services.NewUserService,
		NewUserHandler,
	))
}
