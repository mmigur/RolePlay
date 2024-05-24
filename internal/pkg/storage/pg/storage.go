package pg

import (
	"RolePlayModule/internal/pkg/models"
	"RolePlayModule/internal/pkg/services"
	"RolePlayModule/internal/utils/config"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"time"
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
	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Coupon{}, &models.Order{}, &models.OrderDetail{}, &models.UserCoupon{}, &models.CodeForEmail{}, &models.Category{})
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

//func createDataType(db *gorm.DB, dataType string, values []string) error {
//	valuesStr := strings.Join(values, "', '")
//	query := fmt.Sprintf("CREATE TYPE %s AS ENUM ('%s')", dataType, valuesStr)
//	err := db.Exec(query).Error
//	if err != nil {
//		if strings.Contains(err.Error(), "already exists") {
//			fmt.Printf("Data type %s already exists\n", dataType)
//			return nil
//		}
//		return err
//	}
//	fmt.Printf("Data type %s created successfully\n", dataType)
//	return nil
//}

func (s *Storage) cleanCodes(email string) error {
	err := s.db.Where("email = ?", email).Delete(&models.CodeForEmail{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) isUserExist(email string) (bool, error) {
	err := s.db.Where("email = ?", email).First(&models.User{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *Storage) CheckUser(email string, cfg config.Config) (bool, error) {
	var foundUser models.User
	err := s.db.Where("email = ?", email).First(&foundUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := s.cleanCodes(email)
			if err != nil {
				return false, err
			}
			code := services.GenerateRandomCode()
			var codeForEmail = models.CodeForEmail{
				Email:     email,
				Code:      code,
				CreatedAt: time.Now().Format(time.RFC3339),
			}

			err = s.db.Create(&codeForEmail).Error
			if err != nil {
				s.db.Delete(&codeForEmail)
				return false, err
			}
			err = services.SendCodeToEmailService(cfg, code, email)
			if err != nil {
				return false, err
			}

			return false, nil
		}
		return false, err
	}
	//code := services.GenerateRandomCode()
	//
	//codeForEmail := models.CodeForEmail{
	//	Email:     email,
	//	Code:      code,
	//	CreatedAt: time.Now().Format(time.RFC3339),
	//}
	//err = services.SendCodeToEmailService(cfg, code, email)
	//if err != nil {
	//	return false, err
	//}
	//err = s.db.Create(&codeForEmail).Error
	return true, nil
}

func (s *Storage) CheckCode(email, code string) (bool, error) {
	var foundCode models.CodeForEmail
	err := s.db.Where("email = ? AND code = ?", email, code).First(&foundCode).Error
	if err != nil {
		return false, err
	}
	err = s.db.Delete(&foundCode).Error
	if err != nil {
		return false, err
	}

	var newUser models.User
	newUser.Email = email
	err = s.db.Create(&newUser).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Storage) FillProfile(user models.User, cfg config.Config) (string, error) {
	var foundUser models.User
	err := s.db.Where("email = ?", user.Email).First(&foundUser).Error
	if err != nil {
		return "", err
	}
	foundUser.FirstName = user.FirstName
	foundUser.MiddleName = user.MiddleName
	foundUser.LastName = user.LastName
	foundUser.Address = user.Address
	foundUser.Username = user.Username
	foundUser.Password = user.Password
	err = s.db.Save(&foundUser).Error
	if err != nil {
		return "", err
	}

	token, err := services.GenerateUserToken([]byte(cfg.SecretKey), foundUser)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Storage) SendCodeAgain(email string, cfg config.Config) error {
	err := s.db.Where("email = ?", email).First(&models.CodeForEmail{}).Error
	if err != nil {
		return err
	}
	err = s.cleanCodes(email)
	if err != nil {
		return err
	}
	var newCode models.CodeForEmail
	code := services.GenerateRandomCode()
	err = services.SendCodeToEmailService(cfg, code, email)
	if err != nil {
		return err
	}
	newCode.Email = email
	newCode.Code = code
	newCode.CreatedAt = time.Now().Format(time.RFC3339)
	err = s.db.Create(&newCode).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) CreateProduct(product models.Product) error {
	err := s.db.Where("id = ?", product.CategoryId).First(&models.Category{}).Error
	if err != nil {
		return err
	}
	err = s.db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := s.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *Storage) GetProductsByCategory(categoryID uint) ([]models.Product, error) {
	var products []models.Product
	if err := s.db.Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (s *Storage) CheckPassword(email, password string, cfg config.Config) (string, error) {
	var foundUser models.User
	err := s.db.Where("email = ? AND password = ?", email, password).First(&foundUser).Error
	if err != nil {
		return "", err
	}
	token, err := services.GenerateUserToken([]byte(cfg.SecretKey), foundUser)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Storage) CreateCategory(name string) error {
	var newCategory models.Category
	newCategory.Name = name
	err := s.db.Create(&newCategory).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetProductById(productId uint) (models.Product, error) {
	var foundProduct models.Product
	err := s.db.Where("id = ?", productId).First(&foundProduct).Error
	if err != nil {
		return models.Product{}, err
	}

	return foundProduct, nil
}

func (s *Storage) GetProfileInfo(userId uint) (models.User, error) {
	var foundUser models.User
	err := s.db.Where("id = ?", userId).First(&foundUser).Error
	if err != nil {
		return models.User{}, err
	}
	return foundUser, nil
}

func (s *Storage) CreateOrder(productsIds map[uint]int, userId uint) error {
	// Проверка наличия всех продуктов
	for productId := range productsIds {
		var product models.Product
		if err := s.db.Where("id = ?", productId).First(&product).Error; err != nil {
			return errors.New("product ID " + strconv.Itoa(int(productId)) + " does not exist")
		}
	}

	newOrder := models.Order{
		UserId: userId,
	}

	if err := s.db.Create(&newOrder).Error; err != nil {
		return err
	}

	var totalPrice float32
	for productId, productCount := range productsIds {
		var product models.Product
		if err := s.db.Where("id = ?", productId).First(&product).Error; err != nil {
			return err
		}

		orderDetail := models.OrderDetail{
			OrderID:      newOrder.ID,
			ProductID:    productId,
			ProductCount: productCount,
		}

		if err := s.db.Create(&orderDetail).Error; err != nil {
			return err
		}

		totalPrice += product.Price * float32(productCount)
	}

	// Обновление общей стоимости заказа
	newOrder.TotalPrice = totalPrice
	if err := s.db.Save(&newOrder).Error; err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetOrders(userId uint) ([]models.Order, error) {
	var orders []models.Order
	err := s.db.Where("user_id = ?", userId).Preload("OrderDetails.Product").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
