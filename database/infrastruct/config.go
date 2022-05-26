package infrastruct

import (
	"fmt"
	"time"
)
import "gorm.io/gorm"

type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Charset  string `json:"charset"`

	MaxIdleConn     int           `json:"max_idle_conn"`
	MaxOpenConn     int           `json:"max_open_conn"`
	ConnMaxLifeTime time.Duration `json:"conn_max_life_time"`

	GormConfig *gorm.Config `json:"gormConfig"`
}

func (cfg *Config) BuildDSN() string {
	if cfg.Charset == "" {
		cfg.Charset = "utf8mb4"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.Charset,
	)
}
