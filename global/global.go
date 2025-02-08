package global

import (
	"github.com/quangdvn/go-ec/pkg/loggers"
	"github.com/quangdvn/go-ec/pkg/settings"
)

var (
	Config settings.Config
	Logger *loggers.ZapLogger
)

/*
	Config
		Redis
		MySQL
		...
*/
