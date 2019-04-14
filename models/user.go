package models

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	ID       int64  `orm:"column(id);pk" json:"id"`
	Name     string `orm:"size(128)"`
	Email    string `orm:"size(128)"`
	Password string `orm:"size(128)"`
	IsActive int64
	Comment  string    `orm:"null;size(16)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
	Updated  time.Time `orm:"auto_now_add;type(datetime)" json:"updated"`
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	m.IsActive = 1
	if hashPass, err := PasswordHash(m.Password); err == nil {
		m.Password = hashPass
		id, err = o.Insert(m)
	}
	return id, err
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(u *User) (err error) {
	o := orm.NewOrm()
	if err = o.QueryTable(new(User)).Filter("ID", u.ID).RelatedSel().One(u); err == nil {
		return nil
	}
	return err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		m.Updated = time.Now()
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
			return nil
		}
	}
	return errors.New("Faild to Update")
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	v := User{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return nil
}

func PasswordHash(pw string) (hashPass string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// AuthenticateUser... 認証処理
func AuthenticateUser(u *User) (b bool) {
	o := orm.NewOrm()
	if hashPass, err := PasswordHash(u.Password); err != nil {
		u.Password = hashPass
		o.QueryTable(new(User)).Filter("ID", u.ID).Filter("Password", u.Password).Filter("IsActive", 1)
	} else {
		u = nil
	}
	return u != nil
}
