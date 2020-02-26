package main

import (
	"fmt"
	"github.com/i2eco/muses"
	"github.com/i2eco/muses/pkg/cmd"
	"github.com/i2eco/muses/pkg/database/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Uid  int
	Name string
}

func main() {
	app := muses.Container(
		cmd.Register,
		mysql.Register,
	)
	err := app.Run()
	if err != nil {
		panic(err)
	}

	db := mysql.Caller("muses")
	err = CreateTable(db)
	if err != nil {
		panic(err)
	}
	err = CreateUser(db)
	if err != nil {
		panic(err)
	}
	err = UpdateUser(db)
	if err != nil {
		panic(err)
	}
	var u *User
	u, err = FindUser(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("u: ", u)

	err = DeleteUser(db)
	if err != nil {
		panic(err)
	}

}

func CreateTable(db *gorm.DB) (err error) {
	db.SingularTable(true)
	if db.Error != nil {
		return db.Error
	}
	Model := User{}

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Model)
	if db.Error != nil {
		return db.Error
	}
	return
}

func CreateUser(db *gorm.DB) (err error) {
	return db.Create(&User{
		Uid:  1,
		Name: "muses",
	}).Error
}

func UpdateUser(db *gorm.DB) (err error) {
	return db.Model(User{}).Where("uid = ?", 1).Updates(map[string]interface{}{
		"name": "musesgo",
	}).Error
}

func FindUser(db *gorm.DB) (user *User, err error) {
	user = &User{}
	err = db.Where("uid = ?", 1).Find(user).Error
	return
}

func DeleteUser(db *gorm.DB) (err error) {
	err = db.Where("uid = ?", 1).Delete(&User{}).Error
	return
}
