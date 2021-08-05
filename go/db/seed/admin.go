package seed

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/infrastructure/security"

	"gorm.io/gorm"
)

type adminSeeder struct {
	conn *gorm.DB
}

func NewAdminSeeder(db *gorm.DB) *adminSeeder {
	return &adminSeeder{conn: db}
}

func (as *adminSeeder) Seeder() {
	pass, err := security.Hash("testuser")
	if err != nil {
		return
	}
	as.create(entity.Admin{
		ID:       1,
		RoleID:   1,
		Name:     "m t",
		Email:    "admin@gmail.com",
		Password: pass,
	})
	as.create(entity.Admin{
		ID:       2,
		RoleID:   10,
		Name:     "管理者",
		Email:    "admin2@exsample.com",
		Password: pass,
	})
	as.create(entity.Admin{
		ID:       3,
		RoleID:   20,
		Name:     "閲覧者",
		Email:    "admin3@exsample.com",
		Password: pass,
	})

}

func (as *adminSeeder) create(entity entity.Admin) error {
	tx := as.conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
