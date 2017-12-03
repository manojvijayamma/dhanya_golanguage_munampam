package controllers

import (
	"myapi/models"	
	"encoding/json"
	//"fmt"
	

)

// Operations about Users
type UserController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.BranchMaster	true		"body for user content"
// @Success 200 {int} models.BranchMaster.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserController) Post() {
   
	_=this.isAuthUser()
	
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	
	models.AddUser(user)

	this.Data["json"] = &map[string]interface{}{"status": "success", "text":"Successfully saved"}
	this.ServeJSON()
	
}




// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *UserController) GetAll() {
	var userId=this.isAuthUser()
	if userId>0 {}
	var limit,_  = this.GetInt64("rowsize")
	var offset,_  = this.GetInt64("page")
	var sortby  = this.GetString("sortField")
	var order  = this.GetString("sortOrder")

	qryparams := map[string]interface{}{
		"name": this.GetString("Name"),
		"email":  this.GetString("City"),
		"phone": this.GetString("Phone"),		
	}

	result,pager := models.GetAllUsers(qryparams,sortby,order,offset,limit)
	//this.Data["json"] = users
	this.Data["json"] = &map[string]interface{}{"result": result, "pager":pager}
	this.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = &map[string]interface{}{"result": user}
			//u.Data["json"] = user
		}
	}
	u.ServeJSON()
}



// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /status/:uid [get]
func (u *UserController) Status() {
	uid := u.GetString(":uid")
	models.UpdateUserStatus(uid)
	u.Data["json"] = uid
	u.ServeJSON()
}


// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete/:uid [get]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = uid
	u.ServeJSON()
}

// @Title Priority
// @Description priority the user
// @Param	id		path 	int64	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /priority [get]
func (this *UserController) Priority() {
	id,_ := this.GetInt64("id")
	priority := this.GetString("priority")
	
	models.UpdateUserPriority(id,priority)
	this.Data["json"] = &map[string]interface{}{"status": "success"}
	this.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}


// @Title CreateUser
// @Description create users
// @Param	body		body 	models.BranchMaster	true		"body for user content"
// @Success 200 {int} models.BranchMaster.Id
// @Failure 403 body is empty
// @router /changepassword [post]
func (this *UserController) Changepassword() {
   
	
	var userId=this.isAuthUser()
	var user models.Changepassword
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	user.UserId=userId
	response:=models.ChangePassword(user)

	if response==true {
		this.Data["json"] = &map[string]interface{}{"status": "success", "text":"Successfully Changed."}
	}	else{
		this.Data["json"] = &map[string]interface{}{"status": "failed", "text":"Something Wrong."}
	}	
	this.ServeJSON()
	
}




