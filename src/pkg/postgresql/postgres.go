package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/yigithankarabulut/simplemedia/src/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	Name    string
	Host    string
	Pass    string
	User    string
	Port    string
	Migrate bool
}

func GetEnvFields() (Database, error) {
	var db Database
	if err := godotenv.Load(); err != nil {
		return db, err
	}
	db.Name = os.Getenv("POSTGRES_DB")
	db.Host = os.Getenv("POSTGRES_HOST")
	db.Pass = os.Getenv("POSTGRES_PASSWORD")
	db.User = os.Getenv("POSTGRES_USER")
	db.Port = os.Getenv("POSTGRES_PORT")
	db.Migrate = os.Getenv("POSTGRES_MIGRATE") == "true"
	if db.Name == "" || db.Host == "" || db.Pass == "" || db.User == "" || db.Port == "" {
		return db, fmt.Errorf("database env variables are empty")
	}
	return db, nil
}

func Connect() (*gorm.DB, error) {
	conf, err := GetEnvFields()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Host, conf.User, conf.Pass, conf.Name, conf.Port)
	sqlDB, err := sql.Open("pgx", dsn)

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if conf.Migrate {
		err = AutoMigrate(db)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	migrate := db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
		&model.Likes{},
		&model.Friends{},
	)
	return migrate
}
