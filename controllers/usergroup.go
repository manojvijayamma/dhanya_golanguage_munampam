package controllers

import (
	"myapi/models"
	
	"encoding/json"
	"fmt"
	"strings"
	


)

// Operations about Users
type UserGroupController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserGroupController) Post() {

	s := strings.SplitN(this.Ctx.Input.Header("Authorization"), " ", 2)
	fmt.Println(s[1])
	var group models.Group
	json.Unmarshal(this.Ctx.Input.RequestBody, &group)
	fmt.Println(group)
	 models.AddGroup(group)
	 this.Data["json"] = &map[string]interface{}{"status": "success", "text":"Successfully saved"}
	this.ServeJSON()
	//u.Data["json"] = map[string]string{"uid": uid}
	//u.ServeJSON()
}



// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *UserGroupController) GetAll() {
	userId:=this.isAuthUser()
	fmt.Println(userId)
	groups := models.GetAllGroups()
	this.Data["json"] = &map[string]interface{}{"result": groups}
	this.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserGroupController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}



// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserGroupController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}


// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
/*
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}*/

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserGroupController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

