package util

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(password string, hash string) bool
}

// BcryptHasher Bcrypt 哈希实现
type BcryptHasher struct{}

// HashPassword 生成密码哈希
func (h *BcryptHasher) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// ComparePassword 验证密码哈希
func (h *BcryptHasher) ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
