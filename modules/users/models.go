package users

import "gorm.io/gorm"

type User struct {
	ID          uint   `gorm:"primaryKey"`
	FirstName   string `gorm:"type:varchar(64); not null"`
	LastName    string `gorm:"type:varchar(64); not null"`
	Email       string `gorm:"type:varchar(128); not null"`
	UserGroupId uint   `gorm:"not null"`
	UserGroup   UserGroup
	Password    string

	gorm.Model
}

type UserMeta struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uint `gorm:"index:user_meta_unique,unique"`
	User      User
	MetaKey   string `gorm:"type:varchar(128); index:user_meta_unique,unique"`
	MetaValue string
}

type UserGroup struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"not null"`

	gorm.Model
}

func Migrate(DB *gorm.DB) {

	// run migrations
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserMeta{})
	DB.AutoMigrate(&UserGroup{})
}
