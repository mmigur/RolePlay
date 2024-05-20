package pg

import (
	"RolePlayModule/internal/pkg/models"
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

	err = createDataType(db, "product_category_type", []string{string(models.MeatCategory), string(models.VegetablesCategory), string(models.MilkCategory), string(models.FishCategory), string(models.GroatsCategory), string(models.NutsCategory)})
	if err != nil {
		fmt.Println("news_category_type is already exist")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("table 'users' is already exist")
	}
	err = db.AutoMigrate(&models.Coupon{})
	if err != nil {
		fmt.Println("table 'coupons' is already exist")
	}
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		fmt.Println("table 'products' is already exist")
	}
	err = db.AutoMigrate(&models.OrderRecord{})
	if err != nil {
		fmt.Println("table 'order_records' is already exist")
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
