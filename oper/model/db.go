package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type DbRepoItem struct {
	Id          uint   `gorm:"primaryKey; not null"`
	Name        string `gorm:"type:varchar(255); index; not null"`
	Path        string `gorm:"type:varchar(255); not null"`
	Author      string `gorm:"type:varchar(255); not null"`
	Version     string `gorm:"type:varchar(255); index; not null"`
	Description string `gorm:"type:text"`
	RepoId      uint   `gorm:"index; not null"`
	Created     time.Time
	LastUpdated time.Time
}

type DbRepo struct {
	Id  uint   `gorm:"primaryKey; not null"`
	Url string `gorm:"type:varchar(255); not null"`
}

func NewSqlite(path string) {
	os.Remove(path)
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&DbRepoItem{}, &DbRepo{})
	if err != nil {
		panic(err)
	}
	ins, _ := db.DB()
	ins.Close()
}
