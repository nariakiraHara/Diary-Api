package controllers

import (
	"diary-app/models"
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"
)

// SessionController operations for Session
type SessionController struct {
	beego.Controller
}

// Post ...
// @Title Create
// @Description create Session
// @Param	body		body 	models.User	true		"body for Session content"
// @Success 201 {object} models.User
// @Failure 403 body is empty
// @router / [post]
func (s *SessionController) Post() {
	var user models.User
	json.Unmarshal(s.Ctx.Input.RequestBody, &user)
	if !models.AuthenticateUser(&user) {
		s.Data["json"] = errors.New("Invalid User Account")
		return
	}
	s.Data["json"] = user
	s.Ctx.SetCookie("auth", "hoge")
}

// Delete ...
// @Title Delete
// @Description delete the Session
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (s *SessionController) Delete() {

}
