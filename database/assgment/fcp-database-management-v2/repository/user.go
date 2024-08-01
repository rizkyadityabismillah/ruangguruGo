package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	// Gunakan metode Create dari GORM untuk menambahkan pengguna ke dalam tabel
	result := u.db.Create(&user)
	// Periksa apakah ada error saat menyimpan data
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *userRepository) CheckAvail(user model.User) error {
	// Inisialisasi variabel untuk menyimpan hasil pencarian
	var existingUser model.User
	// Gunakan metode Where dari GORM untuk memeriksa ketersediaan pengguna berdasarkan username dan password
	result := u.db.Where("username = $1 AND password = $2", user.Username, user.Password).First(&existingUser)
	// Jika tidak ditemukan (record not found), kembalikan error
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	// Jika ditemukan, kembalikan nil
	return nil
}
