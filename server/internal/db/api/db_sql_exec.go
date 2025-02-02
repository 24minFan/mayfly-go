package api

import (
	"mayfly-go/internal/db/application"
	"mayfly-go/internal/db/domain/entity"
	"mayfly-go/pkg/ginx"
	"mayfly-go/pkg/req"
)

type DbSqlExec struct {
	DbSqlExecApp application.DbSqlExec
}

func (d *DbSqlExec) DbSqlExecs(rc *req.Ctx) {
	queryCond, page := ginx.BindQueryAndPage(rc.GinCtx, new(entity.DbSqlExecQuery))
	queryCond.CreatorId = rc.LoginAccount.Id
	rc.ResData = d.DbSqlExecApp.GetPageList(queryCond, page, new([]entity.DbSqlExec))
}
