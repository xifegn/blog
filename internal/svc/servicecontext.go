package svc

import (
	"blog/bkmodel/dao/query"
	"blog/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	Bkmodel *query.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	return &ServiceContext{
		Config:  c,
		Bkmodel: query.Use(db),
	}
}
