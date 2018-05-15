package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"chat/lib"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DbParams struct {
	DATABASE_HOST string `default:"localhost"`
}

func InitDB() *gorm.DB {
	db_host := lib.Getenv("DATABASE_HOST", "localhost")
	db_port := lib.Getenv("DATABASE_PORT", "32769")
	db_user := lib.Getenv("DATABASE_USER", "postgres")
	db_name := lib.Getenv("DATABASE_NAME", "chat-development")
	db_password := lib.Getenv("DATABASE_PASSWORD", "pineapple")
	sslmode := lib.Getenv("sslmode", "disable")
	gorm_params := "host="+db_host+" port="+db_port+" user="+db_user+" dbname="+db_name+" password="+db_password+" sslmode="+sslmode
	fmt.Println(gorm_params)
	db, err := gorm.Open("postgres", gorm_params)
	if err != nil {
		panic(fmt.Errorf("connected to open database, error: (%s)", err))
	}
	if err := db.DB().Ping(); err != nil {
		panic(fmt.Errorf("connected to ping database, error: (%s)", err))
	}
	Migrate(db)
	return db
}
