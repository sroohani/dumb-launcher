package db

import (
	"gorm.io/gorm"
)

type AppGroup struct {
	gorm.Model
	Name      string
	ParentID  *uint
	Parent    *AppGroup  `gorm:"foreignKey:ParentID"`
	SubGroups []AppGroup `gorm:"foreignKey:ParentID"`
	Apps      []App      `gorm:"foreignKey:GroupID"`
	IsPinned  bool
	Shortcut  string
}

type App struct {
	gorm.Model
	Name       string
	Command    string
	Parameters string
	GroupID    uint
	Group      AppGroup `gorm:"foreignKey:GroupID"`
	IsPinned   bool
	Shortcut   string
}
