package controllers

import (
	"myapi/models"	
	"encoding/json"
	"strconv"
	

)

// Operations about Users
type ProductController struct {
	BaseController
}


// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *ProductController) GetAll() {
	var userId=this.isAuthUser()
	if userId>0 {}
	var limit,_  = this.GetInt64("rowsize")
	var offset,_  = this.GetInt64("page")
	var sortby  = this.GetString("sortField")
	var order  = this.GetString("sortOrder")

	qryparams := map[string]interface{}{
		"Title": this.GetString("Title"),
		"Branch": this.GetString("Branch"),
		"Category": this.GetString("Category"),		
	}

	result,pager := models.GetAllProduct(qryparams,sortby,order,offset,limit)
	//this.Data["json"] = users
	this.Data["json"] = &map[string]interface{}{"result": result, "pager":pager}
	this.ServeJSON()
}



// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *ProductController) Post() {
   
	var userId=this.isAuthUser()
	var Product models.ProductMaster
	var LogMaster models.LogMaster

	json.Unmarshal(this.Ctx.Input.RequestBody, &Product)
	LogMaster.RowId=models.AddProduct(Product)
	
	LogMaster.UserId=userId
	LogMaster.EndPoint="product"
	LogMaster.Action="create"
	if Product.Id >0 {
		LogMaster.Action="update"
	}	
	LogMaster.Data=string(this.Ctx.Input.RequestBody)	
	models.ActivityLog(LogMaster)

	this.Data["json"] = &map[string]interface{}{"status": "success", "text":"Successfully saved"}
	this.ServeJSON()
	//u.Data["json"] = uid
	//u.ServeJSON()
}






// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *ProductController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetProduct(uid)
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
// @router /delete/:uid [get]
func (u *ProductController) Delete() {
	


	uid := u.GetString(":uid")
	count:=models.DeleteProduct(uid)
	

	if count==0{
		u.Data["json"]=&map[string]interface{}{"status": "failed","message":"Could not delete this data."}
	}else{

		var userId=u.isAuthUser()
		var LogMaster models.LogMaster
		LogMaster.UserId=userId
		LogMaster.EndPoint="ledger"
		LogMaster.Action="delete"
		if j, err := strconv.ParseInt(uid, 10, 64); err == nil {
			LogMaster.RowId=j
		}
		models.ActivityLog(LogMaster)

		u.Data["json"]=&map[string]interface{}{"status": "success","message":"Successfully Deleted. "}
	}
	u.ServeJSON()
}


// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /status/:uid [get]
func (u *ProductController) Status() {
	uid := u.GetString(":uid")
	models.UpdateProductStatus(uid)
	u.Data["json"] = uid
	u.ServeJSON()
}


// @Title Priority
// @Description priority the user
// @Param	id		path 	int64	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /priority [get]
func (this *ProductController) Priority() {
	id,_ := this.GetInt64("id")
	priority := this.GetString("priority")
	
	models.UpdateProductPriority(id,priority)
	this.Data["json"] = &map[string]interface{}{"status": "success"}
	this.ServeJSON()
}


// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /options [get]
func (this *ProductController) Options() {
	id,_ := this.GetInt64("branchId")
	Result := models.GetProductOptions(id)
	this.Data["json"] = &map[string]interface{}{"result": Result}
	this.ServeJSON()
}





