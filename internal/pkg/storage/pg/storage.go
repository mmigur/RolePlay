package pg

import (
	"RolePlayModule/internal/utils/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func buildDSN(cfg *config.Config) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	return dsn
}

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := buildDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MustNewPostgresDB(cfg *config.Config) *gorm.DB {
	db, err := NewPostgresDB(cfg)
	if err != nil {
		panic(err)
	}

	return db
}

func createDataType(db *gorm.DB, dataType string, values []string) error {
	valuesStr := strings.Join(values, "', '")
	query := fmt.Sprintf("CREATE TYPE %s AS ENUM ('%s')", dataType, valuesStr)
	err := db.Exec(query).Error
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			fmt.Printf("Data type %s already exists\n", dataType)
			return nil
		}
		return err
	}
	fmt.Printf("Data type %s created successfully\n", dataType)
	return nil
}
