package mysql

import (
	"github.com/atopx/serverlib/database/infrastruct"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQLConnect 创建MySQL连接
func NewMySQLConnect(config infrastruct.Config) (*gorm.DB, error) {
	dsn := config.BuildDSN()
	db, err := gorm.Open(mysql.Open(dsn), config.GormConfig)
	if err != nil {
		return db, err
	}
	sdb, err := db.DB()
	if err != nil {
		return db, err
	}
	if config.MaxIdleConn > 0 {
		sdb.SetMaxIdleConns(config.MaxIdleConn)
	}
	if config.MaxOpenConn > 0 {
		sdb.SetMaxOpenConns(config.MaxOpenConn)
	}
	if config.ConnMaxLifeTime > 0 {
		sdb.SetConnMaxLifetime(config.ConnMaxLifeTime)
	}
	return db, nil
}
