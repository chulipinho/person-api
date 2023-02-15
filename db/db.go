package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database[T interface{}] struct {
	DB    *gorm.DB
	model *T
}

func NewDatabase[T interface{}](model *T) *Database[T] {
	dsn := "host=localhost user=root password=example dbname=person port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create database")
	}

	db.AutoMigrate(model)
	return &Database[T]{db, model}
}

func (db Database[T]) Get() ([]T, error) {
	var data []T

	result := db.DB.Find(&data)
	if err := result.Error; err != nil {
		return nil, result.Error
	}

	return data, nil
}

func (db Database[T]) GetById(id int) (*T, error) {
	var data T
	result := db.DB.First(&data, id)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (db Database[T]) Post(data T) error {
	result := db.DB.Create(&data)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (db Database[T]) Put(id int, data T) error {
	var p T
	result := db.DB.First(&p, id)
	if err := result.Error; err != nil {
		return db.Post(data)
	}

	result = db.DB.Model(&p).Updates(&data)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (db Database[T]) Delete(id int) error {
	result := db.DB.Delete(&db.model, id)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}
