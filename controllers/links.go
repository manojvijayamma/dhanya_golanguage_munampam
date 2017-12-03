package controllers

import (
	"myapi/models"
	
	"encoding/json"
	"fmt"
	"strings"
	


)

// Operations about Users
type LinksController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *LinksController) Post() {
	
	type selectedItems struct {
		Id    int64 
		Title string 
		ActivityId int64
		Selected bool
	}

	type linkData struct {
		Id    string 
		Title string 
		Activities []selectedItems
	}

	links := linkData{}

	s := strings.SplitN(this.Ctx.Input.Header("Authorization"), " ", 2)
	fmt.Println(s[1])
	
	

	//fmt.Println("CreateContainer,RequestBody:", string(this.Ctx.Input.RequestBody))

	json.Unmarshal(this.Ctx.Input.RequestBody, &links)
	fmt.Println(links)
	 //models.AddLinks(links)
	 this.Data["json"] = &map[string]interface{}{"status": "success", "text":"Successfully saved"}
	this.ServeJSON()
	//u.Data["json"] = map[string]string{"uid": uid}
	//u.ServeJSON()
}



// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *LinksController) GetAll() {
	userId:=this.isAuthUser()
	fmt.Println(userId)
	groups := models.GetAllLinks()
	this.Data["json"] = &map[string]interface{}{"result": groups}
	this.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *LinksController) Get() {
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
func (u *LinksController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}




