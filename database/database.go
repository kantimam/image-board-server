package database

import (
	"gorm.io/gorm"
)

// DBConn shared reference to the database
var (
	DBConn *gorm.DB
)
