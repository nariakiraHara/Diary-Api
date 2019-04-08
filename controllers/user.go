package controllers

import (
	"diary-app/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

// Post ...
// @Title Create
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {object} models.User
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if res, err := models.AddUser(&user); res == 0 && err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// Get...
// @Title Get
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (u *UserController) Get() {
	id, _ := strconv.ParseInt(u.GetString(":id"), 10, 64)
	user := models.User{ID: id}
	err := models.GetUserById(&user)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (u *UserController) Put() {
	id, _ := strconv.ParseInt(u.GetString(":id"), 10, 64)
	if id != 0 {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		if err := models.UpdateUserById(&user); err != nil {
			u.Data["json"] = err
			return
		}
		u.Data["json"] = id
	}
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (u *UserController) Delete() {
	id, _ := strconv.ParseInt(u.GetString(":id"), 10, 64)
	if id != 0 {
		if err := models.DeleteUser(id); err != nil {
			u.Data["json"] = err
			return
		}
	}
	u.Data["json"] = id
}
