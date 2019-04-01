package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	DiaryList  map[int64]*Diary
	DiaryRange []*Diary
)

func init() {
	DiaryList = make(map[int64]*Diary)
}

type Diary struct {
	ID       int64 `orm:"column(id);pk" json:"id"`
	Title    string
	Content  string `orm:"null" json:"content"`
	IsActive int64
	Created  time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
	Updated  time.Time `orm:"auto_now_add;type(datetime)" json:"updated"`
}

func AddDiary(diary *Diary) (res int64, err error) {
	o := orm.NewOrm()
	o.Using("diary")
	diary.IsActive = 1
	if res, err := o.Insert(diary); err != nil {
		return res, nil
	}
	return 0, errors.New("Faild")
}

func GetDiary(d *Diary) (err error) {
	o := orm.NewOrm()
	if err := o.Read(d); err != nil {
		return errors.New("Diary not exists")
	}
	return nil
}

func GetAllDiary() []*Diary {
	o := orm.NewOrm()
	o.QueryTable("diary").Filter("IsActive", 1).All(&DiaryRange)
	return DiaryRange
}

func UpdateDiary(id int64, diary *Diary) (num int64, err error) {
	o := orm.NewOrm()
	d := Diary{ID: id}
	if o.Read(&d) == nil {
		d.Content = diary.Content
		d.Title = diary.Title
		d.Updated = time.Now()
		if num, err := o.Update(&d, "Title", "Content", "Updated"); err == nil {
			return num, nil
		}
	}
	return -1, errors.New("Diary Not Exist")
}

func DeleteDiary(id int64) (err error) {
	return nil
}
