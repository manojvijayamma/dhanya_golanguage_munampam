package controllers

import (
	"myapi/models"	
	"encoding/json"
	
	
)

// Operations about Users
type ReceiptController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.OrderMaster	true		"body for user content"
// @Success 200 {int} models.OrderMaster.Id
// @Failure 403 body is empty
// @router / [post]
func (this *ReceiptController) Post() {
   
	_=this.isAuthUser()
	
	var Order models.OrderMaster

	var FromId string
	var ToId string

	json.Unmarshal(this.Ctx.Input.RequestBody, &Order)
	Order.Type="receipt"

	FromId=Order.ToId
	ToId=Order.FromId

	Order.ToId=ToId
	Order.FromId=FromId
	
	models.AddOrder(Order)

	this.Data["json"] = &map[string]interface{}{"status": "success", "text":"Successfully saved"}
	this.ServeJSON()
	
}



// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *ReceiptController) GetAll() {
	var userId=this.isAuthUser()
	if userId>0 {}
	var limit,_  = this.GetInt64("rowsize")
	var offset,_  = this.GetInt64("page")
	var sortby  = this.GetString("sortField")
	var order  = this.GetString("sortOrder")

	qryparams := map[string]interface{}{
		"branch_id": this.GetString("BranchId"),
		"product_id":  this.GetString("ProductId"),
		"payment_date": this.GetString("PaymentDate"),
		"type":"receipt",
		
	}

	result,pager := models.GetAllOrders(qryparams,sortby,order,offset,limit)
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
func (u *ReceiptController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetOrder(uid)
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
func (u *ReceiptController) Status() {
	uid := u.GetString(":uid")
	models.UpdateOrderStatus(uid)
	u.Data["json"] = uid
	u.ServeJSON()
}


// @Title Priority
// @Description priority the user
// @Param	id		path 	int64	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /priority [get]
func (this *ReceiptController) Priority() {
	id,_ := this.GetInt64("id")
	priority := this.GetString("priority")
	
	models.UpdateOrderPriority(id,priority)
	this.Data["json"] = &map[string]interface{}{"status": "success"}
	this.ServeJSON()
}


// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete/:uid [get]
func (u *ReceiptController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteOrder(uid)
	u.Data["json"] = uid
	u.ServeJSON()
}








