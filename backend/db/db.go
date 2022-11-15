package db

import (
	"errors"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"backend/db/models"
)

var Conn *gorm.DB

func acquireConnection() (gorm.Dialector, error) {
	postgresDsn, found := os.LookupEnv("POSTGRES_DSN")

	if found {
		return postgres.Open(postgresDsn), nil
	}

	sqliteDsn, found := os.LookupEnv("SQLITE_DSN")

	if found {
		return sqlite.Open(sqliteDsn), nil
	}

	return nil, errors.New("please set either POSTGRES_DSN or SQLITE_DSN in the environment")
}

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dialector, err := acquireConnection()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(models.All...)

	Conn = db
}
