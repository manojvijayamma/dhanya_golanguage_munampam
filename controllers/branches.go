package controllers

import (
	"myapi/models"	
	"encoding/json"
	//"fmt"
	"strconv"

)

// Operations about Users
type BranchController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.BranchMaster	true		"body for user content"
// @Success 200 {int} models.BranchMaster.Id
// @Failure 403 body is empty
// @router / [post]
func (this *BranchController) Post() {
   
	var userId=this.isAuthUser()
	var LogMaster models.LogMaster

	var branch models.BranchMaster
	json.Unmarshal(this.Ctx.Input.RequestBody, &branch)
	
	LogMaster.RowId=models.AddBranch(branch)

	LogMaster.UserId=userId
	LogMaster.EndPoint="branch"
	LogMaster.Action="create"
	if branch.Id >0 {
		LogMaster.Action="update"
	}	
	LogMaster.Data=string(this.Ctx.Input.RequestBody)	
	models.ActivityLog(LogMaster)

	this.Data["json"] = &map[string]interface{}{"status": "success", "text":"Successfully saved"}
	this.ServeJSON()
	
}




// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *BranchController) GetAll() {
	var userId=this.isAuthUser()
	if userId>0 {}
	var limit,_  = this.GetInt64("rowsize")
	var offset,_  = this.GetInt64("page")
	var sortby  = this.GetString("sortField")
	var order  = this.GetString("sortOrder")

	qryparams := map[string]interface{}{
		"name": this.GetString("Name"),
		"city":  this.GetString("City"),
		"phone": this.GetString("Phone"),
		"email": this.GetString("Email"),
	}

	result,pager := models.GetAllBranches(qryparams,sortby,order,offset,limit)
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
func (u *BranchController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetBranch(uid)
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
func (u *BranchController) Status() {
	uid := u.GetString(":uid")
	models.UpdateBranchStatus(uid)
	u.Data["json"] = uid
	u.ServeJSON()
}


// @Title Priority
// @Description priority the user
// @Param	id		path 	int64	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /priority [get]
func (this *BranchController) Priority() {
	id,_ := this.GetInt64("id")
	priority := this.GetString("priority")
	
	models.UpdateBranchPriority(id,priority)
	this.Data["json"] = &map[string]interface{}{"status": "success"}
	this.ServeJSON()
}

// @Title GetAll
// @Description get all Branches
// @Success 200 {object} models.Branches
// @router /options [get]
func (this *BranchController) Options() {
	Result := models.GetBranchOptions()
	this.Data["json"] = &map[string]interface{}{"result": Result}
	this.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete/:uid [get]
func (u *BranchController) Delete() {
	uid := u.GetString(":uid")
	count:=models.DeleteBranch(uid)
	if count==0{
		u.Data["json"]=&map[string]interface{}{"status": "failed","message":"Could not delete this data."}
	}else{


		var userId=u.isAuthUser()
		var LogMaster models.LogMaster
		LogMaster.UserId=userId
		LogMaster.EndPoint="branch"
		LogMaster.Action="delete"
		if j, err := strconv.ParseInt(uid, 10, 64); err == nil {
			LogMaster.RowId=j
		}
		
		models.ActivityLog(LogMaster)

		u.Data["json"]=&map[string]interface{}{"status": "success","message":"Successfully Deleted. "}
	}
	
	u.ServeJSON()
}



