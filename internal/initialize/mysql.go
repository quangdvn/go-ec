package initialize

import (
	"fmt"
	"net/url"
	"time"

	"github.com/quangdvn/go-ec/global"
	"github.com/quangdvn/go-ec/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
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
	genTableDAO()
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
		// &po.User{},
		// &po.Role{},
		&model.GoCrmUserV2{},
	)
	if err != nil {
		global.Logger.Error("failed to migrate tables", zap.Error(err))
		fmt.Println("Migrate error", err.Error())
	}
}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb)            // reuse your gorm db
	g.GenerateModel("go_crm_user") // Will create go_crm_user.gen.go
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}

func checkPanicError(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(fmt.Sprintf("%s: %s", errString, err.Error()))
	}
}
