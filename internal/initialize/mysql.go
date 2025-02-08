package initialize

import (
	"fmt"
	"net/url"
	"time"

	"github.com/quangdvn/go-ec/global"
	"github.com/quangdvn/go-ec/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySql() {
	m := global.Config.Mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Dbname,
		url.QueryEscape("Asia/Tokyo"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	checkPanicError(err, "failed to connect database")

	global.Logger.Info("mysql connected")
	global.Mdb = db

	// Set pool - Pre-open connection
	setPool()
	migrateTables()
}

func setPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		global.Logger.Error("failed to set pool", zap.Error(err))
		panic("failed to set pool")
	}
	// Purpose: Sets the maximum amount of time a connection can stay idle (unused) before being closed.
	// Why?: Prevents the database from holding too many idle connections, which can waste resources.
	sqlDb.SetConnMaxIdleTime(time.Duration(m.ConnMaxIdleTime)) // 5-10m

	// Purpose: Defines the maximum number of open connections to the database.
	// Why?: Helps manage MySQL connection limits and prevents overloading the server.
	sqlDb.SetMaxOpenConns(m.MaxOpenConns) // 50-200

	// Purpose: Sets the maximum duration a connection can stay open before being closed and replaced.
	// Why?: Prevents long-lived connections from staying open indefinitely, which can cause stale connections or memory leaks.
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime)) // 30m-1h
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		global.Logger.Error("failed to migrate tables", zap.Error(err))
		fmt.Println("Migrate error", err.Error())
	}
}

func checkPanicError(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(fmt.Sprintf("%s: %s", errString, err.Error()))
	}
}
