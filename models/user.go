package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID     int64  `gorm:"column:id;primaryKey" json:"-" form:"id"`
	Name   string `gorm:"column:name;" json:"name" form:"name" binding:"required" uri:"name"`
	Email  string `gorm:"column:email;" json:"email" form:"email"`
	Active int    `gorm:"column:active;" json:"active" form:"active"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// TODO
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	// TODO
	return
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// TODO
	return
}

func (u *User) AfterSave(tx *gorm.DB) (err error) {
	// TODO
	return
}

//func (user User) GetUserList() interface{} {
//	users := make([]User, 0)
//	//users := make([]User, 0)
//	db := libs.ConnectMySQL()
//	rows, err := db.Query("select `id`,`name`,`email`,`active` from `user`")
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	for rows.Next() {
//		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Active)
//		users = append(users, user)
//		//data, err := json.Marshal(user)
//		//if err != nil {
//		//	panic (err)
//		//}
//		//fmt.Println(string(data))
//	}
//	return users
//	//fmt.Println(users)
//}
//
//func (user User) CreateUser(name string, email string, password string) bool {
//	db := core.ConnectMySQL()
//	stmt, err := db.Prepare("INSERT `user` SET name=?, email=?, password=?")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	res, err := stmt.Exec(name, email, password)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	fmt.Println(res)
//
//	return true
//}
