package controllers

import (
	"diary-app/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about Diaries
type DiaryController struct {
	beego.Controller
}

// @Title CreateDiary
// @Description create diary
// @Param	body		body 	models.diary	true		"body for diary content"
// @Success 200 {int} models.Diary.Id
// @Failure 403 body is empty
// @router / [post]
func (d *DiaryController) Post() {
	var diary models.Diary
	json.Unmarshal(d.Ctx.Input.RequestBody, &diary)
	if res, err := models.AddDiary(&diary); res == 0 && err != nil {
		d.Data["json"] = err
	} else {
		d.Data["json"] = diary
	}
	d.ServeJSON()
}

// @Title GetAll
// @Description get all Diaries
// @Success 200 {object} models.Diary
// @router / [get]
func (d *DiaryController) GetAll() {
	diaries := models.GetAllDiary()
	d.Data["json"] = diaries
	d.ServeJSON()
}

// @Title Get
// @Description get diary by id
// @Param id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Diary
// @Failure 403 :id is empty
// @router /:id [get]
func (d *DiaryController) Get() {
	id, _ := strconv.ParseInt(d.GetString(":id"), 10, 64)
	if id != 0 {
		diary := models.Diary{ID: id}
		err := models.GetDiary(&diary)
		if err != nil {
			d.Data["json"] = err.Error()
		} else {
			d.Data["json"] = diary
		}
	}
	d.ServeJSON()
}

// @Title Update
// @Description update the diary
// @Param id		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Diary	true		"body for user content"
// @Success 200 {object} models.Diary
// @Failure 403 :id is not int
// @router /:id [put]
func (d *DiaryController) Put() {
	id, _ := strconv.ParseInt(d.GetString(":id"), 10, 64)
	if id != 0 {
		var diary models.Diary
		json.Unmarshal(d.Ctx.Input.RequestBody, &diary)
		res, err := models.UpdateDiary(id, &diary)
		if err != nil {
			d.Data["json"] = err.Error()
		} else {
			d.Data["json"] = res
		}
	}
	d.ServeJSON()
}

// @Title Delete
// @Description delete the diary
// @Param	id		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (d *DiaryController) Delete() {
	id, _ := strconv.ParseInt(d.GetString(":id"), 10, 64)
	if err := models.DeleteDiary(id); err != nil {
		d.Data["json"] = err
	}
	d.Data["json"] = "delete success!"
	d.ServeJSON()
}
