package model

import (
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
