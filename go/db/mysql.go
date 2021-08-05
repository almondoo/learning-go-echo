package db

import (
	"learning-go-echo/db/seed"
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/repository"
	"learning-go-echo/infrastructure/persistence"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repositories struct {
	User         repository.UserRepository
	Admin        repository.AdminRepository
	Product      repository.ProductRepository
	ProductImage repository.ProductImageRepository
	Role         repository.RoleRepository
	Article      repository.ArticleRepository
	Artist       repository.ArtistRepository
	DB           *gorm.DB
}

//- connetion確認func
func InitDB() (*Repositories, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      false,
		},
	)

	dns := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@" + os.Getenv("DB_PROTO") + "/" + os.Getenv("DB") + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	//- 自動的にテーブルを生成する
	db.AutoMigrate(&entity.Admin{}, &entity.Article{}, &entity.Artist{}, &entity.ArtistAccessLog{}, &entity.ArtistCategory{}, &entity.DeleteUser{}, &entity.PasswordReset{}, &entity.ProductAccessLog{}, &entity.ProductFavorite{}, &entity.ProductCategory{}, &entity.ProductImage{}, &entity.Product{}, &entity.Role{}, &entity.SubscriptionItem{}, &entity.Subscription{}, &entity.User{})

	return &Repositories{
		User:         persistence.NewUserRepository(db),
		Admin:        persistence.NewAdminRepository(db),
		Product:      persistence.NewProductRepository(db),
		ProductImage: persistence.NewProductImageRepository(db),
		Role:         persistence.NewRoleRepository(db),
		Article:      persistence.NewArticleRepository(db),
		Artist:       persistence.NewArtistRepository(db),
		DB:           db,
	}, nil

}

func (r *Repositories) Close() error {
	db, err := r.DB.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (r *Repositories) Seeder() {
	db := r.DB
	seed.NewProductCategorySeeder(db).Seeder()
	seed.NewArtistCategorySeeder(db).Seeder()
	seed.NewRoleSeeder(db).Seeder()
	seed.NewAdminSeeder(db).Seeder()
	seed.NewArtistSeeder(db).Seeder()
	seed.NewUserSeeder(db).Seeder()
	seed.NewProductSeeder(db).Seeder()
	seed.NewProductImageSeeder(db).Seeder()
	seed.NewArticleSeeder(db).Seeder()
}
