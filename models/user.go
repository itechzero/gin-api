package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	UUID        uuid.UUID `gorm:"column:uuid;comment:用户UUID" json:"uuid"`
	Username    string    `gorm:"column:username;comment:用户登录名" json:"username"`
	Password    string    `gorm:"column:password;comment:用户登录密码" json:"-"`
	NickName    string    `gorm:"column:nick_name;default:系统用户;comment:用户昵称" json:"nickName"`
	HeaderImg   string    `gorm:"column:header_img;comment:用户头像" json:"headerImg"`
	AuthorityId string    `gorm:"column:authority_id;default:888;comment:用户角色ID" json:"authorityId"`
}

func (User) TableName() string {
	return "sys_users"
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
