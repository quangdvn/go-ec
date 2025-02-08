package benchmark

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestUser struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

func BenchmarkMaxOpenConns10(b *testing.B) {
	dsn := "admin:mysql@tcp(localhost:8811)/go-ec?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Check and drop table if it exists
	resetTable(db)

	// Set max open connections to 10
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}
	sqlDB.SetMaxOpenConns(10)
	defer sqlDB.Close()

	// Benchmark test logic
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns1(b *testing.B) {
	dsn := "admin:mysql@tcp(localhost:8811)/go-ec?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Check and drop table if it exists
	resetTable(db)

	// Set max open connections to 1
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}
	sqlDB.SetMaxOpenConns(1)
	defer sqlDB.Close()

	// Benchmark test logic
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

// resetTable checks if the table exists, drops it if it does, and performs auto-migration.
func resetTable(db *gorm.DB) {
	// Check if the table exists
	if db.Migrator().HasTable(&TestUser{}) {
		if err := db.Migrator().DropTable(&TestUser{}); err != nil {
			log.Fatalf("failed to drop table: %v", err)
		}
		log.Println("Table dropped successfully.")
	}

	// Auto-migrate the table
	if err := db.AutoMigrate(&TestUser{}); err != nil {
		log.Fatalf("failed to auto-migrate table: %v", err)
	}
	log.Println("Table auto-migrated successfully.")
}

// insertRecord inserts a record into the table for benchmarking.
func insertRecord(b *testing.B, db *gorm.DB) {
	record := TestUser{Name: "Benchmark User"}
	if err := db.Create(&record).Error; err != nil {
		b.Errorf("failed to insert record: %v", err)
	}
}
