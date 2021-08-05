package persistence

import (
	"errors"
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/repository"

	"gorm.io/gorm"
)

type AdminRepository struct {
	Conn *gorm.DB
}

func NewAdminRepository(conn *gorm.DB) repository.AdminRepository {
	return &AdminRepository{Conn: conn}
}

//- 作成
func (ur *AdminRepository) CreateAdmin(admin *entity.Admin) (*entity.Admin, error) {
	tx := ur.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(admin).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return admin, tx.Commit().Error
}

//- 更新
func (ur *AdminRepository) UpdateAdmin(admin *entity.Admin) (*entity.Admin, error) {
	tx := ur.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Updates(admin).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return admin, tx.Commit().Error
}

//- 削除
func (ur *AdminRepository) DeleteAdmin(admin *entity.Admin) error {
	tx := ur.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(admin).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (ur *AdminRepository) FindByEmail(email string) (*entity.Admin, error) {
	admin := &entity.Admin{}
	if err := ur.Conn.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, errors.New("レコードが存在しません。")
	}
	return admin, nil
}

//- IDで取得
func (ur *AdminRepository) FindByID(id uint) (*entity.Admin, error) {
	admin := &entity.Admin{ID: id}

	if err := ur.Conn.First(&admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}
