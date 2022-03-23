package boltdb

import (
	"PocketBot/pkg/repository"
	"errors"
	"github.com/boltdb/bolt"
	"strconv"
)

// Вся логика для работы с базой данных

type TokenRepository struct {
	db *bolt.DB
}

func NewTokenRepository(db *bolt.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

// Сохранить токен в бд
func (r *TokenRepository) Save(chatId int64, token string, bucket repository.Bucket) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToByte(chatId), []byte(token))
	})

}

// Извлечь токен из бд
func (r *TokenRepository) Get(chatId int64, bucket repository.Bucket) (string, error) {
	var token string
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		token = string(b.Get(intToByte(chatId)))
		return nil
	})
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", errors.New("token not found")
	}
	return token, nil
}

// Преобразование типов int64 в byte
func intToByte(x int64) []byte {
	return ([]byte(strconv.FormatInt(x, 10)))
}
