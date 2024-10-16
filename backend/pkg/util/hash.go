package util

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(hash, password string) bool
}

// BcryptHasher Bcrypt 哈希实现
type BcryptHasher struct{}

// HashPassword 生成密码哈希
func (h *BcryptHasher) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// ComparePassword 验证密码哈希
func (h *BcryptHasher) ComparePassword(hashedPassword, passwrod string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwrod))
	return err == nil
}

var _ PasswordHasher = (*BcryptHasher)(nil)

func HashSPassword(hasher PasswordHasher, password string) (string, error) {
	return hasher.HashPassword(password)
}
func CompareSPassword(hasher PasswordHasher, hashedPassword, password string) bool {
	return hasher.ComparePassword(hashedPassword, password)
}

var Bcrypt = &BcryptHasher{}

func HashPassword(password string) (string, error) {
	return HashSPassword(Bcrypt, password)
}

func ComparePassword(hashedPassword, password string) bool {
	return CompareSPassword(Bcrypt, hashedPassword, password)
}
