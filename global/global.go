package global

import (
	"github.com/quangdvn/go-ec/pkg/loggers"
	"github.com/quangdvn/go-ec/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	// Store the env variables in config YAML file
	Config settings.Config

	// Pointers for all the global services
	Logger *loggers.ZapLogger
	Mdb    *gorm.DB
	Cache  *redis.Client
)

/*
	Config
		Redis
		MySQL
		...
*/
