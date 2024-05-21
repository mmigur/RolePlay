package pg

import (
	"RolePlayModule/internal/pkg/models"
	"RolePlayModule/internal/pkg/services"
	"RolePlayModule/internal/utils/config"
	"errors"
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

	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Coupon{}, &models.Order{}, &models.OrderDetail{}, &models.UserCoupon{}, &models.CodeForEmail{})
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

func (s *Storage) CheckUser(email string, cfg config.Config) (bool, error) {
	var foundUser models.User
	err := s.db.Where("email = ?", email).First(&foundUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code := services.GenerateRandomCode()
			var codeForEmail = models.CodeForEmail{
				Email: email,
				Code:  code,
			}

			err = s.db.Create(&codeForEmail).Error
			if err != nil {
				s.db.Delete(&codeForEmail)
				return false, err
			}
			//err = services.SendCodeToEmailService(cfg, code, email)
			//if err != nil {
			//	return false, err
			//}

			return false, nil
		}
		return false, err
	}
	code := services.GenerateRandomCode()

	codeForEmail := models.CodeForEmail{
		Email: email,
		Code:  code,
	}
	err = services.SendCodeToEmailService(cfg, code, email)
	if err != nil {
		return false, err
	}
	err = s.db.Create(&codeForEmail).Error
	return true, nil
}
